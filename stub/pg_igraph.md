## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_igraph/pg_igraph-1.1.0/README.md)
- [Official extension control file (pg_igraph.control)](https://api.pgxn.org/src/pg_igraph/pg_igraph-1.1.0/pg_igraph.control)
- [Official extension SQL (pg_igraph--1.0.sql)](https://api.pgxn.org/src/pg_igraph/pg_igraph-1.1.0/pg_igraph--1.0.sql)

`pg_igraph` — **High-Performance Graph Traversal Engine for PostgreSQL**. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_igraph;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `graph_add_complex_field(type_id SMALLINT, pos SMALLINT, field_name TEXT)` is an extension function and returns `VOID`.
- `graph_add_complex_type(type_name TEXT)` is an extension function and returns `SMALLINT`.
- `graph_add_edge(from_id BIGINT, to_id BIGINT, rel_name TEXT)` is an extension function and returns `VOID`.
- `graph_add_edge(from_id BIGINT, to_id BIGINT, rel_name TEXT, table_prefix TEXT DEFAULT '')` is an extension function and returns `VOID`.
- `graph_add_edge(from_id INT, to_id INT, rel_name TEXT)` is an extension function and returns `VOID`.
- `graph_add_edge(from_id INT, to_id INT, rel_name TEXT, table_prefix TEXT DEFAULT '')` is an extension function and returns `VOID`.
- `graph_add_node(label_name TEXT)` is an extension function and returns `BIGINT`.
- `graph_add_node(label_name TEXT, table_prefix TEXT DEFAULT '')` is an extension function and returns `BIGINT`.
- `graph_delete_node(node_id BIGINT)` is an extension function and returns `VOID`.
- `graph_delete_node(node_id BIGINT, table_prefix TEXT DEFAULT '')` is an extension function and returns `VOID`.
- `graph_delete_node(node_id INT)` is an extension function and returns `VOID`.
- `graph_delete_node(node_id INT, table_prefix TEXT DEFAULT '')` is an extension function and returns `VOID`.
- `graph_delete_property(node_id BIGINT, prop_name TEXT)` is an extension function and returns `VOID`.
- `graph_delete_property(node_id INT, prop_name TEXT)` is an extension function and returns `VOID`.

### Requirements and Caveats

- The reviewed control file declares default version `1.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
