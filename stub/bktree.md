## Usage

Sources:

- [Official README](https://github.com/meniam/pg_bktree/blob/59db54ff5cd0b9f1b60c95be6f04414282699237/README.md)
- [Version 1.0 SQL API](https://github.com/meniam/pg_bktree/blob/59db54ff5cd0b9f1b60c95be6f04414282699237/bktree--1.0.sql)
- [Hamming-distance implementation](https://github.com/meniam/pg_bktree/blob/59db54ff5cd0b9f1b60c95be6f04414282699237/bktree.c)

`bktree` provides a BK-tree operator class for SP-GiST indexes over 64-bit integer hashes. Its distance is the Hamming distance between the two bit patterns, making it useful for nearest-neighbor-like filtering of perceptual hashes rather than ordinary numeric ranges.

### Core Workflow

Store each 64-bit hash as `bigint`, build the `bktree_ops` index, and query with a `bktree_area` containing the center and maximum distance.

```sql
CREATE EXTENSION bktree;
CREATE TABLE images (id bigint PRIMARY KEY, phash bigint NOT NULL);
CREATE INDEX images_phash_bktree ON images USING spgist (phash bktree_ops);

SELECT id, phash <-> 123456789::bigint AS distance
FROM images
WHERE phash <@ ROW(123456789::bigint, 8::bigint)::bktree_area;
```

The `<@` operator tests whether the Hamming distance is at most the requested radius. The `<->` operator returns the distance, while the extension's `=` operator tests bitwise equality. `int64_to_bitstring()` and `bitstring_to_int64()` are conversion helpers.

### Caveats

The operator class is only for `int8` and the implementation assumes a 64-bit Hamming distance from 0 through 64. A radius is not a numeric tolerance: nearby integer values may have many differing bits, and distant integer values may have few. The reviewed fork specifically targets PostgreSQL 15; validate index creation, plans, equality semantics, and signed hash conversion on the exact server build before production use.
