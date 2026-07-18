## 用法

来源：

- [官方扩展控制文件](https://github.com/carderne/upid/blob/d820a74227adb803d378d89da9eacba78c1a5e09/upid_pg/upid_pg.control)
- [官方上游文档](https://github.com/carderne/upid/blob/d820a74227adb803d378d89da9eacba78c1a5e09/README.md)
- [官方 Rust 包清单](https://github.com/carderne/upid/blob/d820a74227adb803d378d89da9eacba78c1a5e09/upid_pg/Cargo.toml)

`upid_pg` — 带语义前缀、可排序的 128 位标识符类型与生成器

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "upid_pg";
SELECT extversion
FROM pg_extension
WHERE extname = 'upid_pg';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
