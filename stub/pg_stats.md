## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/s-hironobu/pg_stats/blob/c7753215e4070a88379775470c888d5eed049d17/README.md)
- [Version 1.0 SQL objects](https://github.com/s-hironobu/pg_stats/blob/c7753215e4070a88379775470c888d5eed049d17/pg_stats--1.0.sql)
- [Extension control file](https://github.com/s-hironobu/pg_stats/blob/c7753215e4070a88379775470c888d5eed049d17/pg_stats.control)

`pg_stats` is a pure-SQL collection of convenience monitoring views. Version 1.0 creates `pg_stat_tables`, `pg_stat_indexes`, `pg_stat_users`, `pg_stat_queries`, `pg_stat_long_trx`, and `pg_stat_waiting_locks`.

### Core workflow

```sql
CREATE EXTENSION pg_stats;
SELECT schemaname, relname, idx_scan_ratio, hit_ratio, total_size
FROM pg_stat_tables
ORDER BY total_size DESC;
```

### View inventory and limits

`pg_stat_tables` adds scan, cache-hit, DML, HOT-update, and size ratios. `pg_stat_indexes` adds index hit ratios and sizes. `pg_stat_users` summarizes backend login duration, `pg_stat_queries` and `pg_stat_long_trx` show transaction duration, and `pg_stat_waiting_locks` reports ungranted locks.

These are derived snapshots, not durable measurements. Several percentages use integer division before rounding, and `pg_stat_indexes` joins I/O rows by table OID rather than index OID, so a table with multiple indexes can produce duplicated or mismatched index metrics. Version 1.0 also uses the removed `pg_stat_activity.waiting` column. Inspect and correct the SQL for every target PostgreSQL major, qualify view names to avoid confusion with the core `pg_stats` system view, and use a maintained monitoring stack for alerts and long-term trends.
