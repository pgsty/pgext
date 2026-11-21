## 用法

来源：

- [3.7.0 pg_lake_copy.control](https://github.com/Snowflake-Labs/pg_lake/blob/v3.7.0/pg_lake_copy/pg_lake_copy.control)
- [3.7.0 v3.7.0](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.7.0)

- [Upstream source](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.6.2)

- [版本 3.6 控制文件](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_copy/pg_lake_copy.control)
- [官方数据湖导入和导出指南](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/docs/data-lake-import-export.md)
- [官方文件格式参考](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/docs/file-formats-reference.md)
- [3.4 至 3.5 SQL 迁移](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_copy/pg_lake_copy--3.4--3.5.sql)

- [Upstream 3.6.0](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.6.0)
- [3.5 to 3.6 SQL](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_copy/pg_lake_copy--3.5--3.6.sql)

`pg_lake_copy` 扩展了 PostgreSQL `COPY`，使得查询、堆表、外部湖表和 Iceberg 表可以与本地路径、HTTP 端点以及配置的对象存储交换 Parquet、CSV 或换行符分隔的 JSON 文件。它通过钩子添加行为，并没有独立的 SQL 函数 API。

本次核验的上游发布为 `3.6.0`，SQL 扩展版本为 `3.6`，支持 PostgreSQL 16–19。现有软件包版本单独跟踪，动态库与查询服务器应使用同一发布版本。

### 启用组件

正常的入口点会一起安装 `pg_lake_copy` 及其精确依赖项：

```sql
CREATE EXTENSION pg_lake CASCADE;
```

它的控制文件需要 `pg_lake_engine`、`pg_lake_iceberg` 和 `pg_lake_table`。部署时还须将 `pg_extension_base` 加入 `shared_preload_libraries`，并启动 `pgduck_server`。

### 导出和导入

格式可以从路径后缀推断或显式选择：

```sql
COPY (
    SELECT event_id, event_time, payload
    FROM events
    WHERE event_time >= DATE '2026-07-01'
)
TO 's3://analytics-bucket/events/july.parquet'
WITH (format 'parquet');

COPY events_archive
FROM 's3://analytics-bucket/events/july.parquet'
WITH (format 'parquet');
```

CSV 与压缩输出使用湖写入器扩展的 `COPY` 选项：

```sql
COPY (SELECT * FROM daily_summary)
TO 's3://analytics-bucket/summary/daily.csv.gz'
WITH (format 'csv', header true, compression 'gzip');
```

目标可以是 PostgreSQL 堆表或 Iceberg 表；源也可以是安装的 pg_lake 堆栈支持的任何查询。

### 格式和运行时边界

- Parquet 是列式的，并保留了受支持的数据类型值；CSV 和换行符分隔的 JSON 有特定格式的推断和转换选项，这些选项在上游文档中有所说明。
- 对象存储访问通过 `pgduck_server` 进行。其凭据链、网络访问以及桶权限必须允许请求的读取或写入。
- `COPY` 是一个语句，并参与外围的 PostgreSQL 事务，但远程文件和清理也依赖于 pg_lake 事务/队列机制。在重试大规模导出之前，请检查失败的操作和孤儿清理。
- 版本 `3.5` 没有为 `pg_lake_copy` 新增独立 SQL API；从 `3.4` 到 `3.5` 的迁移脚本为空。COPY 行为仍取决于匹配的湖扩展动态库和查询服务器二进制。

### 上游 3.6.0

上游 3.6.0 增加 PostgreSQL 19 支持，并使用 DuckDB 1.5.6。应部署匹配动态库与 `pgduck_server`，保持预加载 `pg_extension_base`，重启加载旧库的进程，并执行所提供的 SQL 升级到 3.6。此组件从 3.5 到 3.6 的 SQL 迁移为空，但运行库及依赖仍须按同一发布协调升级。

### 上游 3.6.2

3.6.2 补丁修正 pgduck_server 逐字节传输 CSV 值的处理。应部署与扩展库匹配的查询服务二进制，并重新连接或重启仍持有旧代码的进程。组件 SQL 版本仍为 3.6，extversion 不变不能证明查询服务修复已经安装。软件包可用性仍是单独的目录事实。

### 上游 3.7.0

当前核对的发行版为 3.7.0，本组件的控制版本为 3.7。前面的旧版本章节描述历史行为和软件包示例。应使用匹配的库与 `pgduck_server` 协同升级整个扩展族，重启持有旧代码的进程，并执行提供的扩展升级脚本，将 SQL 版本更新至 3.7。这不代表相应软件包已经可用。
