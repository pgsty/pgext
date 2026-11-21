## Usage

Sources:

- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/README.md)
- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/pg_bolt.control)
- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/sql/pg_bolt--0.1.0.sql)
- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/Cargo.toml)
- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/LICENSE)
- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/docs/SQL_API.md)
- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/docs/CONFIGURATION.md)
- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/docs/MVCC_AND_CACHE.md)

`pg_bolt` 0.1.0 is an experimental GPU analytical-query extension over explicitly cached Arrow tables. Its source targets PostgreSQL 15–18 through pgrx 0.19. The upstream README reports host and CUDA testing but says live PostgreSQL pgrx tests were not run on the author's Windows host; treat compatibility as a declared source target, not validated production readiness.

### Load and Query a Snapshot

The default background-worker mode requires preload and a server restart. Preserve other preload entries. Create the extension as a superuser; it creates the non-relocatable `bolt` schema. GPU execution needs an NVIDIA driver and CUDA Toolkit 12 or later. The `cuda-stub` mode is for GPU-less tests.

For an existing table with a region_id column, start with explicit queries and disable automatic planner offload:

```ini
shared_preload_libraries = 'pg_bolt'
pg_bolt.max_vram_mb = 4096
```

```sql
CREATE EXTENSION pg_bolt;
SET pg_bolt.enable_custom_scan = off;
SELECT * FROM bolt.load_table('public.sales', 'sales');
SELECT bolt.explain_tier('SELECT region_id, count(*) FROM sales GROUP BY region_id');
SELECT bolt.query_json('SELECT region_id, count(*) FROM sales GROUP BY region_id');
SELECT * FROM bolt.tables();
SELECT bolt.status();
```

`bolt.load_table` replaces the named cached snapshot, `bolt.cache_table` reuses a clean cached entry, and `bolt.sync_table` refreshes it. `bolt.evict_table` removes a cached table; `bolt.reset_vram` releases runtime/cache state. `bolt.query` streams JSONB rows, whereas `bolt.query_json` returns one JSONB array. `bolt.explain_sql` exposes the full Bolt plan.

### Configuration and Boundaries

`pg_bolt.enabled` controls execution, and `pg_bolt.enable_custom_scan` enables planner integration. `pg_bolt.min_offload_rows` defaults to 50000. `pg_bolt.max_vram_mb` defaults to zero, meaning no configured budget; a positive budget uses LRU eviction and rejects a single table exceeding it. `pg_bolt.statement_timeout_ms` provides cooperative cancellation. `pg_bolt.use_bgworker` defaults on and requires the worker initialized by preload; single-backend development is a separate mode.

Cached Arrow data is a snapshot. The upstream design tracks DML, DDL and transaction invalidation, but still lists cross-backend catalog/snapshot synchronization as follow-up work. Validate freshness, transaction visibility, permissions, RLS and fallback behavior against ordinary PostgreSQL queries before granting application access. The installation SQL does not revoke PUBLIC execution, so an administrator should review and restrict the SQL entry points for the intended roles.

The source is Apache-2.0, while the linked craton-bolt dependency is AGPL-3.0-only; upstream documents different licensing for combined binaries.
