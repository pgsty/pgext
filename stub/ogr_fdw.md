## Usage

Sources:

- [ogr_fdw v1.1.9 README](https://github.com/pramsey/pgsql-ogr-fdw/blob/v1.1.9/README.md)
- [Extension control file](https://github.com/pramsey/pgsql-ogr-fdw/blob/v1.1.9/ogr_fdw.control)
- [v1.1.8 to v1.1.9 comparison](https://github.com/pramsey/pgsql-ogr-fdw/compare/v1.1.8...v1.1.9)
- [GDAL vector drivers](https://gdal.org/en/stable/drivers/vector/index.html)

`ogr_fdw` exposes vector data supported by GDAL/OGR as PostgreSQL foreign tables. It can read files and remote data sources through OGR drivers and can write when the selected driver and data source support updates. Install PostGIS before `ogr_fdw` for native geometry columns; otherwise geometry is exposed as WKB `bytea`.

### Discover and Import a Layer

Use the installed helper to inspect a source and generate matching SQL:

```console
ogr_fdw_info -s /srv/gis/cities.gpkg
ogr_fdw_info -s /srv/gis/cities.gpkg -l cities
```

A minimal equivalent definition is:

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION ogr_fdw;

CREATE SERVER city_source
  FOREIGN DATA WRAPPER ogr_fdw
  OPTIONS (
    datasource '/srv/gis/cities.gpkg',
    format 'GPKG'
  );

CREATE FOREIGN TABLE city (
  fid bigint,
  geom geometry,
  name text
) SERVER city_source
  OPTIONS (layer 'cities');

SELECT fid, name FROM city WHERE geom && ST_MakeEnvelope(-10, 35, 30, 60, 4326);
```

The PostgreSQL server account needs filesystem permissions for file-backed data sources and network/credential access for remote drivers.

### Import and Mapping

```sql
CREATE SCHEMA gis_import;

IMPORT FOREIGN SCHEMA ogr_all
  LIMIT TO (cities)
  FROM SERVER city_source
  INTO gis_import;
```

`ogr_all` means all OGR layers. Import normally launders table and column names; use `launder_table_names` and `launder_column_names` options when exact remote names are required. A foreign column can map to a different source name with `OPTIONS (column_name 'RemoteName')`.

### Important Options and Objects

- Server options: required `datasource`, optional `format`, `updateable`, `config_options`, `open_options`, and `character_encoding`.
- Table options: `layer` identifies the OGR layer and `updateable` can disable writes.
- `fid` identifies a feature and is required for writable foreign tables.
- `ogr_fdw_info` lists drivers and layers and emits server/table definitions.
- `ogr_fdw_version()` reports the extension and GDAL version.
- `ogr_fdw_drivers()` lists the compiled OGR drivers.

### Performance and Write Boundaries

Simple comparisons and bounding-box `&&` predicates can be pushed down, but more complex filters may be evaluated locally. The FDW retrieves all selected source columns and opens two OGR connections per query rather than pooling them. Use `EXPLAIN`, project only needed columns, and benchmark the actual driver and data source.

Writes depend on driver capability and require source-level write permissions plus `fid`. Set `updateable = false` when a source must remain read-only. Version 1.1.9 simplifies the version string relative to 1.1.8 and has no documented SQL workflow change; the control file remains at SQL extension version 1.1.
