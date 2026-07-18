## 用法

来源：

- [官方上游文档](https://github.com/ab1smo/pg_normalize_email/blob/709ccbe4cc2733adeb629bc4873d3f3afda868c4/README.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/pg_normalize_email/)

`pg_normalize_email` — 按常见邮件服务商规则统一大小写、别名后缀、点号和域名的纯 PL/pgSQL 邮箱规范化扩展。

已复核目录快照记录的版本为 `1.0.9`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_normalize_email";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_normalize_email';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
