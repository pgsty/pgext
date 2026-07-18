## 用法

来源：

- [官方扩展控制文件](https://github.com/unionlabs/pg_ibc/blob/14abdfdfc6e0fb63500147ed16de14195b0610dc/pg_ibc_0_1.control)
- [官方 Rust 包清单](https://github.com/unionlabs/pg_ibc/blob/14abdfdfc6e0fb63500147ed16de14195b0610dc/Cargo.toml)

`pg_ibc_0_1` — 用于解码和处理 IBC 协议数据类型的 PostgreSQL 扩展

已复核目录快照记录的版本为 `0.1.1`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_ibc_0_1";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ibc_0_1';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
