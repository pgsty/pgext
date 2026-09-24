## Usage

Sources:

- [pgGraph v1.2.1 README](https://github.com/Evokoa/pgGraph/blob/v1.2.1/README.md)
- [v1.2.1 release notes](https://github.com/Evokoa/pgGraph/releases/tag/v1.2.1)
- [SQL API Reference](https://github.com/Evokoa/pgGraph/blob/v1.2.1/docs/user_guide/api-reference.mdx)
- [Schema Registration](https://github.com/Evokoa/pgGraph/blob/v1.2.1/docs/user_guide/schema-registration.mdx)
- [Administration and Security](https://github.com/Evokoa/pgGraph/blob/v1.2.1/docs/user_guide/administration-and-security.mdx)
- [Troubleshooting](https://github.com/Evokoa/pgGraph/blob/v1.2.1/docs/user_guide/troubleshooting.mdx)
- [Extension control file](https://github.com/Evokoa/pgGraph/blob/v1.2.1/graph/graph.control)
- [v1.2.0 to v1.2.1 upgrade SQL](https://github.com/Evokoa/pgGraph/blob/v1.2.1/graph/sql/graph--1.2.0--1.2.1.sql)

`pggraph` is the package and PGXN distribution name, but the installed PostgreSQL extension is `graph`. The extension builds derived graph artifacts from ordinary PostgreSQL tables, keeps those tables as the source of truth, and exposes graph search, traversal, shortest path, GQL-style reads, and selected mapped writes through the `graph` schema.

Version 1.2.1 supports PostgreSQL 14-18, named graphs, graph-scoped grants and quotas, durable synchronization, bounded traversal and analytics, maintenance jobs, and selected GQL read/write profiles. It also removes the historical 254-label relationship-type ceiling through a bounded open-vocabulary type dictionary. It does not claim full ISO GQL, full openCypher, or a public SQL/PGQ `GRAPH_TABLE` surface. Standard PostgreSQL SQLSTATEs are paired with stable `PGxxx` details for application diagnostics.

### Basic Graph Build

```sql
CREATE EXTENSION IF NOT EXISTS graph;
SELECT graph.reset();

CREATE TABLE companies (
  id   text PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE people (
  id         text PRIMARY KEY,
  name       text NOT NULL,
  company_id text REFERENCES companies(id)
);

INSERT INTO companies VALUES
  ('c1', 'Acme Bank'),
  ('c2', 'Northwind Trading');

INSERT INTO people VALUES
  ('p1', 'Alice', 'c1'),
  ('p2', 'Bob', 'c1'),
  ('p3', 'Carol', 'c2');

SELECT * FROM graph.auto_discover('public');
SELECT * FROM graph.build();

SELECT node_count, edge_count, edge_types
FROM graph.status();
```

`graph.auto_discover('public')` scans primary keys and foreign keys in the selected schema, registers discovered source tables and edges, and prepares the graph for `graph.build()`. In production schemas, prefer explicit registration so labels, search columns, filter columns, weights, tenant behavior, and graph identity are deliberate.

### Manual and Named-Graph Registration

```sql
SELECT graph.create_graph('customer_360', namespace := 'analytics');
SELECT graph.set_current_graph('customer_360', namespace := 'analytics');

SELECT graph.add_table(
  table_name := 'public.people'::regclass,
  id_column  := 'id',
  columns    := ARRAY['name'],
  tenant_column := NULL
);

SELECT graph.add_table_to_graph(
  graph_name := 'customer_360',
  table_name := 'public.companies'::regclass,
  id_column  := 'id',
  columns    := ARRAY['name'],
  graph_namespace := 'analytics'
);

SELECT graph.add_edge_to_graph(
  graph_name := 'customer_360',
  from_table := 'public.people'::regclass,
  from_column := 'company_id',
  to_table := 'public.companies'::regclass,
  to_column := 'id',
  label := 'works_at',
  bidirectional := true,
  graph_namespace := 'analytics'
);

SELECT * FROM graph.build_graph('customer_360', graph_namespace := 'analytics');
```

Registration applies to the current graph selection unless you use the explicit `*_to_graph` and `*_from_graph` helpers. Node identifiers must match a primary key or a unique `NOT NULL` index. `columns` controls searchable and GQL-visible properties; traversal filter pushdown uses separate `graph.add_filter_column()` registrations. Edge-table and junction-table relationships are also supported, and `label_column` can provide dynamic edge labels within the documented public limit.

### Search, Traversal, and Paths

```sql
SELECT node_table_name, node_id, node
FROM graph.search(
  property_key   := 'name',
  property_value := 'Alice',
  table_filter   := 'public.people'::regclass,
  mode           := 'exact',
  hydrate        := true
);

SELECT depth, node_table_name, node_id, edge_path
FROM graph.traverse(
  'public.people'::regclass,
  'p1',
  2,
  hydrate := false
);

SELECT step, node_table_name, node_id, edge_label
FROM graph.shortest_path(
  'public.people'::regclass,
  'p1',
  'public.companies'::regclass,
  'c1',
  hydrate := false
);
```

With `hydrate := false`, graph functions return compact graph coordinates. With hydration enabled, PostgreSQL source-table ACLs and RLS still govern which source rows are visible. Stale coordinates fail closed rather than fabricating rows.

### Relationship Types and Registration Recovery

Version 1.2.0 permits up to 1,000,000 distinct relationship types in one graph. Each UTF-8 label is limited to 1,024 bytes, the cumulative dictionary is limited to 256 MiB, and relationship-type filter arrays are limited to 4,096 entries and 4 MiB before allocation. `graph.status()` returns only the first 64 committed types as a preview; page the complete effective dictionary in stable ID order with:

```sql
SELECT type_id, label
FROM graph.edge_types(after_type_id := 0, max_rows := 1000);
```

Dynamic labels committed through trigger-backed synchronization are interned by `graph.apply_sync()` without requiring a blanket rebuild. An absent relationship type returns no match, while an ambiguous endpoint mapping fails closed.

The zero-argument reset removes the selected graph's derived engine and artifacts but preserves registrations. Use the boolean overload only to recover from stale relation identities after a table recreation or logical restore:

```sql
SELECT graph.reset();

-- This also clears table, edge, and filter registrations for the selected graph.
SELECT graph.reset(true);
-- Reapply reviewed graph.add_table(...), graph.add_edge(...), and
-- graph.add_filter_column(...) calls before rebuilding.
SELECT * FROM graph.build();
```

Neither form modifies PostgreSQL source tables or other named graphs. `graph.reset(true)` is destructive to the selected graph's registration catalog, so keep the reviewed registration SQL before using it.

### GQL Queries and Relationship Writes

```sql
SELECT row
FROM graph.gql(
  'MATCH (p:people)-[:works_at]->(c:companies)
   WHERE p.name = $name
   RETURN p.id AS person_id, c.name AS company
   ORDER BY company',
  params  := '{"name":"Alice"}'::jsonb,
  hydrate := true
);
```

`graph.gql()` returns one `jsonb` object per SQL row. Node labels map to registered table names and relationship types map to registered edge labels. The supported mutable GQL profile includes registered relationship creation: mapped writes still go through PostgreSQL source-table DML, and source tables remain authoritative. Unsupported openCypher or SQL/PGQ shapes fail with explicit capability errors instead of partial behavior.

### Administration and Operations

```sql
GRANT USAGE, CREATE ON SCHEMA graph TO graph_admin;

SELECT * FROM graph.grant_graph('customer_360', 'app_reader', 'read', namespace := 'analytics');
SELECT * FROM graph.set_graph_quota('owner', 'max_named_graphs', 25, scope_key := 'app_owner');
SELECT * FROM graph.select_graph('customer_360', namespace := 'analytics');
SELECT * FROM graph.add_sync_policy('customer_360', schedule_interval_secs := 300, graph_namespace := 'analytics');
SELECT * FROM graph.run_due_jobs();
SELECT * FROM graph.projection_status();
```

Graph administration covers catalog mutation, builds, sync replay, maintenance, quotas, runtime graph loading, and global analytics. Named graph privileges are `read`, `write`, `build`, and `admin`, but graph `read` is not enough by itself: hydrated reads still require `SELECT` on source tables. A selected graph tenant also scopes traversal, search, GQL, and Cypher calls unless an explicit matching tenant is supplied.

### Upgrading to 1.2.1

Back up PostgreSQL before installing the matching 1.2.1 package. Update each database and rebuild every registered graph. Named graphs must each be selected before their rebuild:

```sql
ALTER EXTENSION graph UPDATE TO '1.2.1';
SELECT * FROM graph.build();
SELECT extversion FROM pg_extension WHERE extname = 'graph';
SELECT * FROM graph.status();
```

This rebuild is mandatory for database-scoped file roots, catalog provenance checks, and transactional generation state. Earlier artifacts are not adopted automatically; source tables and registrations remain authoritative. Logical restores also require rebuilding. The update preserves existing SQL object identities, owners, and explicit grants. In-place binary downgrade is unsupported: restore the pre-upgrade backup with its matching older package, then rebuild derived graph state.

New `graph.sync_retention()` is an administrator diagnostic for the selected graph's sync-log retention and pruning blockers. It neither prunes nor takes the writer lock. Version 1.2.1 also fixes sync replay, transaction rollback, cross-database artifact isolation, and several GQL result-correctness issues.

### Caveats

- Source tables remain the source of truth. Graph artifacts, projection files, sync state, and runtime engines are derived and rebuildable.
- Use `graph.build()` or graph-scoped build helpers after registration changes, and use sync/maintenance APIs when relying on incremental projection state.
- Internal catalog tables such as `graph._graphs`, grants, quotas, jobs, sync logs, and projection metadata are implementation details; use public SQL functions instead.
- Version 1.2.0 uses Rust 1.96 and `cargo-pgrx` 0.19.1 for source builds. PostgreSQL 14 through 18 are supported upstream, with PostgreSQL 17 as the default release-gate target.
