## Usage

Sources:

- [Official upstream README](https://github.com/jbylund/gomap_tam/blob/a06f376f5f8b10e8e0daa3ede248edd03eec0178/README.md)
- [Official extension control file (treedb_pgext.control)](https://github.com/jbylund/gomap_tam/blob/a06f376f5f8b10e8e0daa3ede248edd03eec0178/treedb_pgext.control)
- [Official extension SQL (treedb_pgext--1.0.sql)](https://github.com/jbylund/gomap_tam/blob/a06f376f5f8b10e8e0daa3ede248edd03eec0178/treedb_pgext--1.0.sql)

`treedb_pgext` — A PostgreSQL Table Access Method backed by TreeDB, a mmap'd B+tree storage engine written in Go. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION treedb_pgext;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `treedb_am_handler(internal)` is an extension function and returns `table_am_handler`.
- `treedb` is an extension-defined access method.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
