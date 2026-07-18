## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/pg_htmldoc/)

`pg_htmldoc` — 在 PostgreSQL 中将 HTML、Markdown、文件或网页转换为 PDF、PostScript 等 HTMLDOC 输出格式。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_htmldoc";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_htmldoc';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
