## 用法

来源：

- [官方上游文档](https://github.com/fabriziomello/pg_normalize_query/blob/bd73d4979dbdb89e152b2a776390369992c8aa42/README.md)

`pg_normalize_query` — 使用 PostgreSQL 解析器逻辑将 SQL 语句中的常量归一化

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_normalize_query";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_normalize_query';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
