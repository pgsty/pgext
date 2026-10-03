## Usage

Sources:

- [2.1.1 README](https://github.com/pganalyze/pg_stat_plans/blob/v2.1.1/README.md)
- [2.1.1 changelog](https://github.com/pganalyze/pg_stat_plans/blob/v2.1.1/CHANGELOG.md)
- [SQL2.1 upgrade](https://github.com/pganalyze/pg_stat_plans/blob/v2.1.1/pg_stat_plans--2.0--2.1.sql)
- [Visibility checks](https://github.com/pganalyze/pg_stat_plans/blob/v2.1.1/pg_stat_plans.c)

`pg_stat_plans` tracks aggregate statistics for PostgreSQL plan shapes. It hashes planned query trees into plan IDs, stores example `EXPLAIN` text in shared memory, and helps identify when the same query ID is executed with different plans.

### Enable

`pg_stat_plans` requires PostgreSQL 16 or newer and must be loaded at server start:

```conf
shared_preload_libraries = 'pg_stat_plans'
pg_stat_plans.compress = 'zstd'
```

```sql
CREATE EXTENSION pg_stat_plans;
```

Restart after changing the preload list. Using `pg_stat_statements` alongside it is recommended so plan IDs can be correlated with query text. Optional zstd compression requires a build with zstd support.

### Query Plans

```sql
SELECT *
FROM pg_stat_plans;
```

The view exposes `userid`, `dbid`, `toplevel`, `queryid`, `planid`, `calls`, `total_exec_time`, `plan`, and `plan_advice`. To omit stored plan text for lighter queries:

```sql
SELECT *
FROM pg_stat_plans(false);
```

Group by `queryid` to see multiple plan shapes chosen for one normalized query:

```sql
SELECT queryid, planid, calls, total_exec_time / NULLIF(calls, 0) AS avg_exec_time
FROM pg_stat_plans(false)
ORDER BY queryid, avg_exec_time DESC;
```

### Running Queries

On PostgreSQL 18 and newer, `pg_stat_plans_activity` can show plan IDs and example plans for currently running queries:

```sql
SELECT *
FROM pg_stat_plans_activity;
```

### Reset And Configure

```sql
SELECT pg_stat_plans_reset();
```

Important settings include `pg_stat_plans.max`, `pg_stat_plans.max_size`, `pg_stat_plans.max_plan_memory`, `pg_stat_plans.track`, `pg_stat_plans.compress`, and `pg_stat_plans.plan_advice`.

### Notes

Statistics use PostgreSQL's cumulative statistics system, so counters are flushed at transaction end and may be delayed. Plan IDs describe plan shape and can change when partitions, casts, or expression details change.

### Version, Memory and Visibility

Release 2.1.1 uses SQL extension version 2.1. After installing new files and reloading the library at restart, update existing SQL objects with `ALTER EXTENSION pg_stat_plans UPDATE`. The 2.1 upgrade drops and recreates the function and views; dependent user objects can prevent the update.

`pg_stat_plans.max_plan_memory` limits total plan-text storage. A full memory budget or oversized plan text can leave the text empty while counters remain tracked. `pg_stat_plans.plan_advice` requires PostgreSQL 19 or later and `pg_plan_advice` in the preload list.

Plan text and query IDs for other users require superuser access or `pg_read_all_stats`; otherwise only the caller's own query details are shown. The reset function is not granted to PUBLIC. Collected plan text can contain sensitive constants; restrict monitoring privileges accordingly.
