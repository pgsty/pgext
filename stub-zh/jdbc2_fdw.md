## 用法

来源：

- [官方扩展控制文件](https://github.com/heimir-sverrisson/jdbc2_fdw/blob/758caf50c674577144b644f36432c4736ec99627/jdbc2_fdw.control)
- [官方上游文档](https://github.com/heimir-sverrisson/jdbc2_fdw/blob/758caf50c674577144b644f36432c4736ec99627/README.md)

`jdbc2_fdw` — 面向可通过 JDBC 访问的数据库的实验性外部数据包装器，支持远程过滤条件下推。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "jdbc2_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'jdbc2_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
