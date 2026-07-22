## 用法

来源：

- [所复核版本的官方 README](https://github.com/ossc-db/pg_plan_advsr/blob/6b41a96c57c3785ef5315d4368f92d17a3488327/README.md)
- [0.1 版扩展 SQL](https://github.com/ossc-db/pg_plan_advsr/blob/6b41a96c57c3785ef5315d4368f92d17a3488327/pg_plan_advsr--0.1.sql)
- [扩展 control 文件](https://github.com/ossc-db/pg_plan_advsr/blob/6b41a96c57c3785ef5315d4368f92d17a3488327/pg_plan_advsr.control)

`pg_plan_advsr` 分析重复的 `EXPLAIN ANALYZE` 运行结果，并通过 `pg_hint_plan` 与 `pg_store_plans` 的反馈循环减少行数估计误差、探索替代计划。上游明确把它限定在执行计划调优的验证环境，而非商业或生产环境。

### 前置条件与重启

所复核分支要求 PostgreSQL 12 或更高版本。安装版本匹配的 `pg_hint_plan`、`pg_store_plans` 与 `pg_plan_advsr`，加入 `shared_preload_libraries` 后重启 PostgreSQL。PostgreSQL 14 或更高版本若需要扩展统计建议，可选装 `pg_qualstats`。

```conf
shared_preload_libraries = 'pg_hint_plan, pg_plan_advsr, pg_store_plans'
max_parallel_workers_per_gather = 0
max_parallel_workers = 0
compute_query_id = on
```

```sql
CREATE EXTENSION pg_hint_plan;
CREATE EXTENSION pg_store_plans;
CREATE EXTENSION pg_plan_advsr;
```

### 反馈流程

使用稳定、无并发变化的测试数据。启用反馈后，重复执行相同的分析计划；随后检查历史，并在实验结束时关闭反馈。

```sql
SELECT pg_plan_advsr_enable_feedback();

EXPLAIN (ANALYZE, VERBOSE)
SELECT a.id, b.value
FROM test_a AS a
JOIN test_b AS b USING (id)
WHERE a.kind = 'target';

SELECT id, pgsp_queryid, pgsp_planid, execution_time,
       rows_hint, scan_hint, join_hint, lead_hint
FROM plan_repo.plan_history
ORDER BY id;

SELECT pg_plan_advsr_disable_feedback();
```

### 重要对象

- `pg_plan_advsr_enable_feedback()`、`pg_plan_advsr_disable_feedback()`：切换 hint 表反馈流程。
- `plan_repo.plan_history`：记录查询 ID、计划 ID、耗时、估计误差与生成的 hints。
- `plan_repo.norm_queries`、`plan_repo.raw_queries`：规范化与原始查询文本。
- `plan_repo.plan_history_pretty`：经过舍入、面向展示的历史视图。
- `plan_repo.get_hint(bigint)`：返回复现某个 `pg_store_plans` 计划 ID 的 hints。
- `plan_repo.get_extstat(bigint)`：在 PostgreSQL 14+ 且已配置 `pg_qualstats` 时建议扩展统计 SQL。
- `pg_plan_advsr.enabled`、`pg_plan_advsr.quieted`、`pg_plan_advsr.widely`：启用分析、抑制消息，或允许从不含 ANALYZE 的计划生成 hints。

### 安全与限制

`EXPLAIN ANALYZE` 会真正执行语句。除非副作用已隔离且确实需要，否则不要把该流程用于写入或破坏性 SQL。反馈会改变 hint 表状态，中间计划可能更差；应先记录基线，并在取得可度量结果后停止。随附的 `sql/base.sql` 会清空 hint 与历史数据，绝不能在共享环境中运行。

扩展不支持并发执行，也不支持 InitPlans、SubPlans、Append、MergeAppend、基础关系估计修正、分组/表达式统计建议以及 14 以前版本的扩展统计建议；所复核版本未测试并行查询、分区表与 JIT。生成的 hints 与统计语句只能视为建议，必须独立验证计划和运行时效果。
