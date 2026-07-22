## Usage

Sources:

- [Official architecture and quick start](https://github.com/mayflower/pg_turboquant/blob/4d27dc6d1c5de7fd3dc4a37fbec89ece597edd63/README.md)
- [Index options reference](https://github.com/mayflower/pg_turboquant/blob/4d27dc6d1c5de7fd3dc4a37fbec89ece597edd63/docs/reference/index-options.md)
- [Compatibility reference](https://github.com/mayflower/pg_turboquant/blob/4d27dc6d1c5de7fd3dc4a37fbec89ece597edd63/docs/reference/compatibility.md)

`pg_turboquant` adds the `turboquant` approximate-nearest-neighbor index access method for pgvector `vector` and `halfvec` columns. Its compact fast path targets normalized cosine or inner-product retrieval, with exact reranking performed explicitly in SQL.

### Core Workflow

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_turboquant;

CREATE INDEX docs_embedding_tq_idx
ON docs USING turboquant (embedding tq_cosine_ops)
WITH (
    bits = 4,
    lists = 0,
    transform = 'hadamard',
    normalized = true
);

SELECT *
FROM tq_rerank_candidates(
    'docs'::regclass,
    'id',
    'embedding',
    '[1,0,0,0]'::vector(4),
    'cosine',
    50,
    10
);
```

Set `lists = 0` for a flat approximate scan or a positive value for IVF routing. The documented production fast lane uses `bits = 4`, `transform = 'hadamard'`, normalized vectors, and automatic lanes; benchmark recall, latency, index size, build time, WAL, and concurrent writes on representative embeddings before choosing values.

### SQL and Observability Surface

- `tq_approx_candidates(...)` returns approximate candidates and `tq_rerank_candidates(...)` applies the documented SQL-side exact rerank workflow.
- `tq_recommended_query_knobs(...)` provides query-setting guidance.
- `tq_index_metadata(...)`, `tq_index_heap_stats(...)`, and `tq_last_scan_stats()` expose index, heap, and most-recent-backend-scan diagnostics.
- `tq_maintain_index(...)` performs the extension's lightweight delta-tier merge or compaction work.

### Compatibility and Maintenance

The reviewed contract targets PostgreSQL 16 and 17 and pgvector `0.8.1`. L2 and non-normalized scans use decoded-vector compatibility scoring, not the faithful quantized fast path. The current format is `v2`; older indexes require `REINDEX` or recreation. Exact reranking remains outside the access method, MVCC visibility remains with PostgreSQL's executor, and the implementation uses generic WAL. Pin versions and validate upgrades, crash recovery, replicas, index-only scans, metadata filters, and maintenance recommendations before production use.
