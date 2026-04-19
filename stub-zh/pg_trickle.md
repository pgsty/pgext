
## 用法

> 来源：[README v0.20.0](https://github.com/grove/pg-trickle/blob/v0.20.0/README.md), [v0.20.0 release notes](https://github.com/grove/pg-trickle/releases/tag/v0.20.0)

`pg_trickle` 为 PostgreSQL 18 提供增量维护的 stream table。它会在可能时使用 differential refresh，同时上游也记录了 full recompute 和事务内立即刷新的模式。

### 启用扩展

```sql
-- postgresql.conf
shared_preload_libraries = 'pg_trickle'
max_worker_processes = 8

CREATE EXTENSION pg_trickle;
```

上游 README 说明默认不要求 `wal_level = logical`。CDC 会先通过 row-level triggers 启动；当 `pg_trickle.cdc_mode = 'auto'` 时，可以切换为基于 WAL 的捕获。

### 创建与刷新 Stream Table

```sql
SELECT pgtrickle.create_stream_table(
    'regional_totals',
    'SELECT region, SUM(amount) AS total, COUNT(*) AS cnt
     FROM orders GROUP BY region'
);

SELECT * FROM regional_totals;
SELECT pgtrickle.refresh_stream_table('regional_totals');
```

文档记录的 refresh mode 包括 `AUTO`、`DIFFERENTIAL`、`FULL` 和 `IMMEDIATE`。

```sql
SELECT pgtrickle.create_stream_table(
    'regional_totals_live',
    'SELECT region, SUM(amount) AS total, COUNT(*) AS cnt
     FROM orders GROUP BY region',
    schedule => NULL,
    refresh_mode => 'IMMEDIATE'
);
```

### 运维与观测

```sql
SELECT pgtrickle.alter_stream_table(
    'regional_totals',
    query => 'SELECT region, SUM(amount) AS total FROM orders GROUP BY region'
);

SELECT * FROM pgtrickle.pgt_status();
SELECT * FROM pgtrickle.health_check();
SELECT * FROM pgtrickle.refresh_timeline(20);
SELECT * FROM pgtrickle.change_buffer_sizes();
SELECT * FROM pgtrickle.dependency_tree();
```

README 还特别说明了较广的 SQL 覆盖面，包括 joins、aggregates、window functions、recursive CTEs、subqueries、set operations，以及 TopK queries。

### v0.20.0 监控增强

`v0.20.0` 新增了内建自监控能力：

- `pgtrickle.setup_dog_feeding()`
- `pgtrickle.teardown_dog_feeding()`
- `pgtrickle.dog_feeding_status()`

release notes 说明新增了五个 monitoring stream tables，用于分析 refresh history，并提供阈值建议与告警。
