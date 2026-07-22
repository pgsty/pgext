## Usage

Sources:

- [Official extension control file](https://github.com/vibhorkum/edb_foreignkeyconstraintmanager/blob/50a89fb337426e810eecca4f3e4f1b00d7bb1023/edb_foreignkeyconstraintmanager.control)
- [Official upstream documentation](https://github.com/vibhorkum/edb_foreignkeyconstraintmanager/blob/50a89fb337426e810eecca4f3e4f1b00d7bb1023/README.md)

`edb_foreignkeyconstraintmanager` manages foreign-key behavior for legacy partitioned tables on EDB Postgres Advanced Server in Oracle compatibility mode. It combines native constraints with `refint` triggers and depends on EDB-specific catalogs, so it is not a portable solution for community PostgreSQL or modern declarative partitioning.

### Core Workflow

Install the required `refint` extension and the manager, then create a relationship through the `edb_util` API:

```sql
CREATE EXTENSION IF NOT EXISTS refint;
CREATE EXTENSION edb_foreignkeyconstraintmanager;

SELECT edb_util.create_fk_constraint(
  'parent',
  ARRAY['id'],
  'child',
  ARRAY['parent_id'],
  'restrict'
);
```

The final argument selects `cascade`, `restrict`, or `setnull` behavior. The function returns a boolean and creates native foreign keys for ordinary relations or trigger-based enforcement when the referenced relation is partitioned.

### Important Objects

- `edb_util.create_fk_constraint(regclass, text[], regclass, text[], text)` creates a single- or multi-column relationship and the required constraints or triggers.
- `edb_util.alter_table_drop_partition(name, name, name[], text)` coordinates removal of an EDB partition while maintaining the managed relationships.
- Managed partition triggers use names beginning with `EDB_partition_`.

### Compatibility and Maintenance

The implementation checks for EnterpriseDB and the `redwood` database dialect and reads EDB-only partition catalogs. Use the exact EDB Postgres Advanced Server generation for which this project was written. After adding a partition, rerun `edb_util.create_fk_constraint(...)`; when dropping one, use the supplied partition-drop function. Constraint or trigger cleanup may otherwise require manual work. Test composite keys, null handling, each cascade mode, concurrent writes, and partition DDL before adopting it, and execute the dynamic-DDL functions only through a tightly controlled administrative role.
