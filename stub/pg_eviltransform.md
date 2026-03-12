

## Usage

> [pg_eviltransform: Coordinate transform between WGS84, GCJ02, and BD09](https://github.com/aiyou178/pg_eviltransform)

`pg_eviltransform` extends PostGIS `ST_Transform` with BD09/GCJ02 support for Chinese coordinate systems. It exposes `ST_EvilTransform` with the same overload interface as `ST_Transform`.

Custom SRIDs:
- `990001`: GCJ02
- `990002`: BD09

### Functions

```sql
ST_EvilTransform(geometry, to_srid integer)
ST_EvilTransform(geometry, to_proj text)
ST_EvilTransform(geometry, from_proj text, to_srid integer)
ST_EvilTransform(geometry, from_proj text, to_proj text)
```

If neither side uses custom coordinates, it delegates directly to `ST_Transform`. If BD09/GCJ02 is involved, it transforms via WGS84 (`4326`) when needed.

### Examples

```sql
-- WGS84 to GCJ02 using text literal
SELECT ST_EvilTransform(ST_SetSRID('POINT(120 30)'::geometry, 4326), 'GCJ02');

-- WGS84 to BD09 using text literal
SELECT ST_EvilTransform(ST_SetSRID('POINT(120 30)'::geometry, 4326), 'BD09');

-- WGS84 to GCJ02 using numeric SRID
SELECT ST_EvilTransform(ST_SetSRID('POINT(120 30)'::geometry, 4326), 990001);

-- BD09 to Web Mercator
SELECT ST_EvilTransform(
  ST_SetSRID('POINT(120.011070620552 30.0038830555128)'::geometry, 990002), 3857
);

-- from_proj / to_proj overload
SELECT ST_EvilTransform('POINT(120 30)'::geometry, 'EPSG:4326', 'GCJ02');
```

### Performance

On PG18 with 200,000 rows, `ST_EvilTransform` is ~30-45x faster than the regex-based SQL approach.
