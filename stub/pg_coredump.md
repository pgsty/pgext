## Usage

Sources:

- [Official upstream README](https://github.com/percona-lab/pg_coredump/blob/e63295452fc6b379302f10193d1edfeee7fa8a94/README.md)
- [Official extension control file (pg_coredump.control)](https://github.com/percona-lab/pg_coredump/blob/e63295452fc6b379302f10193d1edfeee7fa8a94/pg_coredump.control)
- [Official extension SQL (pg_coredump--1.0.sql)](https://github.com/percona-lab/pg_coredump/blob/e63295452fc6b379302f10193d1edfeee7fa8a94/pg_coredump--1.0.sql)

`pg_coredump` — The PostgreSQL extension makes it easier to generate coredump files for PostgreSQL in the advent of a crash. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_coredump;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_coredump(dumpdir text)` is an extension function and returns `boolean`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
