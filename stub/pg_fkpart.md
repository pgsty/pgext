

## Usage

> [pg_fkpart: Table partitioning by foreign key utility](https://github.com/lemoineat/pg_fkpart)

`pg_fkpart` enables partitioning PostgreSQL tables based on a foreign key relationship.
All functions reside in the `pgfkpart` schema.

### Create the Extension

```sql
CREATE EXTENSION pg_fkpart;
```

### Partition a Table by Foreign Key

```sql
SELECT pgfkpart.partition_with_fk(
    'public',           -- schema of the table to partition
    'my_table',         -- table to partition
    'public',           -- schema of the foreign key table
    'fk_table',         -- foreign key table
    true                -- support SQL RETURNING
);
```

### Unpartition a Table

```sql
SELECT pgfkpart.unpartition_with_fk('public', 'my_table');
```

### Index Management

Propagate indexes and constraints from parent to all child tables:

```sql
SELECT pgfkpart.dispatch_index('public', 'my_table');
```

Drop an index across all child tables:

```sql
SELECT pgfkpart.drop_index('public', 'my_table', 'my_index_name');
```

Drop a unique constraint across all child tables:

```sql
SELECT pgfkpart.drop_unique_constraint('public', 'my_table', 'my_constraint_name');
```

### Partition Tracking

The extension maintains a `pgfkpart.partition` table that records all partitioned tables,
including schema, table name, foreign key column, and foreign table information.
