## Usage

Sources:

- [README.md](https://github.com/pg-cuvs/pg_cuvs/blob/v0.9.2/README.md)
- [pg_cuvs.control](https://github.com/pg-cuvs/pg_cuvs/blob/v0.9.2/pg_cuvs.control)
- [docs/playbooks/release-upgrade.md](https://github.com/pg-cuvs/pg_cuvs/blob/v0.9.2/docs/playbooks/release-upgrade.md)
- [sql/pg_cuvs--0.9.0--0.9.1.sql](https://github.com/pg-cuvs/pg_cuvs/blob/v0.9.2/sql/pg_cuvs--0.9.0--0.9.1.sql)
- [sql/pg_cuvs--0.9.1--0.9.2.sql](https://github.com/pg-cuvs/pg_cuvs/blob/v0.9.2/sql/pg_cuvs--0.9.1--0.9.2.sql)
- [Release](https://github.com/pg-cuvs/pg_cuvs/releases/tag/v0.9.2)

`pg_cuvs` 0.9.2 uses a separate NVIDIA cuVS daemon for GPU vector search and GPU-assisted HNSW construction. The documented target is Linux/PostgreSQL 16 with pgvector 0.7.0 or newer. Match the extension, daemon, CUDA/cuVS libraries and required SDK patches; a CPU CI pass is not evidence of GPU correctness.

### Enable and Query

Append `pg_cuvs` to the existing preload list, preserve other libraries, and configure matching backend/daemon paths:

```ini
shared_preload_libraries = 'pg_cuvs'
cuvs.socket_path = '/var/run/postgresql/.s.pg_cuvs'
cuvs.index_dir = '/var/lib/postgresql/cuvs_indexes'
```

Restart PostgreSQL, start the matching `pg_cuvs_server` with access to the same socket and persistent index directory, and wait until it is listening. A privileged installer then enables the non-relocatable extension:

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_cuvs;
CREATE TABLE items (id bigint PRIMARY KEY, embedding vector(3));
INSERT INTO items VALUES (1, '[1,0,0]'), (2, '[0,1,0]'), (3, '[0,0,1]');
CREATE INDEX items_cagra ON items USING cagra (embedding vector_l2_ops);
SELECT id FROM items ORDER BY embedding <-> '[1,0,0]'::vector LIMIT 2;
EXPLAIN SELECT id FROM items ORDER BY embedding <-> '[1,0,0]'::vector LIMIT 2;
```

`cagra` provides approximate GPU search, `flat` exhaustive GPU search, and `ivfpq` compressed approximate search. `pg_cuvs_hnsw` builds with the GPU and serves standard pgvector HNSW pages on CPU. `cuvs.k` controls candidate count independently of SQL LIMIT; filters and approximate search require workload-specific recall checks. Use EXPLAIN to confirm the actual route.

### Operations and Upgrade

`pg_stat_gpu_search` exposes search statistics. `pg_cuvs_hnsw_connectivity_status(regclass)` returns an optional structural snapshot recorded at build time; it does not track later INSERT/DELETE/VACUUM or guarantee recall. The older two-step import API was removed; use `CREATE INDEX ... USING pg_cuvs_hnsw` for current builds.

Install both matching binaries, restart PostgreSQL and then the daemon, and run `ALTER EXTENSION pg_cuvs UPDATE TO '0.9.2'` in each database. The repository includes the SQL upgrade chain from 0.1.0. Version 0.9.2 changes no SQL object; it fixes socket connection retries and the daemon backlog for concurrent micro-batches. Its real-GPU evidence covers a specific A100/PG16 environment, not every GPU or driver.

### Failure and Persistence Boundaries

Daemon loss after planning raises an error; the executing statement cannot silently swap to a CPU node. Roll back before retrying and inspect the new plan. `enable_cuvs=off` disables the GPU planning path. GPU sidecars are outside PostgreSQL WAL and need their own backup, restore and replica procedure. IVF-PQ inserts are not indexed incrementally; rebuild after data changes and verify sidecar recovery. Stale-marker delivery remains best effort during outages, and memory preflight does not reserve memory against other processes. The project is active; a deprecated helper is not a deprecation of the extension.
