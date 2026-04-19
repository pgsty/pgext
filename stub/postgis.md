## Usage

Source: [Official manual](https://postgis.net/documentation/manual/), [current manual HTML](https://postgis.net/docs/postgis-en.html), [release notes](https://postgis.net/docs/release_notes.html), [patch release announcement](https://postgis.net/2026/04/PostGIS-Patch-Releases/)

`postgis` adds spatial types, indexes, and SQL functions to PostgreSQL. The main user-facing split is between `geometry` for planar/projected work and `geography` for spherical calculations on longitude/latitude data.

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
- `PostGIS 3.6.3` is a patch release dated 2026-04-14. The release notes describe fixes and a security hardening change, not a new stub-level usage surface, so this refresh mostly trims and aligns the stub with the current manual.
