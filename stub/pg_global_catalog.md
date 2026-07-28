## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_global_catalog/pg_global_catalog-0.0.1/README.md)
- [Official extension control file (pg_global_catalog.control)](https://api.pgxn.org/src/pg_global_catalog/pg_global_catalog-0.0.1/pg_global_catalog.control)
- [Official extension SQL (pg_global_catalog--0.0.1.sql)](https://api.pgxn.org/src/pg_global_catalog/pg_global_catalog-0.0.1/pg_global_catalog--0.0.1.sql)

`pg_global_catalog` — PostgreSQL extension to consolidate pg_catalog for each database in a single schema named global_catalog. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_global_catalog;

select
datname, count(*)
from global_catalog.pg_class c
join pg_database d
on c.dbid = d.oid
group by datname
order by datname;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pggc_create_fdws()` is an extension function and returns `void`.
- `pggc_create_global_views()` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
