## 用法

来源：

- [官方扩展控制文件](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_rust_ext/conda_pgsql_rust_ext.control)
- [官方 Rust 包清单](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_rust_ext/Cargo.toml)
- [官方上游 README](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/README.md)

`conda_pgsql_rust_ext` — 基于 pgrx 和 rattler_conda_types 的 Conda 版本数据类型及辅助函数。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "conda_pgsql_rust_ext";
SELECT extversion
FROM pg_extension
WHERE extname = 'conda_pgsql_rust_ext';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
