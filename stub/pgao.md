## Usage

Sources:

- [Official upstream README](https://github.com/zvdy/pgao/blob/45f09972a1d6e551125d95279c0fd863a5533aa3/extension/README.md)
- [Official extension control file (pgao.control)](https://github.com/zvdy/pgao/blob/45f09972a1d6e551125d95279c0fd863a5533aa3/extension/pgao.control)
- [Official extension SQL (pgao--0.1.0.sql)](https://github.com/zvdy/pgao/blob/45f09972a1d6e551125d95279c0fd863a5533aa3/extension/pgao--0.1.0.sql)

`pgao` — All functions are SQL-only, STABLE/IMMUTABLE, and require no superuser. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgao;
SELECT * FROM pgao.health();
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgao.health()` is an extension function and returns `TABLE`.
- `pgao.replication_lag_ms()` is an extension function and returns `bigint`.
- `pgao.table_bloat()` is an extension function and returns `TABLE`.
- `pgao.version()` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
