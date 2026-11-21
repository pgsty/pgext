## 用法

来源：

- [3.7.0 pg_lake_table.control](https://github.com/Snowflake-Labs/pg_lake/blob/v3.7.0/pg_lake_table/pg_lake_table.control)
- [3.7.0 v3.7.0](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.7.0)
- [3.7.0 table-options.md](https://github.com/Snowflake-Labs/pg_lake/blob/v3.7.0/docs/table-options.md)
- [3.7.0 data-types.md](https://github.com/Snowflake-Labs/pg_lake/blob/v3.7.0/docs/data-types.md)
- [3.7.0 iceberg-catalogs.md](https://github.com/Snowflake-Labs/pg_lake/blob/v3.7.0/docs/iceberg-catalogs.md)

- [当前上游](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.6.1)
- [当前上游](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.1/docs/table-options.md)
- [当前上游](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.1/docs/iceberg-catalogs.md)

- [官方数据湖文件查询指南](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/docs/query-data-lake-files.md)
- [官方Iceberg表指南](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/docs/iceberg-tables.md)
- [版本3.6控制文件](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_table/pg_lake_table.control)
- [3.4 至 3.5 目录索引升级](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_table/pg_lake_table--3.4--3.5.sql)
- [FDW、服务器、实用工具和访问方法SQL](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_table/pg_lake_table--3.0.sql)

- [Upstream 3.6.0](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.6.0)
- [3.5 to 3.6 SQL](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_table/pg_lake_table--3.5--3.6.sql)

`pg_lake_table` 将对象存储文件作为 PostgreSQL 外部表暴露，并提供 `USING iceberg` 表语法。它拥有 `pg_lake` 和 `pg_lake_iceberg` 外部服务器、文件检查/缓存实用工具、表目录和事务挂钩；Iceberg 元数据编码委托给 `pg_lake_iceberg`。


本次核验的上游发布为 `3.6.0`，SQL 扩展版本为 `3.6`，支持 PostgreSQL 16–19。现有软件包版本单独跟踪，动态库与查询服务器应使用同一发布版本。

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

### 运维注意事项

- 必须预加载 `pg_extension_base`，并在每个执行湖查询的 PostgreSQL 主机上运行 `pgduck_server`，为其配置所需位置的访问凭据。
- `lake_read`、`lake_write` 和 `lake_read_write` 控制对外部服务器、模式和实用工具的访问。为每个应用程序授予最窄的角色权限。
- 外部表是文件的引用，而不是导入的副本。文件替换、跨区域访问和缓存失效可以独立于 PostgreSQL 目录状态改变延迟或结果。
- Iceberg 插入适合批量处理。对于高频率逐行插入，使用暂存堆表，并定期刷新批次。
- 内部 `lake_table.*` 目录跟踪文件、字段 ID、分区和恢复状态。不要直接修改它们。


版本 `3.5` 为文件删除映射、列统计与分区值目录新增索引，避免批量删除时反复扫描子目录。除安装新动态库外，还须执行 SQL 迁移。

### 上游 3.6.0

上游 3.6.0 增加 PostgreSQL 19 支持，并使用 DuckDB 1.5.6。应部署匹配动态库与 `pgduck_server`，保持预加载 `pg_extension_base`，重启加载旧库的进程，并执行所提供的 SQL 升级到 3.6。表组件注册独立的对象存储目录导出 worker，并将目录名称辅助函数改为 STABLE。必须执行 SQL 迁移，仅替换二进制不会注册该 worker。

### 上游 3.6.1 目录行为

通过 `catalog = 'snowflake'` 挂接且自动推断列的只读表，现在默认将 `lowercase_column_names` 设为 true。其他文档描述的只读目录场景保留原有默认值；该选项在建表时固定，不会重命名已有列。挂接新表时应检查推断列名与大小写冲突。

整个家族的 SQL 版本仍为 3.6。扩展库与 `pgduck_server` 需要匹配部署，该补丁也在查询服务中包含 Brotli 解码器安全修正。既有软件包版本单独记录，不能证明已安装这一上游补丁。

### 上游 3.7.0

当前核对的发行版为 3.7.0，本组件的控制版本为 3.7。前面的旧版本章节描述历史行为和软件包示例。应使用匹配的库与 `pgduck_server` 协同升级整个扩展族，重启持有旧代码的进程，并执行提供的扩展升级脚本，将 SQL 版本更新至 3.7。这不代表相应软件包已经可用。

新挂接的只读 `catalog = 'snowflake'` 表默认启用 `lowercase_column_names`，该选项在建表时固定，大小写折叠冲突会报错。对于 PostgreSQL 17+ 的无限间隔，`out_of_range_values = 'error'` 会拒绝写入，`clamp` 则静默保存为 NULL。启用前应评估这种信息丢失。
