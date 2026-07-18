## 用法

来源：

- [官方扩展控制文件](https://github.com/zilder/tuple_fdw/blob/9c737a2043bfd78fe085f0966e3a58b62e38fd70/tuple_fdw.control)
- [官方上游文档](https://github.com/zilder/tuple_fdw/blob/9c737a2043bfd78fe085f0966e3a58b62e38fd70/README.md)

`tuple_fdw` — 用于追加写、LZ4 压缩二进制元组文件冷存储的概念验证型 C 外部数据包装器。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "tuple_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'tuple_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
