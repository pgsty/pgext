## 用法

来源：

- [官方扩展控制文件](https://github.com/hiyenwong/pg_knowledge_graph/blob/904e15c6500e23897dc626338b6469d639270d76/pg_knowledge_graph.control)
- [官方上游文档](https://github.com/hiyenwong/pg_knowledge_graph/blob/904e15c6500e23897dc626338b6469d639270d76/README.md)
- [官方 Rust 包清单](https://github.com/hiyenwong/pg_knowledge_graph/blob/904e15c6500e23897dc626338b6469d639270d76/Cargo.toml)

`pg_knowledge_graph` — 面向 PostgreSQL 的知识图谱扩展，提供图算法与 pgvector 集成。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
应先安装并验证声明的扩展依赖：`vector`。
整理后的兼容版本集合为 `16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_knowledge_graph";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_knowledge_graph';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
