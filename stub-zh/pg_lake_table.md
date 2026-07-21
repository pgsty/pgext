## 用法

来源：

- [官方数据湖文件查询指南](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/query-data-lake-files.md)
- [官方Iceberg表指南](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/iceberg-tables.md)
- [版本3.4控制文件](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_table/pg_lake_table.control)
- [FDW、服务器、实用工具和访问方法SQL](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_table/pg_lake_table--3.0.sql)

`pg_lake_table` 将对象存储文件作为 PostgreSQL 外部表暴露，并提供 `USING iceberg` 表语法。它拥有 `pg_lake` 和 `pg_lake_iceberg` 外部服务器、文件检查/缓存实用工具、表目录和事务挂钩；Iceberg 元数据编码委托给 `pg_lake_iceberg`。

### 查询外部文件

安装完整的堆栈，包括所需的 `pg_lake_engine`、`pg_lake_iceberg` 和 `btree_gist` 依赖项：

```sql
CREATE EXTENSION pg_lake CASCADE;
```

空列列表要求 pg_lake 推断文件模式：

```sql
CREATE FOREIGN TABLE external_events ()
SERVER pg_lake
OPTIONS (
    path 's3://analytics-bucket/events/*.parquet',
    filename 'true'
);

SELECT count(*) FROM external_events;
```

通过扩展提供的表访问方法创建可写的 Iceberg 表：

```sql
CREATE TABLE managed_events (
    event_time timestamptz,
    payload jsonb
) USING iceberg;
```

### 文件和表实用工具索引

- `lake_file.list(pattern)`: 列出匹配的对象路径、大小、修改时间和 ETags。
- `lake_file.size(path)` 和 `lake_file.exists(path)`: 检查一个远程对象。
- `lake_file.preview(url, format, compression)`: 返回推断的列名和类型。
- `lake_file.delete(url)`: 删除一个远程对象；仅限允许删除数据的角色执行。
- `lake_file_cache.add(path, refresh)`, `remove(path)` 和 `list()`: 管理 `lake_read` 成员的本地文件缓存。
- `lake_iceberg.table_size(regclass)`: 计算 Iceberg 表当前数据文件的总大小。
- 外部服务器 `pg_lake` 和 `pg_lake_iceberg`: 读取文件和可写的 Iceberg 入口点。
- 访问方法 `iceberg` 和 `pg_lake_iceberg`: 用于将 `CREATE TABLE ... USING` 转换为扩展的外部表表示形式的别名。

```sql
SELECT path, file_size
FROM lake_file.list('s3://analytics-bucket/events/**/*.parquet');

SELECT *
FROM lake_file.preview('s3://analytics-bucket/events/sample.parquet');
```

### 运营注意事项

- 必须预加载 `pg_extension_base` 并在每个引用位置运行 `pgduck_server`，并带有相应的凭据。
- `lake_read`、`lake_write` 和 `lake_read_write` 控制对外部服务器、模式和实用工具的访问。为每个应用程序授予最窄的角色权限。
- 外部表是文件的引用，而不是导入的副本。文件替换、跨区域访问和缓存失效可以独立于 PostgreSQL 目录状态改变延迟或结果。
- Iceberg 插入优化批量插入而非单行插入。对于高频率逐行插入，使用暂存堆表，并定期刷新批次。
- 内部 `lake_table.*` 目录跟踪文件、字段 ID、分区和恢复状态。不要直接修改它们。

文档结束
