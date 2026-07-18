## 用法

来源：

- [官方扩展控制文件](https://github.com/pierreforstmann/pg_logqueryid/blob/306da69ba0dc73f49ab08feaf61b83e5f521cdd1/pg_logqueryid.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/pg_logqueryid/)

`pg_logqueryid` — pg_logqueryid：统计观测类 PostgreSQL 扩展；在 auto_explain 日志中输出 pg_stat_statements queryid。

已复核目录快照记录的版本为 `0.0.1`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`auto_explain`, `pg_stat_statements`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_logqueryid";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_logqueryid';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
