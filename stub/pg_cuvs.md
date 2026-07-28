## Usage

Sources:

- [Official upstream README](https://github.com/pg-cuvs/pg_cuvs/blob/f8bfcc9a0d594c28257870ac23f9f130b90130e7/README.md)
- [Official extension control file (pg_cuvs.control)](https://github.com/pg-cuvs/pg_cuvs/blob/f8bfcc9a0d594c28257870ac23f9f130b90130e7/pg_cuvs.control)
- [Official extension SQL (pg_cuvs--0.4.0--0.5.0.sql)](https://github.com/pg-cuvs/pg_cuvs/blob/f8bfcc9a0d594c28257870ac23f9f130b90130e7/sql/pg_cuvs--0.4.0--0.5.0.sql)

`pg_cuvs` — GPU-accelerated vector search for PostgreSQL via NVIDIA cuVS — a heterogeneous acceleration path that keeps Postgres as the control plane. Use it for the corresponding vector, model, or retrieval workflow. The reviewed upstream material marks this capability deprecated.

### Core Workflow

```sql
CREATE EXTENSION pg_cuvs;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `cuvs_filtered_knn(index_rel regclass, query vector, filter_tids bigint[], k integer)` is an extension function and returns `TABLE`.
- `cuvs_filtered_knn(index_rel regclass, query vector, filter_tids tid[], k integer)` is an extension function and returns `TABLE`.
- `cuvsamhandler(internal)` is an extension function and returns `index_am_handler`.
- `flatamhandler(internal)` is an extension function and returns `index_am_handler`.
- `ivfpqamhandler(internal)` is an extension function and returns `index_am_handler`.
- `pg_cuvs_batch_search(rel regclass, queries vector[], k integer, OUT query_idx integer, OUT ctid tid, OUT distance real)` is an extension function and returns `SETOF`.
- `pg_cuvs_build_hnsw(cagra_oid regclass, mode text DEFAULT 'nsw')` is an extension function and returns `regclass`.
- `pg_cuvs_compact(index_rel regclass)` is an extension function and returns `void`.
- `pg_cuvs_eat_vram(leave_bytes bigint)` is an extension function and returns `void`.
- `pg_cuvs_free_vram()` is an extension function and returns `void`.
- `pg_cuvs_gc_orphans(do_delete boolean DEFAULT false, OUT db_oid oid, OUT index_oid oid, OUT reason text, OUT action text)` is an extension function and returns `SETOF`.
- `pg_cuvs_gpu_cache_stats(OUT gpu_device_id integer, OUT hits bigint, OUT misses bigint, OUT evictions bigint, OUT reloads bigint, OUT persist_failures bigint, OUT resident_count integer, OUT vram_used_mb bigint, OUT vram_budget_mb bigint, OUT bf_vram_mb bigint, OUT bf_precision text)` is an extension function and returns `SETOF`.
- `pg_cuvs_gpu_fallback_stats(OUT index_oid regclass, OUT fallback_count bigint, OUT last_reason text, OUT last_fallback_at timestamptz)` is an extension function and returns `SETOF`.
- `pg_cuvs_gpu_shard_stats(OUT database_oid oid, OUT index_oid oid, OUT index_name text, OUT shard_id integer, OUT gpu_device_id integer, OUT n_vecs bigint, OUT tid_offset bigint, OUT vram_used_mb bigint, OUT search_count bigint, OUT error_count bigint, OUT resident boolean, OUT last_st…)` is an extension function and returns `SETOF`.

### Requirements and Caveats

- The reviewed control file declares default version `0.5.0`.
- Install the confirmed extension dependencies first: `vector`.
- The control file marks the extension as non-relocatable.
- Upstream material contains an explicit deprecation boundary.
- Upstream material indicates that the project is abandoned or unmaintained.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
