## Usage

Sources:

- [Official upstream README](https://github.com/robertmu/pg-lakebase/blob/a5baec33934b069b0832644ff9cee64b429c14cb/README.md)
- [Official extension control file (pg_delta_am.control)](https://github.com/robertmu/pg-lakebase/blob/a5baec33934b069b0832644ff9cee64b429c14cb/pg-delta-am/pg_delta_am.control)
- [Official implementation source](https://github.com/robertmu/pg-lakebase/blob/a5baec33934b069b0832644ff9cee64b429c14cb/pg-delta-am/src/lib.rs)

`pg_delta_am` — Delta table access method skeleton. Use it for the corresponding analytical or storage workflow. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION pg_delta_am;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `delta_table_am_handler()` is an extension function and returns `pg_lakebase_core::TableAmRoutine`.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- Install the confirmed extension dependencies first: `pg_lakebase_runtime`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
