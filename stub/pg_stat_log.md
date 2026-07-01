


## Usage

Sources: [README](https://github.com/fabriziomello/pg_stat_log/blob/main/README.md), [SQL objects](https://github.com/fabriziomello/pg_stat_log/blob/main/pg_stat_log--0.1.sql), [control file](https://github.com/fabriziomello/pg_stat_log/blob/main/pg_stat_log.control)

`pg_stat_log` collects cumulative statistics about PostgreSQL log messages. It hooks into `emit_log_hook` and counts messages by backend type, database, user, severity, SQLSTATE, and SQLSTATE condition name.

### Enable

`pg_stat_log` requires PostgreSQL 18 or newer and must be preloaded:

```conf
shared_preload_libraries = 'pg_stat_log'
```

Restart PostgreSQL, then create the extension:

```sql
CREATE EXTENSION pg_stat_log;
```

### View Statistics

```sql
SELECT *
FROM pg_stat_log
ORDER BY count DESC;
```

The view exposes `backend_type`, `database_oid`, `database_name`, `user_oid`, `user_name`, `elevel`, `sqlerrcode`, `sqlerrcode_name`, and `count`.

### Common Queries

```sql
SELECT elevel, sqlerrcode, sqlerrcode_name, sum(count) AS total
FROM pg_stat_log
GROUP BY elevel, sqlerrcode, sqlerrcode_name
ORDER BY total DESC
LIMIT 10;

SELECT backend_type, elevel, sqlerrcode_name, count
FROM pg_stat_log
WHERE backend_type <> 'client backend'
ORDER BY count DESC;
```

### Reset And Capacity

```sql
SELECT pg_stat_log_reset();
SELECT * FROM pg_stat_log_info();
```

`pg_stat_log_info()` reports `max_entries`, `num_entries`, `n_dropped`, and `stats_reset`. Increase `pg_stat_log.max_entries` if `n_dropped` grows.

### Configuration

Settings include `pg_stat_log.enabled`, `pg_stat_log.min_error_level`, and `pg_stat_log.max_entries`.

`emit_log_hook` only sees messages that reach the server log. `log_min_messages` therefore acts as a floor for what can be counted.
