## Usage

Sources:

- [Pinned upstream overview](https://github.com/darthunix/pg_fusion/blob/08ad73af18c909bc50b95114c3e4316ffe01a3ed/README.md)
- [Pinned PostgreSQL 17 quick start](https://github.com/darthunix/pg_fusion/blob/08ad73af18c909bc50b95114c3e4316ffe01a3ed/docs/quickstart.md)
- [Pinned query-support reference](https://github.com/darthunix/pg_fusion/blob/08ad73af18c909bc50b95114c3e4316ffe01a3ed/docs/query-support.md)
- [Pinned limitations](https://github.com/darthunix/pg_fusion/blob/08ad73af18c909bc50b95114c3e4316ffe01a3ed/docs/limitations.md)
- [Pinned extension control file](https://github.com/darthunix/pg_fusion/blob/08ad73af18c909bc50b95114c3e4316ffe01a3ed/pg/extension/pg_fusion.control)

pg_fusion version 24.0.0 is an experimental analytical execution extension. PostgreSQL still owns parsing, catalogs, snapshots, MVCC visibility, heap scans, TOAST handling, and client-facing tuples; eligible scan rows cross shared-memory Arrow pages to one background worker, where Apache DataFusion runs selected operators.

### Preload and first query

The library must be loaded at postmaster start because it installs hooks, reserves shared memory, and starts a worker:

```conf
shared_preload_libraries = 'pg_fusion'
```

After restarting, create the extension and enable its strict path only in the session being tested:

```sql
CREATE EXTENSION IF NOT EXISTS pg_fusion;

CREATE TABLE fusion_demo AS
SELECT i AS id, i % 10 AS group_id, i::double precision AS value
FROM generate_series(1, 1000000) AS i;

ANALYZE fusion_demo;
SET pg_fusion.enable = on;

EXPLAIN
SELECT count(*), avg(value)
FROM fusion_demo
WHERE group_id >= 0;

SELECT count(*), avg(value)
FROM fusion_demo
WHERE group_id >= 0;
```

An eligible plan contains Custom Scan (PgFusionScan). The current quick start targets a pgrx PostgreSQL 17 development cluster. Installation is superuser-only.

### Experimental boundary

The supported surface is a conservative subset of analytical SELECT statements. With the runtime enabled, unsupported shapes fail with a pg_fusion planning error rather than silently falling back. Current exclusions include non-SELECT work, modifying CTEs, bound or prepared parameters, SPI and PL/pgSQL-internal execution, table functions, surviving subqueries, and unsupported PostgreSQL types or semantics.

Tuple-to-Arrow conversion and transport are real overhead. Wide projections, raw-row returns, weak filter pushdown, and light computation can be slower than native PostgreSQL. One worker and fixed shared-memory pools are shared cluster resources; memory limits, spill directories, page counts, control slots, and runtime-filter sizing require capacity planning. Worker spill uses operating-system temporary storage, not PostgreSQL temporary-file controls.

Treat pg_fusion as engineering-evaluation software. Compare results and EXPLAIN ANALYZE output with the runtime both off and on, test documented type and collation restrictions, and begin with controlled analytical workloads rather than production traffic.
