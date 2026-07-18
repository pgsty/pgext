## 用法

来源：

- [官方扩展控制文件](https://github.com/pgcodekeeper/pg_dbo_timestamp/blob/8392fee066d972cb97363a1e822a54251cd55840/pg_dbo_timestamp.control)
- [官方上游文档](https://github.com/pgcodekeeper/pg_dbo_timestamp/blob/8392fee066d972cb97363a1e822a54251cd55840/README.md)

`pg_dbo_timestamp` — 记录数据库结构变更时间及执行者的 PostgreSQL 扩展。

已复核目录快照记录的版本为 `1.0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_dbo_timestamp";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_dbo_timestamp';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
