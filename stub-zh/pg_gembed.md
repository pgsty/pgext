## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/pg_gembed/)
- [官方主文档](https://pgxn.org/dist/pg_gembed/1.0.0/)

`pg_gembed` — 在 PostgreSQL 内调用本地 Rust 推理核心生成文本与多模态向量嵌入。

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`vector`。
整理后的兼容版本集合为 `18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_gembed";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_gembed';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
