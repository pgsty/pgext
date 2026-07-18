## 用法

来源：

- [官方扩展控制文件](https://github.com/adunstan/file_text_array_fdw/blob/dde60712f74a249e9d3675f8c33ee18c33d23bd1/file_textarray_fdw.control)
- [官方上游文档](https://github.com/adunstan/file_text_array_fdw/blob/dde60712f74a249e9d3675f8c33ee18c33d23bd1/README.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/file_textarray_fdw/)

`file_textarray_fdw` — file_textarray_fdw：将文本文件的每行解析为 text[] 的 PostgreSQL 外部数据包装器。

已复核目录快照记录的版本为 `1.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "file_textarray_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'file_textarray_fdw';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
