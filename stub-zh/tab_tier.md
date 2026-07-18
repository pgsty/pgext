## 用法

来源：

- [官方上游文档](https://pgxn.org/dist/tab_tier/doc/tab_tier.html)
- [官方 PGXN 分发页](https://pgxn.org/dist/tab_tier/)

`tab_tier` — 以作业驱动的 PL/pgSQL 扩展，用于保留策略、数据分层和基于日期的表分区维护。

已复核目录快照记录的版本为 `1.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "tab_tier";
SELECT extversion
FROM pg_extension
WHERE extname = 'tab_tier';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
