## Usage

Sources:

- [Official README](https://github.com/asotolongo/vacuum_utils/blob/8de352f1960ad0d32c9d0d8cf24fa7417ef6c057/README.md)
- [Version 0.1.0 extension SQL](https://github.com/asotolongo/vacuum_utils/blob/8de352f1960ad0d32c9d0d8cf24fa7417ef6c057/vacuum_utils--0.1.0.sql)
- [PGXN distribution](https://pgxn.org/dist/vacuum_utils/)

`vacuum_utils` provides SQL helpers for estimating autovacuum/analyze thresholds, reading the last maintenance time and count, toggling per-table autovacuum, and running `ANALYZE`. The canonical function name `diff_to_autovaccum` intentionally retains upstream's misspelling.

### Core Workflow

```sql
CREATE EXTENSION vacuum_utils;

SELECT vacuum_utils.get_table_vacuum_threshold('public', 'customers');
SELECT vacuum_utils.diff_to_autovaccum('public', 'customers');
SELECT * FROM vacuum_utils.last_autovacuum_count('public', 'customers');

SELECT vacuum_utils.execute_analyze('public', 'customers');
```

The inspection helpers read `pg_stat_user_tables`, `pg_class.reltuples`, and the current global autovacuum settings. The mutating helpers execute `ALTER TABLE` or `ANALYZE` with the caller's privileges.

### Object Index

- `diff_to_autovaccum(schema, table)` estimates remaining dead tuples before the global vacuum threshold.
- `get_table_vacuum_threshold` and `get_table_analyze_threshold` calculate global-setting-based thresholds.
- `last_vacuum_count`, `last_autovacuum_count`, `last_analyze_count`, and `last_autoanalyze_count` return a `datetext_and_int` composite.
- `disable_autovacuum` and `enable_autovacuum` change the table's `autovacuum_enabled` reloption.
- `execute_analyze` runs `ANALYZE` for the named table.

### Operational Notes

The upstream README explicitly says version `0.1.0` has known bugs. Although the control file says relocatable, the SQL script creates and uses a fixed `vacuum_utils` schema. The threshold calculations use current global GUCs and do not account for per-table storage parameters, inserts-based triggers, wraparound vacuum, partition behavior, or planner-estimate staleness; treat results as rough diagnostics only.

Do not grant the mutating functions to untrusted roles. Their dynamic SQL concatenates schema and table arguments without identifier quoting, creating syntax and SQL-injection risk, and the functions run with invoker privileges. Use only administrator-supplied simple identifiers, confirm the target table separately, and prefer native catalog queries plus safely quoted `ALTER TABLE`/`ANALYZE` statements in production automation.
