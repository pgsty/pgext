## 用法

来源：

- [官方扩展控制文件](https://github.com/RekGRpth/pg_mupdf/blob/a15db6ca07c5e95b8f2d78fe92774c659a3f3c54/pg_mupdf.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/pg_mupdf/)

`pg_mupdf` — pg_mupdf：基于 MuPDF 在 PostgreSQL 中将 HTML 转换为 PDF。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_mupdf";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_mupdf';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
