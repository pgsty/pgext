## 用法

来源：

- [Official README](https://github.com/pinecone-io/pg_pinecone/blob/v0.12.1/README.md)
- [Control file](https://github.com/pinecone-io/pg_pinecone/blob/v0.12.1/pg_pinecone/pg_pinecone.control)
- [Changelog](https://github.com/pinecone-io/pg_pinecone/blob/v0.12.1/CHANGELOG.md)
- [Upgrade runbook](https://github.com/pinecone-io/pg_pinecone/blob/v0.12.1/docs/runbook.md)
- [Hosted embedding tutorial](https://github.com/pinecone-io/pg_pinecone/blob/v0.12.1/docs/tutorial-gemini-document-search.md)
- [License](https://github.com/pinecone-io/pg_pinecone/blob/v0.12.1/LICENSE)

`pg_pinecone` 0.12.1 将 Pinecone 兼容的向量和全文搜索引擎带入 PostgreSQL。它依赖 `vector`，需要超级用户安装，支持 PostgreSQL 17 和 18，并提供本地 `pgvector`、`native` 以及 `remote` 远程 Pinecone 引擎。

### 创建 Native 索引

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_pinecone;
CREATE TABLE items (id bigint PRIMARY KEY, embedding vector(3));
INSERT INTO items VALUES (1, '[1,2,3]'), (2, '[4,5,6]');
SELECT pinecone.create_index('items', 'embedding');
SELECT id FROM items ORDER BY embedding <=> '[1,2,4]' LIMIT 2;
```

0.12.1 的本地默认引擎仍为 native，不访问网络；`pgvector` 引擎保留已有 HNSW 或 IVFFlat 路径。remote 引擎将表同步到 Pinecone serverless 索引，同时由 PostgreSQL 保留行数据。函数和设置位于 `pinecone` 模式，包括 `create_index`、`set_engine`、`topk`、`status` 和 `pinecone.api_key`。

### 安装与升级

上游 README 说明预构建归档、Docker 镜像和 pgrx 构建方式，源码构建使用 cargo-pgrx 0.19.2。0.12.1 是 0.12.0 在新标签下的重新发布，用于修复发布流程；扩展代码和空的 0.11.0 到 0.12.0 升级脚本不变。

从 0.10.0 或更高版本原地升级时，先安装匹配的文件，重启 PostgreSQL 以加载新库，再在每个数据库运行 `ALTER EXTENSION pg_pinecone UPDATE`。用 `pinecone.worker_version()` 检查编译版本与已安装版本一致。0.11.0 之前创建的 native 内积索引仍需 `REINDEX`；更新扩展不会重建这些索引。旧的 `pinegres` 扩展不支持原地升级。

### 预加载与安全

native 与 `pgvector` 不需要预加载。remote worker 需要 `shared_preload_libraries = 'pg_pinecone'` 并重启，应保留已有库。remote 模式会向 Pinecone 发送向量和元数据，应保护 `pinecone.api_key`，评估网络与保留策略，并将仅超级用户可安装的扩展视为可信组件。扩展不可迁移。

### 0.12 的托管文档嵌入

用 `embed` 模型选项跟踪文本列。Pinecone 会在写入和查询时在服务端生成嵌入，无需 pgvector 列或本地 `pinecone.embed` 调用。这会将配置的文本发送到 Pinecone，需要 remote worker 以及通过 `pinecone.set_api_key` 保存的 API 密钥。本地引擎会拒绝 `mode => 'embed'`，因为 PostgreSQL 全文搜索无法运行托管模型。`gemini-embedding-2` 要求 Builder 或更高 Pinecone 套餐，不支持独立 `pinecone.embed` API。

```sql
CREATE TABLE docs (id bigint PRIMARY KEY, body text);
INSERT INTO docs VALUES (1, 'PostgreSQL stores relational data.');
SELECT pinecone.track_docs(
    'docs',
    text_columns => '{"body": {"embed": {"model": "gemini-embedding-2"}}}'
);
SELECT pinecone.flush_docs('docs');
SELECT pinecone.search_docs('docs', 'query text', mode => 'embed');
```

`pinecone.flush_docs` 等待数据投递完成，但远程索引为最终一致，因此立即搜索可能看不到新行。每个索引最多一个文本列可使用 `embed`，且不能同时提供向量列。
