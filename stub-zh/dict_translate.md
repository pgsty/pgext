## 用法

来源：

- [官方上游文档](https://github.com/za-arthur/dict_translate/blob/c95381026d760bdfae6a8cb2bab94c7ca9c06575/README.md)

`dict_translate` — 提供用于翻译词条的全文搜索词典模板，可将输入词规范化并替换为配置的译词组。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "dict_translate";
SELECT extversion
FROM pg_extension
WHERE extname = 'dict_translate';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
