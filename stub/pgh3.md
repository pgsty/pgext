## Usage

Sources:

- [Official README](https://github.com/dlr-eoc/pgh3/blob/a225ac6da928cb0668780221da64061d3c4b1d54/README.md)
- [SQL API reference](https://github.com/dlr-eoc/pgh3/blob/a225ac6da928cb0668780221da64061d3c4b1d54/doc/pgh3.md)
- [Extension SQL](https://github.com/dlr-eoc/pgh3/blob/a225ac6da928cb0668780221da64061d3c4b1d54/sql/pgh3.sql)

`pgh3` 0.3.0 is an archived binding between PostGIS geometries and an older Uber H3 C API. It represents H3 indexes as `text` and provides point indexing, cell geometry, hierarchy, neighborhood, polygon filling, compaction, and average-size helpers.

### Point and Cell Workflow

Install PostGIS first. Coordinates are interpreted as longitude in X and latitude in Y, in degrees; the wrapper does not transform or validate a geometry's SRID.

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pgh3;

WITH cell AS (
    SELECT h3_geo_to_h3index(
        ST_SetSRID(ST_Point(9.4069, 52.1233), 4326),
        5
    ) AS h3index
)
SELECT h3index,
       h3_get_resolution(h3index),
       h3_h3index_to_geo(h3index),
       h3_h3index_to_geoboundary(h3index)
FROM cell;
```

Transform source geometries to longitude/latitude before calling the extension. Returned PostGIS geometries are constructed without preserving an SRID, so set the expected SRID explicitly where downstream spatial operations require it.

### Main Function Groups

- `h3_geo_to_h3index(point, integer)` and its `geometry` overload create an index.
- `h3_h3index_to_geo(text)` and `h3_h3index_to_geoboundary(text)` return cell center and boundary geometry.
- `h3_kring(text, integer)`, `h3_to_parent(text, integer)`, and `h3_to_children(text, integer)` navigate neighbors and hierarchy.
- `h3_compact(text[])` and `h3_uncompact(text[], integer)` change hierarchical representation.
- `h3_hexagon_area_km2(integer)`, `h3_hexagon_area_m2(integer)`, `h3_edge_length_km(integer)`, and `h3_edge_length_m(integer)` return average size measures.

### Polygon Filling and Memory

Estimate cardinality before materializing a fill, especially at fine resolutions.

```sql
WITH region AS (
    SELECT ST_GeomFromText(
        'POLYGON((9.35 52.08,9.46 52.08,9.46 52.17,9.35 52.17,9.35 52.08))',
        4326
    ) AS geom
)
SELECT h3_polyfill_estimate(geom, 8) FROM region;

WITH region AS (
    SELECT ST_GeomFromText(
        'POLYGON((9.35 52.08,9.46 52.08,9.46 52.17,9.35 52.17,9.35 52.08))',
        4326
    ) AS geom
)
SELECT h3_polyfill(geom, 8) FROM region;
```

`h3_polyfill` supports Polygon and MultiPolygon geometries and omits holes. It preallocates the estimated output. `pgh3.polyfill_mem` limits that allocation and defaults to PostgreSQL's approximately 1 GB `MaxAllocSize`; increasing it can create very large backend allocations. Prefer lower resolution or split input with PostGIS before raising the limit.

### Compatibility Boundary

The extension targets pre-H3-v4 function names and is tightly coupled to the H3 library ABI used at build time. Its repository is archived and does not establish support beyond the historical PostgreSQL range. Build against an explicitly matched H3 release, verify known cell IDs and geometry orientation, and prefer a maintained H3 extension for new production systems.
