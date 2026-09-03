


## Usage

Sources:

- [pg_trickle v0.87.16 release](https://github.com/trickle-labs/pg-trickle/releases/tag/v0.87.16)
- [pg_trickle v0.87.16 README](https://github.com/trickle-labs/pg-trickle/blob/v0.87.16/README.md)
- [pg_trickle v0.87.16 SQL reference](https://github.com/trickle-labs/pg-trickle/blob/v0.87.16/docs/SQL_REFERENCE.md)
- [pg_trickle v0.87.16 configuration guide](https://github.com/trickle-labs/pg-trickle/blob/v0.87.16/docs/CONFIGURATION.md)
- [pg_trickle v0.87.16 GUC catalog](https://github.com/trickle-labs/pg-trickle/blob/v0.87.16/docs/GUC_CATALOG.md)
- [pg_trickle v0.87.16 changelog](https://github.com/trickle-labs/pg-trickle/blob/v0.87.16/CHANGELOG.md)
- [pg_trickle v0.87.15 to v0.87.16 upgrade SQL](https://github.com/trickle-labs/pg-trickle/blob/v0.87.16/sql/pg_trickle--0.87.15--0.87.16.sql)
- [pg_trickle v0.87.16 control file](https://github.com/trickle-labs/pg-trickle/blob/v0.87.16/pg_trickle.control)
- [pg_trickle v0.87.16 Cargo metadata](https://github.com/trickle-labs/pg-trickle/blob/v0.87.16/Cargo.toml)

`pg_trickle` v0.87.16 provides stream tables for PostgreSQL 18: regular queryable tables whose contents are maintained from a defining SQL query. It uses incremental view maintenance when possible, can fall back to full recompute, and also supports `IMMEDIATE` mode for same-transaction maintenance.

Upstream v0.87.16 is still pre-1.0 and says APIs and configuration options may change before a stable 1.0 release. The Rust package is `pg_trickle` version `0.87.16`, uses Rust edition 2024, defaults to the `pg18` feature, and pins `pgrx = 0.18.0`. Build prerequisites in the README are PostgreSQL 18.x plus Rust 1.85+ with pgrx 0.18.x.

### Enable the Extension

Add the extension to PostgreSQL startup configuration and restart:

```sql
-- postgresql.conf
shared_preload_libraries = 'pg_trickle'
max_worker_processes = 8
```

```sql
CREATE EXTENSION pg_trickle;
```

`shared_preload_libraries` is required because the extension registers GUCs and a background worker at startup. `wal_level = logical` and replication slots are not required by default: `pg_trickle.cdc_mode = 'auto'` starts with trigger-based CDC and transitions to WAL-based capture only when logical WAL is available.

### Create and Refresh Stream Tables

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

The main refresh modes are `AUTO`, `DIFFERENTIAL`, `FULL`, and `IMMEDIATE`. `AUTO` chooses differential maintenance when the query is differentiable and falls back to full recompute when needed. `DIFFERENTIAL` applies deltas only. `FULL` truncates and reloads from the defining query. `IMMEDIATE` uses statement-level IVM triggers and is maintained inside the same transaction as base-table DML.

```sql
SELECT pgtrickle.create_stream_table(
    'regional_totals_live',
    'SELECT region, SUM(amount) AS total, COUNT(*) AS cnt
     FROM orders GROUP BY region',
    schedule => NULL,
    refresh_mode => 'IMMEDIATE'
);
```

Schedules accept duration strings such as `'30s'`, `'5m'`, `'1h'`, cron expressions such as `'@hourly'`, or the default `'calculated'` schedule inherited from downstream dependents.

```sql
SELECT pgtrickle.create_stream_table(
    name         => 'hourly_totals',
    query        => 'SELECT region, SUM(amount) AS total FROM orders GROUP BY region',
    schedule     => '@hourly',
    refresh_mode => 'FULL'
);
```

### Declare a Freshness Target

Version 0.86.0 adds `target_freshness` to the function-based create, alter, and preview APIs. A positive interval is translated into the existing schedule and deadline controls; `'on_commit'` selects same-transaction `IMMEDIATE` maintenance, and `'manual'` disables scheduled refreshes. Do not pass both a non-default `schedule` and `target_freshness`. Calendar-month intervals are rejected because their duration is not stable, and the release explicitly does not yet provide closed-loop freshness control.

```sql
SELECT pgtrickle.create_stream_table(
    name             => 'fresh_totals',
    query            => 'SELECT region, SUM(amount) AS total FROM orders GROUP BY region',
    target_freshness => '5m'
);

SELECT pgtrickle.alter_stream_table(
    'fresh_totals',
    target_freshness => 'on_commit'
);
```

### Lifecycle, SQL Coverage, and Operators

```sql
SELECT pgtrickle.alter_stream_table(
    'regional_totals',
    query => 'SELECT region, SUM(amount) AS total FROM orders GROUP BY region'
);

SELECT pgtrickle.drop_stream_table('regional_totals');
```

The SQL reference documents lifecycle calls such as `pgtrickle.create_stream_table()`, `pgtrickle.create_stream_table_if_not_exists()`, `pgtrickle.create_or_replace_stream_table()`, `pgtrickle.bulk_create()`, `pgtrickle.alter_stream_table()`, `pgtrickle.drop_stream_table()`, `pgtrickle.resume_stream_table()`, `pgtrickle.refresh_stream_table()`, and `pgtrickle.repair_stream_table()`.

Release `0.81.0` also documents preset wrappers for common refresh profiles:

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

The documented SQL coverage includes joins, aggregates, window functions, set operations, scalar and table subqueries, CTEs including `WITH RECURSIVE`, LATERAL/SRFs, `JSON_TABLE`, TopK queries with `ORDER BY ... LIMIT`, views as sources, tables without primary keys, and stream-table dependency DAGs. No custom SQL operator is the main user-facing API; users primarily interact through functions, views, catalog tables, GUCs, and normal SQL queries over stream tables.

### Operations and Introspection

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
SELECT pgtrickle.explain('regional_totals');
SELECT pgtrickle.explain_json('regional_totals');
SELECT * FROM pgtrickle.pg_stat_pgtrickle;
```

Other documented views and catalog tables include `pgtrickle.stream_tables_info`, `pgtrickle.quick_health`, `pgtrickle.pgt_cdc_status`, `pgtrickle.pgt_stream_tables`, `pgtrickle.pgt_dependencies`, `pgtrickle.pgt_refresh_history`, `pgtrickle.pgt_change_tracking`, `pgtrickle.pgt_source_gates`, and `pgtrickle.pgt_refresh_groups`.

### Outbox, Inbox, Relay, and Snapshots

`pg_trickle` can publish stream-table deltas through the transactional outbox pattern and consume idempotent inbox tables.

```sql
SELECT pgtrickle.enable_outbox('public.regional_totals');
SELECT pgtrickle.create_consumer_group('billing_workers', 'public.regional_totals');
SELECT * FROM pgtrickle.poll_outbox('billing_workers', 'worker-1');
SELECT pgtrickle.commit_offset('billing_workers', 'worker-1', 42);

SELECT pgtrickle.create_inbox('orders_inbox');
SELECT pgtrickle.inbox_health('orders_inbox');
```

The SQL reference also documents snapshot operations and relay configuration helpers:

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

### Important GUCs

The generated catalog exposes many configuration parameters. Common operational GUCs include:

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
- `pg_trickle.pipeline_batch_size`
- `pg_trickle.memory_budget_mb`
- `pg_trickle.load_shed_threshold`
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
- `pg_trickle.explain_annotations`
- `pg_trickle.warn_join_sources`
- `pg_trickle.warn_write_path_overhead_us`
- `pg_trickle.cdc_capture_mode`
- `pg_trickle.commit_timestamp_tracking`
- `pg_trickle.l1_cache_max_entries`
- `pg_trickle.self_heal_oom`
- `pg_trickle.self_heal_lock_timeout`

`pg_trickle.event_driven_wake` and `pg_trickle.wake_debounce_ms` are preserved for upgrade compatibility but are formally deprecated and have no effect, because PostgreSQL background workers cannot use `LISTEN`; the scheduler uses latch-based polling.

### Diagnostics, 0.87 Changes, and Upgrade

`pgtrickle.explain()` returns a bounded text explanation of a stream table's requested and effective refresh mode, pending changes, dominant cost, expected refresh time, lag, next refresh, and most recent FULL fallback reason. `pgtrickle.explain_json()` returns the same snapshot with evidence sources and sample counts for automation. The current implementation reports write-path overhead as unknown; `pg_trickle.warn_write_path_overhead_us` is reserved for compatible sampled trigger statistics and defaults to `0.0`, which disables that warning.

`pgtrickle.pg_stat_pgtrickle` exposes bounded cumulative refresh counts, duration percentiles, current lag, target freshness, fallback details, the last operational error, and the reset timestamp without scanning the full refresh history. A stream-table owner can reset one table's counters with `pgtrickle.stat_reset()`; resetting all tables with `pgtrickle.stat_reset_all()` requires superuser or extension-owner privilege.

```sql
SELECT pgtrickle.explain('regional_totals');
SELECT pgtrickle.explain_json('regional_totals');
SELECT * FROM pgtrickle.pg_stat_pgtrickle;

SELECT pgtrickle.stat_reset(42);
SELECT pgtrickle.stat_reset_all();
```

Set `pg_trickle.explain_annotations = on` to add compact lag and refresh-mode properties to PostgreSQL `EXPLAIN`; it is off by default. Creation and preview warnings now flag always-FULL plans, missing source identity, row-level security, excessive join sources, and other expensive or unsafe query shapes. `pg_trickle.warn_join_sources` defaults to `6`, with `0` disabling that warning.

Version 0.87 bounds ordinary MERGE work through a cursor pipeline, makes `pg_trickle.pipeline_batch_size` canonical, adds memory-budget and load-shedding controls, and hardens lifecycle, snapshot, publication, and refresh execution around the original caller and stream-table owner. Version 0.87.16 carries V2 row identity through new or rebuilt stream-table storage, trigger/WAL capture, refresh, differential and `IMMEDIATE` paths. Existing V1 physical state is rejected fail-closed and is not converted in place.

After installing the v0.87.16 binaries and SQL files, apply the catalog migration in every database:

```sql
ALTER EXTENSION pg_trickle UPDATE TO '0.87.16';
```

The 0.87.16 migration adds probe-version metadata and marks existing identity metadata unknown so the V2 runtime cannot consume V1 relations. It does not recreate those relations; upstream assigns that workflow to 0.87.17. Back up first and follow adjacent upgrade scripts rather than assuming an older multi-version jump is safe. Before crossing 0.87.10, run `pgtrickle.lifecycle_preflight()` and apply its remediation. `pgtrickle.migrate()` is read-only and does not replace `ALTER EXTENSION`.

### Caveats

- `pg_trickle` v0.87.16 is PostgreSQL 18 only; the release packages are named for `pg18`, and Cargo defaults to the `pg18` pgrx feature.
- Upstream Cargo metadata pins `pgrx` 0.18.0; use a compatible pgrx toolchain and keep pgrx schema metadata from linker garbage collection when rebuilding packages.
- The extension control file marks it `superuser = true` and `trusted = false`.
- Direct DML on stream tables is not allowed because their contents are managed by the refresh engine.
- `IMMEDIATE` mode bypasses CDC and uses statement-level IVM triggers; WAL CDC is asynchronous and incompatible with in-transaction maintenance.
- Materialized views in `DIFFERENTIAL` mode require `pg_trickle.matview_polling = on`; `FULL` mode works without that snapshot-comparison path.
- `LIMIT` or `OFFSET` without `ORDER BY` is rejected for stream-table definitions; use `ORDER BY ... LIMIT` for TopK.
- Volatile functions are rejected by default in defining queries according to `pg_trickle.volatile_function_policy`.
