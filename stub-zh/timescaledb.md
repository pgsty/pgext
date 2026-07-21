## 用法

来源：

- [TimescaleDB v2.28.3 README](https://github.com/timescale/timescaledb/blob/2.28.3/README.md)
- [TimescaleDB 2.28.3 发行版](https://github.com/timescale/timescaledb/releases/tag/2.28.3)
- [2.28.3 变更日志](https://github.com/timescale/timescaledb/blob/2.28.3/CHANGELOG.md)
- [CREATE TABLE API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_table/)
- [create_hypertable() API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_hypertable/)
- [连续聚合API](https://www.tigerdata.com/docs/reference/timescaledb/continuous-aggregates/create_materialized_view/)
- [add_columnstore_policy() API](https://www.tigerdata.com/docs/reference/timescaledb/hypercore/add_columnstore_policy/)
- [TimescaleDB GUCs](https://www.tigerdata.com/docs/reference/timescaledb/configuration/gucs/)

`timescaledb` 是一个用于时间序列和事件分析的PostgreSQL扩展。当前文档强调了 `CREATE TABLE ... WITH (tsdb.hypertable)`、连续聚合、自动化任务以及将数据块移入列存储。

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

要转换现有的PostgreSQL表，请使用通用超表API：

```sql
CREATE TABLE ts_existing (
  ts timestamptz NOT NULL,
  id bigint,
  v integer
);
SELECT create_hypertable('ts_existing', by_range('ts'));
```

- `CREATE TABLE ... WITH (tsdb.hypertable)` 从TimescaleDB 2.20.0版本开始被文档化，是创建新超表的最佳实践路径。
- 对于TimescaleDB 2.23.0及更高版本，第一个`TIMESTAMP`或`TIMESTAMPTZ`列将自动选择为分区列，除非多个候选者使选择变得模糊。
- `create_hypertable()` 仍然可以用于转换现有表。

### 连续聚合和任务

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

- 连续聚合需要在超表的时间维度上使用`time_bucket(...)`。
- 连续聚合的`WITH`子句支持 `timescaledb.materialized_only`；当前API默认值为`TRUE`，因此除非另行配置，否则实时聚合不会启用。
- TimescaleDB 2.28.0允许手动调用 `refresh_continuous_aggregate()` 分批增量执行。使用 `buckets_per_batch`、`max_batches_per_execution` 和 `refresh_newest_first` 来将大型手动刷新拆分为较小的工作单元。
- TimescaleDB 2.28.0还允许通过 `ALTER MATERIALIZED VIEW ... ADD COLUMN ... GENERATED ALWAYS AS (...) STORED` 向现有连续聚合添加新的生成聚合列；现有行在刷新之前为`NULL`。

### 列存储

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

- 除非 `CREATE TABLE ... WITH (tsdb.hypertable)`，否则默认情况下 `tsdb.columnstore = false` 启用列存储。
- `add_columnstore_policy()` 替换了较旧的 `add_compression_policy()` API，并需要选择 `after` 或 `created_before` 中的一个，而不是两个。

### 重要GUCs

```sql
SET timescaledb.enable_direct_compress_insert = on;
SET timescaledb.enable_cagg_rewrites = on;
SET timescaledb.enable_columnar_scan_filter_pushdown = on;
```

`timescaledb.enable_direct_compress_insert` 和 `timescaledb.enable_direct_compress_copy` 在数据导入期间启用技术预览版的直接压缩。TimescaleDB 2.27.0 添加了 `timescaledb.enable_cagg_rewrites` 和 `timescaledb.cagg_rewrites_debug_info`，并记录了 `timescaledb.enable_columnar_scan_filter_pushdown` 默认启用。

### 版本2.28.3及相关注意事项

- 使用 `2.28.3` 而不是早期的 `2.28.x` 构建。它包含列存储中 `NULL` 处理、排序负常量和排序值、`stddev`/`avg` 结果、非默认排序规则下的压缩数据DML操作、压缩竞争以及涉及数组和 `NULL` 的直接删除情况的正确性修复。
- TimescaleDB 2.28.0通过从列存储批处理元数据中推导结果而不是解压批处理来加快对压缩数据的 `first(value, time)` 和 `last(value, time)` 聚合。
- 2.28.0中的列存储执行器可以在压缩数据上评估 `CASE ... WHEN` 表达式，保持条件聚合和计算表达式的向量化路径。
- TimescaleDB 2.28.0移除了自适应分块，并删除了 `_timescaledb_catalog.chunk_constraint`，暂时用兼容视图替换。使用稳定的信息视图而不是依赖于该系统目录对象。
- 2.28.x是支持PostgreSQL 15的最后一个次要系列；计划中的下一个次要版本仅支持PostgreSQL 16、17和18。
