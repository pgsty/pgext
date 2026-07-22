## Usage

Sources:

- [Official upstream documentation](https://github.com/hailelagi/pg-sweep-stats/blob/12bc23a94f15db0e55298ccb51364637ac959326/README.md)
- [Official Rust implementation](https://github.com/hailelagi/pg-sweep-stats/blob/12bc23a94f15db0e55298ccb51364637ac959326/src/lib.rs)
- [Official Rust package manifest](https://github.com/hailelagi/pg-sweep-stats/blob/12bc23a94f15db0e55298ccb51364637ac959326/Cargo.toml)

`pg_sweep` is an early pgrx extension that exposes PostgreSQL database and table statistics as JSON. It provides on-demand SQL functions only; it does not include the background collector, exporter, or durable time-series store suggested by the broader monitoring idea in its README.

### Core Workflow

The package declares pgrx build features for PostgreSQL 14 through 16. After building and installing the matching binary, create the extension and call its two functions:

```sql
CREATE EXTENSION pg_sweep;

SELECT collect_database_stats();
SELECT collect_table_stats();
```

`collect_database_stats()` returns a JSON object with an epoch timestamp, connection states, current-database transaction counts, block and tuple counters, temporary-file counters, deadlocks, and block I/O timing. `collect_table_stats()` is intended to return a JSON object keyed by `schema.table`, with scan, row-change, live/dead-row, and block-I/O counters.

### Known Source Defect

The reviewed `collect_table_stats()` implementation selects `heap_blks_read`, `heap_blks_hit`, `idx_blks_read`, and `idx_blks_hit` from `pg_stat_user_tables`. On standard PostgreSQL those I/O columns are exposed by `pg_statio_user_tables`, not `pg_stat_user_tables`, so the current function should fail before returning data. Patch and test that query or wait for an upstream correction; do not build dashboards that assume this function works as published.

### Operational Notes

Statistics depend on PostgreSQL collector settings such as `track_activities` and `track_counts`; block I/O timing is meaningful only when `track_io_timing` is enabled. Values are cumulative and can reset after restart or an explicit statistics reset, so consumers must record reset boundaries. The Rust code unwraps SPI results aggressively, making catalog changes and unexpected NULLs failure points. Treat version 0.0.0 as work in progress, validate both functions on the exact server release, and apply least-privilege grants because activity and workload statistics can reveal operational information.
