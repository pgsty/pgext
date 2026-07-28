## Usage

Sources:

- [Official upstream README](https://github.com/xingjianwei/pg_analytics/blob/924deb4044b4e1c40766cd342ec990bf47066702/README.md)
- [Official extension control file (th_dbdm.control)](https://github.com/xingjianwei/pg_analytics/blob/924deb4044b4e1c40766cd342ec990bf47066702/th_dbdm.control)
- [Official implementation source](https://github.com/xingjianwei/pg_analytics/blob/924deb4044b4e1c40766cd342ec990bf47066702/src/lib.rs)

`th_dbdm` — pg_analytics (formerly named pg_lakehouse) puts DuckDB inside Postgres. With pg_analytics installed, Postgres can query foreign object stores like AWS S3 and table formats like Iceberg or Delta Lake. Queries are pushed down to DuckDB, a high performance analytical query engine. Use it for the corresponding analytical or storage workflow. The reviewed upstream material marks this capability deprecated.

### Core Workflow

```sql
CREATE EXTENSION th_dbdm;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `1.3.3`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Upstream material contains an explicit deprecation boundary.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
