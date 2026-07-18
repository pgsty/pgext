## Usage

Sources:

- [Upstream README and function reference](https://github.com/decision-labs/a5pg/blob/main/README.md)
- [Extension control file](https://github.com/decision-labs/a5pg/blob/main/a5pg.control)
- [Cargo package metadata](https://github.com/decision-labs/a5pg/blob/main/Cargo.toml)

`a5pg` exposes the equal-area [A5](https://a5geo.org/) discrete global grid in PostgreSQL. It converts longitude/latitude coordinates to 64-bit cell identifiers, recovers cell centers or boundaries, and navigates the cell hierarchy. The core API is compatible with the DuckDB A5 extension.

### Core operations

```sql
CREATE EXTENSION a5pg;

SELECT a5_lonlat_to_cell(-73.9857, 40.7580, 10);
SELECT a5_cell_to_lonlat(2742822465196523520);
SELECT a5_get_resolution(2742822465196523520);
SELECT a5_cell_to_parent(2742822465196523520, 8);
SELECT a5_cell_to_children(2742822465196523520, 12);
```

`a5_cell_to_boundary()` returns the boundary as a two-dimensional `double precision` array. If a PostGIS `geometry` type is already available, `a5_point_to_cell(geometry, resolution)` provides a convenience wrapper; PostGIS is not required for the core longitude/latitude functions.

The upstream package and this catalog both identify version `0.6.1`. Cargo features cover PostgreSQL 13 through 18. The extension does not require preloading.
