## 用法

来源：

- [官方上游 README](https://github.com/pg-cuvs/pg_cuvs/blob/f8bfcc9a0d594c28257870ac23f9f130b90130e7/README.md)
- [官方扩展控制文件 (pg_cuvs.control)](https://github.com/pg-cuvs/pg_cuvs/blob/f8bfcc9a0d594c28257870ac23f9f130b90130e7/pg_cuvs.control)
- [官方扩展 SQL (pg_cuvs--0.4.0--0.5.0.sql)](https://github.com/pg-cuvs/pg_cuvs/blob/f8bfcc9a0d594c28257870ac23f9f130b90130e7/sql/pg_cuvs--0.4.0--0.5.0.sql)

`pg_cuvs` — 通过 NVIDIA cuVS 在 PostgreSQL 上实现 GPU 加速向量搜索的扩展，这是一种异构加速路径，保持 PostgreSQL 作为控制平面。使用它来进行相应的向量、模型或检索工作流。经过审核的上游材料已将此功能标记为弃用。

### 核心工作流

```sql
CREATE EXTENSION pg_cuvs;
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `cuvs_filtered_knn(index_rel regclass, query vector, filter_tids bigint[], k integer)` 是一个扩展函数，返回 `TABLE`。
- `cuvs_filtered_knn(index_rel regclass, query vector, filter_tids tid[], k integer)` 是一个扩展函数，返回 `TABLE`。
- `cuvsamhandler(internal)` 是一个扩展函数，返回 `index_am_handler`。
- `flatamhandler(internal)` 是一个扩展函数，返回 `index_am_handler`。
- `ivfpqamhandler(internal)` 是一个扩展函数，返回 `index_am_handler`。
- `pg_cuvs_batch_search(rel regclass, queries vector[], k integer, OUT query_idx integer, OUT ctid tid, OUT distance real)` 是一个扩展函数，返回 `SETOF`。
- `pg_cuvs_build_hnsw(cagra_oid regclass, mode text DEFAULT 'nsw')` 是一个扩展函数，返回 `regclass`。
- `pg_cuvs_compact(index_rel regclass)` 是一个扩展函数，返回 `void`。
- `pg_cuvs_eat_vram(leave_bytes bigint)` 是一个扩展函数，返回 `void`。
- `pg_cuvs_free_vram()` 是一个扩展函数，返回 `void`。
- `pg_cuvs_gc_orphans(do_delete boolean DEFAULT false, OUT db_oid oid, OUT index_oid oid, OUT reason text, OUT action text)` 是一个扩展函数，返回 `SETOF`。
- `pg_cuvs_gpu_cache_stats(OUT gpu_device_id integer, OUT hits bigint, OUT misses bigint, OUT evictions bigint, OUT reloads bigint, OUT persist_failures bigint, OUT resident_count integer, OUT vram_used_mb bigint, OUT vram_budget_mb bigint, OUT bf_vram_mb bigint, OUT bf_precision text)` 是一个扩展函数，返回 `SETOF`。
- `pg_cuvs_gpu_fallback_stats(OUT index_oid regclass, OUT fallback_count bigint, OUT last_reason text, OUT last_fallback_at timestamptz)` 是一个扩展函数，返回 `SETOF`。
- `pg_cuvs_gpu_shard_stats(OUT database_oid oid, OUT index_oid oid, OUT index_name text, OUT shard_id integer, OUT gpu_device_id integer, OUT n_vecs bigint, OUT tid_offset bigint, OUT vram_used_mb bigint, OUT search_count bigint, OUT error_count bigint, OUT resident boolean, OUT last_st…)` 是一个扩展函数，返回 `SETOF`。

### 要求与注意事项

- 经过审核的控制文件声明默认版本为 `0.5.0`。
- 先安装确认的扩展依赖项：`vector`。
- 控制文件将该扩展标记为不可重定位。
- 上游材料包含显式的弃用边界。
- 上游材料表明该项目已被放弃或不再维护。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
