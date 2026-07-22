## Usage

Sources:

- [Official project description](https://github.com/nmandery/dumppoints/blob/96343c7e4eb76765da69737f79611fa7355c9e5b/README)
- [Extension control file](https://github.com/nmandery/dumppoints/blob/96343c7e4eb76765da69737f79611fa7355c9e5b/dumppoints.control)
- [Extension SQL](https://github.com/nmandery/dumppoints/blob/96343c7e4eb76765da69737f79611fa7355c9e5b/dumppoints--1.0.sql)

`dumppoints` version `1.0` installs a C implementation of the PostGIS `ST_DumpPoints(geometry)` function. Use it to expand every vertex of a geometry into PostGIS `geometry_dump` rows containing a path and point geometry.

### Core Workflow

PostGIS is a declared dependency and must be available first.

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION dumppoints;

SELECT (d).path, ST_AsText((d).geom)
FROM (
  SELECT ST_DumpPoints(
    ST_GeomFromText('POLYGON((0 0, 2 0, 2 2, 0 0))')
  ) AS d
) AS q;
```

The path array identifies the component, ring, and vertex position using PostGIS dump conventions; the geometry field is the corresponding point.

### Objects

- `ST_DumpPoints(geometry) returns setof geometry_dump`: emit one `geometry_dump` value for each input vertex.

### Compatibility Notes

The installation script executes `CREATE OR REPLACE FUNCTION` for the same signature exposed by PostGIS. Installing `dumppoints` therefore replaces that database function in the target extension schema rather than adding a separately named API. Test behavior against the exact PostGIS version in use, and account for extension ownership before upgrading or removing either extension. No preload is required.
