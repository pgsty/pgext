## 用法

来源：

- [官方扩展控制文件](https://github.com/pierreforstmann/pg_timeout/blob/cabf38c88b48da64e84d226793da762e5c6e14a1/pg_timeout.control)
- [官方上游文档](https://github.com/pierreforstmann/pg_timeout/blob/cabf38c88b48da64e84d226793da762e5c6e14a1/README.md)

`pg_timeout` — 用于终止空闲数据库会话的 PostgreSQL 后台工作进程扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_timeout";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_timeout';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
