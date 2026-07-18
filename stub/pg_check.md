## Usage

Sources:

- [Project README](https://github.com/tvondra/pg_check/blob/db34fe20ddf5c597fe97f634a765360b022d13ab/README.md)
- [Extension control file](https://github.com/tvondra/pg_check/blob/db34fe20ddf5c597fe97f634a765360b022d13ab/pg_check.control)
- [Version 0.1.0 installation SQL](https://github.com/tvondra/pg_check/blob/db34fe20ddf5c597fe97f634a765360b022d13ab/sql/pg_check--0.1.0.sql)

`pg_check` 0.1.0 performs basic consistency checks on heap tables and B-tree indexes. It examines page headers, overlapping items, tuple attribute bounds, and invalid varlena sizes, and can cross-check heap tuples against index entries to find missing or superfluous items.

### Run a conservative check

```sql
CREATE EXTENSION pg_check;

SELECT pg_check_table('public.orders'::regclass, false, false);
SELECT pg_check_index('public.orders_pkey'::regclass);
```

`pg_check_table(regclass,bool,bool,bigint,bigint)` returns an issue count and can optionally check indexes and cross-check them against the heap. `pg_check_index(regclass,bigint,bigint)` checks all or a selected block range. Detail is emitted as server messages according to `client_min_messages`; `pg_check.debug` and `pg_check.bitmap_format` control additional cross-check diagnostics.

### Locking and scope

Do not omit the Boolean arguments without understanding the defaults: both index checking and cross-checking default to true. A cross-check takes `SHARE ROW EXCLUSIVE` on the table, blocks data changes for the duration, and can deadlock with sessions that acquire table and index locks in another order. A non-cross-check uses `ACCESS SHARE`.

The extension reports problems but does not repair them, and structural index validation is limited to B-tree page/item checks. The reviewed C source predates current PostgreSQL releases and publishes no modern compatibility matrix; test it against the exact server build and take a backup before using low-level diagnostic tooling on important data.
