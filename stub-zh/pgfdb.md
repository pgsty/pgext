## 用法

来源：

- [官方扩展控制文件](https://github.com/fabianlindfors/pgfdb/blob/598aca8937e5c61c805592c90719baaa5e1a4d73/pgfdb.control)

`pgfdb` — 使用 FoundationDB 为 PostgreSQL 提供实验性的分布式、容错、水平扩展存储与事务能力。

已复核目录快照记录的版本为 `0.0.2`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgfdb";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgfdb';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
