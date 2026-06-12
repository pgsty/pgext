## 用法

来源：[README](https://github.com/timescale/timescaledb/blob/main/README.md)、[TimescaleDB 2.27.2 release](https://github.com/timescale/timescaledb/releases/tag/2.27.2)、[2.27.2 changelog](https://github.com/timescale/timescaledb/blob/2.27.2/CHANGELOG.md)、[CREATE TABLE API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_table/)、[create_hypertable() API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_hypertable/)、[continuous aggregate API](https://www.tigerdata.com/docs/reference/timescaledb/continuous-aggregates/create_materialized_view/)、[add_columnstore_policy() API](https://www.tigerdata.com/docs/reference/timescaledb/hypercore/add_columnstore_policy/)、[GUCs](https://www.tigerdata.com/docs/reference/timescaledb/configuration/gucs/)

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

- 本项目 CSV 跟踪 TimescaleDB `2.27.2`，覆盖 PostgreSQL 15-18。
- TimescaleDB 2.27.0 增加 Hypercore columnstore 性能工作：vectorized filters、用于 `UPDATE`/`DELETE` equality predicates 的 bloom-filter pruning，以及用于 `UPSERT` 的 bloom-filter pruning。
- 2.27.1 和 2.27.2 bugfix releases 修复 columnstore 与 catalog 边界问题，包括 grouped columnar scans 的错误结果或崩溃、当 quals 含 subplans 或不支持的 grouping forms 时跳过 `ColumnarIndexScan`、composite bloom-filter migration scripts、孤立的 compression settings，以及 job/policy information leaks。
- 2.27.0 release notes 为受影响的 compressed `int2` bloom sparse indexes 和 v2.26 生成的 composite bloom metadata 列出 backward-incompatible upgrade caveats；2.27.1 增加 composite bloom-filter migration scripts。
- 2.27.0 release notes 宣布，计划中的 2026 年 6 月 TimescaleDB release 将是最后一个支持 PostgreSQL 15 的版本。
