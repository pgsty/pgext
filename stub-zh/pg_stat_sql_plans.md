## 用法

来源：

- [官方扩展控制文件](https://github.com/legrandlegrand/pg_stat_sql_plans/blob/d0401cdf5ff3f01ceb0e0b7a0d8bc3edbb620d24/pg_stat_sql_plans.control)
- [官方上游文档](https://github.com/legrandlegrand/pg_stat_sql_plans/blob/d0401cdf5ff3f01ceb0e0b7a0d8bc3edbb620d24/README.md)

`pg_stat_sql_plans` — pg_stat_sql_plans：按规范化 SQL 查询和执行计划 ID 跟踪执行统计信息。

已复核目录快照记录的版本为 `0.2`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_stat_sql_plans";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_stat_sql_plans';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
