


## Usage

Sources:

- [PGXN metadata](https://api.pgxn.org/dist/pg_stat_ch/0.4.1/META.json)
- [Official README](https://api.pgxn.org/src/pg_stat_ch/pg_stat_ch-0.4.1/README.md)
- [Control file](https://api.pgxn.org/src/pg_stat_ch/pg_stat_ch-0.4.1/pg_stat_ch.control)
- [Installation guide](https://api.pgxn.org/src/pg_stat_ch/pg_stat_ch-0.4.1/INSTALL.md)

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

The README says PostgreSQL 16, 17, and 18 are fully supported; the distribution `0.4.1` keeps the SQL/control version at `0.4`; the PGXN release changes exporter defaults and ClickHouse schema setup.

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

Release `0.3.6` replaces the normalize cache with a queryId-keyed LRU cache and replaces the OpenTelemetry SDK with a direct-protobuf exporter plus sampling support. Existing SQL usage remains centered on the event tables, summary views, and reset/version helper functions above.

### Distribution 0.4.1

The `0.4.1` README raises the default `pg_stat_ch.queue_capacity` to `131072`, sets `flush_interval_ms` to `500`, and caps `batch_max` at the queue capacity. Manual ClickHouse setup now creates the `pg_stat_ch` database and applies the Goose migrations. Queries use the `ts` column for the recent-event window.

### Caveats

- The design intentionally drops events on queue overflow instead of blocking the foreground query path.
- ClickHouse schema creation is a required part of setup; upstream quickstart scripts preload it automatically, but manual deployments must load the schema separately.
- Pigsty's current RPM refresh for `pg_stat_ch` is PostgreSQL 16-18 on EL9/EL10 only; Debian/Ubuntu packages remain PostgreSQL 16-18 on the active targets.
