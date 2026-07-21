## 用法

来源：

- [DocumentDB v0.114-0 README](https://github.com/documentdb/documentdb/blob/v0.114-0/README.md)
- [DocumentDB v0.114-0 changelog](https://github.com/documentdb/documentdb/blob/v0.114-0/CHANGELOG.md)
- [`documentdb` control file](https://github.com/documentdb/documentdb/blob/v0.114-0/pg_documentdb/documentdb.control)
- [Official preload helper](https://github.com/documentdb/documentdb/blob/v0.114-0/scripts/preload_libraries.sh)

`documentdb` 是 PostgreSQL 的公共 API 扩展，用于 DocumentDB，这是一个基于 PostgreSQL 开源的 MongoDB 兼容文档数据库。它存储 BSON 文档并实现 CRUD、聚合、全文搜索、地理空间和向量工作流。MongoDB 驱动程序需要单独的 DocumentDB 网关；仅安装此扩展不会暴露 Wire-Protocol 监听器，而是 PostgreSQL API。

### 配置与安装

官方部署助手使用 `pg_cron` 预加载核心库和 API 库。更改此设置后，请重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_cron, pg_documentdb_core, pg_documentdb'
```

安装公共扩展及其声明的依赖项：

```sql
CREATE EXTENSION documentdb CASCADE;
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

v0.114-0 标签的变更日志默认启用了模式验证，修复了验证器传播和缓存问题，并启用了非阻塞唯一有序索引构建。它还记录了网关配置、连通性检查、TLS 和凭证处理改进。该变更日志中的两个 RUM 优化仍受功能开关控制且默认禁用；请勿将它们描述为已启用行为。

MongoDB 兼容性并不等同于所有 MongoDB 服务器版本。应测试应用实际使用的操作符、索引行为、事务、模式验证、身份验证和驱动行为。请确保 `documentdb`、`documentdb_core`、网关及可选的分布式/索引组件来自同一发行系列。
