## 用法

来源：

- [官方扩展控制文件](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_trace_hooks/pgx_trace_hooks.control)
- [官方 Rust 包清单](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_trace_hooks/Cargo.toml)
- [官方上游 README](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/README.md)

`pgx_trace_hooks` — 通过预加载的 pgrx 库跟踪 PostgreSQL 执行器与工具命令钩子的调用

已复核目录快照记录的版本为 `0.0.0`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgx_trace_hooks";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgx_trace_hooks';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
