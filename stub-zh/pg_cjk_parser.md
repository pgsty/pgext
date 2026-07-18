## 用法

来源：

- [官方扩展控制文件](https://github.com/huangjimmy/pg_cjk_parser/blob/master/pg_cjk_parser.control)

`pg_cjk_parser` — pg_cjk_parser：基于 PostgreSQL 默认解析器的全文搜索解析器，将 CJK 文本切分为 2-gram tokens。

已复核目录快照记录的版本为 `0.2.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_cjk_parser";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_cjk_parser';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
