## Usage

Sources:

- [Official extension control file](https://github.com/vibhorkum/edb_foreignkeyconstraintmanager/blob/50a89fb337426e810eecca4f3e4f1b00d7bb1023/edb_foreignkeyconstraintmanager.control)
- [Official upstream documentation](https://github.com/vibhorkum/edb_foreignkeyconstraintmanager/blob/50a89fb337426e810eecca4f3e4f1b00d7bb1023/README.md)

`edb_foreignkeyconstraintmanager` — Manage trigger-based foreign-key behavior for partitioned tables on EDB Postgres Advanced Server

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `refint`.

```sql
CREATE EXTENSION "edb_foreignkeyconstraintmanager";
SELECT extversion
FROM pg_extension
WHERE extname = 'edb_foreignkeyconstraintmanager';
```

The upstream project is associated with `EDB`; verify its current support, license, packaging, and deployment boundary from the linked source.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
