## 用法

来源：

- [官方扩展控制文件](https://github.com/philopon/cas/blob/e23971e2a65e5c9e50fa238cce9bb4209f041b36/cas.control)

`cas` — cas：为 PostgreSQL 提供 CAS Registry Number 数据类型。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "cas";
SELECT extversion
FROM pg_extension
WHERE extname = 'cas';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
