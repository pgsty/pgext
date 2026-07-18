## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/TantorLabs/pg_dphyp/blob/f76770ca001acca74f80ff6111d224aa44f415ff/README.md)
- [Extension control file](https://github.com/TantorLabs/pg_dphyp/blob/f76770ca001acca74f80ff6111d224aa44f415ff/pg_dphyp.control)
- [Join-search hook implementation](https://github.com/TantorLabs/pg_dphyp/blob/f76770ca001acca74f80ff6111d224aa44f415ff/src/pg_dphyp.c)

`pg_dphyp` is a headless planner module implementing hypergraph-based dynamic-programming join enumeration. It is enabled by loading the library; the catalog records no SQL extension objects.

```conf
shared_preload_libraries = 'pg_dphyp'
pg_dphyp.enabled = on
pg_dphyp.min_relations = 8
pg_dphyp.max_relations = 16
pg_dphyp.cj_strategy = 'pass'
```

The module can reduce search space for some multi-join queries, but upstream warns that its join order is not necessarily optimal. Cross joins and disconnected relation sets require special handling; beyond configured thresholds it falls back to PostgreSQL dynamic programming or GEQO. Other controls include `pg_dphyp.count_cc` and `pg_dphyp.geqo_cc_threshold`.

Planner hooks and internal structures are major-version-sensitive. The reviewed source claims PostgreSQL 12 through 17 compatibility but also says there is no real test suite and that plan-output tests are version-specific. Benchmark planning time and execution time separately, compare plans with the module disabled, and test outer joins, lateral references, disconnected graphs, prepared statements, partitioning, and upgrades before enabling it for a production cluster.
