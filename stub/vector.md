


## Usage

Sources:

- [pgvector v0.8.4 release](https://github.com/pgvector/pgvector/releases/tag/v0.8.4)
- [pgvector v0.8.4 README](https://github.com/pgvector/pgvector/blob/v0.8.4/README.md)
- [pgvector v0.8.4 CHANGELOG](https://github.com/pgvector/pgvector/blob/v0.8.4/CHANGELOG.md)
- [Local package metadata](../db/extension.csv)

`pgvector` provides vector similarity search inside PostgreSQL. The extension name is `vector`, while Pigsty packages it as `pgvector`. It supports exact search, approximate nearest-neighbor search with HNSW and IVFFlat indexes, and multiple vector representations for dense, half-precision, binary, and sparse embeddings.

v0.8.4 is a maintenance release after the 0.8.x HNSW/vacuum fixes. Use it instead of older 0.8.x builds when maintaining HNSW indexes under write-heavy workloads.

### Create and Query Vectors

```sql
CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE items (
  id bigserial PRIMARY KEY,
  embedding vector(3)
);

INSERT INTO items (embedding)
VALUES ('[1,2,3]'), ('[4,5,6]');

SELECT *
FROM items
ORDER BY embedding <-> '[3,1,2]'
LIMIT 5;
```

Common distance operators:

- `<->` for L2 distance
- `<#>` for negative inner product
- `<=>` for cosine distance
- `<+>` for L1 distance
- `<~>` for Hamming distance on binary vectors
- `<%>` for Jaccard distance on binary vectors

Because PostgreSQL indexes scan in ascending order, `<#>` returns the negative inner product; multiply by `-1` when displaying the actual inner product.

### Vector Types

```sql
CREATE TABLE embeddings (
  id bigserial PRIMARY KEY,
  dense      vector(768),
  half_dense halfvec(768),
  binary_sig bit(1024),
  sparse     sparsevec(100000)
);
```

`vector` is the standard single-precision type. Use `halfvec` to reduce storage and memory pressure, `bit` for binary signatures, and `sparsevec` for high-dimensional sparse vectors.

Aggregates such as `avg()` and `sum()` can be used with vector columns:

```sql
SELECT avg(embedding) FROM items;
```

### HNSW Indexes

HNSW gives strong speed/recall tradeoffs and does not require a training step.

```sql
CREATE INDEX items_embedding_hnsw
ON items USING hnsw (embedding vector_l2_ops);

SET hnsw.ef_search = 100;

SELECT *
FROM items
ORDER BY embedding <-> '[3,1,2]'
LIMIT 10;
```

Choose the operator class that matches the distance:

```sql
CREATE INDEX ON items USING hnsw (embedding vector_ip_ops);
CREATE INDEX ON items USING hnsw (embedding vector_cosine_ops);
CREATE INDEX ON items USING hnsw (embedding vector_l1_ops);
CREATE INDEX ON embeddings USING hnsw (half_dense halfvec_l2_ops);
CREATE INDEX ON embeddings USING hnsw (sparse sparsevec_l2_ops);
CREATE INDEX ON embeddings USING hnsw (binary_sig bit_hamming_ops);
```

Useful tuning settings include `hnsw.ef_search`, `hnsw.iterative_scan`, `hnsw.max_scan_tuples`, and `hnsw.scan_mem_multiplier`.

### IVFFlat Indexes

IVFFlat requires representative data before index creation because it trains cluster lists at build time.

```sql
CREATE INDEX items_embedding_ivfflat
ON items USING ivfflat (embedding vector_l2_ops)
WITH (lists = 100);

SET ivfflat.probes = 10;

SELECT *
FROM items
ORDER BY embedding <-> '[3,1,2]'
LIMIT 10;
```

Increase `lists` for larger tables and increase `ivfflat.probes` for higher recall. For filtered queries, test whether an exact btree filter, a partial vector index, or partitioning gives better plans.

### Filtering and Hybrid Search

Normal PostgreSQL filters can be combined with vector ordering:

```sql
CREATE INDEX ON items (tenant_id);

SELECT *
FROM items
WHERE tenant_id = 42
ORDER BY embedding <=> '[0.1,0.2,0.3]'
LIMIT 20;
```

For hybrid search, combine `pgvector` with PostgreSQL full text search, trigram search, or an external ranking expression:

```sql
SELECT id,
       ts_rank_cd(text_tsv, plainto_tsquery('database')) AS text_score,
       1 - (embedding <=> '[0.1,0.2,0.3]') AS vector_score
FROM docs
WHERE text_tsv @@ plainto_tsquery('database')
ORDER BY vector_score DESC
LIMIT 20;
```

### Maintenance

```sql
VACUUM items;
REINDEX INDEX CONCURRENTLY items_embedding_hnsw;
ANALYZE items;
```

HNSW indexes can be large and expensive to build. Use `maintenance_work_mem` for builds, monitor build notices, and schedule `REINDEX` when index bloat or recall drift matters.

### Caveats

- Pigsty local metadata may lag this upstream version; this stub tracks upstream pgvector 0.8.4 while the local package row may still show an older package version until the package catalog is refreshed.
- Use the operator class that matches the query operator. A cosine index will not accelerate an L2 `ORDER BY`.
- Approximate indexes trade exact recall for speed. Validate recall with representative data and query filters.
- Build IVFFlat after loading data. If data distribution changes substantially, rebuild the index.
- Keep pgvector updated when using HNSW with heavy writes and vacuum activity; v0.8.x includes important HNSW maintenance fixes.
