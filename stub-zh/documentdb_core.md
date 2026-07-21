## 用法

来源：

- [DocumentDB v0.114-0 README](https://github.com/documentdb/documentdb/blob/v0.114-0/README.md)
- [`documentdb_core`控制文件](https://github.com/documentdb/documentdb/blob/v0.114-0/pg_documentdb_core/documentdb_core.control)
- [BSON SQL定义](https://github.com/documentdb/documentdb/blob/v0.114-0/pg_documentdb_core/sql/udfs/bson_io/bson_io--latest.sql)
- [官方预加载助手](https://github.com/documentdb/documentdb/blob/v0.114-0/scripts/preload_libraries.sh)

`documentdb_core`是DocumentDB使用的低层BSON类型和操作符层。通常它作为`documentdb`的依赖项安装，自身不提供集合CRUD、MongoDB网络协议或网关。

### 配置与安装

必须通过`pg_documentdb_core`加载`shared_preload_libraries`，然后重启PostgreSQL：

```conf
shared_preload_libraries = 'pg_documentdb_core'
```

对于完整的单节点堆栈，官方助手还会预加载`pg_cron`和`pg_documentdb`。在常规部署中安装父扩展：

```sql
CREATE EXTENSION documentdb CASCADE;
```

直接安装仅适用于低层BSON工作：

```sql
CREATE EXTENSION documentdb_core;
```

该扩展为超级用户专用且不可重定位。

### BSON 工作流

```sql
SELECT '{"name":"Ada","score":42}'::documentdb_core.bson;

SELECT documentdb_core.bson_get_value_text(
  '{"name":"Ada","score":42}'::documentdb_core.bson,
  'name'
);
```

除非`documentdb_core`在`search_path`中，否则请显式使用模式限定名。

### 重要对象

- `documentdb_core.bson`存储BSON文档。
- `documentdb_core.bsonquery`表示用于DocumentDB计划器和操作符层的BSON查询值。
- `documentdb_core.bsonsequence`表示BSON值序列。
- `bson_get_value` 和 `bson_get_value_text`，通过`->`和`->>`也暴露出来，从BSON文档中提取路径。
- `bson_from_bytea`, `bson_to_bytea`, `bson_json_to_bson` 和 `bson_to_json_string` 支持序列化边界。
- `bson_btree_ops` 和 `bson_hash_ops` 提供更高层所需的比较和哈希支持。

### 操作边界

BSON比较、索引和数值语义遵循DocumentDB的实现，不应假设与PostgreSQL `jsonb`匹配。大多数对象是`documentdb`的基础架构；寻求集合和MongoDB命令的应用程序应使用父扩展或网关而非直接构建在内部类型上。

版本0.114-0保持`documentdb_core`与整个DocumentDB堆栈一致。上游变更日志未标识此发布单独的用户核心API迁移，因此没有新的独立工作流程声明。
