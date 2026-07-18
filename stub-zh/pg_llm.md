## 用法

来源：

- [官方扩展控制文件](https://github.com/onelazyteam/pg_llm/blob/d69b9d4f4da3ec1cc4df480f95483545e89ea309/pg_llm.control)
- [官方上游文档](https://github.com/onelazyteam/pg_llm/blob/d69b9d4f4da3ec1cc4df480f95483545e89ea309/README.md)

`pg_llm` — 将本地及远程大语言模型与聊天、会话、文本转 SQL、审计和向量知识库功能集成到 PostgreSQL。

已复核目录快照记录的版本为 `1.1`、类型为 `preload`、实现语言为 `C++`。
应先安装并验证声明的扩展依赖：`vector`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_llm";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_llm';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
