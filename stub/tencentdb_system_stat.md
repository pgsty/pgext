## Usage

Sources:

- [TencentDB for PostgreSQL process monitoring](https://cloud.tencent.com/document/product/409/96894)
- [TencentDB supported extension versions](https://cloud.tencent.com/document/product/409/75121)

`tencentdb_system_stat` 1.0 is a TencentDB for PostgreSQL provider extension for sampled process resource monitoring. It exposes the `tencentdb_process_system_usage` view with backend identity, current query text, CPU usage, and memory bytes. The current provider matrix lists version 1.0 for PostgreSQL 11 through 18 and no version for PostgreSQL 10.

### Inspect sampled backend usage

```sql
CREATE EXTENSION tencentdb_system_stat;

SELECT pid, username, datname, backend_type, query,
       cpu_usage, memory_bytes
FROM tencentdb_process_system_usage;

SHOW tencentdb_system_stat.sampling_interval;
```

`cpu_usage` covers the current sampling period. The `tencentdb_system_stat.sampling_interval` setting is measured in milliseconds and defaults to 1000. `memory_bytes` is reported in bytes. Treat these values as sampled observations rather than cumulative accounting totals.

### Provider and visibility boundary

This extension is supplied and versioned by Tencent Cloud inside TencentDB; it is not a portable source package for self-managed PostgreSQL. Confirm availability on the exact instance engine version before putting the view in monitoring automation, and tolerate provider upgrades, missing samples, and backend-type additions.

The view includes usernames, database names, process identifiers, and current SQL text. Restrict access according to tenant and query-confidentiality requirements, and avoid exporting raw query text into broadly readable telemetry. A shorter sampling interval can increase monitoring work; change it only after measuring overhead on the managed instance and following the provider's parameter controls.
