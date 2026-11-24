## Usage

Sources:

- [PostGIS MEST README](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/contrib/postgis/README.md)
- [Control](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/contrib/postgis/postgis_mest.control)
- [SQL 1.0](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/contrib/postgis/postgis_mest--1.0.sql)

`postgis_mest` 1.0 adds multi-entry R-tree indexing for PostGIS `geometry`. Geometry collections contribute both a whole-geometry bounding box and per-element boxes. Upstream requires PostgreSQL 17 or 18, PostGIS 3 and `mest`.

### Enable and Index

Upstream requires preloading the installed PostGIS library. Add it to the existing setting and restart PostgreSQL; the documented library name is:

```ini
shared_preload_libraries = 'postgis-3'
```

After installing the extension libraries, use a privileged database session:

```sql
CREATE EXTENSION postgis_mest CASCADE;
CREATE TABLE regions (id bigint, geom geometry);
CREATE INDEX regions_mgist_geom ON regions USING mgist (geom);
```

The control file requires `postgis` and `mest`. `CASCADE` installs available dependencies but does not install server library files.

### Query Surface

`geometry_mgist_ops_2d` is the default `geometry` class for `mgist`. Its useful query operators include bounding-box overlap `&&` and distance ordering `<->`.

```sql
SELECT id
FROM regions
WHERE geom && ST_MakeEnvelope(0, 0, 10, 10, 4326)
  AND ST_Intersects(geom, ST_MakeEnvelope(0, 0, 10, 10, 4326));
```

Use a query geometry with the same spatial reference as the stored data.

### Limits

PostGIS planner support functions are not implemented for this access method. A predicate containing only `ST_Intersects` or `ST_Contains` will not select the MGiST index through that support mechanism; add an appropriate explicit overlap condition when applicable. Benefits depend on collection shape and query selectivity. Single geometries generally contribute one bounding box.
