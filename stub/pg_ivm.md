## Usage

Sources: [README](https://github.com/sraoss/pg_ivm/blob/master/README.md), [release 1.14](https://github.com/sraoss/pg_ivm/releases/tag/v1.14)

`pg_ivm` provides immediate Incremental View Maintenance for PostgreSQL materialized views. Instead of recomputing the whole view, it applies deltas in `AFTER` triggers and stores metadata in the `pgivm` schema.

```sql
CREATE EXTENSION pg_ivm;
```

### Required Setup

Upstream says `pg_ivm` should be preloaded so IMMVs are maintained correctly:

```conf
shared_preload_libraries = 'pg_ivm'
session_preload_libraries = 'pg_ivm'
```

The current README says the extension is compatible with PostgreSQL 13 through 18, and the latest GitHub release is `1.14` dated March 31, 2026.

### Main Functions

- `pgivm.create_immv(name, query)` creates an incrementally maintainable materialized view (IMMV), its maintenance triggers, and a unique index when possible.
- `pgivm.refresh_immv(name, with_data)` fully refreshes the IMMV and can disable or re-enable maintenance.
- `pgivm.get_immv_def(regclass)` reconstructs the stored `SELECT` definition.
- `pgivm.pg_ivm_immv` stores IMMV metadata including `immvrelid`, `viewdef`, `ispopulated`, and `lastivmupdate`.

### Common Patterns

Create an IMMV:

```sql
SELECT pgivm.create_immv(
  'immv_agg',
  'SELECT bid, count(*), sum(abalance), avg(abalance)
   FROM pgbench_accounts JOIN pgbench_branches USING(bid)
   GROUP BY bid'
);
```

Query the maintained result after base-table changes:

```sql
UPDATE pgbench_accounts SET abalance = abalance + 1000 WHERE aid = 4112345;
SELECT * FROM immv_agg WHERE bid = 42;
```

Inspect or refresh IMMVs:

```sql
SELECT immvrelid AS immv, pgivm.get_immv_def(immvrelid)
FROM pgivm.pg_ivm_immv;

SELECT pgivm.refresh_immv('immv_agg', true);
```

Pause maintenance for bulk work, then rebuild:

```sql
SELECT pgivm.refresh_immv('myview', false);
-- bulk changes
SELECT pgivm.refresh_immv('myview', true);
```

### Caveats

- Upstream only supports a restricted subset of view definitions: joins, `DISTINCT`, simple subqueries/CTEs, and built-in aggregates `count`, `sum`, `avg`, `min`, and `max`.
- Unsupported constructs include `HAVING`, window functions, `ORDER BY`, `LIMIT/OFFSET`, `UNION`/`INTERSECT`/`EXCEPT`, `DISTINCT ON`, and user-defined aggregates.
- Efficient maintenance depends on having a suitable unique index; `create_immv` creates one automatically only when the definition allows it.
