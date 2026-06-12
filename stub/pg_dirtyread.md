## Usage

Sources: [upstream README](https://github.com/df7cb/pg_dirtyread), [Debian changelog](https://github.com/df7cb/pg_dirtyread/blob/master/debian/changelog), [tags](https://github.com/df7cb/pg_dirtyread/tags).

`pg_dirtyread` reads dead or updated heap rows that have not yet been vacuumed away. The function returns `record`, so every call needs a table alias that declares the columns you want back.

### Basic Usage

```sql
CREATE EXTENSION pg_dirtyread;

SELECT *
FROM pg_dirtyread('foo') AS t(bar bigint, baz text);
```

Columns are matched by name, so the alias can omit columns or place them in a different order.

### Example

```sql
CREATE TABLE foo (bar bigint, baz text);
ALTER TABLE foo SET (autovacuum_enabled = false, toast.autovacuum_enabled = false);

INSERT INTO foo VALUES (1, 'Test'), (2, 'New Test');
DELETE FROM foo WHERE bar = 1;

SELECT * FROM pg_dirtyread('foo') AS t(bar bigint, baz text);
```

The deleted row can remain visible until vacuum removes it.

### Dropped Columns

Dropped column contents can be retrieved as long as the table has not been rewritten by operations such as `VACUUM FULL` or `CLUSTER`. Use `dropped_N`, where `N` is the original 1-based column position:

```sql
CREATE TABLE ab(a text, b text);
INSERT INTO ab VALUES ('Hello', 'World');
ALTER TABLE ab DROP COLUMN b;
DELETE FROM ab;

SELECT *
FROM pg_dirtyread('ab') AS ab(a text, dropped_2 text);
```

Only limited type checks are possible because PostgreSQL removes the dropped column's original type metadata.

### System Columns

Include system columns in the alias to retrieve them. A special `dead boolean` column reports rows that are surely dead:

```sql
SELECT *
FROM pg_dirtyread('foo') AS t(
  tableoid oid,
  ctid tid,
  xmin xid,
  xmax xid,
  cmin cid,
  cmax cid,
  dead boolean,
  bar bigint,
  baz text
);
```

The `dead` column is not usable during recovery, including on standby servers. The `oid` system column is only available on PostgreSQL 11 and earlier.

### Caveats

- Pigsty packages `pg_dirtyread` 2.8 as RPMs for PostgreSQL 14-18; DEB availability comes from PGDG as `postgresql-$v-dirtyread`.
- Upstream 2.8 is a PostgreSQL 19 compatibility release for the default TOAST compression change to `lz4`; no new user-facing SQL function is documented.
- `pg_dirtyread` is for forensic and recovery-style inspection. It bypasses normal MVCC visibility expectations and should not be used for application reads.
