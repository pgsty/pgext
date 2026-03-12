

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
