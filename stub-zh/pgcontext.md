## 用法

来源：

- [PGXN pgContext 0.1.0 发布页](https://pgxn.org/dist/pgcontext/0.1.0/)
- [官方 pgContext 0.1.0 README](https://github.com/Evokoa/pgContext/blob/v0.1.0/README.md)
- [官方 pgcontext 控制文件](https://github.com/Evokoa/pgContext/blob/v0.1.0/pgcontext.control)
- [官方 PostgreSQL 17 快速入门](https://github.com/Evokoa/pgContext/blob/v0.1.0/docs/user_guide/quickstart.md)
- [官方配置指南](https://github.com/Evokoa/pgContext/blob/v0.1.0/docs/user_guide/configuration.md)
- [官方索引指南](https://github.com/Evokoa/pgContext/blob/v0.1.0/docs/user_guide/indexes.md)
- [官方 0.1.0 已知限制](https://github.com/Evokoa/pgContext/blob/v0.1.0/docs/user_guide/limitations.md)
- [官方 pgvector 共存指南](https://github.com/Evokoa/pgContext/blob/v0.1.0/docs/user_guide/pgvector_coexist.md)

`pgcontext` 0.1.0 是面向 PostgreSQL 17 的扩展，在普通 PostgreSQL 表上提供精确向量检索、注册字段过滤、持久化 HNSW 索引，以及稠密向量与全文搜索混合检索。应用表仍然是数据、MVCC 可见性、ACL/RLS、备份与复制的权威来源；HNSW 只是可重建的派生索引。V1 的稳定能力以基于表的精确检索为中心，`pgcontext_hnsw`、非稠密向量类型和若干高级服务路径仍有明确的实验性边界。

### 核心流程

把扩展文件安装到 PostgreSQL 17 后，创建基于应用表的 collection，并且只注册检索可以使用的字段：

```sql
CREATE EXTENSION pgcontext;

CREATE TABLE public.docs (
    id text PRIMARY KEY,
    embedding vector(3) NOT NULL,
    status text NOT NULL,
    body text NOT NULL,
    metadata jsonb NOT NULL
);

INSERT INTO public.docs VALUES
    ('postgres', '[1,0,0]', 'published', 'postgres vector search', '{"topic":"database"}'),
    ('rust',     '[0,1,0]', 'published', 'rust extension guide',  '{"topic":"systems"}'),
    ('draft',    '[1,1,0]', 'draft',     'internal notes',        '{"topic":"database"}');

SELECT * FROM pgcontext.create_collection('docs', 'public.docs');
SELECT pgcontext.register_vector('docs', 'embedding', 'embedding', 3, 'cosine');
SELECT pgcontext.register_filter_column('docs', 'status', 'status');
SELECT pgcontext.register_jsonb_path('docs', 'topic', 'metadata', ARRAY['topic']);
SELECT pgcontext.upsert_points('docs', ARRAY['postgres', 'rust', 'draft']);

SELECT source_key, score
FROM pgcontext.search(
    'docs',
    '[1,0,0]'::vector,
    '{"must":[{"key":"status","match":"published"}]}'::jsonb,
    10
);
```

过滤器只能使用已注册的列和 JSONB 路径。搜索、计数与分面聚合共享同一套过滤语法：

```sql
SELECT pgcontext.count(
    'docs',
    '{"must":[{"key":"topic","match":"database"}]}'::jsonb
);

SELECT * FROM pgcontext.facet('docs', 'topic', NULL, 10);
```

### HNSW 与混合检索

应先以精确搜索作为工作负载的正确性基准，再添加实验性的 HNSW 访问方法：

```sql
CREATE INDEX docs_embedding_hnsw
ON public.docs USING pgcontext_hnsw (
    embedding pgcontext.vector_hnsw_cosine_ops
);

SELECT source_key, score
FROM pgcontext.search('docs', '[1,0,0]'::vector, 10);

SELECT source_key, score
FROM pgcontext.query(
    'docs',
    '[1,0,0]'::vector,
    'postgres search',
    'body',
    10
);
```

稠密向量 HNSW 操作符类覆盖 L2、余弦、负内积与 L1。把工作负载切换到近似检索路径前，应使用 `pgcontext.index_status`、`pgcontext.index_diagnostics` 与 `pgcontext.recall_check` 校验。`pgcontext.query` 使用倒数排名融合组合稠密向量与 PostgreSQL 全文检索；单向量请求应继续使用 `pgcontext.search`。

### 重要对象

- `vector` 是稳定的稠密向量类型；`halfvec`、`sparsevec` 与 `bitvec` 在 V1 中可见但仍属实验性能力。
- `pgcontext.create_collection`、`collection_info`、`drop_collection` 与 collection alias 管理对应用自有表的注册。
- `register_vector`、`register_filter_column`、`register_jsonb_path` 与 `upsert_points` 建立显式检索契约。
- `search`、`count`、`facet`、`scroll`、`grouped_search`、`recommend` 与 `discover` 提供基于表的检索和导航。
- `query` 与 `explain` 提供混合检索及其 SQL 可见执行阶段。
- `pgcontext_hnsw` 以及 `vector_hnsw_ops`、`vector_hnsw_cosine_ops`、`vector_hnsw_ip_ops`、`vector_hnsw_l1_ops` 操作符类提供稠密向量 ANN 索引。
- `optimization_status`、`index_status`、`index_diagnostics`、`vacuum_advice`、`recall_check` 与 `telemetry` 用于运维检查。

### PostgreSQL 与 pgvector 边界

- V1 只支持 PostgreSQL 17。PostgreSQL 15、16 与 18 是路线图目标，并非 0.1.0 已支持的大版本。
- 控制文件设置 `superuser = false` 与 `relocatable = false`；扩展安装固定的 `pgcontext` 模式和 `$libdir/pgcontext` 动态库，不需要 `shared_preload_libraries`、`LOAD` 或重启服务。
- 把扩展文件安装到主机上的 PostgreSQL 17 目录仍然需要相应文件系统权限。源码构建固定使用 Rust 1.96.0 与 `cargo-pgrx` 0.19.1。
- `pgcontext` 可以提供自己的向量类型。若数据库已经使用 pgvector，应先执行 `CREATE EXTENSION vector`，再执行 `CREATE EXTENSION pgcontext`，使 pgContext 绑定到 pgvector 的既有类型。
- pgContext 不要求外部服务；普通 PostgreSQL 的内存、WAL、认证、备份和 ACL/RLS 配置仍是运维基础。

### V1 注意事项

- `pgcontext_hnsw` 仍属实验性能力，不承诺长期磁盘格式兼容。升级改变页面格式时应重建受影响的 HNSW 索引，并用贴近生产的数据验证召回率、延迟、MVCC、RLS、重启、VACUUM 与副本行为。
- HNSW delta segment 填满后，一次插入可能同步触发压缩，停顿时间会随索引规模增长。可调整 `pgcontext.hnsw_delta_segment_limit` 与 `pgcontext.hnsw_compact_on_threshold_max_mb`，或关闭阈值自动压缩，改为有计划地执行 `pgcontext.compact()` 或 `REINDEX`。
- 稀疏向量及其他非稠密 ANN 覆盖仍不完整。量化辅助函数不代表已经具备量化 HNSW 服务能力，内部维护的晚期交互或内存映射服务也不属于稳定 V1 路径。
- 删除扩展前，应先移除 collection 注册并检查依赖的应用表。`DROP EXTENSION pgcontext` 不会把用户表当作可随意丢弃的加速状态。
