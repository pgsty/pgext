## 用法

来源：

- [官方扩展控制文件](https://github.com/dsarafan/pg_dirty_hands/blob/3d5bb01c39e987bc394fc0f093c37d673744ba7e/pg_dirty_hands.control)
- [官方上游文档](https://github.com/dsarafan/pg_dirty_hands/blob/3d5bb01c39e987bc394fc0f093c37d673744ba7e/README.md)

`pg_dirty_hands` — 提供用于冻结指定堆元组的函数，包括不写 WAL 的变体。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_dirty_hands";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_dirty_hands';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
