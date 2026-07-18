## 用法

来源：

- [官方扩展控制文件](https://github.com/2ndQuadrant/array_textsort/blob/dd60ea1a5403e1ff938aeb6d26d0e9a901ea30f4/array_textsort.control)
- [官方上游文档](https://github.com/2ndQuadrant/array_textsort/blob/dd60ea1a5403e1ff938aeb6d26d0e9a901ea30f4/README.md)

`array_textsort` — 提供 array_textsort 和 array_distinct 函数，用于排序和去重一维 text 数组。

已复核目录快照记录的版本为 `1.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "array_textsort";
SELECT extversion
FROM pg_extension
WHERE extname = 'array_textsort';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
