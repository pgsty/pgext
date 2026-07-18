## 用法

来源：

- [官方上游文档](https://github.com/MasahikoSawada/pg_simula/blob/532efbc57ebd57de8308af8278e77581bb6be885/README.md)

`pg_simula` — 向 PostgreSQL 操作注入 WAIT、ERROR、FATAL 或 PANIC 故障

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_simula";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_simula';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
