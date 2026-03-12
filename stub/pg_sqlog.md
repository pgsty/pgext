

## Usage

> [pg_sqlog: access PostgreSQL CSV logs via SQL](https://github.com/kouber/pg_sqlog)

pg_sqlog provides a SQL interface for querying PostgreSQL CSV log files using a foreign table backed by `file_fdw`.

### Querying Logs

```sql
-- Current day's log
SELECT * FROM sqlog.log();

-- Specific day's log
SELECT * FROM sqlog.log('yesterday');
SELECT * FROM sqlog.log('2024-01-15');

-- Error summary
SELECT error_severity, COUNT(*) FROM sqlog.log() GROUP BY 1;
```

### Slow Query Analysis

```sql
-- Top 3 slowest query patterns
SELECT
  AVG(sqlog.duration(message)) AS avg_duration,
  COUNT(*) AS count,
  sqlog.preparable_query(message) AS query_pattern
FROM sqlog.log()
WHERE message ~ '^duration'
GROUP BY 3
ORDER BY 1 DESC
LIMIT 3;

-- Query summary with timing
SELECT
  log_time::time,
  sqlog.duration(message),
  sqlog.summary(message)
FROM sqlog.log('yesterday')
WHERE message ~ '^duration';
```

### Functions

| Function | Description |
|----------|-------------|
| `sqlog.log([timestamp])` | Returns log contents for a given day |
| `sqlog.set_date([timestamp])` | Set the date for `sqlog.log` table queries |
| `sqlog.duration(text)` | Extract query duration from message (ms) |
| `sqlog.preparable_query(text)` | Replace arguments with `?` for grouping |
| `sqlog.summary(text, int, int)` | Strip metadata, show first/last N chars |
| `sqlog.temporary_file_size(text)` | Extract temp file size from message |
| `sqlog.autovacuum([timestamp])` | Autovacuum report for a given day |
| `sqlog.autoanalyze([timestamp])` | Autoanalyze report for a given day |

### Autovacuum Reports

```sql
SELECT * FROM sqlog.autovacuum() LIMIT 5;
SELECT * FROM sqlog.autoanalyze() LIMIT 5;
```

### Prerequisites

Required `postgresql.conf` settings:

```
log_destination = 'syslog,csvlog'
log_filename = 'postgresql.%F'
logging_collector = 'on'
log_rotation_age = '1d'
log_rotation_size = 0
log_truncate_on_rotation = 'on'
log_min_duration_statement = 1000
```
