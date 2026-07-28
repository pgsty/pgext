## Usage

Sources:

- [Official upstream README](https://github.com/darthunix/pg_vx/blob/5916aafcfde5b7a5c5a90b6ca7ea1879faf85a4d/README.md)
- [Official extension control file (pg_vx.control)](https://github.com/darthunix/pg_vx/blob/5916aafcfde5b7a5c5a90b6ca7ea1879faf85a4d/pg_vx.control)
- [Official extension SQL (pg_vx--0.1.sql)](https://github.com/darthunix/pg_vx/blob/5916aafcfde5b7a5c5a90b6ca7ea1879faf85a4d/pg_vx--0.1.sql)

`pg_vx` — PG_VX means PostgreSQL Vectorized Executor. It is a test project to understand how PostgreSQL CustomScanAPI works and is it possible to use it for vectorized acceleration of scans, aggregates and joins on OLAP workloads. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_vx;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
