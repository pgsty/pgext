


## Usage

> [log_fdw: Foreign data wrapper for Postgres log file access](https://github.com/aws/postgresql-logfdw)

### Create Server

```sql
CREATE EXTENSION log_fdw;

CREATE SERVER log_fdw_server FOREIGN DATA WRAPPER log_fdw;
```

### List Available Log Files

```sql
SELECT * FROM list_postgres_log_files();
```

Returns the file name and size of each log file in the PostgreSQL log directory.

### Create Foreign Table for CSV Logs

```sql
SELECT * FROM create_foreign_table_for_log_file(
  'postgresql_2024_01_15_csv',   -- foreign table name
  'log_fdw_server',               -- server name
  'postgresql-2024-01-15.csv'     -- log file name
);
```

### Create Foreign Table for Plain Text Logs

```sql
SELECT * FROM create_foreign_table_for_log_file(
  'postgresql_2024_01_15_log',
  'log_fdw_server',
  'postgresql-2024-01-15.log'
);
```

### Query Log Data

```sql
-- Query CSV-format logs (structured columns)
SELECT log_time, error_severity, message
FROM postgresql_2024_01_15_csv
WHERE error_severity = 'ERROR'
ORDER BY log_time DESC
LIMIT 20;

-- Query plain text logs
SELECT * FROM postgresql_2024_01_15_log LIMIT 10;
```

### Granting Access to Non-Superusers

Only superusers can create the extension, but access can be granted:

```sql
GRANT pg_monitor TO monitoring_user;
GRANT CREATE ON SCHEMA public TO monitoring_user;
GRANT USAGE ON FOREIGN SERVER log_fdw_server TO monitoring_user;
```

### Functions Reference

| Function | Description |
|----------|-------------|
| `list_postgres_log_files()` | List available log files and their sizes |
| `create_foreign_table_for_log_file(table_name, server_name, file_name)` | Create a foreign table from a log file |
