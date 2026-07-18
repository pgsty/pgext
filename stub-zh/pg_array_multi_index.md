## 用法

来源：

- [官方扩展控制文件](https://github.com/rlichtenwalter/pg_array_multi_index/blob/d315364e110db6ec50e822e2eb5957d988d5a5a6/pg_array_multi_index.control)

`pg_array_multi_index` — pg_array_multi_index：为 PostgreSQL 数组提供多下标访问函数。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_array_multi_index";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_array_multi_index';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
