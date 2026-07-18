## Usage

Sources:

- [Version 1.0 SQL objects](https://github.com/johto/pgbrotli/blob/21c84a7d3dc0777c60db1d6266736e5e8f4e5906/pgbrotli--1.0.sql)
- [C implementation](https://github.com/johto/pgbrotli/blob/21c84a7d3dc0777c60db1d6266736e5e8f4e5906/pgbrotli.c)
- [Extension control file](https://github.com/johto/pgbrotli/blob/21c84a7d3dc0777c60db1d6266736e5e8f4e5906/pgbrotli.control)

`pgbrotli` provides Brotli compression inside PostgreSQL. `pgbrotli_compress(text)` and `pgbrotli_compress(bytea)` return compressed `bytea`; `pgbrotli_decompress(bytea)` returns the original bytes.

```sql
CREATE EXTENSION pgbrotli;
WITH packed AS (
  SELECT pgbrotli_compress('hello world'::text) AS value
)
SELECT convert_from(pgbrotli_decompress(value), 'UTF8')
FROM packed;
```

Decompression occurs in the database backend, and a small input can expand to a very large output. Reject untrusted or oversized payloads before calling the function, apply statement and resource limits, and avoid repeated decompression in scans. Use the `bytea` result deliberately; `convert_from` fails when bytes are not valid in the chosen encoding.

This abandoned 1.0 project has no current PostgreSQL/Brotli compatibility matrix or documented wire-format contract beyond its source and tests. Verify malformed input, truncated streams, expansion limits, empty values, architecture, library ABI, memory exhaustion, and parallel execution. PostgreSQL TOAST may already compress large values, so measure end-to-end storage and CPU before adding application-level Brotli.
