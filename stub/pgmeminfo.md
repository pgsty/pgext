## Usage

Sources:

- [1.0.1 README](https://github.com/okbob/pgmeminfo/blob/VERSION_1_0_1/README.md)
- [Installation SQL](https://github.com/okbob/pgmeminfo/blob/VERSION_1_0_1/pgmeminfo--1.0.sql)
- [Control file](https://github.com/okbob/pgmeminfo/blob/VERSION_1_0_1/pgmeminfo.control)

`pgmeminfo` reports allocator statistics and the memory-context hierarchy of the current PostgreSQL backend. It does not aggregate memory across the cluster.

### Inspect Memory

Have a superuser install the extension, then inspect the current connection:

```sql
CREATE EXTENSION pgmeminfo;
SELECT * FROM pgmeminfo();
SELECT * FROM pgmeminfo_contexts();
SELECT * FROM pgmeminfo_contexts(deep => 1);
SELECT * FROM pgmeminfo_contexts(deep => -1, accum_mode => 'off');
```

`pgmeminfo()` returns allocator counters such as `arena`, `uordblks`, `fordblks`, and `keepcost`. They are allocator-specific measurements, not process RSS or an estimate of free cluster memory.

`pgmeminfo_contexts(deep, accum_mode)` returns context names, parents, levels and byte counts. The default accumulation mode is `all`; `off` reports contexts without descendant accumulation, while `deep => -1` removes the depth limit.

### Operation and Version

No preload or restart is required. Upstream release 1.0.1 keeps SQL extension version 1.0; do not use package release 1.0.1 as an SQL update target. Context names and output values can change as the backend executes queries. Review function access if internal context names should not be visible to ordinary users.
