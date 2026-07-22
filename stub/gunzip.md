## Usage

Sources:

- [Version 0.2 SQL objects](https://gitlab.com/nicklasaven/pg_gunzip/-/blob/c46688ee9dd6e23702b631a8e0b08964df418b94/gunzip--0.2.sql)
- [C implementation](https://gitlab.com/nicklasaven/pg_gunzip/-/blob/c46688ee9dd6e23702b631a8e0b08964df418b94/gunzip.c)
- [Extension control file](https://gitlab.com/nicklasaven/pg_gunzip/-/blob/c46688ee9dd6e23702b631a8e0b08964df418b94/gunzip.control)

`gunzip` exposes in-database decompression for gzip-compressed byte strings. Version 0.2 defines `gunzip_text(bytea)` for text output, `gunzip_bytea(bytea)` for binary output, and `gunzip(bytea)` as an alias of the text function.

### Core workflow

```sql
CREATE EXTENSION gunzip;
SELECT gunzip_text(compressed_payload),
       gunzip_bytea(compressed_payload)
FROM archive_item
WHERE id = 42;
```

### Data and backend safety

Neither output function is binary-safe. Both use the C `strlen` result as the PostgreSQL datum length, so the first zero byte truncates the result; `gunzip_bytea(bytea)` therefore does not preserve arbitrary decompressed bytes. Accept only known NUL-free payloads and validate the result outside this extension.

The SQL functions are not strict while the C code does not check for a SQL `NULL` argument. Empty or malformed input returns `NULL` after a notice, allocation failure calls `exit`, and decompression grows a backend-local buffer without an output bound. A bad call can therefore terminate its backend, while a small compressed value can exhaust memory. Do not expose these functions to untrusted input. Version 0.2 is old and the catalog records PostgreSQL 10 coverage only; production use requires fixing these defects and testing the exact PostgreSQL and zlib versions.
