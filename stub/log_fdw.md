
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION log_fdw;
> CREATE SERVER log_fdw_server FOREIGN DATA WRAPPER log_fdw;
> SELECT * FROM list_postgres_log_files();
> ```
>
> Source: [README](https://github.com/aws/postgresql-logfdw)

`log_fdw` is a PostgreSQL foreign data wrapper for reading PostgreSQL log files through SQL. It provides helper functions to list files in the server log directory and to create foreign tables for individual log files.

### Core Functions

The upstream README documents two SQL entry points:

```sql
create_foreign_table_for_log_file(table_name text, server_name text, log_file_name text)
list_postgres_log_files()
```

`list_postgres_log_files()` is a compatibility wrapper around PostgreSQL core's `pg_ls_logdir()`.

## Basic Workflow

Create the extension and a foreign server:

```sql
CREATE EXTENSION log_fdw;
CREATE SERVER log_fdw_server FOREIGN DATA WRAPPER log_fdw;
```

List files available in the PostgreSQL log directory:

```sql
SELECT * FROM list_postgres_log_files() ORDER BY 1 DESC LIMIT 10;
```

Create foreign tables for CSV logs or plain `.log` files:

```sql
SELECT * FROM create_foreign_table_for_log_file(
  'postgresql_2022_11_28_csv',
  'log_fdw_server',
  'postgresql-2022-11-28.csv'
);

SELECT * FROM create_foreign_table_for_log_file(
  'postgresql_2022_11_28_log',
  'log_fdw_server',
  'postgresql-2022-11-28.log'
);
```

## Querying

Foreign tables created from plain log files expose a single log-entry style column, while CSV log files expose structured columns such as `log_time`, `error_severity`, `message`, and session metadata.

Typical usage is straightforward:

```sql
SELECT * FROM postgresql_2022_11_28_log LIMIT 2;

SELECT log_time, error_severity, message
FROM postgresql_2022_11_28_csv
WHERE error_severity = 'ERROR'
ORDER BY log_time DESC
LIMIT 20;
```

## Privileges

Only superusers can create the extension. The README also notes that superusers can delegate access to non-superusers with the minimum required grants, for example:

```sql
CREATE ROLE foo;
GRANT pg_monitor TO foo;
GRANT CREATE ON SCHEMA bar TO foo;
GRANT USAGE ON FOREIGN SERVER log_fdw_server TO foo;
```

`pg_monitor` is specifically needed when `list_postgres_log_files()` is used, because the underlying `pg_ls_logdir()` function requires it.

## Build Notes

The project can be built standalone with PGXS:

```bash
export USE_PGXS=1
make
make install
```

The source can also be copied into PostgreSQL's `contrib` tree and built there as part of a larger distribution.
