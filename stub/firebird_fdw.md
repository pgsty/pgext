


## Usage

> [firebird_fdw: Foreign data wrapper for Firebird](https://github.com/ibarwick/firebird_fdw)

### Create Server

```sql
CREATE EXTENSION firebird_fdw;

CREATE SERVER firebird_server FOREIGN DATA WRAPPER firebird_fdw
  OPTIONS (address 'localhost', database '/path/to/database.fdb');
```

**Server Options:** `address` (default `localhost`), `port` (default `3050`), `database` (required, path to Firebird database file), `updatable` (default `true`), `disable_pushdowns` (disable WHERE clause pushdown), `quote_identifiers`, `implicit_bool_type` (enable integer-to-boolean conversion), `batch_size` (PostgreSQL 14+).

### Create User Mapping

```sql
CREATE USER MAPPING FOR CURRENT_USER SERVER firebird_server
  OPTIONS (username 'sysdba', password 'masterke');
```

### Create Foreign Table

```sql
CREATE FOREIGN TABLE fb_test (
  id smallint,
  val varchar(2048)
)
SERVER firebird_server
OPTIONS (table_name 'fdw_test');
```

With column name mapping:

```sql
CREATE FOREIGN TABLE fb_mapped (
  id smallint OPTIONS (column_name 'test_id'),
  val varchar(2048) OPTIONS (column_name 'test_val')
)
SERVER firebird_server
OPTIONS (table_name 'fdw_test');
```

With a custom query (read-only):

```sql
CREATE FOREIGN TABLE fb_query (
  id smallint,
  val varchar(2048)
)
SERVER firebird_server
OPTIONS (query $$ SELECT id, val FROM fdw_test WHERE id > 10 $$);
```

**Table Options:** `table_name`, `query` (mutually exclusive with `table_name`, read-only), `updatable`, `estimated_row_count`, `quote_identifier`, `batch_size`.

**Column Options:** `column_name`, `quote_identifier`, `implicit_bool_type`.

### Import Foreign Schema

```sql
IMPORT FOREIGN SCHEMA someschema
  FROM SERVER firebird_server
  INTO public
  OPTIONS (import_views 'true', verbose 'true');
```

**Import Options:** `import_not_null` (default `true`), `import_views` (default `true`), `updatable`, `verbose`.

The schema parameter has no particular meaning in Firebird and can be set to any value.

### CRUD Operations

```sql
SELECT * FROM fb_test WHERE id > 5;
INSERT INTO fb_test VALUES (10, 'new record');
UPDATE fb_test SET val = 'updated' WHERE id = 10;
DELETE FROM fb_test WHERE id = 10;
```
