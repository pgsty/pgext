Source: [pgGraph v0.1.7 README](https://github.com/Evokoa/pgGraph/blob/v0.1.7/README.md), [Quickstart](https://github.com/Evokoa/pgGraph/blob/v0.1.7/docs/quickstart.mdx), [SQL API Reference](https://github.com/Evokoa/pgGraph/blob/v0.1.7/docs/user_guide/api-reference.mdx), [Schema Registration](https://github.com/Evokoa/pgGraph/blob/v0.1.7/docs/user_guide/schema-registration.mdx), [Configuration](https://github.com/Evokoa/pgGraph/blob/v0.1.7/docs/user_guide/configuration.mdx).

## Usage

`pggraph` is the package and PGXN distribution name, but the installed PostgreSQL extension is `graph`. The extension builds a derived graph index from ordinary PostgreSQL tables, keeps those tables as the source of truth, and exposes graph search, traversal, GQL-style pattern reads, and path functions through the `graph` schema.

The upstream project labels v0.1.7 as early alpha. Use it first in a disposable or development database and rebuild the graph from source tables instead of treating the generated graph artifact as authoritative data.

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

SELECT node_count, edge_count, edge_types
FROM graph.status();
```

`graph.auto_discover('public')` scans primary keys and foreign keys in the schema, registers the discovered source tables and edges, then builds the graph. For production schemas, prefer explicit registration so labels, search columns, weights, and tenant behavior are intentional.

### Manual Registration

```sql
SELECT graph.reset();

SELECT graph.add_table(
  table_name := 'public.people'::regclass,
  id_column  := 'id',
  columns    := ARRAY['name']
);

SELECT graph.add_table(
  table_name := 'public.companies'::regclass,
  id_column  := 'id',
  columns    := ARRAY['name']
);

SELECT graph.add_edge(
  from_table    := 'public.people'::regclass,
  from_column   := 'company_id',
  to_table      := 'public.companies'::regclass,
  to_column     := 'id',
  label         := 'works_at',
  bidirectional := true
);

SELECT * FROM graph.build();
```

The node identifier must match a primary key or a unique `NOT NULL` index. `columns` controls the source-table properties available to search and GQL. Traversal filter pushdown uses separate `graph.add_filter_column()` registrations.

### Search, Traversal, and Paths

```sql
SELECT node_table_name, node_id, node
FROM graph.search(
  property_key  := 'name',
  property_value := 'Alice',
  table_filter  := 'public.people'::regclass,
  mode          := 'exact',
  hydrate       := true
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

`hydrate := false` returns compact graph coordinates. With hydration enabled, source-row visibility is still governed by PostgreSQL ACLs and RLS, and stale coordinates fail closed rather than fabricating rows.

### GQL Queries

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

`graph.gql()` returns one `jsonb` object per SQL row. Node labels map to registered table names and relationship types map to registered edge labels. The supported GQL/openCypher subset covers common reads, bounded paths, selected aggregates, and narrow mapped writes when mutable overlays are enabled.

### Operational Caveats

- Rebuild with `graph.build()` after changing registrations or after source-table changes that are not covered by the selected sync mode.
- Dynamic edge labels use compact IDs; v0.1.7 supports up to 254 user-facing edge labels.
- Weighted shortest paths require a numeric `weight_column`; missing or NULL weights default to `1`.
- Important GUCs include `graph.max_nodes`, `graph.max_frontier`, `graph.memory_limit_mb`, `graph.query_freshness`, `graph.default_projection_mode`, and `graph.mutable_enabled`.
- Mapped GQL writes require `graph.default_projection_mode = 'mutable_overlay'` and `graph.mutable_enabled = on`.
