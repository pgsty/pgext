
## Usage

TimescaleDB Toolkit provides specialized functions for time-series analytics using a **two-step aggregation pattern**. Most functions create intermediate representations that accessor functions then query, enabling efficient reuse and multiple analyses.

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

-- Hourly percentiles with continuous aggregation
CREATE MATERIALIZED VIEW hourly_percentiles AS
SELECT 
    time_bucket('1 hour', timestamp) as hour,
    tdigest(200, response_time) as digest
FROM requests GROUP BY hour;
```

#### UddSketch - Bounded Error Quantiles

Quantile estimation with guaranteed maximum relative error bounds.

```sql
-- CPU utilization percentiles with 1% max error
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

-- Extrapolated rate for partial buckets
SELECT 
    extrapolated_rate(
        counter_agg(timestamp, bytes_sent, 
                   bounds => time_bucket_range('1 hour', timestamp))
    ) as bytes_per_second
FROM network_stats;
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

Efficient intermediate representation for time series operations.

```sql
-- Create and manipulate timevector
CREATE VIEW cpu_series AS 
SELECT host_id, timevector(timestamp, cpu_percent) as ts
FROM system_metrics GROUP BY host_id;

-- Chain operations on timevector
SELECT host_id, unnest(lttb(ts, 100)) 
FROM cpu_series;
```

## Integration Patterns

### Continuous Aggregation Support

Most toolkit functions work seamlessly with TimescaleDB continuous aggregates:

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
    average(stats) as avg_response,
    slope(stats) as size_correlation
FROM daily_summaries;
```

All functions in the **experimental** schema (`toolkit_experimental`) may change between versions. Use stable functions for production workloads requiring API stability.
