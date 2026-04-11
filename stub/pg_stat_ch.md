
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pg_stat_ch;
> SELECT pg_stat_ch_version();
> SELECT * FROM pg_stat_ch_stats();
> ```
>
> Sources: [README](https://github.com/ClickHouse/pg_stat_ch), [Blog post](https://clickhouse.com/blog/pg_stat_ch-postgres-extension-stats-to-clickhouse)

`pg_stat_ch` captures per-query execution telemetry in PostgreSQL and exports raw events to ClickHouse in real time. The upstream project contrasts this with `pg_stat_statements`: instead of aggregating inside PostgreSQL, it sends raw events to ClickHouse for downstream analysis.

## Architecture

The README describes a single pipeline:

```text
PostgreSQL hooks -> shared memory queue -> background worker -> ClickHouse
```

Design goals called out upstream include:

- no network I/O on the query path
- bounded memory via a fixed-size ring buffer
- raw event export instead of local aggregation
- graceful degradation when the queue overflows or ClickHouse is unavailable

## Setup

The extension must be preloaded and configured with ClickHouse connection settings:

```ini
shared_preload_libraries = 'pg_stat_ch'
track_io_timing = on

pg_stat_ch.clickhouse_host = 'localhost'
pg_stat_ch.clickhouse_port = 9000
pg_stat_ch.clickhouse_database = 'pg_stat_ch'
pg_stat_ch.clickhouse_use_tls = on
pg_stat_ch.clickhouse_skip_tls_verify = off
```

After PostgreSQL restart and ClickHouse schema setup:

```sql
CREATE EXTENSION pg_stat_ch;
```

## SQL API

The README documents these SQL functions:

- `pg_stat_ch_version()`
- `pg_stat_ch_stats()`
- `pg_stat_ch_reset()`

`pg_stat_ch_stats()` exposes queue and exporter counters so you can verify that events are being captured and flushed.

## What Gets Captured

The current README lists support for:

- query timing and row counts
- buffer usage and WAL usage
- CPU time
- DML, DDL, and utility statements
- SQLSTATE and error levels
- JIT instrumentation on PostgreSQL 15+
- parallel worker statistics on PostgreSQL 18+
- client context such as application name and client IP

The project currently states full support for PostgreSQL 16, 17, and 18.
