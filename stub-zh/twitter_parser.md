## 用法

来源：

- [官方扩展控制文件](https://github.com/3manuek/twitter_parser/blob/1fafaae28bd7dc2c0d01617ddf992a19027f2c43/twitter_parser.control)
- [官方上游文档](https://github.com/3manuek/twitter_parser/blob/1fafaae28bd7dc2c0d01617ddf992a19027f2c43/README)

`twitter_parser` — 用于 PostgreSQL 全文检索、解析 Twitter 风格话题标签和用户提及的轻量 C 解析器。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "twitter_parser";
SELECT extversion
FROM pg_extension
WHERE extname = 'twitter_parser';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
