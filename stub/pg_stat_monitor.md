

## Usage

> [pg_stat_monitor: query performance monitoring tool for PostgreSQL](https://github.com/percona/pg_stat_monitor)

pg_stat_monitor is an advanced replacement for `pg_stat_statements` that aggregates statistics into configurable time-based buckets, provides query origin information, actual parameter capture, and query plan details.

### Querying Statistics

```sql
-- Basic query monitoring
SELECT application_name, userid::regrole AS user_name,
       datname AS database_name, substr(query, 0, 50) AS query,
       calls, client_ip
FROM pg_stat_monitor;

-- Bucket-based time analysis
SELECT bucket, bucket_start_time, query, calls,
       mean_exec_time, total_exec_time
FROM pg_stat_monitor
ORDER BY total_exec_time DESC;

-- Show query plans
SELECT query, query_plan FROM pg_stat_monitor
WHERE query_plan IS NOT NULL;
```

### Key Features

- **Time-based buckets**: Statistics are grouped into configurable intervals for more accurate analysis
- **Client IP tracking**: Each query records the originating client IP address
- **Actual parameters**: Optionally capture real query parameter values instead of placeholders
- **Query plans**: Each query is accompanied by its actual execution plan
- **Top query tracking**: Identify the most resource-intensive queries per bucket
- **Histogram support**: Visual timing distribution via histogram function

### Functions

```sql
-- Reset all statistics
SELECT pg_stat_monitor_reset();

-- View internal info
SELECT * FROM pg_stat_monitor_info;
```

### Configuration

Key parameters (set in `postgresql.conf`):

| Parameter | Description |
|-----------|-------------|
| `pg_stat_monitor.pgsm_max` | Maximum number of statements tracked |
| `pg_stat_monitor.pgsm_query_max_len` | Maximum query length |
| `pg_stat_monitor.pgsm_bucket_time` | Bucket duration in seconds |
| `pg_stat_monitor.pgsm_max_buckets` | Maximum number of buckets |
| `pg_stat_monitor.pgsm_enable_query_plan` | Enable query plan capture |
| `pg_stat_monitor.pgsm_track` | Track: `top`, `all`, or `none` |
