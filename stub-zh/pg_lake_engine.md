## 用法

来源：

- [3.7.0 pg_lake_engine.control](https://github.com/Snowflake-Labs/pg_lake/blob/v3.7.0/pg_lake_engine/pg_lake_engine.control)
- [3.7.0 v3.7.0](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.7.0)

- [Upstream source](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.6.2)

- [当前上游](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.6.1)
- [当前上游](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.1/docs/table-options.md)
- [当前上游](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.1/docs/iceberg-catalogs.md)

- [官方pg_lake架构概述](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/README.md#architecture)
- [版本3.6控制文件](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_engine/pg_lake_engine.control)
- [基础SQL对象](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_engine/pg_lake_engine--3.0.sql)
- [3.5 清理重试状态迁移](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_engine/pg_lake_engine--3.4--3.5.sql)

- [Upstream 3.6.0](https://github.com/Snowflake-Labs/pg_lake/releases/tag/v3.6.0)
- [3.5 to 3.6 SQL](https://github.com/Snowflake-Labs/pg_lake/blob/v3.6.0/pg_lake_engine/pg_lake_engine--3.5--3.6.sql)

`pg_lake_engine` 是pg_lake表、复制和Iceberg扩展共享的执行层。它重写符合条件的PostgreSQL工作以供 `pgduck_server` 使用，映射PostgreSQL和DuckDB值，并跟踪在中止或表变更后必须删除的远程文件。它是内部依赖项而非独立的分析接口。

本次核验的上游发布为 `3.6.0`，SQL 扩展版本为 `3.6`，支持 PostgreSQL 16–19。现有软件包版本单独跟踪，动态库与查询服务器应使用同一发布版本。

### 部署边界

通过顶级扩展安装它，使其依赖图和预加载指令保持一致：

```ini
shared_preload_libraries = 'pg_extension_base'
```

```sql
CREATE EXTENSION pg_lake CASCADE;
```

`pg_lake_engine` 需要 `pg_extension_base` 和 `pg_map`。查询执行还需要一个运行中的本地 `pgduck_server`；仅创建此扩展不会提供可用的湖表。

### 用户可见对象

- 角色 `lake_read`、`lake_write` 和 `lake_read_write`：被其他组件消费的共享权限组。
- `to_postgres(any)`：返回其输入并强制该表达式在PostgreSQL中而不是推送到湖引擎中进行评估。
- `to_date(double precision)`：将常见的Parquet中的天数自Unix纪元值转换为PostgreSQL `date`。
- `lake_engine.deletion_queue`：跟踪已提交的孤儿文件清理；可由 `lake_write` 读取。
- `lake_engine.in_progress_files`：跟踪未提交事务产生的文件。
- `lake_engine.flush_deletion_queue(regclass)` 和 `flush_in_progress_queue()`：用于维护路径的特权清理函数。

```sql
SELECT to_postgres(application_only_function(payload))
FROM external_events;
```

仅在表达式无法或不应下推时使用 `to_postgres()`；将数据拉回PostgreSQL可能会显著增加传输和执行成本。

### 内部状态与注意事项

- `__lake__internal__nsp__` 函数是规划器/解析器占位符，不是受支持的直接SQL API。
- 不要手动更新或删除队列行。清理函数需要扩展的对象存储凭证和特权角色，并且仅应在操作工具文档中指定的情况下调用。
- 版本 `3.5` 在 `lake_engine.deletion_queue` 中新增 `last_attempt_at`。清理重试由 `pg_lake_engine.vacuum_file_remove_retry_interval` 控制间隔，`pg_lake_engine.vacuum_file_remove_max_retries` 限制实际尝试次数，而不是统计无关的 VACUUM 执行轮次。
- 角色是集群范围的对象，在一个数据库中可能会超出扩展实例的存在；在移除pg_lake时需单独审查成员资格。

### 上游 3.6.0

上游 3.6.0 增加 PostgreSQL 19 支持，并使用 DuckDB 1.5.6。应部署匹配动态库与 `pgduck_server`，保持预加载 `pg_extension_base`，重启加载旧库的进程，并执行所提供的 SQL 升级到 3.6。引擎迁移增加内部 UUID 下推包装函数，用于规划与反解析，不是通用应用 API。

### 上游 3.6.1 运行时更新

将匹配的 3.6.1 `pgduck_server` 与扩展库一同部署。本补丁回移 Brotli 解码器安全修正，并将查询服务 S3 上传器的最大文件大小降至 50 GB。SQL 版本仍为 3.6，SQL 版本不变不能证明已部署修补后的查询服务。C 层 COPY 查询导出接口属于内部集成工作，不是该组件新增的 SQL 函数。

### 上游 3.6.2

3.6.2 补丁修正 pgduck_server 逐字节传输 CSV 值的处理。应部署与扩展库匹配的查询服务二进制，并重新连接或重启仍持有旧代码的进程。组件 SQL 版本仍为 3.6，extversion 不变不能证明查询服务修复已经安装。软件包可用性仍是单独的目录事实。

### 上游 3.7.0

当前核对的发行版为 3.7.0，本组件的控制版本为 3.7。前面的旧版本章节描述历史行为和软件包示例。应使用匹配的库与 `pgduck_server` 协同升级整个扩展族，重启持有旧代码的进程，并执行提供的扩展升级脚本，将 SQL 版本更新至 3.7。这不代表相应软件包已经可用。
