## 用法

来源：

- [官方扩展控制文件](https://github.com/ossc-db/pg_dbms_stats/blob/REL14_0/pg_dbms_stats.control)
- [官方上游文档](https://github.com/ossc-db/pg_dbms_stats/blob/REL14_0/doc/pg_dbms_stats-en.md)
- [官方上游 README](https://github.com/ossc-db/pg_dbms_stats/blob/e041b7cf92f105e41746e8e06b106f1246cc51b0/README.md)

`pg_dbms_stats` — pg_dbms_stats：统计观测类 PostgreSQL 扩展；上游说明为“使用固定统计信息稳定 PostgreSQL 执行计划的统计信息管理扩展”。

已复核目录快照记录的版本为 `14.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `14`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_dbms_stats";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_dbms_stats';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
