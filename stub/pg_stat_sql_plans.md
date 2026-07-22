## Usage

Sources:

- [Official extension control file](https://github.com/legrandlegrand/pg_stat_sql_plans/blob/d0401cdf5ff3f01ceb0e0b7a0d8bc3edbb620d24/pg_stat_sql_plans.control)
- [Official upstream documentation](https://github.com/legrandlegrand/pg_stat_sql_plans/blob/d0401cdf5ff3f01ceb0e0b7a0d8bc3edbb620d24/README.md)

`pg_stat_sql_plans` 0.2 is an alpha fork of `pg_stat_statements` that tracks statistics by normalized query ID and plan ID. It can distinguish plan changes for the same SQL shape and retain planning, execution, extension, failure, first-call, and last-call information; upstream explicitly says not to use it in production.

### Core Workflow

For the reviewed branch, use PostgreSQL 14 or later, preload the module, disable core query-ID computation, restart, and create the extension.

```conf
shared_preload_libraries = 'pg_stat_sql_plans'
compute_query_id = off
```

```sql
CREATE EXTENSION pg_stat_sql_plans;

SELECT queryid, planid, calls, plan_time, exec_time, first_call, last_call
FROM pg_stat_sql_plans
ORDER BY exec_time DESC;

SELECT * FROM pg_stat_sql_plans_agg;
```

`pg_stat_sql_plans` stores one row per query-and-plan pair. `pg_stat_sql_plans_agg` rolls rows up by query. `pgssp_backend_qpid(pid)` exposes the most recent query-plan identifier for a backend, `pgssp_normalize_query(text)` normalizes SQL text, and `pg_stat_sql_plans_reset()` clears collected entries.

### Settings and Caveats

Important settings include `pg_stat_sql_plans.max`, `pg_stat_sql_plans.plan_type`, `pg_stat_sql_plans.track`, `pg_stat_sql_plans.track_errors`, `pg_stat_sql_plans.track_pid`, `pg_stat_sql_plans.track_utility`, `pg_stat_sql_plans.save`, and `pg_stat_sql_plans.explain`. Plan hashing and optional plan logging add planning, memory, and log overhead; stored query text can contain literals. Restrict access, measure overhead, and treat IDs as implementation-derived diagnostics rather than portable identities.
