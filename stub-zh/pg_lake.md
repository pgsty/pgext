## 用法

来源：

- [官方 pg_lake README](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/README.md)
- [版本3.4控制文件](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake/pg_lake.control)
- [官方构建和启动指南](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/building-from-source.md)
- [官方项目文档索引](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/README.md)

`pg_lake`是Snowflake的PostgreSQL湖库堆栈中的顶级扩展。它安装了查询对象存储文件所需的表、Iceberg、复制、查询引擎、基础扩展和映射组件，从而创建事务性的Iceberg表。这些PostgreSQL扩展协调规划和事务处理，而一个单独的本地`pgduck_server`进程使用DuckDB执行向量化工作。

### 启动堆栈

版本`3.4`支持PostgreSQL 16至18。预加载通用扩展基础设施，重启PostgreSQL，并在数据库主机上启动`pgduck_server`：

```conf
shared_preload_libraries = 'pg_extension_base'
```

```shell
pgduck_server --cache_dir /var/cache/pg_lake
```

在目标数据库中创建完整的依赖树：

```sql
CREATE EXTENSION pg_lake CASCADE;
SELECT lake.version();
```

为`pgduck_server`配置对象存储凭证，然后选择托管的Iceberg位置：

```sql
SET pg_lake_iceberg.default_location_prefix =
    's3://analytics-bucket/warehouse';
```

### 核心工作流

创建和修改事务性的Iceberg表：

```sql
CREATE TABLE measurements (
    station_name text NOT NULL,
    measured_at timestamptz NOT NULL,
    value double precision
) USING iceberg;

INSERT INTO measurements VALUES
    ('Istanbul', now(), 18.5),
    ('Haarlem', now(), 9.3);
```

通过`COPY`导入或导出Parquet、CSV或换行符分隔的JSON文件：

```sql
COPY (SELECT * FROM measurements)
TO 's3://analytics-bucket/export/measurements.parquet';

COPY measurements
FROM 's3://analytics-bucket/import/measurements.parquet';
```

查询文件而不将其加载到PostgreSQL中：

```sql
CREATE FOREIGN TABLE external_events ()
SERVER pg_lake
OPTIONS (path 's3://analytics-bucket/events/*.parquet');

SELECT count(*) FROM external_events;
```

### 组件索引

- `pg_lake`：元扩展和`lake.version()`。
- `pg_lake_table`：数据湖FDW、Iceberg表语法、文件工具和表目录。
- `pg_lake_iceberg`：Iceberg的元数据、快照、清单和目录集成。
- `pg_lake_copy`：对对象存储文件和湖格式的`COPY`拦截。
- `pg_lake_engine`：共享查询重写、类型转换、清理和`pgduck_server`客户端层。
- `pg_extension_base`：预加载和生命周期工作进程基础设施。
- `pg_map`：用于嵌套湖数据的生成PostgreSQL映射类型。

### 运营注意事项

- 对于湖查询，`pgduck_server`是必需的，并且必须具有从PostgreSQL到对象存储的有效凭证和本地套接字连接。
- S3及其兼容凭证由DuckDB的秘密/凭证链解析。仅授予工作负载所需的桶权限。
- Iceberg写入按语句创建Parquet文件。批量插入并定期运行`VACUUM`以避免产生许多小文件。
- 这些PostgreSQL扩展、`pgduck_server`对象存储数据和Iceberg目录形成一个部署单元。单独创建扩展不足以证明外部服务可用，需要分别备份和升级它们。
