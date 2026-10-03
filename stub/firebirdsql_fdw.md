## Usage

Sources:

- [Official README](https://github.com/nakagami/firebirdsql_fdw/blob/c223920d389c3a247d89388472f73d65f5660a0c/README.md)
- [Extension control file](https://github.com/nakagami/firebirdsql_fdw/blob/c223920d389c3a247d89388472f73d65f5660a0c/firebirdsql_fdw.control)
- [Implementation (lib.rs)](https://github.com/nakagami/firebirdsql_fdw/blob/c223920d389c3a247d89388472f73d65f5660a0c/src/lib.rs)
- [Cargo package metadata](https://github.com/nakagami/firebirdsql_fdw/blob/c223920d389c3a247d89388472f73d65f5660a0c/Cargo.toml)

`firebirdsql_fdw` is a Rust/pgrx wrapper for Firebird. Creating this extension installs an SQL FDW named `firebird_fdw`; it can therefore conflict with another extension that already owns that wrapper name.

### Core Workflow

Install a build matching PostgreSQL 13–18. Remote credentials are stored as server options, so restrict catalog access and server-definition privileges.

```sql
CREATE EXTENSION firebirdsql_fdw;
CREATE SERVER my_firebird FOREIGN DATA WRAPPER firebird_fdw
  OPTIONS (host 'localhost', port '3050', db_name '/firebird/data/mydb.fdb',
           username 'app_user', password 'replace-with-password');
CREATE FOREIGN TABLE employees (id integer, name text, hired date)
  SERVER my_firebird OPTIONS (table 'EMPLOYEES', rowid_column 'ID');
SELECT * FROM employees WHERE id = 1;
```

### Objects and Limits

`IMPORT FOREIGN SCHEMA firebird` discovers remote tables. Reads support predicate/order/limit pushdown; inserts, updates and deletes are exposed, with `rowid_column` required to identify rows for updates and deletes. The client uses `firebirust`; no separate SQL extension dependency is declared. Control version `0.1.0` is not relocatable and sets `superuser=false`, but creating a C-backed FDW still requires the appropriate PostgreSQL privileges. Do not assume cross-system transaction atomicity.
