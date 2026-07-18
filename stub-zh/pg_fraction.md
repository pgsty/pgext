## 用法

来源：

- [官方扩展控制文件](https://github.com/nuko-yokohama/pg_fraction/blob/812f0599f82a47024b96ed0c745c0464bf1e2f2b/pg_fraction.control)

`pg_fraction` — 提供分数数据类型，以及算术、比较、聚合和 B-tree 运算符类支持。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_fraction";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_fraction';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
