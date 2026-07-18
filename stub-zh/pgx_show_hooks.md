## 用法

来源：

- [官方扩展控制文件](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_show_hooks/pgx_show_hooks.control)
- [官方上游文档](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/README.md)
- [官方 Rust 包清单](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_show_hooks/Cargo.toml)

`pgx_show_hooks` — 列出 PostgreSQL 钩子函数与 PL/pgSQL 钩子回调在进程内的地址。

已复核目录快照记录的版本为 `0.0.3`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgx_show_hooks";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgx_show_hooks';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
