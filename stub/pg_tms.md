## Usage

Sources:

- [Pinned official README](https://github.com/jkeifer/pg_tms/blob/3e11a51f42095632d36bcd8b9775da41edaa8cde/README.md)
- [Pinned extension SQL](https://github.com/jkeifer/pg_tms/blob/3e11a51f42095632d36bcd8b9775da41edaa8cde/pg_tms--0.0.4.sql)

`pg_tms` is a SQL/PLpgSQL PostGIS extension that converts rasters into 256-pixel Global Mercator tiles compatible with the Tile Map Service coordinate model. It creates tile rasters and coordinate helpers inside PostgreSQL; it does not provide an HTTP tile server.

### Core Workflow

PostGIS 2.1 or later is required. `CASCADE` can install that dependency when available:

```sql
CREATE EXTENSION pg_tms CASCADE;

SELECT t.x, t.y, t.z, t.rast
FROM source_rasters AS s
CROSS JOIN LATERAL tms_tile_raster_to_zoom(
  s.rast,
  12,
  true,
  'NearestNeighbor'
) AS t;
```

`tms_tile_raster_to_zoom(raster, zoom, drop_blanks, algorithm)` returns `tms_tile` rows containing `rast`, `x`, `y`, and `z`. A `zoom` of `-1` asks the extension to derive a native zoom from pixel size. `drop_blanks` omits tiles whose bands contain no data, and `algorithm` is passed to PostGIS resampling.

### Important Objects

- Coordinate types: `tms_latlon`, `tms_meters`, `tms_pixels`, `tms_tilecoord`, `tms_tilecoordz`, and their extent forms.
- Tile type: `tms_tile(rast raster, x int, y int, z int)`.
- Resolution and bounds: `tms_initialresolution`, `tms_originshift`, `tms_resolution`, `tms_zoomforpixelsize`, `tms_tilebounds_meters`, and `tms_tilebounds_latlon`.
- Conversion helpers: `tms_latlon2meters`, `tms_meters2latlon`, `tms_pixels2meters`, `tms_meters2pixels`, `tms_pixels2tile`, and `tms_meters2tile`.
- Raster workflow: `tms_raster2tilecoord_ext`, `tms_tilecoordz_from_raster`, `tms_copy_to_tile`, `tms_has_data`, and `tms_tile_raster_to_zoom`.

The extension also defines assignment casts from `tms_tilecoordz` to `raster`, `geometry`, meter bounds, and latitude/longitude bounds.

### Caveats

Only the Global Mercator path is implemented; generated tile rasters and geometries use SRID 3857. Reproject source data and inspect its georeference, pixel size, band nodata values, and dateline behavior before tiling. Large rasters or high zoom levels can generate many CPU- and storage-intensive tiles.

The project describes version 0.0.4 as under development and is now abandoned. Its SQL targets PostgreSQL 9.3/PostGIS 2.1-era APIs, so validate every used function against the deployed PostGIS version. Place an HTTP service or cache in front of PostgreSQL to serve the returned rasters. No preload or restart is required.
