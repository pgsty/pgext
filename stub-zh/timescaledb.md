
## 用法

创建一张表并将其转换为超表（hypertable）

```sql
DROP TABLE IF EXISTS ts_test;
CREATE TABLE ts_test
(
    ts TIMESTAMPTZ NOT NULL,
    id BIGINT,
    v  INTEGER -- 载荷数据
);
SELECT create_hypertable('ts_test', by_range('ts'));

INSERT INTO ts_test
    SELECT now() + (i || ' seconds')::INTERVAL, i, i % 100
    FROM generate_series(1, 1000000) i;
```

连续聚合示例：

```sql
CREATE MATERIALIZED VIEW ts_hourly
WITH (timescaledb.continuous) AS
  SELECT time_bucket('1 hour', ts) AS bucket,
         count(*) AS cnt,
         avg(v)   AS avg_v
  FROM ts_test
  GROUP BY bucket;

-- 添加刷新策略以保持连续聚合的数据是最新的
SELECT add_continuous_aggregate_policy('ts_hourly',
    start_offset    => INTERVAL '3 hours',
    end_offset      => INTERVAL '1 hour',
    schedule_interval => INTERVAL '1 hour');
```

任务调度示例：

```sql
SELECT add_job('SELECT 1','1h', initial_start => now()::timestamptz);
```

压缩示例：

```sql
ALTER TABLE ts_test SET (
    timescaledb.compress,
    timescaledb.compress_segmentby = 'id',
    timescaledb.compress_orderby = 'ts'
);

-- 添加压缩策略，自动压缩超过 1 天的数据块
SELECT add_compression_policy('ts_test', INTERVAL '1 day');
```
