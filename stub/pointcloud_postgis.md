

## Usage

> [pointcloud_postgis: PostGIS integration for pgPointcloud](https://github.com/pgpointcloud/pointcloud)

`pointcloud_postgis` is a bridge extension that integrates the pgPointcloud extension with PostGIS. It enables conversion between pointcloud geometry types and PostGIS geometry types.

```sql
CREATE EXTENSION pointcloud_postgis;
```

This extension requires both `pointcloud` and `postgis` to be installed. It adds functions to convert between `pcpoint`/`pcpatch` types and PostGIS `geometry` types, enabling spatial queries on point cloud data using PostGIS operators and functions.
