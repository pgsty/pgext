## 用法

来源：[README](https://github.com/documentdb/documentdb/blob/v0.113-0/README.md)、[CHANGELOG](https://github.com/documentdb/documentdb/blob/v0.113-0/CHANGELOG.md)、[documentdb.control](https://github.com/documentdb/documentdb/blob/v0.113-0/pg_documentdb/documentdb.control)、[documentdb_core.control](https://github.com/documentdb/documentdb/blob/v0.113-0/pg_documentdb_core/documentdb_core.control)、[documentdb_extended_rum.control](https://github.com/documentdb/documentdb/blob/v0.113-0/pg_documentdb_extended_rum/documentdb_extended_rum.control)

`documentdb` 是以 PostgreSQL 扩展实现的 MongoDB 兼容文档数据库。它在 PostgreSQL 中加入 BSON 存储和 API，并可通过可选 gateway 层服务 MongoDB wire-protocol 客户端。FerretDB 2.0+ 可以使用 DocumentDB 作为后端。

### 组件

公开扩展接口拆分为多个相关扩展：

- `documentdb_core`：BSON 数据类型与底层 BSON 操作。
- `documentdb`：面向文档 CRUD 与查询行为的公开 API。
- `documentdb_extended_rum`：DocumentDB 索引使用的 extended RUM access method。
- `pg_documentdb_gw`：本地 Docker 镜像和 MongoDB 兼容客户端使用的 gateway protocol layer。

在需要 API 的每个数据库中安装 SQL 扩展：

```sql
CREATE EXTENSION IF NOT EXISTS documentdb CASCADE;
```

`0.113-0` 的 `documentdb.control` 声明依赖 `documentdb_core`、`pg_cron`、`tsm_system_rows`、`vector` 和 `postgis`。gateway container 监听 MongoDB 兼容端口；README 示例使用 `10260`，避免与本机 MongoDB 服务冲突。

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

上游 README 也展示了通过普通 MongoDB driver 执行 aggregation pipeline：

```python
pipeline = [
    {"$match": {"name": "Alice"}},
    {"$project": {"_id": 0, "name": 1, "email": 1}},
]

for doc in coll.aggregate(pipeline):
    print(doc)
```

### 版本说明

本项目 CSV 跟踪 DocumentDB `0.113`，覆盖 PostgreSQL 15-18。上游 tag 为 `v0.113-0`；control files 报告 `default_version = '0.113-0'`。

`0.111` 到 `0.113` 的 changelog 主要是 query planner、collation 与 index correctness 方面的工作：

- `0.113-0` 为带 `$in` 和 `$nin` 的非唯一 ordered indexes 增加 opt-in collation support，并在 feature flag 后支持清理 ordered TTL indexes 上的 dead index entries。
- `0.112-0` 移除较旧的 composite-returning `bson_update_document` UDF 路径，扩展非唯一 ordered-index collation support，并改进 `$group` 和 accumulator 执行。
- `0.111-0` 增加后台初始化 job infrastructure、更多 `$group` 校验、collation/index pushdown 改进，以及多个 crash fix。

### 注意事项

- DocumentDB 是多扩展栈；`CREATE EXTENSION documentdb CASCADE` 是通常入口，但如果需要 MongoDB wire compatibility，实际部署还需要 gateway/runtime pieces。
- changelog 中部分功能受 `documentdb.*` feature flags 控制。将行为写成 always-on 前，先核对确切安装版本中的 flag 默认值。
- `documentdb_extended_rum` 可 relocation，但 `documentdb` 和 `documentdb_core` 不可 relocation。
