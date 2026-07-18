## 用法

来源：

- [官方扩展控制文件](https://github.com/zdk/pg-search-thai/blob/e943164a3640a65c5714028447189f9a3a7b4f82/thai_parser/thai_parser.control)
- [官方上游文档](https://github.com/zdk/pg-search-thai/blob/e943164a3640a65c5714028447189f9a3a7b4f82/README.md)

`thai_parser` — 为 PostgreSQL 全文检索提供泰语分词解析器和文本搜索配置。

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "thai_parser";
SELECT extversion
FROM pg_extension
WHERE extname = 'thai_parser';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
