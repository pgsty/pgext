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

The project describes itself as experimental and is heavily based on pgSphere. Confirm coordinate units, right-ascension wraparound, declination bounds, spherical distance, null behavior, and edge cases near poles before trusting results.

There is no current compatibility or numerical-accuracy matrix. Compare indexed and sequential results over a known reference catalog, test GiST recheck behavior and concurrent updates, and benchmark both selectivity and KNN ordering. If index-key representation or distance semantics change, rebuild every dependent index. Prefer maintained pgSphere or another supported spatial astronomy stack for production.
