## Usage

Sources:

- [Project README](https://github.com/rogerwelin/pg_column_tetris/blob/e70f9867c63e932cdaf87b2d34b6504adad9ce12/README.md)
- [Extension control file](https://github.com/rogerwelin/pg_column_tetris/blob/e70f9867c63e932cdaf87b2d34b6504adad9ce12/pg_column_tetris.control)
- [Version 0.1.0 SQL implementation](https://github.com/rogerwelin/pg_column_tetris/blob/e70f9867c63e932cdaf87b2d34b6504adad9ce12/pg_column_tetris--0.1.0.sql)

`pg_column_tetris` 0.1.0 is a pure-SQL extension for PostgreSQL 14 through 18. An event trigger estimates alignment padding after `CREATE TABLE` and can warn about or reject inefficient column order. It also provides inspection and rewrite-suggestion functions.

### Inspect and choose enforcement

The default mode is `warn`; `strict` rejects a newly created table that the estimator considers suboptimal, and `off` disables the event-trigger check.

```sql
CREATE EXTENSION pg_column_tetris;

SELECT column_tetris.mode();
SELECT * FROM column_tetris.check('public.measurement'::regclass);
SELECT column_tetris.padding_wasted('public.measurement'::regclass);

SELECT column_tetris.set_mode('warn');
```

Use `column_tetris.exclude()` for tables that must not be checked. Temporary and system tables are skipped, and the event trigger checks table creation rather than every later alteration.

### Treat estimates and rewrites as advisory

The estimator models tuple headers and type alignment, but it cannot fully predict real storage for null bitmaps, variable-length or toasted values, compression, and workload-specific row populations. A reported byte saving is therefore a design signal, not measured disk reclamation.

`column_tetris.suggest_rewrite()` returns a migration script; it does not preserve every foreign key, index, trigger, or default. The generated sequence renames the original table, creates and copies a replacement, and drops the old table, which can require an exclusive lock and downtime. Never execute that output without reviewing dependent objects, privileges, identity and sequence behavior, replication, rollback, and a realistic staging rehearsal. Column order can also be part of application contracts such as positional inserts and row decoding.
