## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_procfs/pg_procfs-0.0.2/README.md)
- [Official extension control file (pg_procfs.control)](https://api.pgxn.org/src/pg_procfs/pg_procfs-0.0.2/pg_procfs.control)
- [Official extension SQL (pg_procfs--0.0.1.sql)](https://api.pgxn.org/src/pg_procfs/pg_procfs-0.0.2/pg_procfs--0.0.1.sql)

`pg_procfs` — PostgreSQL extension to display /proc FS data from SQL. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_procfs;

select * from pg_procfs('/proc/version');
 line |                                                                                     message

------+--------------------------------------------------------------------------------------------------------------------------------------
--------------------------------------------
    0 | Linux version 4.18.0-372.19.1.el8_6.x86_64 (mockbuild@49c5e54ed716424c9ae8c1a3d1fef96f) (gcc version 8.5.0 20210514 (Red Hat 8.5.0-10
) (GCC)) #1 SMP Tue Aug 2 13:42:59 EDT 2022
(1 row)
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_procfs(IN filename cstring, OUT line integer, OUT data text)` is an extension function and returns `SETOF record`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
