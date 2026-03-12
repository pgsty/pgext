


## Usage

> [oracle_fdw: Foreign data wrapper for Oracle access](https://github.com/laurenz/oracle_fdw)

### Create Server

```sql
CREATE EXTENSION oracle_fdw;

CREATE SERVER oradb FOREIGN DATA WRAPPER oracle_fdw
  OPTIONS (dbserver '//dbserver.mydomain.com:1521/ORADB');
```

**Server Options:** `dbserver` (required, Oracle connection string), `isolation_level` (`serializable`, `read_committed`, or `read_only`, default `serializable`), `nchar` (expensive character conversion, default `off`), `set_timezone` (sync timezone with Oracle, default `off`).

### Create User Mapping

```sql
CREATE USER MAPPING FOR pguser SERVER oradb
  OPTIONS (user 'orauser', password 'orapwd');
```

Use an empty string for `user` to enable external (OS) authentication.

### Create Foreign Table

```sql
CREATE FOREIGN TABLE oratab (
  id integer OPTIONS (key 'true') NOT NULL,
  text character varying(30),
  floating double precision NOT NULL
)
SERVER oradb
OPTIONS (schema 'ORAUSER', table 'ORATAB');
```

**Table Options:** `table` (required, Oracle table name in uppercase), `schema` (table owner), `dblink` (Oracle DB link), `max_long` (max LONG column length, default 32767), `readonly` (default `false`), `sample_percent` (ANALYZE sampling, default 100), `prefetch` (rows per round-trip, default 50).

**Column Options:** `key` (mark as primary key, required for UPDATE/DELETE), `strip_zeros` (remove ASCII 0 from strings).

You can also use a query instead of a table name by enclosing it in parentheses:

```sql
CREATE FOREIGN TABLE oraquery (
  id integer,
  text character varying(30)
)
SERVER oradb
OPTIONS (table '(SELECT id, text FROM ORAUSER.ORATAB WHERE id > 10)');
```

### Import Foreign Schema

```sql
IMPORT FOREIGN SCHEMA "ORAUSER"
  FROM SERVER oradb INTO local_schema;
```

**Import Options:** `case` (`keep`, `lower`, or `smart`, default `smart`), `readonly`, `skip_tables`, `skip_views`, `skip_matviews`, `max_long`, `sample_percent`, `prefetch`.

### Utility Functions

```sql
SELECT oracle_diag();              -- Show versions and environment info
SELECT oracle_diag('oradb');       -- Include Oracle server version
SELECT oracle_close_connections(); -- Close all cached Oracle connections
SELECT oracle_execute('oradb', 'TRUNCATE TABLE ORAUSER.ORATAB'); -- Execute arbitrary Oracle SQL
```

### Data Type Mapping

| Oracle Type | PostgreSQL Types |
|---|---|
| CHAR, VARCHAR2, NVARCHAR2 | char, varchar, text |
| CLOB, NCLOB | text, json |
| NUMBER | numeric, float4, float8, int2, int4, int8, boolean |
| DATE, TIMESTAMP | date, timestamp, timestamptz |
| BLOB, LONG RAW | bytea |
| XMLTYPE | xml, text |
| SDO_GEOMETRY | geometry (PostGIS) |
