## 用法

来源：

- [官方扩展控制文件](https://github.com/uptimejp/pg_stat_statements2/blob/db2e94c49c9be23137f6a8925616535950f2ae89/pg_stat_statements2.control)
- [官方上游文档](https://github.com/uptimejp/pg_stat_statements2/blob/db2e94c49c9be23137f6a8925616535950f2ae89/README)

`pg_stat_statements2` — 基于 PostgreSQL 9.4 的 pg_stat_statements 分支，与 sql_firewall 集成以统计 SQL 执行信息

已复核目录快照记录的版本为 `1.2`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`sql_firewall`。

```sql
CREATE EXTENSION "pg_stat_statements2";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_stat_statements2';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
