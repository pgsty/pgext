## Usage

Sources:

- [Official README.md](https://github.com/Mateus504R35/pgfaiss/blob/8cf777c2f38379958d0be447037ae37affbe78a3/README.md)
- [Official faiss_pg.control](https://github.com/Mateus504R35/pgfaiss/blob/8cf777c2f38379958d0be447037ae37affbe78a3/faiss_pg.control)
- [Official faiss_pg--0.2.sql](https://github.com/Mateus504R35/pgfaiss/blob/8cf777c2f38379958d0be447037ae37affbe78a3/faiss_pg--0.2.sql)
- [Official faiss_pg.cpp](https://github.com/Mateus504R35/pgfaiss/blob/8cf777c2f38379958d0be447037ae37affbe78a3/faiss_pg.cpp)

`faiss_pg` 0.2, distributed as pgfaiss, builds Faiss files from PostgreSQL arrays and searches them through SQL. It is a research extension, not a PostgreSQL index access method; there is no planner-integrated index scan.

### Build and Search

```sql
CREATE EXTENSION faiss_pg;
CREATE TABLE items (id bigint PRIMARY KEY, embedding real[] NOT NULL);
INSERT INTO items VALUES (1, ARRAY[0.1,0.2,0.3]::real[]),
                         (2, ARRAY[0.8,0.7,0.6]::real[]);
SELECT faiss_build_index('items_flat_l2', 'public.items'::regclass,
  'id'::name, 'embedding'::name, 'l2', 'flat', false,
  '{"fetchBatchSize":10000}');
SELECT * FROM faiss_search('items_flat_l2', ARRAY[0.1,0.2,0.3]::real[], 2);
```

### Inputs and Results

The source key must be an integer or bigint; embeddings are one-dimensional `real[]` values of equal length, without null keys, arrays or elements. Flat gives exact search; HNSW and IVF-Flat give approximate search. L2 results are squared distances with smaller values preferred; inner product and normalized cosine return similarities with larger values preferred. `faiss_search_batch()` accepts flattened query vectors and returns a zero-based query number, ID and distance. Search parameters include `efSearch` and `nprobe`.

### Files, Updates and Access

`faiss_indexes` stores file paths and source metadata. The default server-side directory is /var/lib/postgresql/faiss_indexes; the PostgreSQL OS user needs write access. Source-table writes mark an index dirty and searches then fail until it is rebuilt. Builds take a SHARE lock on the source table. Files are cached per backend; `faiss_clear_cache()` and `faiss_clear_all_cache()` release that cache. External files need their own backup and lifecycle management. Install as an administrator and restrict function, metadata-table and filesystem access to trusted operators; this prototype does not provide a separate file-access permission boundary.

### Requirements

The source uses the Faiss CPU library 1.13.1, C++17, OpenMP and BLAS/LAPACK. PL/pgSQL implements dirty-tracking helpers. The README shows PostgreSQL 16 on Linux, but does not define a supported major-version range. No shared-preload step is documented.

The SQL leaves default `PUBLIC` function execution in place, and the C file operations have no explicit superuser or file-role ACL check. Treat caller-selected `indexDir` and metadata file paths as trusted inputs; ordinary table privileges alone are not a complete filesystem-access boundary.
