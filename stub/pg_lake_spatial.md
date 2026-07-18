## Usage

Sources:

- [Official upstream README](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/README.md)

`pg_lake_spatial` — Adds geospatial file-format and spatial-function support to pg_lake through PostGIS and GDAL-backed workflows.

The reviewed catalog snapshot records version `3.4`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pg_lake`, `postgis`.
The curated compatibility set is `16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_lake_spatial";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_lake_spatial';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
