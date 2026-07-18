## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/pg_natural_sort_order/)

`pg_natural_sort_order` — 通过 C 语言归一化函数提供可用于索引的自然排序支持。

已复核目录快照记录的版本为 `1.0.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_natural_sort_order";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_natural_sort_order';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
