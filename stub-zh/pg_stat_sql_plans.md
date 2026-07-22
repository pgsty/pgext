## 用法

来源：

- [Official extension control file](https://github.com/legrandlegrand/pg_stat_sql_plans/blob/d0401cdf5ff3f01ceb0e0b7a0d8bc3edbb620d24/pg_stat_sql_plans.control)
- [Official upstream documentation](https://github.com/legrandlegrand/pg_stat_sql_plans/blob/d0401cdf5ff3f01ceb0e0b7a0d8bc3edbb620d24/README.md)

`pg_stat_sql_plans` 0.2 是 `pg_stat_statements` 的 alpha 分支，按规范化查询 ID 与计划 ID 记录统计信息。它能区分同一 SQL 形态的计划变化，并保留规划、执行、扩展开销、失败、首次调用和最近调用信息；上游明确说明不要用于生产环境。

### 核心流程

对于这里审阅的分支，应使用 PostgreSQL 14 或更高版本，预加载模块、关闭内核查询 ID 计算、重启，然后创建扩展。

```conf
shared_preload_libraries = 'pg_stat_sql_plans'
compute_query_id = off
```

```sql
CREATE EXTENSION pg_stat_sql_plans;

SELECT queryid, planid, calls, plan_time, exec_time, first_call, last_call
FROM pg_stat_sql_plans
ORDER BY exec_time DESC;

SELECT * FROM pg_stat_sql_plans_agg;
```

`pg_stat_sql_plans` 为每个查询与计划组合保存一行。`pg_stat_sql_plans_agg` 按查询聚合。`pgssp_backend_qpid(pid)` 暴露后端最近的查询计划标识，`pgssp_normalize_query(text)` 规范化 SQL 文本，`pg_stat_sql_plans_reset()` 清空已收集条目。

### 参数与注意事项

重要参数包括 `pg_stat_sql_plans.max`、`pg_stat_sql_plans.plan_type`、`pg_stat_sql_plans.track`、`pg_stat_sql_plans.track_errors`、`pg_stat_sql_plans.track_pid`、`pg_stat_sql_plans.track_utility`、`pg_stat_sql_plans.save` 和 `pg_stat_sql_plans.explain`。计划哈希和可选计划日志会增加规划、内存与日志开销；保存的查询文本可能包含字面量。应限制访问、测量开销，并把这些 ID 当作实现派生的诊断值，而不是可移植身份。
