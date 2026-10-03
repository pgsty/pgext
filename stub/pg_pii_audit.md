## Usage

Sources:

- [Official README.md](https://github.com/jgsn13/piipg/blob/661ae53fa2a110cc535cf9b33a15d9f8ce28f026/README.md)
- [Official pg_pii_audit.control](https://github.com/jgsn13/piipg/blob/661ae53fa2a110cc535cf9b33a15d9f8ce28f026/pg_pii_audit.control)
- [Official pg_pii_audit--1.0.sql](https://github.com/jgsn13/piipg/blob/661ae53fa2a110cc535cf9b33a15d9f8ce28f026/sql/pg_pii_audit--1.0.sql)
- [Official pg_pii_audit.c](https://github.com/jgsn13/piipg/blob/661ae53fa2a110cc535cf9b33a15d9f8ce28f026/src/pg_pii_audit.c)
- [Official Dockerfile](https://github.com/jgsn13/piipg/blob/661ae53fa2a110cc535cf9b33a15d9f8ce28f026/Dockerfile)

`pg_pii_audit` 1.0 is an academic PostgreSQL 16 prototype that samples email-shaped literal values in INSERT execution plans. It records counters, not a comprehensive classification of stored personal data.

### Core Workflow

Install as a superuser, then load the hook in each session to be observed:

```sql
CREATE EXTENSION pg_pii_audit;
LOAD 'pg_pii_audit';

CREATE TABLE pii_demo (email text);
INSERT INTO pii_demo VALUES ('alice@example.org');

SELECT table_name, column_name, total_samples,
       matched_samples, column_probability
FROM pii_audit.identified_table_columns;
```

The control file fixes the schema to `pii_audit`. The SQL script creates `identified_table_columns`; it does not itself call a C function that loads the executor hook, so the explicit `LOAD` is essential.

### Interpretation and Limits

`total_samples` counts inspected plan constants; `matched_samples` counts regex matches. `column_probability` is their ratio. `table_probability` is updated with the same per-column ratio and is not an independent table-level estimate.

Only literal text, varchar and char expressions reached by the hook are inspected. Multi-row VALUES, parameters, INSERT SELECT, existing data and non-INSERT changes must not be assumed covered. Audit writes execute with the caller’s privileges; missing permissions on the audit table or sequence can abort the original write.

Use disposable evaluation data. Table and column names are interpolated into SQL without escaping, and the recursion guard is not reset on every error path. There are no documented configuration GUCs, broader PostgreSQL compatibility guarantees, or declared license in this revision.
