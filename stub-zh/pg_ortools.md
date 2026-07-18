## 用法

来源：

- [官方上游文档](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ortools/README.md)
- [官方 Rust 包清单](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ortools/Cargo.toml)

`pg_ortools` — 在 PostgreSQL 中调用 HiGHS MIP/LP 求解器进行约束优化，并支持可选异步执行。

已复核目录快照记录的版本为 `0.2.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_ortools";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ortools';
```

该上游项目与 `Matroid` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
