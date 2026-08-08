## Usage

Sources:

- [pg_turbovec v1.29.0 README](https://codeberg.org/gregburd/pg_turbovec/src/tag/v1.29.0/README.md)
- [pg_turbovec v1.29.0 changelog](https://codeberg.org/gregburd/pg_turbovec/src/tag/v1.29.0/CHANGELOG.md)
- [pg_turbovec v1.29.0 control file](https://codeberg.org/gregburd/pg_turbovec/src/tag/v1.29.0/pg_turbovec.control)
- [Partitioned-scale guide](https://codeberg.org/gregburd/pg_turbovec/src/tag/v1.29.0/docs/PARTITIONED_SCALE.md)
- [Pigsty package matrix](https://pgext.cloud/ext/pg_turbovec)

`pg_turbovec` provides a dense `turbovec.vector` type and a `turbovec` approximate-nearest-neighbor index access method. It quantizes floating-point coordinates to 2, 3, or 4 bits and reranks candidates against heap vectors. Use it when index footprint and search throughput matter more than exact nearest-neighbor recall.

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

### Build a TurboVec Index

```sql
CREATE INDEX items_embedding_idx ON items
USING turbovec (embedding vec_cosine_ops)
WITH (bit_width = 4);

SET turbovec.probes = 32;

SELECT id
FROM items
ORDER BY embedding <=> '[3,1,2]'::turbovec.vector
LIMIT 10;
```

Use `vec_cosine_ops` with `<=>` or `vec_ip_ops` with `<#>`. `bit_width = 4` is the default and generally favors recall; 2-bit indexes are smaller but need workload-specific recall testing. `CREATE INDEX CONCURRENTLY` is supported.

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
- Upstream v1.29 targets PostgreSQL 13-18 and labels PostgreSQL 19 support experimental; current Pigsty packages cover PostgreSQL 14-18 and provide the matching OpenBLAS-linked binary.
- The current Pigsty package/source metadata remains at `1.28.3`, while upstream is `1.29.0`. On a packaged 1.28.3 installation, do not assume the 1.28.4 integrity checker or 1.29 features are present.
- Upstream 1.28.4 fixes persisted row-count drift that could corrupt the index ID table and adds `turbovec.turbovec_check(regclass)`. Prioritize a package upgrade to at least 1.28.4 before write-heavy production use. An already corrupt index still needs `REINDEX` or drop/recreate recovery.
- Version 1.29.0 is additive and does not require reindexing when upgrading from a healthy 1.28.4 index. `ALTER EXTENSION pg_turbovec UPDATE TO '1.29.0'` is sufficient after the new files are installed.
- Although the 1.29 reloption parser accepts `bit_width = 1`, end-to-end one-bit indexing is not implemented and `CREATE INDEX` intentionally errors. Use 2, 3, or 4.
- The on-disk ID table still has a documented crash-safety gap after an unclean shutdown. Treat integrity errors as actionable and follow the upstream recovery guidance.
