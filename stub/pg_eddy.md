## Usage

Sources:

- [Official upstream README](https://github.com/trickle-labs/pg-eddy/blob/5afc17c9b630788dffeee5a6a603a6df99f22183/README.md)
- [Official extension control file (pg_eddy.control)](https://github.com/trickle-labs/pg-eddy/blob/5afc17c9b630788dffeee5a6a603a6df99f22183/pg_eddy/pg_eddy.control)
- [Official extension SQL (pg_eddy--0.1.0.sql)](https://github.com/trickle-labs/pg-eddy/blob/5afc17c9b630788dffeee5a6a603a6df99f22183/pg_eddy/sql/pg_eddy--0.1.0.sql)

`pg_eddy` — A Postgres extension for labeled property graphs with index-free adjacency and built-in materialized views. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_eddy;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_eddy_edge_handler(internal)` is an extension function and returns `table_am_handler`.
- `pg_eddy_node_handler(internal)` is an extension function and returns `table_am_handler`.
- `edges` is an extension-defined view.
- `nodes` is an extension-defined view.
- `_pg_eddy.edge_type_dst` is a table installed or managed by the extension.
- `_pg_eddy.edge_type_src` is a table installed or managed by the extension.
- `_pg_eddy.edges` is a table installed or managed by the extension.
- `_pg_eddy.label_index` is a table installed or managed by the extension.
- `_pg_eddy.label_registry` is a table installed or managed by the extension.
- `_pg_eddy.nodes` is a table installed or managed by the extension.
- `_pg_eddy.property_key_registry` is a table installed or managed by the extension.
- `_pg_eddy.rel_type_registry` is a table installed or managed by the extension.
- `edges` is a table installed or managed by the extension.
- `nodes` is a table installed or managed by the extension.

### Requirements and Caveats

- The catalog records version `0.6.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
