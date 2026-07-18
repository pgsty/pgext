## 用法

来源：

- [官方扩展控制文件](https://github.com/powa-team/pg_sortstats/blob/b9228cca5947b010c8114ecd5091e76d57ef28c3/pg_sortstats.control)
- [官方上游文档](https://github.com/powa-team/pg_sortstats/blob/b9228cca5947b010c8114ecd5091e76d57ef28c3/README.md)

`pg_sortstats` — 收集按查询聚合的排序统计并估算所需 work_mem

已复核目录快照记录的版本为 `0.0.1`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`pg_stat_statements`。

```sql
CREATE EXTENSION "pg_sortstats";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_sortstats';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
