## 用法

来源：

- [官方扩展控制文件](https://github.com/AjinkyaTaranekar/ai_toolkit/blob/49c3e95059e2c050e61400829004b70dc08992b2/ai_toolkit.control)
- [官方上游 README](https://github.com/AjinkyaTaranekar/ai_toolkit/blob/49c3e95059e2c050e61400829004b70dc08992b2/README.md)

`ai_toolkit` — 在 PostgreSQL 中提供由 AI 驱动的自然语言查询生成、SQL 解释与错误说明工具。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C++`。
整理后的兼容版本集合为 `18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "ai_toolkit";
SELECT extversion
FROM pg_extension
WHERE extname = 'ai_toolkit';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
