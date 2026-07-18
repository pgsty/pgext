## Usage

Sources:

- [Version 1.0 SQL objects](https://github.com/amutu/fnvhash/blob/f536e411fa8b23cca1fd664e017c80ee6df42448/fnvhash--1.0.sql)
- [Wrapper implementation](https://github.com/amutu/fnvhash/blob/f536e411fa8b23cca1fd664e017c80ee6df42448/fnvhash.c)
- [Extension control file](https://github.com/amutu/fnvhash/blob/f536e411fa8b23cca1fd664e017c80ee6df42448/fnvhash.control)

`fnvhash` exposes Fowler–Noll–Vo hash variants for `varchar`, `text`, and `bytea`. The 32-bit functions `fnv032`, `fnv132`, and `fnv1a32` return `bigint`; the 64-bit counterparts `fnv064`, `fnv164`, and `fnv1a64` return a textual representation.

```sql
CREATE EXTENSION fnvhash;
SELECT fnv1a32('example'::text),
       fnv1a64('example'::text),
       fnv1a32(E'\\x0001'::bytea);
```

FNV is a fast non-cryptographic hash. Collisions are expected, especially in 32 bits; never use these values for passwords, authentication, signatures, tamper detection, or as a uniqueness guarantee. Preserve the original value and use a database constraint when identity matters.

The upstream README contains no compatibility or byte-order contract beyond the code. Before using hashes for partitioning, cross-system joins, or persisted protocol fields, pin the implementation and verify exact outputs across architecture, encoding, PostgreSQL version, and client representation. Changing algorithm or return formatting requires a coordinated data migration.
