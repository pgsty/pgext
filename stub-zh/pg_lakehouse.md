## 用法

来源：

- [官方 0.7.6 版 README](https://github.com/paradedb/paradedb/blob/v0.7.6/pg_lakehouse/README.md)
- [官方 0.7.6 版 control 文件](https://github.com/paradedb/paradedb/blob/v0.7.6/pg_lakehouse/pg_lakehouse.control)
- [官方 0.7.6 版源码目录](https://github.com/paradedb/paradedb/tree/v0.7.6/pg_lakehouse)

`pg_lakehouse` 版本 `0.7.6` 是 ParadeDB 历史上的分析型 FDW，通过 Apache DataFusion 查询对象存储文件和 Delta Lake 表。它支持 S3 兼容、Azure、GCS 和本地存储，以及 Parquet、CSV、JSON、Avro 文件。该 API 属于固定标签下的子项目，不能与当前 ParadeDB 搜索文档混用。

### 核心流程

扩展使用执行器钩子，因此需要预加载并重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_lakehouse'
```

创建对象存储封装器、服务器和字段与远程文件完全一致的外部表：

```sql
CREATE EXTENSION pg_lakehouse;
CREATE FOREIGN DATA WRAPPER s3_wrapper
  HANDLER s3_fdw_handler VALIDATOR s3_fdw_validator;

CREATE SERVER s3_server FOREIGN DATA WRAPPER s3_wrapper
OPTIONS (region 'us-east-1', allow_anonymous 'true');

CREATE FOREIGN TABLE trips (
  "VendorID" integer,
  "tpep_pickup_datetime" timestamp,
  "trip_distance" double precision,
  "total_amount" double precision
)
SERVER s3_server
OPTIONS (
  path 's3://paradedb-benchmarks/yellow_tripdata_2024-01.parquet',
  extension 'parquet'
);

SELECT count(*) FROM trips;
```

如果 `path` 指向分区目录，末尾必须有 `/`。DataFusion 区分大小写，因此混合大小写字段名需要双引号。

### 辅助函数与格式

- `arrow_schema(server => ..., path => ..., extension => ...)`：选择 PostgreSQL 字段类型前查看 Arrow 字段。
- `CALL connect_table('schema.table')`：提前建立对象存储连接，并验证表凭据。
- 对象存储：Amazon S3/S3 兼容服务、Azure Blob/ADLS Gen2、Google Cloud Storage 与本地文件。
- 0.7.6 的格式：Parquet、CSV、JSON、Avro 与 Delta Lake；README 把 ORC 和 Iceberg 列为尚未实现。

### 注意事项

版本 `0.7.6` 只记录支持 PostgreSQL 14–16。预加载会影响整个服务器，查询下推还会改变规划与执行路径；必须用固定二进制验证每类负载。模式/类型不匹配会在查询时失败，`date`、`timestamp`、`timestamptz` 尤其需要谨慎。

对象存储凭据、网络出口、文件一致性、分区布局和冷连接成本都会影响行为。应严格限制外部服务器/表权限，避免把长期密钥放在广泛可见的服务器选项中，并通过 `EXPLAIN` 确认实际下推到 DataFusion 的操作。
