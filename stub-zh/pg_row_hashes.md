## 用法

来源：

- [官方上游文档](https://github.com/ibotty/pg-row-hashes/blob/a654208eba84d767f8d25656ab01d154a86d594c/README.md)
- [官方 Rust 包清单](https://github.com/ibotty/pg-row-hashes/blob/a654208eba84d767f8d25656ab01d154a86d594c/Cargo.toml)

`pg_row_hashes` — 使用 SeaHash 与 FarmHash 变体计算确定性的行指纹、校验值和 XOR 聚合。

已复核目录快照记录的版本为 `0.3.2`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_row_hashes";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_row_hashes';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
