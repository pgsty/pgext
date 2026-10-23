## Usage

Sources:

- [Official source](https://github.com/ExaDev/pg_warren/blob/0a14927e2fa6336439133c386ea7abb27e60a74b/README.md)
- [Official source](https://github.com/ExaDev/pg_warren/blob/0a14927e2fa6336439133c386ea7abb27e60a74b/Cargo.toml)
- [Official source](https://github.com/ExaDev/pg_warren/blob/0a14927e2fa6336439133c386ea7abb27e60a74b/crates/pg_warren/pg_warren.control)
- [Official source](https://github.com/ExaDev/pg_warren/blob/0a14927e2fa6336439133c386ea7abb27e60a74b/sql/pg_warren--0.4.0.sql)
- [Official source](https://github.com/ExaDev/pg_warren/blob/0a14927e2fa6336439133c386ea7abb27e60a74b/docs/limits.md)

`pg_warren` builds versioned adjacency snapshots from ordinary SQL tables for graph traversal and computation. Source revision 0a14927e declares development version 0.4.0. The README's 0.3.0 release reference could not be resolved, so this documentation describes the pinned development source, not a confirmed stable release.

### Core Workflow

The trusted extension supports PostgreSQL 15–18. A role with `CREATE` privilege on the database can install it; no preload is required.

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
```

Traversal reads an immutable version, not current source rows. Refresh after source changes. Building and promoting versions participates in transactions; use `warren.info` and `warren.diff` to inspect state and differences.

### Computation and Memory

`warren.pagerank`, `warren.components` and `warren.shortest_path` cover common algorithms. `warren.spmv`, `warren.gather` and vector helpers compose graph calculations with SQL. Identifier types must agree with the graph; functions without an identifier argument may need an explicit result-type witness.

The cache is per backend: connection pools multiply memory usage. `warren.max_cache_memory`, `warren.max_compute_memory` and `warren.max_build_memory` bound different allocations and are not one global memory cap. `warren.threads` controls compute parallelism. The source is pre-1.0 and does not provide automatic incremental maintenance or a shared cache across backends.

### Access, Backups and Upgrades

Graph source queries execute as the graph owner under the stored `search_path`. Use schema-qualified objects and schemas writable only by trusted roles. Public functions enforce graph access; installation alone does not grant direct access to private catalog tables.

Dump as a superuser to include graph catalogs, versions and vectors; ordinary-role dumps can fail on private tables. Restore required owner/reader roles first and follow the normal restore path, retaining catalog checks. Rebuilding graphs from source tables is a separate recovery option, not equivalent to preserving historical versions. After installing new library files, update extension SQL before calling its functions. Dropping the extension removes its stored graphs.
