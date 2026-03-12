


## Usage

> [odbc_fdw: Foreign data wrapper for accessing remote databases using ODBC](https://github.com/CartoDB/odbc_fdw)

### Create Server

Connect using a DSN defined in your ODBC configuration:

```sql
CREATE EXTENSION odbc_fdw;

CREATE SERVER odbc_server
  FOREIGN DATA WRAPPER odbc_fdw
  OPTIONS (dsn 'test');
```

Or specify connection attributes directly without a DSN:

```sql
CREATE SERVER odbc_server
  FOREIGN DATA WRAPPER odbc_fdw
  OPTIONS (
    odbc_DRIVER 'MySQL',
    odbc_SERVER '192.168.1.17',
    encoding 'iso88591'
  );
```

**Server Options:** `dsn` (ODBC data source name), `driver` (ODBC driver name, required if no DSN), `odbc_*` (driver-specific attributes), `encoding` (remote database character encoding).

Prefix driver-specific options with `odbc_`. Attributes DSN, DRIVER, UID, and PWD are automatically uppercased.

### Create User Mapping

```sql
CREATE USER MAPPING FOR postgres
  SERVER odbc_server
  OPTIONS (odbc_UID 'root', odbc_PWD '');
```

### Create Foreign Table

```sql
CREATE FOREIGN TABLE odbc_table (
  id integer,
  name varchar(255),
  description text,
  users float4,
  created timestamp
)
SERVER odbc_server
OPTIONS (
  odbc_DATABASE 'mydb',
  schema 'test',
  sql_query 'SELECT id, name, description, created, users FROM test.mytable',
  sql_count 'SELECT count(id) FROM test.mytable'
);

SELECT * FROM odbc_table;
```

**Table Options:** `schema` (remote schema), `table` (remote table name), `sql_query` (custom SQL query, overrides `table`), `sql_count` (custom count SQL).

### Import Foreign Schema

```sql
IMPORT FOREIGN SCHEMA test
  FROM SERVER odbc_server
  INTO public
  OPTIONS (odbc_DATABASE 'mydb');
```
