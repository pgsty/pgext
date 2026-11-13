## Usage

Sources:

- [README](https://api.pgxn.org/src/svs/svs-0.1.0/README.md)
- [Control](https://api.pgxn.org/src/svs/svs-0.1.0/svs.control)
- [Install SQL](https://api.pgxn.org/src/svs/svs-0.1.0/sql/svs--0.1.0.sql)
- [Server setup](https://api.pgxn.org/src/svs/svs-0.1.0/docs/build_guide/README.md)
- [User guide](https://api.pgxn.org/src/svs/svs-0.1.0/docs/USER_GUIDE.md)

`svs` adds the Vamana approximate-nearest-neighbor index for pgvector, with optional LeanVec and LVQ compression. This 0.1.0 distribution requires Linux x86-64, PostgreSQL 18.1 or later in the 18 series, pgvector 0.8.2 or later, and the Intel SVS C API. Compressed indexes require Intel CPUs; ARM, macOS and Windows are unsupported.

### Server and database setup

Preload `svs`, enable logical WAL and restart. Each Vamana index needs a logical replication slot; size `max_replication_slots` for these indexes and all other consumers. Without logical WAL the launcher stops. In the launcher database, selected by `svs.launcher_database` and defaulting to `postgres`, install the dependencies and register the target database in `vamana_databases`. Creating the extension requires a superuser.

```ini
shared_preload_libraries = 'svs'
wal_level = logical
max_replication_slots = 10
```

```sql
CREATE EXTENSION vector;
CREATE EXTENSION svs;
CREATE DATABASE svs_demo;
INSERT INTO vamana_databases (datname) VALUES ('svs_demo');
```

### Build and query

Connect to the registered database, install both extensions there, then create an index. Use an ordered distance expression with `LIMIT`. The example uses a complete three-dimensional query vector. The same index method supports `halfvec`; operator classes pair L2, inner product and cosine with `<->`, `<#>` and `<=>` respectively. Inner product is returned negated for ascending scans.

```sql
CREATE EXTENSION vector;
CREATE EXTENSION svs;
CREATE TABLE items (id bigint GENERATED ALWAYS AS IDENTITY, embedding vector(3));
INSERT INTO items (embedding) VALUES ('[1,0,0]'), ('[0,1,0]'), ('[0,0,1]');
CREATE INDEX items_vamana ON items USING vamana (embedding vector_l2_ops);
SELECT id FROM items ORDER BY embedding <-> '[1,0,0]'::vector LIMIT 2;
```

### Tuning and monitoring

Index options include `graph_degree`, `build_window_size` and `compression_type` (0: none, 1: LeanVec, 2: LVQ). The `svs.search_window_size` GUC controls query breadth; the similarly named index option is not the runtime fallback. Build memory is controlled by `svs.max_build_memory`, not `maintenance_work_mem`. Residency and search-scratch budgets are separate. Inspect `pg_stat_vamana_worker` and `pg_stat_vamana_worker_slot` views; `svs_restart_worker` requires update permission on the launcher registry.

### Operational boundaries

Indexes are owned by per-database background workers and replay logical WAL after crashes. Replication slots can retain WAL; monitor retention and resource limits. Pausing a registered database is different from removing it. Permanent removal uses `svs_teardown_database()` in the target before deleting the launcher registration; it drops Vamana indexes and checks ownership. Do not remove registration while retaining indexes. Concurrent writes may make searches wait; ANN results are approximate.
