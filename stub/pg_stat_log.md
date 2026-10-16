## Usage

Sources:

- [README.md](https://github.com/fabriziomello/pg_stat_log/blob/20b57e2e2eff884bfcf3ede3c095762e614c51d1/README.md)
- [pg_stat_log.control](https://github.com/fabriziomello/pg_stat_log/blob/20b57e2e2eff884bfcf3ede3c095762e614c51d1/pg_stat_log.control)
- [pg_stat_log--0.1.sql](https://github.com/fabriziomello/pg_stat_log/blob/20b57e2e2eff884bfcf3ede3c095762e614c51d1/pg_stat_log--0.1.sql)
- [pg_stat_log--0.1--0.3.sql](https://github.com/fabriziomello/pg_stat_log/blob/20b57e2e2eff884bfcf3ede3c095762e614c51d1/pg_stat_log--0.1--0.3.sql)

`pg_stat_log` 0.3 collects PostgreSQL 18+ log-message counts by backend, database, role, severity and SQLSTATE, with optional grouping by query identifier. Add the library to the existing preload list and restart before creating the SQL extension. The default SQL grants allow `pg_read_all_stats` to inspect the counters.

### Core Workflow

```ini
shared_preload_libraries = 'pg_stat_log'
```

```sql
CREATE EXTENSION pg_stat_log;
SELECT database_name, user_name, query_id, elevel, sqlerrcode_name, count
FROM pg_stat_log
ORDER BY count DESC;
SELECT * FROM pg_stat_log_info();
```

### Operational Boundaries

The `pg_stat_log` view resolves names from `pg_stat_log_data()`. `pg_stat_log_info()` reports `max_entries`, `num_entries`, `n_dropped` and `stats_reset`; `pg_stat_log_reset()` clears counters and frees capacity, and is restricted by default. `pg_stat_log.enabled` and `pg_stat_log.min_error_level` control collection, with defaults true and warning. Only messages reaching the server log are counted: `log_min_messages` can filter messages before the hook runs, and caught exceptions are absent. Parallel-worker messages can count the same logical failure more than once.

`pg_stat_log.compute_query_id` defaults to off and requires restart to change; PostgreSQL's `compute_query_id` must not be off. A missing query identifier appears as NULL. `pg_stat_log.max_entries` defaults to 1024, requires restart, and drops new combinations when full; inspect `n_dropped`. Changing this capacity resets all PostgreSQL cumulative statistics on the next start. Counters survive clean restarts, but drop telemetry does not, and crash recovery discards cumulative statistics.

Version 0.3 maps connection authentication, authorization and disconnection to `PSL01`, `PSL02` and `PSL03`. Collection needs the corresponding server logging plus `pg_stat_log.min_error_level = 'log'` or lower; the role attribution is the login role.

Upstream releases 0.1 and 0.2 both install SQL version 0.1. Moving to 0.3 requires matching library files and a restart; the first start resets all PostgreSQL cumulative statistics. Then run `ALTER EXTENSION pg_stat_log UPDATE TO '0.3'` in each installed database. The update recreates the reporting view and function: drop and recreate dependent views and reapply custom grants; the default grants to `pg_read_all_stats` are restored.
