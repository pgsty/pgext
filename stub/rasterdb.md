## Usage

Sources:

- [Official README](https://github.com/hexiaoting/rasterdb/blob/7241bd185509371b08651f9c04876144da59772c/README.md)
- [Official control file](https://github.com/hexiaoting/rasterdb/blob/7241bd185509371b08651f9c04876144da59772c/rasterdb.control)
- [Official installation SQL](https://github.com/hexiaoting/rasterdb/blob/7241bd185509371b08651f9c04876144da59772c/rasterdb--0.1.sql)
- [Official FDW implementation](https://github.com/hexiaoting/rasterdb/blob/7241bd185509371b08651f9c04876144da59772c/rasterdb.c)
- [Official raster loader](https://github.com/hexiaoting/rasterdb/blob/7241bd185509371b08651f9c04876144da59772c/load_raster.c)

`rasterdb` is an experimental read-only foreign data wrapper that opens raster files on the PostgreSQL server, converts them through GDAL and PostGIS raster internals, and returns one `raster` value per tile. It does not access a remote database; the data source is a server-local file or directory.

### Core Workflow

Install PostGIS and the extension. Create a server-side configuration file containing an absolute location, then point a one-column foreign table at that file.

```text
tile_size=100x100
location=/srv/postgresql/rasters
batchsize=50
```

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION rasterdb;

CREATE SERVER raster_server
  FOREIGN DATA WRAPPER rasterdb_fdw;

CREATE FOREIGN TABLE raster_tiles (rast raster)
  SERVER raster_server
  OPTIONS (conf_file '/etc/postgresql/rasterdb.conf');

SELECT count(*) FROM raster_tiles;
```

`tile_size` selects tile dimensions; when omitted, each source raster is returned as one tile. `location` names a single raster file or a directory scanned for GDAL-recognized files. `batchsize` controls how many tiles the FDW materializes at a time.

### Objects and Query Behavior

The extension creates `rasterdb_fdw` plus its handler and validator. A foreign table must provide the `conf_file` option and must have one PostGIS `raster` column. Scans read and transform the raster files on demand; the FDW has no write callbacks, rescan support, statistics collection, or predicate pushdown.

### Security and Compatibility

The PostgreSQL server process opens both the configuration file and every path selected by `location`. Restrict who can create or alter these foreign tables, keep configuration files outside writable application directories, and grant the database operating-system account only the filesystem access it needs.

The reviewed `0.1` code enables all PostGIS GDAL drivers process-wide, assumes a single-band input raster, accepts arbitrary options in a no-op validator, and uses PostgreSQL and PostGIS internal APIs from 2018. It also contains fixed-size filename handling and incomplete planner/executor callbacks. Treat it as a prototype for isolated data conversion, not as a maintained production FDW; test exact file formats, paths, memory use, and server-version compatibility before running it.
