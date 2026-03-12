


## Usage

> [informix_fdw: Foreign data wrapper for Informix access](https://github.com/credativ/informix_fdw)

### Create Server

```sql
CREATE EXTENSION informix_fdw;

CREATE SERVER informix_server
  FOREIGN DATA WRAPPER informix_fdw
  OPTIONS (
    informixserver 'informix_hostname',
    informixdir '/opt/informix/csdk'
  );
```

**Server Options:** `informixserver` (required, Informix server identifier from sqlhosts file), `informixdir` (required, path to Informix Client SDK), `disable_predicate_pushdown` (disable WHERE pushdown), `gl_datetime` (custom datetime format, default `%iY-%m-%d %H:%M:%S`), `gl_date` (custom date format, default `%iY-%m-%d`).

### Create User Mapping

```sql
CREATE USER MAPPING FOR CURRENT_USER
  SERVER informix_server
  OPTIONS (username 'informix', password 'secret');
```

### Create Foreign Table

```sql
CREATE FOREIGN TABLE remote_table (
  id bigint NOT NULL,
  name varchar(255),
  amount numeric(10,2),
  created timestamp
)
SERVER informix_server
OPTIONS (
  database 'mydb',
  table 'remote_table',
  client_locale 'en_US.utf8',
  db_locale 'en_US.819'
);
```

**Table Options:** `database` (required, remote database name), `table` or `query` (one required), `client_locale` (required, must match PostgreSQL server encoding), `db_locale` (required, should match Informix locale), `disable_rowid` (use updatable cursor instead of ROWID), `enable_blobs` (include if table contains BLOBs).

Use `query` instead of `table` for view-like access:

```sql
CREATE FOREIGN TABLE remote_view (
  id bigint,
  total numeric(10,2)
)
SERVER informix_server
OPTIONS (
  database 'mydb',
  query 'SELECT id, SUM(amount) AS total FROM orders GROUP BY id',
  client_locale 'en_US.utf8',
  db_locale 'en_US.819'
);
```

### CRUD Operations

```sql
SELECT * FROM remote_table WHERE id > 100;
INSERT INTO remote_table (id, name, amount) VALUES (1, 'test', 99.99);
UPDATE remote_table SET amount = 100.00 WHERE id = 1;
DELETE FROM remote_table WHERE id = 1;
```

### Supported Data Types

Queries: BOOLEAN, DATE, DATETIME, INTERVAL, SMALLINT, INTEGER, BIGINT, SERIAL, VARCHAR, CHARACTER, TEXT, NUMERIC, MONEY.

DML operations: SERIAL, INTEGER, BIGINT, INTERVAL, TEXT, BYTEA, VARCHAR, NVARCHAR, TIMESTAMP, DATE, NUMERIC, MONEY.
