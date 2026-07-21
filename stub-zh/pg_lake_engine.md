## 用法

来源：

- [官方pg_lake架构概述](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/README.md#architecture)
- [版本3.4控制文件](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_engine/pg_lake_engine.control)
- [基础SQL对象](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_engine/pg_lake_engine--3.0.sql)
- [版本3.4清理队列变更](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_engine/pg_lake_engine--3.3--3.4.sql)

`pg_lake_engine`是pg_lake表、复制和Iceberg扩展共享的执行层。它重写符合条件的PostgreSQL工作以供`pgduck_server`使用，映射PostgreSQL和DuckDB值，并跟踪在中止或表变更后必须删除的远程文件。它是内部依赖项而非独立的分析接口。

### 部署边界

通过顶级扩展安装它，使其依赖图和预加载指令保持一致：

```conf
shared_preload_libraries = 'pg_extension_base'
```

```sql
CREATE EXTENSION pg_lake CASCADE;
```

`pg_lake_engine`需要`pg_extension_base`和`pg_map`。查询执行还需要一个运行中的本地`pgduck_server`；仅创建此扩展不会提供可用的湖表。

### 用户可见对象

- 角色`lake_read`、`lake_write`和`lake_read_write`：被其他组件消费的共享权限组。
- `to_postgres(any)`：返回其输入并强制该表达式在PostgreSQL中而不是推送到湖引擎中进行评估。
- `to_date(double precision)`：将常见的Parquet中的天数自Unix纪元值转换为PostgreSQL `date`。
- `lake_engine.deletion_queue`：跟踪已提交的孤儿文件清理；可由`lake_write`读取。
- `lake_engine.in_progress_files`：跟踪未提交事务产生的文件。
- `lake_engine.flush_deletion_queue(regclass)`和`flush_in_progress_queue()`：用于维护路径的特权清理函数。

```sql
SELECT to_postgres(application_only_function(payload))
FROM external_events;
```

仅在表达式无法或不应下推时使用`to_postgres()`；将数据拉回PostgreSQL可能会显著增加传输和执行成本。

### 内部状态与注意事项

- `__lake__internal__nsp__`函数是规划器/解析器占位符，不是受支持的直接SQL API。
- 不要手动更新或删除队列行。清理函数需要扩展的对象存储凭证和特权角色，并且仅应在操作工具文档中指定的情况下调用。
- 版本`3.4`将`resolve_metadata`添加到删除队列，以便在`VACUUM`期间可以将Iceberg元数据展开为精确引用的文件，从而将对象存储遍历移出`DROP`路径。
- 角色是集群范围的对象，在一个数据库中可能会超出扩展实例的存在；在移除pg_lake时需单独审查成员资格。
