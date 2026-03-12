


## Usage

> [db2_fdw: Foreign data wrapper for DB2 access](https://github.com/wolfgangbrandl/db2_fdw)

### Create Server

```sql
CREATE EXTENSION db2_fdw;

CREATE SERVER db2srv FOREIGN DATA WRAPPER db2_fdw
  OPTIONS (dbserver 'SAMPLE');
```

**Server Options:** `dbserver` (required, DB2 database connection string).

### Create User Mapping

```sql
CREATE USER MAPPING FOR PUBLIC SERVER db2srv
  OPTIONS (user 'db2inst1', password 'secret');
```

Use empty strings for `user` and `password` to enable external authentication.

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

**Table Options:** `table` (required, DB2 table name, case-sensitive, typically uppercase), `schema` (table owner), `readonly` (default `false`), `prefetch` (rows per round-trip, default 200, range 0-10240), `max_long` (max LONG column length, default 32767).

**Column Options:** `key` (set to `true` for primary key columns, required for UPDATE/DELETE).

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

### Data Type Mapping

| DB2 Type | PostgreSQL Types |
|----------|------------------|
| CHAR | char |
| VARCHAR | varchar |
| CLOB | text |
| BLOB | bytea |
| SMALLINT, INTEGER, BIGINT | smallint, integer, bigint |
| DOUBLE | numeric, float |
| DATE | date |
| TIMESTAMP | timestamp |
| TIME | time |

WHERE conditions and column projections are pushed down to DB2 to minimize data transfer.
