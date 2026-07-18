## 用法

来源：

- [官方扩展控制文件](https://github.com/bigsql/pgtsql/blob/3240980dff20736a38101eb0806c8580ff169d9c/pgtsql.control)
- [官方上游文档](https://github.com/bigsql/pgtsql/blob/3240980dff20736a38101eb0806c8580ff169d9c/README.md)

`pgtsql` — 为 PostgreSQL 提供 Transact-SQL 兼容的过程语言、函数与系统目录视图。

已复核目录快照记录的版本为 `3.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `11`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgtsql";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgtsql';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
