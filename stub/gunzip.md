## Usage

Sources:

- [Version 0.2 SQL objects](https://gitlab.com/nicklasaven/pg_gunzip/-/blob/master/gunzip--0.2.sql)
- [C implementation](https://gitlab.com/nicklasaven/pg_gunzip/-/blob/master/gunzip.c)
- [Extension control file](https://gitlab.com/nicklasaven/pg_gunzip/-/blob/master/gunzip.control)

`gunzip` exposes in-database decompression for gzip-compressed byte strings. Version 0.2 defines `gunzip_text(bytea)` for text output, `gunzip_bytea(bytea)` for binary output, and `gunzip(bytea)` as an alias of the text function.

```sql
CREATE EXTENSION gunzip;
SELECT gunzip_text(compressed_payload),
       gunzip_bytea(compressed_payload)
FROM archive_item
WHERE id = 42;
```

Use `gunzip_bytea(bytea)` when decompressed bytes may not be valid database text. The extension performs decompression inside a backend, so input size alone does not bound output size: reject untrusted or oversized compressed values before calling it, and enforce statement and resource limits to reduce decompression-bomb risk.

The upstream SQL labels these functions immutable, but failures and memory pressure still affect the executing statement. Version 0.2 is old and the catalog records PostgreSQL 10 coverage only. Validate build compatibility, malformed-stream handling, encoding behavior, and maximum expansion on the exact PostgreSQL and zlib versions before production use.
