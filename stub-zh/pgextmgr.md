## 用法

来源：

- [官方扩展控制文件](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgextmgr/pgextmgr.control)
- [官方上游文档](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/README.md)
- [官方 Rust 包清单](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgextmgr/Cargo.toml)

`pgextmgr` — pgextmgr：基于 pgrx 的扩展管理与 Hook 框架，用于注册、排序、启用和禁用 PostgreSQL 扩展插件。

已复核目录快照记录的版本为 `0.0.0`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgextmgr";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgextmgr';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
