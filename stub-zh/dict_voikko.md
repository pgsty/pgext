## 用法

来源：

- [官方扩展控制文件](https://github.com/Houston-Inc/dict_voikko/blob/746eeef7f40cadbe5a527ccee4250a2f8daf2a94/dict_voikko.control)
- [官方上游文档](https://github.com/Houston-Inc/dict_voikko/blob/746eeef7f40cadbe5a527ccee4250a2f8daf2a94/README.md)

`dict_voikko` — 使用 Voikko 形态分析的芬兰语 PostgreSQL 全文搜索词典。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "dict_voikko";
SELECT extversion
FROM pg_extension
WHERE extname = 'dict_voikko';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
