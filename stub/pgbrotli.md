## Usage

Sources:

- [Version 1.0 SQL objects](https://github.com/johto/pgbrotli/blob/21c84a7d3dc0777c60db1d6266736e5e8f4e5906/pgbrotli--1.0.sql)
- [C implementation](https://github.com/johto/pgbrotli/blob/21c84a7d3dc0777c60db1d6266736e5e8f4e5906/pgbrotli.c)
- [Extension control file](https://github.com/johto/pgbrotli/blob/21c84a7d3dc0777c60db1d6266736e5e8f4e5906/pgbrotli.control)

`pgbrotli` 1.0 provides in-backend Brotli compression for PostgreSQL text. `pgbrotli_compress(text)` returns compressed `bytea`, and `pgbrotli_decompress(bytea)` returns the original bytes as `bytea`.

### Core Workflow

```sql
CREATE EXTENSION pgbrotli;
WITH packed AS (
  SELECT pgbrotli_compress('hello world'::text) AS value
)
SELECT convert_from(pgbrotli_decompress(value), 'UTF8')
FROM packed;
```

The text compressor uses Brotli's maximum quality, default window, and text mode; these settings are not SQL parameters. Although version 1.0 declares `pgbrotli_compress(bytea)`, its C entry point always raises `not implemented yet`. Binary input is therefore unsupported. Decompression always returns `bytea`; use `convert_from` only when the original bytes are valid in the chosen encoding.

### Resource and Compatibility Limits

Compression and decompression run in the database backend. The code caps an input or expanded value at 1 GiB, but a tiny untrusted stream can still drive repeated allocation toward that limit. Reject untrusted or oversized payloads before calling it, apply statement and resource limits, and avoid repeated decompression in scans. Malformed and truncated streams raise errors.

This abandoned project has no current PostgreSQL/Brotli compatibility matrix or documented wire-format contract beyond its source. Verify empty values, library ABI, architecture, malformed streams, memory exhaustion, and parallel execution on the exact build. PostgreSQL TOAST may already compress large values, so measure end-to-end storage and CPU before adding application-level Brotli.
