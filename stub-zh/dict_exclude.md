## 用法

来源：

- [官方扩展控制文件](https://github.com/no0p/dict_exclude/blob/2058a16ca44942a0c11eae8f7bffa9a733d7cc5a/dict_exclude.control)
- [官方上游文档](https://github.com/no0p/dict_exclude/blob/2058a16ca44942a0c11eae8f7bffa9a733d7cc5a/README.md)

`dict_exclude` — 使用正则表达式规则过滤停用词的全文搜索词典。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "dict_exclude";
SELECT extversion
FROM pg_extension
WHERE extname = 'dict_exclude';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
