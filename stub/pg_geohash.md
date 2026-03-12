

## Usage

> [pg_geohash: Geohash functions for PostgreSQL](https://github.com/jistok/pg_geohash)

C-based geohash functions for PostgreSQL (also supports HAWQ and Greenplum). Based on the [libgeohash](https://github.com/lyokato/libgeohash) C library.

Background on geohash: [Wikipedia: Geohash](http://en.wikipedia.org/wiki/Geohash)

### Functions

Encode latitude/longitude to a geohash string with specified precision:

```sql
SELECT LAT_LON_TO_GEOHASH_WITH_LEN(latitude, longitude, 5) AS geohash;
```

Encode latitude/longitude to a full-precision geohash:

```sql
SELECT LAT_LON_TO_GEOHASH(latitude, longitude) AS geohash;
```

Decode a geohash back to latitude/longitude:

```sql
SELECT GEOHASH_TO_LAT_LON('dp3w7') AS lat_lon;
```

### Example

Compute geohash-based aggregates using 5-character precision (~2.4km x 4.9km cells):

```sql
SELECT LAT_LON_TO_GEOHASH_WITH_LEN(latitude, longitude, 5) AS geohash,
       COUNT(*)
FROM crimes
GROUP BY 1
ORDER BY 2 DESC
LIMIT 10;
```

```
 geohash | count
---------+-------
 dp3w7   | 72404
 dp3tt   | 70713
 dp3tw   | 63642
 dp3wm   | 62332
 dp3wk   | 56467
```

Recover coordinates from a geohash:

```sql
SELECT location,
       GEOHASH_TO_LAT_LON(LAT_LON_TO_GEOHASH(latitude, longitude))
FROM crimes
LIMIT 5;
```
