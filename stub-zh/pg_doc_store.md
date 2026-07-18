## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/pg_doc_store/)

`pg_doc_store` — 提供基于 JSONB、支持全文检索的文档存储 API。

已复核目录快照记录的版本为 `0.2.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`pgcrypto`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_doc_store";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_doc_store';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
