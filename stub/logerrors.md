


## Usage

> [logerrors: Collect statistics about log messages](https://github.com/munakoiso/logerrors)

`logerrors` collects statistics about WARNING, ERROR, and FATAL messages in PostgreSQL log files, making it easy to monitor error rates without parsing logs.

```sql
CREATE EXTENSION logerrors;
```

### Configuration Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `logerrors.interval` | `5000` (5s) | Time between writing stats to buffer (ms, max 60s) |
| `logerrors.intervals_count` | `120` | Number of intervals to keep in buffer (max 360) |
| `logerrors.excluded_errcodes` | (empty) | Error codes to exclude, comma-separated |

### Querying Error Statistics

```sql
SELECT * FROM pg_log_errors_stats();
 time_interval |  type   |       message        | count | username | database | sqlstate
---------------+---------+----------------------+-------+----------+----------+----------
               | WARNING | TOTAL                |     0 |          |          |
               | ERROR   | TOTAL                |     1 |          |          |
               | FATAL   | TOTAL                |     0 |          |          |
             5 | ERROR   | ERRCODE_SYNTAX_ERROR |     1 | postgres | postgres | 42601
           600 | ERROR   | ERRCODE_SYNTAX_ERROR |     1 | postgres | postgres | 42601
```

### Slow Log Statistics

```sql
SELECT * FROM pg_slow_log_stats();
 slow_count |         reset_time
------------+----------------------------
          1 | 2020-06-13 00:19:31.084923
```

### Reset Statistics

```sql
SELECT pg_log_errors_reset();
```
