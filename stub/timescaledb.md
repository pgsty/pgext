


## Usage

Sources: [README](https://github.com/timescale/timescaledb/blob/main/README.md), [TimescaleDB 2.28.0 release](https://github.com/timescale/timescaledb/releases/tag/2.28.0), [2.28.0 changelog](https://github.com/timescale/timescaledb/blob/2.28.0/CHANGELOG.md), [CREATE TABLE API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_table/), [create_hypertable() API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_hypertable/), [continuous aggregate API](https://www.tigerdata.com/docs/reference/timescaledb/continuous-aggregates/create_materialized_view/), [add_columnstore_policy() API](https://www.tigerdata.com/docs/reference/timescaledb/hypercore/add_columnstore_policy/), [GUCs](https://www.tigerdata.com/docs/reference/timescaledb/configuration/gucs/)

`timescaledb` is a PostgreSQL extension for time-series and event analytics. The current docs emphasize `CREATE TABLE ... WITH (tsdb.hypertable)`, continuous aggregates, automation jobs, and moving chunks into the columnstore.

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

SELECT add_job('user_defined_action', '1h');
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

### Caveats

- This project's CSV tracks TimescaleDB `2.28.0` for PostgreSQL 15-18.
- TimescaleDB 2.28.0 speeds up `first(value, time)` and `last(value, time)` aggregates on compressed data by deriving results from columnstore batch metadata instead of decompressing batches.
- The columnar executor in 2.28.0 can evaluate `CASE ... WHEN` expressions on compressed data, keeping conditional aggregates and computed expressions on the vectorized path.
- TimescaleDB 2.28.0 removes adaptive chunking and drops `_timescaledb_catalog.chunk_constraint`, temporarily replacing it with a compatibility view. Use stable informational views instead of depending on that catalog object.
- TimescaleDB 2.28.x is the final minor series supporting PostgreSQL 15; the next planned minor line supports PostgreSQL 16, 17, and 18 only.
