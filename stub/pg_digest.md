## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/ratbox/pg-digest/blob/44ebe696700423cdd64ccc8f1d8b9dc18728934a/README.md)
- [Version 1.0 SQL objects](https://github.com/ratbox/pg-digest/blob/44ebe696700423cdd64ccc8f1d8b9dc18728934a/pg_digest--1.0.sql)
- [Type implementation](https://github.com/ratbox/pg-digest/blob/44ebe696700423cdd64ccc8f1d8b9dc18728934a/pg_digest.c)

`pg_digest` defines the variable-length `digest` type for compact storage of hash digests, checksums, encryption keys, and other binary blobs. Its representation is binary-compatible with `bytea`, with implicit casts in both directions.

```sql
CREATE EXTENSION pg_digest;
CREATE TABLE file_upload (
  path text PRIMARY KEY,
  hash digest NOT NULL UNIQUE
);
INSERT INTO file_upload VALUES ('manual.pdf', '\x0123456789abcdef'::digest);
```

The extension stores bytes; it does not calculate, validate, size, or label a hash algorithm. Enforce the expected byte length separately, and use `pgcrypto` or application cryptography to compute the digest. Never treat a checksum or short/non-cryptographic digest as authentication.

Because casts with `bytea` are implicit, review overloaded expressions and indexes for unexpected coercion. Pin the binary/text representation across upgrades and test dump/restore and logical replication before using the type in durable schemas; `bytea` remains the more portable choice when the small storage distinction is not important.
