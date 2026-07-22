## Usage

Sources:

- [Official README](https://github.com/decision-labs/a5pg/blob/2555ce912ae6d135b831b21385b6fc2d4bcff989/README.md)
- [Official control file](https://github.com/decision-labs/a5pg/blob/2555ce912ae6d135b831b21385b6fc2d4bcff989/a5pg.control)
- [Official build manifest](https://github.com/decision-labs/a5pg/blob/2555ce912ae6d135b831b21385b6fc2d4bcff989/Cargo.toml)

`a5pg` exposes A5 equal-area discrete global grid functions in PostgreSQL. Use it to convert longitude/latitude coordinates into hierarchical 64-bit cell identifiers for aggregation and mapping; it is not a PostGIS index access method, and the optional geometry wrapper is available only when a compatible `geometry` type exists.

### Core Workflow

```sql
CREATE EXTENSION a5pg;

SELECT a5_lonlat_to_cell(-73.9857, 40.7580, 10) AS cell_id;

SELECT a5_cell_to_lonlat(2742822465196523520) AS center,
       a5_get_resolution(2742822465196523520) AS resolution,
       a5_cell_to_parent(2742822465196523520, 8) AS parent;
```

Choose a resolution for the analytical scale, persist the returned `bigint` cell ID with source rows, and group or join on that value. Validate resolution changes explicitly: a parent cell is coarser, while `a5_cell_to_children` expands to finer cells.

### Important Objects

- `a5_lonlat_to_cell` converts longitude, latitude, and resolution to a cell ID.
- `a5_cell_to_lonlat` and `a5_cell_to_boundary` return the center and polygon coordinate arrays.
- `a5_get_resolution`, `a5_cell_to_parent`, and `a5_cell_to_children` navigate the hierarchy.
- `a5_point_to_cell` is a PostGIS convenience wrapper when `geometry` is already defined.
- `a5pg_version` and `a5pg_info` report extension information.

### Caveats

The control file makes version `0.6.1` non-relocatable, untrusted, and superuser-only to install. The manifest declares build features for PostgreSQL 13 through 18, while the reviewed README reports CI for 15 through 17; test the exact packaged server target.

Longitude/latitude order, coordinate reference system, cell-boundary handling, and chosen resolution are application semantics. Normalize inputs to WGS84 degrees, reject invalid coordinates, and regression-test joins near cell edges. A5 cell IDs are implementation identifiers rather than distances or globally sortable geographic coordinates.
