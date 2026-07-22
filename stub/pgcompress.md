## Usage

Sources:

- [Official README](https://github.com/chet0xhenry/pgcompress/blob/8264107a26c07b6415f75c4a8eb84f9d27e77ea3/README.md)
- [Extension SQL API](https://github.com/chet0xhenry/pgcompress/blob/8264107a26c07b6415f75c4a8eb84f9d27e77ea3/pgcompress--1.0.sql)
- [Compression implementation](https://github.com/chet0xhenry/pgcompress/blob/8264107a26c07b6415f75c4a8eb84f9d27e77ea3/pgcompress.c)

`pgcompress` 1.0 compresses `text`, `bytea`, `json`, and `jsonb` values into `bytea` using zlib, raw DEFLATE, gzip, or Brotli, and provides matching decode functions. It is an old, abandoned C extension; use it only when an application needs one of these external wire formats, not as a replacement for PostgreSQL TOAST compression.

### Core Workflow

The generic interface uses zlib by default. When an algorithm code is supplied, use the same code to decode:

```sql
CREATE EXTENSION pgcompress;

SELECT decode_compressed(encode_compressed('hello PostgreSQL'));

SELECT decode_compressed(
  encode_compressed('hello PostgreSQL', 2, 6),
  2
);

SELECT decode_compressed_jsonb(
  encode_compressed('{"ok":true}'::jsonb, 3, 5),
  3
);
```

Algorithm codes are `0` for zlib, `1` for raw zlib/DEFLATE, `2` for gzip, and `3` for Brotli. zlib/gzip levels are 0–9; Brotli levels are 0–11. Out-of-range types fall back to zlib, and invalid levels fall back to the library default, so validate parameters before storing data rather than relying on silent fallback.

### Function Families

`encode_compressed` and `decode_compressed*` select the algorithm by code. Explicit families are `encode_zlib`/`decode_zlib*`, `encode_zlib_raw`/`decode_zlib_raw*`, `encode_gzip`/`decode_gzip*`, and `encode_br`/`decode_br*`. Suffixes such as `_bytea`, `_json`, and `_jsonb` select the decoded PostgreSQL return type; encoded data is always `bytea`.

### Operational Boundaries

The module must be linked against zlib and Brotli. Compression runs synchronously in the calling backend and adds CPU plus memory allocation; PostgreSQL may then TOAST-compress the resulting value independently. Record the algorithm alongside every payload because compressed bytes do not carry a PostgreSQL-level type contract, and using the wrong decoder raises an error or produces unusable data.

Do not decompress untrusted, unbounded payloads without application size limits: highly compressed input can expand enough to exhaust backend memory. The repository does not establish compatibility with current PostgreSQL majors and uses server-internal C APIs, so rebuild and test against the exact target server before relying on stored values. Prefer built-in column compression when the goal is merely transparent database storage reduction.
