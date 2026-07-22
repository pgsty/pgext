## Usage

Sources:

- [Official C-subproject README](https://github.com/bdrouvot/pg_subxact_counters/blob/main/c/README.md)
- [C extension SQL](https://github.com/bdrouvot/pg_subxact_counters/blob/main/c/pg_subxact_counters--1.0.sql)
- [C counter implementation](https://github.com/bdrouvot/pg_subxact_counters/blob/main/c/pg_subxact_counters.c)

`pg_subxact_counters` 1.0 is a preload-only C extension that exposes four instance-wide atomic counters for subtransaction activity. Sample it periodically to measure rates and detect subtransaction-cache overflow; it does not retain history or identify the responsible database, session, or statement.

### Preload and Setup

Add the library to the preload list, restart PostgreSQL, then create the SQL objects in the monitoring database:

```conf
shared_preload_libraries = 'pg_subxact_counters'
```

```sql
CREATE EXTENSION pg_subxact_counters;

SELECT * FROM pg_subxact_counters;
```

The query function explicitly requires superuser, so the view does as well. Export samples through a tightly controlled monitoring path rather than granting general access.

### Counter Index

- `subxact_start` counts subtransactions started.
- `subxact_commit` counts subtransactions committed.
- `subxact_abort` counts subtransactions aborted.
- `subxact_overflow` counts transitions where a top-level transaction's cached subtransaction IDs overflow.

The overflow counter is not the number of subtransactions beyond the cache limit.

### Interpretation and Caveats

Counters are global across all databases and cumulative from shared-memory initialization until postmaster restart; there is no SQL reset. The rough expression starts minus commits minus aborts can suggest currently open subtransactions, but separate atomic reads and concurrent activity mean one query is not a transactionally consistent snapshot. Monitor deltas and rates, correlate them with workload telemetry, and expect a restart to reset the series. This document covers only the repository's C subproject, which upstream reports testing on PostgreSQL 13 and later.
