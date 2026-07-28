## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_cext/pg_cext-1.0.0/README.md)
- [Official extension control file (pg_cext.control)](https://api.pgxn.org/src/pg_cext/pg_cext-1.0.0/pg_cext.control)
- [Official extension SQL (pg_cext--1.0.0.sql)](https://api.pgxn.org/src/pg_cext/pg_cext-1.0.0/pg_cext--1.0.0.sql)

`pg_cext` — This is a PostgreSQL extension implemented in C that adds two numbers. It demonstrates how to create a simple extension for PostgreSQL using the C language. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_cext;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `add_nums(int,int)` is an extension function and returns `int`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
