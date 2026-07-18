## 用法

来源：

- [官方扩展控制文件](https://github.com/sangli00/pg_get_page_tuple/blob/a4330098dfc5535e8f5e416658a17261cafff71e/pg_fuck_block.control)
- [官方上游文档](https://github.com/sangli00/pg_get_page_tuple/blob/a4330098dfc5535e8f5e416658a17261cafff71e/README)

`pg_fuck_block` — 读取关系指定数据块中全部元组的 C 扩展

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_fuck_block";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_fuck_block';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
