## Usage

Sources:

- [Integration guide](https://github.com/moriyoshi/yesnodb/blob/6e7ac29bc45256ec3d4e8d15c1a992ca71af19f6/docs/integrations.md)
- [Extension control file](https://github.com/moriyoshi/yesnodb/blob/6e7ac29bc45256ec3d4e8d15c1a992ca71af19f6/yesno-pg/yesno_pg.control)
- [Installation SQL](https://github.com/moriyoshi/yesnodb/blob/6e7ac29bc45256ec3d4e8d15c1a992ca71af19f6/yesno-pg/sql/yesno_pg--0.1.0.sql)
- [Official README](https://github.com/moriyoshi/yesnodb/blob/6e7ac29bc45256ec3d4e8d15c1a992ca71af19f6/yesno-pg/README.md)
- [Operational guide](https://github.com/moriyoshi/yesnodb/blob/6e7ac29bc45256ec3d4e8d15c1a992ca71af19f6/docs/operations.md)

`yesno_pg` is the experimental PostgreSQL integration for yesnodb ordinal sets. It exposes a foreign data wrapper, an index access method and a single-column table access method, with an important durability boundary between PostgreSQL and yesnodb.

### Core Workflow

Install the ABI-matched upstream artifacts for PostgreSQL 17 or 18, then create the extension as a superuser. A foreign table represents one remote key as a set of ordinal values.

```sql
CREATE EXTENSION yesno_pg;
SELECT yesno_pg_version();
CREATE SERVER yesno FOREIGN DATA WRAPPER yesno_fdw
  OPTIONS (endpoint 'https://yesno.internal:50051');
CREATE FOREIGN TABLE rust_docs (ordinal bigint NOT NULL)
  SERVER yesno OPTIONS (key '42');
SELECT count(*) FROM rust_docs;
```

### Access Methods

`yesno` provides equality posting lists and bitmap index scans. `yesno_table` stores a single `bigint` set. A server dictionary option supports named-key import through `IMPORT FOREIGN SCHEMA`. Inspect query plans to verify access-method selection.

```sql
CREATE TABLE selected_ordinals (ordinal bigint) USING yesno_table;
INSERT INTO selected_ordinals VALUES (1), (5), (9);
SELECT * FROM selected_ordinals ORDER BY ordinal;
```

### Isolation and Backup

The engines do not share WAL or a commit clock. Remote writes are not atomic with PostgreSQL heap writes; back up yesnodb separately. For yesno tables, `REPEATABLE READ` pins a version and `READ COMMITTED` takes a statement snapshot, but a transaction spanning both engines can observe different commit moments. `UPDATE`, `SELECT FOR UPDATE`, `ON CONFLICT`, `CLUSTER` and `TABLESAMPLE` are unsupported for the table method. Build artifacts are produced by the upstream Bazel/Docker gate; this is not a general packaged PostgreSQL extension.
