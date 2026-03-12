


## Usage

> [tds_fdw: Foreign data wrapper for querying a TDS database (Sybase or Microsoft SQL Server)](https://github.com/tds-fdw/tds_fdw)

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

**Table Options:** `table_name` or `query` (one required, mutually exclusive), `schema_name`, `match_column_names` (map by name vs position), `use_remote_estimate`, `row_estimate_method` (`execute` or `showplan_all`).

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
  INTO public;
```
