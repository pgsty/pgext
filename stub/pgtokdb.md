## Usage

Sources:

- [Official upstream README](https://github.com/hughhyndman/pgtokdb/blob/5e79929cbc13fcb69cdb36d90ce2af0723687ab5/README.md)
- [Official extension control file (pgtokdb.control)](https://github.com/hughhyndman/pgtokdb/blob/5e79929cbc13fcb69cdb36d90ce2af0723687ab5/pgtokdb.control)
- [Official extension SQL (pgtokdb--0.0.1.sql)](https://github.com/hughhyndman/pgtokdb/blob/5e79929cbc13fcb69cdb36d90ce2af0723687ab5/pgtokdb--0.0.1.sql)

`pgtokdb` — This project is the implementation of a PostgreSQL extension that allows Postgres processes to access kdb+ data through its SQL interface. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgtokdb;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgtokdb.genddl(varchar, varchar, varchar, varchar)` is an extension function and returns `setof`.
- `pgtokdb.getstatus(varchar)` is an extension function and returns `setof`.
- `pgtokdb.genddl_t` is an extension-defined type.
- `pgtokdb.getstatus_t` is an extension-defined type.
- `pgtokdb` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
