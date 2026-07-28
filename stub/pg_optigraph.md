## Usage

Sources:

- [Official upstream README](https://github.com/cloverdb/pg_optigraph/blob/83e362f68774c0c87fc2de1508795011c4170dfe/README.md)
- [Official extension control file (pg_optigraph.control)](https://github.com/cloverdb/pg_optigraph/blob/83e362f68774c0c87fc2de1508795011c4170dfe/extension/pg_optigraph.control)
- [Official extension SQL (pg_optigraph--0.1.0.sql)](https://github.com/cloverdb/pg_optigraph/blob/83e362f68774c0c87fc2de1508795011c4170dfe/extension/sql/pg_optigraph--0.1.0.sql)

`pg_optigraph` — pg Optigraph is a PostgreSQL extension which uses OptiGraph ML model. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_optigraph;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `optigraph.extract_plans(query_text TEXT)` is an extension function and returns `TABLE`.
- `optigraph.health_check()` is an extension function and returns `TABLE`.
- `optigraph.reset_stats()` is an extension function and returns `VOID`.
- `optigraph.stats()` is an extension function and returns `TABLE`.
- `optigraph.status()` is an extension function and returns `TABLE`.
- `optigraph.test_optimize()` is an extension function and returns `TABLE`.
- `optigraph.configuration` is an extension-defined view.
- `optigraph` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
