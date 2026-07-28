## Usage

Sources:

- [Official upstream README](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/README.md)
- [Official extension control file (pgturbohybrid.control)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/pgturbohybrid.control)
- [Official extension SQL (pgturbohybrid--0.1.2--0.2.0.sql)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/sql/pgturbohybrid--0.1.2--0.2.0.sql)

`pgturbohybrid` — This README helps you understand what pgturbohybrid does, when hybrid search is useful, how to install it, how to create your first index, and how to check whether the fast path is working. Use it for the corresponding vector, model, or retrieval workflow. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION pgturbohybrid;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `turbohybrid_cosine_distance(vector, turbohybrid_query)` is an extension function and returns `pg_catalog`.
- `turbohybrid_dense_query(vector_query vector, final_k pg_catalog.int4 DEFAULT NULL, dense_k pg_catalog.int4 DEFAULT NULL)` is an extension function and returns `turbohybrid_query`.
- `turbohybrid_distance(vector, turbohybrid_query)` is an extension function and returns `pg_catalog`.
- `turbohybrid_estimate_memory(pg_catalog.regclass)` is an extension function and returns `pg_catalog`.
- `turbohybrid_graph_repair_dry_run(index pg_catalog.regclass, sample_nodes pg_catalog.int4 DEFAULT 1000, search_ef pg_catalog.int4 DEFAULT 400, candidate_limit pg_catalog.int4 DEFAULT 200)` is an extension function and returns `pg_catalog`.
- `turbohybrid_handler(pg_catalog.internal)` is an extension function and returns `pg_catalog`.
- `turbohybrid_hybrid_query(vector_query vector, text_query pg_catalog.tsquery, final_k pg_catalog.int4 DEFAULT NULL, dense_k pg_catalog.int4 DEFAULT NULL, bm25_k pg_catalog.int4 DEFAULT NULL)` is an extension function and returns `turbohybrid_query`.
- `turbohybrid_index_stats(pg_catalog.regclass)` is an extension function and returns `pg_catalog`.
- `turbohybrid_l2_distance(vector, turbohybrid_query)` is an extension function and returns `pg_catalog`.
- `turbohybrid_last_build_stats()` is an extension function and returns `pg_catalog`.
- `turbohybrid_last_scan_diagnosis()` is an extension function and returns `pg_catalog`.
- `turbohybrid_last_scan_stats()` is an extension function and returns `pg_catalog`.
- `turbohybrid_negative_inner_product(vector, turbohybrid_query)` is an extension function and returns `pg_catalog`.
- `turbohybrid_prewarm(pg_catalog.regclass)` is an extension function and returns `pg_catalog`.

### Requirements and Caveats

- The reviewed control file declares default version `0.2.0`.
- Install the confirmed extension dependencies first: `vector`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
