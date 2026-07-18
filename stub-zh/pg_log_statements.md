## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/pg_log_statements/)

`pg_log_statements` — 可按后端进程或连接过滤条件，为选中的 PostgreSQL 会话单独启用 log_statement=all。

已复核目录快照记录的版本为 `0.0.2`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_log_statements";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_log_statements';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
