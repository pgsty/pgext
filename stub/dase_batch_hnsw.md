## Usage

Sources:

- [README.md](https://github.com/ljmzlh/DASE/blob/b2d0d4e9e2627867bdc9c607df7d3324c316b500/README.md)
- [dase_batch_hnsw.control](https://github.com/ljmzlh/DASE/blob/b2d0d4e9e2627867bdc9c607df7d3324c316b500/batch_hnsw/dase_batch_hnsw.control)
- [dase_batch_hnsw--1.0.sql](https://github.com/ljmzlh/DASE/blob/b2d0d4e9e2627867bdc9c607df7d3324c316b500/batch_hnsw/dase_batch_hnsw--1.0.sql)
- [Makefile](https://github.com/ljmzlh/DASE/blob/b2d0d4e9e2627867bdc9c607df7d3324c316b500/batch_hnsw/Makefile)

`dase_batch_hnsw` provides batched HNSW range probes for the DASE research system. Its SQL version is 1.0. It compiles HNSW code from the same checkout's modified pgvector and must use that matching fork and PostgreSQL server.

### Core workflow

Install the matching extension files first. As a superuser, enable the dependency and inspect the functions:

```sql
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION IF NOT EXISTS dase_batch_hnsw;
SELECT n.nspname, p.proname
FROM pg_proc p JOIN pg_namespace n ON n.oid = p.pronamespace
WHERE n.nspname = 'bh';
SELECT bh.last_stats();
```

The fixed `bh` schema exposes `range_join_all` and `last_stats`. The range-probe function accepts `left_sql`, `right_index`, `radius`, `ef`, `mode` and `n_seeds`. The left query supplies an int8 ID, a vector and an int4 group in processing order. Results contain `left_id`, `right_tid` and `dist`. Statistics describe the current backend's most recent range-probe call.

### Integration and boundaries

Upstream documents PostgreSQL 15 or later and supplies a PostgreSQL 18 Docker environment. The bundled `vector` fork supplies extra HNSW behavior; a stock pgvector installation is not an interchangeable dependency. The extension uses shared traversal to construct semantic join candidates; approximate probing can return a different pair set from the other DASE method.

The control file is nonrelocatable and the functions are native C code. The documented workflow uses a superuser. Only pass trusted SQL, matching vector dimensions and the intended HNSW index; preserve the matching fork for upgrades and recovery. No mandatory preload is documented. DASE's datasets, Python runner and materialized join tables are separate from installing this extension, and the repository does not provide an explicit license for this component.
