## 用法

来源：

- [官方上游 README](https://github.com/RekGRpth/pg_wthtmltopdf/blob/dc68969a4821945d7743cd4ab4b93258dd365251/README.md)

`pg_wthtmltopdf` — 通过 PostgreSQL 函数将 HTML 渲染为 PDF。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C++`。

```sql
CREATE EXTENSION "pg_wthtmltopdf";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_wthtmltopdf';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
