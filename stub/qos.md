## Usage

Sources:

- [README.md](https://github.com/appstonia/pg_qos/blob/fd3462b7fa81f8bb8aed2113a1f75bcf0e5dfe00/README.md)
- [qos.control](https://github.com/appstonia/pg_qos/blob/fd3462b7fa81f8bb8aed2113a1f75bcf0e5dfe00/qos.control)
- [qos--1.0--1.1.sql](https://github.com/appstonia/pg_qos/blob/fd3462b7fa81f8bb8aed2113a1f75bcf0e5dfe00/qos--1.0--1.1.sql)
- [qos--1.1.sql](https://github.com/appstonia/pg_qos/blob/fd3462b7fa81f8bb8aed2113a1f75bcf0e5dfe00/qos--1.1.sql)

`qos` 1.1 (distribution 1.1.0) applies per-role and per-database resource limits on PostgreSQL 15+. Merge `qos` into `shared_preload_libraries` and restart before creating its SQL objects as an administrator. CPU affinity limits require Linux.

### Configure Limits

```sql
CREATE EXTENSION qos;
ALTER ROLE app_user SET qos.work_mem_limit = '32MB';
ALTER ROLE app_user SET qos.max_concurrent_select = '100';
ALTER ROLE app_user SET qos.max_select_rate = '10/500ms';
SELECT * FROM qos_stat_rate;
```

### Limit Semantics

`qos.work_mem_limit` caps effective work memory; `qos.cpu_core_limit` controls CPU affinity on Linux and limits parallel workers on other platforms. `qos.max_concurrent_tx`, `qos.max_concurrent_select`, `qos.max_concurrent_update`, `qos.max_concurrent_delete` and `qos.max_concurrent_insert` cap concurrent operations.

`qos.max_tx_rate`, `qos.max_select_rate`, `qos.max_update_rate`, `qos.max_delete_rate` and `qos.max_insert_rate` use count/window pairs such as 100/1s. The default -1 disables each rate limit. Windows range from 100 ms to one day. Rate and concurrency violations raise SQLSTATE 54000; clients should use the retry hint. The most restrictive applicable role/database setting wins, and rate pairs are compared by normalized rate.

### Observability and Upgrade

`qos_stat_rate` exposes live windows; the other `qos_stat` views expose activity and counters. `qos_prometheus_metrics()` renders Prometheus exposition text. Counters reset at server restart. Version 1.1 replaces the old nonfunctional `qos_get_stats()` with these views.

Upgrading requires replacing the library and restarting PostgreSQL because the shared-memory layout changes, followed by `ALTER EXTENSION qos UPDATE TO '1.1'` in each database. New rate limits stay disabled until configured. These controls do not replace application admission limits or operating-system isolation.
