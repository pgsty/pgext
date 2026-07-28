## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/hexgrid/hexgrid-1.0.3/README.md)
- [Official extension control file (hexgrid.control)](https://api.pgxn.org/src/hexgrid/hexgrid-1.0.3/hexgrid.control)
- [Official extension SQL (hexgrid.sql)](https://api.pgxn.org/src/hexgrid/hexgrid-1.0.3/sql/hexgrid.sql)

`hexgrid` — Configurable hex grid on abstract surface. Use it for the corresponding spatial data or geospatial workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION hexgrid;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hex_Orientation(name text, f float[], b float[], start_angle float)` is an extension function and returns `hex_orientation`.
- `hex_OrientationFlat()` is an extension function and returns `hex_orientation`.
- `hex_OrientationPointy()` is an extension function and returns `hex_orientation`.
- `ST_Centroid(hexagon hexagon)` is an extension function and returns `geometry`.
- `ST_Hexagon(point geometry(point), grid_id int default 1)` is an extension function and returns `hexagon`.
- `ST_HexagonCoverage(region geometry, grid_id int default 1)` is an extension function and returns `setof`.
- `hex_orientation` is an extension-defined type.
- `hexagon` is an extension-defined type.
- `hexgrid` is an extension-defined type.
- `hexgrids` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.3`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
