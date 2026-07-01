


## Usage

Sources: [PGHydro README](https://github.com/pghydro/pghydro/blob/master/README.md), [pgh_raster SQL](https://github.com/pghydro/pghydro/blob/master/pgh_raster--6.6.sql)

`pgh_raster` adds raster-based hydrological products to PgHydro. It stores DEM, flow-path, flow-direction, and flow-accumulation rasters, then exposes functions for upstream/downstream pixel tracing, drainage-area analysis, elevation lookup, elevation profiles, and raster clipping.

### Enable

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_raster;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_raster;
```

### Raster Tables

The extension creates the `pgh_raster` schema with tables such as:

| Table | Purpose |
|-------|---------|
| `pghrt_elevation` | Elevation raster data |
| `pghrt_flowpath` | Flow-path raster data |
| `pghrt_flowdirection` | Flow-direction raster data |
| `pghrt_flowaccumulation` | Flow-accumulation raster data |

Load raster products with PostGIS Raster tooling such as `raster2pgsql`, using the SRID and tiling scheme that matches the drainage dataset.

### Analysis Functions

Function groups include downstream and upstream pixel traversal, upstream area lookup, downstream drainage point and area lookup, elevation pixel values, elevation profiles, and raster clipping.

```sql
SELECT nspname, proname
FROM pg_proc p
JOIN pg_namespace n ON n.oid = p.pronamespace
WHERE nspname = 'pgh_raster'
ORDER BY proname;
```

### Notes

`pgh_raster` depends on both `pghydro` and `postgis_raster`. Keep raster alignment, SRID, and nodata handling consistent with the vector drainage layers before deriving hydrogeomorphological metrics.
