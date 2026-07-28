## 用法

来源：

- [官方上游 README](https://github.com/sahilchug/rxpgvector/blob/bef81a80011f1abcf197f8aea7db1e514fca04a0/README.md)
- [官方扩展控制文件 (pgvector.control)](https://github.com/sahilchug/rxpgvector/blob/bef81a80011f1abcf197f8aea7db1e514fca04a0/pgvector.control)
- [官方实现源代码](https://github.com/sahilchug/rxpgvector/blob/bef81a80011f1abcf197f8aea7db1e514fca04a0/src/lib.rs)

`pgvector` — 自定义向量类型，支持欧几里得距离、余弦相似度和 IVFFlat 帮助器。使用它来对应向量、模型或检索工作流。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgvector;

CREATE TABLE test_vectors (
    id SERIAL PRIMARY KEY,
    embedding PGVector
);
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `create_vector` 是一个扩展函数。
- `euclidean_distance` 是一个扩展函数。
- `ivfflat_incex_create` 是一个扩展函数。
- `ivfflat_index_search` 是一个扩展函数。
- `vector_cosine_similarity` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
