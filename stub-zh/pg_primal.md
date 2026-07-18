## 用法

来源：

- [官方扩展控制文件](https://github.com/PrimalHQ/primal-server/blob/4b6f384c00fb68d7df10a38bdc33e0e951d2d446/pg_primal/pg_primal.control)
- [官方上游 README](https://github.com/PrimalHQ/primal-server/blob/4b6f384c00fb68d7df10a38bdc33e0e951d2d446/README.md)

`pg_primal` — pg_primal：Primal Nostr 缓存服务器使用的 PostgreSQL 数据库扩展。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_primal";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_primal';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
