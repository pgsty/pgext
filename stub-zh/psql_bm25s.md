


## 用法

来源：[README v0.4.13](https://github.com/Intelligent-Internet/psql_bm25s/blob/v0.4.13/README.md), [API reference](https://github.com/Intelligent-Internet/psql_bm25s/blob/v0.4.13/docs/api-reference.md), [query semantics](https://github.com/Intelligent-Internet/psql_bm25s/blob/v0.4.13/docs/query-semantics.md), [input types](https://github.com/Intelligent-Internet/psql_bm25s/blob/v0.4.13/docs/input-types.md), [index parameters](https://github.com/Intelligent-Internet/psql_bm25s/blob/v0.4.13/docs/index-parameters.md), [index policy](https://github.com/Intelligent-Internet/psql_bm25s/blob/v0.4.13/docs/index-policy.md)

`psql_bm25s` 是 PostgreSQL-native index access method，用于 BM25-family lexical retrieval。它通过基于 corpus statistics 的 ranking、精确 top-k retrieval APIs，以及面向可变表的 PostgreSQL storage/maintenance 行为，保持 BM25 语义明确。

本 catalog 为 PostgreSQL 17 和 18 打包版本 `0.4.13`。

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

不带 `WITH (...)` 选项时，索引使用 Lucene-style BM25 和 IDF defaults，并使用 `realtime` consistency policy。

### Indexed Inputs

`psql_bm25s` 支持五种被索引源列类型：

- `text` 和 `varchar`：直接索引普通标量文本列。
- `text[]` 和 `varchar[]`：用于应用自行维护的 token streams。
- `int4[]`：用于应用在外部管理 token IDs 的场景。

标量 `text` 和 `varchar` 列会在索引边界 tokenized。Pretokenized arrays 可避免标量重复 tokenization，当应用已经拥有 tokenization 结果时是首选形状。

### Retrieval APIs

规范的精确 BM25 API 是：

- `psql_bm25s_query_tokens(regclass, text[], k, weight_mask)`：用于 token-text indexes。
- `psql_bm25s_query_ids(regclass, int4[], k, weight_mask)`：用于 token-ID indexes。

主要 SQL 便利 API 是：

- `psql_bm25s_query(regclass, query_text, k, weight_mask, ...)`
- `psql_bm25s_prepared_query(regclass, query_text, ...)`
- `psql_bm25s_query_prepared(prepared_query, k, weight_mask)`

这些 rowset APIs 返回包含 `ctid`、`doc_id` 和 `score` 的 `psql_bm25s_result_hit` 行。当查询同时需要行数据和 query-time score 时，用 `ctid` 将 hits 连接回应用表。

### 操作符

操作符表面适用于 SQL-native filtering 和 ordering：

- `tokens @@ 'query text'` 是布尔 document-match predicate。
- `tokens @@@ psql_bm25s_prepared_query(...)` 是 index-bound prepared predicate。
- `ORDER BY tokens <=> psql_bm25s_order_tokens(...) ASC LIMIT k` 是 ordered retrieval 形式。

`@@` 不是 ranking API。只有当 PostgreSQL 选择真正的 `psql_bm25s` index scan 时，`<=>` 才与真实 BM25 ordering 对齐；需要最清晰的精确 top-k 契约时，请使用 rowset retrieval APIs。

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

Scoring reloptions 包括：

- `method` 和 `idf_method`，默认 `lucene`；支持的 variants 是 `robertson`、`lucene`、`atire`、`bm25l` 和 `bm25+`。
- `k1`，默认 `1.5`。
- `b`，默认 `0.75`。
- `delta`，默认 `0.5`，由 BM25L 和 BM25+ 使用。

标量文本处理 reloptions 包括：

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

### Multi-Field 与 Hybrid Search

对于独立的 title、abstract 和 body indexes，使用 late fusion helpers，让每个字段先产生 query-scoped hits，然后再组合分数：

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

Hybrid vector/BM25 search 也是 late-fusion 层。BM25 和 vector indexes 各自产生 candidates，`psql_bm25s_hybrid_fuse_candidates(...)` 默认使用 `rrf` 将它们组合。核心扩展不要求安装 `pgvector`、VectorChord 或任何 vector type。

### 维护与缓存

公开维护开关是 `consistency` reloption：

- `realtime`：默认值，使已提交写入立即可搜索。
- `eventual`：偏向前台读写延迟，允许维护收敛期间 BM25 结果短期陈旧。
- `manual`：将刷新留给显式 operator 或 scheduler 控制。

运维辅助函数包括：

- `psql_bm25s_index_details(regclass)`
- `psql_bm25s_index_policy_recommend(regclass, profile)`
- `psql_bm25s_index_refresh(regclass)`
- `psql_bm25s_index_maintain(regclass)`
- `psql_bm25s_index_try_maintain(regclass)`
- `psql_bm25s_index_maintain_due(max_indexes)`

大型 immutable index payloads 可使用 shared generation cache。零配置 DSM cache 不要求 `shared_preload_libraries`。对于大型 connection-pool 部署，可通过 `shared_preload_libraries = 'psql_bm25s'` 和 `psql_bm25s.shared_generation_cache_size` 启用可选 shared-preload arena，但正常使用不需要该 arena。

相关全局 GUC 包括：

- `psql_bm25s.maintenance_worker_limit`
- `psql_bm25s.preload_timer_interval_ms`
- `psql_bm25s.maintenance_timer_interval_ms`
- `psql_bm25s.maintenance_rebuild_memory_budget`

### 注意事项

- 该扩展创建后不可 relocatable；如有需要，请在 `CREATE EXTENSION` 时选择非 `public` schema。
- `eventual` 和 `manual` consistency 会有意用 immediate freshness 换取更低前台成本或显式刷新控制。
- Logical replication 遵循 PostgreSQL 行为：表行会复制，但 index relations 不会作为 logical data objects 复制，因此 subscribers 上应创建或重建 indexes。
- 可选 shared-preload cache 需要 PostgreSQL 配置并重启，因为 shared arena 在服务器启动时分配。
