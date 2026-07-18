## 用法

来源：

- [官方上游文档](https://github.com/TypedLambda/pg_secret/blob/eef18ef9f504fa55fb19011e9f4844495f89709a/NOTES)

`pg_secret` — 定义顺序可揭示加密域、比较运算符和 B-tree 运算符类。

已复核目录快照记录的版本为 `0.5`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_secret";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_secret';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
