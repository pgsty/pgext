## 用法

来源：

- [官方 Rust 包清单](https://github.com/askyx/pg_tpch/blob/5684779c89b29b4674604c447a30ec07aaebf307/Cargo.toml)

`pg_tpch` — 在 PostgreSQL 内快速创建、装载 TPC-H 数据集并提供 22 条基准查询元数据。

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_tpch";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_tpch';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
