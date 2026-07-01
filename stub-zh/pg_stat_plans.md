


## 用法

来源：[README](https://github.com/pganalyze/pg_stat_plans/blob/main/README.md)，[v2.1.0 release](https://github.com/pganalyze/pg_stat_plans/releases/tag/v2.1.0)，[SQL objects](https://github.com/pganalyze/pg_stat_plans/blob/main/pg_stat_plans--2.0.sql)

`pg_stat_plans` 用于跟踪 PostgreSQL 执行计划形状的聚合统计。它把计划树哈希为 plan ID，把示例 `EXPLAIN` 文本存放在共享内存中，并帮助识别同一个 query ID 是否出现了不同计划。

### 启用

`pg_stat_plans` 需要 PostgreSQL 16 或更新版本，并且必须在服务器启动时加载：

```conf
shared_preload_libraries = 'pg_stat_plans'
pg_stat_plans.compress = 'zstd'
```

```sql
CREATE EXTENSION pg_stat_plans;
```

通常建议和 `pg_stat_statements` 一起使用，便于把 plan ID 与查询文本关联起来。

### 查询计划

```sql
SELECT *
FROM pg_stat_plans;
```

视图包含 `userid`、`dbid`、`toplevel`、`queryid`、`planid`、`calls`、`total_exec_time`、`plan`。如果只需要统计而不需要计划文本：

```sql
SELECT *
FROM pg_stat_plans(false);
```

按 `queryid` 汇总可以查看同一个归一化查询选择过哪些计划形状：

```sql
SELECT queryid, planid, calls, total_exec_time / NULLIF(calls, 0) AS avg_exec_time
FROM pg_stat_plans(false)
ORDER BY queryid, avg_exec_time DESC;
```

### 运行中查询

在 PostgreSQL 18 及更新版本上，`pg_stat_plans_activity` 可以显示当前运行中查询的 plan ID 和示例计划：

```sql
SELECT *
FROM pg_stat_plans_activity;
```

### 重置与配置

```sql
SELECT pg_stat_plans_reset();
```

重要设置包括 `pg_stat_plans.max`、`pg_stat_plans.max_size`、`pg_stat_plans.max_plan_memory`、`pg_stat_plans.track`、`pg_stat_plans.compress`、`pg_stat_plans.plan_advice`。

### 注意事项

统计基于 PostgreSQL cumulative statistics 系统，因此计数器会在事务结束时刷新，并可能有延迟。Plan ID 描述的是计划形状，分区、类型转换或表达式细节变化都可能导致 plan ID 变化。
