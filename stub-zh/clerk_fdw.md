## 用法

来源：

- [官方扩展控制文件](https://github.com/Jayko001/clerk_fdw/blob/ff55a536f658b2e4461407b8e10b3e47a6dfb5d1/clerk_fdw.control)
- [官方上游文档](https://github.com/Jayko001/clerk_fdw/blob/ff55a536f658b2e4461407b8e10b3e47a6dfb5d1/README.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/clerk_fdw/)

`clerk_fdw` — 用于从 PostgreSQL 查询 Clerk 用户管理 API 的外部数据包装器。

已复核目录快照记录的版本为 `0.3.3`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "clerk_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'clerk_fdw';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
