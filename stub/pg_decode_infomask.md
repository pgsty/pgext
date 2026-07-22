## Usage

Sources:

- [Official README](https://github.com/begriffs/pg_decode_infomask/blob/3cac9efcb408743601c9eb98b24704854210e027/README.md)
- [Extension SQL](https://github.com/begriffs/pg_decode_infomask/blob/3cac9efcb408743601c9eb98b24704854210e027/pg_decode_infomask--0.1.sql)
- [C implementation](https://github.com/begriffs/pg_decode_infomask/blob/3cac9efcb408743601c9eb98b24704854210e027/pg_decode_infomask.c)
- [Extension control file](https://github.com/begriffs/pg_decode_infomask/blob/3cac9efcb408743601c9eb98b24704854210e027/pg_decode_infomask.control)

`pg_decode_infomask` decodes PostgreSQL 9.6 heap tuple `t_infomask` and `t_infomask2` integers into named boolean fields. It helps inspect page-level MVCC hint and lock bits, but it does not read heap pages or determine transaction status by itself.

### Core Workflow

Install `pageinspect` separately to obtain tuple-header values, then pass the integers to the decoder:

```sql
CREATE EXTENSION pageinspect;
CREATE EXTENSION pg_decode_infomask;

SELECT h.lp, h.t_xmin, h.t_xmax, x.*
FROM heap_page_items(get_raw_page('public.demo', 0)) AS h
CROSS JOIN LATERAL pg_get_xact_infomask(h.t_infomask) AS x;
```

`pageinspect` is used by the example but is not declared as an extension dependency. Its raw-page functions normally require elevated privileges; apply its own security guidance.

### Functions

- `pg_get_xact_infomask_bits(integer)` returns booleans for committed, invalid, and frozen `xmin` bits and committed or invalid `xmax` bits.
- `pg_get_xact_infomask(integer)` adds `xmin_status text[]` and `xmax_status text[]` summaries.
- `pg_get_lock_infomask_bits(integer)` exposes tuple-lock and multixact-related bit tests.
- `pg_get_infomask2_bits(integer)` returns the attribute count and key-updated, HOT-updated, and heap-only flags.

These functions interpret only the supplied bit masks. A hint bit is not a live lookup of `pg_xact`, and the decoder cannot prove current commit status, visibility, or data consistency.

### Version and Safety Boundaries

The SQL file explicitly targets PostgreSQL 9.6 internal bit definitions. Heap tuple flags and server C APIs are major-version-specific, so do not use this build to interpret values from another PostgreSQL major. Build and validate a version-specific port before relying on the output.

The C entry points are not declared `STRICT` and read their first argument without checking for SQL null. Never pass `NULL`. The source also uses the old two-argument `CreateTemplateTupleDesc` API and has not been updated since 2017, so it will not build unchanged on modern PostgreSQL releases. Use the extension for controlled forensic work, cross-check decoded values against the matching server source, and avoid treating it as a production integrity checker.
