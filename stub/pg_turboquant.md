## Usage

Sources:

- [Current upstream architecture and quick start](https://github.com/mayflower/pg_turboquant/blob/4d27dc6d1c5de7fd3dc4a37fbec89ece597edd63/README.md)
- [Index options reference](https://github.com/mayflower/pg_turboquant/blob/4d27dc6d1c5de7fd3dc4a37fbec89ece597edd63/docs/reference/index-options.md)
- [Compatibility reference](https://github.com/mayflower/pg_turboquant/blob/4d27dc6d1c5de7fd3dc4a37fbec89ece597edd63/docs/reference/compatibility.md)

`pg_turboquant` adds the `turboquant` approximate-nearest-neighbor index access method for pgvector `vector` and `halfvec` columns. Its compact four-bit path targets normalized cosine or inner-product retrieval, with SQL-side exact reranking.

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_turboquant;
CREATE INDEX docs_embedding_tq_idx
ON docs USING turboquant (embedding tq_cosine_ops)
WITH (bits = 4, lists = 0, transform = 'hadamard', normalized = true);
```

The reviewed compatibility contract targets PostgreSQL 16/17 and pgvector `0.8.1`. L2 and non-normalized scans use a fallback rather than the faithful fast path, and exact reranking remains outside the access method. Format changes can require `REINDEX`; measure recall, latency, WAL, maintenance, concurrent writes, and index size on representative embeddings before deployment.
