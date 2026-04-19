## 用法

来源：[official docs](https://ossc-db.github.io/pg_store_plans/)，[repo](https://github.com/ossc-db/pg_store_plans)，[1.10 release notes](https://github.com/ossc-db/pg_store_plans/releases/tag/1.10)

`pg_store_plans` 用于跟踪 SQL 语句的执行计划统计信息，作用类似于 `pg_stat_statements` 对语句统计的补充。上游 `1.10` 发布说明指出该版本新增 PostgreSQL 18 支持，其余 SQL 接口与当前文档保持一致。

### 必要服务器设置

```ini
shared_preload_libraries = 'pg_store_plans'
compute_query_id = 'on'
```

`pg_store_plans` 需要共享内存，因此增删该扩展都需要重启。文档还说明如果 `compute_query_id = 'no'`，模块会被静默禁用。

### 查看已存储计划

```sql
SELECT queryid, planid, plan, calls, total_time, rows
FROM pg_store_plans
ORDER BY total_time DESC;

SELECT * FROM pg_store_plans_info;
```

文档将 `queryid` 作为与 `pg_stat_statements` 的关联键，并将 `pg_store_plans_info` 描述为单行视图，用于暴露 `dealloc`、`stats_reset` 等模块级统计。

### 辅助函数

```sql
SELECT pg_store_plans_reset();
SELECT pg_store_hash_query('SELECT 1');
SELECT pg_store_plans_textplan(plan);
SELECT pg_store_plans_jsonplan(plan);
SELECT pg_store_plans_xmlplan(plan);
SELECT pg_store_plans_yamlplan(plan);
```

当 `pg_store_plans.plan_format = 'raw'` 时，`pg_store_plans_*plan()` 尤其实用。

### 关键 GUC

- `pg_store_plans.max`
- `pg_store_plans.track`
- `pg_store_plans.max_plan_length`
- `pg_store_plans.plan_storage`
- `pg_store_plans.plan_format`
- `pg_store_plans.min_duration`
- `pg_store_plans.log_analyze`
- `pg_store_plans.log_buffers`
- `pg_store_plans.log_timing`
- `pg_store_plans.save`

文档说明 `plan_storage` 取值为 `file` 或 `shmem`，`plan_format` 取值为 `text`、`json`、`xml`、`yaml` 或 `raw`。

### 注意事项

- 非超级用户不能查看其他用户执行语句对应的 `plan`、`queryid` 或 `planid`。
- `pg_store_plans` 与 `pg_stat_statements` 独立维护条目，因此低频记录不一定总能在另一侧找到对应行。
