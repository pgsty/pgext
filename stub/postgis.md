## Usage

Sources:

- [PostGIS 3.7.0 release](https://postgis.net/2026/10/PostGIS-3.7.0/)
- [PostGIS 3.7 manual](https://postgis.net/docs/manual-3.7/)
- [Official release notes](https://postgis.net/docs/release_notes.html)
- [Current manual HTML](https://postgis.net/docs/postgis-en.html)

`postgis` 3.7.0 adds spatial types, indexes, and SQL functions to PostgreSQL. The main user-facing split is between `geometry` for planar/projected work and `geography` for spherical calculations on longitude/latitude data. This major release requires PostgreSQL 14 or later, GEOS 3.10 or later, and PROJ 6.1 or later.

### Basic setup

```sql
CREATE EXTENSION postgis;
SELECT PostGIS_Full_Version();
```

### Core types and functions

```sql
CREATE TABLE sensors (
  id bigserial PRIMARY KEY,
  geom geometry(Point, 4326),
  geog geography(Point, 4326)
);

SELECT ST_SetSRID(ST_MakePoint(-73.985, 40.748), 4326);
SELECT ST_Intersects(a.geom, b.geom) FROM a, b;
SELECT ST_DWithin(a.geom, b.geom, 100);
SELECT ST_Distance(a.geog, b.geog);
SELECT ST_Transform(geom, 3857) FROM sensors;
```

- constructors: `ST_MakePoint`, `ST_GeomFromText`, `ST_GeomFromGeoJSON`
- relationships: `ST_Intersects`, `ST_Contains`, `ST_Within`, `ST_DWithin`
- measurements and transforms: `ST_Distance`, `ST_Area`, `ST_Length`, `ST_Transform`
- processing: `ST_Buffer`, `ST_Intersection`, `ST_Union`

### Spatial indexes

```sql
CREATE INDEX idx_sensors_geom ON sensors USING GIST (geom);
```

The official manual continues to recommend GiST as the general-purpose spatial index, with BRIN and SP-GiST available for specific data distributions and tradeoffs.

### Caveats

- Use `geometry` in an appropriate projected SRID for planar distances and areas; use `geography` when you need meter-based spheroidal calculations.
- `PostGIS 3.7.0` is a major release. It drops PostgreSQL 12/13 and GEOS versions older than 3.10; review the topology tolerance changes and the move of `address_standardizer` and `tiger_geocoder` packaging before upgrading. After installing matching files, run `SELECT postgis_extensions_upgrade();` where the upstream upgrade procedure requires it.
