## Usage

Sources:

- [Official upstream README](https://github.com/olirice/pg_spi/blob/307212d23d931775d122c1559bd8be041a97c67e/README.md)
- [Official extension control file (pg_spi.control)](https://github.com/olirice/pg_spi/blob/307212d23d931775d122c1559bd8be041a97c67e/pg_spi.control)
- [Official extension SQL (pg_spi--0.0.1.sql)](https://github.com/olirice/pg_spi/blob/307212d23d931775d122c1559bd8be041a97c67e/pg_spi--0.0.1.sql)

`pg_spi` — Failing test case for SPI rollback. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_spi;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `execute_and_commit(text)` is an extension function and returns `bigint`.
- `execute_and_rollback(text)` is an extension function and returns `bigint`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
