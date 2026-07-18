## Usage

Sources:

- [pg_sentence_transformer README at the reviewed commit](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/README.md)
- [pg_sentence_transformer.control at the reviewed commit](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pg_sentence_transformer.control)
- [Version 0.1 install SQL](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pg_sentence_transformer--0.1.sql)
- [Background-worker implementation](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/src/pg_sentence_transformer.c)
- [Python model implementation](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pg_sentence_transformer/model.py)
- [Python dependency metadata](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pyproject.toml)

`pg_sentence_transformer` runs a background worker that turns text from one configured table into sentence embeddings. It must be loaded through `shared_preload_libraries` and depends on `plpython3u`, `vector`, a Python virtual environment, and the model's Python packages.

The install script creates the `sentence_transformer` schema, a queue, source-table triggers, and the `prepare()` and `process()` procedures. The reviewed implementation loads `all-MiniLM-L6-v2`, writes `vector(384)` values, and creates an `ivfflat` index.

### Configuration

```conf
shared_preload_libraries = 'pg_sentence_transformer'
sentence_transformer.venv_path = '/opt/pg_sentence_transformer/venv'
sentence_transformer.database_name = 'appdb'
sentence_transformer.schema_name = 'public'
sentence_transformer.src_table_name = 'posts'
sentence_transformer.src_column = 'body'
sentence_transformer.id_column = 'id'
```

Restart PostgreSQL after setting the postmaster parameters, then connect to the configured database:

```sql
CREATE EXTENSION pg_sentence_transformer CASCADE;
```

Install the Python dependencies into the configured environment first. Prepare the `all-MiniLM-L6-v2` model and NLTK tokenizer data before deploying into a network-restricted environment.

### Caveats

- `prepare()` downloads the NLTK `punkt` and `punkt_tab` datasets, and model initialization may download model files when they are not cached. Perform and verify this provisioning outside ordinary query traffic.
- The current `process()` loop upserts every sentence from one source value under the same `ref_id`; for multi-sentence text, later sentences overwrite earlier embeddings. Treat this as prototype behavior rather than document-level pooling.
- The control file declares version `0.1` and says the extension is relocatable, while the SQL creates a fixed `sentence_transformer` schema. The worker also targets one configured database and table per server start; test schema placement, restart behavior, permissions, Python ABI, and model memory use on the exact PostgreSQL build.
