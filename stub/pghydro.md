## Usage

> Sources: [GitHub repo](https://github.com/pghydro/pghydro), [README](https://raw.githubusercontent.com/pghydro/pghydro/master/README.md), [releases](https://github.com/pghydro/pghydro/releases)
> Lead extension for the PgHydro suite.

PgHydro extends PostGIS and PostgreSQL for drainage network analysis and water-resources decision making. The project covers river network modeling, flow-direction analysis, Otto Pfafstetter basin coding, upstream and downstream stretch selection, distance-to-mouth calculations, upstream area analysis, river orders, and basin levels.

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_raster;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_raster;
CREATE EXTENSION pgh_hgm;
CREATE EXTENSION pgh_consistency;
CREATE EXTENSION pgh_output;
```

### Components

- `pghydro` is the core drainage-network analysis extension.
- `pgh_raster` uses hydrological products derived from a digital elevation model.
- `pgh_hgm` combines `pghydro` and `pgh_raster` for hydrogeomorphological analysis.
- `pgh_output` provides reporting objects.
- `pgh_consistency` adds Pfafstetter consistency checks.

### Requirements

- PostgreSQL 9.1 or newer.
- PostGIS 3.x.
- PostGIS Raster.

### Notes

- The README says the master branch tracks the latest minor release, 6.6.
- The CSV lead row is the core `pghydro` package, but the repository also ships companion extensions in the same release tree.
