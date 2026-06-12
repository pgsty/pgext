## 用法

来源：[README v0.81.0](https://github.com/trickle-labs/pg-trickle/blob/v0.81.0/README.md)、[v0.81.0 release notes](https://github.com/trickle-labs/pg-trickle/releases/tag/v0.81.0)、[SQL reference](https://github.com/trickle-labs/pg-trickle/blob/v0.81.0/docs/SQL_REFERENCE.md)、[configuration guide](https://github.com/trickle-labs/pg-trickle/blob/v0.81.0/docs/CONFIGURATION.md)、[GUC catalog](https://github.com/trickle-labs/pg-trickle/blob/v0.81.0/docs/GUC_CATALOG.md)、[Cargo.toml](https://github.com/trickle-labs/pg-trickle/blob/v0.81.0/Cargo.toml)

`pg_trickle` 为 PostgreSQL 18 提供 stream tables：它们是可正常查询的表，内容由定义它们的 SQL query 维护。扩展会在可行时使用 incremental view maintenance，也可以回退到 full recompute，并支持在同一事务内维护的 `IMMEDIATE` 模式。

上游 v0.81.0 仍处于 pre-1.0 阶段，并说明 API 与配置项在稳定的 1.0 release 前仍可能变化。Rust 包名为 `pg_trickle`，版本 `0.81.0`，使用 Rust edition 2024，默认启用 `pg18` feature，并固定 `pgrx = 0.18.0`。README 中的构建前提是 PostgreSQL 18.x，以及带 pgrx 0.18.x 的 Rust 1.85+。

### 启用扩展

将扩展加入 PostgreSQL 启动配置并重启：

```sql
-- postgresql.conf
shared_preload_libraries = 'pg_trickle'
max_worker_processes = 8
```

```sql
CREATE EXTENSION pg_trickle;
```

`shared_preload_libraries` 是必需的，因为扩展会在启动时注册 GUC 和后台 worker。默认不要求 `wal_level = logical` 或 replication slots：`pg_trickle.cdc_mode = 'auto'` 会先使用 trigger-based CDC，只有在 logical WAL 可用时才切换到 WAL-based capture。

### 创建和刷新 Stream Tables

```sql
CREATE TABLE orders (
    id int PRIMARY KEY,
    region text,
    amount numeric
);

SELECT pgtrickle.create_stream_table(
    'regional_totals',
    'SELECT region, SUM(amount) AS total, COUNT(*) AS cnt
     FROM orders GROUP BY region'
);

SELECT * FROM regional_totals;
SELECT pgtrickle.refresh_stream_table('regional_totals');
```

主要刷新模式包括 `AUTO`、`DIFFERENTIAL`、`FULL` 和 `IMMEDIATE`。`AUTO` 会在查询可微分时选择 differential maintenance，必要时回退到 full recompute。`DIFFERENTIAL` 只应用增量。`FULL` 会 truncate 后按定义查询重新加载。`IMMEDIATE` 使用 statement-level IVM triggers，并在 base-table DML 所在的同一事务内维护。

```sql
SELECT pgtrickle.create_stream_table(
    'regional_totals_live',
    'SELECT region, SUM(amount) AS total, COUNT(*) AS cnt
     FROM orders GROUP BY region',
    schedule => NULL,
    refresh_mode => 'IMMEDIATE'
);
```

调度可以接受 `'30s'`、`'5m'`、`'1h'` 这类 duration strings，`'@hourly'` 这类 cron expressions，或从下游依赖继承的默认 `'calculated'` schedule。

```sql
SELECT pgtrickle.create_stream_table(
    name         => 'hourly_totals',
    query        => 'SELECT region, SUM(amount) AS total FROM orders GROUP BY region',
    schedule     => '@hourly',
    refresh_mode => 'FULL'
);
```

### 生命周期、SQL 覆盖与操作符

```sql
SELECT pgtrickle.alter_stream_table(
    'regional_totals',
    query => 'SELECT region, SUM(amount) AS total FROM orders GROUP BY region'
);

SELECT pgtrickle.drop_stream_table('regional_totals');
```

SQL reference 记录了生命周期调用，例如 `pgtrickle.create_stream_table()`、`pgtrickle.create_stream_table_if_not_exists()`、`pgtrickle.create_or_replace_stream_table()`、`pgtrickle.bulk_create()`、`pgtrickle.alter_stream_table()`、`pgtrickle.drop_stream_table()`、`pgtrickle.resume_stream_table()`、`pgtrickle.refresh_stream_table()` 和 `pgtrickle.repair_stream_table()`。

Release `0.81.0` 还记录了常见刷新 profile 的 preset wrappers：

```sql
SELECT pgtrickle.create_stream_table_realtime(
    'regional_totals_rt',
    'SELECT region, SUM(amount) AS total FROM orders GROUP BY region'
);

SELECT pgtrickle.create_stream_table_batch(
    'regional_totals_batch',
    'SELECT region, SUM(amount) AS total FROM orders GROUP BY region'
);

SELECT pgtrickle.create_stream_table_cost_optimized(
    'regional_totals_cost',
    'SELECT region, SUM(amount) AS total FROM orders GROUP BY region'
);
```

已记录的 SQL 覆盖范围包括 joins、aggregates、window functions、set operations、scalar and table subqueries、包含 `WITH RECURSIVE` 的 CTEs、LATERAL/SRFs、`JSON_TABLE`、带 `ORDER BY ... LIMIT` 的 TopK queries、以 views 作为 sources、无 primary keys 的 tables，以及 stream-table dependency DAGs。主要面向用户的 API 不是自定义 SQL operator；用户主要通过 functions、views、catalog tables、GUCs，以及对 stream tables 的普通 SQL 查询交互。

### 运维与自省

```sql
SELECT * FROM pgtrickle.pgt_status();
SELECT * FROM pgtrickle.refresh_timeline(20);
SELECT * FROM pgtrickle.health_check();
SELECT * FROM pgtrickle.health_summary();
SELECT * FROM pgtrickle.pg_stat_stream_tables;
SELECT * FROM pgtrickle.change_buffer_sizes();
SELECT * FROM pgtrickle.dependency_tree();
SELECT * FROM pgtrickle.explain_st('regional_totals');
SELECT * FROM pgtrickle.slot_health();
SELECT * FROM pgtrickle.check_cdc_health();
SELECT * FROM pgtrickle.commit_latency_stats();
SELECT * FROM pgtrickle.tune_recommendations();
SELECT * FROM pgtrickle.preview_stream_table(
    'SELECT region, SUM(amount) FROM orders GROUP BY region'
);
```

其他已记录的 views 和 catalog tables 包括 `pgtrickle.stream_tables_info`、`pgtrickle.quick_health`、`pgtrickle.pgt_cdc_status`、`pgtrickle.pgt_stream_tables`、`pgtrickle.pgt_dependencies`、`pgtrickle.pgt_refresh_history`、`pgtrickle.pgt_change_tracking`、`pgtrickle.pgt_source_gates` 和 `pgtrickle.pgt_refresh_groups`。

### Outbox、Inbox、Relay 与 Snapshots

`pg_trickle` 可以通过 transactional outbox pattern 发布 stream-table deltas，也可以消费幂等的 inbox tables。

```sql
SELECT pgtrickle.enable_outbox('public.regional_totals');
SELECT pgtrickle.create_consumer_group('billing_workers', 'public.regional_totals');
SELECT * FROM pgtrickle.poll_outbox('billing_workers', 'worker-1');
SELECT pgtrickle.commit_offset('billing_workers', 'worker-1', 42);

SELECT pgtrickle.create_inbox('orders_inbox');
SELECT pgtrickle.inbox_health('orders_inbox');
```

SQL reference 还记录了 snapshot 操作与 relay 配置辅助函数：

```sql
SELECT pgtrickle.snapshot_stream_table('public.regional_totals');
SELECT pgtrickle.restore_from_snapshot(
    'public.regional_totals',
    'pgtrickle.regional_totals_snapshot'
);

SELECT pgtrickle.set_relay_outbox(
    'orders-to-nats',
    'public.regional_totals',
    'relay_group_1',
    '{"type": "nats", "subject": "orders.deltas", "url": "nats://nats:4222"}'::jsonb
);
```

### 重要 GUC

v0.81.0 release 记录了 129 个配置参数。常见运维 GUC 包括：

- `pg_trickle.enabled`
- `pg_trickle.cdc_mode`
- `pg_trickle.scheduler_interval_ms`
- `pg_trickle.min_schedule_seconds`
- `pg_trickle.default_schedule_seconds`
- `pg_trickle.max_consecutive_errors`
- `pg_trickle.wal_transition_timeout`
- `pg_trickle.slot_lag_warning_threshold_mb`
- `pg_trickle.slot_lag_critical_threshold_mb`
- `pg_trickle.differential_max_change_ratio`
- `pg_trickle.refresh_strategy`
- `pg_trickle.cost_model_safety_margin`
- `pg_trickle.planner_aggressive`
- `pg_trickle.merge_join_strategy`
- `pg_trickle.merge_strategy`
- `pg_trickle.auto_backoff`
- `pg_trickle.tiered_scheduling`
- `pg_trickle.cleanup_use_truncate`
- `pg_trickle.block_source_ddl`
- `pg_trickle.buffer_alert_threshold`
- `pg_trickle.compact_threshold`
- `pg_trickle.max_buffer_rows`
- `pg_trickle.auto_index`
- `pg_trickle.aggregate_fast_path`
- `pg_trickle.template_cache`
- `pg_trickle.buffer_partitioning`
- `pg_trickle.ivm_topk_max_limit`
- `pg_trickle.ivm_recursive_max_depth`
- `pg_trickle.parallel_refresh_mode`
- `pg_trickle.max_dynamic_refresh_workers`
- `pg_trickle.max_concurrent_refreshes`
- `pg_trickle.worker_pool_size`
- `pg_trickle.merge_batch_size`
- `pg_trickle.change_buffer_schema`
- `pg_trickle.foreign_table_polling`
- `pg_trickle.matview_polling`
- `pg_trickle.log_delta_sql`
- `pg_trickle.metrics_port`
- `pg_trickle.outbox_enabled`
- `pg_trickle.inbox_enabled`
- `pg_trickle.citus_st_lock_lease_ms`
- `pg_trickle.citus_worker_retry_ticks`
- `pg_trickle.enable_vector_agg`
- `pg_trickle.enable_trace_propagation`
- `pg_trickle.otel_endpoint`
- `pg_trickle.trace_id`
- `pg_trickle.cdc_capture_mode`
- `pg_trickle.commit_timestamp_tracking`
- `pg_trickle.l1_cache_max_entries`
- `pg_trickle.self_heal_oom`
- `pg_trickle.self_heal_lock_timeout`

`pg_trickle.event_driven_wake` 和 `pg_trickle.wake_debounce_ms` 为升级兼容而保留，但已正式废弃且没有效果，因为 PostgreSQL background workers 不能使用 `LISTEN`；scheduler 使用基于 latch 的 polling。

### v0.81.0 运维说明

v0.81.0 release 增加了面向 operator 的自省与调优辅助函数，包括 `pgtrickle.commit_latency_stats()`、`pgtrickle.tune_recommendations()` 和 `pgtrickle.preview_stream_table(query text)`。它还通过 `pg_trickle.l1_cache_max_entries` 增加 bounded LRU DVM caches，新增 `pg_trickle.merge_batch_size` GUC，并为 out-of-memory 和 lock-timeout 场景增加 self-healing circuit-breaker settings。

release notes 说明这些新代码路径不需要 schema migration；替换扩展二进制后，现有安装使用 `ALTER EXTENSION pg_trickle UPDATE` 升级。

### 注意事项

- `pg_trickle` v0.81.0 仅支持 PostgreSQL 18；release packages 以 `pg18` 命名，Cargo 默认启用 `pg18` pgrx feature。
- Pigsty builds 使用 `pgrx` 0.18.0；重新构建包时需要防止 pgrx schema metadata 被 linker garbage collection 移除。
- 扩展 control file 标记为 `superuser = true` 且 `trusted = false`。
- 不允许对 stream tables 直接执行 DML，因为其内容由 refresh engine 管理。
- `IMMEDIATE` 模式绕过 CDC，使用 statement-level IVM triggers；WAL CDC 是异步的，不能与事务内维护兼容。
- `DIFFERENTIAL` 模式中的 materialized views 需要 `pg_trickle.matview_polling = on`；`FULL` 模式不需要该 snapshot-comparison 路径。
- stream-table 定义中没有 `ORDER BY` 的 `LIMIT` 或 `OFFSET` 会被拒绝；TopK 应使用 `ORDER BY ... LIMIT`。
- 根据 `pg_trickle.volatile_function_policy`，定义查询默认会拒绝 volatile functions。
