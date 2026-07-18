## 用法

来源：

- [官方扩展控制文件](https://github.com/ossc-db/pg_plan_advsr/blob/6b41a96c57c3785ef5315d4368f92d17a3488327/pg_plan_advsr.control)

`pg_plan_advsr` — pg_plan_advsr：功能增强类 PostgreSQL 扩展；上游说明为“自动执行计划调优扩展”。

已复核目录快照记录的版本为 `0.1`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`pg_hint_plan`, `pg_store_plans`。
整理后的兼容版本集合为 `12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_plan_advsr";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_plan_advsr';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
