## 用法

来源：

- [官方扩展控制文件](https://github.com/stella3d/pgzstd_rs/blob/d9b56279f41f1375e9db6a3d397c89346dade5f8/pgzstd_rs.control)
- [官方上游文档](https://github.com/stella3d/pgzstd_rs/blob/d9b56279f41f1375e9db6a3d397c89346dade5f8/README.md)
- [官方 Rust 包清单](https://github.com/stella3d/pgzstd_rs/blob/d9b56279f41f1375e9db6a3d397c89346dade5f8/Cargo.toml)

`pgzstd_rs` — 在 PostgreSQL 中使用 Zstandard 压缩和解压 bytea，并提供并行批量压缩函数

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgzstd_rs";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgzstd_rs';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
