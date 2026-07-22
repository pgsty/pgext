## Usage

Sources:

- [Official control file](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag_jina_reranker_v1_tiny_en/rag_jina_reranker_v1_tiny_en.control)
- [Official pgrag README](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/README.md)
- [Official package manifest](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag_jina_reranker_v1_tiny_en/Cargo.toml)

`rag_jina_reranker_v1_tiny_en` runs the 33-million-parameter jina-reranker-v1-tiny-en model locally and exposes scalar and batch reranking scores or distances. It is deprecated and no longer maintained.

### Core Workflow

The model executes in a background worker, so preload the library and restart PostgreSQL before creating the extension:

```conf
shared_preload_libraries = 'rag_jina_reranker_v1_tiny_en.so'
```

```sql
CREATE EXTENSION vector;
CREATE EXTENSION rag_jina_reranker_v1_tiny_en;

SELECT rag_jina_reranker_v1_tiny_en.rerank_distance(
    'What did the quick brown fox jump over?',
    'The quick brown fox jumps over the lazy dog'
);
```

On macOS, the upstream instructions use `.dylib` instead of `.so` in `shared_preload_libraries`.

### SQL API

- `rag_jina_reranker_v1_tiny_en.rerank_score(text, text)` returns one score; the `text[]` overload returns scores in input order.
- `rag_jina_reranker_v1_tiny_en.rerank_distance(text, text)` returns one distance; the `text[]` overload returns distances in input order.
- For every overload, distance equals negative score. Lower distance therefore corresponds to a higher score.

### Operational Notes

The reviewed version is 0.0.0, non-relocatable, requires `vector`, and sets `superuser = true`. The worker starts with PostgreSQL and lazy-loads the model at first use, so account for first-call latency and memory use. Default builds embed the ONNX model; the optional `remote_onnx` build downloads it after restart and should point to controlled storage. The project pins patched ONNX Runtime components and is archived and deprecated, so deploy a fixed, tested artifact and plan a replacement.
