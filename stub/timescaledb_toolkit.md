## Usage

Sources:

- [Documentation index](https://github.com/timescale/timescaledb-toolkit/blob/1.27.0%2Bp2/docs/README.md)
- [README](https://github.com/timescale/timescaledb-toolkit/blob/1.27.0%2Bp2/README.md)
- [Control](https://github.com/timescale/timescaledb-toolkit/blob/1.27.0%2Bp2/extension/timescaledb_toolkit.control)
- [Cargo version](https://github.com/timescale/timescaledb-toolkit/blob/1.27.0%2Bp2/extension/Cargo.toml)
- [Changelog](https://github.com/timescale/timescaledb-toolkit/blob/1.27.0%2Bp2/CHANGELOG.md)
- [Counter aggregates](https://github.com/timescale/timescaledb-toolkit/blob/1.27.0%2Bp2/docs/counter_agg.md)
- [Gauge aggregates](https://github.com/timescale/timescaledb-toolkit/blob/1.27.0%2Bp2/docs/gauge_agg.md)
- [UddSketch](https://github.com/timescale/timescaledb-toolkit/blob/1.27.0%2Bp2/docs/uddsketch.md)
- [Statistics](https://github.com/timescale/timescaledb-toolkit/blob/1.27.0%2Bp2/docs/stats_agg.md)
- [Timevector pipelines](https://github.com/timescale/timescaledb-toolkit/blob/1.27.0%2Bp2/docs/timeseries_pipeline_elements.md)
- [1.27.0+p2 release](https://github.com/timescale/timescaledb-toolkit/releases/tag/1.27.0%2Bp2)

TimescaleDB Toolkit provides specialized functions for time-series analytics using a **two-step aggregation pattern**. Most functions create intermediate representations that accessor functions then query, enabling efficient reuse and multiple analyses.

### Enablement and Compatibility

`timescaledb_toolkit` 1.27.0 is published under tag 1.27.0+p2; upstream CI covers PostgreSQL 16–18. After installing the matching library, enable it in each database:

```sql
CREATE EXTENSION timescaledb_toolkit;
```

The control file marks it trusted, permitting installation by a role with CREATE privilege on the database; shared preload is not required. Examples using `time_bucket` below also assume TimescaleDB is enabled. Upgrade existing installations with matching libraries and supplied SQL upgrade paths; experimental objects have no upgrade compatibility guarantee.

### Approximate Analytics

#### HyperLogLog - Distinct Count Estimation

Probabilistic distinct counting with configurable precision for high-cardinality datasets.

```sql
-- Estimate unique users per day
SELECT
    date_trunc('day', timestamp) as day,
    distinct_count(hyperloglog(64, user_id)) as unique_users
FROM events
GROUP BY day;

-- Combine counts across partitions
SELECT distinct_count(rollup(hll))
FROM (SELECT hyperloglog(32, session_id) as hll FROM events_2023
      UNION ALL
      SELECT hyperloglog(32, session_id) FROM events_2024) t;
```

#### T-Digest - Quantile Approximation

High-accuracy percentile estimation optimized for tail quantiles (P95, P99).

```sql
-- Track response time percentiles
SELECT
    service_name,
    approx_percentile(0.50, tdigest(100, response_time)) as p50,
    approx_percentile(0.95, tdigest(100, response_time)) as p95,
    approx_percentile(0.99, tdigest(100, response_time)) as p99
FROM api_metrics
GROUP BY service_name;

-- Store hourly percentiles in a materialized view
CREATE MATERIALIZED VIEW hourly_percentiles AS
SELECT
    time_bucket('1 hour', timestamp) as hour,
    tdigest(200, response_time) as digest
FROM requests GROUP BY hour;
```

#### UddSketch - Bounded Error Quantiles

Quantile estimation with a requested starting error bound. Too few buckets can increase the actual bound; check `error` on the resulting sketch.

```sql
-- CPU percentiles with a starting 1% error target
SELECT
    host_id,
    approx_percentile(0.95, uddsketch(100, 0.01, cpu_percent)) as p95_cpu,
    error(uddsketch(100, 0.01, cpu_percent)) as actual_error
FROM system_metrics
GROUP BY host_id;
```

### Counter Analytics

#### Counter Aggregates - Monotonic Metrics

Handle counters that increase monotonically with automatic reset detection.

```sql
-- Request rate calculation
SELECT
    time_bucket('5 min', timestamp) as bucket,
    rate(counter_agg(timestamp, request_count)) as requests_per_sec,
    delta(counter_agg(timestamp, request_count)) as total_requests
FROM metrics
GROUP BY bucket;

-- Extrapolated rate for each hourly bucket
SELECT bucket, extrapolated_rate(summary, method => 'prometheus') AS bytes_per_second
FROM (
    SELECT time_bucket('1 hour', timestamp) AS bucket,
           counter_agg(timestamp, bytes_sent,
                       bounds => time_bucket_range('1 hour', timestamp)) AS summary
    FROM network_stats
    GROUP BY time_bucket('1 hour', timestamp)
) hourly;
```

#### Gauge Aggregates - Varying Metrics

Analytics for metrics that vary up and down (temperature, memory usage).

```sql
-- Temperature change analysis
SELECT
    sensor_id,
    delta(gauge_agg(timestamp, temperature)) as temp_delta,
    rate(gauge_agg(timestamp, temperature)) as temp_rate_per_sec
FROM weather_data
GROUP BY sensor_id;
```

### Time-Weighted Analytics

#### Time-Weighted Averages

Handle irregularly sampled data with interpolation methods (LOCF, Linear).

```sql
-- Weighted average for irregular sensor readings
SELECT
    device_id,
    average(time_weight('LOCF', timestamp, sensor_value)) as weighted_avg,
    average(time_weight('Linear', timestamp, sensor_value)) as linear_avg
FROM iot_readings
GROUP BY device_id;

-- Combining multiple time ranges
SELECT average(rollup(tw))
FROM (SELECT time_weight('LOCF', ts, val) as tw FROM readings_2023
      UNION ALL
      SELECT time_weight('LOCF', ts, val) FROM readings_2024) t;
```

### Data Visualization

#### LTTB Downsampling

Downsample time series while preserving visual similarity for charts.

```sql
-- Reduce 100K points to 1K for visualization
SELECT time, value
FROM unnest((
    SELECT lttb(timestamp, price, 1000)
    FROM stock_prices
    WHERE symbol = 'AAPL'
));
```

#### ASAP Smoothing

Generate human-readable graphs by reducing noise while preserving trends.

```sql
-- Smooth daily data to weekly resolution
SELECT time, value
FROM unnest((
    SELECT asap_smooth(date, daily_sales, 52)
    FROM sales_data
    WHERE date >= '2023-01-01'
));
```

### Statistical Analysis

#### Stats Aggregates

Comprehensive statistical analysis with 1D and 2D regression capabilities.

```sql
-- Multi-variable analysis
SELECT
    -- Basic statistics
    average(stats_agg(response_time)) as avg_response,
    stddev(stats_agg(response_time)) as response_stddev,

    -- Regression analysis
    slope(stats_agg(response_time, request_size)) as size_impact,
    corr(stats_agg(response_time, request_size)) as correlation,
    determination_coeff(stats_agg(response_time, request_size)) as r_squared
FROM performance_data;
```

### Timevector Data Type

Timevectors hold intermediate time-series values. Pipeline operators such as `toolkit_experimental.lttb` remain experimental; sort before downsampling.

```sql
-- Create and manipulate timevector
CREATE VIEW cpu_series AS
SELECT host_id, timevector(timestamp, cpu_percent) as ts
FROM system_metrics GROUP BY host_id;

-- Chain operations on timevector
SELECT host_id, unnest(ts -> toolkit_experimental.sort()
                         -> toolkit_experimental.lttb(100))
FROM cpu_series;
```

### Integration Patterns

### Materialized Aggregates

The following is an ordinary materialized view; refresh it explicitly when source data changes. A TimescaleDB continuous aggregate additionally requires a hypertable, its continuous option and a suitable refresh policy.

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

-- Query pre-computed aggregates
SELECT
    hour,
    approx_percentile(0.95, response_digest) as p95_response,
    rate(request_counter) as req_per_sec,
    distinct_count(unique_users) as unique_users
FROM hourly_analytics
WHERE hour >= NOW() - INTERVAL '24 hours';
```

### Two-Step Analysis Pattern

Store intermediate aggregates for multiple analyses:

```sql
-- Step 1: Create aggregates
CREATE TABLE daily_summaries AS
SELECT
    date_trunc('day', timestamp) as day,
    tdigest(200, response_time) as response_digest,
    stats_agg(response_time, request_size) as stats
FROM requests GROUP BY day;

-- Step 2: Multiple analyses from same data
SELECT
    day,
    approx_percentile(0.50, response_digest) as median,
    approx_percentile(0.99, response_digest) as p99,
    average_y(stats) as avg_response,
    slope(stats) as size_correlation
FROM daily_summaries;
```

All functions in the **experimental** schema (`toolkit_experimental`) may change between versions. Use stable functions for production workloads requiring API stability.

### Version Notes

Since 1.23, `max_n`, `max_n_by`, `min_n` and `min_n_by` return empty results on empty input, while `stats_agg(bigint)` rejects values outside the exact double-precision integer range. Retain application tests for these behaviors when upgrading.

1.24 adds equality for `heartbeat_agg` and repairs empty frequency aggregates and statistical overflow handling. 1.25 stabilizes `gauge_agg` and fixes empty t-digest input and several type-handling defects. 1.26 removes combine/serialize/deserialize support from `time_weight` and its rollup, retaining their nonparallel aggregation boundary.

1.27 adds equality for `statssummary1d` and `statssummary2d`, fixes sparse HyperLogLog rollups and prevents an ignored counter point at the same timestamp from counting as a reset. It rejects invalid UTF-8 state names, lambdas longer than 1024 bytes or nested beyond 256 parenthesis levels, frequency rollups of different element types and a mismatched NULL type argument to `topn`. It also restores the pipeline map ACL check.
