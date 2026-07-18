## 用法

来源：

- [官方扩展控制文件](https://github.com/snaga/monetdb_fdw/blob/fbff386be673bb6f1d43cddbceb0081a20e9b459/monetdb_fdw.control)
- [官方上游文档](https://github.com/snaga/monetdb_fdw/blob/fbff386be673bb6f1d43cddbceb0081a20e9b459/README.md)

`monetdb_fdw` — monetdb_fdw：用于访问 MonetDB 的 PostgreSQL 外部数据包装器。

已复核目录快照记录的版本为 `0.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "monetdb_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'monetdb_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
