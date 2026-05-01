## Usage

Sources: [README](https://github.com/viggy28/pg_savior/blob/v0.1.0/README.md), [release 0.1.0](https://github.com/viggy28/pg_savior/releases/tag/v0.1.0), [PGXN 0.1.0](https://pgxn.org/dist/pg_savior/0.1.0/), [SQL file](https://github.com/viggy28/pg_savior/blob/v0.1.0/pg_savior--0.1.0.sql), [C source](https://github.com/viggy28/pg_savior/blob/v0.1.0/pg_savior.c), [pg_savior.control](https://github.com/viggy28/pg_savior/blob/v0.1.0/pg_savior.control)

`pg_savior` is a PostgreSQL safety extension for blocking specific high-risk DML and DDL statements before they run. Version `0.1.0` is a deliberate PGXN release and a major rewrite from `0.0.1`; the README still labels it pre-1.0 and not production-ready.

### Activation

`CREATE EXTENSION` alone does not activate the checks. The SQL file only documents that protection lives in the loaded shared library, so each backend must load `pg_savior` by one of the upstream-supported paths:

Cluster-wide activation uses `shared_preload_libraries` and requires a PostgreSQL restart:

```conf
shared_preload_libraries = 'pg_savior'
```

Per-session activation for new connections can use `session_preload_libraries` after a config reload:

```conf
session_preload_libraries = 'pg_savior'
```

For development or test sessions, load the library manually:

```sql
LOAD 'pg_savior';
CREATE EXTENSION pg_savior;
```

Once the library is loaded, `_PG_init` installs `post_parse_analyze_hook`, `ExecutorStart_hook`, and `ProcessUtility_hook` for that backend.

### DML Guards

`pg_savior` blocks `DELETE` and `UPDATE` statements that have no `WHERE` clause. The parser hook checks the analyzed `Query` tree and raises `ERROR`, so the transaction aborts and the application sees the failure.

```sql
CREATE TABLE emp (id int);
INSERT INTO emp VALUES (1), (2), (3);

DELETE FROM emp;
-- ERROR: pg_savior: DELETE without WHERE clause is blocked

UPDATE emp SET id = id + 1;
-- ERROR: pg_savior: UPDATE without WHERE clause is blocked

DELETE FROM emp WHERE id = 1;
-- allowed
```

The optional row-count guard applies to `DELETE` and `UPDATE` statements whose planner estimate exceeds `pg_savior.max_rows_affected`. It runs from `ExecutorStart_hook`, after planning and before tuples are touched.

```sql
SET pg_savior.max_rows_affected = 100;

DELETE FROM emp WHERE id > 0;
-- blocked if the planner estimate is greater than 100 rows
```

### DDL Guards

The `ProcessUtility_hook` guards only the DDL operations listed by upstream:

- `CREATE INDEX` without `CONCURRENTLY` is always blocked.
- `DROP DATABASE` is always blocked.
- `ALTER TABLE ADD COLUMN ... DEFAULT` is blocked when the target table is larger than `pg_savior.large_table_threshold_rows`.
- `ALTER TABLE ALTER COLUMN TYPE` is blocked for large tables.
- `TRUNCATE` is blocked when any target table is large.
- `DROP TABLE` is blocked when any target table is large.

Large-table checks use `pg_class.reltuples > pg_savior.large_table_threshold_rows`.

```sql
CREATE INDEX emp_idx ON emp (id);
-- ERROR: pg_savior: CREATE INDEX without CONCURRENTLY is blocked

CREATE INDEX CONCURRENTLY emp_idx ON emp (id);
-- allowed by this guard

ALTER TABLE big_emp ADD COLUMN status text DEFAULT 'active';
-- blocked when big_emp is over pg_savior.large_table_threshold_rows

TRUNCATE big_emp;
-- blocked when big_emp is over pg_savior.large_table_threshold_rows
```

### Configuration

All documented GUCs are session-scoped `USERSET` variables:

| GUC | Default | Effect |
|---|---:|---|
| `pg_savior.enabled` | `on` | Master switch; when `off`, checks do not run. |
| `pg_savior.bypass` | `off` | Allows the current session through the guards. |
| `pg_savior.max_rows_affected` | `0` | Blocks estimated `DELETE`/`UPDATE` row counts above this value; `0` disables the check. |
| `pg_savior.large_table_threshold_rows` | `1000000` | Defines "large" for the guarded large-table DDL operations. |

Use `SET LOCAL` for a deliberate one-transaction bypass:

```sql
BEGIN;
SET LOCAL pg_savior.bypass = on;
DELETE FROM staging_table;
COMMIT;
```

### Caveats

- The library must be loaded in the backend before protection exists; `CREATE EXTENSION pg_savior` only registers extension metadata.
- The row-count and large-table guards depend on planner/catalog estimates. Run `ANALYZE` when recent changes make estimates stale.
- `UPDATE` coverage is limited to unguarded `UPDATE` and the optional planner-estimate threshold; the README does not claim semantic review of every `WHERE` predicate.
- DDL coverage is limited to the listed `ProcessUtility_hook` cases. Do not assume other schema changes are blocked.
- The `ADD COLUMN ... DEFAULT` guard is conservative and blocks any default on a large table, including non-volatile defaults that newer PostgreSQL versions may handle without a full table rewrite.
