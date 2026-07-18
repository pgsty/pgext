## 用法

来源：

- [官方 Rust 包清单](https://github.com/zvectorlabs/pgauthz/blob/a24a3bae33d1bac320c233e99ceb9153a4d0667d/crates/pgauthz/Cargo.toml)

`pgauthz` — 在 PostgreSQL 内提供 Zanzibar 风格的关系型访问控制

已复核目录快照记录的版本为 `0.3.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgauthz";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgauthz';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
