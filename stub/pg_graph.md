## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_graph/pg_graph.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_graph/pg_graph--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_graph/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_graph/pg_graph.c)

`pg_graph` `1.0` provides property-graph traversal and algorithms over ordinary tables. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_graph;
SELECT graph_create('demo');
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

graph_add_vertex_label and graph_add_edge_label bind existing tables and can install foreign-key/check constraints. graph_bfs, graph_pagerank and graph_sample_neighbors expose traversal, ranking and sampling; temporal columns enable as_of selection. Table privileges and RLS apply to graph reads. Graph algorithms load edges into backend memory; pg_graph.max_edges bounds that load. Validate existing data before enabling constraints and review the cost of large graphs.

This module corresponds to the former upstream name `alohadb_graph`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
