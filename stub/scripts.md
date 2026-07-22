## Usage

Sources:

- [Official repository README at the catalog revision](https://github.com/a5221985/PhaseZero_Scripts/blob/443e3a5e59d2816b7b8d7ee0129c8685f3f11d8a/README.md)
- [Build recipe at the catalog revision](https://github.com/a5221985/PhaseZero_Scripts/blob/443e3a5e59d2816b7b8d7ee0129c8685f3f11d8a/Makefile)
- [Table DDL at the catalog revision](https://github.com/a5221985/PhaseZero_Scripts/blob/443e3a5e59d2816b7b8d7ee0129c8685f3f11d8a/sql/create_table_ddl.sql)
- [Index DDL at the catalog revision](https://github.com/a5221985/PhaseZero_Scripts/blob/443e3a5e59d2816b7b8d7ee0129c8685f3f11d8a/sql/create_index.sql)
- [Constraint DDL at the catalog revision](https://github.com/a5221985/PhaseZero_Scripts/blob/443e3a5e59d2816b7b8d7ee0129c8685f3f11d8a/sql/create_constraint.sql)

`scripts` 1.4 is not a general-purpose function library. It packages PhaseZero application DDL that creates the pzv_aftermarket schema, a fixed set of business/staging tables, indexes, and foreign keys. Install it only for that legacy application after reviewing the complete generated SQL.

### Core Workflow

The Makefile generates the extension script by concatenating the table, index, and constraint files. Inspect that generated artifact before installation and use a disposable database first.

```sql
CREATE EXTENSION scripts;

SELECT n.nspname, c.relname, c.relkind
FROM pg_class AS c
JOIN pg_namespace AS n ON n.oid = c.relnamespace
WHERE n.nspname = 'pzv_aftermarket'
ORDER BY c.relkind, c.relname;
```

### Installed Objects

- Core tables include company, category, attribute, part, scope_part, display_metadata, orders, and order_parts.
- Staging/support tables include stage_category, stage_attribute, stage_part, and temp_part_child.
- The script adds application-specific indexes and foreign keys between several of those relations. It exposes no user-facing SQL functions.

### Critical Caveats

- Object names, columns, and constraints are hard-coded. Installation can collide with an existing pzv_aftermarket schema or partially matching objects, and the DDL is not written as a repeatable schema migration.
- The repository's standalone database-creation file is not included by the extension build recipe; CREATE EXTENSION does not create a database.
- The control-file comment describes an unrelated key/value data type, while the actual generated SQL is application schema DDL. Trust the reviewed SQL, not that metadata comment.
- Upstream is abandoned and targets an old PGXS path. Take a backup, test install/uninstall behavior, and prefer a maintained migration system for new deployments.
