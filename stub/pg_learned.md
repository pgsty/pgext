## Usage

Sources:

- [Official upstream README](https://github.com/baofuhann/pg-learned/blob/d1433bf3261d1adee216c2a0162081b21a0be59f/README.md)
- [Official extension control file (pg_learned.control)](https://github.com/baofuhann/pg-learned/blob/d1433bf3261d1adee216c2a0162081b21a0be59f/pg_learned.control)
- [Official implementation source](https://github.com/baofuhann/pg-learned/blob/d1433bf3261d1adee216c2a0162081b21a0be59f/c_impl/pg_learned.c)

`pg_learned` — **A PostgreSQL extension demonstrating learned index technology for faster database queries.**. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_learned;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
