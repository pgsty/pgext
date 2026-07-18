## 用法

来源：

- [官方上游文档](https://docs.pgedge.com/pgedge-vectorizer/)

`pgedge_vectorizer` — 在 PostgreSQL 中异步分块文本，并通过后台 worker 与 pgvector 生成和维护嵌入向量。

已复核目录快照记录的版本为 `1.1`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`vector`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgedge_vectorizer";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgedge_vectorizer';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
