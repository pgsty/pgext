## Usage

Sources:

- [Official archived README](https://github.com/CartoDB/cartodb-postgresql/blob/e0a0597061bef6cdd7d7239ffa07ab44491817f1/README.md)
- [Official extension control template](https://github.com/CartoDB/cartodb-postgresql/blob/e0a0597061bef6cdd7d7239ffa07ab44491817f1/cartodb.control.in)
- [Official 0.37.1 release notes](https://github.com/CartoDB/cartodb-postgresql/blob/e0a0597061bef6cdd7d7239ffa07ab44491817f1/NEWS.md)
- [Official SQL source directory](https://github.com/CartoDB/cartodb-postgresql/tree/e0a0597061bef6cdd7d7239ffa07ab44491817f1/scripts-available)

`cartodb` is the legacy server-side extension that turned a PostgreSQL database into a CartoDB user database. It combines PostGIS helpers with table metadata, quotas, CartoDB role conventions, table preparation, spatial grids/bins, overviews, synchronization, and federated-server administration; it is not a small general-purpose geospatial library.

### Installation and Basic Geometry

The generated control file fixes objects in the `cartodb` schema, requires PostGIS plus untrusted PL/Python, and requires superuser installation. PostgreSQL 12+ builds substitute `plpython3u`; PostgreSQL 11 builds use the older language name.

```sql
CREATE EXTENSION cartodb CASCADE;

SELECT cartodb.CDB_LatLng(40.4168, -3.7038);

SELECT cartodb.CDB_TransformToWebmercator(
  cartodb.CDB_LatLng(40.4168, -3.7038)
);
```

`CASCADE` may install the declared PostGIS and PL/Python dependencies. Review the resulting extension set and trusted-language policy before enabling it in an existing database.

### Major Surfaces

- `CDB_CartodbfyTable` prepares a table for legacy CartoDB conventions; it can add or rebuild database objects and should be tested on a copy first.
- `CDB_UserTables`, `CDB_ColumnNames`, `CDB_ColumnType`, and `CDB_EstimateRowCount` expose catalog/reporting helpers.
- `CDB_XYZ_Extent`, `CDB_RectangleGrid`, `CDB_HexagonGrid`, binning functions, and `CDB_TransformToWebmercator` support map-oriented queries.
- Synchronization, ghost-table, groups/OAuth, overview, quota, and federated-server functions mutate metadata or contact Redis/HTTP/remote databases when configured.

This API assumes the surrounding legacy CartoDB role and configuration model, including organization/public-user conventions. Installing the extension alone does not create a current CARTO platform deployment.

### Security and Lifecycle

Several administration paths use dynamic SQL, untrusted PL/Python, credentials in configuration tables, and outbound Redis/HTTP/database connections. Restrict schema/function privileges, protect configuration rows, constrain egress, and do not expose administrative helpers to ordinary SQL users.

Version `0.37.1` raised the supported baseline to PostgreSQL 11 and depends on PostGIS and a Python Redis module. The repository was functionally last updated in 2022 and is now archived/read-only with no support. Treat it as a migration/compatibility component; do not assume compatibility with current PostgreSQL, PostGIS, Python, Redis, or CARTO services without a maintained fork and full integration tests.
