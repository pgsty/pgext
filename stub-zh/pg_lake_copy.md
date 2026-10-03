## 用法

来源：

- [版本 3.5 控制文件](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/pg_lake_copy/pg_lake_copy.control)
- [官方数据湖导入和导出指南](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/docs/data-lake-import-export.md)
- [官方文件格式参考](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/docs/file-formats-reference.md)
- [3.4 至 3.5 SQL 迁移](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/pg_lake_copy/pg_lake_copy--3.4--3.5.sql)

`pg_lake_copy` 扩展了 PostgreSQL `COPY`，使得查询、堆表、外部湖表和 Iceberg 表可以与本地路径、HTTP 端点以及配置的对象存储交换 Parquet、CSV 或换行符分隔的 JSON 文件。它通过钩子添加行为，并没有独立的 SQL 函数 API。

pg_lake 发布与软件包版本为 `3.5.3`，SQL 扩展版本为 `3.5`。动态库与查询服务器应使用同一发布版本。

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
