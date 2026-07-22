## Usage

Sources:

- [Official Azure Database for PostgreSQL guide](https://learn.microsoft.com/en-us/azure/postgresql/extensions/how-to-use-pgdiskann)
- [Official Microsoft Research DiskANN project](https://www.microsoft.com/en-us/research/project/project-akupara-approximate-nearest-neighbor-search-for-large-scale-semantic-search/)

`pg_diskann` is the Azure Database for PostgreSQL Flexible Server extension for DiskANN approximate nearest-neighbor indexes over `vector` columns. It is provider-managed rather than a cataloged open-source PostgreSQL extension; availability, supported versions, and operational limits are controlled by Azure.

### Enable and Query

Allow `pg_diskann` for the Azure server, then create it in each database together with its `vector` dependency:

```sql
CREATE EXTENSION IF NOT EXISTS pg_diskann CASCADE;

CREATE TABLE demo (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  embedding public.vector(3)
);

CREATE INDEX demo_embedding_diskann_idx
ON demo USING diskann (embedding vector_cosine_ops);

SELECT id, embedding
FROM demo
ORDER BY embedding <=> '[2.0,3.0,4.0]'
LIMIT 5;
```

The index is approximate. Measure recall and latency with production-shaped vectors, filters, and concurrency, and verify index use with `EXPLAIN (ANALYZE, BUFFERS)`.

### Tuning Index

Index storage parameters include `max_neighbors`, `l_value_ib`, `product_quantized`, `pq_param_num_chunks`, and `pq_param_training_samples`. Product quantization is available from version `0.6`, reduces memory, and is required by the Azure guide for dimensions above 2,000, up to 16,000:

```sql
CREATE INDEX demo_embedding_diskann_pq_idx
ON demo USING diskann (embedding vector_cosine_ops)
WITH (
  product_quantized = true,
  pq_param_num_chunks = 0,
  pq_param_training_samples = 0
);
```

`diskann.l_value_is` trades query work for recall. `diskann.iterative_search` accepts `relaxed_order`, `strict_order`, or `off`; relaxed order may not emit exact distance order, while noniterative search cannot return more than its search-list size. Full-vector reranking of a larger inner candidate set can improve final ordering after a quantized search.

### Operational Boundaries

Index construction can require substantial memory, storage, and time; Azure documents `maintenance_work_mem`, table `parallel_workers`, and server worker limits as build controls. Changing `max_worker_processes` requires a restart. Do not disable `enable_seqscan` globally to force an index: a local setting affects every table in that transaction and can hide planning or recall problems. The catalog version `0.6.5` should be confirmed against the extension version actually offered in the target Azure region before using version-specific options, and index rebuild, backup/restore, replica, failover, and upgrade behavior should be tested through Azure's supported procedures.
