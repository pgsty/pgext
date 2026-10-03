## Usage

Sources:

- [4.5.0 PostGIS API](https://github.com/postgis/h3-pg/blob/v4.5.0/docs/api.md)
- [Dependencies and extension definition](https://github.com/postgis/h3-pg/blob/v4.5.0/h3_postgis/CMakeLists.txt)
- [4.5.0 migration SQL](https://github.com/postgis/h3-pg/blob/v4.5.0/h3_postgis/sql/updates/h3_postgis--4.2.3--4.5.0.sql)
- [4.5.0 release](https://github.com/postgis/h3-pg/releases/tag/v4.5.0)

`h3_postgis` bridges H3 cells with PostGIS geometry, geography, and raster. It requires `h3`, `postgis`, and `postgis_raster`, including for geometry-only use. Input geometries must use SRID 4326 with longitude, latitude coordinates; the functions do not reproject inputs.

### Convert Points and Cells

```sql
CREATE EXTENSION h3_postgis CASCADE;
SET h3.strict = true;

SELECT h3_latlng_to_cell(
    ST_SetSRID(ST_MakePoint(-122.0553238, 37.3615593), 4326), 9
);
SELECT h3_cell_to_geometry('85283473fffffff'::h3index);
SELECT h3_cell_to_boundary_geometry('85283473fffffff'::h3index);
```

Transform other coordinate systems to SRID 4326 with PostGIS before calling the H3 functions. Setting an SRID label alone does not transform coordinates.

### Core API

| Task | Functions |
| --- | --- |
| Point to cell | `h3_latlng_to_cell(geometry, integer)`, `h3_latlng_to_cell(geography, integer)` |
| Cell center | `h3_cell_to_geometry`, `h3_cell_to_geography` |
| Cell boundary | `h3_cell_to_boundary_geometry`, `h3_cell_to_boundary_geography` |
| Polygon coverage | `h3_polygon_to_cells`, `h3_cells_to_multi_polygon_geometry`, `h3_cells_to_multi_polygon_geography` |
| Continuous raster summaries | `h3_raster_summary`, `h3_raster_summary_stats_agg` |
| Categorical raster summaries | `h3_raster_class_summary`, `h3_raster_class_summary_item_agg` |

The geometry-to-resolution `@` operator also maps a location to an H3 cell. Validate polygon inputs with `ST_IsValid()`; repairs through `ST_MakeValid()` can change topology and produce geometry collections, so retain and review polygonal components before coverage calculations. Invalid polygons have undefined behavior.

### Summarize Raster Data

```sql
SELECT (summary).h3,
       (h3_raster_summary_stats_agg((summary).stats)).*
FROM (
    SELECT h3_raster_summary(rast, 8) AS summary
    FROM rasters
) AS r
GROUP BY (summary).h3;
```

The default summary chooses a method; explicit clip, centroid, and subpixel variants let you control how raster pixels are assigned to cells. Review the choice against raster resolution and the requested H3 resolution.

### Upgrade and Boundaries

```sql
ALTER EXTENSION h3 UPDATE TO '4.5.0';
ALTER EXTENSION h3_postgis UPDATE TO '4.5.0';
```

The base extension update performs affected btree rebuilds and distance-dependent refreshes, so plan its maintenance window before updating the companion. Version 4.5.0 fixes restricted-search-path maintenance on PostgreSQL 17+, expression-index dump/restore, and several geometry and polygonization errors. Both extension versions should match. Install the matching 4.5.0 package files before updating either extension.

`h3.extend_antimeridian` should normally remain false for planar overlays. Both extensions are relocatable; ensure the schemas containing H3 and PostGIS objects are visible when issuing unqualified SQL. Neither extension requires shared preloading.
