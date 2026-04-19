
## Usage

> Sources: [README v0.20.0](https://github.com/grove/pg-trickle/blob/v0.20.0/README.md), [v0.20.0 release notes](https://github.com/grove/pg-trickle/releases/tag/v0.20.0)

`pg_trickle` provides incrementally maintained stream tables for PostgreSQL 18. It keeps query results fresh using differential refresh when possible, with full recompute and immediate in-transaction modes also documented upstream.

### Enable the Extension

```sql
-- postgresql.conf
shared_preload_libraries = 'pg_trickle'
max_worker_processes = 8

CREATE EXTENSION pg_trickle;
```

The upstream README notes that `wal_level = logical` is not required by default. CDC starts with row-level triggers and can switch to WAL-based capture when `pg_trickle.cdc_mode = 'auto'`.

### Create and Refresh Stream Tables

```sql
SELECT pgtrickle.create_stream_table(
    'regional_totals',
    'SELECT region, SUM(amount) AS total, COUNT(*) AS cnt
     FROM orders GROUP BY region'
);

SELECT * FROM regional_totals;
SELECT pgtrickle.refresh_stream_table('regional_totals');
```

The documented refresh modes are `AUTO`, `DIFFERENTIAL`, `FULL`, and `IMMEDIATE`.

```sql
SELECT pgtrickle.create_stream_table(
    'regional_totals_live',
    'SELECT region, SUM(amount) AS total, COUNT(*) AS cnt
     FROM orders GROUP BY region',
    schedule => NULL,
    refresh_mode => 'IMMEDIATE'
);
```

### Operations and Introspection

```sql
SELECT pgtrickle.alter_stream_table(
    'regional_totals',
    query => 'SELECT region, SUM(amount) AS total FROM orders GROUP BY region'
);

SELECT * FROM pgtrickle.pgt_status();
SELECT * FROM pgtrickle.health_check();
SELECT * FROM pgtrickle.refresh_timeline(20);
SELECT * FROM pgtrickle.change_buffer_sizes();
SELECT * FROM pgtrickle.dependency_tree();
```

The README also documents broad SQL coverage including joins, aggregates, window functions, recursive CTEs, subqueries, set operations, and TopK queries.

### v0.20.0 Monitoring Additions

The `v0.20.0` release adds built-in self-monitoring:

- `pgtrickle.setup_dog_feeding()`
- `pgtrickle.teardown_dog_feeding()`
- `pgtrickle.dog_feeding_status()`

Release notes describe five monitoring stream tables that analyze refresh history and can emit threshold advice and alerts.
