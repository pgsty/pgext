## Usage

Sources:

- [Official upstream README](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/README)
- [Official extension control file (alohadb_timeseries.control)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_timeseries/alohadb_timeseries.control)
- [Official extension SQL (alohadb_timeseries--1.0.sql)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_timeseries/alohadb_timeseries--1.0.sql)

`alohadb_timeseries` — Auto-partition management for time-series tables using a background worker. Use it for the corresponding scheduling, temporal, or time-series workflow. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION alohadb_timeseries;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `alohadb_time_bucket(bucket_width interval, ts timestamp)` is an extension function and returns `timestamp`.
- `alohadb_time_bucket(bucket_width interval, ts timestamptz)` is an extension function and returns `timestamptz`.
- `alohadb_timeseries_maintain_now()` is an extension function and returns `void`.
- `alohadb_timeseries_manage(p_table_name regclass, p_partition_column text, p_partition_interval interval, p_retention_interval interval DEFAULT NULL)` is an extension function and returns `void`.
- `alohadb_timeseries_status(OUT table_name text, OUT partition_column text, OUT partition_interval interval, OUT retention_interval interval, OUT partition_count int, OUT enabled boolean)` is an extension function and returns `SETOF`.
- `alohadb_timeseries_unmanage(p_table_name regclass)` is an extension function and returns `void`.
- `alohadb_timeseries_config` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
