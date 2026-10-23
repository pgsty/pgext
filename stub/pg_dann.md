## Usage

Sources:

- [Official source](https://github.com/A-Jimider/pg_dann/blob/351737cd0d9561a290e1c06fa687e1caeea0ddb6/README.md)
- [Official source](https://github.com/A-Jimider/pg_dann/blob/351737cd0d9561a290e1c06fa687e1caeea0ddb6/pg_dann.control)
- [Official source](https://github.com/A-Jimider/pg_dann/blob/351737cd0d9561a290e1c06fa687e1caeea0ddb6/sql/pg_dann--0.1.0.sql)

`pg_dann` 0.1.0 is an experimental disk-oriented approximate vector index using IVF lists and SQ8 or RaBitQ quantization. It depends on `vector`. The reviewed source was exercised on PostgreSQL 16; wider compatibility is not established here.

### Core Workflow

Install with a superuser after installing the dependency. No preload is required.

```sql
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION pg_dann;
CREATE TABLE items (id bigint, embedding vector(3));
INSERT INTO items VALUES (1, '[1,2,3]'), (2, '[3,2,1]');
CREATE INDEX items_dann ON items USING dann (embedding dann_l2_ops)
  WITH (nlist = 1, quantizer = 'sq8');
SELECT * FROM items ORDER BY embedding <-> '[1,2,3]' LIMIT 10;
```

For useful training, build indexes on representative data. `dann_l2_ops`, `dann_ip_ops` and `dann_cosine_ops` correspond to L2, inner product and cosine operators respectively; the query operator must match the chosen operator class.

### Tuning and Inspection

`dann.probes` controls the lists searched. `dann.rerank_factor` and `dann.enable_rerank` control candidate reranking, while `dann.max_latency_ms` sets a latency budget. Index options include `nlist` and `quantizer`; automatic quantization chooses according to dimension count. `dann_scan_stats()` and `dann_list_stats()` expose scan and list information.

Reranking uses full-precision candidate vectors, but cannot recover neighbors never selected as candidates. Measure recall as well as latency for your own data. Non-finite vector values are rejected.

### Maintenance Boundary

Inserts are serialized through the metapage in this revision. Parallel scans and builds are not implemented. Vacuum removes dead entries but does not reclaim empty index pages; rebuilding may be needed to recover space. Treat this as a prototype and verify recovery and update behavior before relying on it.
