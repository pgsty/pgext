## Usage

Sources:

- [README.md](https://api.pgxn.org/src/tessera_pg/tessera_pg-1.2.1/README.md)
- [tessera_pg.control](https://api.pgxn.org/src/tessera_pg/tessera_pg-1.2.1/tessera_pg.control)
- [tessera_pg--1.0.sql](https://api.pgxn.org/src/tessera_pg/tessera_pg-1.2.1/tessera_pg--1.0.sql)
- [Makefile](https://api.pgxn.org/src/tessera_pg/tessera_pg-1.2.1/Makefile)

`tessera_pg` replaces supported join-aggregate plans with sparse matrix multiplication and exposes item-similarity functions. Distribution 1.2.1 installs SQL version 1.0. It targets PostgreSQL 18; the native engine requires OpenMP and AVX2 or newer on x86-64.

### Core Workflow

Enable the planner hook in the session, then install the SQL functions:

```sql
LOAD 'tessera_pg';
CREATE EXTENSION tessera_pg;
CREATE TABLE edges (k bigint NOT NULL, x bigint NOT NULL);
INSERT INTO edges VALUES (1,10),(1,20),(2,10),(2,30);
EXPLAIN SELECT a.x, b.x, COUNT(*)
FROM edges a JOIN edges b ON a.k = b.k
GROUP BY a.x, b.x;
SELECT * FROM tessera_sim('edges', 'k', 'x', 'jaccard', 0.5);
```

`Custom Scan (Tessera)` identifies an accepted rewrite. The hook requires valid NOT NULL constraints for the relevant grouping and summed columns. Unsupported query shapes use the normal planner. Supported cases include count, weighted products, HAVING, ordered pairs and two input relations; do not assume arbitrary SQL can be rewritten.

### Similarity Functions and Settings

`tessera_sim` returns bigint item pairs and a score; `tessera_sim_text` returns text pairs. Their arguments identify the relation, key and item columns, similarity measure, threshold, optional weight column, pair mode and thread count. Both are volatile and parallel-unsafe SQL functions.

`tessera.enabled` defaults on, `tessera.threads` defaults to 0, and `tessera.frontier` defaults to -1 for cost-model selection. Other controls include `tessera.min_expansion`, `tessera.min_expansion_cross` and `tessera.trace`. Default threading accounts for the leader and the query's `max_parallel_workers_per_gather` workers.

### Operational Boundaries

`LOAD` or `session_preload_libraries` activates the hook; a shared preload/restart is not mandatory. `CREATE EXTENSION` installs the relocatable SQL objects. Plan memory can reach `work_mem * hash_mem_multiplier`; spilling uses `temp_tablespaces` and remains subject to `temp_file_limit`. Measure execution plans, temporary disk and concurrency on the actual workload; source benchmarks do not guarantee production speedups.
