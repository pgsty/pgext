## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/pmpp/)

`pmpp` — 通过异步 libpq 连接并行运行手工拆分的 SQL 工作并汇总结果

已复核目录快照记录的版本为 `1.2.3`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pmpp";
SELECT extversion
FROM pg_extension
WHERE extname = 'pmpp';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
