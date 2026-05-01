## Usage

Source: [README](https://github.com/timescale/timescaledb/blob/main/README.md), [TimescaleDB 2.26.4 release](https://github.com/timescale/timescaledb/releases/tag/2.26.4), [create_hypertable() API](https://docs.timescale.com/api/latest/hypertable/create_hypertable/), [CREATE TABLE hypertable API](https://docs.timescale.com/api/latest/hypertable/create_table/), [continuous aggregates guide](https://docs.timescale.com/use-timescale/latest/continuous-aggregates/create-a-continuous-aggregate/), [add_job() API](https://docs.timescale.com/api/latest/jobs-automation/add_job/), [add_columnstore_policy() API](https://docs.timescale.com/api/latest/hypercore/add_columnstore_policy/)

`timescaledb` is a PostgreSQL extension for time-series and event analytics. The current docs emphasize hypertables, continuous aggregates, automation jobs, and moving older chunks into the columnstore.

### Hypertables

```sql
CREATE EXTENSION timescaledb;

CREATE TABLE ts_test (
  ts timestamptz NOT NULL,
  id bigint,
  v integer
);

SELECT create_hypertable('ts_test', by_range('ts'));
```

- `create_hypertable()` still works, but the API docs mark it as old since TimescaleDB 2.20.0 and point new users toward `CREATE TABLE ... WITH (...)`.
- The current README also shows the newer pattern: `CREATE TABLE ... WITH (tsdb.hypertable)`.

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
- In TimescaleDB 2.13 and later, real-time aggregates are disabled by default unless configured otherwise.

### Columnstore

```sql
ALTER TABLE ts_test SET (
  timescaledb.enable_columnstore,
  timescaledb.orderby = 'ts DESC'
);

CALL add_columnstore_policy('ts_test', after => INTERVAL '1 day');
```

- The docs treat `add_columnstore_policy()` and `convert_to_columnstore()` as the current APIs.
- Older compression functions such as `add_compression_policy()` are documented as old APIs replaced by the columnstore interface.

### Caveats

- TimescaleDB 2.26.4 is a bug-fix release over 2.26.3; the release notes recommend upgrading but do not add a new SQL usage surface for hypertables, continuous aggregates, jobs, or columnstore.
- Relevant 2.26.4 fixes touch continuous aggregate planning and refresh, runtime chunk exclusion, background worker restart after restore, sparse indexes on `orderby`, and compressed chunk merge safety. Keep using the documented APIs above rather than changing SQL for the patch release itself.
