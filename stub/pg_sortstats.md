## Usage

Sources:

- [Official README](https://github.com/powa-team/pg_sortstats/blob/b9228cca5947b010c8114ecd5091e76d57ef28c3/README.md)
- [Extension SQL API](https://github.com/powa-team/pg_sortstats/blob/b9228cca5947b010c8114ecd5091e76d57ef28c3/pg_sortstats--0.0.1.sql)
- [Shared-memory and GUC implementation](https://github.com/powa-team/pg_sortstats/blob/b9228cca5947b010c8114ecd5091e76d57ef28c3/pg_sortstats.c)

`pg_sortstats` 0.0.1 collects cumulative sort-node statistics keyed by query ID, user, database, and textual sort keys. It distinguishes top-N heapsort, quicksort, external sort/merge, disk and memory space, parallel workers, and estimates the `work_mem` needed to keep observed sorts in memory. Upstream marks it under development and not production-ready.

### Preload and Configure

Both dependencies require startup loading and a restart:

```conf
shared_preload_libraries = 'pg_stat_statements,pg_sortstats'
pg_sortstats.max = 10000
pg_sortstats.save = on
```

Then install it in databases where the reporting objects are needed:

```sql
CREATE EXTENSION pg_sortstats;

SELECT queryid, userid, dbid, sort_keys,
       external_sorts, space_disk, work_mems
FROM pg_sortstats
ORDER BY space_disk DESC;

SELECT * FROM pg_sortstats(false);

SELECT pg_sortstats_reset();
```

`pg_sortstats` is a view over `pg_sortstats(true)`. Passing false omits textual sort keys when their disclosure or collection cost is unwanted. Counters are cumulative totals, not per-execution values; relate `queryid` to `pg_stat_statements` and compare changes over a measured interval.

### GUCs and Report Fields

`pg_sortstats.enabled` is a user-settable switch and defaults on. `pg_sortstats.max` is the postmaster-time maximum number of tracked entries, defaults to 10,000, and affects shared memory. `pg_sortstats.save` defaults on and controls saving statistics across clean shutdown/restart.

Besides identity and `sort_keys`, the function reports `nb_keys`, input `lines`, `lines_to_sort`, estimated `work_mems`, algorithm counts, `nb_tapes`, `space_disk`, `space_memory`, `non_parallels`, and `nb_workers`.

### Operational Boundaries

Collection adds executor-hook and shared-memory overhead. Sort-key text and cross-user query identifiers can expose workload structure, while `pg_sortstats_reset()` clears shared statistics; revoke reporting/reset execution from roles that should not have those capabilities. Query IDs require `pg_stat_statements` and suitable query-ID computation.

Use the estimate as tuning evidence, not as a direct cluster-wide `work_mem` setting: increasing `work_mem` multiplies across concurrent operations and sessions. Validate extension compatibility and counter accuracy on the exact PostgreSQL major before deployment.
