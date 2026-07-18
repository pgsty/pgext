## Usage

Sources:

- [Pinned current upstream README](https://github.com/vyruss/pg_statviz/blob/b9af6214a10d9295de54b27842a7e0da1e3066fa/README.md)
- [Version 1.1 installation SQL](https://github.com/vyruss/pg_statviz/blob/b9af6214a10d9295de54b27842a7e0da1e3066fa/pg_statviz--1.1.sql)
- [Pinned extension control file](https://github.com/vyruss/pg_statviz/blob/b9af6214a10d9295de54b27842a7e0da1e3066fa/pg_statviz.control)
- [Formal PGXN distribution](https://pgxn.org/dist/pg_statviz/)

pg_statviz version 1.1 is a pure-SQL statistics snapshot extension paired with an external Python visualization utility. Each snapshot stores selected buffer, configuration, connection, database, I/O, lock, replication, SLRU, wait-event, and WAL counters under the fixed pgstatviz schema.

### Capture snapshots

```sql
CREATE EXTENSION pg_statviz;

SELECT pgstatviz.snapshot();

SELECT snapshot_tstamp, conn_total, conn_active,
       max_query_age_seconds, max_xact_age_seconds
FROM pgstatviz.conn
ORDER BY snapshot_tstamp DESC
LIMIT 10;
```

Schedule snapshot calls with an external job runner at an interval appropriate for the required resolution. The separately installed pg_statviz command reads a time range and renders charts. The current README states Python 3.11+ for the utility and supports recent PostgreSQL versions through 18.

### Retention, privileges, and optional AI

Snapshots accumulate until rows are deleted or delete_snapshots truncates all snapshot tables. Set a retention job and estimate storage before frequent collection. Statistics are cumulative, can reset independently, and may require elevated visibility; interpret deltas rather than treating raw counters as rates.

The install SQL grants the built-in pg_monitor role schema usage, execute on all functions, and SELECT, INSERT, DELETE, and TRUNCATE on all pgstatviz tables. That lets any member collect and remove history, including calling the all-data delete function. Use a dedicated role or revised grants when collection and retention administration should be separated.

The utility's optional AI mode can send chart data, captured configuration, host/role context, and deterministic findings to Claude or Gemini; local Ollama is also supported. This mode is not required for normal visualization. Review data classification, provider retention, prompt content, credentials, and outbound network policy before enabling a cloud provider.
