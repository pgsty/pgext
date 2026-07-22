## Usage

Sources:

- [Official control file](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag_bge_small_en_v15/rag_bge_small_en_v15.control)
- [Official pgrag README](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/README.md)
- [Official package manifest](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag_bge_small_en_v15/Cargo.toml)

`rag_bge_small_en_v15` runs the 33-million-parameter bge-small-en-v1.5 model locally to tokenize text, split model-aware chunks, and generate 384-dimensional passage or query embeddings. It is deprecated and no longer maintained.

### Core Workflow

The model executes in a background worker, so preload the library and restart PostgreSQL before creating the extension:

```conf
shared_preload_libraries = 'rag_bge_small_en_v15.so'
```

```sql
CREATE EXTENSION vector;
CREATE EXTENSION rag_bge_small_en_v15;

SELECT rag_bge_small_en_v15.embedding_for_query(
    'What did the quick brown fox jump over?'
);
```

On macOS, the upstream instructions use `.dylib` instead of `.so` in `shared_preload_libraries`.

### SQL API

- `rag_bge_small_en_v15.chunks_by_token_count(text, integer, integer)` returns chunks bounded by model tokens with the requested overlap.
- `rag_bge_small_en_v15.embedding_for_passage(text)` returns `vector(384)` for stored passages.
- `rag_bge_small_en_v15.embedding_for_query(text)` returns `vector(384)` for search queries.

### Operational Notes

The reviewed version is 0.0.0, non-relocatable, requires `vector`, and sets `superuser = true`. The worker starts with PostgreSQL but lazy-loads the model on first use; expect first-call latency and substantial memory use. The default build embeds the large ONNX model. A `remote_onnx` build downloads it after restart and should use a controlled URL. The repository pins patched ONNX Runtime dependencies and documents possible platform-specific symbol failures. Pin the entire build because the project is archived and deprecated.
