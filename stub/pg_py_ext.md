## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_py_ext/pg_py_ext-1.0.0/README.md)
- [Official extension control file (pg_py_ext.control)](https://api.pgxn.org/src/pg_py_ext/pg_py_ext-1.0.0/pg_py_ext.control)
- [Official extension SQL (pg_py_ext--1.0.0.sql)](https://api.pgxn.org/src/pg_py_ext/pg_py_ext-1.0.0/pg_py_ext--1.0.0.sql)

`pg_py_ext` — **A PostgreSQL extension using PL/Python3U to add numbers**. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_py_ext;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `add_numbers(a integer, b integer)` is an extension function and returns `integer`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
