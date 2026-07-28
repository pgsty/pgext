## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/logfmt/logfmt-1.0.0/README.md)
- [Official extension control file (logfmt.control)](https://api.pgxn.org/src/logfmt/logfmt-1.0.0/logfmt.control)
- [Official extension SQL (logfmt--1.0.0.sql)](https://api.pgxn.org/src/logfmt/logfmt-1.0.0/logfmt--1.0.0.sql)

`logfmt` — logfmt is known to work with PostgreSQL 16beta1. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION logfmt;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `emit_test_logs()` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
