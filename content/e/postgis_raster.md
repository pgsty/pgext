---
title: "postgis_raster"
linkTitle: "postgis_raster"
description: "PostGIS raster types and functions"
weight: 1502
categories: ["GIS"]
width: full
---

[**postgis**](https://git.osgeo.org/gitea/postgis/postgis) : PostGIS raster types and functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1502** | {{< badge content="postgis_raster" link="https://git.osgeo.org/gitea/postgis/postgis" >}} | {{< ext "postgis_raster" "postgis" >}} | `3.6.3` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "postgis" >}} |
|    **Need By**    | {{< ext "h3_postgis" >}} |
|   **See Also**    | {{< ext "pointcloud_postgis" >}} {{< ext "pointcloud" >}} {{< ext "pgrouting" >}} {{< ext "h3" >}} {{< ext "q3c" >}} {{< ext "ogr_fdw" >}} {{< ext "geoip" >}} {{< ext "pg_polyline" >}} |
|    **Siblings**   | {{< ext "postgis" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `postgis` | `postgis` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.3` | {{< bg "18" "postgis36_18" "green" >}} {{< bg "17" "postgis36_17" "green" >}} {{< bg "16" "postgis36_16" "green" >}} {{< bg "15" "postgis36_15" "green" >}} {{< bg "14" "postgis36_14" "green" >}} | `postgis36_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.3` | {{< bg "18" "postgresql-18-postgis-3" "green" >}} {{< bg "17" "postgresql-17-postgis-3" "green" >}} {{< bg "16" "postgresql-16-postgis-3" "green" >}} {{< bg "15" "postgresql-15-postgis-3" "green" >}} {{< bg "14" "postgresql-14-postgis-3" "green" >}} | `postgresql-$v-postgis-3` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 6" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 6" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 5" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 5" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://git.osgeo.org/gitea/postgis/postgis" title="Repository" icon="link" subtitle="git.osgeo.org/gitea/postgis/postgis" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install postgis;		# install via package name, for the active PG version
pig install postgis_raster;		# install by extension name, for the current active PG version

pig install postgis_raster -v 18;   # install for PG 18
pig install postgis_raster -v 17;   # install for PG 17
pig install postgis_raster -v 16;   # install for PG 16
pig install postgis_raster -v 15;   # install for PG 15
pig install postgis_raster -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION postgis_raster CASCADE; -- requires postgis
```



## Usage

> [PostGIS Raster: Raster data support for PostGIS](https://github.com/postgis/postgis)

PostGIS Raster extends PostGIS with support for raster (gridded) data stored directly in PostgreSQL. It enables raster analysis, raster/vector interaction, and map algebra operations within SQL.

- [Raster Reference](https://postgis.net/docs/RT_reference.html)

### Setup

```sql
CREATE EXTENSION postgis_raster;
```

--------

## Loading Raster Data

The `raster2pgsql` command-line tool imports raster files (GeoTIFF, etc.) into PostgreSQL:

```bash
# Load a GeoTIFF as tiled 100x100 rasters, create spatial index, use COPY
raster2pgsql -s 4326 -t 100x100 -I -C -M elevation.tif public.dem | psql mydb

# Append to existing table
raster2pgsql -s 4326 -t 100x100 -a more_data.tif public.dem | psql mydb
```

Key flags:
- `-s <srid>` -- Set the SRID
- `-t <width>x<height>` -- Tile the raster into chunks
- `-I` -- Create a spatial GiST index
- `-C` -- Apply raster constraints
- `-M` -- Vacuum analyze after loading

--------

## Querying Raster Data

### Raster Metadata

```sql
-- Get raster dimensions and pixel size
SELECT rid,
    ST_Width(rast) AS width,
    ST_Height(rast) AS height,
    ST_ScaleX(rast) AS pixel_size_x,
    ST_ScaleY(rast) AS pixel_size_y,
    ST_NumBands(rast) AS bands,
    ST_SRID(rast) AS srid
FROM dem LIMIT 5;
```

### Pixel Values

```sql
-- Get the value at a specific point
SELECT ST_Value(rast, ST_SetSRID(ST_MakePoint(-73.985, 40.748), 4326)) AS elevation
FROM dem
WHERE ST_Intersects(rast, ST_SetSRID(ST_MakePoint(-73.985, 40.748), 4326));

-- Get value at column/row position (band 1)
SELECT ST_Value(rast, 1, 10, 20) FROM dem WHERE rid = 1;
```

### Band Statistics

```sql
SELECT (ST_SummaryStats(rast)).*
FROM dem WHERE rid = 1;
-- Returns: count, sum, mean, stddev, min, max
```

--------

## Raster Processing

### Clipping Rasters by Vector Geometry

```sql
-- Clip raster to a polygon boundary
SELECT ST_Clip(rast, geom) AS clipped_rast
FROM dem, boundaries
WHERE ST_Intersects(rast, geom);
```

### Map Algebra

Apply pixel-by-pixel operations:

```sql
-- Single-raster map algebra: classify elevation
SELECT ST_MapAlgebra(rast, 1, NULL,
    'CASE WHEN [rast] > 100 THEN 1 WHEN [rast] > 50 THEN 2 ELSE 3 END') AS classified
FROM dem;

-- Two-raster map algebra: difference between two DEMs
SELECT ST_MapAlgebra(a.rast, 1, b.rast, 1, '[rast1] - [rast2]') AS diff
FROM dem_old a, dem_new b
WHERE ST_Intersects(a.rast, b.rast);
```

### Raster/Vector Interaction

```sql
-- Convert raster pixels to vector points
SELECT (ST_PixelAsPoints(rast)).*
FROM dem WHERE rid = 1;

-- Convert raster to polygons (one per unique value)
SELECT (ST_DumpAsPolygons(rast)).*
FROM dem WHERE rid = 1;

-- Intersect raster with vector and get values
SELECT p.name, ST_Value(d.rast, p.geom) AS elevation
FROM dem d, points p
WHERE ST_Intersects(d.rast, p.geom);
```

### Resampling and Reprojection

```sql
-- Resample to a different pixel size
SELECT ST_Rescale(rast, 0.001, -0.001) FROM dem;

-- Reproject to a different SRID
SELECT ST_Transform(rast, 3857) FROM dem;
```

--------

## Exporting Rasters

```sql
-- Export as GeoTIFF (binary)
SELECT ST_AsTIFF(rast) FROM dem WHERE rid = 1;

-- Export as PNG
SELECT ST_AsPNG(rast) FROM dem WHERE rid = 1;

-- Export as JPEG
SELECT ST_AsJPEG(rast) FROM dem WHERE rid = 1;
```
