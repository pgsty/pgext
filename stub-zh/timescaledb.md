## 用法

来源：[README](https://github.com/timescale/timescaledb/blob/main/README.md)，[TimescaleDB changelog](https://github.com/timescale/docs/blob/latest/about/changelog.md)，[create_hypertable() API](https://docs.timescale.com/api/latest/hypertable/create_hypertable/)，[CREATE TABLE hypertable API](https://docs.timescale.com/api/latest/hypertable/create_table/)，[continuous aggregates guide](https://docs.timescale.com/use-timescale/latest/continuous-aggregates/create-a-continuous-aggregate/)，[add_job() API](https://docs.timescale.com/api/latest/jobs-automation/add_job/)，[add_columnstore_policy() API](https://docs.timescale.com/api/latest/hypercore/add_columnstore_policy/)

`timescaledb` 是 PostgreSQL 的时序与事件分析扩展。当前文档重点强调 hypertables、continuous aggregates、automation jobs，以及将旧 chunk 移入 columnstore。

### Hypertables

```sql
CREATE EXTENSION timescaledb;

CREATE TABLE ts_test (
  ts timestamptz NOT NULL,
  id bigint,
  v integer
);

SELECT create_hypertable('ts_test', by_range('ts'));
```

- `create_hypertable()` 仍可使用，但 API 文档从 TimescaleDB 2.20.0 起将其标记为 old，并建议新用户转向 `CREATE TABLE ... WITH (...)`。
- 当前 README 也展示了更新模式：`CREATE TABLE ... WITH (tsdb.hypertable)`。

### Continuous aggregates 与 jobs

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

- Continuous aggregates 要求在 hypertable 的时间维度上使用 `time_bucket(...)`。
- 从 TimescaleDB 2.13 开始，real-time aggregates 默认关闭，除非另行配置。

### Columnstore

```sql
ALTER TABLE ts_test SET (
  timescaledb.enable_columnstore,
  timescaledb.orderby = 'ts DESC'
);

CALL add_columnstore_policy('ts_test', after => INTERVAL '1 day');
```

- 文档将 `add_columnstore_policy()` 与 `convert_to_columnstore()` 视为当前 API。
- 较早的压缩函数如 `add_compression_policy()` 已被文档标记为 old API，并由 columnstore 接口取代。

### 注意事项

- 上游 tag 包含 `2.26.3`，但公开 changelog 与可见 release notes 目前主要记录 `2.26.0` 与 `2.26.2` 这条 `2.26` 分支；这些说明主要涉及性能与 bug 修复，而不是根本不同的使用接口。
- `2.26` changelog 强调了更快的 columnstore 查询、更好的压缩数据过滤以及 chunk exclusion 改进，因此这次刷新聚焦当前文档 API，而不是 patch-level 内部细节。
