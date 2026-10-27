## 用法

来源：

- [Official README](https://github.com/pinecone-io/pg_pinecone/blob/ca5fc2261429aee036b899e372c5772794c9b963/README.md)
- [Control file](https://github.com/pinecone-io/pg_pinecone/blob/ca5fc2261429aee036b899e372c5772794c9b963/pg_pinecone/pg_pinecone.control)
- [Changelog](https://github.com/pinecone-io/pg_pinecone/blob/ca5fc2261429aee036b899e372c5772794c9b963/CHANGELOG.md)
- [License](https://github.com/pinecone-io/pg_pinecone/blob/ca5fc2261429aee036b899e372c5772794c9b963/LICENSE)

`pg_pinecone` 0.11.0 将 Pinecone 兼容的向量和全文搜索引擎带入 PostgreSQL。它依赖 `vector`，需要超级用户安装，支持 PostgreSQL 17 和 18，并提供本地 `pgvector`、`native` 以及 `remote` 远程 Pinecone 引擎。

### 创建 Native 索引

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_pinecone;
CREATE TABLE items (id bigint PRIMARY KEY, embedding vector(3));
SELECT pinecone.create_index('items', 'embedding');
SELECT id FROM items ORDER BY embedding <=> '[1,2,4]' LIMIT 2;
```

0.11.0 默认使用 native 引擎，不访问网络；`pgvector` 引擎保留已有 HNSW 或 IVFFlat 路径。remote 引擎将表同步到 Pinecone serverless 索引，同时由 PostgreSQL 保留行数据。函数和设置位于 `pinecone` 模式，包括 `create_index`、`set_engine`、`topk`、`status` 和 `pinecone.api_key`。

### 安装与升级

上游 README 说明预构建归档、Docker 镜像和 pgrx 构建方式，源码构建使用 cargo-pgrx 0.19.2。安装匹配文件后，`ALTER EXTENSION pg_pinecone UPDATE` 可将 0.10.0 升级到 0.11.0；本版本调整 native 默认值，并修复删除 hnsw 索引后的清理。变更记录要求时，应重建 0.11.0 之前创建的 native 内积索引。

### 预加载与安全

native 与 `pgvector` 不需要预加载。remote worker 需要 `shared_preload_libraries = 'pg_pinecone'` 并重启，应保留已有库。remote 模式会向 Pinecone 发送向量和元数据，应保护 `pinecone.api_key`，评估网络与保留策略，并将仅超级用户可安装的扩展视为可信组件。扩展不可迁移。
