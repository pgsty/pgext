## Usage

Sources:

- [Official EDB Postgres Tuner documentation](https://www.enterprisedb.com/docs/pg_extensions/pg_tuner/)
- [Official configuration guide](https://www.enterprisedb.com/docs/pg_extensions/pg_tuner/configuring/)
- [Official usage guide](https://www.enterprisedb.com/docs/pg_extensions/pg_tuner/using/)

`edb_pg_tuner` observes host resources and workload behavior, then recommends or optionally applies PostgreSQL configuration changes.

### Enablement

Install EDB Postgres Tuner 1.3.2, preload it, restart, and create the extension:

```ini
shared_preload_libraries = 'edb_pg_tuner'
edb_pg_tuner.autotune = false
```

```sql
CREATE EXTENSION edb_pg_tuner;
```

Keep `edb_pg_tuner.autotune` disabled until recommendations have been reviewed against workload and memory budgets.

### Review Recommendations

`edb_pg_tuner_recommendations` returns current recommendations or SQL commands. PostgreSQL 14+ also exposes query-level spill statistics.

```sql
SELECT * FROM edb_pg_tuner_recommendations();
SELECT edb_pg_tuner_recommendations('sql');

SELECT *
FROM edb_pg_tuner_query_stats()
WHERE sort_spill > 0 OR hash_spill > 0;
```

`edb_pg_tuner_global_stats()` reports buffer-level tracking totals.

### Operational Boundaries

Automatic tuning can change restart-required GUCs and can raise `work_mem` based on prior spill observations. `work_mem` applies per plan node and per concurrent query, so the observed spill multiplier is not a cluster memory guarantee. Bound `edb_pg_tuner.work_mem_pool` and `edb_pg_tuner.max_wal_size_limit`, review generated SQL, and roll changes through the same change-control path as manual PostgreSQL configuration.

