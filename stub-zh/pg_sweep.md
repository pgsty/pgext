## 用法

来源：

- [官方上游文档](https://github.com/hailelagi/pg-sweep-stats/blob/12bc23a94f15db0e55298ccb51364637ac959326/README.md)
- [官方 Rust 实现](https://github.com/hailelagi/pg-sweep-stats/blob/12bc23a94f15db0e55298ccb51364637ac959326/src/lib.rs)
- [官方 Rust 软件包清单](https://github.com/hailelagi/pg-sweep-stats/blob/12bc23a94f15db0e55298ccb51364637ac959326/Cargo.toml)

`pg_sweep` 是一个早期 pgrx 扩展，以 JSON 形式暴露 PostgreSQL 数据库和表统计信息。它只提供按需 SQL 函数；README 中更广泛的监控设想所暗示的后台采集器、导出器或持久时序存储并不存在。

### 核心流程

软件包声明了面向 PostgreSQL 14 至 16 的 pgrx 构建特性。构建并安装与服务器匹配的二进制文件后，创建扩展并调用两个函数：

```sql
CREATE EXTENSION pg_sweep;

SELECT collect_database_stats();
SELECT collect_table_stats();
```

`collect_database_stats()` 返回 JSON 对象，其中包含 epoch 时间戳、连接状态、当前数据库事务计数、块与元组计数器、临时文件计数器、死锁和块 I/O 计时。`collect_table_stats()` 预期返回以 `schema.table` 为键的 JSON 对象，包含扫描、行变更、活跃或死亡行以及块 I/O 计数器。

### 已知源码缺陷

已复核的 `collect_table_stats()` 实现从 `pg_stat_user_tables` 查询 `heap_blks_read`、`heap_blks_hit`、`idx_blks_read` 和 `idx_blks_hit`。在标准 PostgreSQL 上，这些 I/O 列由 `pg_statio_user_tables` 而不是 `pg_stat_user_tables` 暴露，因此当前函数应当会在返回数据前失败。应修补并测试该查询，或等待上游修正；不要构建假设公开函数可以工作的仪表盘。

### 运维说明

统计信息依赖 `track_activities`、`track_counts` 等 PostgreSQL 收集器设置；只有启用 `track_io_timing` 后，块 I/O 计时才有意义。这些值是累积值，服务器重启或显式重置统计后可能归零，因此使用方必须记录重置边界。Rust 代码会直接解包 SPI 结果，目录变化和意外空值都可能成为失败点。应将 0.0.0 版视为开发中状态，在准确的服务器发布上验证两个函数，并应用最小权限授权，因为活动与工作负载统计可能暴露运维信息。
