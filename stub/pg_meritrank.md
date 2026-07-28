## Usage

Sources:

- [Official upstream README](https://github.com/vsradkevich/pg_meritrank/blob/6157eac1d22b52357bd84c5c6b0dac76eb33160a/README.md)
- [Official extension control file (pg_meritrank.control)](https://github.com/vsradkevich/pg_meritrank/blob/6157eac1d22b52357bd84c5c6b0dac76eb33160a/pg_meritrank.control)
- [Official implementation source](https://github.com/vsradkevich/pg_meritrank/blob/6157eac1d22b52357bd84c5c6b0dac76eb33160a/src/lib.rs)

`pg_meritrank` — Postgres Merit Rank is an extension for PostgreSQL that provides functionality for calculating and ranking merits. This README provides instructions for testing the extension using cargo pgx test and installing it in a PostgreSQL database. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_meritrank;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.0.1`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
