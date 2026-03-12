


## Usage

> [pgdd: Introspect pg data dictionary via standard SQL](https://github.com/rustprooflabs/pgdd)

PgDD provides data dictionary views in the `dd` schema for introspecting database objects via standard SQL.

### Database Overview

```sql
SELECT * FROM dd.database;
```

Returns: `db_name`, `db_size`, `schema_count`, `table_count`, `size_in_tables`, `view_count`, `size_in_views`, `extension_count`.

### Schemas

```sql
SELECT s_name, table_count, view_count, function_count, size_plus_indexes, description
  FROM dd.schemas;
```

### Tables

```sql
SELECT t_name, size_pretty, rows, bytes_per_row
  FROM dd.tables
  WHERE s_name = 'public';
```

### Views

```sql
SELECT s_name, v_name, description FROM dd.views;
```

### Columns

```sql
SELECT source_type, s_name, t_name, c_name, data_type
  FROM dd.columns
  WHERE data_type LIKE 'int%';
```

### Functions

```sql
SELECT s_name, f_name, argument_data_types, result_data_types FROM dd.functions;
```

### Partitioned Tables

```sql
SELECT * FROM dd.partition_parents WHERE s_name = 'public';
SELECT * FROM dd.partition_children WHERE s_name = 'public';
```

The `partition_parents` view shows aggregate partition stats (count, total size, total rows). The `partition_children` view shows per-partition details with percentage calculations against the parent.

System objects are excluded by default. To include them, query the underlying functions directly: `SELECT * FROM dd.tables() WHERE system_object;`
