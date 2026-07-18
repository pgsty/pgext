## 用法

来源：

- [官方上游文档](https://github.com/adjust/pg_querylog/blob/f639adb6dc1686044b76d82139df1c3a49076053/README.md)

`pg_querylog` — 在共享内存中捕获当前及最近完成的后端查询，供 SQL 检查。

已复核目录快照记录的版本为 `0.1`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_querylog";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_querylog';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
