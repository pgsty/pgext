## Usage

Sources:

- [Official upstream README](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/README.md)
- [Official extension control file (gp_exttable_fdw.control)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_exttable_fdw/gp_exttable_fdw.control)
- [Official extension SQL (gp_exttable_fdw--1.0.sql)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_exttable_fdw/gp_exttable_fdw--1.0.sql)

`gp_exttable_fdw` — External Table Foreign Data Wrapper for Greenplum-family databases. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION gp_exttable_fdw;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gp_exttable_fdw_handler()` is an extension function and returns `fdw_handler`.
- `gp_exttable_permission_check(text[], oid)` is an extension function and returns `void`.
- `pg_exttable` is an extension-defined view.
- `gp_exttable_fdw` is an extension-defined foreign data wrapper.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
