## 用法

来源：

- [官方扩展控制文件](https://github.com/zulip/tsearch_extras/blob/f566e9606bac00a22c4b62e9511b022c861db05c/tsearch_extras.control)
- [官方上游文档](https://github.com/zulip/tsearch_extras/blob/f566e9606bac00a22c4b62e9511b022c861db05c/README.md)

`tsearch_extras` — 提供全文检索匹配位置提取和 tsvector 词元枚举函数。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "tsearch_extras";
SELECT extversion
FROM pg_extension
WHERE extname = 'tsearch_extras';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
