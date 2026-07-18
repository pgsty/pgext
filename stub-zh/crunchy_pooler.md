## 用法

来源：

- [官方上游文档](https://docs.crunchybridge.com/how-to/pgbouncer)
- [官方项目或服务商页面](https://www.crunchydata.com/products/crunchy-bridge)

`crunchy_pooler` — Crunchy Bridge 的 PgBouncer 认证辅助扩展，通过 PostgreSQL 的用户存储验证连接身份。

已复核目录快照记录的版本为 `unknown`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "crunchy_pooler";
SELECT extversion
FROM pg_extension
WHERE extname = 'crunchy_pooler';
```

这是 `Crunchy Data` 的服务商专用组件；可用性、启用、权限与升级遵循该服务，而不是可移植的社区软件包。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
