
--------

## Usage

pg_duckdb embeds DuckDB's columnar-vectorized analytics engine directly into PostgreSQL, enabling high-performance analytics within PostgreSQL while maintaining full SQL compatibility and ACID properties.

### Query Acceleration

Force analytical queries to use DuckDB's vectorized execution engine for dramatic performance improvements.

```sql
-- Enable DuckDB execution for current session
SET duckdb.force_execution = true;

-- Analytical queries automatically use DuckDB engine
SELECT 
    product_category,
    AVG(price) as avg_price,
    COUNT(*) as total_items
FROM products 
GROUP BY product_category
ORDER BY avg_price DESC;

-- Check execution plan
EXPLAIN SELECT count(*) FROM large_table;
-- Shows: Custom Scan (DuckDBScan) with DuckDB execution plan
```

### Data Lake Integration

Query cloud storage and modern data formats directly from PostgreSQL without ETL.

```sql
-- Read Parquet files from S3
SELECT customer_id, SUM(amount) as total_spent
FROM read_parquet('s3://data-lake/transactions/*.parquet')
WHERE purchase_date >= '2024-01-01'
GROUP BY customer_id;

-- Query CSV files with schema inference
SELECT * FROM read_csv('https://example.com/sales_data.csv')
WHERE region = 'North America';

-- Read JSON files with automatic flattening
SELECT event_type, COUNT(*) as event_count
FROM read_json('s3://events/user_events.json')
GROUP BY event_type;

-- Scan Apache Iceberg tables
SELECT * FROM iceberg_scan('s3://warehouse/orders/metadata.json')
WHERE order_date >= CURRENT_DATE - INTERVAL '30 days';

-- Query Delta Lake tables
SELECT product_id, SUM(quantity) as total_sales
FROM delta_scan('s3://delta-lake/sales/')
GROUP BY product_id;
```

### Hybrid Storage Tables

Create DuckDB-backed tables for analytical workloads with columnar storage and compression.

```sql
-- Create analytics table using DuckDB storage
CREATE TABLE sales_analytics (
    transaction_id BIGINT,
    product_id INT,
    customer_id INT,
    sale_date DATE,
    amount DECIMAL(10,2),
    quantity INT
) USING duckdb;

-- Insert data (stored in columnar format)
INSERT INTO sales_analytics 
SELECT id, product_id, customer_id, sale_date, price * quantity, quantity
FROM transactions
WHERE sale_date >= '2024-01-01';

-- High-performance analytical queries
SELECT 
    product_id,
    DATE_TRUNC('month', sale_date) as month,
    SUM(amount) as monthly_revenue,
    AVG(quantity) as avg_quantity
FROM sales_analytics
GROUP BY product_id, month
ORDER BY monthly_revenue DESC;
```

### Cloud Storage Security

Manage secure credentials for cloud data access.

```sql
-- Create S3 credentials
SELECT duckdb.create_simple_secret(
    type := 'S3',
    key_id := 'AKIAIOSFODNN7EXAMPLE',
    secret := 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY'
);

-- Create Azure credentials
SELECT duckdb.create_simple_secret(
    type := 'AZURE',
    connection_string := 'DefaultEndpointsProtocol=https;AccountName=myaccount;AccountKey=mykey'
);

-- Use credential chain for AWS
CREATE SERVER s3_server FOREIGN DATA WRAPPER duckdb_fdw 
OPTIONS (type 'S3', provider 'credential_chain');
```

### MotherDuck Cloud Integration

Connect to MotherDuck cloud analytics platform for distributed analytics.

```sql
-- Enable MotherDuck connection
CALL duckdb.enable_motherduck('your_token_here', 'production_db');

-- Create cloud-synchronized table
CREATE TABLE cloud_metrics(
    timestamp TIMESTAMP,
    metric_name TEXT,
    value DOUBLE,
    tags JSON
) USING duckdb;

-- Query across local and cloud data
SELECT 
    local.user_id,
    local.session_count,
    cloud.conversion_rate
FROM user_sessions local
JOIN ddb$shared_analytics.user_metrics cloud 
  ON local.user_id = cloud.user_id;

-- Force cloud synchronization
SELECT duckdb.force_motherduck_sync();
```

### Extension Management

Leverage DuckDB's extension ecosystem for additional capabilities.

```sql
-- Install extensions (superuser required)
SELECT duckdb.install_extension('iceberg');
SELECT duckdb.install_extension('delta'); 
SELECT duckdb.install_extension('azure');

-- View installed extensions
SELECT name, loaded, installed 
FROM duckdb.extensions();

-- Extensions auto-load when functions are used
SELECT * FROM delta_scan('azure://container/delta-table/');
```

### Direct DuckDB Queries

Execute raw DuckDB queries for advanced functionality.

```sql
-- Execute DuckDB query directly
SELECT duckdb.query('PRAGMA version');

-- Raw query with complex DuckDB syntax
SELECT duckdb.raw_query($$
    WITH monthly_sales AS (
        SELECT DATE_TRUNC('month', sale_date) as month, 
               SUM(amount) as total
        FROM sales_data 
        GROUP BY month
    )
    SELECT month, total,
           LAG(total) OVER (ORDER BY month) as prev_month,
           total - LAG(total) OVER (ORDER BY month) as growth
    FROM monthly_sales
$$);

-- Reset DuckDB instance
SELECT duckdb.recycle_ddb();
```

## Configuration

### Performance Tuning

```sql
-- Memory allocation (default: 80% of system memory)
SET duckdb.max_memory = '8GB';

-- Thread configuration
SET duckdb.threads = 8;

-- PostgreSQL scan optimization
SET duckdb.max_workers_per_postgres_scan = 4;

-- Disable specific execution for debugging
SET duckdb.force_execution = false;
```

### Security Controls

```sql
-- Restrict access to specific PostgreSQL roles
SET duckdb.postgres_role = 'analytics_team';

-- Disable filesystem access
SET duckdb.disabled_filesystems = 'LocalFileSystem';

-- Control extension installation
SET duckdb.autoinstall_known_extensions = false;
SET duckdb.allow_community_extensions = false;

-- External access control
SET duckdb.enable_external_access = true;
```

### Data Type Compatibility

| PostgreSQL Type        | DuckDB Support | Notes                           |
|------------------------|----------------|---------------------------------|
| INTEGER, BIGINT        | Full           | Direct mapping                  |
| REAL, DOUBLE PRECISION | Full           | Direct mapping                  |
| TEXT, VARCHAR          | Full           | Direct mapping                  |
| BOOLEAN                | Full           | Direct mapping                  |
| UUID                   | Full           | Direct mapping                  |
| JSON/JSONB             | Full           | JSONB converted to JSON         |
| DATE, TIMESTAMP        | Full           | Microsecond precision           |
| NUMERIC                | Limited        | May lose precision              |
| ARRAYS (1D)            | Full           | Multi-dimensional not supported |
| ENUM                   | None           | Not supported                   |

## Use Cases

- **Real-time Analytics**: Accelerate analytical queries on transactional data by 10-100x
- **Data Lake Analytics**: Query Parquet, Iceberg, and Delta Lake tables directly
- **Hybrid Workloads**: Combine OLTP applications with high-performance analytics
- **Operational Dashboards**: Build responsive dashboards on live PostgreSQL data
- **ETL Elimination**: Remove data movement overhead for analytical workloads
- **Cloud Data Integration**: Seamlessly access cloud storage from PostgreSQL

pg_duckdb transforms PostgreSQL into a powerful hybrid OLTP/OLAP database, enabling organizations to run analytical workloads at scale without sacrificing transactional consistency or requiring separate analytics infrastructure.