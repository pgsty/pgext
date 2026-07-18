## 用法

来源：

- [官方扩展控制文件](https://github.com/alexeyklyukin/translate_isbn/blob/dd65576710245916ab14e7ff00d8f94058e63e23/translate_isbn.control)
- [官方上游文档](https://github.com/alexeyklyukin/translate_isbn/blob/dd65576710245916ab14e7ff00d8f94058e63e23/README.md)

`translate_isbn` — 以 C 函数在 ISBN-10 与 ISBN-13 文本编码之间转换，用于兼容 Evergreen 的检索扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "translate_isbn";
SELECT extversion
FROM pg_extension
WHERE extname = 'translate_isbn';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
