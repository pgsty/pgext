

## 用法

> [documentdb_core: DocumentDB for PostgreSQL 的核心 API 接口](https://github.com/documentdb/documentdb)

DocumentDB 在 PostgreSQL 上提供 MongoDB 兼容的文档数据库功能。`documentdb_core` 扩展为原生 Postgres 引入了 BSON 数据类型支持和操作。

### BSON 数据类型

该扩展为 PostgreSQL 添加了原生 BSON（二进制 JSON）数据类型，支持 MongoDB 风格文档的存储和操作。

### 基本文档操作

文档通过 DocumentDB API 层的 MongoDB 兼容 CRUD 操作进行管理：

```python
import pymongo

client = pymongo.MongoClient(
    'mongodb://user:pass@localhost:10260/?tls=true&tlsAllowInvalidCertificates=true'
)

db = client["myDatabase"]
collection = db.create_collection("myCollection")

# 插入文档
collection.insert_one({
    'name': 'John Doe',
    'email': 'john@email.com',
    'address': '123 Main St'
})

collection.insert_many([
    {'name': 'Jane Smith', 'email': 'jane@email.com'},
    {'name': 'Alice Johnson', 'email': 'alice@email.com'}
])

# 查询文档
for doc in collection.find():
    print(doc)

single = collection.find_one({'name': 'John Doe'})
```

### 聚合管道

```python
pipeline = [
    {'$match': {'name': 'Alice Johnson'}},
    {'$project': {'_id': 0, 'name': 1, 'email': 1}}
]

results = collection.aggregate(pipeline)
for doc in results:
    print(doc)
```

### 组件

- **documentdb_core**：原生 Postgres 的 BSON 数据类型支持和操作
- **documentdb (pg_documentdb)**：提供 CRUD 功能的公共 API 接口
- **pg_documentdb_gw**：网关协议转换层（MongoDB 线协议到 PostgreSQL）

该扩展支持对 BSON 文档的全文搜索、地理空间查询和向量搜索。
