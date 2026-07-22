## 用法

来源：

- [0.3.7 版官方 README](https://github.com/paradedb/pg_analytics/blob/v0.3.7/README.md)
- [扩展控制文件](https://github.com/paradedb/pg_analytics/blob/v0.3.7/pg_analytics.control)
- [0.3.7 版 Rust 清单](https://github.com/paradedb/pg_analytics/blob/v0.3.7/Cargo.toml)
- [SQL 可见的辅助 API](https://github.com/paradedb/pg_analytics/tree/v0.3.7/src/api)

`pg_analytics` 0.3.7 在 PostgreSQL 中嵌入 DuckDB，通过只读 FDW 暴露对象存储文件和数据湖格式。ParadeDB 已于 2025 年 3 月停止并归档该项目；它不再维护，应视为遗留集成。

### 核心流程

为了让每个后端都获得基于钩子的查询下推，上游建议先在 `postgresql.conf` 中加入动态库并重启 PostgreSQL，再创建扩展：

```text
shared_preload_libraries = 'pg_analytics'
```

创建特定格式的包装器、服务器和外部表。空列列表会要求扩展发现 Parquet 模式：

```sql
CREATE EXTENSION pg_analytics;

CREATE FOREIGN DATA WRAPPER parquet_wrapper
HANDLER parquet_fdw_handler
VALIDATOR parquet_fdw_validator;

CREATE SERVER parquet_server
FOREIGN DATA WRAPPER parquet_wrapper;

CREATE FOREIGN TABLE trips ()
SERVER parquet_server
OPTIONS (files 's3://paradedb-benchmarks/yellow_tripdata_2024-01.parquet');

SELECT count(*) FROM trips;
```

凭据和对象存储配置应放入用户映射，而不是表定义。对于 S3 兼容存储，相关映射选项包括凭据，以及端点需要的 `type`、`region`、`endpoint`、`use_ssl` 和 `url_style`。

### 主要对象

- 为 `csv`、`json`、`parquet`、`delta`、`iceberg` 和 `spatial` 数据源提供了成对的 FDW 处理器/校验器。
- 外部表选项 `files` 指定本地路径、HTTP URL、对象存储 URL 或特定格式的表位置。
- `sniff_csv` 返回检测到的 CSV 配置和列。
- `parquet_describe` 和 `parquet_schema` 检查 Parquet 关系。
- `duckdb_execute` 在嵌入式 DuckDB 连接上运行语句；`duckdb_settings` 和 `duckdb_extensions` 暴露其状态。

版本 0.3.7 嵌入 DuckDB 1.1.0，并声明支持 PostgreSQL 13+。上游列出的格式包括 Parquet、CSV、JSON、地理空间文件、Delta Lake 和 Apache Iceberg；数据源包括本地文件系统、HTTP、S3 兼容存储、Google Cloud Storage、Azure 存储和 Hugging Face。

### 运维边界

这些 FDW 只读。文件系统访问和出站请求由 PostgreSQL 服务器进程发起，DuckDB 查询执行还会与事务工作负载共享后端资源。应限制函数执行权、用户映射、文件路径和远程端点；不要向不受信任角色开放 `duckdb_execute`。

修改 `shared_preload_libraries` 需要重启。已归档项目明确不支持 Windows，也不会再提供兼容性修复。应固定使用确切的 0.3.7 构建，在目标 PostgreSQL 版本上测试规划器下推和类型转换，并规划迁移到维护中的方案，而不是将其用于新的生产系统。
