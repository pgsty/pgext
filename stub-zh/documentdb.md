


## 用法

来源：[README](https://github.com/documentdb/documentdb/blob/v0.113-0/README.md)、[CHANGELOG](https://github.com/documentdb/documentdb/blob/v0.113-0/CHANGELOG.md)、[documentdb.control](https://github.com/documentdb/documentdb/blob/v0.113-0/pg_documentdb/documentdb.control)、[documentdb_core.control](https://github.com/documentdb/documentdb/blob/v0.113-0/pg_documentdb_core/documentdb_core.control)、[documentdb_extended_rum.control](https://github.com/documentdb/documentdb/blob/v0.113-0/pg_documentdb_extended_rum/documentdb_extended_rum.control)

`documentdb` 是以 PostgreSQL 扩展实现的 MongoDB 兼容文档数据库。它在 PostgreSQL 中加入 BSON 存储和 API，并可通过可选网关层服务 MongoDB 线协议客户端。FerretDB 2.0+ 可以使用 DocumentDB 作为后端。

### 组件

公开扩展接口拆分为多个相关扩展：

- `documentdb_core`：BSON 数据类型与底层 BSON 操作。
- `documentdb`：面向文档 CRUD 与查询行为的公开 API。
- `documentdb_extended_rum`：DocumentDB 索引使用的扩展 RUM 访问方法。
- `pg_documentdb_gw`：本地 Docker 镜像和 MongoDB 兼容客户端使用的网关协议层。

在需要 API 的每个数据库中安装 SQL 扩展：

```sql
CREATE EXTENSION IF NOT EXISTS documentdb CASCADE;
```

`0.113-0` 的 `documentdb.control` 声明依赖 `documentdb_core`、`pg_cron`、`tsm_system_rows`、`vector` 和 `postgis`。网关容器监听 MongoDB 兼容端口；README 示例使用 `10260`，避免与本机 MongoDB 服务冲突。

### MongoDB 客户端示例

```python
import pymongo

client = pymongo.MongoClient(
    "mongodb://user:pass@localhost:10260/?tls=true&tlsAllowInvalidCertificates=true"
)

db = client["quickStartDatabase"]
coll = db.create_collection("quickStartCollection")

coll.insert_one({"name": "Alice", "email": "alice@example.com"})
print(coll.find_one({"name": "Alice"}))
```

上游 README 也展示了通过普通 MongoDB 驱动程序执行聚合流水线：

```python
pipeline = [
    {"$match": {"name": "Alice"}},
    {"$project": {"_id": 0, "name": 1, "email": 1}},
]

for doc in coll.aggregate(pipeline):
    print(doc)
```

### 版本说明

本项目 CSV 跟踪 DocumentDB `0.113`，覆盖 PostgreSQL 15-18。上游标签为 `v0.113-0`；控制文件报告 `default_version = '0.113-0'`。

`0.111` 到 `0.113` 的变更日志主要是查询规划器、排序规则与索引正确性方面的工作：

- `0.113-0` 为带 `$in` 和 `$nin` 的非唯一有序索引增加可选排序规则支持，并在功能开关后支持清理有序 TTL 索引上的失效索引项。
- `0.112-0` 移除较旧的返回复合类型的 `bson_update_document` UDF 路径，扩展非唯一有序索引的排序规则支持，并改进 `$group` 和累加器执行。
- `0.111-0` 增加后台初始化作业基础设施、更多 `$group` 校验、排序规则/索引下推改进，以及多个崩溃修复。

### 注意事项

- DocumentDB 是多扩展栈；`CREATE EXTENSION documentdb CASCADE` 是通常入口，但如果需要 MongoDB 线协议兼容性，实际部署还需要网关和运行时组件。
- 变更日志中部分功能受 `documentdb.*` 功能开关控制。将行为写成始终启用前，先核对确切安装版本中的开关默认值。
- `documentdb_extended_rum` 可迁移，但 `documentdb` 和 `documentdb_core` 不可迁移。
