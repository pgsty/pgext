


## Usage

Sources: [upstream README](https://github.com/ibarwick/firebird_fdw), [1.4.2 README](https://github.com/ibarwick/firebird_fdw/blob/REL_1_4_STABLE/README.md), [1.4.2 tag](https://github.com/ibarwick/firebird_fdw/tree/1.4.2).

`firebird_fdw` connects PostgreSQL to Firebird databases through the foreign data wrapper API. It supports reads, writes, `IMPORT FOREIGN SCHEMA`, predicate pushdown for supported expressions, connection caching, and foreign-table `TRUNCATE` on PostgreSQL 14+.

### Create Server

```sql
CREATE EXTENSION firebird_fdw;

CREATE SERVER firebird_server FOREIGN DATA WRAPPER firebird_fdw
  OPTIONS (address 'localhost', database '/path/to/database.fdb');
```

Server options include:

- `address`, default `localhost`.
- `port`, default `3050`.
- `database`, the Firebird database name or path.
- `updatable`, default `true`; table-level settings can override it.
- `disable_pushdowns`, useful for debugging or benchmarking.
- `quote_identifiers`, to quote table and column identifiers by default.
- `implicit_bool_type`, for integer-backed Firebird boolean columns.
- `batch_size`, for multi-row inserts on PostgreSQL 14+.

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

With a custom query, the foreign table is read-only:

```sql
CREATE FOREIGN TABLE fb_query (
  id smallint,
  val varchar(2048)
)
SERVER firebird_server
OPTIONS (query $$ SELECT id, val FROM fdw_test WHERE id > 10 $$);
```

Table options include `table_name`, `query`, `updatable`, `estimated_row_count`, `quote_identifier`, and `batch_size`. Column options include `column_name`, `quote_identifier`, and `implicit_bool_type`.

### Import Foreign Schema

```sql
IMPORT FOREIGN SCHEMA someschema
  FROM SERVER firebird_server
  INTO public
  OPTIONS (import_views 'true', verbose 'true');
```

Import options include `import_not_null`, `import_views`, `updatable`, and `verbose`. The schema name has no special Firebird meaning and can be any value accepted by PostgreSQL's `IMPORT FOREIGN SCHEMA` syntax.

### CRUD Operations

```sql
SELECT * FROM fb_test WHERE id > 5;
INSERT INTO fb_test VALUES (10, 'new record');
UPDATE fb_test SET val = 'updated' WHERE id = 10;
DELETE FROM fb_test WHERE id = 10;
TRUNCATE fb_test;
```

`UPDATE` and `DELETE` use Firebird's `RDB$DB_KEY`. `TRUNCATE` is implemented as an unqualified Firebird `DELETE` because Firebird has no native `TRUNCATE` statement.

### Utility Functions

- `firebird_fdw_version()` returns the FDW version as an integer.
- `firebird_fdw_close_connections()` closes cached Firebird connections for the current PostgreSQL session.
- `firebird_fdw_server_options(servername text)` shows effective server option values and whether each was explicitly provided.
- `firebird_fdw_diag()` returns diagnostic key/value data, including FDW and `libfq` versions.
- `firebird_version()` reports Firebird server versions for configured foreign servers and may open connections to do so.

### Caveats

- Pigsty packages `firebird_fdw` 1.4.2 for PostgreSQL 14-18. Upstream documents compatibility with PostgreSQL 10-19, with newer FDW API features available only on newer PostgreSQL releases.
- Upstream supports Firebird 2.5 and later; older Firebird versions are not a tested target.
- `firebird_fdw` and `libfq` are developed together, so package compatibility depends on matching those libraries.
- Custom-query foreign tables cannot be updated.
- The `implicit_bool_type` feature is experimental and is designed around integer columns representing boolean values.
