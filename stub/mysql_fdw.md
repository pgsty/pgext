## Usage

Sources:

- [Official source](https://github.com/EnterpriseDB/mysql_fdw/blob/REL-2_10_0/README.md)
- [Official source](https://github.com/EnterpriseDB/mysql_fdw/blob/REL-2_10_0/mysql_fdw.control)
- [Official source](https://github.com/EnterpriseDB/mysql_fdw/blob/REL-2_10_0/mysql_fdw--1.2.sql)
- [Official source](https://github.com/EnterpriseDB/mysql_fdw/releases/tag/REL-2_10_0)

`mysql_fdw` 2.10.0 exposes MySQL tables through PostgreSQL 14–19. Its SQL/control version is 1.2. The server needs the compatible MySQL client library. Installation requires a superuser, without shared preload.

### Connect and Map Tables

Create a server and a restricted remote account mapping for an existing table:

```sql
CREATE EXTENSION mysql_fdw;
CREATE SERVER mysql_server FOREIGN DATA WRAPPER mysql_fdw
  OPTIONS (host '127.0.0.1', port '3306');
CREATE USER MAPPING FOR CURRENT_USER SERVER mysql_server
  OPTIONS (username 'mysql_user', password 'replace_with_remote_password');
CREATE FOREIGN TABLE warehouse (warehouse_id integer, warehouse_name text)
  SERVER mysql_server OPTIONS (dbname 'mydb', table_name 'warehouse');
SELECT * FROM warehouse;
```

`dbname` is required; `table_name` defaults to the foreign-table name. The FDW also supports `IMPORT FOREIGN SCHEMA`. Grant server and table access only to intended roles. Protect mapping credentials and use certificate-verified TLS for untrusted networks.

### Writes and Pushdown

The first foreign-table column is used as the row identifier for `UPDATE` and `DELETE`; it must identify one remote row. Version 2.10.0 rejects those operations when that column belongs to a composite key, because older versions could affect multiple rows sharing its value. Match local constraints to remote data. Remote writes have their own transaction/failure boundaries.

The FDW can push filters, joins, aggregates, ordering and limits. Check `EXPLAIN` before relying on an index or pushdown. `fetch_size`, `use_remote_estimate` and `reconnect` affect remote access. Ordering optimization trusts local NOT NULL declarations; a nullable remote column can therefore sort differently if the declarations disagree.

### Upgrade and Security Changes

2.10.0 fixes remote/local SQL injection, generated-column INSERT handling, repeated-parameter queries and binary/text buffer reads. `init_command`, `mysql_default_file` and SSL path/options are restricted to superusers. Existing foreign servers created by other users with restricted options need a superuser to alter or recreate them.

`NO_BACKSLASH_ESCAPES` is always added to the remote `sql_mode`. Imported defaults are now string literals; restore expression defaults such as `CURRENT_TIMESTAMP` on the foreign table after import. Install the new library and reconnect existing sessions; SQL version 1.2 does not change. Recheck writable-table keys, defaults, escaping and affected-row counts before returning the connection to applications.
