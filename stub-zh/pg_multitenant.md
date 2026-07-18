## 用法

来源：

- [官方上游文档](https://github.com/gurkanindibay/pg_multitenant/blob/b24ebb692a61f8ae30da2eb7c455edd3cc9c85ac/README.md)
- [官方 Rust 包清单](https://github.com/gurkanindibay/pg_multitenant/blob/b24ebb692a61f8ae30da2eb7c455edd3cc9c85ac/Cargo.toml)

`pg_multitenant` — pg_multitenant：基于 PostgreSQL 行级安全策略的共享 Schema 多租户辅助函数。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_multitenant";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_multitenant';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
