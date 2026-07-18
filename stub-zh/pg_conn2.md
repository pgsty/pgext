## 用法

来源：

- [官方扩展控制文件](https://github.com/serpent7776/pg_conn2/blob/master/pg_conn2.control)
- [官方上游文档](https://github.com/serpent7776/pg_conn2/blob/master/README.md)

`pg_conn2` — 从 SQL 中创建并管理独立的本地 PostgreSQL 数据库连接。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_conn2";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_conn2';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
