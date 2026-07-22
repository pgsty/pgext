## Usage

Sources:

- [Official version 0.4.0 README](https://github.com/tensorchord/pgvecto.rs/blob/2b290b34e8ba69104ea2f800fa53328c6ed6c236/README.md)
- [Official version 0.4.0 extension SQL](https://github.com/tensorchord/pgvecto.rs/blob/2b290b34e8ba69104ea2f800fa53328c6ed6c236/sql/install/vectors--0.4.0.sql)
- [Official migration guide from PGVecto.rs to VectorChord](https://docs.vectorchord.ai/vectorchord/admin/migration.html)

`vectors` is the canonical extension name of PGVecto.rs 0.4.0, a Rust/pgrx vector-similarity extension. It supplies dense, half-precision, sparse, and binary vector types plus its own `vectors` index access method. Upstream now advises users to migrate to VectorChord; choose `vectors` mainly for maintaining an existing PGVecto.rs deployment.

### Setup and Search

The library must be included in `shared_preload_libraries` and PostgreSQL restarted before creating the extension. Its objects are installed in the `vectors` schema.

```sql
CREATE EXTENSION vectors;

CREATE TABLE items (
    id bigserial PRIMARY KEY,
    embedding vectors.vector(3) NOT NULL
);

INSERT INTO items(embedding)
VALUES ('[1,2,3]'), ('[4,5,6]');

SELECT id
FROM items
ORDER BY embedding <-> '[3,2,1]'::vectors.vector
LIMIT 5;

CREATE INDEX items_embedding_idx
ON items USING vectors (embedding vectors.vector_l2_ops)
WITH (options = '[indexing.hnsw]');
```

For dense `vector`, `<->` is squared Euclidean distance, `<#>` is negative dot product, and `<=>` is cosine distance. Select the matching `vector_l2_ops`, `vector_dot_ops`, or `vector_cos_ops` class for the query operator. Other types include `vecf16`, `svector`, and `bvector`, with corresponding operator classes; binary vectors also support Hamming/Jaccard-oriented operations.

Search GUCs include `vectors.hnsw_ef_search`, `vectors.ivf_nprobe`, `vectors.search_mode`, and quantization rerank/fast-scan settings. Tune recall and latency with representative data rather than treating defaults or published benchmarks as guarantees.

### Lifecycle and Migration

PGVecto.rs index storage is managed separately from PostgreSQL heap storage, and the upstream overview warns that index WAL support is incomplete. Exercise crash recovery, replication, backup/restore, upgrades, `REINDEX`, and extension removal on the exact 0.4.0 build. Index columns require a fixed dimension.

The official VectorChord migration path installs `vchord`/`vector` in a different schema, records and drops PGVecto.rs indexes, converts columns through documented casts, and recreates indexes. Type conversion takes an `ACCESS EXCLUSIVE` lock and therefore causes downtime. `svector` also needs index-number conversion because PGVecto.rs uses zero-based sparse indexes while pgvector uses one-based indexes. Keep the old extension until every dependent column and index has been migrated and verified.
