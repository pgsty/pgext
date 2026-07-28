## Usage

Sources:

- [Official upstream README](https://github.com/ansilo-data/ansilo/blob/819a32f1782c4d8d4c97e01fe908e2694a546f35/README.md)
- [Official extension control file (ansilo_pgx.control)](https://github.com/ansilo-data/ansilo/blob/819a32f1782c4d8d4c97e01fe908e2694a546f35/ansilo-pgx/ansilo_pgx.control)
- [Official implementation source](https://github.com/ansilo-data/ansilo/blob/819a32f1782c4d8d4c97e01fe908e2694a546f35/ansilo-pgx/src/lib.rs)

`ansilo_pgx` — A Postgres interface into any database. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION ansilo_pgx;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hello_ansilo()` is an extension function.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
