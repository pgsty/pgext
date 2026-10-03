## Usage

Sources:

- [Official README](https://github.com/polydbms/pg_xdbc_fdw/blob/fac1fc6b5275d327843a3efb4d519f9afe6a161b/README.md)
- [Extension control file](https://github.com/polydbms/pg_xdbc_fdw/blob/fac1fc6b5275d327843a3efb4d519f9afe6a161b/pg_xdbc_fdw.control)
- [Installation SQL](https://github.com/polydbms/pg_xdbc_fdw/blob/fac1fc6b5275d327843a3efb4d519f9afe6a161b/pg_xdbc_fdw--0.1.sql)
- [Official SQL example](https://github.com/polydbms/pg_xdbc_fdw/blob/fac1fc6b5275d327843a3efb4d519f9afe6a161b/test/test_fdw_create.sql)
- [Official README](https://github.com/polydbms/pg_xdbc_fdw/blob/fac1fc6b5275d327843a3efb4d519f9afe6a161b/docker/README.md)

`pg_xdbc_fdw` is a research connector between PostgreSQL and the XDBC client/server data-transfer stack. Upstream provides a PostgreSQL 13 container topology and experiment setup, rather than a general-purpose ODBC compatibility layer.

### Core Workflow

Install the XDBC client dependencies and run the corresponding server before querying. The local JSON schema must match both the declared columns and remote data.

```sql
CREATE EXTENSION pg_xdbc_fdw;
CREATE SERVER xdbcserver FOREIGN DATA WRAPPER pg_xdbc_fdw;
CREATE FOREIGN TABLE transfer_data (id integer, value text)
  SERVER xdbcserver
  OPTIONS (schema_file_path '/srv/xdbc/schema.json', server_host 'xdbcserver', table 'transfer_data');
SELECT * FROM transfer_data;
```

### Boundaries

The extension installs `pg_xdbc_fdw_handler()` and its foreign data wrapper. `schema_file_path` is read from the PostgreSQL host; `server_host` and `table` select the XDBC source. A matching schema file and reachable service are prerequisites, not created by this SQL. The upstream test setup also uses buffer settings; choose them according to the exact XDBC build. No general schema-import, write, security or cross-system transaction guarantees are established by the project documentation.
