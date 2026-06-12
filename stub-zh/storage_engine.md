## 用法

来源：[README v2.4.0](https://github.com/saulojb/storage_engine/blob/v2.4.0/README.md)、[release v2.4.0](https://github.com/saulojb/storage_engine/releases/tag/v2.4.0)、[PGXN 2.4.0](https://pgxn.org/dist/storage_engine/2.4.0/)、[current README](https://github.com/saulojb/storage_engine/blob/main/README.md)

`storage_engine` 2.4.0 在 `engine` schema 中提供两个 PostgreSQL table access methods：

- `colcompress`：面向列的压缩存储，支持 vectorized filtering、vectorized aggregation、parallel scans，以及 stripe/chunk min/max pruning。
- `rowcompress`：面向 row-batch compression，支持 parallel scans、index scans 和 batch metadata。

```sql
CREATE EXTENSION storage_engine;
```

### 快速开始

使用任一 access method 创建表。版本 2.2 及之后可在 `CREATE TABLE ... WITH (...)` 中直接传入 per-table options。

```sql
CREATE TABLE events (
  ts timestamptz NOT NULL,
  user_id bigint,
  event_type text,
  value float8
) USING colcompress
  WITH (compression = 'zstd', compression_level = 9, orderby = 'ts ASC');

CREATE TABLE logs (
  id bigserial,
  logged_at timestamptz NOT NULL,
  message text
) USING rowcompress
  WITH (batch_size = 10000, compression = 'zstd');

SELECT event_type, count(*), avg(value)
FROM events
WHERE ts > now() - interval '1 day'
GROUP BY 1;
```

版本 2.4 保留了 2.3 的 vectorized-aggregation 工作，并为 TPC-H Q7/Q18/Q20/Q21 以及 Q9 风格的 post-join aggregates 增加 planner-hook 改进。它还为重复的 `rowcompress` index probes 和重复的 `colcompress` scans 增加 backend-local reread caches。

### 主要调优项

上游文档记录的 session-level GUC 包括：

- `storage_engine.compression`
- `storage_engine.compression_level`
- `storage_engine.stripe_row_limit`
- `storage_engine.chunk_group_row_limit`
- `storage_engine.enable_parallel_execution`
- `storage_engine.min_parallel_processes`
- `storage_engine.enable_vectorization`
- `storage_engine.enable_vectorized_groupagg`
- `storage_engine.enable_automatic_plan`
- `storage_engine.enable_dml`
- `storage_engine.enable_custom_scan`
- `storage_engine.enable_qual_pushdown`
- `storage_engine.qual_pushdown_correlation_threshold`
- `storage_engine.max_custom_scan_paths`
- `storage_engine.enable_engine_index_scan`
- `storage_engine.enable_column_cache`
- `storage_engine.column_cache_size`
- `storage_engine.debug_vectorized_groupagg_fallback`
- `storage_engine.planner_debug_level`
- `storage_engine.maintenance_auto_enabled`
- `storage_engine.maintenance_auto_naptime`
- `storage_engine.maintenance_auto_database`

README 说明这些 GUC 会在 shared library 加载后可见；如果希望它们在每个会话中立即可用，或需要内置 maintenance background worker，请将 `storage_engine` 加入 `shared_preload_libraries`。

### 类型与操作符

`engine.uint8` 为 `colcompress` 工作负载存储 unsigned 64-bit values，适用于需要完整 `0` 到 `2^64 - 1` 范围的场景。上游记录了 comparison operators（`=`、`<>`、`<`、`<=`、`>`、`>=`）、B-tree 和 hash opclasses、与 `bigint`、`numeric`、`text` 的双向 casts，以及 `engine.min`、`engine.max`、`engine.sum` aggregates。vectorized planner 可在 `colcompress` 表上派发 `engine.vmin`、`engine.vmax` 和 `engine.vsum`。

### 常用管理函数

针对 `colcompress` 表：

```sql
SELECT engine.alter_colcompress_table_set(
  'events'::regclass,
  orderby => 'ts ASC, user_id ASC',
  compression => 'zstd',
  compression_level => 9
);

SELECT engine.colcompress_merge('events');
CALL engine.colcompress_repack('events');
CALL engine.colcompress_repack('events', min_fill_ratio => 0.7);
CALL engine.colcompress_merge_incremental('events', max_stripes => 64);
CALL engine.smart_update(
  'events'::regclass,
  'value = value * 1.1',
  'event_type = ''purchase'''
);
```

批量加载后，如果 `orderby` key 需要全局排序以便 pruning，请使用 `engine.colcompress_merge()`。使用 `CALL engine.colcompress_repack()` 压缩低填充率 stripes；使用 `CALL engine.colcompress_merge_incremental()` 以较低锁级别分批处理 dirty stripes。

针对 `rowcompress` 表：

```sql
SELECT engine.alter_rowcompress_table_set(
  'logs'::regclass,
  batch_size => 10000,
  compression => 'zstd',
  compression_level => 5
);

SELECT engine.rowcompress_repack('logs');
CALL engine.rowcompress_merge_incremental('logs', max_batches => 128);
SELECT * FROM engine.rowcompress_scan_stats();
```

运维视图包括 `engine.colcompress_options`、`engine.colcompress_stripes`、`engine.rowcompress_options`、`engine.rowcompress_batches` 和 `engine.storage_health`。`engine.storage_maintenance_recommendation(table)` 会返回单表 health metrics 与推荐操作；`CALL engine.storage_maintenance_auto(...)` 可手工或通过内置 background worker 分发维护任务。在 v2.4 中，重复的 `rowcompress` probes 可以复用 backend-local metadata 和 decompressed batches，`engine.rowcompress_scan_stats()` 也能更可靠地报告这些 cache 效果。

### 何时使用哪种 AM

- 对于 projection、vectorization 和 stripe/chunk pruning 能带来收益的分析扫描、聚合和 range predicates，使用 `colcompress`。
- 对于 append-heavy logs 或通常一起读取的宽行，如果 compression 比列投影更重要，使用 `rowcompress`。
- 对于 `colcompress` 上的 point lookups，使用 per-table `index_scan => true` 或 session-level `storage_engine.enable_engine_index_scan = on`；对于分析型 range scans，优先使用 `index_scan => false`，并配合 `engine.colcompress_merge()` 和 `orderby` key。

### 注意事项

- 本仓库打包版本为 `2.4.0`，覆盖 PostgreSQL 15 到 18。上游 v2.4 validation 也覆盖 PostgreSQL 19 devel，但 PG19 不在本仓库 package matrix 中。PostgreSQL 12、13 和 14 用户应停留在上游 1.3.4。
- 此 stub 遵循 `extension.csv` 与 v2.4.0 release/PGXN docs。
- 现有安装使用 `ALTER EXTENSION storage_engine UPDATE TO '2.4.0';` 升级。
- `colcompress` 和 `rowcompress` 不支持 foreign keys 或 `AFTER ROW` triggers。
- 这些 table access methods 不能使用 `pg_repack`。`engine.colcompress_repack()` 会获取 `AccessExclusiveLock`，因此大表应安排在维护窗口执行；对 dirty stripes 或 batches，incremental merge procedures 是较低锁级别选项。
- 不支持 `VACUUM FULL`、`CLUSTER` 和 `CREATE UNLOGGED TABLE ... USING colcompress`；上游建议改用扩展自带的 merge/repack 函数。
- 在 `colcompress` 上，将 `orderby` 与 B-tree indexes 组合可能禁用 sort-on-write 路径；ordered columns 上的 B-tree indexes 可能削弱 range queries 的 stripe pruning。全局排序重要时，在加载数据后使用 `engine.colcompress_merge()`；分析型表优先使用 `index_scan => false`。
- 如果同时 preload `citus` 或 `pg_cron`，上游文档给出的加载顺序是 `shared_preload_libraries = 'pg_cron,citus,storage_engine'`；`citus` 必须出现在 `storage_engine` 之前。
