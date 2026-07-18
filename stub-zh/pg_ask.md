## 用法

来源：

- [官方扩展控制文件](https://github.com/Abiji-2020/pg_ask/blob/7c58aeb31042b75c4f1f675b621329c689003e2c/pg/pg_ask.control)
- [官方上游文档](https://github.com/Abiji-2020/pg_ask/blob/7c58aeb31042b75c4f1f675b621329c689003e2c/README.md)

`pg_ask` — 使用 AI 将自然语言请求生成 SQL 的 PostgreSQL 扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C++`。
整理后的兼容版本集合为 `16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_ask";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ask';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
