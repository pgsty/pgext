## Usage

Sources:

- [README.md](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/README.md)
- [sql/pg_vectorsearch.sql](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/sql/pg_vectorsearch.sql)
- [docs/maintenance.md](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/docs/maintenance.md)
- [docs/tuning.md](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/docs/tuning.md)
- [CHANGELOG.md](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/CHANGELOG.md)
- [VERSION](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/VERSION)
- [meson.build](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/meson.build)
- [pg_vectorsearch.control.in](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/pg_vectorsearch.control.in)

`pg_vectorsearch` 0.2.0 is an alpha extension for approximate nearest-neighbor search with the PRISM index. It supports PostgreSQL 18 and 19, native `vec32` / `vec16` types and optional pgvector interoperability. It is intended for testing, with evolving behavior and maintenance requirements.

### Core Workflow

```sql
CREATE EXTENSION pg_vectorsearch;
CREATE TABLE vectorsearch_items (
    id serial PRIMARY KEY,
    embedding vec32(3) STORAGE PLAIN
);
INSERT INTO vectorsearch_items (embedding)
SELECT ARRAY[i::real, (i % 7)::real, (i % 5)::real]::vec32
FROM generate_series(1, 1000) AS i;
CREATE INDEX vectorsearch_items_prism ON vectorsearch_items
USING prism (embedding vec32_l2_ops);
SET prism.nprobe = 40;
SELECT id FROM vectorsearch_items
ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
```

### Operational Boundaries

Install as a superuser; no preload is required and the extension is not relocatable. The `prism` access method has `vec32_l2_ops`, `vec32_ip_ops`, `vec32_cosine_ops` and corresponding vec16 operator classes; `<->`, `<#>` and `<=>` express L2, negative inner product and cosine distance. `prism.nprobe` trades recall for query work; zero selects automatic tuning, not exact recall.

Load representative data before building the index. L2 and cosine indexes support up to 8192 dimensions: any dimension through 1968, then a power of two times an odd factor at most eight. Inner-product indexes remain limited to 1968 for vec32 and 3584 for vec16. Expression and unlogged indexes are unsupported. Keep `prism.leaf_refine_threshold` at zero. Filtering can return fewer rows than LIMIT, and candidate buffers bounded by `work_mem` can truncate large or unbounded results. Measure against an exact baseline.

Inspect `prism_index_settings` and follow the maintenance guide for churn. `prism_rebalance` supports only flat RaBitQ centroid-page indexes with `centroid_fastscan` off, not the default multilevel format. DELETE leaves tombstones, and retired pages are reclaimed only by REINDEX. There is no in-place upgrade from the earlier release candidates.

### Upgrade to 0.2.0

Install matching files, start a fresh session that has not loaded the old library, update the SQL extension and rebuild every PRISM index:

```sql
ALTER EXTENSION pg_vectorsearch UPDATE TO '0.2.0';
REINDEX INDEX vectorsearch_items_prism;
```

Old index format 0x06 is rejected by the 0x07 reader. The two library versions cannot share one session. `vec32_norm` is renamed to `l2_norm`; the vec16 overload of `vec32_dims` becomes `vec16_dims`, while `vector_dims` provides the pgvector-compatible interface. Update direct callers of renamed functions; dependent views retain function OIDs.

`posting_reference_format` defaults to float16. A reference element beyond ±65504 causes a build or split error; use float32 when needed. The vec16 type now rejects out-of-range input. Existing filter, buffer-size and maintenance limits still apply.
