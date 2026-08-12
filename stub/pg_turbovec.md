## Usage

Sources:

- [pg_turbovec v1.29.0 README](https://codeberg.org/gregburd/pg_turbovec/src/tag/v1.29.0/README.md)
- [pg_turbovec v1.29.0 changelog](https://codeberg.org/gregburd/pg_turbovec/src/tag/v1.29.0/CHANGELOG.md)
- [pg_turbovec v1.29.0 control file](https://codeberg.org/gregburd/pg_turbovec/src/tag/v1.29.0/pg_turbovec.control)
- [Partitioned-scale guide](https://codeberg.org/gregburd/pg_turbovec/src/tag/v1.29.0/docs/PARTITIONED_SCALE.md)
- [Filtering guide](https://codeberg.org/gregburd/pg_turbovec/src/tag/v1.29.0/docs/FILTERING.md)
- [Pigsty package matrix](https://pgext.cloud/ext/pg_turbovec)

`pg_turbovec` 1.29.0 provides a dense `turbovec.vector` type and a `turbovec` nearest-neighbor index access method. It quantizes floating-point coordinates to 2, 3, or 4 bits and reranks candidates against heap vectors. Use it for storage-constrained cosine or inner-product search; choose the index kind deliberately because the default flat scan is linear in row count.

### Create and Query Vectors

```sql
CREATE EXTENSION pg_turbovec;
SET search_path = public, turbovec;

CREATE TABLE items (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  embedding turbovec.vector
);

INSERT INTO items (embedding)
VALUES ('[1,2,3]'), ('[4,5,6]');

SELECT id, embedding <=> '[3,1,2]'::turbovec.vector AS cosine_distance
FROM items
ORDER BY embedding <=> '[3,1,2]'::turbovec.vector
LIMIT 10;
```

Distance operators are `<->` for L2, `<#>` for negative inner product, `<=>` for cosine distance, and `<+>` for L1. The current index supports inner-product and cosine ordering; L2 and L1 are exact-only operations.

The `turbovec.vector` type accepts 1–16,000 coordinates. Indexed vectors must have a fixed dimension that is a multiple of 8; use a check constraint or application validation when the column itself is variable-dimensional.

### Choose and Build an Index Kind

```sql
-- Default flat quantized scan
CREATE INDEX items_embedding_flat_idx ON items
USING turbovec (embedding vec_cosine_ops)
WITH (bit_width = 4);

-- Out-of-core IVF alternative
CREATE INDEX items_embedding_ivf_idx ON items
USING turbovec (embedding vec_cosine_ops)
WITH (bit_width = 4, lists = 1024);

-- Navigable-graph alternative
CREATE INDEX items_embedding_graph_idx ON items
USING turbovec (embedding vec_cosine_ops)
WITH (bit_width = 4, graph = true);

SET turbovec.probes = 32;

SELECT id
FROM items
ORDER BY embedding <=> '[3,1,2]'::turbovec.vector
LIMIT 10;
```

These `CREATE INDEX` statements are alternatives, not a recommendation to keep all three. The default flat kind performs an `O(n * dim)` quantized scan and can reach exact recall after heap reranking, but it is a poor latency choice at large row counts. `WITH (lists = N)` enables an out-of-core IVF layer; `WITH (graph = true)` enables the Vamana graph for lower-latency ANN at moderate scale.

Use `vec_cosine_ops` with `<=>` or `vec_ip_ops` with `<#>`. `bit_width = 4` is the default and generally favors recall; 2-bit indexes are smaller but need workload-specific recall testing. Three-bit indexes are also supported. `CREATE INDEX CONCURRENTLY` is supported.

Important tuning controls include `turbovec.probes`, `turbovec.search_k`, `turbovec.oversample`, `turbovec.hi_dim_rerank`, `turbovec.iterative_scan`, and `turbovec.cache_size_mb`. Change one dimension at a time and compare approximate results with an exact baseline.

### Filtering and Partitioning

Use PostgreSQL partial indexes for stable filter values, the documented `turbovec.knn(..., allowed)` surface for an explicit candidate allowlist, or iterative scan for normal filtered `ORDER BY ... LIMIT` queries.

Version 1.29 documents native PostgreSQL partitioning for larger-than-single-table datasets. A parent query can use `Merge Append` across per-partition TurboVec indexes:

```sql
SELECT id
FROM partitioned_items
ORDER BY embedding <=> $1::turbovec.vector
LIMIT 20;
```

Build, vacuum, and reindex each partition independently. Partition pruning based on a coarse vector quantizer is only a design in 1.29.0, not a shipped feature.

### Version and Integrity Boundaries

- The control file installs objects in schema `turbovec`, is not relocatable, and does not require `shared_preload_libraries` or a server restart.
- Upstream v1.29 targets PostgreSQL 13-18 and labels PostgreSQL 19 support experimental; current Pigsty 1.29.0 packages cover PostgreSQL 14-18 and provide the matching OpenBLAS-linked binary.
- Upstream 1.28.4 fixes persisted row-count drift that could corrupt the index ID table and adds `turbovec.turbovec_check(regclass)`. An already corrupt index still needs `REINDEX` or drop/recreate recovery.
- Version 1.29.0 is additive, keeps wire format 7, and does not require reindexing when upgrading from a healthy 1.28.4 index. `ALTER EXTENSION pg_turbovec UPDATE TO '1.29.0'` is sufficient after the new files are installed.
- Although the 1.29 reloption parser accepts `bit_width = 1`, end-to-end one-bit indexing is not implemented and `CREATE INDEX` intentionally errors. Use `bit_width = 2`, `bit_width = 3`, or `bit_width = 4`.
- The on-disk ID table still has a documented crash-safety gap after an unclean shutdown. Treat integrity errors as actionable and follow the upstream recovery guidance.

```sql
SELECT *
FROM turbovec.turbovec_check('items_embedding_flat_idx'::regclass);

REINDEX INDEX CONCURRENTLY items_embedding_flat_idx;
```

Only the index owner can run the integrity checker. Alert on `is_corrupt` and rebuild the affected index when the checker or a scan reports corruption; a successful version upgrade does not repair an already damaged index.
