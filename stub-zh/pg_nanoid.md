## 用法

来源：

- [官方上游文档](https://github.com/lameferret/pg_nanoid/blob/d216e0b481d8ffb8a2c93b1dc3420bdafe31b486/Examples.md)
- [官方 Rust 包清单](https://github.com/lameferret/pg_nanoid/blob/d216e0b481d8ffb8a2c93b1dc3420bdafe31b486/Cargo.toml)

`pg_nanoid` — 为 PostgreSQL 增加 NanoID 类型及 NanoID 生成函数。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_nanoid";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_nanoid';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
