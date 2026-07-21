## Usage

Sources:

- [odbc_fdw 0.6.1 README](https://github.com/devrimgunduz/odbc_fdw/blob/0.6.1/README.md)
- [odbc_fdw changelog](https://github.com/devrimgunduz/odbc_fdw/blob/0.6.1/NEWS.md)
- [Extension control file](https://github.com/devrimgunduz/odbc_fdw/blob/0.6.1/odbc_fdw.control)
- [0.6.0 to 0.6.1 comparison](https://github.com/devrimgunduz/odbc_fdw/compare/0.6.0...0.6.1)

`odbc_fdw` exposes tables or driver-specific queries from an ODBC data source as PostgreSQL foreign tables. It is primarily a read/query bridge across heterogeneous systems; validate data-type conversions and remote-driver behavior before relying on it for production queries.

### Core Workflow

```sql
CREATE EXTENSION odbc_fdw;

-- In 0.6.1 a superuser must set the server-level dsn or driver option.
CREATE SERVER warehouse_odbc
  FOREIGN DATA WRAPPER odbc_fdw
  OPTIONS (dsn 'warehouse');

CREATE USER MAPPING FOR analyst
  SERVER warehouse_odbc
  OPTIONS (odbc_UID 'reporter', odbc_PWD 'secret');

CREATE FOREIGN TABLE remote_customer (
  id bigint,
  name text,
  created_at timestamp
) SERVER warehouse_odbc
  OPTIONS (schema 'sales', table 'customer');

SELECT * FROM remote_customer WHERE id = 42;
```

Use `driver` instead of `dsn` for a DSN-less connection. Other driver attributes use the `odbc_` prefix and may be placed on the server, user mapping, or foreign table. Put credentials in a user mapping. Quote case-sensitive attribute names, and wrap values containing `=` or `;` in braces as required by the driver.

### Queries and Import

`sql_query` overrides `table`; pair it with `sql_count` when the FDW needs an explicit row-count query:

```sql
CREATE FOREIGN TABLE active_customer (
  id bigint,
  name text
) SERVER warehouse_odbc
  OPTIONS (
    sql_query 'SELECT id, name FROM sales.customer WHERE active = 1',
    sql_count 'SELECT count(*) FROM sales.customer WHERE active = 1'
  );

IMPORT FOREIGN SCHEMA sales
  FROM SERVER warehouse_odbc
  INTO imported
  OPTIONS (prefix 'odbc_');
```

### Important Objects and Options

- `dsn` or `driver` selects the ODBC data source; 0.6.1 restricts these server options to superusers because the driver manager loads shared libraries.
- `schema`, `table`, `sql_query`, and `sql_count` select the remote relation or query.
- `prefix` changes local names created by `IMPORT FOREIGN SCHEMA`.
- `ODBCTablesList(server_name, ...)` lists visible remote tables.
- `ODBCTableSize(server_name, table_name)` and `ODBCQuerySize(server_name, query)` return remote row counts.

Version 0.6.0 restores compatibility and fixes crashes on recent PostgreSQL releases. Version 0.6.1 escapes remote literals and identifiers to prevent SQL injection, restricts driver selection, and redacts common credential attributes in debug connection strings. Upgrade before allowing delegated FDW use, while retaining normal server ownership and user-mapping controls.

Only the ODBC types listed by the upstream README are fully supported. Identifier length, driver SQL dialect, encodings, null handling, and binary values can vary. The source/package release is 0.6.1, while the control file and install SQL continue to declare extension version 0.5.2.
