## 用法

来源：

- [官方扩展控制文件](https://github.com/obartunov/dict_regex/blob/a8b8d3270c3d50d59eb8064efd1154128b39c548/dict_regex.control)
- [官方上游文档](https://github.com/obartunov/dict_regex/blob/a8b8d3270c3d50d59eb8064efd1154128b39c548/README.dict_regex)

`dict_regex` — 使用正则表达式规则的 PostgreSQL 全文搜索词典。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "dict_regex";
SELECT extversion
FROM pg_extension
WHERE extname = 'dict_regex';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
