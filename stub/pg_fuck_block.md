## Usage

Sources:

- [Pinned upstream README](https://github.com/sangli00/pg_get_page_tuple/blob/a4330098dfc5535e8f5e416658a17261cafff71e/README)
- [Pinned version 1.0 installation SQL](https://github.com/sangli00/pg_get_page_tuple/blob/a4330098dfc5535e8f5e416658a17261cafff71e/pg_fuck_block--1.0.sql)
- [Pinned C implementation](https://github.com/sangli00/pg_get_page_tuple/blob/a4330098dfc5535e8f5e416658a17261cafff71e/pg_fuck_block.c)
- [Pinned extension control file](https://github.com/sangli00/pg_get_page_tuple/blob/a4330098dfc5535e8f5e416658a17261cafff71e/pg_fuck_block.control)

`pg_fuck_block` version `1.0` exposes the low-level function `pg_get_page_tuple(regclass, bigint)` for inspecting tuple slots in one heap block. It is a legacy diagnostic experiment over PostgreSQL storage internals, not a snapshot-consistent table-reading API.

### Core Workflow

Because the function returns anonymous `record`, supply a column definition list matching the target relation when decomposing its output:

```sql
CREATE EXTENSION pg_fuck_block;

CREATE TABLE page_demo (
  id bigint,
  note text,
  created_at timestamptz
);
INSERT INTO page_demo VALUES (1, 'first row', clock_timestamp());

SELECT *
FROM pg_get_page_tuple('page_demo'::regclass, 0)
  AS t(id bigint, note text, created_at timestamptz);
```

The second argument is a zero-based heap block number. A block outside the current relation size raises `block number out of range`.

### Important Objects

- `pg_get_page_tuple(regclass, bigint) RETURNS SETOF record`: scans normal line pointers in a heap page and returns tuples that the implementation does not classify as `HEAPTUPLE_DEAD`.

### Operational Notes

The function does not enforce table `SELECT` privileges in its C code, yet extension functions are executable by `PUBLIC` by default. Revoke execution from untrusted roles. Its result is not governed by an ordinary MVCC snapshot: recently dead, in-progress, or otherwise invisible tuples may be exposed, while non-normal line pointers are skipped.

The reviewed 2016 implementation depends on removed or changed backend APIs, releases the relation lock and buffer before multi-call iteration completes, and does not populate the output null bitmap from `heap_getattr()`. Concurrent modification, rows containing `NULL`, TOASTed data, or modern PostgreSQL builds can therefore produce incorrect data, memory-safety failures, or backend crashes. Use only on an isolated copy with the exact tested server source; prefer maintained tools such as `pageinspect` for production diagnostics.
