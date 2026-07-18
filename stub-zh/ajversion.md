## 用法

来源：

- [官方扩展控制文件](https://github.com/adjust/pg-ajversion/blob/502776cf0c17d7ad034a2b5bb36ba742f3e8334d/ajversion.control)
- [官方上游文档](https://github.com/adjust/pg-ajversion/blob/502776cf0c17d7ad034a2b5bb36ba742f3e8334d/README.md)

`ajversion` — PostgreSQL 语义版本号类型，使用紧凑的整数形式存储。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "ajversion";
SELECT extversion
FROM pg_extension
WHERE extname = 'ajversion';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
