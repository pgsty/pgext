


## 用法

来源：[README](https://github.com/timescale/timescaledb/blob/main/README.md)、[TimescaleDB 2.28.0 release](https://github.com/timescale/timescaledb/releases/tag/2.28.0)、[2.28.0 changelog](https://github.com/timescale/timescaledb/blob/2.28.0/CHANGELOG.md)、[CREATE TABLE API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_table/)、[create_hypertable() API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_hypertable/)、[continuous aggregate API](https://www.tigerdata.com/docs/reference/timescaledb/continuous-aggregates/create_materialized_view/)、[add_columnstore_policy() API](https://www.tigerdata.com/docs/reference/timescaledb/hypercore/add_columnstore_policy/)、[GUCs](https://www.tigerdata.com/docs/reference/timescaledb/configuration/gucs/)

`timescaledb` 是用于 time-series 与 event analytics 的 PostgreSQL 扩展。当前文档强调 `CREATE TABLE ... WITH (tsdb.hypertable)`、continuous aggregates、automation jobs，以及将 chunks 移入 columnstore。

### Hypertables

```sql
CREATE EXTENSION timescaledb;

CREATE TABLE ts_test (
  ts timestamptz NOT NULL,
  id bigint,
  v integer
) WITH (
  tsdb.hypertable,
  tsdb.orderby = 'ts DESC'
);
```

转换已有 PostgreSQL 表时，使用 generalized hypertable API：

```sql
CREATE TABLE ts_existing (
  ts timestamptz NOT NULL,
  id bigint,
  v integer
);
SELECT create_hypertable('ts_existing', by_range('ts'));
```

- `CREATE TABLE ... WITH (tsdb.hypertable)` 自 TimescaleDB 2.20.0 起已有文档记录，是新建 hypertables 的 best-practice 路径。
- TimescaleDB 2.23.0 及之后版本中，除非存在多个候选列导致选择 ambiguous，否则第一个 `TIMESTAMP` 或 `TIMESTAMPTZ` 列会自动选为 partition column。
- `create_hypertable()` 仍可用于转换已有表。

### Continuous Aggregates 与 Jobs

```sql
CREATE MATERIALIZED VIEW ts_hourly
WITH (timescaledb.continuous) AS
SELECT time_bucket('1 hour', ts) AS bucket,
       count(*) AS cnt,
       avg(v)   AS avg_v
FROM ts_test
GROUP BY bucket;

SELECT add_continuous_aggregate_policy(
  'ts_hourly',
  start_offset => INTERVAL '3 hours',
  end_offset => INTERVAL '1 hour',
  schedule_interval => INTERVAL '1 hour'
);

SELECT add_job('user_defined_action', '1h');
```

- Continuous aggregates 要求在 hypertable 的 time dimension 上使用 `time_bucket(...)`。
- Continuous aggregate 的 `WITH` 子句支持 `timescaledb.materialized_only`；当前 API 默认值为 `TRUE`，因此除非另行配置，否则不会启用 real-time aggregation。
- TimescaleDB 2.28.0 允许手动 `refresh_continuous_aggregate()` 调用以增量批次执行。可使用 `buckets_per_batch`、`max_batches_per_execution` 和 `refresh_newest_first` 将大型手动刷新拆成更小的工作单元。
- TimescaleDB 2.28.0 还允许通过 `ALTER MATERIALIZED VIEW ... ADD COLUMN ... GENERATED ALWAYS AS (...) STORED` 向已有 continuous aggregate 添加新的 generated aggregate column；已有行在刷新前为 `NULL`。

### Columnstore

```sql
CREATE TABLE crypto_ticks (
  "time" timestamptz,
  symbol text,
  price double precision,
  day_volume numeric
) WITH (
  tsdb.hypertable,
  tsdb.segmentby = 'symbol',
  tsdb.orderby = 'time DESC'
);

CALL add_columnstore_policy('crypto_ticks', after => INTERVAL '60 days');
```

- `CREATE TABLE ... WITH (tsdb.hypertable)` 默认启用 columnstore，除非设置 `tsdb.columnstore = false`。
- `add_columnstore_policy()` 替代较旧的 `add_compression_policy()` API，并要求 `after` 或 `created_before` 二选一，不能同时使用。
- Bloom filters 对新的 columnstore chunks 默认启用。已有 chunks 需要 recompression 后才会拥有 bloom indexes。

### 相关 GUC

```sql
SET timescaledb.enable_direct_compress_insert = on;
SET timescaledb.enable_cagg_rewrites = on;
SET timescaledb.enable_columnar_scan_filter_pushdown = on;
```

`timescaledb.enable_direct_compress_insert` 和 `timescaledb.enable_direct_compress_copy` 会启用 ingestion 期间 tech-preview direct compression。TimescaleDB 2.27.0 增加 `timescaledb.enable_cagg_rewrites` 和 `timescaledb.cagg_rewrites_debug_info`，并记录 `timescaledb.enable_columnar_scan_filter_pushdown` 默认启用。

### 注意事项

- 本项目 CSV 跟踪 TimescaleDB `2.28.0`，覆盖 PostgreSQL 15-18。
- TimescaleDB 2.28.0 通过从 columnstore batch metadata 派生结果，而不是解压 batches，加速 compressed data 上的 `first(value, time)` 和 `last(value, time)` 聚合。
- 2.28.0 的 columnar executor 可以在 compressed data 上计算 `CASE ... WHEN` 表达式，使 conditional aggregates 和 computed expressions 保持在 vectorized path 上。
- TimescaleDB 2.28.0 移除了 adaptive chunking，并删除 `_timescaledb_catalog.chunk_constraint`，临时以 compatibility view 替代。不要依赖该 catalog object，改用稳定的信息视图。
- TimescaleDB 2.28.x 是最后一个支持 PostgreSQL 15 的 minor series；下一条计划中的 minor line 只支持 PostgreSQL 16、17 和 18。
