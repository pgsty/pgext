## 用法

来源：

- [官方扩展控制文件](https://github.com/ergo70/pg_etag/blob/444c21d9df4475179d9ee40ec07fcf51a48e5501/pg_etag.control)
- [官方上游文档](https://github.com/ergo70/pg_etag/blob/444c21d9df4475179d9ee40ec07fcf51a48e5501/README.md)

`pg_etag` — 使用 BLAKE2 哈希为行和查询结果集快速生成 ETag，提供单值函数和聚合函数。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_etag";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_etag';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
