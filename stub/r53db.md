## Usage

Sources:

- [Official upstream README](https://github.com/apparentorder/r53db/blob/8c66a3f954918c2d292700b7494f5315506b169d/README.md)
- [Official extension control file (r53db.control)](https://github.com/apparentorder/r53db/blob/8c66a3f954918c2d292700b7494f5315506b169d/r53db.control)
- [Official extension SQL (r53db--0.1.sql)](https://github.com/apparentorder/r53db/blob/8c66a3f954918c2d292700b7494f5315506b169d/r53db--0.1.sql)

`r53db` — *r53db* is a Foreign Data Wrapper for PostgreSQL that enables you to access AWS Route53 Database Zones like SQL tables. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION r53db;
CREATE SERVER route53 FOREIGN DATA WRAPPER r53db;
CREATE SCHEMA route53;
IMPORT FOREIGN SCHEMA dummy FROM SERVER route53 INTO route53;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `r53db_fdw_handler()` is an extension function and returns `fdw_handler`.
- `r53db` is an extension-defined foreign data wrapper.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
