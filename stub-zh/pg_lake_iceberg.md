## 用法

来源：

- [官方Iceberg表指南](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/docs/iceberg-tables.md)
- [版本3.6控制文件](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_iceberg/pg_lake_iceberg.control)
- [Iceberg元数据SQL API](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_iceberg/pg_lake_iceberg--3.0.sql)
- [版本3.4目录FDW SQL](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_iceberg/pg_lake_iceberg--3.3--3.4.sql)

- [Upstream 3.6.0](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.6.0)
- [3.5 to 3.6 SQL](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_iceberg/pg_lake_iceberg--3.5--3.6.sql)

`pg_lake_iceberg` 在PostgreSQL内部实现了Iceberg的元数据、快照、清单、分区规范和目录集成。依赖它的 `pg_lake_table` 组件提供常用的 `CREATE TABLE ... USING iceberg` 语法；用户通常通过 `pg_lake` 安装这两个组件。

本次核验的上游发布为 `3.6.0`，SQL 扩展版本为 `3.6`，支持 PostgreSQL 16–19。现有软件包版本单独跟踪，动态库与查询服务器应使用同一发布版本。

### 创建并检查一个Iceberg表

```sql
CREATE EXTENSION pg_lake CASCADE;

SET pg_lake_iceberg.default_location_prefix =
    's3://analytics-bucket/warehouse';

CREATE TABLE events (
    event_time timestamptz NOT NULL,
    user_id bigint NOT NULL,
    payload jsonb
) USING iceberg
WITH (partition_by = 'day(event_time), bucket(32, user_id)');

SELECT table_namespace, table_name, metadata_location
FROM iceberg_tables
WHERE table_name = 'events';
```

检查Iceberg元数据文件及其引用的状态：

```sql
SELECT lake_iceberg.metadata(metadata_location)
FROM iceberg_tables
WHERE table_name = 'events';

SELECT f.*
FROM iceberg_tables AS t
CROSS JOIN LATERAL lake_iceberg.files(t.metadata_location) AS f
WHERE t.table_name = 'events';
```

### 元数据和目录API

- `iceberg_tables`: `pg_catalog` 视图，结合本地管理表和外部目录条目。
- `iceberg_namespace_properties`: 目录命名空间属性。
- `lake_iceberg.metadata(uri)`: 原始Iceberg元数据JSON。
- `lake_iceberg.files(uri)`: 清单路径、内容类型、数据文件路径/格式、规范ID、记录数和文件大小。
- `lake_iceberg.snapshots(uri)`: 序列号、快照ID、时间戳和清单列表路径。
- `lake_iceberg.data_file_stats(uri)`: 每个文件的序列号及下限/上限；执行权限授予 `lake_read` 而非 `PUBLIC`。
- `iceberg_catalog`: 版本3.4 FDW，用于命名PostgreSQL、对象存储或REST目录配置。

定义一个用户管理的REST目录服务器，并将凭据保存在用户映射中：

```sql
CREATE SERVER my_polaris TYPE 'rest'
FOREIGN DATA WRAPPER iceberg_catalog
OPTIONS (rest_endpoint 'https://polaris.example.com');

CREATE USER MAPPING FOR app_role SERVER my_polaris
OPTIONS (client_id 'app', client_secret 'secret');

CREATE TABLE catalog_events (id bigint)
USING iceberg
WITH (catalog = 'my_polaris');
```

### 目录和存储注意事项

- 用户创建的目录服务器需要自己的 `USER MAPPING` 凭证，不会回退到内置的REST目录凭证GUC。
- 内置的 `postgres`、`object_store` 和 `rest` 目录映射到不可变的扩展拥有者服务器。通过文档中的GUC进行配置，而不是修改这些服务器。
- 对 `iceberg_tables` 的外部修改默认被阻止，因为更改元数据可能会破坏pg_lake的事务和查询引擎一致性。
- Iceberg写操作应分批执行。每次语句可以添加Parquet文件和快照；常规 `VACUUM` 会压缩小文件并根据表/GUC策略过期数据。
- Iceberg对某些PostgreSQL值有更窄的表现形式。默认的 `out_of_range_values = 'error'` 保持完整性；`clamp` 默默改变超出范围的时间戳值，并用 `NULL` 替换一些不支持的值。

Iceberg 从 `3.4` 到 `3.5` 的 SQL 迁移脚本为空；补丁级目录与对象删除修复仍须依靠匹配的 `3.5.3` 二进制。

### 上游 3.6.0

上游 3.6.0 增加 PostgreSQL 19 支持，并使用 DuckDB 1.5.6。应部署匹配动态库与 `pgduck_server`，保持预加载 `pg_extension_base`，重启加载旧库的进程，并执行所提供的 SQL 升级到 3.6。此组件从 3.5 到 3.6 的 SQL 迁移为空，但运行库及依赖仍须按同一发布协调升级。
