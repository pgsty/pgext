## 用法

来源：

- [官方扩展控制文件](https://github.com/arunchaganty/pg-span/blob/0ff0b6f1070cf6410299701b7fcfe6368513eaa2/span.control)
- [官方上游文档](https://github.com/arunchaganty/pg-span/blob/0ff0b6f1070cf6410299701b7fcfe6368513eaa2/README.md)

`span` — 用于可索引文档文本区间的 C 数据类型扩展，提供转换、比较、B-tree/hash 运算符类、聚合与访问函数。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "span";
SELECT extversion
FROM pg_extension
WHERE extname = 'span';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
