


## 用法

来源：[README](https://github.com/timescale/timescaledb/blob/main/README.md)、[TimescaleDB 2.28.0 发行版](https://github.com/timescale/timescaledb/releases/tag/2.28.0)、[2.28.0 变更日志](https://github.com/timescale/timescaledb/blob/2.28.0/CHANGELOG.md)、[CREATE TABLE API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_table/)、[create_hypertable() API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_hypertable/)、[连续聚合 API](https://www.tigerdata.com/docs/reference/timescaledb/continuous-aggregates/create_materialized_view/)、[add_columnstore_policy() API](https://www.tigerdata.com/docs/reference/timescaledb/hypercore/add_columnstore_policy/)、[GUC](https://www.tigerdata.com/docs/reference/timescaledb/configuration/gucs/)

`timescaledb` 是用于时间序列与事件分析的 PostgreSQL 扩展。当前文档强调 `CREATE TABLE ... WITH (tsdb.hypertable)`、连续聚合、自动化作业，以及将数据块移入列存。

### 超表

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

转换已有 PostgreSQL 表时，使用通用超表 API：

```sql
CREATE TABLE ts_existing (
  ts timestamptz NOT NULL,
  id bigint,
  v integer
);
SELECT create_hypertable('ts_existing', by_range('ts'));
```

- `CREATE TABLE ... WITH (tsdb.hypertable)` 自 TimescaleDB 2.20.0 起已有文档记录，是新建超表的最佳实践路径。
- TimescaleDB 2.23.0 及之后版本中，除非存在多个候选列导致选择不明确，否则第一个 `TIMESTAMP` 或 `TIMESTAMPTZ` 列会自动选为分区列。
- `create_hypertable()` 仍可用于转换已有表。

### 连续聚合与作业

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

- 连续聚合要求在超表的时间维度上使用 `time_bucket(...)`。
- 连续聚合的 `WITH` 子句支持 `timescaledb.materialized_only`；当前 API 默认值为 `TRUE`，因此除非另行配置，否则不会启用实时聚合。
- TimescaleDB 2.28.0 允许手动 `refresh_continuous_aggregate()` 调用以增量批次执行。可使用 `buckets_per_batch`、`max_batches_per_execution` 和 `refresh_newest_first` 将大型手动刷新拆成更小的工作单元。
- TimescaleDB 2.28.0 还允许通过 `ALTER MATERIALIZED VIEW ... ADD COLUMN ... GENERATED ALWAYS AS (...) STORED` 向已有连续聚合添加新的生成聚合列；已有行在刷新前为 `NULL`。

### 列存

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

- `CREATE TABLE ... WITH (tsdb.hypertable)` 默认启用列存，除非设置 `tsdb.columnstore = false`。
- `add_columnstore_policy()` 替代较旧的 `add_compression_policy()` API，并要求 `after` 或 `created_before` 二选一，不能同时使用。
- 新的列存数据块默认启用布隆过滤器。已有数据块需要重新压缩后才会拥有布隆索引。

### 相关 GUC

```sql
SET timescaledb.enable_direct_compress_insert = on;
SET timescaledb.enable_cagg_rewrites = on;
SET timescaledb.enable_columnar_scan_filter_pushdown = on;
```

`timescaledb.enable_direct_compress_insert` 和 `timescaledb.enable_direct_compress_copy` 会启用数据摄取期间的技术预览版直接压缩。TimescaleDB 2.27.0 增加 `timescaledb.enable_cagg_rewrites` 和 `timescaledb.cagg_rewrites_debug_info`，并记录 `timescaledb.enable_columnar_scan_filter_pushdown` 默认启用。

### 注意事项

- 本项目 CSV 跟踪 TimescaleDB `2.28.0`，覆盖 PostgreSQL 15-18。
- TimescaleDB 2.28.0 通过从列存批次元数据派生结果，而不是解压批次，加速压缩数据上的 `first(value, time)` 和 `last(value, time)` 聚合。
- 2.28.0 的列式执行器可以在压缩数据上计算 `CASE ... WHEN` 表达式，使条件聚合和计算表达式保持在向量化路径上。
- TimescaleDB 2.28.0 移除了自适应分块，并删除 `_timescaledb_catalog.chunk_constraint`，临时以兼容视图替代。不要依赖该目录对象，改用稳定的信息视图。
- TimescaleDB 2.28.x 是最后一个支持 PostgreSQL 15 的次要版本系列；下一条计划中的次要版本线只支持 PostgreSQL 16、17 和 18。
