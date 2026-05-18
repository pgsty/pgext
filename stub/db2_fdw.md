
## Usage

Sources: [README](https://github.com/wolfgangbrandl/db2_fdw), [current upstream README](https://github.com/pg-fdw/db2_fdw/blob/master/README.md)

`db2_fdw` is a PostgreSQL foreign data wrapper for querying and modifying IBM Db2 tables from PostgreSQL. It pushes down required columns and `WHERE` conditions where possible, and provides helper functions for connection cleanup and diagnostics.

### Create Server

```sql
CREATE EXTENSION db2_fdw;

CREATE SERVER db2srv FOREIGN DATA WRAPPER db2_fdw
  OPTIONS (dbserver 'SAMPLE');
```

**Server options:** `dbserver` (required Db2 connection string), `batch_size` (currently reserved for future batch behavior), and `no_encoding_error` (`ON`, `OFF`, `YES`, `NO`, `TRUE`, or `FALSE`).

### Create User Mapping

```sql
CREATE USER MAPPING FOR PUBLIC SERVER db2srv
  OPTIONS (user 'db2inst1', password 'secret');
```

Use empty strings for `user` and `password` to enable external authentication through the Db2 client environment.

### Create Foreign Table

```sql
CREATE FOREIGN TABLE employee (
  empno char(6) OPTIONS (key 'true'),
  firstname varchar(12),
  lastname varchar(15),
  salary numeric
)
SERVER db2srv
OPTIONS (schema 'DB2INST1', table 'EMPLOYEE');
```

**Table options:** `table` (required, Db2 table name or simple query, case-sensitive, typically uppercase), `schema` (table owner), `readonly` (default `false`), `sample_percent` (ANALYZE sampling), `prefetch` (rows per round-trip, default `100`, range `0`-`1024`), `fetch_size` (accepted but currently fixed at `1`), `batch_size`, and `no_encoding_error`. `max_long` is documented upstream as deprecated and no longer used.

**Column options:** `key` (set to `true` for all primary key columns, required for `UPDATE` and `DELETE`), plus Db2 metadata options such as `db2type`, `db2size`, `db2bytes`, `db2chars`, `db2scale`, `db2null`, and `db2ccsid` on imported tables.

### Import Foreign Schema

```sql
IMPORT FOREIGN SCHEMA "DB2INST1" FROM SERVER db2srv INTO public;
```

**Import Options:** `case` (`keep`, `lower`, or `smart`, default `smart`), `readonly`.

### CRUD Operations

```sql
SELECT * FROM employee WHERE empno = '000010';
INSERT INTO employee (empno, firstname, lastname, salary) VALUES ('999999', 'John', 'Doe', 50000);
UPDATE employee SET salary = 55000 WHERE empno = '999999';
DELETE FROM employee WHERE empno = '999999';
```

### Connection Helpers

```sql
SELECT db2_close_connections();
SELECT db2_diag();
SELECT db2_diag('db2srv');
```

`db2_close_connections()` closes cached Db2 connections in the current session. `db2_diag()` reports db2_fdw, PostgreSQL, Db2 client, and optionally remote server diagnostic details.

### Data Type Mapping

| DB2 Type | PostgreSQL Types |
|----------|------------------|
| CHAR | char |
| VARCHAR | varchar |
| CLOB | text |
| VARGRAPHIC, GRAPHIC | text |
| BLOB | bytea |
| SMALLINT, INTEGER, BIGINT | smallint, integer, bigint |
| DOUBLE | numeric, float |
| DATE | date |
| TIMESTAMP | timestamp |
| TIME | time |

WHERE conditions and column projections are pushed down to DB2 to minimize data transfer.
