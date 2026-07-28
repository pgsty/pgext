## Usage

Sources:

- [Official upstream README](https://github.com/ohdsi/trex/blob/c9e2b35ff8ebaf2320b060c9d3d32f0007045511/README.md)
- [Official extension control file (pg_trex.control)](https://github.com/ohdsi/trex/blob/c9e2b35ff8ebaf2320b060c9d3d32f0007045511/plugins/pg_trex/pg_trex.control)
- [Official extension SQL (pg_trex--0.1.0.sql)](https://github.com/ohdsi/trex/blob/c9e2b35ff8ebaf2320b060c9d3d32f0007045511/plugins/pg_trex/sql/pg_trex--0.1.0.sql)

`pg_trex` — pg trex: PostgreSQL extension for distributed trexsql. Use it for the corresponding analytical or storage workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_trex;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
