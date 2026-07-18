## 用法

来源：

- [官方扩展控制文件](https://github.com/mpajkowski/natord_pg/blob/a1f2923934a1ae4179e559ea651641f1b975bef0/natord_pg.control)
- [官方 Rust 包清单](https://github.com/mpajkowski/natord_pg/blob/a1f2923934a1ae4179e559ea651641f1b975bef0/Cargo.toml)

`natord_pg` — 为 PostgreSQL text 值提供自然排序比较操作符和 btree 操作符类。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "natord_pg";
SELECT extversion
FROM pg_extension
WHERE extname = 'natord_pg';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
