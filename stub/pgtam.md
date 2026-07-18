## Usage

Sources:

- [Upstream README](https://github.com/eatonphil/pgtam/blob/b32fd6ef442d2cd046bc609cfb2575a88ec82211/README.md)
- [Version 0.0.1 install SQL](https://github.com/eatonphil/pgtam/blob/b32fd6ef442d2cd046bc609cfb2575a88ec82211/pgtam--0.0.1.sql)
- [In-memory table access method](https://github.com/eatonphil/pgtam/blob/b32fd6ef442d2cd046bc609cfb2575a88ec82211/pgtam.c)
- [Upstream demonstration](https://github.com/eatonphil/pgtam/blob/b32fd6ef442d2cd046bc609cfb2575a88ec82211/test.sql)
- [MIT license](https://github.com/eatonphil/pgtam/blob/b32fd6ef442d2cd046bc609cfb2575a88ec82211/LICENSE.md)

pgtam 0.0.1 is a teaching prototype for PostgreSQL's table access-method API. Its mem access method stores a tiny amount of data in ordinary process memory and demonstrates callbacks; it is not a storage engine.

### Disposable demonstration

Use a dedicated single-session test server whose contents can be lost:

```sql
CREATE EXTENSION pgtam;
CREATE TABLE mem_demo(a integer) USING mem;
INSERT INTO mem_demo VALUES (23), (101);
SELECT * FROM mem_demo;
```

The rows exist only in the current backend's private memory and disappear when that process exits.

### Caveats

- There is no WAL, disk persistence, shared state, MVCC, crash recovery, replication, or transaction rollback. Other sessions see different memory, and aborted inserts are not undone.
- Update, delete, locking, speculative insertion, bulk insertion, index lookup, vacuum, analyze, sampling, and many other callbacks are empty or return hard-coded success values.
- Storage is capped at 100 tables and 100 rows, with incomplete bound checks. Only non-null 32-bit integer data is assumed, and scans populate only the first column.
- Tables are matched only by unqualified relation name, so same-named relations in different schemas collide. Dropping a relation does not free its private table data.
- Loading the handler appends debug output to a fixed /tmp/pgtam.log path without checking file-open failure, creating reliability and local file-safety concerns.
- The TableAmRoutine structure is major-version-specific and the Makefile hard-codes a developer pg_config path. Build only against the exact reviewed PostgreSQL source tree.
