
## Usage

Sources: [README](https://github.com/ClickHouse/pg_stat_ch/blob/main/README.md), [release 0.3.6](https://github.com/ClickHouse/pg_stat_ch/releases/tag/v0.3.6)

`pg_stat_ch` captures per-query PostgreSQL execution telemetry and exports raw events to ClickHouse through a shared-memory queue and background worker. The upstream project positions it as a raw-event alternative to `pg_stat_statements`, with aggregation and dashboards handled in ClickHouse instead of PostgreSQL.

```sql
CREATE EXTENSION pg_stat_ch;
```

### Required Setup

`pg_stat_ch` must be preloaded and pointed at a ClickHouse database:

```conf
shared_preload_libraries = 'pg_stat_ch'
track_io_timing = on

pg_stat_ch.clickhouse_host = 'localhost'
pg_stat_ch.clickhouse_port = 9000
pg_stat_ch.clickhouse_database = 'pg_stat_ch'
pg_stat_ch.clickhouse_use_tls = on
pg_stat_ch.clickhouse_skip_tls_verify = off
```

The README says PostgreSQL 16, 17, and 18 are fully supported; the latest release is `0.3.6` from April 15, 2026.

### SQL API

- `pg_stat_ch_version()` returns the extension version.
- `pg_stat_ch_stats()` exposes queue and exporter counters.
- `pg_stat_ch_reset()` clears queue counters.
- `pg_stat_ch_flush()` triggers an immediate export flush.

```sql
SELECT pg_stat_ch_version();
SELECT * FROM pg_stat_ch_stats();
SELECT pg_stat_ch_flush();
```

### Important GUCs

- `pg_stat_ch.enabled` toggles collection.
- `pg_stat_ch.queue_capacity` and `pg_stat_ch.string_area_size` size the shared-memory buffers.
- `pg_stat_ch.flush_interval_ms` and `pg_stat_ch.batch_max` control exporter cadence and batch size.
- `pg_stat_ch.log_min_elevel` controls which errors are captured.

### What Gets Captured

- Query timing, row counts, buffer usage, WAL usage, and CPU time.
- DML, DDL, and utility statements.
- SQLSTATE and error levels.
- JIT metrics on PostgreSQL 15+.
- Parallel worker stats on PostgreSQL 18+.
- Application name, client IP, and query text up to the upstream truncation limit.

### Caveats

- The design intentionally drops events on queue overflow instead of blocking the foreground query path.
- ClickHouse schema creation is a required part of setup; upstream quickstart scripts preload it automatically, but manual deployments must load the schema separately.
