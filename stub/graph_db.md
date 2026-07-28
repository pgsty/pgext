## Usage

Sources:

- [Official upstream README](https://github.com/dmitriy-m1319/graphextension/blob/e4446ba491752b7ac280f52a32c1533d9e652721/README.md)
- [Official extension control file (graph_db.control)](https://github.com/dmitriy-m1319/graphextension/blob/e4446ba491752b7ac280f52a32c1533d9e652721/extension/graph_db.control)
- [Official extension SQL (graph_db--1.0.sql)](https://github.com/dmitriy-m1319/graphextension/blob/e4446ba491752b7ac280f52a32c1533d9e652721/extension/graph_db--1.0.sql)

`graph_db` — A PostgreSQL extension for graph-database usage. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION graph_db;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `cypher(cstring, cstring)` is an extension function and returns `cstring`.
- `graph_node_in(cstring)` is an extension function and returns `graph_node`.
- `graph_node_out(graph_node)` is an extension function and returns `cstring`.
- `graph_nodes()` is an extension function and returns `graph_node`.
- `key_value(cstring, cstring)` is an extension function and returns `cstring`.
- `graph_node` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
