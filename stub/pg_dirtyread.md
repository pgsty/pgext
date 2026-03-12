


## Usage

> [pg_dirtyread: Read dead but unvacuumed rows from table](https://github.com/df7cb/pg_dirtyread)

`pg_dirtyread` allows reading dead (deleted/updated) rows that have not yet been vacuumed. The function returns RECORD, so a table alias describing the schema is required.

### Basic Usage

```sql
SELECT * FROM pg_dirtyread('foo') AS t(bar bigint, baz text);
```

### Example

```sql
CREATE TABLE foo (bar bigint, baz text);
ALTER TABLE foo SET (autovacuum_enabled = false, toast.autovacuum_enabled = false);

INSERT INTO foo VALUES (1, 'Test'), (2, 'New Test');
DELETE FROM foo WHERE bar = 1;

SELECT * FROM pg_dirtyread('foo') AS t(bar bigint, baz text);
 bar |   baz
-----+----------
   1 | Test
   2 | New Test
```

### Dropped Columns

Access dropped column content using `dropped_N` (Nth column, 1-based), as long as the table has not been rewritten (e.g., via `VACUUM FULL` or `CLUSTER`):

```sql
ALTER TABLE ab DROP COLUMN b;
DELETE FROM ab;
SELECT * FROM pg_dirtyread('ab') ab(a text, dropped_2 text);
```

### System Columns

Include system columns in the alias to retrieve them. A special `dead` boolean column reports dead rows:

```sql
SELECT * FROM pg_dirtyread('foo')
    AS t(tableoid oid, ctid tid, xmin xid, xmax xid, cmin cid, cmax cid, dead boolean,
         bar bigint, baz text);
```

The `dead` column is not usable during recovery (e.g., on standby servers).
