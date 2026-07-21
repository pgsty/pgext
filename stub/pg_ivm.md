## Usage

Sources:

- [Official v1.15 README](https://github.com/sraoss/pg_ivm/blob/v1.15/README.md)
- [v1.15 release notes](https://github.com/sraoss/pg_ivm/releases/tag/v1.15)
- [v1.14 to v1.15 upgrade SQL](https://github.com/sraoss/pg_ivm/blob/v1.15/pg_ivm--1.14--1.15.sql)
- [pg_ivm_dump_metadata utility](https://github.com/sraoss/pg_ivm/blob/v1.15/scripts/pg_ivm_dump_metadata)

`pg_ivm` provides immediate incremental view maintenance for PostgreSQL. An Incrementally Maintainable Materialized View (IMMV) is stored as a table with triggers and metadata in the `pgivm` schema; base-table changes update the IMMV inside the same transaction instead of recomputing the complete query.

### Enable and Create an IMMV

Load the library for every session that can modify an IMMV's base tables. A cluster-wide setup requires a restart:

```conf
shared_preload_libraries = 'pg_ivm'
```

`session_preload_libraries = 'pg_ivm'` is also supported when managed consistently for all relevant sessions.

```sql
CREATE EXTENSION pg_ivm;

SELECT pgivm.create_immv(
    'account_totals',
    'SELECT branch_id, count(*) AS accounts, sum(balance) AS balance
     FROM accounts
     GROUP BY branch_id'
);

UPDATE accounts
SET balance = balance + 100
WHERE account_id = 42;

SELECT * FROM account_totals;
```

### Manage and Inspect IMMVs

- `pgivm.create_immv(name, query)`: creates and populates an IMMV, returning its row count.
- `pgivm.refresh_immv(name, with_data)`: fully rebuilds the IMMV; `false` disables maintenance until a later populated refresh.
- `pgivm.get_immv_def(regclass)`: returns the stored view definition.
- `pgivm.restore_immv(name, query, populate)`: version 1.15 function that reconstructs metadata, triggers, and indexes for an existing IMMV table.
- `pgivm.get_create_immv_commands()` and `pgivm.get_restore_immv_commands()`: emit SQL for rebuilding IMMVs or restoring their metadata.

Version 1.15 includes a helper for dump or `pg_upgrade` workflows:

```shell
pg_ivm_dump_metadata -d application > pg_ivm_metadata.sql
```

The script emits `pgivm.restore_immv()` calls. Restore the table data first, then execute the saved metadata SQL so incremental maintenance resumes without recreating the tables.

### Restrictions and Operational Caveats

- Supported definitions include selected joins, `DISTINCT`, simple subqueries/CTEs, and built-in `count`, `sum`, `avg`, `min`, and `max` aggregates. Unsupported constructs include `HAVING`, window functions, `ORDER BY`, `LIMIT/OFFSET`, set operations, `DISTINCT ON`, and user-defined aggregates.
- Efficient maintenance depends on a suitable unique index. `create_immv()` creates one automatically only when the definition supplies usable grouping, distinct, or base-table primary-key columns.
- Creation and refresh take `AccessExclusiveLock`. Upstream warns about consistency risks for creation under `REPEATABLE READ` or `SERIALIZABLE`; use `READ COMMITTED` or refresh afterward.
- `restore_immv()` fails when the relation is already registered or its table definition does not match the supplied query.
- Version 1.15 also fixes incorrect maintenance after repeated trigger-driven modifications and a v1.14 outer-join maintenance crash.
