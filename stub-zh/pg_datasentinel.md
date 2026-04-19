## 用法

来源：[README](https://github.com/datasentinel/pg_datasentinel/blob/main/README.md)，[Release 1.0](https://github.com/datasentinel/pg_datasentinel/releases/tag/1.0)

`pg_datasentinel` 在 PostgreSQL 的 activity、maintenance、temporary-file、checkpoint、wraparound 与 container resource 数据之上提供可观测性视图。由于它会分配共享内存并挂接 activity logging，因此必须预加载。

### 所需设置

```sql
-- postgresql.conf
shared_preload_libraries = 'pg_datasentinel'
log_autovacuum_min_duration = 0
log_temp_files = 0
log_checkpoints = on

CREATE EXTENSION pg_datasentinel;
```

### 主要视图

- `ds_stat_activity`：在 `pg_stat_activity` 基础上扩展 backend RSS、可选 PSS、temp-file bytes，以及 PostgreSQL 18+ 上的 `plan_id`。
- `ds_container_resources`：报告 cgroup CPU 与内存限制，以及当前使用量。
- `ds_wraparound_risk`：基于每小时快照估算到 aggressive vacuum 与 wraparound 的 XID、MXID ETA。
- `ds_xid_snapshots`：`ds_wraparound_risk` 使用的原始快照历史。
- `ds_vacuum_activity`、`ds_analyze_activity`、`ds_tempfile_activity`、`ds_checkpoint_activity`：用于维护与 checkpoint 事件的共享内存 ring buffer。
- `ds_activity_summary`：返回 ring buffer 占用和时间戳的一行汇总。

### 常用 GUCs

- `pg_datasentinel.enabled`：启用或禁用采集。
- `pg_datasentinel.max_entries`：每类 activity stream 的 ring buffer 容量；修改后需要重启。
- `pg_datasentinel.maintenance_force_verbose`：为手工 `VACUUM` 与 `ANALYZE` 添加 `VERBOSE`，以便被捕获。
- `pg_datasentinel.ignore_system_schemas`：抑制 `pg_catalog` 与 `information_schema` 噪声。
- `pg_datasentinel.enable_pss_memory`：读取 `/proc/*/smaps_rollup` 以获取每个 backend 的 PSS。

### 注意事项

- 上游要求 PostgreSQL 15+。
- 内存与 temp-file 增强信息依赖 Linux `/proc`；container 指标依赖 cgroups。
- `VACUUM` 与 `ANALYZE` 的解析依赖英文日志关键字，因此不完全支持已翻译的 server message locale。
- `plan_id` 仅在 PostgreSQL 18+ 上填充，并且与 README 中链接的官方 `pg_store_plans` fork 搭配时最有价值。
