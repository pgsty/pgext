## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/range_agg/)

`range_agg` — range_agg：合并相邻或重叠的 PostgreSQL 范围值

已复核目录快照记录的版本为 `1.2.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "range_agg";
SELECT extversion
FROM pg_extension
WHERE extname = 'range_agg';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
