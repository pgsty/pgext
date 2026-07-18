## 用法

来源：

- [官方扩展控制文件](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/pgjwt_rs.control)
- [官方上游文档](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/README.md)
- [官方 Rust 包清单](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/Cargo.toml)

`pgjwt_rs` — 在 PostgreSQL 中验证 RS256 与 Ed25519 JWT 令牌的安全扩展。

已复核目录快照记录的版本为 `0.1.2`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgjwt_rs";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgjwt_rs';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
