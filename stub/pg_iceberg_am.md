## Usage

Sources:

- [Official upstream README](https://github.com/robertmu/pg-lakebase/blob/a5baec33934b069b0832644ff9cee64b429c14cb/pg-iceberg-am/README.md)
- [Official extension control file (pg_iceberg_am.control)](https://github.com/robertmu/pg-lakebase/blob/a5baec33934b069b0832644ff9cee64b429c14cb/pg-iceberg-am/pg_iceberg_am.control)
- [Official implementation source](https://github.com/robertmu/pg-lakebase/blob/a5baec33934b069b0832644ff9cee64b429c14cb/pg-iceberg-am/src/lib.rs)

`pg_iceberg_am` — pg-iceberg-am is the current SQL-facing extension in pg-lakebase. It exposes Apache Iceberg tables through PostgreSQL's **Table Access Method (TAM)** interface, so applications use ordinary SQL and PostgreSQL transaction semantics while metadata and data files are managed as Iceberg and Parquet. Use it for the corresponding analytical or storage workflow. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION pg_iceberg_am;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- Install the confirmed extension dependencies first: `pg_lakebase_runtime`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
