

## Usage

> [ogr_fdw: OGR Foreign Data Wrapper for PostgreSQL](https://github.com/pramsey/pgsql-ogr-fdw)

OGR is the **vector** half of the [GDAL](http://www.gdal.org/) spatial data access library. It allows access to a [large number of GIS data formats](http://www.gdal.org/ogr_formats.html) using a simple C API. Since OGR exposes a simple table structure and PostgreSQL foreign data wrappers allow access to table structures, the fit is pretty perfect.

### Quick Start

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION ogr_fdw;
```

Use the `ogr_fdw_info` tool to read an OGR data source and output server/table definitions:

```bash
ogr_fdw_info -s /tmp/test -l pt_two
```

```sql
CREATE SERVER "myserver"
  FOREIGN DATA WRAPPER ogr_fdw
  OPTIONS (
    datasource '/tmp/test',
    format 'ESRI Shapefile' );

CREATE FOREIGN TABLE "pt_two" (
  fid integer,
  "geom" geometry(Point, 4326),
  "name" varchar,
  "age" integer,
  "height" real,
  "birthdate" date )
  SERVER "myserver"
  OPTIONS (layer 'pt_two');

SELECT * FROM pt_two;
```

Filter pushdown is supported — both simple predicates and bounding box filters (`&&`):

```sql
SET client_min_messages = debug1;

SELECT name, age, height
FROM pt_two
WHERE height < 5.7
AND geom && ST_MakeEnvelope(0, 0, 1, 1);
```

```
DEBUG:  OGR SQL: (height < 5.7)
DEBUG:  OGR spatial filter (0 0, 1 1)
```


## Limitations

- PostgreSQL 11 or higher required
- Limited non-spatial query restrictions are pushed down to OGR (only `>`, `<`, `<=`, `>=`, `=`)
- Only bounding box filters (`&&`) are pushed down for spatial filtering
- OGR connections are made for every query (no pooling)
- All columns are retrieved every time


## Examples

### WFS (Web Feature Service)

```sql
CREATE SERVER geoserver
  FOREIGN DATA WRAPPER ogr_fdw
  OPTIONS (
    datasource 'WFS:https://demo.geo-solutions.it/geoserver/wfs',
    format 'WFS' );

CREATE FOREIGN TABLE topp_states (
  fid bigint,
  the_geom Geometry(MultiSurface,4326),
  gml_id varchar,
  state_name varchar,
  state_fips varchar,
  state_abbr varchar,
  land_km double precision,
  persons double precision )
  SERVER "geoserver"
  OPTIONS (layer 'topp:states');
```

### File Geodatabase

```sql
CREATE SERVER fgdbtest
  FOREIGN DATA WRAPPER ogr_fdw
  OPTIONS (
    datasource '/tmp/Querying.gdb',
    format 'OpenFileGDB' );

CREATE FOREIGN TABLE cities (
  fid integer,
  geom geometry(Point, 4326),
  city_name varchar,
  state_name varchar,
  elevation integer,
  pop1990 integer )
  SERVER fgdbtest
  OPTIONS (layer 'Cities');
```


## Advanced Features

### Writeable Tables

If the OGR driver supports it, you can insert/update/delete records. Writeable tables require a `fid` column in the table definition.

```sql
ALTER SERVER myserver
  OPTIONS (ADD updateable 'true');
```

### Column Name Mapping

Map remote column names to local names:

```sql
CREATE FOREIGN TABLE typetest_fdw_mapped (
  fid bigint,
  supertime time OPTIONS (column_name 'clock'),
  thebestname varchar OPTIONS (column_name 'name') )
  SERVER wraparound
  OPTIONS (layer 'typetest');
```

### Automatic Table Import

Use `IMPORT FOREIGN SCHEMA` to auto-create foreign table definitions:

```sql
CREATE SCHEMA fgdball;

-- Import all tables
IMPORT FOREIGN SCHEMA ogr_all
  FROM SERVER fgdbtest
  INTO fgdball;

-- Import specific tables
IMPORT FOREIGN SCHEMA ogr_all
  LIMIT TO(cities)
  FROM SERVER fgdbtest
  INTO fgdball;
```

### GDAL Options

Control driver behavior with config and open options:

```sql
CREATE SERVER myserver_latin1
  FOREIGN DATA WRAPPER ogr_fdw
  OPTIONS (
    datasource '/tmp/test',
    format 'ESRI Shapefile',
    config_options 'SHAPE_ENCODING=LATIN1' );
```

Multiple config options can be passed as a space-separated list.
