## Usage

Sources:

- [README](https://github.com/ExaDev/pg_warren/blob/v0.8.0/README.md)
- [Control file](https://github.com/ExaDev/pg_warren/blob/v0.8.0/crates/pg_warren/pg_warren.control)
- [SQL API](https://github.com/ExaDev/pg_warren/blob/v0.8.0/sql/pg_warren--0.8.0.sql)
- [Limits and shared cache](https://github.com/ExaDev/pg_warren/blob/v0.8.0/docs/limits.md)

`pg_warren` 0.8.0 builds versioned adjacency snapshots from ordinary PostgreSQL tables for graph traversal and computation. It supports PostgreSQL 15–18 and keeps graph versions, vectors and their changes within PostgreSQL transactions. Traversal reads a built snapshot; source-table changes require an explicit refresh.

### Core Workflow

This is a trusted extension: a role with `CREATE` privilege on the database can install it. A basic graph needs node and edge queries:

```sql
CREATE EXTENSION pg_warren;
CREATE TABLE people (id bigint PRIMARY KEY, name text);
CREATE TABLE follows (src bigint, dst bigint);
INSERT INTO people VALUES (1, 'Ada'), (2, 'Bob'), (3, 'Cy');
INSERT INTO follows VALUES (1, 2), (2, 3);
SELECT warren.create_graph('social',
  'SELECT id FROM public.people', 'SELECT src, dst FROM public.follows');
SELECT warren.refresh('social');
SELECT * FROM warren.bfs('social', ARRAY[1::bigint], 3);
SELECT warren.pagerank('social', 'rank');
SELECT * FROM warren.vector_topk('social', 'rank', 3, NULL::bigint);
```

The graph stores its source queries. Run `warren.refresh` after changing the source tables; use `warren.info` to inspect versions and `warren.diff` to compare them. Building and promoting versions rolls back with the surrounding transaction.

### Traversal and Computation

| Interface | Purpose |
| --- | --- |
| `warren.bfs`, `warren.shortest_path` | Breadth-first traversal and shortest paths |
| `warren.pagerank`, `warren.components` | Rank nodes and find connected components |
| `warren.leiden` | Detect communities |
| `warren.spmv`, `warren.gather` | Compose graph calculations with SQL |
| `warren.vector_topk` | Read the largest values from a stored vector |
| `warren.shared_cache`, `warren.backends` | Inspect shared entries and backend cache usage |

Identifier types must agree with the graph. Functions without a node identifier argument may require a type witness such as `NULL::bigint`. In 0.8.0, PageRank and Leiden use native kernels; they write their results to named vectors.

### Memory and Preload

PostgreSQL 17–18 initialize the shared cache on first use. PostgreSQL 15–16 need the following setting and a server restart to share graph data between backends:

```ini
shared_preload_libraries = 'pg_warren'
```

Without preload on PostgreSQL 15–16, graph operations still work using private backend caches. Add the library to any existing preload list.

`warren.shared_cache_memory` limits the cluster-wide shared cache and takes effect on configuration reload. `warren.max_cache_memory`, `warren.max_compute_memory` and `warren.max_build_memory` limit distinct private-cache, computation and build allocations; they are not one combined memory budget. `warren.threads` controls computation parallelism. On Linux, shared entries also need sufficient `/dev/shm`; entries that do not fit fall back to private caches.

### Access, Backups and Upgrades

Source queries run as the graph owner with the stored `search_path`. Use schema-qualified source objects and controlled schemas. Public functions enforce graph access; installation does not grant direct access to private catalog tables.

Dump as a superuser to include graph catalogs, versions and vectors, and restore the required owner/reader roles first. Rebuilding from source tables does not preserve historical graph versions. After installing newer library files, run `ALTER EXTENSION pg_warren UPDATE` in each database using the extension. Dropping it removes stored graphs. The project remains pre-1.0 and does not maintain graph snapshots incrementally.
