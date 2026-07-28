## Usage

Sources:

- [Official upstream README](https://github.com/neurdb/neurdb/blob/d9aefea3681a8c754ca5836354a314f0dfea23a9/aiengine/pgext/nr_modelmanager/readme.md)
- [Official extension control file (nr_model.control)](https://github.com/neurdb/neurdb/blob/d9aefea3681a8c754ca5836354a314f0dfea23a9/aiengine/pgext/nr_modelmanager/nr_model.control)
- [Official extension SQL (nr_model--1.0.0.sql)](https://github.com/neurdb/neurdb/blob/d9aefea3681a8c754ca5836354a314f0dfea23a9/aiengine/pgext/nr_modelmanager/sql/nr_model--1.0.0.sql)

`nr_model` — pg-model is a simple extension to Postgres that allows model management and in-database inference. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION nr_model;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgm_get_model_id_by_name(model_name text)` is an extension function and returns `INT`.
- `pgm_predict_float4(model_name text, input anyarray)` is an extension function and returns `SETOF`.
- `pgm_predict_table(model_name text, batch_size int, table_name text, column_names text[])` is an extension function and returns `SETOF`.
- `pgm_register_model(model_name text, model_path text)` is an extension function and returns `BOOL`.
- `pgm_store_model(model_name text, model_path text)` is an extension function and returns `BOOL`.
- `pgm_unregister_model(model_name text)` is an extension function and returns `BOOL`.
- `layer` is a table installed or managed by the extension.
- `model` is a table installed or managed by the extension.
- `router` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
