
## Usage

- Sources: [README](https://github.com/tds-fdw/tds_fdw/blob/master/README.md), [foreign server](https://github.com/tds-fdw/tds_fdw/blob/master/ForeignServerCreation.md), [foreign table](https://github.com/tds-fdw/tds_fdw/blob/master/ForeignTableCreation.md), [user mapping](https://github.com/tds-fdw/tds_fdw/blob/master/UserMappingCreation.md), [foreign schema](https://github.com/tds-fdw/tds_fdw/blob/master/ForeignSchemaImporting.md), [variables](https://github.com/tds-fdw/tds_fdw/blob/master/Variables.md)

`tds_fdw` is a foreign data wrapper for querying TDS databases such as Sybase and Microsoft SQL Server through a DB-Library implementation such as FreeTDS.

### Create Server

```sql
CREATE EXTENSION tds_fdw;

CREATE SERVER mssql_svr
  FOREIGN DATA WRAPPER tds_fdw
  OPTIONS (servername '127.0.0.1', port '1433',
           database 'tds_fdw_test', tds_version '7.1');
```

**Server Options:** `servername` (server address or DSN, supports comma-separated failover list), `port`, `database`, `dbuse` (0 for direct connection, non-0 for dbuse()), `tds_version` (protocol version), `language`, `character_set`, `msg_handler` (`notice` or `blackhole`), `sqlserver_ansi_mode`, `fdw_startup_cost`, `fdw_tuple_cost`.

### Create User Mapping

```sql
CREATE USER MAPPING FOR postgres
  SERVER mssql_svr
  OPTIONS (username 'sa', password 'secret');
```

For Azure SQL databases, use the format `username@servername` for the `username` option.

### Create Foreign Table

Map a remote table directly:

```sql
CREATE FOREIGN TABLE mssql_table (
  id integer,
  name varchar(255),
  value numeric(10,2)
)
SERVER mssql_svr
OPTIONS (schema_name 'dbo', table_name 'mytable');
```

Or use a custom SQL query:

```sql
CREATE FOREIGN TABLE mssql_query (
  id integer,
  name varchar(255),
  total numeric(10,2)
)
SERVER mssql_svr
OPTIONS (query 'SELECT id, name, SUM(amount) AS total FROM orders GROUP BY id, name');
```

**Table Options:** `table_name` or `query` (one required, mutually exclusive), `schema_name`, `match_column_names` (map by name vs position), `use_remote_estimate`, `local_tuple_estimate`, `row_estimate_method` (`execute` or `showplan_all`).

**Column Options:** `column_name` (remote column name if different from local).

### Query and Debug

```sql
SELECT * FROM mssql_table WHERE id > 100;

-- View the remote query sent to SQL Server
EXPLAIN (VERBOSE) SELECT * FROM mssql_table WHERE id > 100;
```

### Import Foreign Schema

```sql
IMPORT FOREIGN SCHEMA dbo
  FROM SERVER mssql_svr
  INTO public
  OPTIONS (import_default 'true');
```

**Import Options:** `import_default`, `import_not_null`, and `keep_custom_types` for preserving Sybase user-defined types when matching PostgreSQL domains already exist.

### Planner And Runtime Notes

The upstream README says the current version does not support join pushdown or write operations. It does support `WHERE` and column pushdown when `match_column_names` is enabled.

Set diagnostic memory-stat variables with PostgreSQL `SET`, for example:

```sql
SET tds_fdw.show_finished_memory_stats = 1;
```

Available variables are `tds_fdw.show_before_row_memory_stats`, `tds_fdw.show_after_row_memory_stats`, and `tds_fdw.show_finished_memory_stats`.

Pigsty package metadata is version `2.0.5` from PGDG for PostgreSQL 14-18. Upstream docs say the FDW should support PostgreSQL 9.2+ and the current build matrix includes PostgreSQL 13-18, but this stub follows the packaged PostgreSQL versions from `db/extension.csv`.
