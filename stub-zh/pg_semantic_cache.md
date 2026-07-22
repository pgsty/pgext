## 用法

来源：

- [官方 README](https://github.com/pgEdge/pg_semantic_cache/blob/38e4a910f1ae7e4050abbb8bc1f9cf5d44abf4cc/README.md)
- [函数参考](https://github.com/pgEdge/pg_semantic_cache/blob/38e4a910f1ae7e4050abbb8bc1f9cf5d44abf4cc/docs/functions.md)
- [0.1.0-beta4 版本扩展 SQL](https://github.com/pgEdge/pg_semantic_cache/blob/38e4a910f1ae7e4050abbb8bc1f9cf5d44abf4cc/sql/pg_semantic_cache--0.1.0-beta4.sql)
- [配置指南](https://github.com/pgEdge/pg_semantic_cache/blob/38e4a910f1ae7e4050abbb8bc1f9cf5d44abf4cc/docs/configuration.md)

`pg_semantic_cache` 使用向量嵌入存储 JSON 查询结果，并为语义相似输入检索结果。嵌入和缓存结果都由应用提供；扩展不会调用嵌入模型，也不会执行保存的查询文本。

### 核心流程

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_semantic_cache;
SELECT semantic_cache.init_schema();

-- Configure a three-dimensional demonstration cache.
SELECT semantic_cache.set_vector_dimension(3);
SELECT semantic_cache.rebuild_index();

SELECT semantic_cache.cache_query(
    query_text  := 'completed orders summary',
    embedding   := '[0.10, 0.20, 0.30]'::text,
    result_data := '{"total": 150}'::jsonb,
    ttl_seconds := 3600,
    tags        := ARRAY['orders', 'analytics']
);

SELECT *
FROM semantic_cache.get_cached_result(
    embedding            := '[0.11, 0.19, 0.31]'::text,
    similarity_threshold := 0.95,
    max_age_seconds      := NULL
);
```

这里只因为演示缓存为空，才适合调用 `rebuild_index()`。同一缓存中的所有嵌入必须符合已配置的向量维度。只有生成权威结果后才应缓存，并应针对该查询类别独立判断返回的相似度是否安全。

### 对象索引

- `cache_query(...) RETURNS bigint` 存储查询文本、嵌入、JSONB 结果、TTL 和标签。
- `get_cached_result(...)` 返回 `found`、`result_data`、`similarity_score` 和 `age_seconds`。
- `invalidate_cache(pattern, tag)`、`evict_expired`、`evict_lru`、`evict_lfu`、`auto_evict` 和 `clear_cache` 删除条目。
- `cache_stats`、`cache_hit_rate`、`cache_health` 和 `recent_cache_activity` 暴露使用情况与健康状态。
- `set_vector_dimension`、`set_index_type` 和 `rebuild_index` 配置向量布局与 IVFFlat/HNSW 索引。
- `cache_config` 保存 `max_cache_size_mb`、`default_ttl_seconds`、`eviction_policy` 和 `similarity_threshold` 等键。

### 运维说明

版本 `0.1.0-beta4` 是预览版。它依赖 `vector`，使用固定的 `semantic_cache` 模式，不可重定位，且未声明预加载要求。`rebuild_index()` 会清空缓存数据；维度或索引变更应按破坏性维护处理。

语义相似并不证明结果等价。应用缓存设计必须包含租户、授权、模型与版本、区域设置、模式以及其他会改变结果的上下文；对于可变数据应使用短 TTL 或显式失效。缓存的 JSONB 可能包含敏感源结果，因此应严格授予模式、表和函数权限。生产使用前，应监控误命中、陈旧数据、存储增长、淘汰和索引召回率。
