## Usage

Sources:

- [Official README](https://github.com/vderic/postgres-kite/blob/ef6334886b99ea83cf8df253da5992732599f551/README.md)
- [Extension control file](https://github.com/vderic/postgres-kite/blob/ef6334886b99ea83cf8df253da5992732599f551/kite_fdw.control)
- [Installation SQL](https://github.com/vderic/postgres-kite/blob/ef6334886b99ea83cf8df253da5992732599f551/kite_fdw--1.0.sql)
- [Installation SQL](https://github.com/vderic/postgres-kite/blob/ef6334886b99ea83cf8df253da5992732599f551/kite_fdw--1.0--1.1.sql)

`kite_fdw` queries a compatible remote Kite service using PostgreSQL foreign tables. The reviewed source is historical and does not establish compatibility with current PostgreSQL releases.

### Core Workflow

Configure Kite endpoints, the remote database and per-user credentials. Foreign-table options describe the remote table or file pattern and data format.

```sql
CREATE EXTENSION kite_fdw;
CREATE SERVER kite_server FOREIGN DATA WRAPPER kite_fdw
  OPTIONS (host '127.0.0.1:7878', dbname 'pgsql', fragcnt '4');
CREATE USER MAPPING FOR CURRENT_USER SERVER kite_server
  OPTIONS (username 'app_user', password 'replace-with-password');
CREATE FOREIGN TABLE warehouse (warehouse_id int, warehouse_name text)
  SERVER kite_server
  OPTIONS (schema_name 'public', table_name 'warehouse*', fmt 'csv', csv_header 'false');
SELECT * FROM warehouse;
```

### Options and Requirements

`host` accepts comma-separated endpoints; `dbname` is required; `fragcnt` controls query fragments. `fetch_size` defaults to 100 and can be set at server or table level. `fmt` supports CSV or Parquet, with CSV delimiter, quote, escape, header and null-string options. The control version is `1.1`, reached from the shipped base SQL and upgrade script. No standalone license declaration or supported-major matrix was found; review the source and external Kite dependencies before adoption.
