


## Usage

> [sqlite_fdw: SQLite Foreign Data Wrapper](https://github.com/pgspider/sqlite_fdw)

### Create Server

```sql
CREATE EXTENSION sqlite_fdw;

CREATE SERVER sqlite_server FOREIGN DATA WRAPPER sqlite_fdw
  OPTIONS (database '/path/to/database.db');
```

**Server Options:** `database` (required, path to SQLite file), `updatable` (default `true`), `truncatable` (default `false`), `keep_connections` (default `true`), `batch_size` (default 1), `force_readonly` (default `false`).

No `CREATE USER MAPPING` is required since SQLite has no authentication model.

### Create Foreign Table

```sql
CREATE FOREIGN TABLE remote_data (
  id integer OPTIONS (key 'true'),
  name text,
  created timestamp OPTIONS (column_type 'INT'),
  data bytea
)
SERVER sqlite_server
OPTIONS (table 'data_table');
```

**Table Options:** `table` (SQLite table name if different from PostgreSQL name), `updatable`, `truncatable`, `batch_size`.

**Column Options:** `column_name` (map to different SQLite column name), `column_type` (SQLite affinity: `INT` for epoch timestamps, `BLOB` for UUIDs), `key` (mark as primary key for UPDATE/DELETE).

### CRUD Operations

```sql
SELECT * FROM remote_data WHERE id > 100;
INSERT INTO remote_data (id, name) VALUES (1, 'test');
UPDATE remote_data SET name = 'updated' WHERE id = 1;
DELETE FROM remote_data WHERE id = 1;
```

### Import Foreign Schema

```sql
IMPORT FOREIGN SCHEMA public
  FROM SERVER sqlite_server INTO local_schema;
```

**Import Options:** `import_default` (default `false`), `import_not_null` (default `true`).

### Data Type Mapping

| SQLite Type | PostgreSQL Type |
|-------------|-----------------|
| int | bigint |
| text, char, clob | text |
| blob | bytea |
| real, float, double | double precision |
| datetime | timestamp |
| uuid | uuid |
| json, jsonb | json, jsonb |

Timestamps can be stored as TEXT (ISO format) or INT (Unix epoch, use `column_type 'INT'`). UUIDs can be stored as TEXT (36 chars) or BLOB (16 bytes). The SQLite database file must be readable (and writable for DML) by the PostgreSQL OS user.
