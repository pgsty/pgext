

## 用法

> [pgdd: 通过标准 SQL 内省 PostgreSQL 数据字典](https://github.com/rustprooflabs/pgdd)

PgDD 在 `dd` 模式中提供数据字典视图，用于通过标准 SQL 内省数据库对象。

### 数据库概览

```sql
SELECT * FROM dd.database;
```

返回：`db_name`、`db_size`、`schema_count`、`table_count`、`size_in_tables`、`view_count`、`size_in_views`、`extension_count`。

### 模式

```sql
SELECT s_name, table_count, view_count, function_count, size_plus_indexes, description
  FROM dd.schemas;
```

### 表

```sql
SELECT t_name, size_pretty, rows, bytes_per_row
  FROM dd.tables
  WHERE s_name = 'public';
```

### 视图

```sql
SELECT s_name, v_name, description FROM dd.views;
```

### 列

```sql
SELECT source_type, s_name, t_name, c_name, data_type
  FROM dd.columns
  WHERE data_type LIKE 'int%';
```

### 函数

```sql
SELECT s_name, f_name, argument_data_types, result_data_types FROM dd.functions;
```

### 分区表

```sql
SELECT * FROM dd.partition_parents WHERE s_name = 'public';
SELECT * FROM dd.partition_children WHERE s_name = 'public';
```

`partition_parents` 视图显示聚合的分区统计信息（数量、总大小、总行数）。`partition_children` 视图显示每个分区的详情及相对于父表的百分比计算。

默认排除系统对象。如需包含系统对象，请直接查询底层函数：`SELECT * FROM dd.tables() WHERE system_object;`
