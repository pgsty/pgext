## Usage

Sources:

- [Tagged README](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/README)
- [Extension control file](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/pg_infer.control)
- [Cargo manifest](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/Cargo.toml)
- [Extension installation SQL](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/sql/pg_infer--1.0.0.sql)
- [Model registration implementation](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/src/model_mgmt.rs)
- [Index access method implementation](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/src/am.rs)
- [Official regression SQL](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/sql/pg_infer.sql)
- [Vindex format and model data](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/crates/infer-vindex/README.md)

`pg_infer` is an experimental PostgreSQL extension that exposes transformer-model features, learned associations, and similarity signals through SQL. It creates the `infer` schema for its model registry and provides an `infer` index access method for model-aware ordering over text. Use it to inspect externally prepared model knowledge from PostgreSQL; it does not bundle a model or vindex, and its results are model-dependent signals rather than verified facts.

### Version and Scope

The repository tag is `v0.1.1-alpha`, while the banner in its tagged README still says `v0.1.0-alpha`. In that same tag, `pg_infer.control` declares `default_version = '1.0.0'`, `Cargo.toml` declares `1.0.0`, and `sql/pg_infer--1.0.0.sql` installs SQL version `1.0.0`. Treat the SQL version and project maturity as separate facts: the upstream README still marks the project experimental, warns that the SQL API may change, and says the vindex format is not frozen.

The tagged Cargo manifest has only the `pg18` PostgreSQL feature and selects `pg18` by default. PostgreSQL 18 is therefore the evidenced server target for this source snapshot; do not assume compatibility with earlier major versions.

Install the extension in each database where it is needed, then inspect the installed SQL version:

```sql
CREATE EXTENSION pg_infer;

SELECT extversion
FROM pg_extension
WHERE extname = 'pg_infer';
```

The control file sets `superuser = true` and `relocatable = false`, so a superuser must create the extension and it cannot be moved to another schema after installation.

### Install and Register a Vindex

`pg_infer` stores registration metadata in `infer.models`, but local model data remains in an external vindex directory. Set `infer.data_directory` to the permitted base directory, register an existing vindex, and select the default model used when a query omits its model argument:

```sql
SET infer.data_directory = '/data';
SELECT infer_create_model('qwen05b', '/data/qwen-0.5b.vindex');
SET infer.default_model = 'qwen05b';
SELECT * FROM infer_models();
```

The PostgreSQL operating-system user must be able to read the directory. The registration implementation validates the vindex before inserting or updating `infer.models`; an absolute local path must remain beneath `infer.data_directory`, while a relative path is resolved below that directory. `infer_drop_model` removes the registration and evicts the process-local cache, but it does not delete the external model files.

A vindex is a directory containing model configuration, tokenizer data, embeddings, gate vectors, and feature metadata arranged for query access. The required files depend on the extraction level: browse data supports model inspection such as `walk` and `describe`, while complete forward prediction requires inference-level model data. Prepare and manage this data outside PostgreSQL, and include it explicitly in backup, restore, replication, and host-migration procedures.

### Query Model Knowledge

After setting `infer.default_model`, call the query functions without a model argument, or pass `model => 'name'` to choose one explicitly. These examples are from the tagged API; their rows and scores depend entirely on the registered vindex:

```sql
SELECT * FROM walk('The capital of France is', top => 10);
SELECT * FROM describe('France');
SELECT similar_to('France', 'Paris');
SELECT implies('France', 'Paris');
```

- `walk` traces the strongest feature activations across layers; `infer_explain_walk` adds band and label detail.
- `describe` returns inferred relations for an entity; `describe_layers` preserves the per-layer breakdown, and `nearest_to` probes a selected layer.
- `similar_to` returns a similarity score, `similar_to_many` scores an array of candidates, and `implies` tests directional support in the model's extracted knowledge.
- `infer_show_layers`, `infer_show_features`, and `infer_show_relations` inspect available model metadata; `infer_diff` compares feature metadata between two registered models.
- `infer` performs forward prediction only when the extension was built with the Cargo feature `inference` and the vindex contains the necessary inference data. A default build without that feature exposes the function but returns an error when it is called.

### Build a Semantic Index

Register the model before building an index. The tagged access method accepts one text column and stores the chosen model name in its metapage. Use the distance operator in an ascending top-N query:

```sql
CREATE TABLE documents (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title text NOT NULL
);

CREATE INDEX documents_title_infer_idx
    ON documents USING infer (title)
    WITH (model = 'qwen05b');

SELECT id, title
FROM documents
ORDER BY title <~> 'artificial intelligence'
LIMIT 5;
```

`<~>` returns a distance, so smaller values sort as more similar. `<~` returns the underlying similarity score, where larger values mean stronger similarity, and `@>` exposes directional implication using the default model. The `infer_text_ops` operator class connects `<~>` to the `infer` access method for `ORDER BY ... LIMIT` plans.

### Important Objects

#### Registry and Configuration

- `infer.models` records the model name, vindex path, dimensions, backend, and registration time.
- `infer.default_model` selects the implicit model; `infer.data_directory` constrains local model paths.
- `infer.max_memory` bounds the per-backend vindex cache; `infer.gate_threshold`, `infer.describe_top_k`, and `infer.walk_embed_mode` tune inspection behavior.

#### Query Functions

- Model lifecycle: `infer_create_model`, `infer_drop_model`, `infer_models`.
- Exploration: `walk`, `describe`, `describe_layers`, `nearest_to`, `similar_to`, `similar_to_many`, `implies`.
- Introspection: `infer_explain_walk`, `infer_show_layers`, `infer_show_features`, `infer_show_relations`, `infer_diff`.
- Optional forward prediction: `infer`.

#### Operators and Access Method

- Distance ordering: `<~>` with `USING infer` and `infer_text_ops`.
- Raw similarity filtering: `<~`.
- Directional implication: `@>`.

### Requirements and Caveats

- The tagged Cargo manifest pins `pgrx 0.17.0` and `rust-version = '1.80'`; the tagged README additionally calls for Rust nightly, PostgreSQL 18 or newer, OpenSSL, and OpenBLAS when building from source.
- No `shared_preload_libraries` setting is declared by the control file or tagged setup instructions. Normal use begins with `CREATE EXTENSION pg_infer`; do not add a preload or restart requirement without evidence from the exact build being deployed.
- The registry row and index metapage are not substitutes for the external vindex directory. Verify that model data, permissions, and paths are present on every server that may execute queries after failover or restore.
- The alpha README explicitly warns of breaking SQL changes, an unfrozen vindex format, hardware-specific compute paths, and no established production deployments for that snapshot. Test representative models and queries, constrain memory use, and treat upgrades as compatibility-sensitive.
- Outputs such as relations, similarity, and implication reflect the selected model and extraction pipeline. They should not be treated as authoritative truth or used for consequential decisions without independent validation.
