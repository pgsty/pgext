## 用法

来源：

- [官方扩展控制文件](https://github.com/einhverfr/pgxn-postgres_monitoring/blob/b2b877bcca7ccc129f281c96587248e9ef5398a5/pg_monitoring.control)
- [官方上游文档](https://github.com/einhverfr/pgxn-postgres_monitoring/blob/b2b877bcca7ccc129f281c96587248e9ef5398a5/README)
- [官方 PGXN 分发页](https://pgxn.org/dist/pg_monitoring/)

`pg_monitoring` — 提供统一 SQL 接口，用于收集 PostgreSQL 复制、表负载与数据库负载监控统计。

已复核目录快照记录的版本为 `0.2`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_monitoring";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_monitoring';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
