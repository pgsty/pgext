## 用法

来源：

- [官方扩展控制文件](https://github.com/akorotkov/pg_aa/blob/4ae0a182dfd66e65e18ef8faf460973c8706ac44/pg_aa.control)
- [官方上游文档](https://github.com/akorotkov/pg_aa)

`pg_aa` — pg_aa：为 PostgreSQL 提供 ASCII art 扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_aa";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_aa';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
