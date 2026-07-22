## Usage

Sources:

- [Official benchmark README](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_benchmark/README.md)
- [Official control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_benchmark/pg_lake_benchmark.control)

`pg_lake_benchmark` supplies ClickBench, TPC-H, and TPC-DS data/query drivers for comparing `pg_lake`, Iceberg, and ordinary PostgreSQL heap tables. It depends on `pg_lake` and is intended for controlled benchmark databases, because its generators create large schemas and data sets and its runners discard query results.

### Core Workflow

Install the dependency and extension, then choose one benchmark and table type:

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

Use a dedicated object-store prefix and disposable database. Increasing `iteration_count` reduces peak memory during large TPC-H generation.

### Benchmark APIs

- `lake_clickbench.create(table_type => ...)` and `lake_clickbench.load()`: create and load the ClickBench table. The load reads an approximately 15 GB Parquet data set from S3.
- `lake_clickbench.show_query(id)` and `lake_clickbench.run_query(id)`: inspect or run one of 43 queries.
- `lake_tpch.gen(...)`, `lake_tpch.queries()`, and `lake_tpch.run_query(id)`: generate and run the 22-query TPC-H suite.
- `lake_tpcds.gen(...)`, `lake_tpcds.queries()`, and `lake_tpcds.run_query(id)`: generate and run the 99-query TPC-DS suite.
- Supported `table_type` values are `iceberg`, `pg_lake`, and `heap`.

### Caveats

The run functions execute a query and discard its result, so enable psql timing or capture server-side metrics explicitly. TPC-H and TPC-DS use fixed query parameters, matching the limitation documented by the underlying DuckDB generators; they are not substitute workloads for an application's parameter distribution.

Data generation can exhaust memory and consume substantial object-storage capacity, bandwidth, and time. Keep scale, credentials, region, caches, concurrency, and server settings identical when comparing results, and do not present a single run as a portable performance guarantee.
