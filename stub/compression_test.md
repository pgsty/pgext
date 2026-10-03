## Usage

Sources:

- [Official README](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/compress_test/README)
- [Control 1.0](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/compress_test/compression_test.control)
- [SQL 1.0](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/compress_test/compression_test--1.0.sql)
- [Implementation](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/compress_test/compression_test.c)

`compression_test` 1.0 exposes backend utilities for experimenting with pglz compression and raw relation pages. It is a developer module from Michael Paquier's plugin collection. The module README states PostgreSQL 9.5 onward, but it relies on server internals; verify the target major when building.

### Core Workflow

Install as a superuser. A separate schema avoids collisions with the module's generic function names:

```sql
CREATE SCHEMA compression_lab;
CREATE EXTENSION compression_test WITH SCHEMA compression_lab;

WITH input AS (
  SELECT convert_to(repeat('abc', 100), 'UTF8') AS original
), packed AS (
  SELECT original, compression_lab.compress_data(original) AS compressed
  FROM input
)
SELECT compression_lab.bytea_size(original) AS original_bytes,
       compression_lab.bytea_size(compressed) AS compressed_bytes
FROM packed;
```

This compares byte counts only. Compression returns the original bytes when the algorithm cannot compress them; the result carries no flag distinguishing that case. Keep the original bytes and length and do not treat every output as a decompressible frame.

### Function Index

- `compress_data(bytea)` uses the built-in always-try pglz strategy. Its seven-argument overload also accepts minimum and maximum input sizes, minimum compression rate, first-success threshold and two match-search controls. These are low-level experiment settings.
- `decompress_data(bytea, smallint)` needs the original uncompressed length as its second argument. The declared type limits that length to a positive signed 16-bit value. Invalid or incompressible input can fail; this is not a self-describing archive format.
- `bytea_size(bytea)` returns payload bytes, excluding the varlena header.
- `get_raw_page(oid, int4, bool)` returns page bytes and a hole offset. It checks superuser privilege. A false final argument removes the page's free-space hole; a true argument preserves a full page. Do not rely on that branch to sanitize unused bytes: the current implementation clears its local buffer after copying the returned page.

### Boundaries

Raw pages expose physical storage contents, not MVCC-filtered query results. Use only controlled test data and retain the schema's access restrictions. These helpers neither enable table compression nor define a stable, portable persistence format. They do not require a background worker or preloading.
