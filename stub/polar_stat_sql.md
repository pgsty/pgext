## Usage

Sources:

- [Official documentation](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/9dfb570c0940ed1db830cc05c2218f3e796c6d4e/external/polar_stat_sql/polar_stat_sql--1.3.sql)
- [Control file](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/9dfb570c0940ed1db830cc05c2218f3e796c6d4e/external/polar_stat_sql/polar_stat_sql.control)
- [Runtime and configuration](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/9dfb570c0940ed1db830cc05c2218f3e796c6d4e/external/polar_stat_sql/polar_stat_sql.c)
- [Upstream test configuration](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/9dfb570c0940ed1db830cc05c2218f3e796c6d4e/external/polar_stat_sql/polar_stat_sql.conf)

`polar_stat_sql` 1.3 collects per-query kernel resource use, plan-node activity, execution phases, lock waits, and storage I/O in the official PolarDB PostgreSQL 11 branch. Its implementation uses PolarDB-specific instrumentation; this is not a stock PostgreSQL compatibility claim.

### Enablement and Core Workflow

On the matching PolarDB build, preload the dependency first, append both modules to the existing list, and restart:

```conf
shared_preload_libraries = 'pg_stat_statements,polar_stat_sql'
polar_stat_sql.enable_stat = on
```

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION polar_stat_sql;
SELECT query, datname, rolname, user_time, system_time, reads, writes
FROM polar_stat_sql
ORDER BY user_time DESC
LIMIT 10;
```

The preload order matters: initialization reads the `pg_stat_statements.max` setting. The SQL view joins the extension's counters with statement text, database names, and role names. `polar_stat_sql()` exposes raw identifiers and counters; `polar_stat_sql_reset()` clears its statistics and has public execution revoked by the installation SQL.

### Configuration and Limits

`polar_stat_sql.sample_rate` controls sampling. `polar_stat_sql.enable_getrusage` enables kernel resource measurements; `polar_stat_sql.enable_gather_plan_info` and `polar_stat_sql.enable_plan_need_time` control plan-node details and timing. `polar_stat_sql.save` controls persistence across shutdowns. Enabling extra measurements adds work to query execution, so choose the required detail deliberately.

Installation and preload changes require administrative access. The collector uses shared state and platform-dependent kernel counters; treat results as monitoring data, not an audit trail. The canonical source and versioned SQL are present in POLARDB_11_STABLE; this module was not found in the currently inspected official 15 and 17 branches. Do not assume that a newer PolarDB or PostgreSQL server provides it.
