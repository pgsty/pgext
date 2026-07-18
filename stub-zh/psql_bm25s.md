


## 用法

来源：[README v0.4.13](https://github.com/Intelligent-Internet/psql_bm25s/blob/v0.4.13/README.md), [API reference](https://github.com/Intelligent-Internet/psql_bm25s/blob/v0.4.13/docs/api-reference.md), [query semantics](https://github.com/Intelligent-Internet/psql_bm25s/blob/v0.4.13/docs/query-semantics.md), [input types](https://github.com/Intelligent-Internet/psql_bm25s/blob/v0.4.13/docs/input-types.md), [index parameters](https://github.com/Intelligent-Internet/psql_bm25s/blob/v0.4.13/docs/index-parameters.md), [index policy](https://github.com/Intelligent-Internet/psql_bm25s/blob/v0.4.13/docs/index-policy.md)

`psql_bm25s` 是 PostgreSQL 原生索引访问方法，用于 BM25 系列的词法检索。它通过基于语料库统计信息的排序、精确的 top-k 检索 API，以及适用于可变表的 PostgreSQL 存储和维护行为，保持明确的 BM25 语义。

本目录为 PostgreSQL 17 和 18 打包版本 `0.4.13`。

### 基本搜索

```sql
CREATE EXTENSION psql_bm25s;

CREATE TABLE docs (
    id integer PRIMARY KEY,
    title text NOT NULL,
    body text NOT NULL
);

INSERT INTO docs (id, title, body) VALUES
    (1, 'red apple', 'fresh red apple fruit'),
    (2, 'green apple', 'green apple slices'),
    (3, 'orange citrus', 'orange citrus fruit'),
    (4, 'cat guide', 'small cat animal care');

CREATE INDEX docs_bm25_idx
    ON docs USING psql_bm25s (body);

SELECT d.id, d.title, h.score
FROM psql_bm25s_query(
    'docs_bm25_idx'::regclass,
    'apple fruit',
    5
) AS h
JOIN docs AS d ON d.ctid = h.ctid
ORDER BY h.score DESC, d.id;
```

不带 `WITH (...)` 选项时，索引使用 Lucene 风格的 BM25 和 IDF 默认值，并采用 `realtime` 一致性策略。

### 索引输入

`psql_bm25s` 支持五种可作为索引输入的源列类型：

- `text` 和 `varchar`：直接索引普通标量文本列。
- `text[]` 和 `varchar[]`：用于应用自行维护的词元流。
- `int4[]`：用于应用在外部管理词元 ID 的场景。

标量 `text` 和 `varchar` 列会在索引边界进行分词。预分词数组可以避免对标量重复分词；当应用已经拥有分词结果时，应优先使用这种输入形式。

### 检索 API

规范的精确 BM25 API 是：

- `psql_bm25s_query_tokens(regclass, text[], k, weight_mask)`：用于文本词元索引。
- `psql_bm25s_query_ids(regclass, int4[], k, weight_mask)`：用于词元 ID 索引。

主要的 SQL 便捷 API 是：

- `psql_bm25s_query(regclass, query_text, k, weight_mask, ...)`
- `psql_bm25s_prepared_query(regclass, query_text, ...)`
- `psql_bm25s_query_prepared(prepared_query, k, weight_mask)`

这些行集 API 返回包含 `ctid`、`doc_id` 和 `score` 的 `psql_bm25s_result_hit` 行。当查询同时需要行数据和查询时分数时，可用 `ctid` 将命中结果连接回应用表。

### 操作符

操作符接口适用于 SQL 原生过滤和排序：

- `tokens @@ 'query text'` 是布尔文档匹配谓词。
- `tokens @@@ psql_bm25s_prepared_query(...)` 是与索引绑定的预备谓词。
- `ORDER BY tokens <=> psql_bm25s_order_tokens(...) ASC LIMIT k` 是有序检索形式。

`@@` 不是排序 API。只有当 PostgreSQL 选择真正的 `psql_bm25s` 索引扫描时，`<=>` 才与真实的 BM25 排序一致；需要最明确的精确 top-k 契约时，请使用行集检索 API。

```sql
WITH rq AS (
    SELECT psql_bm25s_ranked_query(
        'docs_bm25_idx'::regclass,
        'apple fruit',
        10
    ) AS q
)
SELECT d.*
FROM docs AS d, rq
WHERE d.body @@@ psql_bm25s_filter_query(rq.q)
ORDER BY d.body <=> psql_bm25s_order_tokens(rq.q) ASC
LIMIT (rq.q).k;
```

### 索引选项

评分相关的索引选项包括：

- `method` 和 `idf_method`，默认 `lucene`；支持 `robertson`、`lucene`、`atire`、`bm25l` 和 `bm25+`。
- `k1`，默认 `1.5`。
- `b`，默认 `0.75`。
- `delta`，默认 `0.5`，由 BM25L 和 BM25+ 使用。

标量文本处理相关的索引选项包括：

- `text_lowercase`，默认 `true`。
- `text_stopwords`，默认 `NULL`。
- `text_stem_english`，默认 `false`。
- `text_fold_diacritics`，默认 `false`。

```sql
CREATE INDEX docs_body_bm25_idx
    ON docs USING psql_bm25s (body)
    WITH (
        method = 'lucene',
        idf_method = 'lucene',
        k1 = 1.5,
        b = 0.75,
        text_stem_english = true
    );
```

### 多字段与混合搜索

对于独立的标题、摘要和正文索引，可使用后期融合辅助函数，让每个字段先生成当前查询范围内的命中结果，再组合分数：

```sql
SELECT d.id, d.title, h.score
FROM psql_bm25s_fusion_query(
    ARRAY[
        'docs_title_bm25_idx'::regclass,
        'docs_body_bm25_idx'::regclass
    ],
    'apple fruit',
    ARRAY[3.0, 1.0]::real[],
    10,
    30,
    NULL
) AS h
JOIN docs AS d ON d.ctid = h.ctid
ORDER BY h.score DESC, d.id;
```

向量/BM25 混合搜索也是一个后期融合层。BM25 和向量索引各自产生候选结果，`psql_bm25s_hybrid_fuse_candidates(...)` 默认使用 `rrf` 将其组合。核心扩展不要求安装 `pgvector`、VectorChord 或任何向量类型。

### 维护与缓存

公开的维护开关是 `consistency` 索引选项：

- `realtime`：默认值，使已提交写入立即可搜索。
- `eventual`：优先降低前台读写延迟，允许 BM25 结果在维护收敛期间短期陈旧。
- `manual`：将刷新交给显式操作或调度器控制。

运维辅助函数包括：

- `psql_bm25s_index_details(regclass)`
- `psql_bm25s_index_policy_recommend(regclass, profile)`
- `psql_bm25s_index_refresh(regclass)`
- `psql_bm25s_index_maintain(regclass)`
- `psql_bm25s_index_try_maintain(regclass)`
- `psql_bm25s_index_maintain_due(max_indexes)`

大型不可变索引载荷可以使用共享代际缓存。零配置 DSM 缓存不要求设置 `shared_preload_libraries`。对于大型连接池部署，可通过 `shared_preload_libraries = 'psql_bm25s'` 和 `psql_bm25s.shared_generation_cache_size` 启用可选的预加载共享内存区，但正常使用不需要该内存区。

相关全局 GUC 包括：

- `psql_bm25s.maintenance_worker_limit`
- `psql_bm25s.preload_timer_interval_ms`
- `psql_bm25s.maintenance_timer_interval_ms`
- `psql_bm25s.maintenance_rebuild_memory_budget`

### 注意事项

- 该扩展创建后不可迁移模式；如有需要，请在 `CREATE EXTENSION` 时选择非 `public` 模式。
- `eventual` 和 `manual` 一致性策略会有意牺牲即时新鲜度，以换取更低的前台成本或显式刷新控制。
- 逻辑复制遵循 PostgreSQL 行为：表行会复制，但索引关系不会作为逻辑数据对象复制，因此应在订阅端创建或重建索引。
- 可选的预加载共享缓存需要修改 PostgreSQL 配置并重启，因为共享内存区在服务器启动时分配。
