
## 用法

TimescaleDB Toolkit 提供了一系列专用于时序数据分析的函数，采用**两步聚合模式**。大多数函数会先创建中间表示，再通过访问器函数进行查询，从而实现高效复用和多维分析。

### 近似分析

#### HyperLogLog - 基数估算

基于概率的去重计数，支持可配置精度，适用于高基数数据集。

```sql
-- 估算每日独立用户数
SELECT
    date_trunc('day', timestamp) as day,
    distinct_count(hyperloglog(64, user_id)) as unique_users
FROM events
GROUP BY day;

-- 跨分区合并计数
SELECT distinct_count(rollup(hll))
FROM (SELECT hyperloglog(32, session_id) as hll FROM events_2023
      UNION ALL
      SELECT hyperloglog(32, session_id) FROM events_2024) t;
```

#### T-Digest - 分位数近似

高精度百分位估算，针对尾部分位数（P95、P99）进行了优化。

```sql
-- 追踪响应时间百分位
SELECT
    service_name,
    approx_percentile(0.50, tdigest(100, response_time)) as p50,
    approx_percentile(0.95, tdigest(100, response_time)) as p95,
    approx_percentile(0.99, tdigest(100, response_time)) as p99
FROM api_metrics
GROUP BY service_name;

-- 结合连续聚合计算每小时百分位
CREATE MATERIALIZED VIEW hourly_percentiles AS
SELECT
    time_bucket('1 hour', timestamp) as hour,
    tdigest(200, response_time) as digest
FROM requests GROUP BY hour;
```

#### UddSketch - 有界误差分位数

带有最大相对误差保证的分位数估算。

```sql
-- 计算 CPU 利用率百分位，最大误差 1%
SELECT
    host_id,
    approx_percentile(0.95, uddsketch(100, 0.01, cpu_percent)) as p95_cpu,
    error(uddsketch(100, 0.01, cpu_percent)) as actual_error
FROM system_metrics
GROUP BY host_id;
```

### 计数器分析

#### 计数器聚合 - 单调递增指标

处理单调递增的计数器，自动检测重置。

```sql
-- 计算请求速率
SELECT
    time_bucket('5 min', timestamp) as bucket,
    rate(counter_agg(timestamp, request_count)) as requests_per_sec,
    delta(counter_agg(timestamp, request_count)) as total_requests
FROM metrics
GROUP BY bucket;

-- 对不完整时间桶进行外推速率计算
SELECT
    extrapolated_rate(
        counter_agg(timestamp, bytes_sent,
                   bounds => time_bucket_range('1 hour', timestamp))
    ) as bytes_per_second
FROM network_stats;
```

#### 仪表聚合 - 波动指标

用于分析上下波动的指标（如温度、内存使用量）。

```sql
-- 温度变化分析
SELECT
    sensor_id,
    delta(gauge_agg(timestamp, temperature)) as temp_delta,
    rate(gauge_agg(timestamp, temperature)) as temp_rate_per_sec
FROM weather_data
GROUP BY sensor_id;
```

### 时间加权分析

#### 时间加权平均

处理不规则采样数据，支持插值方法（LOCF 末次观测值保持、Linear 线性插值）。

```sql
-- 对不规则传感器读数计算加权平均
SELECT
    device_id,
    average(time_weight('LOCF', timestamp, sensor_value)) as weighted_avg,
    average(time_weight('Linear', timestamp, sensor_value)) as linear_avg
FROM iot_readings
GROUP BY device_id;

-- 合并多个时间范围
SELECT average(rollup(tw))
FROM (SELECT time_weight('LOCF', ts, val) as tw FROM readings_2023
      UNION ALL
      SELECT time_weight('LOCF', ts, val) FROM readings_2024) t;
```

### 数据可视化

#### LTTB 降采样

在保持视觉相似性的前提下对时序数据进行降采样，适用于图表展示。

```sql
-- 将 10 万个数据点降采样为 1000 个用于可视化
SELECT time, value
FROM unnest((
    SELECT lttb(timestamp, price, 1000)
    FROM stock_prices
    WHERE symbol = 'AAPL'
));
```

#### ASAP 平滑

通过降噪生成易于阅读的图表，同时保留趋势特征。

```sql
-- 将每日数据平滑为周粒度
SELECT time, value
FROM unnest((
    SELECT asap_smooth(date, daily_sales, 52)
    FROM sales_data
    WHERE date >= '2023-01-01'
));
```

### 统计分析

#### 统计聚合

提供全面的统计分析能力，支持一维和二维回归分析。

```sql
-- 多变量分析
SELECT
    -- 基础统计量
    average(stats_agg(response_time)) as avg_response,
    stddev(stats_agg(response_time)) as response_stddev,

    -- 回归分析
    slope(stats_agg(response_time, request_size)) as size_impact,
    corr(stats_agg(response_time, request_size)) as correlation,
    determination_coeff(stats_agg(response_time, request_size)) as r_squared
FROM performance_data;
```

### Timevector 数据类型

用于时序操作的高效中间表示。

```sql
-- 创建和操作 timevector
CREATE VIEW cpu_series AS
SELECT host_id, timevector(timestamp, cpu_percent) as ts
FROM system_metrics GROUP BY host_id;

-- 对 timevector 进行链式操作
SELECT host_id, unnest(lttb(ts, 100))
FROM cpu_series;
```

## 集成模式

### 连续聚合支持

大多数 Toolkit 函数可与 TimescaleDB 连续聚合无缝配合使用：

```sql
CREATE MATERIALIZED VIEW hourly_analytics AS
SELECT
    time_bucket('1 hour', timestamp) as hour,
    service_name,
    tdigest(100, response_time) as response_digest,
    counter_agg(timestamp, request_count) as request_counter,
    hyperloglog(64, user_id) as unique_users
FROM api_events
GROUP BY hour, service_name;

-- 查询预计算的聚合结果
SELECT
    hour,
    approx_percentile(0.95, response_digest) as p95_response,
    rate(request_counter) as req_per_sec,
    distinct_count(unique_users) as unique_users
FROM hourly_analytics
WHERE hour >= NOW() - INTERVAL '24 hours';
```

### 两步分析模式

存储中间聚合结果，支持多维度分析复用：

```sql
-- 第一步：创建聚合
CREATE TABLE daily_summaries AS
SELECT
    date_trunc('day', timestamp) as day,
    tdigest(200, response_time) as response_digest,
    stats_agg(response_time, request_size) as stats
FROM requests GROUP BY day;

-- 第二步：基于同一份数据进行多维分析
SELECT
    day,
    approx_percentile(0.50, response_digest) as median,
    approx_percentile(0.99, response_digest) as p99,
    average(stats) as avg_response,
    slope(stats) as size_correlation
FROM daily_summaries;
```

所有位于 **experimental** 模式（`toolkit_experimental`）中的函数可能会在版本间发生变化。生产环境中如需 API 稳定性保证，请使用稳定版函数。
