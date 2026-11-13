## 用法

来源：

- [DocumentDB v1.0-0 README](https://github.com/documentdb/documentdb/blob/v1.0-0/README.md)
- [DocumentDB v1.0-0 changelog](https://github.com/documentdb/documentdb/blob/v1.0-0/CHANGELOG.md)
- [`documentdb` control file](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb/documentdb.control)
- [Official preload helper](https://github.com/documentdb/documentdb/blob/v1.0-0/scripts/preload_libraries.sh)
- [DocumentDB 1.0 release](https://github.com/documentdb/documentdb/releases/tag/v1.0-0)
- [DocumentDB 1.2 release](https://github.com/documentdb/documentdb/releases/tag/v1.2-0)
- [1.2 control](https://github.com/documentdb/documentdb/blob/v1.2-0/pg_documentdb/documentdb.control)
- [1.1 to 1.2 SQL](https://github.com/documentdb/documentdb/blob/v1.2-0/pg_documentdb/sql/documentdb--1.1-0--1.2-0.sql)
- [Official upgrade procedure](https://github.com/documentdb/documentdb/blob/v1.2-0/packaging/README.md)

`documentdb` 是 PostgreSQL 的公共 API 扩展，用于 DocumentDB，这是一个基于 PostgreSQL 开源的 MongoDB 兼容文档数据库。它存储 BSON 文档并实现 CRUD、聚合、全文搜索、地理空间和向量工作流。MongoDB 驱动程序需要单独的 DocumentDB 网关；仅安装此扩展不会暴露 Wire-Protocol 监听器，而是 PostgreSQL API。

### 配置与安装

官方部署助手使用 `pg_cron` 预加载核心库和 API 库。更改此设置后，请重启 PostgreSQL：

```ini
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

1.0-0 是首个正式可用的发行系列。它修复了零批量读取的权限检查、合并与输出目标的写权限，以及崩溃、地理排序和索引清理问题。自定义角色 CRUD 与连接后嵌套查找仍受设置控制，默认关闭。应使用匹配的全部库，并应用该版本的升级 SQL。

官方 DocumentDB Local 容器默认启用使用情况遥测；原生软件包不受影响。容器中可设置 DOCUMENTDB_USAGE_TELEMETRY=false、NO_ANALYTICS=1 或 DO_NOT_TRACK=1 禁用。

MongoDB 兼容性并不等同于所有 MongoDB 服务器版本。应测试应用实际使用的操作符、索引行为、事务、模式验证、身份验证和驱动行为。请确保 `documentdb`、`documentdb_core`、网关及可选的分布式/索引组件来自同一发行系列。

### DocumentDB 1.2 变化

DocumentDB 1.2 的 SQL 版本为 `1.2-0`。升级前应备份数据目录；不支持将已升级数据退回 1.0。官方 1.2 容器先升级扩展再启动网关，并要求显式提供密码，已有数据卷也不例外。对于接管的 PostgreSQL 实例，应遵循包指南的重启与升级顺序，不要重启或修改无关实例。

只有设置 `documentdb_core.enableCollation = on` 后才启用排序规则，即使显式指定 `simple` 也一样。非唯一排序规则索引还需要 `documentdb.enableCollationWithNonUniqueOrderedIndexes`；唯一索引还需要 `documentdb.enable_collated_unique_indexes` 与 `documentdb.enable_composite_unique_optional_key`。此版本扩展了逐命令、逐更新的排序规则行为，并修复降序唯一索引的空值与缺失值冲突、`$dateTrunc` 的天边界以及重扫 `$densify` 的结果。升级后应复核相关查询和索引语义。
