## Usage

Sources: [README](https://github.com/pghydro/pghydro/blob/master/README.md), [repo](https://github.com/pghydro/pghydro), [releases](https://github.com/pghydro/pghydro/releases)

`pghydro` is the core extension in the PgHydro suite for drainage-network analysis and water-resources decision support on top of PostgreSQL and PostGIS.

### Install the PgHydro suite

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_raster;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_raster;
CREATE EXTENSION pgh_hgm;
CREATE EXTENSION pgh_consistency;
CREATE EXTENSION pgh_output;
```

The upstream README presents these companion extensions together:

- `pghydro` for drainage-network analysis
- `pgh_raster` for DEM-derived hydrological products
- `pgh_hgm` for hydrogeomorphological analysis
- `pgh_consistency` for Pfafstetter consistency checks
- `pgh_output` for reporting objects

### What upstream says it covers

The README describes support for:

- flow-direction correction in river networks
- Otto Pfafstetter basin coding
- upstream and downstream stretch selection
- distance-to-mouth calculations
- upstream area calculations
- river orders and basin levels

### Requirements

- PostgreSQL 9.1+
- PostGIS 3.x
- PostGIS Raster

### Notes

- The current upstream README status section still says the master branch tracks release `6.6` and the develop branch tracks `6.7-dev`.
- The repository also publishes newer tags, but the user-facing README remains centered on the `6.6` installation and tutorial flow.
