## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/pg_icu_parser/)

`pg_icu_parser` — 使用 ICU 边界分析的全文搜索解析器

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_icu_parser";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_icu_parser';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
