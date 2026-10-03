## Usage

Sources:

- [TimescaleDB v2.30.2 README](https://github.com/timescale/timescaledb/blob/2.30.2/README.md)
- [TimescaleDB 2.29.0 release](https://github.com/timescale/timescaledb/releases/tag/2.29.0)
- [TimescaleDB 2.30.2 release](https://github.com/timescale/timescaledb/releases/tag/2.30.2)
- [Version 2.30.2 changelog](https://github.com/timescale/timescaledb/blob/2.30.2/CHANGELOG.md)
- [TimescaleDB v2.30.2 control file](https://github.com/timescale/timescaledb/blob/2.30.2/timescaledb.control.in)
- [CREATE TABLE API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_table/)
- [create_hypertable() API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_hypertable/)
- [Continuous aggregate API](https://www.tigerdata.com/docs/reference/timescaledb/continuous-aggregates/create_materialized_view/)
- [add_columnstore_policy() API](https://www.tigerdata.com/docs/reference/timescaledb/hypercore/add_columnstore_policy/)
- [TimescaleDB GUCs](https://www.tigerdata.com/docs/reference/timescaledb/configuration/gucs/)

`timescaledb` is a PostgreSQL extension for time-series and event analytics. The current docs emphasize `CREATE TABLE ... WITH (tsdb.hypertable)`, continuous aggregates, automation jobs, and moving chunks into the columnstore.

### Enable the Extension

Append `timescaledb` to the existing preload list and restart PostgreSQL before creating the extension. Version `2.30.2` supports PostgreSQL 16, 17, and 18.

```conf
shared_preload_libraries = 'timescaledb'
```

### Hypertables

```sql
CREATE EXTENSION timescaledb;

CREATE TABLE ts_test (
  ts timestamptz NOT NULL,
  id bigint,
  v integer
) WITH (
  tsdb.hypertable,
  tsdb.orderby = 'ts DESC'
);
```

To convert an existing PostgreSQL table, use the generalized hypertable API:

```sql
CREATE TABLE ts_existing (
  ts timestamptz NOT NULL,
  id bigint,
  v integer
);
SELECT create_hypertable('ts_existing', by_range('ts'));
```

- `CREATE TABLE ... WITH (tsdb.hypertable)` has been documented since TimescaleDB 2.20.0 and is the best-practice path for new hypertables.
- For TimescaleDB 2.23.0 and later, the first `TIMESTAMP` or `TIMESTAMPTZ` column is selected automatically as the partition column unless more than one candidate makes the choice ambiguous.
- `create_hypertable()` still works for converting existing tables.

### Continuous aggregates and jobs

```sql
CREATE MATERIALIZED VIEW ts_hourly
WITH (timescaledb.continuous) AS
SELECT time_bucket('1 hour', ts) AS bucket,
       count(*) AS cnt,
       avg(v)   AS avg_v
FROM ts_test
GROUP BY bucket;

SELECT add_continuous_aggregate_policy(
  'ts_hourly',
  start_offset => INTERVAL '3 hours',
  end_offset => INTERVAL '1 hour',
  schedule_interval => INTERVAL '1 hour'
);

```

- Continuous aggregates require `time_bucket(...)` on the hypertable's time dimension.
- The continuous aggregate `WITH` clause supports `timescaledb.materialized_only`; the current API default is `TRUE`, so real-time aggregation is not enabled unless configured otherwise.
- TimescaleDB 2.28.0 lets manual `refresh_continuous_aggregate()` calls run incrementally in batches. Use `buckets_per_batch`, `max_batches_per_execution`, and `refresh_newest_first` to break large manual refreshes into smaller work units.
- TimescaleDB 2.28.0 also allows adding a new generated aggregate column to an existing continuous aggregate with `ALTER MATERIALIZED VIEW ... ADD COLUMN ... GENERATED ALWAYS AS (...) STORED`; existing rows are `NULL` until refreshed.

### Columnstore

```sql
CREATE TABLE crypto_ticks (
  "time" timestamptz,
  symbol text,
  price double precision,
  day_volume numeric
) WITH (
  tsdb.hypertable,
  tsdb.segmentby = 'symbol',
  tsdb.orderby = 'time DESC'
);

CALL add_columnstore_policy('crypto_ticks', after => INTERVAL '60 days');
```

- `CREATE TABLE ... WITH (tsdb.hypertable)` enables columnstore by default unless `tsdb.columnstore = false`.
- `add_columnstore_policy()` replaces the older `add_compression_policy()` API and requires either `after` or `created_before`, not both.
- Bloom filters are enabled by default for new columnstore chunks. Existing chunks need recompression before they have bloom indexes.

### Relevant GUCs

```sql
SET timescaledb.enable_direct_compress_insert = on;
SET timescaledb.enable_cagg_rewrites = on;
SET timescaledb.enable_columnar_scan_filter_pushdown = on;
```

`timescaledb.enable_direct_compress_insert` and `timescaledb.enable_direct_compress_copy` enable tech-preview direct compression during ingestion. TimescaleDB 2.27.0 adds `timescaledb.enable_cagg_rewrites` and `timescaledb.cagg_rewrites_debug_info`, and documents `timescaledb.enable_columnar_scan_filter_pushdown` as enabled by default.

### Version 2.30.2 and Upgrades

- TimescaleDB 2.29 supports PostgreSQL 16, 17, and 18. PostgreSQL 15 support ended with the 2.28 line, so upgrade PostgreSQL before moving a PG15 database to 2.29.
- Version 2.29.0 adds `compact_chunk()` and a compaction policy for merging small columnstore batches, plus optimized DML chunk exclusion and small-`LIMIT` columnstore scans. Review the release notes before enabling compaction policies on existing workloads.
- The 2.29 line adds `alter_job(..., config_merge => ...)`, direct-compression and unordered-recompression controls, and concurrent refresh policies for hierarchical continuous aggregates.
- The 2.30 line adds `DeferredChunkAppend` for LIMIT queries and concurrent compaction DML. Version 2.30.2 fixes chunk-merge and compressed-query crashes, orphaned compressed chunks after schema removal, and non-deterministic text-collation handling. Granular-refresh option names now use the `timescaledb.cagg_granular_refresh_*` prefix; review any existing granular-refresh configuration.
- The control file marks `timescaledb` trusted and non-relocatable. The server library still has to be preloaded and PostgreSQL restarted according to the packaged deployment configuration.

Install the matching package, then use a fresh `psql -X` connection for the migration so startup commands do not load an older library first. Apply the update in each database and verify its SQL version:

```sql
ALTER EXTENSION timescaledb UPDATE TO '2.30.2';
SELECT extversion FROM pg_extension WHERE extname = 'timescaledb';
```
