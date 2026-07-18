## Usage

Sources:

- [Pinned upstream README](https://github.com/MasahikoSawada/pg_foobar/blob/1cf8a2ea6c701726b14241a16d95ece30e643ad8/README.md)
- [Version 1.0 installation SQL](https://github.com/MasahikoSawada/pg_foobar/blob/1cf8a2ea6c701726b14241a16d95ece30e643ad8/pg_foobar--1.0.sql)
- [Pinned parallel-worker implementation](https://github.com/MasahikoSawada/pg_foobar/blob/1cf8a2ea6c701726b14241a16d95ece30e643ad8/pg_foobar.c)

pg_foobar version 1.0 is an instructional dynamic-parallel-worker example. One function launches the requested number of workers; each worker writes a sequence of foo and bar messages to the PostgreSQL server log. The function itself always returns null text.

### Test-only example

Use very small values on a disposable PostgreSQL instance and inspect the server log rather than the query result:

```sql
CREATE EXTENSION pg_foobar;

SELECT pg_foobar(2, 3, 4);
```

This asks two workers to log three foo messages and four bar messages each. Available worker slots, server logging configuration, and concurrent background-worker use determine whether it can start.

### API defect and compatibility

The installation SQL declares the first argument as regclass, but the C function reads it as a plain 32-bit worker count. The README nevertheless passes an integer. This catalog/API mismatch can cause confusing coercion, display, and function-resolution behavior; inspect the installed signature and never pass a relation name or OID-derived value as the worker count.

The extension can consume many dynamic worker slots and generate a large volume of logs, so do not expose the function to ordinary users. Revoke public execution and cap all three arguments in any wrapper used for demonstrations.

Upstream documents only PostgreSQL 9.6 and 10 beta 1. The source depends on old parallel-worker internals, has not changed since 2017, and has no current compatibility matrix, although the repository is not archived. It is sample code rather than a production logging, scheduling, or monitoring facility.
