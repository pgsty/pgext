## 用法

来源：

- [TimescaleDB v2.30.2 README](https://github.com/timescale/timescaledb/blob/2.30.2/README.md)
- [TimescaleDB 2.29.0 发行说明](https://github.com/timescale/timescaledb/releases/tag/2.29.0)
- [TimescaleDB 2.30.2 发布说明](https://github.com/timescale/timescaledb/releases/tag/2.30.2)
- [2.30.2 版变更日志](https://github.com/timescale/timescaledb/blob/2.30.2/CHANGELOG.md)
- [TimescaleDB v2.30.2 控制文件](https://github.com/timescale/timescaledb/blob/2.30.2/timescaledb.control.in)
- [CREATE TABLE API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_table/)
- [create_hypertable() API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_hypertable/)
- [连续聚合 API](https://www.tigerdata.com/docs/reference/timescaledb/continuous-aggregates/create_materialized_view/)
- [add_columnstore_policy() API](https://www.tigerdata.com/docs/reference/timescaledb/hypercore/add_columnstore_policy/)
- [TimescaleDB GUC](https://www.tigerdata.com/docs/reference/timescaledb/configuration/gucs/)

`timescaledb` 是用于时序与事件分析的 PostgreSQL 扩展。当前文档重点介绍 `CREATE TABLE ... WITH (tsdb.hypertable)`、连续聚合、自动化作业，以及将数据块迁移到列存储的用法。

### 启用扩展

将 `timescaledb` 加入现有预加载列表，重启 PostgreSQL 后再创建扩展。版本 `2.30.2` 支持 PostgreSQL 16、17 和 18。

```conf
shared_preload_libraries = 'timescaledb'
```

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

要转换现有 PostgreSQL 表，请使用通用超表 API：

```sql
CREATE TABLE ts_existing (
  ts timestamptz NOT NULL,
  id bigint,
  v integer
);
SELECT create_hypertable('ts_existing', by_range('ts'));
```

- `CREATE TABLE ... WITH (tsdb.hypertable)` 自 TimescaleDB 2.20.0 起已有文档说明，是新建超表的最佳实践。
- 在 TimescaleDB 2.23.0 及更高版本中，系统会自动选择第一个 `TIMESTAMP` 或 `TIMESTAMPTZ` 列作为分区列；如果存在多个候选列而导致选择有歧义，则不会自动选择。
- `create_hypertable()` 仍可用于转换现有表。

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

```

- 连续聚合要求在超表的时间维度上使用 `time_bucket(...)`。
- 连续聚合的 `WITH` 子句支持 `timescaledb.materialized_only`；当前 API 默认值为 `TRUE`，因此除非另行配置，否则不会启用实时聚合。
- TimescaleDB 2.28.0 支持以批次增量执行手工 `refresh_continuous_aggregate()` 调用。可使用 `buckets_per_batch`、`max_batches_per_execution` 和 `refresh_newest_first` 将大型手工刷新拆分为较小的工作单元。
- TimescaleDB 2.28.0 还允许通过 `ALTER MATERIALIZED VIEW ... ADD COLUMN ... GENERATED ALWAYS AS (...) STORED`，向现有连续聚合添加新的生成聚合列；在刷新之前，现有行的该列为 `NULL`。

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

- `CREATE TABLE ... WITH (tsdb.hypertable)` 默认启用列存储，除非设置 `tsdb.columnstore = false`。
- `add_columnstore_policy()` 取代了旧的 `add_compression_policy()` API，且要求在 `after` 和 `created_before` 中二选一，不能同时指定。
- 新建列存储数据块默认启用布隆过滤器。现有数据块必须重新压缩，之后才会拥有布隆索引。

### 相关 GUC

```sql
SET timescaledb.enable_direct_compress_insert = on;
SET timescaledb.enable_cagg_rewrites = on;
SET timescaledb.enable_columnar_scan_filter_pushdown = on;
```

`timescaledb.enable_direct_compress_insert` 和 `timescaledb.enable_direct_compress_copy` 可在写入期间启用技术预览版的直接压缩。TimescaleDB 2.27.0 新增了 `timescaledb.enable_cagg_rewrites` 与 `timescaledb.cagg_rewrites_debug_info`，并说明 `timescaledb.enable_columnar_scan_filter_pushdown` 默认启用。

### 版本 2.30.2 与升级

- TimescaleDB 2.29 支持 PostgreSQL 16、17 和 18。PostgreSQL 15 支持止于 2.28 系列，因此在将 PG15 数据库迁移到 2.29 之前，应先升级 PostgreSQL。
- 版本 2.29.0 新增 `compact_chunk()` 以及用于合并小型列存储批次的压实策略，并优化了 DML 数据块排除和小 `LIMIT` 列存储扫描。在现有工作负载上启用压实策略之前，请先审阅发行说明。
- 2.29 系列新增了 `alter_job(..., config_merge => ...)`、直接压缩与无序重压缩控制，以及分层连续聚合的并发刷新策略。
- 2.30 系列新增面向 LIMIT 查询的 `DeferredChunkAppend`，并支持并发压实期间的 DML。2.30.2 修复了数据块合并与压缩查询崩溃、删除模式后遗留压缩块，以及非确定性文本排序规则处理问题。细粒度刷新选项现统一使用 `timescaledb.cagg_granular_refresh_*` 前缀；应核对已有相关配置。
- 控制文件将 `timescaledb` 标记为受信任且不可重定位。服务器库仍须根据打包部署配置进行预加载，并重启 PostgreSQL。

安装匹配的软件包后，使用新的 `psql -X` 连接执行迁移，避免启动命令提前加载旧库。在每个数据库更新并核对 SQL 版本：

```sql
ALTER EXTENSION timescaledb UPDATE TO '2.30.2';
SELECT extversion FROM pg_extension WHERE extname = 'timescaledb';
```
