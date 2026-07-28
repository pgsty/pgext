## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_rand_ext/pg_rand_ext-1.0.1/README.md)
- [Official extension control file (pg_rand_ext.control)](https://api.pgxn.org/src/pg_rand_ext/pg_rand_ext-1.0.1/pg_rand_ext.control)
- [Official extension SQL (pg_rand_ext--1.0.sql)](https://api.pgxn.org/src/pg_rand_ext/pg_rand_ext-1.0.1/pg_rand_ext--1.0.sql)

`pg_rand_ext` — The common way to build extension modules in PostgreSQL. It has been tested with PostgreSQL v14 and above. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_rand_ext;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `rand_ext.random_exponential(bigint, bigint, double precision)` is an extension function and returns `bigint`.
- `rand_ext.random_gaussian(bigint, bigint, double precision)` is an extension function and returns `bigint`.
- `rand_ext.random_zipfian(bigint, bigint, double precision)` is an extension function and returns `bigint`.
- `rand_ext` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
