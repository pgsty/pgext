## 用法

来源：

- [官方扩展控制文件](https://github.com/xaneets/pg-debyte/blob/main/pg_debyte_ext/pg_debyte_ext.control)
- [官方 Rust 包清单](https://github.com/xaneets/pg-debyte/blob/e0b6e4d4c98781acc0cfbc458fc95de9e0fbb1af/pg_debyte_ext/Cargo.toml)
- [官方上游 README](https://github.com/xaneets/pg-debyte/blob/e0b6e4d4c98781acc0cfbc458fc95de9e0fbb1af/README.md)

`pg_debyte_ext` — 使用 pg_debyte Rust 组件将 bytea 载荷解码为 JSON 的 PostgreSQL 示例扩展。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `15,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_debyte_ext";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_debyte_ext';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
