

--------

## Usage

pg_mooncake creates real-time columnstore mirrors of PostgreSQL tables using Apache Iceberg format, enabling fast analytics queries with sub-second data freshness while maintaining full PostgreSQL compatibility.

> **Warning**: This extension conflicts with `pg_duckdb` & `duckdb_fdw` due to shared `libduckdb.so` library usage.

### Real-time Columnstore Mirrors

Create analytical mirrors of existing PostgreSQL tables that automatically sync data changes.

```sql
-- Create extension
CREATE EXTENSION pg_mooncake;

-- Create source table
CREATE TABLE sales_data(
    id bigint PRIMARY KEY,
    product_id int,
    sale_date timestamp,
    amount decimal(10,2),
    region text
);

-- Create real-time columnstore mirror
CALL mooncake.create_table('sales_analytics', 'sales_data');

-- Insert data (automatically synced to mirror)
INSERT INTO sales_data VALUES 
    (1, 101, '2024-01-15 10:30:00', 1250.00, 'North'),
    (2, 102, '2024-01-15 11:45:00', 890.50, 'South');

-- Query columnstore mirror for analytics
SELECT 
    region,
    DATE_TRUNC('month', sale_date) as month,
    SUM(amount) as total_sales,
    AVG(amount) as avg_sale
FROM sales_analytics 
GROUP BY region, month;
```

### Direct Columnstore Tables

Create columnstore tables directly for analytical workloads.

```sql
-- Create columnstore table
CREATE TABLE user_activity(
    user_id BIGINT,
    activity_type TEXT,
    activity_timestamp TIMESTAMP,
    duration INT
) USING columnstore;

-- Insert data
INSERT INTO user_activity VALUES
    (1, 'login', '2024-01-01 08:00:00', 120),
    (2, 'page_view', '2024-01-01 08:05:00', 30),
    (3, 'logout', '2024-01-01 08:30:00', 60);

-- Run analytical queries
SELECT
    user_id,
    activity_type,
    SUM(duration) AS total_duration,
    COUNT(*) AS activity_count
FROM user_activity
GROUP BY user_id, activity_type
ORDER BY user_id, activity_type;
```

### Cloud Storage Integration

Configure pg_mooncake to use cloud storage backends like S3.

```sql
-- Create S3 credentials
SELECT mooncake.create_secret(
    'my_s3_secret', 
    'S3', 
    'ACCESS_KEY_ID', 
    'SECRET_ACCESS_KEY', 
    '{"REGION": "us-west-2"}'
);

-- Set default S3 bucket
SET mooncake.default_bucket = 's3://my-analytics-bucket';

-- Disable local cache for cloud environments
SET mooncake.enable_local_cache = false;

-- Create columnstore table using S3 storage
CREATE TABLE metrics_data(
    timestamp TIMESTAMP,
    metric_name TEXT,
    value DOUBLE PRECISION,
    tags JSONB
) USING columnstore;
```

### Multi-Format Data Loading

Load data from various file formats into columnstore tables.

```sql
-- Load from Parquet files
COPY metrics_data FROM 's3://data-lake/metrics/*.parquet';

-- Load from CSV files  
COPY user_events FROM '/path/to/events.csv' WITH CSV HEADER;

-- Load from JSON files
COPY log_data FROM 's3://logs/application.json' WITH (FORMAT json);
```

### External Data Lake Access

Query external Iceberg and Delta Lake tables directly.

```sql
-- Query external Iceberg table
SELECT COUNT(*) FROM iceberg_scan('s3://datalake/sales_iceberg/');

-- Query Delta Lake table
SELECT * FROM delta_scan('s3://datalake/user_events_delta/') 
WHERE event_date >= '2024-01-01';
```

### Hybrid Workloads

Combine transactional and analytical queries seamlessly.

```sql
-- Join columnstore table with regular PostgreSQL table
SELECT 
    t.symbol,
    t.price,
    c.company_name,
    c.sector
FROM trades_analytics t 
JOIN companies c ON t.symbol = c.symbol
WHERE t.trade_date >= CURRENT_DATE - INTERVAL '7 days'
  AND t.volume > 1000000;

-- Mixed workload: OLTP inserts with OLAP queries
BEGIN;
INSERT INTO orders (customer_id, product_id, quantity, price) 
VALUES (1001, 5502, 2, 29.99);

-- Immediate analytics on fresh data
SELECT 
    product_id,
    SUM(quantity * price) as revenue,
    COUNT(*) as order_count
FROM orders_analytics
WHERE order_date >= CURRENT_DATE
GROUP BY product_id;
COMMIT;
```


--------

## Performance Features

### Query Acceleration

Queries on columnstore tables use DuckDB execution engine for high performance:

```sql
EXPLAIN SELECT 
    region,
    AVG(amount) as avg_sale,
    PERCENTILE_CONT(0.95) WITHIN GROUP (ORDER BY amount) as p95
FROM sales_analytics 
GROUP BY region;

-- Shows: Custom Scan (DuckDBScan) with optimized execution plan
```

### Configuration Options

```sql
-- Set maximum memory usage (default: 1GB)
SET mooncake.max_memory = '4GB';

-- Configure thread count for parallel processing
SET mooncake.threads = 8;

-- Enable/disable local caching
SET mooncake.enable_local_cache = true;

-- Set data compression level
SET mooncake.compression = 'zstd';
```


--------

## Use Cases

- **Real-time Analytics**: Live dashboards and reports with sub-second data freshness
- **Hybrid Workloads**: Applications requiring both OLTP and OLAP capabilities  
- **Data Lake Integration**: Bridge PostgreSQL with modern data lake architectures
- **Performance Optimization**: Accelerate analytical queries without application changes
- **Operational Analytics**: Real-time monitoring and alerting on transactional data

pg_mooncake provides PostgreSQL-native columnstore capabilities while maintaining compatibility with existing applications and enabling integration with modern data lake ecosystems through Apache Iceberg format.