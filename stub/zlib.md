## Usage

Sources:

- [Official version 1.0 SQL](https://github.com/overplumbum/postgresql-zlib/blob/74f0e5126266105fea05dcd248400e73bc7bacc2/zlib--1.0.sql)
- [Official C implementation](https://github.com/overplumbum/postgresql-zlib/blob/74f0e5126266105fea05dcd248400e73bc7bacc2/zlib.c)
- [Official control file](https://github.com/overplumbum/postgresql-zlib/blob/74f0e5126266105fea05dcd248400e73bc7bacc2/zlib.control)
- [Official repository](https://github.com/overplumbum/postgresql-zlib)

`zlib` version `1.0` exposes one function that decompresses a zlib-wrapped `bytea` value into PostgreSQL `text`. It is useful for reading trusted application payloads already compressed with zlib; it does not compress data and does not implement gzip/archive handling.

### Core Workflow

Pass a complete zlib stream as `bytea`:

```sql
CREATE EXTENSION zlib;

SELECT zlib_decompress(
  decode('789ccb48cdc9c90700062c0215', 'hex')
);
```

The example returns `hello`. In application tables, retain the compressed column as `bytea` and call the function only when text output is actually needed.

### API

`zlib_decompress(bytea) RETURNS text` calls zlib's `uncompress`. It is declared `IMMUTABLE STRICT`: null input returns null, invalid streams raise an external-routine error containing zlib's diagnostic, and successful output is returned as database text.

### Caveats

There is no maximum decompressed size. The implementation starts with an output buffer six times the compressed size and repeatedly doubles it on `Z_BUF_ERROR`; a highly compressible or malicious payload can allocate large backend memory and cause denial of service. Restrict the function to trusted data or enforce compressed/output limits outside it.

The decompressed bytes must be valid in the database text encoding. The function accepts a zlib stream, not a gzip header, ZIP file, or arbitrary deflate fragment. The source dates from 2013 and has no published modern PostgreSQL compatibility matrix, so test the exact package/server build and malformed inputs before deployment.
