## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/stanislawbartkowski/db2odbc_fdw/blob/fefa0d2181e4089fa39f0743c473d49dcf7a5e3f/README.md)
- [Version 1.0 SQL objects](https://github.com/stanislawbartkowski/db2odbc_fdw/blob/fefa0d2181e4089fa39f0743c473d49dcf7a5e3f/db2odbc_fdw--1.0.sql)
- [FDW implementation](https://github.com/stanislawbartkowski/db2odbc_fdw/blob/fefa0d2181e4089fa39f0743c473d49dcf7a5e3f/db2odbc_fdw.c)

`db2odbc_fdw` is a PostgreSQL 12-era read wrapper for DB2 through an operating-system ODBC data source. A foreign table maps to a user-supplied remote `sql_query`; server options include the ODBC `dsn` and an optional cached-connection setting.

```sql
CREATE EXTENSION db2odbc_fdw;
CREATE SERVER db2_server
  FOREIGN DATA WRAPPER db2odbc_fdw OPTIONS (dsn 'REPORTING');
CREATE USER MAPPING FOR report_role SERVER db2_server
  OPTIONS (username 'reporter', password 'secret');
CREATE FOREIGN TABLE db2_customer (id integer, name text)
  SERVER db2_server OPTIONS (sql_query 'SELECT ID, NAME FROM CUSTOMER');
```

Use a real secret workflow rather than keeping the example password in SQL, restrict access to user mappings and ODBC configuration, and never grant the FDW or server to `PUBLIC`. Treat `sql_query` as trusted administrator input because it defines remote SQL.

Upstream points advanced users to a fuller DB2 FDW and documents PostgreSQL 12 only. Validate ODBC driver ABI, encodings, numeric and timestamp conversion, nulls, remote errors, connection retry/caching, cancellation, timeouts, credential rotation, transaction boundaries, and predicate behavior. Assume no write or distributed-transaction semantics unless the exact code and tests demonstrate them.
