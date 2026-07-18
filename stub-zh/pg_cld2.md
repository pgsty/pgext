## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/pg_cld2/)

`pg_cld2` — 在 PostgreSQL 中使用 CLD2 检测文本的自然语言和文字系统。

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `C++`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_cld2";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_cld2';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
