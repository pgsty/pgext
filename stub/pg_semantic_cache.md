## Usage

Sources:

- [Official README](https://github.com/pgEdge/pg_semantic_cache/blob/38e4a910f1ae7e4050abbb8bc1f9cf5d44abf4cc/README.md)
- [Function reference](https://github.com/pgEdge/pg_semantic_cache/blob/38e4a910f1ae7e4050abbb8bc1f9cf5d44abf4cc/docs/functions.md)
- [Version 0.1.0-beta4 extension SQL](https://github.com/pgEdge/pg_semantic_cache/blob/38e4a910f1ae7e4050abbb8bc1f9cf5d44abf4cc/sql/pg_semantic_cache--0.1.0-beta4.sql)
- [Configuration guide](https://github.com/pgEdge/pg_semantic_cache/blob/38e4a910f1ae7e4050abbb8bc1f9cf5d44abf4cc/docs/configuration.md)

`pg_semantic_cache` stores JSON query results with vector embeddings and retrieves them for semantically similar inputs. The application supplies both the embeddings and cached result; the extension does not call an embedding model or execute the stored query text.

### Core Workflow

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

`rebuild_index()` is appropriate here only because the demonstration cache is empty. All embeddings in a cache must match the configured vector dimension. Cache only after producing an authoritative result, and independently decide whether the returned similarity score is safe for that query class.

### Object Index

- `cache_query(...) RETURNS bigint` stores query text, embedding, JSONB result, TTL, and tags.
- `get_cached_result(...)` returns `found`, `result_data`, `similarity_score`, and `age_seconds`.
- `invalidate_cache(pattern, tag)`, `evict_expired`, `evict_lru`, `evict_lfu`, `auto_evict`, and `clear_cache` remove entries.
- `cache_stats`, `cache_hit_rate`, `cache_health`, and `recent_cache_activity` expose usage and health.
- `set_vector_dimension`, `set_index_type`, and `rebuild_index` configure the vector layout and IVFFlat/HNSW index.
- `cache_config` stores keys such as `max_cache_size_mb`, `default_ttl_seconds`, `eviction_policy`, and `similarity_threshold`.

### Operational Notes

Version `0.1.0-beta4` is a preview release. It requires `vector`, uses the fixed `semantic_cache` schema, is non-relocatable, and declares no preload requirement. `rebuild_index()` clears cached data; treat dimension or index changes as destructive maintenance.

Semantic similarity does not prove result equivalence. Include tenant, authorization, model/version, locale, schema, and other result-changing context in application cache design; apply short TTLs or explicit invalidation for mutable data. Cached JSONB can contain sensitive source results, so grant schema/table/function access narrowly. Monitor false hits, stale data, storage growth, eviction, and index recall before production use.
