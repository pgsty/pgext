


## 用法

来源：[README](https://github.com/datasentinel/pg_datasentinel/blob/main/README.md)，[1.0 发行版](https://github.com/datasentinel/pg_datasentinel/releases/tag/1.0)

`pg_datasentinel` 在 PostgreSQL 的活动、维护、临时文件、检查点、事务回卷与容器资源数据之上提供可观测性视图。由于它会分配共享内存并挂接活动日志记录，因此必须预加载。

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

- `ds_stat_activity`：在 `pg_stat_activity` 基础上扩展后端 RSS、可选 PSS、临时文件字节数，以及 PostgreSQL 18+ 上的 `plan_id`。
- `ds_container_resources`：报告 cgroup CPU 与内存限制，以及当前使用量。
- `ds_wraparound_risk`：基于每小时快照估算距离激进清理与事务回卷的 XID、MXID 预计时间。
- `ds_xid_snapshots`：`ds_wraparound_risk` 使用的原始快照历史。
- `ds_vacuum_activity`、`ds_analyze_activity`、`ds_tempfile_activity`、`ds_checkpoint_activity`：用于维护与检查点事件的共享内存环形缓冲区。
- `ds_activity_summary`：返回环形缓冲区占用和时间戳的一行汇总。

### 常用 GUCs

- `pg_datasentinel.enabled`：启用或禁用采集。
- `pg_datasentinel.max_entries`：每类活动流的环形缓冲区容量；修改后需要重启。
- `pg_datasentinel.maintenance_force_verbose`：为手工 `VACUUM` 与 `ANALYZE` 添加 `VERBOSE`，以便被捕获。
- `pg_datasentinel.ignore_system_schemas`：抑制 `pg_catalog` 与 `information_schema` 噪声。
- `pg_datasentinel.enable_pss_memory`：读取 `/proc/*/smaps_rollup` 以获取每个后端的 PSS。

### 注意事项

- 上游要求 PostgreSQL 15+。
- 内存与临时文件增强信息依赖 Linux `/proc`；容器指标依赖 cgroups。
- VACUUM 与 ANALYZE 的解析依赖英文日志关键字，因此不完全支持已翻译的服务器消息区域设置。
- `plan_id` 仅在 PostgreSQL 18+ 上填充，并且与 README 中链接的官方 `pg_store_plans` 分支搭配时最有价值。
