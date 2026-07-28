## Usage

Sources:

- [Official upstream README](https://github.com/scokmen/pg_mask/blob/e614f6ebe6ac98f89c977694cd4624c978c1caf1/README.md)
- [Official extension control file (pg_mask.control)](https://github.com/scokmen/pg_mask/blob/e614f6ebe6ac98f89c977694cd4624c978c1caf1/pg_mask.control)
- [Official extension SQL (pg_mask--1.0.0.sql)](https://github.com/scokmen/pg_mask/blob/e614f6ebe6ac98f89c977694cd4624c978c1caf1/pg_mask--1.0.0.sql)

`pg_mask` — This project is being developed with C language using C-Language Functions of PostgreSQL, that is also provide an extension makefile that manages library directories and installation targets. In order to include PostgreSQL extension makefile, the binary must be available. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_mask;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_mask()` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
