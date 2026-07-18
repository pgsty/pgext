## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/s-hironobu/pg_stats/blob/c7753215e4070a88379775470c888d5eed049d17/README.md)
- [Version 1.0 SQL objects](https://github.com/s-hironobu/pg_stats/blob/c7753215e4070a88379775470c888d5eed049d17/pg_stats--1.0.sql)
- [Extension control file](https://github.com/s-hironobu/pg_stats/blob/c7753215e4070a88379775470c888d5eed049d17/pg_stats.control)

`pg_stats` is a pure-SQL collection of convenience monitoring views. It combines PostgreSQL statistics and size functions into `pg_stat_tables`, `pg_stat_indexes`, `pg_stat_users`, and `pg_stat_queries`.

```sql
CREATE EXTENSION pg_stats;
SELECT schemaname, relname, idx_scan_ratio, hit_ratio, total_size
FROM pg_stat_tables
ORDER BY total_size DESC;
```

`pg_stat_tables` adds scan, cache-hit, DML, HOT-update, and size ratios. `pg_stat_indexes` adds index hit ratios and sizes; the user and query views summarize active backends and running-query duration.

These are derived snapshots, not durable measurements. Ratios can be null or misleading after statistics resets, with small denominators, or when activity is not represented by the underlying views. Version 1.0 was written against older statistics columns such as the historical waiting flag. Inspect the SQL on every target PostgreSQL major, qualify view names to avoid confusion with the core `pg_stats` system view, and use a maintained monitoring stack for alerts and long-term trends.
