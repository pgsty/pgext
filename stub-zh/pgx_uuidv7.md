## 用法

来源：

- [官方扩展控制文件](https://github.com/kaznak/pgx_uuidv7/blob/3b3828ae89d634c2e828cc872a1ac96394e6cf60/pgx_uuidv7.control)
- [官方上游文档](https://github.com/kaznak/pgx_uuidv7/blob/3b3828ae89d634c2e828cc872a1ac96394e6cf60/README.md)
- [官方 Rust 包清单](https://github.com/kaznak/pgx_uuidv7/blob/3b3828ae89d634c2e828cc872a1ac96394e6cf60/Cargo.toml)

`pgx_uuidv7` — pgx_uuidv7 是基于 pgrx 的 PostgreSQL 扩展，用于生成 UUIDv7，并在 UUID 与时间戳之间转换。

已复核目录快照记录的版本为 `0.1.7`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgx_uuidv7";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgx_uuidv7';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
