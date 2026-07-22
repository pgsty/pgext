## 用法

来源：

- [官方基准测试 README](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_benchmark/README.md)
- [官方 control 文件](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_benchmark/pg_lake_benchmark.control)

`pg_lake_benchmark` 提供 ClickBench、TPC-H、TPC-DS 的数据与查询驱动，用于比较 `pg_lake`、Iceberg 和普通 PostgreSQL 堆表。它依赖 `pg_lake`，应只在受控基准数据库中使用，因为生成器会创建大型模式与数据集，查询运行器还会丢弃结果。

### 核心流程

安装依赖和扩展，再选择一种基准与表类型：

```sql
CREATE EXTENSION pg_lake;
CREATE EXTENSION pg_lake_benchmark;

SELECT lake_tpch.gen(
  location => 's3://my-benchmark/tpch',
  table_type => 'pg_lake',
  scale_factor => 0.01,
  iteration_count => 10
);

SELECT * FROM lake_tpch.queries();
\timing on
SELECT lake_tpch.run_query(1);
```

应使用专用对象存储前缀和一次性数据库。生成大规模 TPC-H 数据时，提高 `iteration_count` 可降低峰值内存。

### 基准 API

- `lake_clickbench.create(table_type => ...)` 与 `lake_clickbench.load()`：创建并加载 ClickBench 表；加载过程会从 S3 读取约 15 GB 的 Parquet 数据集。
- `lake_clickbench.show_query(id)` 与 `lake_clickbench.run_query(id)`：查看或运行 43 条查询之一。
- `lake_tpch.gen(...)`、`lake_tpch.queries()`、`lake_tpch.run_query(id)`：生成并运行包含 22 条查询的 TPC-H。
- `lake_tpcds.gen(...)`、`lake_tpcds.queries()`、`lake_tpcds.run_query(id)`：生成并运行包含 99 条查询的 TPC-DS。
- 支持的 `table_type` 值为 `iceberg`、`pg_lake`、`heap`。

### 注意事项

运行函数会执行查询并丢弃结果，因此需要启用 psql 计时或明确采集服务器指标。TPC-H 与 TPC-DS 使用固定查询参数，与底层 DuckDB 生成器所记录的限制相同；它们不能替代应用自身的参数分布。

数据生成可能耗尽内存，并消耗大量对象存储容量、带宽与时间。比较结果时，应保持规模、凭据、区域、缓存、并发和服务器设置一致；不能把单次运行描述为可移植的性能保证。
