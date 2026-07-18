## 用法

来源：

- [官方扩展控制文件](https://github.com/joelonsql/pg-timeit/blob/d83ab65bc9ae7b919bd99de1ebf3e2177bdea607/timeit.control)
- [官方上游文档](https://github.com/joelonsql/pg-timeit/blob/d83ab65bc9ae7b919bd99de1ebf3e2177bdea607/README.md)

`timeit` — 对 PostgreSQL 内置函数与 SQL 表达式进行高分辨率基准测量。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "timeit";
SELECT extversion
FROM pg_extension
WHERE extname = 'timeit';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
