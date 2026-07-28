## Usage

Sources:

- [Official upstream README](https://github.com/activeloopai/deeplake/blob/f432041fabbf4a1fc1d342aeb550e0a7de41b0da/postgres/README.md)
- [Official extension control file (pg_deeplake.control)](https://github.com/activeloopai/deeplake/blob/f432041fabbf4a1fc1d342aeb550e0a7de41b0da/postgres/pg_deeplake.control)
- [Official extension SQL (pg_deeplake--1.0.sql)](https://github.com/activeloopai/deeplake/blob/f432041fabbf4a1fc1d342aeb550e0a7de41b0da/postgres/pg_deeplake--1.0.sql)

`pg_deeplake` — PostgreSQL extension for vector similarity search, full-text search, and hybrid search using DeepLake. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_deeplake;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `array_ndims(anyarray)` is an extension function and returns `int`.
- `contains(col text, search text)` is an extension function and returns `boolean`.
- `create_deeplake_table(tablename TEXT, path TEXT)` is an extension function and returns `void`.
- `deeplake_bm25_cmp(text, text)` is an extension function and returns `int4`.
- `deeplake_bm25_similarity(text, text)` is an extension function and returns `float4`.
- `deeplake_cosine_similarity(float4[], float4[])` is an extension function and returns `float4`.
- `deeplake_hybrid_cmp(deeplake_hybrid_record, deeplake_hybrid_record_weighted)` is an extension function and returns `int4`.
- `deeplake_hybrid_record(embedding float4[], text_value text)` is an extension function and returns `deeplake_hybrid_record_weighted`.
- `deeplake_hybrid_record(embedding float4[], text_value text, embedding_weight float8, text_weight float8)` is an extension function and returns `deeplake_hybrid_record_weighted`.
- `deeplake_hybrid_record_to_weighted(deeplake_hybrid_record)` is an extension function and returns `deeplake_hybrid_record_weighted`.
- `deeplake_hybrid_record_weighted_to_simple(deeplake_hybrid_record_weighted)` is an extension function and returns `deeplake_hybrid_record`.
- `deeplake_hybrid_search(deeplake_hybrid_record, deeplake_hybrid_record_weighted)` is an extension function and returns `float4`.
- `deeplake_index_handler(INTERNAL)` is an extension function and returns `index_am_handler`.
- `deeplake_maxsim(float4[][], float4[][])` is an extension function and returns `float4`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
