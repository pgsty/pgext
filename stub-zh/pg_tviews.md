## 用法

来源：

- [官方扩展控制文件](https://github.com/fraiseql/pg_tviews/blob/5f4288db5d1fdc565b01468ec3a5c9ca4124c447/pg_tviews.control)
- [官方上游文档](https://github.com/fraiseql/pg_tviews/blob/5f4288db5d1fdc565b01468ec3a5c9ca4124c447/README.md)
- [官方 Rust 包清单](https://github.com/fraiseql/pg_tviews/blob/5f4288db5d1fdc565b01468ec3a5c9ca4124c447/Cargo.toml)

`pg_tviews` — pg_tviews：通用类 PostgreSQL 扩展；上游说明为“为 PostgreSQL 提供带自动刷新触发器的事务视图（TVIEW）”。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_tviews";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_tviews';
```

整理后的生命周期为 `preview`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
