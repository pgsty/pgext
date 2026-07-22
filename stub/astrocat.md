## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/karpov-sv/astrocat/blob/5c714e05f7ee11f7cf61c482c189a919ec37edd9/README.md)
- [Version 1.0 SQL objects](https://github.com/karpov-sv/astrocat/blob/5c714e05f7ee11f7cf61c482c189a919ec37edd9/astrocat--1.0.sql)
- [GiST implementation](https://github.com/karpov-sv/astrocat/blob/5c714e05f7ee11f7cf61c482c189a919ec37edd9/astrocat.c)

`astrocat` is a testbed for indexing astronomical catalog points. It supports radial cone searches and k-nearest-neighbor ordering with a functional GiST index, using right ascension and declination expressed in degrees.

```sql
CREATE EXTENSION astrocat;
CREATE TABLE star (ra double precision, dec double precision);
CREATE INDEX star_sky_idx ON star USING gist (skypoint(ra, dec));
SELECT *
FROM star
WHERE skypoint(ra, dec) @ skycircle(11.2, 67.0, 2.0)
ORDER BY skypoint(ra, dec) <-> skypoint(11.2, 67.0)
LIMIT 10;
```

The project describes itself as experimental and is heavily based on pgSphere. Its constructors and text input routines do not validate parse success or enforce right-ascension, declination, and radius ranges. Equality compares the internal floating-point representation byte for byte. Validate inputs before constructing values, and do not assume mathematically equivalent coordinates compare equal.

The reviewed GiST code has correctness defects. Direct `@` containment includes a point exactly on the circle boundary, but the leaf consistency function uses a strict comparison and sets `recheck = false`, so an indexed scan can omit that point. The internal-node KNN distance code compares the query's y coordinate with the key's x upper bound in one branch, which can invalidate nearest-neighbor ordering. Compare indexed and sequential cone-search results and compare KNN output with an independently sorted spherical-distance calculation; exact-boundary tests are mandatory.

There is no current compatibility or numerical-accuracy matrix. Do not use this operator class for authoritative results without fixing and independently validating it. If the index-key representation, containment, or distance code changes, rebuild every dependent index. Prefer maintained pgSphere or another supported spatial astronomy stack for production.
