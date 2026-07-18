## 用法

来源：

- [官方 Rust 包清单](https://github.com/dotvezz/pg_smolid/blob/b9840b6764545c4e59353506cfa4bcf2cfb2bf4f/Cargo.toml)

`pg_smolid` — 为 PostgreSQL 添加 smolid 类型及辅助函数的 pgrx 扩展。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_smolid";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_smolid';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
