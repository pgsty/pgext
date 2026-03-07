
> [!WARNING] 此扩展已于 2025 年 3 月归档，不再维护。ParadeDB 的分析功能已整合到 pg_search 中。建议使用 pg_duckdb 或 pg_mooncake 作为数据湖分析的替代方案。

## 用法

https://github.com/paradedb/pg_analytics

示例：从 S3 读取 Parquet 文件：

```bash
CREATE EXTENSION pg_analytics;
CREATE FOREIGN DATA WRAPPER parquet_wrapper HANDLER parquet_fdw_handler VALIDATOR parquet_fdw_validator;

-- 提供 S3 凭证
CREATE SERVER parquet_server FOREIGN DATA WRAPPER parquet_wrapper;

-- 创建外部表并自动生成表结构
CREATE FOREIGN TABLE trips ()
SERVER parquet_server
OPTIONS (files 's3://paradedb-benchmarks/yellow_tripdata_2024-01.parquet');

-- 完成！现在可以像查询普通 PostgreSQL 表一样查询远程 Parquet 文件
SELECT COUNT(*) FROM trips;
  count
---------
 2964624
(1 row)
```

该外部数据包装器目前仅支持只读操作。



----

## Iceberg 支持

```sql
CREATE EXTENSION pg_analytics;

CREATE FOREIGN DATA WRAPPER iceberg_wrapper
    HANDLER iceberg_fdw_handler
    VALIDATOR iceberg_fdw_validator;

CREATE SERVER iceberg_server
    FOREIGN DATA WRAPPER iceberg_wrapper;

-- 请将示例中的占位结构替换为实际的表结构定义
CREATE FOREIGN TABLE iceberg_table (x INT)
    SERVER iceberg_server
    OPTIONS (files 's3://bucket/iceberg_folder');

-- 完成！现在可以查询 Iceberg 表了
SELECT COUNT(*) FROM iceberg_table;
```
