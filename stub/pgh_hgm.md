


## Usage

Sources: [PGHydro README](https://github.com/pghydro/pghydro/blob/master/README.md), [pgh_hgm SQL](https://github.com/pghydro/pghydro/blob/master/pgh_hgm--2.2.6.sql)

`pgh_hgm` adds hydrogeomorphological analysis routines to PgHydro. It combines drainage-line, drainage-area, DEM, and derived raster products to calculate basin and river-network metrics used in water-resources analysis.

### Enable

Install it after the core PgHydro and raster extensions:

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_raster;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_raster;
CREATE EXTENSION pgh_hgm;
```

### Initialize Objects

The extension creates the `pgh_hgm` schema with tables such as `pghft_drn_elevationprofile`, `pghft_upn_elevationprofile`, and `pghft_hydro_intel`.

```sql
SELECT pgh_hgm.pghfn_tables_initialize();
```

### Metrics

The SQL objects include functions for elevation profiles, drainage-area elevation summaries, average length of overland flow, axis length, circularity, compacity, drainage density, elevation statistics, form factor, hydrodensity, Kirpich time, perimeter, reach gradient, relief ratio, shape factor, time of concentration, slope methods, sinuosity, wave travel, and AMHG depth/width estimates.

```sql
SELECT nspname, proname
FROM pg_proc p
JOIN pg_namespace n ON n.oid = p.pronamespace
WHERE nspname = 'pgh_hgm'
ORDER BY proname;
```

### Notes

Load and validate DEM, flow-direction, flow-accumulation, and PgHydro drainage layers before running the metrics. `pgh_hgm` depends on `pgh_raster` for raster-derived hydrological products.
