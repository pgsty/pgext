## 用法

来源：

- [官方上游文档](https://github.com/dgillis/pg_vartype/blob/536d1581db6c9e6651f014797bf6391a5231e9be/README.md)

`pg_vartype` — 在单个 PostgreSQL 列中存储多种标量类型的 VARTYPE 数据类型扩展。

已复核目录快照记录的版本为 `0.4`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_vartype";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_vartype';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
