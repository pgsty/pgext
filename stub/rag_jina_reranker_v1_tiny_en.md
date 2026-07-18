## Usage

Sources:

- [Official extension control file](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag_jina_reranker_v1_tiny_en/rag_jina_reranker_v1_tiny_en.control)
- [Official upstream documentation](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/README.md)
- [Official Rust package manifest](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag_jina_reranker_v1_tiny_en/Cargo.toml)

`rag_jina_reranker_v1_tiny_en` — Runs jina-reranker-v1-tiny-en locally in a preloaded background worker and exposes SQL rerank scores and distances.

The reviewed catalog snapshot records version `0.0.0`, kind `preload`, and implementation language `Rust`.
Install and validate the declared extension dependencies first: `vector`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "rag_jina_reranker_v1_tiny_en";
SELECT extversion
FROM pg_extension
WHERE extname = 'rag_jina_reranker_v1_tiny_en';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
