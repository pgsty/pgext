


## Usage

> [mysql_fdw: Foreign data wrapper for querying a MySQL server](https://github.com/EnterpriseDB/mysql_fdw)

### Create Server

```sql
CREATE EXTENSION mysql_fdw;

CREATE SERVER mysql_server
  FOREIGN DATA WRAPPER mysql_fdw
  OPTIONS (host '127.0.0.1', port '3306');
```

**Server Options:** `host` (default `127.0.0.1`), `port` (default `3306`), `secure_auth` (default `true`), `init_command`, `use_remote_estimate` (default `false`), `reconnect` (default `false`), `sql_mode` (default `ANSI_QUOTES`), `fetch_size` (default `100`), `character_set`, `truncatable` (default `true`), plus SSL options (`ssl_key`, `ssl_cert`, `ssl_ca`, `ssl_capath`, `ssl_cipher`).

### Create User Mapping

```sql
CREATE USER MAPPING FOR pguser
  SERVER mysql_server
  OPTIONS (username 'mysqluser', password 'mysqlpass');
```

### Create Foreign Table

```sql
CREATE FOREIGN TABLE warehouse (
  warehouse_id int,
  warehouse_name text,
  warehouse_created timestamp
)
SERVER mysql_server
OPTIONS (dbname 'mydb', table_name 'warehouse');
```

**Table Options:** `dbname` (required, MySQL database name), `table_name` (defaults to foreign table name), `fetch_size` (overrides server setting), `max_blob_size`, `truncatable` (default `true`).

### CRUD Operations

```sql
INSERT INTO warehouse VALUES (1, 'UPS', current_date);
SELECT * FROM warehouse ORDER BY warehouse_id;
UPDATE warehouse SET warehouse_name = 'NEW_NAME' WHERE warehouse_id = 1;
DELETE FROM warehouse WHERE warehouse_id = 3;
```

### Pushdown Features

mysql_fdw optimizes queries through several pushdown mechanisms:

- **WHERE clause** pushdown to minimize data transfer
- **Column** pushdown to retrieve only requested columns
- **JOIN** pushdown for joins between foreign tables on the same MySQL server
- **AGGREGATE** pushdown for `min`, `max`, `sum`, `avg`, `count`
- **ORDER BY** and **LIMIT/OFFSET** pushdown to reduce network traffic
