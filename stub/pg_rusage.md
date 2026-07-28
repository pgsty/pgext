## Usage

Sources:

- [Official upstream README](https://github.com/michaelpq/pg_plugins/blob/626fb56b0a0b833d4f23ca55359ce56d38162864/pg_rusage/README)
- [Official extension control file (pg_rusage.control)](https://github.com/michaelpq/pg_plugins/blob/626fb56b0a0b833d4f23ca55359ce56d38162864/pg_rusage/pg_rusage.control)
- [Official extension SQL (pg_rusage--1.0.sql)](https://github.com/michaelpq/pg_plugins/blob/626fb56b0a0b833d4f23ca55359ce56d38162864/pg_rusage/pg_rusage--1.0.sql)

`pg_rusage` — This module is a PostgreSQL extension that can enable CPU measurements, with one SQL function to enable the measurement and one to disable it. When disabling, the existing accumulated results will show up. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_rusage;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_rusage_print()` is an extension function and returns `void`.
- `pg_rusage_reset()` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
