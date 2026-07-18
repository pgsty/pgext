## Usage

Sources:

- [Official upstream documentation](https://github.com/AyoubKaz07/pg_microbench/blob/b33b11fc47da522b26dde2bf3023c2b48ee1374a/README.md)

`pg_microbench` — Perf stat-like hardware performance counters for PostgreSQL planner, executor, utility, and SPI query phases using Linux perf_event_open

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_microbench";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_microbench';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
