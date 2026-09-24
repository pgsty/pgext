## 用法

来源：

- [DocumentDB v0.117-0 README](https://github.com/documentdb/documentdb/blob/v0.117-0/README.md)
- [DocumentDB v0.117-0 changelog](https://github.com/documentdb/documentdb/blob/v0.117-0/CHANGELOG.md)
- [`documentdb` control file](https://github.com/documentdb/documentdb/blob/v0.117-0/pg_documentdb/documentdb.control)
- [Official preload helper](https://github.com/documentdb/documentdb/blob/v0.117-0/scripts/preload_libraries.sh)

`documentdb` 是 PostgreSQL 的公共 API 扩展，用于 DocumentDB，这是一个基于 PostgreSQL 开源的 MongoDB 兼容文档数据库。它存储 BSON 文档并实现 CRUD、聚合、全文搜索、地理空间和向量工作流。MongoDB 驱动程序需要单独的 DocumentDB 网关；仅安装此扩展不会暴露 Wire-Protocol 监听器，而是 PostgreSQL API。

### 配置与安装

官方部署助手使用 `pg_cron` 预加载核心库和 API 库。更改此设置后，请重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_cron, pg_documentdb_core, pg_documentdb, pg_documentdb_extended_rum'
```

安装公共扩展及其声明的依赖项：

```sql
CREATE EXTENSION documentdb CASCADE;
CREATE EXTENSION documentdb_extended_rum;
```

`CASCADE` 可以在文件存在时安装 `documentdb_core`、`pg_cron`、`tsm_system_rows`、`vector` 和 `postgis`。安装仅限超级用户且不可重定位。

### 原生 SQL 工作流

SQL 接口使用数据库名、集合名和 BSON 命令文档：

```sql
SELECT documentdb_api.create_collection('appdb', 'people');

SELECT documentdb_api.insert_one(
  'appdb',
  'people',
  '{"_id": 1, "name": "Ada", "team": "storage"}',
  NULL
);

SELECT document
FROM documentdb_api_catalog.bson_aggregation_find(
  'appdb',
  '{"find":"people","filter":{"team":"storage"}}'
);
```

为了应用程序兼容性，请运行网关并在其配置的 TLS 端点上使用受支持的 MongoDB 驱动程序。网关将 Wire-Protocol 命令转换为此 PostgreSQL API。

### 重要对象

- `documentdb_api` 包含集合管理及命令函数，如 `create_collection` 和 `insert_one`。
- `documentdb_api_catalog.bson_aggregation_find` 执行 MongoDB 风格的查找规范并返回 BSON 文档。
- `documentdb_core.bson` 是由 `documentdb_core` 提供的存储和交换类型。
- DocumentDB 角色和内部模式将公共读写操作与管理及实现对象分开。
- `documentdb.enableNonBlockingUniqueIndexBuild` 控制 v0.114 路径下的后台唯一有序索引构建，并在该版本中默认启用。

### 版本与运行注意事项

0.117-0 增加了考虑排序规则的分组与最小/最大值处理，并包含 0.116-0 引入的 JSON Schema enum 与 oneOf 支持。标量聚合索引下推受功能开关控制，默认关闭。所有受支持 PostgreSQL 主版本的 `documentdb.rum_library_load_option` 现在默认是 `require_documentdb_extended_rum`，部署时须提供匹配的扩展 RUM 库。

MongoDB 兼容性并不等同于所有 MongoDB 服务器版本。应测试应用实际使用的操作符、索引行为、事务、模式验证、身份验证和驱动行为。请确保 `documentdb`、`documentdb_core`、网关及可选的分布式/索引组件来自同一发行系列。
