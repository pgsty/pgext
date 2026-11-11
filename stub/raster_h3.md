## Usage

Sources:

- [PostgreSQL interface](https://github.com/dmuldrew/raster_h3/blob/507e980c7a91feae74bdd40607acfccb64a1bbcb/docs/engines/postgresql.md)
- [Extension control](https://github.com/dmuldrew/raster_h3/blob/507e980c7a91feae74bdd40607acfccb64a1bbcb/crates/raster-h3-pg/raster_h3.control)
- [SQL definitions](https://github.com/dmuldrew/raster_h3/blob/507e980c7a91feae74bdd40607acfccb64a1bbcb/crates/raster-h3-pg/sql/raster_h3--0.7.0.sql)

`raster_h3` converts GeoTIFF and cloud-optimized rasters into H3 cell aggregates. Version `0.7.0` provides a native Rust library using PostgreSQL's function-manager ABI, without PGRX. It supports PostgreSQL 13–18 on little-endian 64-bit systems; the library must match the server major version.

### Enable and aggregate

Install the matching library and extension files, then enable the extension as a superuser:

```sql
CREATE EXTENSION raster_h3;

SELECT h3_hex, mean, count, min, max, sum
FROM h3_raster_continuous_aggregate('/data/elevation.tif', 8);

SELECT h3_hex, majority_class, majority_fraction, histogram
FROM h3_raster_categorical_aggregate('/data/landcover.tif', 8);
```

The continuous function returns summary statistics; the categorical function returns class counts, proportions, diversity and a JSON-text histogram. Both also return a cell index, resolution and WKB geometry. Each call computes one resolution, from 0 through 15. Optional arguments include `source_crs`, `nodata` and `bbox`; the bounding box uses longitude/latitude limits.

```sql
SELECT h3_hex, mean
FROM h3_raster_continuous_aggregate(
    '/data/temperature.tif', h3_res => 8,
    source_crs => 'EPSG:4326', nodata => -9999,
    bbox => '-122.50,37.70,-122.35,37.85'
);
```

### Privileges and helpers

Raster paths are opened by the database server, including supported HTTP(S) and S3 URLs. The two aggregation functions revoke `EXECUTE` from `PUBLIC` and grant it to `pg_read_server_files`. That role also permits broader server-file access, so grant it only to appropriate operators. Scalar helpers do not need raster-file access.

The scalar API includes `h3_to_string`, `string_to_h3`, `h3_to_lat`, `h3_to_lng`, `h3_get_resolution`, `h3_is_valid`, `h3_to_wkb` and `h3_cell_to_parent`. PostGIS is optional for consuming the returned WKB with `ST_GeomFromWKB`; it is not an installation dependency.

### Limits and maintenance

No shared preload or server restart is required by the extension. Existing installations from `0.4.0` onward have SQL upgrade paths via `ALTER EXTENSION raster_h3 UPDATE`. The PostgreSQL interface does not expose the project's Parquet/PMTiles writers, band selection or quantile options; use the separate upstream interfaces for those operations. Statistics weight pixel contributions, not ground area. Check the source CRS and any transformation warnings when interpreting results.
