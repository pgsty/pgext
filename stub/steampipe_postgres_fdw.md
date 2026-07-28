## Usage

Sources:

- [Official upstream README](https://github.com/turbot/steampipe-postgres-fdw/blob/f8098e5e79cda44af01fcb9cd77b9ac684e70690/fdw/README.md)
- [Official extension control file (steampipe_postgres_fdw.control)](https://github.com/turbot/steampipe-postgres-fdw/blob/f8098e5e79cda44af01fcb9cd77b9ac684e70690/fdw/steampipe_postgres_fdw.control)
- [Official extension SQL (steampipe_postgres_fdw--1.0.sql)](https://github.com/turbot/steampipe-postgres-fdw/blob/f8098e5e79cda44af01fcb9cd77b9ac684e70690/fdw/steampipe_postgres_fdw--1.0.sql)

`steampipe_postgres_fdw` — Fdw is a Postgres Foreign Data Wrapper interface written in Go. Dynamic Foreign Tables are defined through gRPC plugins, making them safe, performant and easy to build. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION steampipe_postgres_fdw;

create server
  fdw_aws
foreign data wrapper
  fdw
options (
  wrapper 'aws'
);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `fdw_handler()` is an extension function and returns `fdw_handler`.
- `fdw_validator(text[], oid)` is an extension function and returns `void`.
- `steampipe_postgres_fdw` is an extension-defined foreign data wrapper.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
