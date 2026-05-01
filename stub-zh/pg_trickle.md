## 用法

来源：[README v0.40.0](https://github.com/grove/pg-trickle/blob/v0.40.0/README.md)，[v0.40.0 release notes](https://github.com/grove/pg-trickle/releases/tag/v0.40.0)，[SQL reference](https://github.com/grove/pg-trickle/blob/v0.40.0/docs/SQL_REFERENCE.md)，[configuration guide](https://github.com/grove/pg-trickle/blob/v0.40.0/docs/CONFIGURATION.md)，[GUC catalog](https://github.com/grove/pg-trickle/blob/v0.40.0/docs/GUC_CATALOG.md)，[SQL API catalog](https://github.com/grove/pg-trickle/blob/v0.40.0/docs/SQL_API_CATALOG.md)，[Cargo.toml](https://github.com/grove/pg-trickle/blob/v0.40.0/Cargo.toml)

`pg_trickle` 为 PostgreSQL 18 提供 stream table：它们是普通可查询表，内容由定义 SQL 查询维护。扩展会在可行时使用 incremental view maintenance，也可回退到 full recompute，并支持同一事务内维护的 `IMMEDIATE` 模式。

上游 `v0.40.0` 仍是 pre-1.0，说明 API 与配置项在稳定 1.0 前可能变化。Rust 包为 `pg_trickle` version `0.40.0`，使用 Rust edition 2024，默认启用 `pg18` feature，并固定 `pgrx = 0.18.0`。README 列出的构建前提是 PostgreSQL 18.x，以及 Rust 1.85+ 和 pgrx 0.18.x。

### 启用扩展与构建范围

可在有 release package 时直接安装，也可使用 pgrx 构建：

```bash
cargo pgrx install --release
cargo pgrx package
```

将扩展加入 PostgreSQL 启动配置并重启：

```sql
-- postgresql.conf
shared_preload_libraries = 'pg_trickle'
max_worker_processes = 8
```

```sql
CREATE EXTENSION pg_trickle;
```

`shared_preload_libraries` 是必需的，因为扩展会在启动时注册 GUC 和 background worker。默认不要求 `wal_level = logical` 或 replication slot：`pg_trickle.cdc_mode = 'auto'` 会先使用 trigger-based CDC，只有 logical WAL 可用时才切换到 WAL-based capture。

### 创建与刷新 Stream Table

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

主要 refresh mode 包括 `AUTO`、`DIFFERENTIAL`、`FULL` 和 `IMMEDIATE`。`AUTO` 会在查询可微分时选择 differential maintenance，需要时回退到 full recompute。`DIFFERENTIAL` 只应用 delta。`FULL` 会截断并从定义查询重新加载。`IMMEDIATE` 使用 statement-level IVM trigger，并在基础表 DML 的同一事务内维护。

```sql
SELECT pgtrickle.create_stream_table(
    'regional_totals_live',
    'SELECT region, SUM(amount) AS total, COUNT(*) AS cnt
     FROM orders GROUP BY region',
    schedule => NULL,
    refresh_mode => 'IMMEDIATE'
);
```

调度可使用 `'30s'`、`'5m'`、`'1h'` 等 duration string，`'@hourly'` 等 cron expression，或默认的 `'calculated'`，由下游依赖继承。

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

SQL reference 记录的生命周期调用包括 `pgtrickle.create_stream_table()`、`pgtrickle.create_stream_table_if_not_exists()`、`pgtrickle.create_or_replace_stream_table()`、`pgtrickle.bulk_create()`、`pgtrickle.alter_stream_table()`、`pgtrickle.drop_stream_table()`、`pgtrickle.resume_stream_table()`、`pgtrickle.refresh_stream_table()` 和 `pgtrickle.repair_stream_table()`。

文档覆盖的 SQL 范围包括 joins、aggregates、window functions、set operations、scalar/table subqueries、包含 `WITH RECURSIVE` 的 CTE、LATERAL/SRF、`JSON_TABLE`、带 `ORDER BY ... LIMIT` 的 TopK 查询、以 view 作为来源、无主键表，以及 stream-table dependency DAG。主要面向用户的 API 不是自定义 SQL operator；用户主要通过函数、view、catalog table、GUC，以及对 stream table 的普通 SQL 查询进行交互。

### 运维与观测

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
```

其他已记录的 view 与 catalog table 包括 `pgtrickle.stream_tables_info`、`pgtrickle.quick_health`、`pgtrickle.pgt_cdc_status`、`pgtrickle.pgt_stream_tables`、`pgtrickle.pgt_dependencies`、`pgtrickle.pgt_refresh_history`、`pgtrickle.pgt_change_tracking`、`pgtrickle.pgt_source_gates` 和 `pgtrickle.pgt_refresh_groups`。

### Outbox、Inbox、Relay 与 Snapshot

`pg_trickle` 可通过 transactional outbox pattern 发布 stream-table delta，并消费幂等 inbox table。

```sql
SELECT pgtrickle.enable_outbox('public.regional_totals');
SELECT pgtrickle.create_consumer_group('billing_workers', 'public.regional_totals');
SELECT * FROM pgtrickle.poll_outbox('billing_workers', 'worker-1');
SELECT pgtrickle.commit_offset('billing_workers', 'worker-1', 42);

SELECT pgtrickle.create_inbox('orders_inbox');
SELECT pgtrickle.inbox_health('orders_inbox');
```

SQL reference 还记录了 snapshot 操作和 relay 配置 helper：

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

`v0.40.0` release 新增了覆盖 125 个配置参数的生成文档。常见运维 GUC 包括：

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

`pg_trickle.event_driven_wake` 和 `pg_trickle.wake_debounce_ms` 在 `v0.40.0` 中为升级兼容而保留，但已正式废弃且没有效果，因为 PostgreSQL background worker 不能使用 `LISTEN`；scheduler 使用基于 latch 的 polling。

### v0.40.0 运维说明

`v0.40.0` release 聚焦 operator confidence。它新增了生成的 GUC 与 SQL API catalog、security model 和 secret-handling guide、带端到端测试的 drain-mode runbook、扩展的 Prometheus alert rules、dbt 与 relay parity helper、严格的 unsafe-block CI gating、更清晰的 template-cache 文档、对 `event_driven_wake` 的正式废弃，以及 CI secret scanning。

生成的 SQL API catalog 列出 24 个 SQL-callable `#[pg_extern]` 函数，包括 `pgtrickle.metrics_summary()`、`pgtrickle.outbox_status()`、`pgtrickle.outbox_rows_consumed()`、`pgtrickle.commit_offset()`、`pgtrickle.consumer_heartbeat()`、`pgtrickle.seek_offset()`、`pgtrickle.inbox_health()`、`pgtrickle.inbox_is_my_partition()`、`pgtrickle.snapshot_stream_table()`、`pgtrickle.restore_from_snapshot()`、`pgtrickle.restore_stream_tables()`、`pgtrickle.recommend_schedule()` 和 `pgtrickle.schedule_recommendations()`。

### 注意事项

- `pg_trickle` v0.40.0 仅支持 PostgreSQL 18；release package 命名为 `pg18`，Cargo 默认使用 `pg18` pgrx feature。
- 扩展 control file 标记为 `superuser = true` 且 `trusted = false`。
- 不允许对 stream table 直接执行 DML，因为其内容由 refresh engine 管理。
- `IMMEDIATE` 模式绕过 CDC，使用 statement-level IVM trigger；WAL CDC 是异步的，不能与事务内维护兼容。
- `DIFFERENTIAL` 模式中的 materialized view 需要 `pg_trickle.matview_polling = on`；`FULL` 模式不需要该 snapshot-comparison 路径。
- stream-table 定义中的 `LIMIT` 或 `OFFSET` 如果没有 `ORDER BY` 会被拒绝；TopK 应使用 `ORDER BY ... LIMIT`。
- 根据 `pg_trickle.volatile_function_policy`，定义查询默认会拒绝 volatile function。
