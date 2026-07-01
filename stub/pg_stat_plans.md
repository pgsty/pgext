


## Usage

Sources: [README](https://github.com/pganalyze/pg_stat_plans/blob/main/README.md), [v2.1.0 release](https://github.com/pganalyze/pg_stat_plans/releases/tag/v2.1.0), [SQL objects](https://github.com/pganalyze/pg_stat_plans/blob/main/pg_stat_plans--2.0.sql)

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

Using `pg_stat_statements` alongside it is recommended so plan IDs can be correlated with query text.

### Query Plans

```sql
SELECT *
FROM pg_stat_plans;
```

The view exposes `userid`, `dbid`, `toplevel`, `queryid`, `planid`, `calls`, `total_exec_time`, and `plan`. To omit stored plan text for lighter queries:

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
