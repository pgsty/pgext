## Usage

Sources: [README v2.3.0](https://github.com/saulojb/storage_engine/blob/v2.3.0/README.md), [release v2.3.0](https://github.com/saulojb/storage_engine/releases/tag/v2.3.0), [PGXN 2.3.0](https://pgxn.org/dist/storage_engine/2.3.0/), [current README](https://github.com/saulojb/storage_engine/blob/main/README.md)

`storage_engine` 2.3.0 provides two PostgreSQL table access methods in the `engine` schema:

- `colcompress` for column-oriented compressed storage with vectorized filtering, vectorized aggregation, parallel scans, and stripe/chunk min/max pruning.
- `rowcompress` for row-batch compression with parallel scans, index scans, and batch metadata.

```sql
CREATE EXTENSION storage_engine;
```

### Quick Start

Create tables using either access method. Version 2.2 and later accept per-table options directly in `CREATE TABLE ... WITH (...)`.

```sql
CREATE TABLE events (
  ts timestamptz NOT NULL,
  user_id bigint,
  event_type text,
  value float8
) USING colcompress
  WITH (compression = 'zstd', compression_level = 9, orderby = 'ts ASC');

CREATE TABLE logs (
  id bigserial,
  logged_at timestamptz NOT NULL,
  message text
) USING rowcompress
  WITH (batch_size = 10000, compression = 'zstd');

SELECT event_type, count(*), avg(value)
FROM events
WHERE ts > now() - interval '1 day'
GROUP BY 1;
```

Version 2.3 expands `colcompress` vectorized aggregation with simple `sum(expression)` shapes such as `sum(amount + price)`, post-aggregation arithmetic such as `sum(amount) + count(*)`, and corrected `avg(int8)` behavior in parallel plans.

### Main Tuning Knobs

Session-level GUCs documented upstream include:

- `storage_engine.compression`
- `storage_engine.compression_level`
- `storage_engine.stripe_row_limit`
- `storage_engine.chunk_group_row_limit`
- `storage_engine.enable_parallel_execution`
- `storage_engine.min_parallel_processes`
- `storage_engine.enable_vectorization`
- `storage_engine.enable_vectorized_groupagg`
- `storage_engine.enable_automatic_plan`
- `storage_engine.enable_dml`
- `storage_engine.enable_custom_scan`
- `storage_engine.enable_qual_pushdown`
- `storage_engine.qual_pushdown_correlation_threshold`
- `storage_engine.max_custom_scan_paths`
- `storage_engine.enable_engine_index_scan`
- `storage_engine.enable_column_cache`
- `storage_engine.column_cache_size`
- `storage_engine.debug_vectorized_groupagg_fallback`
- `storage_engine.planner_debug_level`
- `storage_engine.maintenance_auto_enabled`
- `storage_engine.maintenance_auto_naptime`
- `storage_engine.maintenance_auto_database`

The README says these GUCs become visible once the shared library is loaded; add `storage_engine` to `shared_preload_libraries` if you want them available immediately in every session or need the built-in maintenance background worker.

### Types and Operators

`engine.uint8` stores unsigned 64-bit values for `colcompress` workloads that need the full `0` through `2^64 - 1` range. Upstream documents comparison operators (`=`, `<>`, `<`, `<=`, `>`, `>=`), B-tree and hash opclasses, casts to and from `bigint`, `numeric`, and `text`, plus `engine.min`, `engine.max`, and `engine.sum` aggregates. The vectorized planner can dispatch `engine.vmin`, `engine.vmax`, and `engine.vsum` on `colcompress` tables.

### Useful Management Functions

For `colcompress` tables:

```sql
SELECT engine.alter_colcompress_table_set(
  'events'::regclass,
  orderby => 'ts ASC, user_id ASC',
  compression => 'zstd',
  compression_level => 9
);

SELECT engine.colcompress_merge('events');
CALL engine.colcompress_repack('events');
CALL engine.colcompress_repack('events', min_fill_ratio => 0.7);
CALL engine.colcompress_merge_incremental('events', max_stripes => 64);
CALL engine.smart_update(
  'events'::regclass,
  'value = value * 1.1',
  'event_type = ''purchase'''
);
```

Use `engine.colcompress_merge()` after bulk loads when the `orderby` key should be globally sorted for pruning. Use `CALL engine.colcompress_repack()` to compact low-fill stripes, and `CALL engine.colcompress_merge_incremental()` for lower-lock maintenance that processes dirty stripes in batches.

For `rowcompress` tables:

```sql
SELECT engine.alter_rowcompress_table_set(
  'logs'::regclass,
  batch_size => 10000,
  compression => 'zstd',
  compression_level => 5
);

SELECT engine.rowcompress_repack('logs');
CALL engine.rowcompress_merge_incremental('logs', max_batches => 128);
SELECT * FROM engine.rowcompress_scan_stats();
```

Operational views include `engine.colcompress_options`, `engine.colcompress_stripes`, `engine.rowcompress_options`, `engine.rowcompress_batches`, and `engine.storage_health`. `engine.storage_maintenance_recommendation(table)` returns health metrics and a recommended action for one table, and `CALL engine.storage_maintenance_auto(...)` can dispatch maintenance manually or through the built-in background worker.

### When to Use Which AM

- Use `colcompress` for analytical scans, aggregates, and range predicates where projection, vectorization, and stripe/chunk pruning pay off.
- Use `rowcompress` for append-heavy logs or wide rows that are usually fetched together, where compression matters more than column projection.
- For point lookups on `colcompress`, use per-table `index_scan => true` or session-level `storage_engine.enable_engine_index_scan = on`; for analytical range scans, prefer `index_scan => false` with `engine.colcompress_merge()` and an `orderby` key.

### Caveats

- The packaged version in this repo is `2.3.0` for PostgreSQL 15 through 18. Upstream 2.x also tests PostgreSQL 19 devel, but PG19 is not in this repo's package matrix. PostgreSQL 12, 13, and 14 users should stay on upstream 1.3.4.
- The upstream default branch README has moved past the packaged 2.3.0 release; this stub follows `extension.csv` and the v2.3.0 release/PGXN docs.
- Upgrade existing installations with `ALTER EXTENSION storage_engine UPDATE TO '2.3.0';`.
- `colcompress` and `rowcompress` do not support foreign keys or `AFTER ROW` triggers.
- `pg_repack` cannot be used on these table access methods. `engine.colcompress_repack()` acquires `AccessExclusiveLock`, so schedule it during maintenance windows for large tables; the incremental merge procedures are the lower-lock option for dirty stripes or batches.
- `VACUUM FULL`, `CLUSTER`, and `CREATE UNLOGGED TABLE ... USING colcompress` are not supported; upstream recommends the extension's merge/repack functions instead.
- On `colcompress`, combining `orderby` with B-tree indexes can disable the sort-on-write path, and B-tree indexes on ordered columns can defeat stripe pruning for range queries. Use `engine.colcompress_merge()` after loading data when global ordering matters, and prefer `index_scan => false` for analytical tables.
- If `citus` or `pg_cron` is also preloaded, upstream documents the load order as `shared_preload_libraries = 'pg_cron,citus,storage_engine'`; `citus` must appear before `storage_engine`.
