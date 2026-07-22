## Usage

Sources:

- [Official README](https://github.com/julmon/pg_track_slow_queries/blob/eeea6f7c65d2e59cfe01e71329016e0bc364bdca/README.md)
- [Extension SQL](https://github.com/julmon/pg_track_slow_queries/blob/eeea6f7c65d2e59cfe01e71329016e0bc364bdca/pg_track_slow_queries--1.0.sql)
- [Hook and GUC implementation](https://github.com/julmon/pg_track_slow_queries/blob/eeea6f7c65d2e59cfe01e71329016e0bc364bdca/pg_track_slow_queries.c)

pg_track_slow_queries is a preload extension that writes statements exceeding a duration threshold, their execution metrics, and optionally JSON plans to a dedicated server-side file. SQL functions read and reset that file. Upstream marks the project as still under development, and the checked source is from 2019.

### Core Workflow

Preload the library, configure a positive threshold, restart PostgreSQL, and create the extension in the database used for inspection:

```ini
shared_preload_libraries = 'pg_track_slow_queries'
pg_track_slow_queries.log_min_duration = 1000
pg_track_slow_queries.log_plan = on
pg_track_slow_queries.compression = on
pg_track_slow_queries.max_file_size = 1048576
```

```sql
CREATE EXTENSION pg_track_slow_queries;

SELECT datetime, duration, username, dbname, query, plan
FROM pg_track_slow_queries()
ORDER BY datetime DESC;

SELECT pg_track_slow_queries_reset();
```

The reset call truncates the shared history; export evidence first if it must be retained.

### Important Objects

- `pg_track_slow_queries` returns timestamp, duration, user, application, database, temporary blocks, cache hit ratio, affected tuples, query text, and JSON plan.
- `pg_track_slow_queries_reset` clears the storage file.
- `pg_track_slow_queries.log_min_duration` sets the millisecond threshold; -1 disables capture.
- `pg_track_slow_queries.compression` controls row compression.
- `pg_track_slow_queries.max_file_size` limits file size in kilobytes; -1 is unlimited.
- `pg_track_slow_queries.log_plan` controls plan capture.
- `pg_track_slow_queries.cost_analyze` enables EXPLAIN-ANALYZE-like work above a cost threshold; -1 disables it.

### Security and Performance Notes

Captured SQL and plans can contain literals, identifiers, and sensitive workload details, while the reset function destroys shared evidence. Restrict function execution and server-file access. The extension does not track utility statements or prepared-statement parameter values. Plan capture, compression, file-size enforcement, and especially analyze-style execution can add substantial overhead; benchmark the exact configuration and monitor disk use before enabling it broadly.
