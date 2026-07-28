## Usage

Sources:

- [Official upstream README](https://github.com/lowdown-labs/pg_fela/blob/247a5038f9b88d783524a83a69afaffcf08fdaa7/README.md)
- [Official extension control file (pg_fela.control)](https://github.com/lowdown-labs/pg_fela/blob/247a5038f9b88d783524a83a69afaffcf08fdaa7/pg_fela.control)
- [Official implementation source](https://github.com/lowdown-labs/pg_fela/blob/247a5038f9b88d783524a83a69afaffcf08fdaa7/src/lib.rs)

`pg_fela` — A full Rust pgrx PostgreSQL extension that runs a frozen tabular foundation model (FelaTab) inside the database: classify, impute, cluster, score anomalies, rank feature importance, and explain predictions with a single SELECT, no training step, nothing leaving Postgres. Use it for the corresponding vector, model, or retrieval workflow. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION pg_fela;

-- AutoML in a SELECT: learns from the labeled rows, predicts the ones where target is NULL
SELECT * FROM fela_automl('my_table', 'target_column');

-- Same, plus a per row trust/OOD score so a confident prediction on unfamiliar data gets flagged
SELECT * FROM fela_predict_trust('my_table', 'target_column');

-- Why did row 42 get this prediction? Top contributing features, signed toward/away from it
SELECT * FROM fela_explain_row('my_table', 'target_column', 42);

-- Implicit AutoML: builds my_table_ml, joining prediction/confidence/trust/ood/cluster back onto the base table
SELECT fela_create_view('my_table', 'target_column');
SELECT * FROM my_table_ml WHERE ood;             -- rows unlike anything the model learned from
SELECT * FROM my_table_ml ORDER BY confidence;   -- triage the least sure predictions first
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `fela_anomaly` is an extension function.
- `fela_argmax` is an extension function.
- `fela_automl` is an extension function.
- `fela_caps()` is an extension function.
- `fela_classify` is an extension function.
- `fela_classify_gated` is an extension function.
- `fela_cluster` is an extension function.
- `fela_cluster_ex` is an extension function.
- `fela_confidence` is an extension function.
- `fela_conformal_regress` is an extension function.
- `fela_conformal_threshold` is an extension function.
- `fela_create_view` is an extension function.
- `fela_detect_task` is an extension function.
- `fela_explain` is an extension function.

### Requirements and Caveats

- The catalog records version `1.0.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
