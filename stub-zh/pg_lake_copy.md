## 用法

来源：

- [Version 3.4控制文件](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_copy/pg_lake_copy.control)
- [官方数据湖导入和导出指南](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/data-lake-import-export.md)
- [官方文件格式参考](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/file-formats-reference.md)
- [Version 3.4 SQL文件](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_copy/pg_lake_copy--3.3--3.4.sql)

`pg_lake_copy` 扩展了 PostgreSQL `COPY`，使得查询、堆表、外部湖表和 Iceberg 表可以与本地路径、HTTP 端点以及配置的对象存储交换 Parquet、CSV 或换行符分隔的 JSON 文件。它通过钩子添加行为，并没有独立的 SQL 函数 API。

### 启用组件

正常的入口点会一起安装 `pg_lake_copy` 及其精确依赖项：

```sql
CREATE EXTENSION pg_lake CASCADE;
```

它的控制文件需要 `pg_lake_engine`、`pg_lake_iceberg` 和 `pg_lake_table`。部署还需要在 `pg_extension_base` 中安装 `shared_preload_libraries`，并运行一个正在运行的 `pgduck_server`。

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

CSV 和压缩输出使用标准看起来的 `COPY` 选项扩展为湖写入器：

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
- Version `3.4` 在 `pg_lake_copy` 中没有增加任何用户可见的 SQL 对象；其从 `3.3` 到 `3.4` 的升级脚本为空。
