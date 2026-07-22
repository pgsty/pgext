## Usage

Sources:

- [Official README](https://github.com/siose-innova/pg_geohash_extra/blob/ea5fcbb319503599ef26f67bb165f99b39add763/README.md)
- [Installed SQL build list](https://github.com/siose-innova/pg_geohash_extra/blob/ea5fcbb319503599ef26f67bb165f99b39add763/Makefile)
- [Geometry-to-geohash SQL](https://github.com/siose-innova/pg_geohash_extra/blob/ea5fcbb319503599ef26f67bb165f99b39add763/sql/functions/geohashfromgeom.sql)
- [Neighbour implementation](https://github.com/siose-innova/pg_geohash_extra/blob/ea5fcbb319503599ef26f67bb165f99b39add763/src/geohash_extra.c)

`pg_geohash_extra` supplements PostGIS with functions for finding neighbouring geohashes and covering a geometry with geohash cells. It also installs helper grid/composite objects used by an alternate geometry-covering function.

### Core Workflow

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_geohash_extra;

SELECT * FROM geohash_neighbours('ezs42');

SELECT *
FROM geohashfromgeom(
    ST_MakeEnvelope(-3.72, 40.40, -3.68, 40.43, 4326),
    6
) AS hash;
```

`geohashfromgeom` transforms a non-4326 geometry to SRID 4326, covers its bounding range at the requested precision, and returns text geohashes. It is a cell-covering helper, not an exact geometry decomposition.

### Object Index

- `geohash_neighbours(text)` returns the eight adjacent geohashes around the input cell.
- `geohashfromgeom(geometry, integer)` returns text geohashes covering a geometry's range.
- `ST_RegularGrid(geometry, numeric, numeric, boolean)` returns grid points or cells as `T_Grid` rows.
- `st_geohashfromgeom(geometry, integer)` returns composite `geohash` rows with ID, precision, and geometry.
- `geohash` and `T_Grid` are composite helper types.

### Operational Notes

Version `0.0.1` requires `postgis`, is relocatable, and declares no preload requirement. Validate precision and input SRID, and separately test antimeridian, polar, empty, and very large geometries; the implementation is early-stage and the README lists substantial future work.

Review installation in a disposable database first. The extension script runs `DROP TYPE IF EXISTS ... CASCADE` for generic type names `geohash` and `T_Grid`, and changes `ST_RegularGrid` ownership to the hard-coded role `postgres`. Those actions can remove dependent objects, conflict with existing names, require elevated privileges, or fail where the `postgres` role does not exist.
