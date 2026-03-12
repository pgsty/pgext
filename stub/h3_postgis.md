

## Usage

> [h3_postgis: PostGIS integration for H3](https://github.com/zachasme/h3-pg)

`h3_postgis` is a bridge extension that integrates the H3 hexagonal hierarchical spatial index with PostGIS. It enables conversion between H3 indexes and PostGIS geometry types.

```sql
CREATE EXTENSION h3_postgis CASCADE;
```

This extension requires both `h3` and `postgis` to be installed. It provides functions to convert between H3 cell indexes and PostGIS geometries, enabling spatial queries that combine H3's hexagonal grid system with PostGIS's spatial capabilities.

### Key Functions

```sql
-- Convert a PostGIS point to an H3 cell index
SELECT h3_latlng_to_cell(ST_MakePoint(-73.985, 40.748)::point, 9);

-- Get the boundary of an H3 cell as a PostGIS geometry
SELECT h3_cell_to_boundary_geometry('892a1008003ffff'::h3index);

-- Convert H3 cells to PostGIS polygons for visualization
SELECT h3_cell_to_geometry('892a1008003ffff'::h3index);
```
