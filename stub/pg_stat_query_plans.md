## Usage

Sources:

- [Official documentation](https://github.com/postgredients/pg_stat_query_plans/blob/f0049ce1d02b1912cd7475c9fc91e1e1a441238c/README.MD)
- [Control file](https://github.com/postgredients/pg_stat_query_plans/blob/f0049ce1d02b1912cd7475c9fc91e1e1a441238c/pg_stat_query_plans.control)
- [Version 1.0 SQL](https://github.com/postgredients/pg_stat_query_plans/blob/f0049ce1d02b1912cd7475c9fc91e1e1a441238c/pg_stat_query_plans--1.0.sql)
- [Runtime and access checks](https://github.com/postgredients/pg_stat_query_plans/blob/f0049ce1d02b1912cd7475c9fc91e1e1a441238c/pg_stat_query_plans.c)

`pg_stat_query_plans` 1.0 tracks statements and their execution plans in shared memory. Use its separate query and plan views to identify changes in execution behavior; it is an independent module, not an additional view installed by pg_stat_statements.

### Enablement and Core Workflow

Add the module to the existing preload list, restart PostgreSQL, and create the extension in each database where the views are needed:

```conf
shared_preload_libraries = 'pg_stat_query_plans'
```

```sql
CREATE EXTENSION pg_stat_query_plans;
SELECT queryid, calls, total_exec_time, query
FROM pg_stat_query_plans_sql
ORDER BY total_exec_time DESC
LIMIT 10;
SELECT queryid, planid, calls, normalized_plan
FROM pg_stat_query_plans
ORDER BY calls DESC
LIMIT 10;
SELECT * FROM pg_stat_query_plans_info;
```

### Objects and Configuration

- `pg_stat_query_plans_sql` groups measurements by database, user, query identifier, and top-level status.
- `pg_stat_query_plans` further separates plan identifiers and includes representative query text, normalized plans, and example plans.
- `pg_stat_query_plans_info` exposes allocation, text storage, eviction, and reset information.
- `pg_stat_query_plans.track` selects which statements to collect; `pg_stat_query_plans.track_planning` enables planning-time measurement and is off by default.
- `pg_stat_query_plans_reset(userid, dbid, queryid)` resets matching statistics; zero arguments select all entries. `pg_stat_query_plans_reset_minmax()` clears extrema. The installation SQL revokes public execution of both reset functions.

### Operational Boundaries

Installation requires administrative privileges and a restart for preload changes. The views are selectable by ordinary roles, but the C readers apply caller checks; do not assume every role sees identical rows. Representative query text and plans may retain literals, so review access before exposing them to monitoring users. Storage is bounded and entries can be evicted; these counters are not an audit log. The pinned source does not declare a complete PostgreSQL-major support matrix; validate the intended server version separately.
