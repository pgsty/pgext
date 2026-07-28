## Usage

Sources:

- [Official upstream README](https://github.com/dalibo/data2pg/blob/b4a2bd14b926b82553ccc5cc4e18fe5f140977ca/README.md)
- [Official extension control file (data2pg.control)](https://github.com/dalibo/data2pg/blob/b4a2bd14b926b82553ccc5cc4e18fe5f140977ca/ext/data2pg.control)
- [Official extension SQL (data2pg--0.3.sql)](https://github.com/dalibo/data2pg/blob/b4a2bd14b926b82553ccc5cc4e18fe5f140977ca/ext/data2pg--0.3.sql)

`data2pg` — Migration framework for discovering, copying, and comparing non-PostgreSQL database content through foreign data wrappers. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION data2pg;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `add_step_parent(p_batchName TEXT, p_step TEXT, p_parent_step TEXT)` is an extension function and returns `INT`.
- `assign_fkey_checks_to_batch(p_batchName TEXT, p_schema TEXT, p_table TEXT, p_fkey TEXT DEFAULT NULL)` is an extension function and returns `INT`.
- `assign_index_to_batch(p_batchName TEXT, p_schema TEXT, p_table TEXT, p_object TEXT)` is an extension function and returns `INTEGER`.
- `assign_sequence_to_batch(p_batchName TEXT, p_schema TEXT, p_sequence TEXT)` is an extension function and returns `INTEGER`.
- `assign_sequences_to_batch(p_batchName TEXT, p_schema TEXT, p_sequencesToInclude TEXT, p_sequencesToExclude TEXT)` is an extension function and returns `INT`.
- `assign_table_checks_to_batch(p_batchName TEXT, p_schema TEXT, p_table TEXT)` is an extension function and returns `INTEGER`.
- `assign_table_part_to_batch(p_batchName TEXT, p_schema TEXT, p_table TEXT, p_partId TEXT)` is an extension function and returns `INTEGER`.
- `assign_table_part_to_batch(p_batchName TEXT, p_schema TEXT, p_table TEXT, p_partNum INTEGER)` is an extension function and returns `INTEGER`.
- `assign_table_to_batch(p_batchName TEXT, p_schema TEXT, p_table TEXT)` is an extension function and returns `INTEGER`.
- `assign_tables_checks_to_batch(p_batchName TEXT, p_schema TEXT, p_tablesToInclude TEXT, p_tablesToExclude TEXT)` is an extension function and returns `INTEGER`.
- `assign_tables_to_batch(p_batchName TEXT, p_schema TEXT, p_tablesToInclude TEXT, p_tablesToExclude TEXT)` is an extension function and returns `INTEGER`.
- `check_schema(p_schema TEXT)` is an extension function and returns `void`.
- `complete_migration_configuration(p_migration TEXT)` is an extension function and returns `INT`.
- `create_migration(p_migration TEXT, p_sourceDbms TEXT, p_extension TEXT, p_serverOptions TEXT, p_userMappingOptions TEXT, p_userHasPrivileges BOOLEAN DEFAULT false, p_importSchemaOptions TEXT DEFAULT NULL)` is an extension function and returns `INTEGER`.

### Requirements and Caveats

- The reviewed control file declares default version `0.7`.
- Install the confirmed extension dependencies first: `dblink`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
