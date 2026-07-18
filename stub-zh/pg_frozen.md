## 用法

来源：

- [官方扩展控制文件](https://github.com/lawrencejones/pg_frozen/blob/68004b4f21b01041487e2a080658565dc4cdc937/pg_frozen.control)
- [官方上游文档](https://github.com/lawrencejones/pg_frozen/blob/68004b4f21b01041487e2a080658565dc4cdc937/README.md)

`pg_frozen` — 提供 frozen(oid, tid) 函数，用于判断指定堆元组是否已被冻结。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_frozen";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_frozen';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
