## 用法

来源：

- [官方扩展控制文件](https://github.com/fraktalio/fmodel-rust-postgres/blob/3f9c55771582e1fc7fd77858b0aded2fcb766cae/fmodel_rust_postgres.control)
- [官方上游文档](https://github.com/fraktalio/fmodel-rust-postgres/blob/3f9c55771582e1fc7fd77858b0aded2fcb766cae/README.md)
- [官方 Rust 包清单](https://github.com/fraktalio/fmodel-rust-postgres/blob/3f9c55771582e1fc7fd77858b0aded2fcb766cae/Cargo.toml)

`fmodel_rust_postgres` — 将 f{model} 事件溯源领域模型作为 pgrx 扩展运行在 PostgreSQL 内。

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "fmodel_rust_postgres";
SELECT extversion
FROM pg_extension
WHERE extname = 'fmodel_rust_postgres';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
