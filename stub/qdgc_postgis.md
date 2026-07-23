## Usage

Sources:

- [PGXN qdgc 0.1.0 release](https://pgxn.org/dist/qdgc/0.1.0/)
- [Official 0.1.0 README](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/README.md)
- [Official qdgc_postgis control file](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/qdgc_postgis.control)
- [Official qdgc_postgis 0.1.0 extension SQL](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/qdgc_postgis--0.1.0.sql)

`qdgc_postgis` 0.1.0 is the PostGIS companion to the pure-SQL `qdgc` core. It converts QDGC cells to and from PostGIS points and polygons, measures cell area on the WGS84 spheroid, and fills arbitrary geometries with QDGC cells. The extension requires both `qdgc` and `postgis`; it does not replace either one.

### Core Workflow

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION qdgc;
CREATE EXTENSION qdgc_postgis;

SELECT qdgc_latlng_to_cell(
    ST_SetSRID(ST_MakePoint(31.4, 2.7), 4326),
    5
);

SELECT qdgc_cell_to_geometry('E031N02ADBAC');
SELECT qdgc_cell_to_boundary_geometry('E031N02ADBAC');
SELECT qdgc_cell_area_km2('E031N02ADBAC');
```

The point overload transforms geometry with a nonzero, non-4326 SRID to EPSG:4326. An SRID of zero is assumed to already contain longitude and latitude.

### Fill an Area of Interest

Estimate the result size before producing a deep fill:

```sql
WITH area AS (
    SELECT ST_GeomFromText(
        'POLYGON((31.0 2.0, 31.5 2.0, 31.5 2.5, 31.0 2.5, 31.0 2.0))',
        4326
    ) AS geom
)
SELECT qdgc_estimate_cell_count(geom, 7)
FROM area;

WITH area AS (
    SELECT ST_GeomFromText(
        'POLYGON((31.0 2.0, 31.5 2.0, 31.5 2.5, 31.0 2.5, 31.0 2.0))',
        4326
    ) AS geom
)
SELECT cell
FROM area
CROSS JOIN LATERAL qdgc_polygon_to_cells(
    geom,
    7,
    'intersects'
) AS cell;
```

The predicate can be:

- `intersects`, the default, for cells intersecting the geometry;
- `centroid`, for cells whose center lies inside the geometry;
- `contains`, for cells wholly contained by the geometry.

The implementation descends a pruning quadtree instead of testing every cell in the geometry's full envelope. Multi-part geometries are filled per part and their cell sets are combined.

### Important Objects

- `qdgc_latlng_to_cell(geometry, level)` and its `geography` overload encode PostGIS points.
- `qdgc_cell_to_geometry` and `qdgc_cell_to_geography` return the cell centroid.
- `qdgc_cell_to_boundary_geometry` and `qdgc_cell_to_boundary_geography` return the rectangular cell boundary.
- `qdgc_cell_area_km2` measures the boundary geography on the WGS84 spheroid.
- `qdgc_polygon_to_cells` fills an area using one of the three documented predicates.
- `qdgc_estimate_cell_count` provides a cheap, envelope-capped guard before materializing a fill.

### Operational Notes

- `qdgc_postgis.control` declares `requires = 'qdgc,postgis'` and `relocatable = true`. Install PostGIS with an appropriately privileged role before delegating use of the companion extension.
- No `shared_preload_libraries`, `LOAD`, or restart is required. The extension is SQL-only, but its PostGIS dependency includes native code.
- Install `qdgc`, `qdgc_postgis`, and their callable dependencies into schemas visible on the active `search_path`, because the relocatable SQL calls functions by unqualified name.
- Upstream tests PostgreSQL 13 through 17. Do not infer PostgreSQL 18 support from the absence of compiled code.
- Deep area fills can still produce enormous sets even with pruning. Treat `qdgc_estimate_cell_count` as an operational guard and apply application-specific limits before executing `qdgc_polygon_to_cells`.

