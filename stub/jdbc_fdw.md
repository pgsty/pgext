## Usage

Sources:

- [jdbc_fdw v0.5.0 README](https://github.com/pgspider/jdbc_fdw/blob/v0.5.0/README.md)
- [Extension control file](https://github.com/pgspider/jdbc_fdw/blob/v0.5.0/jdbc_fdw.control)
- [jdbc_fdw v0.5.0 release](https://github.com/pgspider/jdbc_fdw/releases/tag/v0.5.0)

`jdbc_fdw` exposes a JDBC data source as PostgreSQL foreign tables and can execute remote SQL through a helper function. Use it when a suitable JDBC driver exists but no more specialized FDW is available; the JVM, driver JAR, credentials, and remote query behavior all run inside a PostgreSQL backend process.

### Core Workflow

```sql
CREATE EXTENSION jdbc_fdw;

CREATE SERVER reporting_jdbc
  FOREIGN DATA WRAPPER jdbc_fdw
  OPTIONS (
    drivername 'org.postgresql.Driver',
    url 'jdbc:postgresql://db.example/reporting',
    jarfile '/opt/jdbc/postgresql.jar',
    querytimeout '10',
    maxheapsize '256'
  );

CREATE USER MAPPING FOR app_user
  SERVER reporting_jdbc
  OPTIONS (username 'reader', password 'secret');

CREATE FOREIGN TABLE remote_orders (
  id bigint OPTIONS (key 'true'),
  created_at timestamptz,
  total numeric
) SERVER reporting_jdbc;

SELECT * FROM remote_orders WHERE id = 42;
```

There are no table-level options in v0.5.0. Foreign columns map by name. Mark the remote primary-key column with `OPTIONS (key 'true')` when `UPDATE` or `DELETE` needs row identity.

### Import and Direct SQL

```sql
IMPORT FOREIGN SCHEMA public
  FROM SERVER reporting_jdbc
  INTO jdbc_import
  OPTIONS (recreate 'true');

SELECT *
FROM jdbc_exec('reporting_jdbc', 'SELECT id, name FROM customer')
  AS t(id bigint, name text);
```

The upstream README says `IMPORT FOREIGN SCHEMA` currently works only with GridDB. `jdbc_exec` returns `record`, so queries returning columns require a column definition list.

### Important Options and Limits

- Server options: required `drivername` and `url`, absolute `jarfile`, plus `querytimeout` and JVM `maxheapsize`.
- User-mapping options: `username` and `password`.
- Column option: `key = true` identifies primary-key columns for writable operations.
- `jdbc_exec(connname, sql)` executes driver-specific SQL and can return a defined record set.

Version 0.5.0 supports predicate, column, and aggregate pushdown according to the upstream project, but not remote `RETURNING`, `GROUP BY`, `ORDER BY`, casts, or transaction-control statements. Arrays and foreign-table `TRUNCATE` are not implemented. Test type conversion and write semantics with the selected driver.

Protect JAR paths and server definitions from untrusted users, keep passwords in user mappings, and bound the JVM heap and remote query time. The source/package release is 0.5.0 while `jdbc_fdw.control` continues to declare SQL extension version 1.2; use `pg_extension.extversion` rather than assuming those version spaces are identical.
