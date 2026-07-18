## 用法

来源：

- [官方上游文档](https://github.com/pgEdge/pg_semantic_cache/blob/38e4a910f1ae7e4050abbb8bc1f9cf5d44abf4cc/docs/index.md)

`pg_semantic_cache` — 使用 pgvector 嵌入缓存查询结果并检索语义相似结果，支持可配置的淘汰策略。

已复核目录快照记录的版本为 `0.1.0-beta4`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`, `vector`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_semantic_cache";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_semantic_cache';
```

整理后的生命周期为 `preview`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
