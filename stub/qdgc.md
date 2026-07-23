## Usage

Sources:

- [PGXN qdgc 0.1.0 release](https://pgxn.org/dist/qdgc/0.1.0/)
- [Official 0.1.0 README](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/README.md)
- [Official qdgc control file](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/qdgc.control)
- [Official qdgc 0.1.0 extension SQL](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/qdgc--0.1.0.sql)

`qdgc` 0.1.0 is the trusted, relocatable, pure-SQL core of the QDGC extension family. It encodes longitude and latitude as Extended Quarter Degree Grid Cell codes, decodes their bounds, navigates the prefix hierarchy, reports level metrics, and fills longitude/latitude bounding boxes. It has no PostGIS or compiled-library dependency; geometry, geography, and polygon-fill operations belong to the companion `qdgc_postgis` extension.

### Core Workflow

```sql
CREATE EXTENSION qdgc;

-- qdgc_encode uses (longitude, latitude, level).
SELECT qdgc_encode(31.4, 2.7, 5);
-- E031N02ADBAC

-- The h3-style alias reverses the coordinate arguments.
SELECT qdgc_latlng_to_cell(2.7, 31.4, 5);

SELECT *
FROM qdgc_cell_to_bounds('E031N02ADBAC');

SELECT qdgc_cell_to_parent('E031N02ADBAC', 3);
SELECT * FROM qdgc_cell_to_children('E031N02AD', 5);
```

QDGC hierarchy is encoded directly in the text: a child code begins with its parent code. That makes prefix filtering useful for rollups and descendant lookups:

```sql
CREATE TABLE observations (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    qdgc_code text NOT NULL
);

CREATE INDEX observations_qdgc_idx ON observations (qdgc_code);

SELECT qdgc_cell_to_parent(qdgc_code, 3) AS level_3_cell,
       count(*)
FROM observations
GROUP BY 1;

SELECT *
FROM observations
WHERE qdgc_code LIKE 'E031N02AB%';
```

### Bounding Boxes and Level Metrics

The core extension can enumerate rectangular coverage without PostGIS. Pass `min_lon > max_lon` for a box that crosses the antimeridian.

```sql
SELECT qdgc_bbox_cell_count(30.0, 1.0, 32.0, 3.0, 7);

SELECT *
FROM qdgc_bbox_to_cells(30.0, 1.0, 32.0, 3.0, 7);

SELECT qdgc_level_degrees(7);
SELECT qdgc_get_num_cells(7);
SELECT qdgc_average_cell_area(7, 2.0, 'km^2');
SELECT qdgc_version();
```

`qdgc_average_cell_area` is a spherical estimate. Use `qdgc_cell_area_km2` from `qdgc_postgis` when a cell-specific WGS84 spheroid measurement is required.

### Important Objects

- `qdgc_encode(lon, lat, level)` and `qdgc_latlng_to_cell(lat, lng, level)` create codes; the argument order is intentionally different.
- `qdgc_is_valid_cell`, `qdgc_get_level`, `qdgc_cell_to_bounds`, `qdgc_cell_to_lonlat`, and `qdgc_cell_to_latlng` inspect or decode a code.
- `qdgc_cell_to_parent` and `qdgc_cell_to_children` navigate the four-way prefix hierarchy.
- `qdgc_bbox_to_cells` enumerates cells meeting a bounding box, while `qdgc_bbox_cell_count` calculates the count without materializing the set.
- `qdgc_level_degrees`, `qdgc_get_num_cells`, and `qdgc_average_cell_area` report grid-level metrics.

### Operational Notes

- Upstream requires PostgreSQL 13 or newer and tests PostgreSQL 13 through 17. PostgreSQL 18 is not part of the published 0.1.0 test matrix.
- The control file sets `trusted = true` and `relocatable = true`. No `shared_preload_libraries`, `LOAD`, server restart, or native library is required.
- Relocatable functions call one another by unqualified name. Install `qdgc` into a schema on the active `search_path`; the default `public` schema satisfies this boundary.
- Coordinates are longitude/latitude degrees. `qdgc_encode` takes longitude first, while `qdgc_latlng_to_cell` takes latitude first.
- Result cardinality grows by four for every additional child level. Count a bounding-box fill before materializing it, and avoid requesting deep descendants without a deliberate result-size bound.
