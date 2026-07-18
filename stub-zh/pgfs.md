## 用法

来源：

- [官方项目或服务商页面](https://www.coredb.io)

`pgfs` — 基于 pgrx 的 PostgreSQL 扩展，可通过 SQL 在服务器文件系统上复制文件。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgfs";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgfs';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
