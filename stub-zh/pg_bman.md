## 用法

来源：

- [官方扩展控制文件](https://github.com/s-hironobu/pg_bman/blob/master/pg_bman.control)

`pg_bman` — PostgreSQL 热备份工具，支持远程全量/增量备份与归档日志辅助操作。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_bman";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_bman';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
