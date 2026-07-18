## 用法

来源：

- [官方扩展控制文件](https://github.com/ratbox/pg-digest/blob/44ebe696700423cdd64ccc8f1d8b9dc18728934a/pg_digest.control)
- [官方上游文档](https://github.com/ratbox/pg-digest/blob/44ebe696700423cdd64ccc8f1d8b9dc18728934a/README.md)

`pg_digest` — pg_digest：为 PostgreSQL 提供高效存储哈希摘要的数据类型。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_digest";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_digest';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
