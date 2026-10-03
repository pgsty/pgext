## Usage

Sources:

- [0.9.1 README](https://github.com/timescale/pgvectorscale/blob/0.9.1/README.md)
- [Control and dependency](https://github.com/timescale/pgvectorscale/blob/0.9.1/pgvectorscale/vectorscale.control)
- [0.9.1 migration SQL](https://github.com/timescale/pgvectorscale/blob/0.9.1/pgvectorscale/sql/vectorscale--0.9.0--0.9.1.sql)
- [0.9.1 security and upgrade notes](https://github.com/timescale/pgvectorscale/releases/tag/0.9.1)

`vectorscale` adds the StreamingDiskANN approximate vector index to pgvector. Its `diskann` access method supports L2, inner product, cosine distance, and label filtering. It requires `vector`; creating the extension requires superuser privileges. Version 0.9.1 validates vector types, dimensions, and stored datum layouts to address crashes, memory disclosure, and out-of-bounds writes.

### Core Workflow

Declare a concrete vector dimension and select the operator class matching the query distance:

```sql
CREATE EXTENSION vectorscale CASCADE;
CREATE TABLE documents (
    id bigserial PRIMARY KEY,
    contents text,
    embedding vector(3),
    labels smallint[]
);
INSERT INTO documents(contents, embedding, labels)
VALUES ('PostgreSQL search', '[1,2,3]', ARRAY[1,3]::smallint[]);
CREATE INDEX documents_diskann ON documents
USING diskann (embedding vector_cosine_ops, labels);

SELECT id, contents FROM documents
WHERE labels && ARRAY[1]::smallint[]
ORDER BY embedding <=> '[1,2,3]'::vector
LIMIT 10;
```

Use `vector_l2_ops` with `<->`, `vector_ip_ops` with `<#>`, or `vector_cosine_ops` with `<=>`. Labels use `smallint[]`, and `&&` means any requested label overlaps. Ordinary WHERE conditions are also supported, but selective filters can reduce the returned result count and need workload-specific testing.

### Tuning and Ordering

`diskann.query_search_list_size` controls extra graph-search candidates (default 100); `diskann.query_rescore` controls exact rescoring (default 50, 0 disables it):

```sql
SET diskann.query_search_list_size = 200;
SET diskann.query_rescore = 100;
```

Build options include `storage_layout`, `num_neighbors`, `search_list_size`, and `num_dimensions`. Compressed memory-optimized storage is the default. Increase maintenance memory only after considering concurrent builds and dataset size. Parallel builds require the supported compression layout and do not support label columns in this version.

DiskANN returns relaxed distance ordering. Sort a materialized result set when strict ordering is required; this reorders the retrieved candidates without making approximate search exhaustive. Null vectors are not indexed, null labels act as empty arrays, and null array elements are ignored. Index creation on UNLOGGED tables is unsupported.

### Upgrade to 0.9.1

Install matching extension files and update each database:

```sql
ALTER EXTENSION vectorscale UPDATE TO '0.9.1';
```

The upgrade binds operator classes to pgvector's actual installation schema and checks existing bindings. If an existing operator class references the wrong type or operators, the upgrade aborts; follow the release instructions to drop the affected operator class and recreate extension objects, accounting for dependent indexes.

**DiskANN now requires a valid `vector(N)` column type.** An unconstrained vector column cannot be indexed. Existing indexes with invalid persisted dimensions raise errors during scans, inserts, and vacuum. Correct the column type, then **drop and recreate the affected indexes**: `REINDEX` cannot repair this condition. Valid existing indexes do not need a blanket rebuild merely because 0.9.1 was installed. The SQL surface is not relocatable after extension creation.
