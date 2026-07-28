## 用法

来源：

- [pgContext 0.2.0 README](https://github.com/evokoa/pgcontext/blob/v0.2.0/README.md)
- [pgContext 0.2.0 发行说明](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/release_notes.md)
- [pgContext collection 快速入门](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/quickstart.md)
- [pgContext 索引指南](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/indexes.md)
- [pgContext 已知限制](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/limitations.md)
- [pgContext 控制文件](https://github.com/evokoa/pgcontext/blob/v0.2.0/pgcontext.control)
- [pgvector 共存指南](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/pgvector_coexist.md)

`pgcontext` 把向量与混合检索保留在 PostgreSQL 内。它提供 pgContext 自有向量类型、基于应用表的 collection 元数据、已注册字段过滤、精确搜索、持久化 HNSW，以及稠密向量与全文搜索融合。应用行仍是 MVCC、ACL/RLS、备份与复制的权威来源；索引和生成产物只是可重建的加速状态。

0.2.0 面向 PostgreSQL 17 与 18，而该版本的受控试点认证重点是 PostgreSQL 17。高级 HNSW、非稠密、量化、映射与晚期交互路径仍有明确的实验性边界。

### 核心流程

```sql
CREATE EXTENSION pgcontext;

CREATE TABLE public.docs (
    id text PRIMARY KEY,
    embedding pgcontext.vector(2) NOT NULL,
    status text NOT NULL,
    body text NOT NULL,
    metadata jsonb NOT NULL
);

INSERT INTO public.docs (id, embedding, status, body, metadata) VALUES
    ('doc-1', '[1,0]'::pgcontext.vector, 'published', 'postgres vector search', '{"topic":"postgres"}'),
    ('doc-2', '[0,1]'::pgcontext.vector, 'published', 'rust extension guide', '{"topic":"rust"}');

SELECT * FROM pgcontext.create_collection('docs', 'public.docs');
SELECT pgcontext.register_vector('docs', 'embedding', 'embedding', 2, 'l2');
SELECT pgcontext.register_filter_column('docs', 'status', 'status');
SELECT pgcontext.register_jsonb_path('docs', 'topic', 'metadata', ARRAY['topic']);
SELECT pgcontext.upsert_points('docs', ARRAY['doc-1', 'doc-2']);

SELECT source_key, score
FROM pgcontext.search(
    'docs',
    '[1,0]'::pgcontext.vector,
    '{"must":[{"key":"status","match":"published"}]}'::jsonb,
    10
);
```

Collection 用于描述应用自有表，不会把这些行复制到另一套权威存储。搜索、计数、分面、分组、滚动、推荐与发现共享已注册的向量和过滤定义。

### HNSW 与混合检索

```sql
SET maintenance_work_mem = '2GB';
CREATE INDEX docs_embedding_hnsw ON public.docs
    USING pgcontext_hnsw
    (embedding pgcontext.vector_hnsw_cosine_ops);
RESET maintenance_work_mem;

SELECT source_key, score
FROM pgcontext.query(
    'docs',
    '[1,0]'::pgcontext.vector,
    'postgres search',
    'body',
    10
);
```

稠密 HNSW 操作符类覆盖 L2、内积、余弦与 L1。索引构建会强制执行 `maintenance_work_mem`；应先确定构建预算，再用精确搜索与 `pgcontext.recall_check` 比较近似结果。`pgcontext.query` 使用倒数排名融合组合稠密向量与 PostgreSQL 全文分支。

### 重要对象

- `pgcontext.vector`、`pgcontext.halfvec`、`pgcontext.sparsevec` 与 `pgcontext.bitvec` 是扩展自有类型；非稠密变体仍属实验性能力。
- `pgcontext.create_collection`、注册函数与点映射函数定义基于源表的检索契约。
- `pgcontext.search`、`count`、`facet`、`scroll`、`grouped_search`、`recommend` 与 `discover` 提供基于表的检索。
- `pgcontext.query` 与 `explain` 提供可组合及混合检索。
- `pgcontext_hnsw` 及按度量区分的操作符类提供 ANN 索引。
- 索引状态、诊断、VACUUM 建议、召回检查、优化状态与有界遥测用于运维复核。

### 升级与 pgvector 边界

0.2.0 把 pgContext 自有类型移动到固定的 `pgcontext` schema。已有独立 0.1.0 安装可以由超级用户执行软件包提供的升级：

```sql
ALTER EXTENSION pgcontext UPDATE TO '0.2.0';
```

升级后，应限定类型名，例如 `pgcontext.vector(1536)`，或有意把该 schema 加入 `search_path`。如果 0.1.0 数据库中的 public 向量类型属于 pgvector，升级会在变更前拒绝执行；应盘点依赖、安装 0.2.0 与独立的 `pgcontext_pgvector` 桥、重建注册和依赖对象，并在不改变 pgvector 列的前提下重建 pgContext 索引。

主扩展不依赖 pgvector。其类型与 `public.vector`、`public.halfvec` 和 `public.sparsevec` 不同；不要假定安装顺序会让一个扩展的类型成为另一个扩展的别名。

### 运维边界

- pgContext 安装访问方法，因此 `CREATE EXTENSION` 与升级需要 PostgreSQL 超级用户；已授权的应用 API 不需要超级用户。
- 主扩展不要求 `shared_preload_libraries`、`LOAD` 或重启。
- 早期版本 HNSW 磁盘格式不稳定；跨版本时应规划并验证索引重建，而不能把索引文件当作可移植数据。
- 精确重排、MVCC、ACL 与 RLS 检查仍是正确性边界，但不能替代针对工作负载的召回率、延迟、重启、VACUUM、复制与故障测试。
- 删除扩展前应移除 collection 注册并检查应用依赖对象；避免未经审查的 `CASCADE`。
