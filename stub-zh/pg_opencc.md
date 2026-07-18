## 用法

来源：

- [官方扩展控制文件](https://github.com/VitoVan/pg_opencc/blob/c919ba0cff11949c682c583d80f72dfb15eca87a/pg_opencc.control)
- [官方上游文档](https://github.com/VitoVan/pg_opencc/blob/c919ba0cff11949c682c583d80f72dfb15eca87a/README.md)

`pg_opencc` — 使用 OpenCC 在简体中文、各地区繁体中文与日文新字体之间转换文本。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `13`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_opencc";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_opencc';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
