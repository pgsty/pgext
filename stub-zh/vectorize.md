## 用法

- 来源：[repo README](https://github.com/ChuckHend/pg_vectorize/blob/main/README.md)，[extension README](https://github.com/ChuckHend/pg_vectorize/blob/main/extension/README.md)，[v0.26.1 release](https://github.com/ChuckHend/pg_vectorize/releases/tag/v0.26.1)

`vectorize` 是 `pg_vectorize` 提供的 PostgreSQL 扩展。上游文档描述了两种模式：独立的 HTTP 服务，以及数据库内的 SQL 扩展。对这里打包的扩展而言，相关的是 SQL 工作流。

### 启用扩展

```sql
ALTER SYSTEM SET shared_preload_libraries = 'vectorize,pg_cron';
ALTER SYSTEM SET cron.database_name = 'postgres';

CREATE EXTENSION vectorize CASCADE;
```

extension README 列出了 `pg_cron`、`pgmq` 和 `pgvector` 这些依赖，以及用于嵌入服务的 `vectorize.embedding_service_url`。

### 创建搜索任务

高级 SQL API 从 `vectorize.table()` 开始：

```sql
SELECT vectorize.table(
  job_name    => 'product_search_hf',
  relation    => 'products',
  primary_key => 'product_id',
  columns     => ARRAY['product_name', 'description'],
  transformer => 'sentence-transformers/all-MiniLM-L6-v2',
  schedule    => 'realtime'
);
```

extension README 说明，这会为源表创建并维护一个 embeddings 列。

### 搜索、RAG 与直接模型调用

搜索示例：

```sql
SELECT * FROM vectorize.search(
  job_name       => 'product_search_hf',
  query          => 'accessories for mobile devices',
  return_columns => ARRAY['product_id', 'product_name'],
  num_results    => 3
);
```

上游还记录了：

- `vectorize.rag()`：用于 retrieval-augmented answers。
- `vectorize.generate()`：用于文本生成。
- `vectorize.encode()`：用于直接生成 embedding。
- `vectorize.import_embeddings()`：用于导入预计算向量。

### 更新行为与 v0.26.1 说明

extension README 说明，`schedule => '* * * * *'` 会每分钟检查一次更新，而 `schedule => 'realtime'` 会创建 trigger，在插入和更新时立即刷新。

v0.26.1 release note 只写了 "update dependencies"，因此除了现有 README 中已经记录的接口外，没有额外的上游用户侧 SQL/API 变化需要补充。
