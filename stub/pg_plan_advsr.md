## Usage

Sources:

- [Official README for the reviewed revision](https://github.com/ossc-db/pg_plan_advsr/blob/6b41a96c57c3785ef5315d4368f92d17a3488327/README.md)
- [Extension SQL for version 0.1](https://github.com/ossc-db/pg_plan_advsr/blob/6b41a96c57c3785ef5315d4368f92d17a3488327/pg_plan_advsr--0.1.sql)
- [Extension control file](https://github.com/ossc-db/pg_plan_advsr/blob/6b41a96c57c3785ef5315d4368f92d17a3488327/pg_plan_advsr.control)

`pg_plan_advsr` analyzes repeated `EXPLAIN ANALYZE` runs and uses a feedback loop with `pg_hint_plan` and `pg_store_plans` to reduce row-estimation errors and explore alternative plans. Upstream explicitly limits it to a plan-tuning validation environment, not a commercial or production environment.

### Prerequisites and Restart

The reviewed branch requires PostgreSQL 12 or later. Install matching builds of `pg_hint_plan`, `pg_store_plans`, and `pg_plan_advsr`, add them to `shared_preload_libraries`, and restart PostgreSQL. `pg_qualstats` is optional for extended-statistics suggestions on PostgreSQL 14 or later.

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

### Feedback Workflow

Use stable, non-concurrent test data. Enable feedback, repeatedly execute the same analyzed plan, then inspect the history and disable feedback when the experiment is complete.

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

### Important Objects

- `pg_plan_advsr_enable_feedback()`, `pg_plan_advsr_disable_feedback()`: switch the hint-table feedback workflow.
- `plan_repo.plan_history`: recorded query IDs, plan IDs, timings, estimation errors, and generated hints.
- `plan_repo.norm_queries`, `plan_repo.raw_queries`: normalized and original query text.
- `plan_repo.plan_history_pretty`: a rounded, display-oriented history view.
- `plan_repo.get_hint(bigint)`: returns hints that reproduce a `pg_store_plans` plan ID.
- `plan_repo.get_extstat(bigint)`: suggests extended-statistics SQL when PostgreSQL 14+ and `pg_qualstats` are configured.
- `pg_plan_advsr.enabled`, `pg_plan_advsr.quieted`, `pg_plan_advsr.widely`: enable analysis, suppress messages, or allow hint generation from plans without ANALYZE.

### Safety and Limitations

`EXPLAIN ANALYZE` executes the statement. Do not point the workflow at writes or destructive SQL unless their effects are isolated and intended. Feedback changes hint-table state and intermediate plans can be worse; capture a baseline and stop after a measured result. The supplied `sql/base.sql` clears hint and history data and must not be run against a shared environment.

Concurrent execution is unsupported. InitPlans, SubPlans, Append, MergeAppend, base-relation estimate fixes, grouping/expression statistics suggestions, and pre-14 extended-statistics suggestions are unsupported; parallel query, partitioned tables, and JIT were not tested in the reviewed revision. Treat generated hints and statistics statements as proposals requiring independent plan and runtime verification.
