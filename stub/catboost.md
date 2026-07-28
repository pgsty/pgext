## Usage

Sources:

- [Official upstream README](https://github.com/akalend/pg_catboost/blob/c5f10185700ed82b6a35a5c327f644c8bf7259fd/README.md)
- [Official extension control file (catboost.control)](https://github.com/akalend/pg_catboost/blob/c5f10185700ed82b6a35a5c327f644c8bf7259fd/catboost.control)
- [Official extension SQL (catboost--0.1.sql)](https://github.com/akalend/pg_catboost/blob/c5f10185700ed82b6a35a5c327f644c8bf7259fd/catboost--0.1.sql)

`catboost` — machine learning module using CatBoost. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION catboost;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ml_learn(name text, model_type int, options text, table_name text, filename text)` is an extension function and returns `float`.
- `ml_meta(OUT name text, OUT loss_function text, OUT model_type char(1), OUT acc real, OUT args text, OUT classes text)` is an extension function and returns `setof`.
- `ml_predict(model text, tablename text, join_field text DEFAULT 'row', OUT index text, OUT predict float, OUT class text)` is an extension function and returns `setof`.
- `ml_predict_internal(model text, tablename text, join_field text DEFAULT 'row', isQuery bool DEFAULT FALSE, OUT index text, OUT predict float, OUT class text)` is an extension function and returns `setof`.
- `ml_predict_query(model text, query text, join_field text DEFAULT 'row', OUT index text, OUT predict float, OUT class text)` is an extension function and returns `setof`.
- `ml_test(name Name)` is an extension function and returns `text`.
- `ml_model` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- The control file marks the extension as trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
