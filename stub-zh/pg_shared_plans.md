## 用法

来源：

- [官方扩展控制文件](https://github.com/rjuju/pg_shared_plans/blob/f9b04f967e5bd64834b53c3a119384346107bfc6/pg_shared_plans.control)
- [官方上游文档](https://github.com/rjuju/pg_shared_plans/blob/f9b04f967e5bd64834b53c3a119384346107bfc6/README.md)

`pg_shared_plans` — PostgreSQL 共享内存执行计划缓存概念验证扩展

已复核目录快照记录的版本为 `0.0.1`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`pg_stat_statements`。
整理后的兼容版本集合为 `12,13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_shared_plans";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_shared_plans';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
