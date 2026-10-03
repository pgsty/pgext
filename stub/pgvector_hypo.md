## Usage

Sources:

- [Official README.md](https://github.com/varadfromeast/pgvector_hypo/blob/154b8b9312fb9c55a5d0843f3d82c5aa5fdc3224/README.md)
- [Official pgvector_hypo.control](https://github.com/varadfromeast/pgvector_hypo/blob/154b8b9312fb9c55a5d0843f3d82c5aa5fdc3224/pgvector_hypo.control)
- [Official pgvector_hypo--0.1.0.sql](https://github.com/varadfromeast/pgvector_hypo/blob/154b8b9312fb9c55a5d0843f3d82c5aa5fdc3224/sql/pgvector_hypo--0.1.0.sql)
- [Official compatibility.md](https://github.com/varadfromeast/pgvector_hypo/blob/154b8b9312fb9c55a5d0843f3d82c5aa5fdc3224/docs/compatibility.md)

`pgvector_hypo` 0.1.0 lets PostgreSQL 16 estimate plans with hypothetical pgvector 0.8.x IVFFlat indexes, without constructing an index. It predicts planner choices and costs, not execution speed, recall or build time.

### Compare Plans

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pgvector_hypo;
CREATE TABLE items (id bigint, embedding vector(3));
SELECT * FROM pgvector_hypo_create_index(
  'CREATE INDEX ON items USING ivfflat (embedding vector_l2_ops) WITH (lists = 100)');
EXPLAIN SELECT id FROM items ORDER BY embedding <-> '[0,0,0]'::vector LIMIT 5;
SELECT * FROM pgvector_hypo_list_indexes;
SELECT pgvector_hypo_reset();
```

### Session State and Coverage

`pgvector_hypo_create_index()` accepts a CREATE INDEX statement and returns a synthetic OID/name. `pgvector_hypo_drop_index()` removes one registration; `pgvector_hypo_reset_index()` is an alias for reset. State belongs to the current session and disappears at disconnect. Supported keys include fixed-dimension `vector`, `halfvec` and `bit`, ordinary columns and supported immutable expressions, with partial and compatible partitioned-table definitions. Only the documented IVFFlat operator classes are covered; HNSW is outside this release.

### Operational Boundaries

Install with administrator privileges after the matching vector extension. No preload is required. Hypothetical indexes participate only in top-level plain `EXPLAIN`; ordinary execution, `EXPLAIN ANALYZE` and `EXPLAIN EXECUTE` cannot use them. No persistent index catalog row or relation file is created. Upstream still records two planner-choice mismatches at a build-sensitive fixed-bit boundary, so this is experimental evidence for planning, not a production performance guarantee.
