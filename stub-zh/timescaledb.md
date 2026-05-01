## 用法

来源：[README](https://github.com/timescale/timescaledb/blob/main/README.md)，[TimescaleDB 2.26.4 release](https://github.com/timescale/timescaledb/releases/tag/2.26.4)，[create_hypertable() API](https://docs.timescale.com/api/latest/hypertable/create_hypertable/)，[CREATE TABLE hypertable API](https://docs.timescale.com/api/latest/hypertable/create_table/)，[continuous aggregates guide](https://docs.timescale.com/use-timescale/latest/continuous-aggregates/create-a-continuous-aggregate/)，[add_job() API](https://docs.timescale.com/api/latest/jobs-automation/add_job/)，[add_columnstore_policy() API](https://docs.timescale.com/api/latest/hypercore/add_columnstore_policy/)

`timescaledb` 是 PostgreSQL 的时序与事件分析扩展。当前文档重点强调 hypertables、continuous aggregates、automation jobs，以及将较旧 chunks 移入 columnstore。

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
- TimescaleDB 2.13 及以后，real-time aggregates 默认关闭，除非另行配置。

### Columnstore

```sql
ALTER TABLE ts_test SET (
  timescaledb.enable_columnstore,
  timescaledb.orderby = 'ts DESC'
);

CALL add_columnstore_policy('ts_test', after => INTERVAL '1 day');
```

- 文档将 `add_columnstore_policy()` 与 `convert_to_columnstore()` 作为当前 API。
- 较早的压缩函数如 `add_compression_policy()` 已被文档标记为 old API，并由 columnstore 接口取代。

### 注意事项

- TimescaleDB 2.26.4 是 2.26.3 之上的 bug-fix release；release notes 建议升级，但没有为 hypertables、continuous aggregates、jobs 或 columnstore 增加新的 SQL 使用接口。
- 相关 2.26.4 修复涉及 continuous aggregate 规划与刷新、运行时 chunk exclusion、restore 后 background worker 重启、`orderby` 上的 sparse indexes，以及 compressed chunk merge safety。继续使用上述文档 API，不需要因该 patch release 本身修改 SQL。
