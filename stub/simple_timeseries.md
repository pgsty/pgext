## Usage

Sources:

- [Official upstream README](https://github.com/gaoweichang/timeseries-extension/blob/7b3d2ca3c31c2cad61ea09712c17874ecad28d87/README.md)
- [Official extension control file (simple_timeseries.control)](https://github.com/gaoweichang/timeseries-extension/blob/7b3d2ca3c31c2cad61ea09712c17874ecad28d87/project/simple_timeseries.control)
- [Official extension SQL (simple_timeseries--1.0.sql)](https://github.com/gaoweichang/timeseries-extension/blob/7b3d2ca3c31c2cad61ea09712c17874ecad28d87/project/sql/simple_timeseries--1.0.sql)

`simple_timeseries` — time-series extension. Use it for the corresponding scheduling, temporal, or time-series workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION simple_timeseries;

CREATE TABLE sensor_data (
    time TIMESTAMPTZ NOT NULL,
    sensor_id INTEGER,
    temperature DOUBLE PRECISION,
    humidity DOUBLE PRECISION
);

SELECT create_hypertable('sensor_data', 'time', INTERVAL '1 day');
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `apply_retention_policies()` is an extension function and returns `VOID`.
- `check_hypertable_exists(schema_name text, table_name text)` is an extension function and returns `boolean`.
- `compress_chunk(chunk_name REGCLASS)` is an extension function and returns `VOID`.
- `create_continuous_aggregate(view_name TEXT, hypertable REGCLASS, view_sql TEXT, bucket_width INTERVAL, refresh_interval INTERVAL DEFAULT NULL)` is an extension function and returns `VOID`.
- `create_hypertable(table_name REGCLASS, time_column_name TEXT, chunk_time_interval INTERVAL)` is an extension function and returns `VOID`.
- `display_all_chunks()` is an extension function and returns `TABLE`.
- `display_all_triggers()` is an extension function and returns `TABLE`.
- `drop_chunks(hypertable REGCLASS, older_than INTERVAL)` is an extension function and returns `INTEGER`.
- `drop_continuous_aggregate(view_name TEXT)` is an extension function and returns `VOID`.
- `drop_hypertable(table_name REGCLASS)` is an extension function and returns `VOID`.
- `get_chunks_by_hid(p_hypertable_id int)` is an extension function and returns `TABLE`.
- `get_current_timestamp_seconds()` is an extension function and returns `bigint`.
- `get_hypertable_info(s_name text, t_name text)` is an extension function and returns `text`.
- `refresh_continuous_aggregate(view_name TEXT, start_time TIMESTAMPTZ, end_time TIMESTAMPTZ)` is an extension function and returns `VOID`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Install the confirmed extension dependencies first: `dblink`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
