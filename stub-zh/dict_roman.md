## 用法

来源：

- [官方扩展控制文件](https://github.com/obartunov/dict_roman/blob/22c2cd1acd8eb077a6d76e543e4ae6a4a0071d61/dict_roman.control)
- [官方上游文档](https://github.com/obartunov/dict_roman/blob/22c2cd1acd8eb077a6d76e543e4ae6a4a0071d61/README.dict_roman)

`dict_roman` — 将罗马数字词元转换为阿拉伯整数的全文搜索词典。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "dict_roman";
SELECT extversion
FROM pg_extension
WHERE extname = 'dict_roman';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
