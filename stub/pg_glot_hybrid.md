## Usage

Sources:

- [extensions/pg_glot_hybrid/pg_glot_hybrid.control](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/extensions/pg_glot_hybrid/pg_glot_hybrid.control)
- [README.md](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/README.md)
- [Cargo.toml](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/Cargo.toml)
- [extensions/pg_glot_hybrid/Cargo.toml](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/extensions/pg_glot_hybrid/Cargo.toml)
- [extensions/pg_glot_hybrid/src/lib.rs](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/extensions/pg_glot_hybrid/src/lib.rs)
- [extensions/pg_glot_hybrid/src/customscan.rs](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/extensions/pg_glot_hybrid/src/customscan.rs)

`pg_glot_hybrid` combines the CJK configurations of `pg_glot`, BM25 from `pg_textsearch` and dense vectors from `vector`. It offers `glot.rank` planner integration and the explicit `glot.hybrid` set-returning interface.

### Core Workflow

```conf
shared_preload_libraries = 'pg_textsearch,pg_glot_hybrid'
```

```sql
CREATE EXTENSION pg_glot_hybrid CASCADE;
CREATE TABLE glot_documents (id bigint PRIMARY KEY, body text, emb vector(3));
CREATE INDEX ON glot_documents USING bm25(body) WITH (text_config = 'public.korean');
CREATE INDEX ON glot_documents USING hnsw(emb vector_cosine_ops);
SELECT id, score FROM glot.hybrid('glot_documents', 'id', 'body', 'emb',
  '형태소 분석', '[1,0,0]'::vector, 60, 60, 10);
```

### Operational Boundaries

Install as a superuser. Preload both libraries shown below and restart for the intended planner path. Literal queries with plain ORDER BY/LIMIT and real column references are required for the custom scan. Without the hook, `glot.rank` falls back to a different, non-RRF score. The explicit `glot.hybrid` interface does not need that hook but requires a matching BM25 index, a vector index and a bigint key. Embeddings come from the application. The reviewed default target is PostgreSQL 17; inherited extension dependencies impose their own version constraints.
