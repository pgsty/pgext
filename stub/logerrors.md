## Usage

Sources:

- [Official source](https://github.com/munakoiso/logerrors/blob/v2.1.6/README.md)
- [Official source](https://github.com/munakoiso/logerrors/blob/v2.1.6/logerrors.control)
- [Official source](https://github.com/munakoiso/logerrors/releases/tag/v2.1.6)

`logerrors` collects counts of WARNING, ERROR and FATAL messages through server hooks. Release 2.1.6 fixes PostgreSQL 19 ABI/build compatibility while retaining SQL extension version 2.1.

### Enable and inspect

Add the library to the existing preload list, restart PostgreSQL, then create the SQL extension:

```ini
shared_preload_libraries = 'logerrors'
```

```sql
CREATE EXTENSION logerrors;
SELECT * FROM pg_log_errors_stats();
SELECT * FROM pg_slow_log_stats();
SELECT pg_log_errors_reset();
```

`pg_log_errors_stats()` returns collection intervals, message types, counts, role/database names and SQLSTATE. `pg_slow_log_stats()` reports slow-log counts and reset time; `pg_log_errors_reset()` resets statistics.

### Configuration

`logerrors.interval` controls the collection interval in milliseconds: default 5000, maximum 60000. `logerrors.intervals_count` retains 120 intervals by default, at most 360. `logerrors.excluded_errcodes` accepts comma-separated error codes.

These are bounded in-memory statistics, not a durable audit archive. Restrict access to operational information according to database policy. Replacing the library requires a restart because it is preloaded; SQL extversion 2.1 alone does not identify the 2.1.6 runtime fix.
