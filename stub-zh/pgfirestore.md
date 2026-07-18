## 用法

来源：

- [官方扩展控制文件](https://github.com/xwkuang5/pgfirestore/blob/bc524d1b5200218b973a157361d74fc55031c907/pgfirestore.control)

`pgfirestore` — 在 PostgreSQL 中实现 Firestore 数据模型、fsvalue 类型和集合/集合组查询语义的 pgrx 扩展。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgfirestore";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgfirestore';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
