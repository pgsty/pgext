## 用法

来源：

- [官方 Rust 包清单](https://github.com/pgbouncer/pg_pgbouncer/blob/e71619b01d391316c98eb8456e5b0b37c76d96d9/Cargo.toml)

`pg_pgbouncer` — 通过 pgrx 扩展的 PostgreSQL 后台工作进程运行和管理外部 PgBouncer 进程。

已复核目录快照记录的版本为 `0.0.0`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_pgbouncer";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_pgbouncer';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
