## 用法

来源：

- [官方上游文档](https://github.com/hailelagi/pg-sweep-stats/blob/12bc23a94f15db0e55298ccb51364637ac959326/README.md)
- [官方 Rust 包清单](https://github.com/hailelagi/pg-sweep-stats/blob/12bc23a94f15db0e55298ccb51364637ac959326/Cargo.toml)

`pg_sweep` — 将 PostgreSQL 数据库与表级统计信息收集为 JSON，用于监控。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_sweep";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_sweep';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
