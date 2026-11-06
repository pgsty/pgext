## Usage

Sources:

- [README v1.3.0](https://github.com/GyuyeongKim/Tessera-public/blob/v1.3.0/README.md)
- [v1.3.0 release](https://github.com/GyuyeongKim/Tessera-public/releases/tag/v1.3.0)
- [tessera_pg.control](https://api.pgxn.org/src/tessera_pg/tessera_pg-1.3.0/tessera_pg.control)
- [tessera_pg--1.0.sql](https://api.pgxn.org/src/tessera_pg/tessera_pg-1.3.0/tessera_pg--1.0.sql)
- [Makefile](https://api.pgxn.org/src/tessera_pg/tessera_pg-1.3.0/Makefile)

`tessera_pg` replaces supported join-aggregate plans with sparse matrix multiplication and exposes item-similarity functions. Distribution 1.3.0 installs SQL version 1.0. It targets PostgreSQL 18; the native engine requires OpenMP and AVX2 or newer on x86-64.

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

`Custom Scan (Tessera)` identifies an accepted rewrite. Version 1.3.0 supports `COUNT(DISTINCT a.k)`, `AVG`, one-sided `SUM`, multicolumn or expression keys, additional grouping columns and filters on either input. Grouping columns may contain NULLs, which form separate groups; summed columns still need valid NOT NULL constraints. Inputs may also use subqueries, views, CTEs, partitioned or foreign tables and row-level-security plans, with pruning for HAVING and top-N queries. `MIN`, `MAX`, outer joins, nondeterministic collations and other unsupported shapes keep the normal plan; the adapter recognizes only PostgreSQL built-in operators and aggregates.

### Similarity Functions and Settings

`tessera_sim` returns bigint item pairs and a score; `tessera_sim_text` returns text pairs. Their arguments identify the relation, key and item columns, similarity measure, threshold, optional weight column, pair mode and thread count. Both are volatile and parallel-unsafe SQL functions.

`tessera.enabled` defaults on, `tessera.threads` defaults to 0, and `tessera.frontier` defaults to -1 for cost-model selection. Other controls include `tessera.min_expansion`, `tessera.min_expansion_cross` and `tessera.trace`. Default threading accounts for the leader and the query's `max_parallel_workers_per_gather` workers.

### Operational Boundaries

`LOAD` or `session_preload_libraries` activates the hook; a shared preload/restart is not mandatory. `CREATE EXTENSION` installs the relocatable SQL objects. Plan memory can reach `work_mem * hash_mem_multiplier`; spilling uses `temp_tablespaces` and remains subject to `temp_file_limit`. Measure execution plans, temporary disk and concurrency on the actual workload; source benchmarks do not guarantee production speedups.
