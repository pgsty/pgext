## Usage

Sources:

- [Official upstream README](https://github.com/siose-innova/pg_landmetrics/blob/20ea7350111e3b8bfbd15e921b8f239503893801/README.md)

`pg_landmetrics` — A PostgreSQL/PostGIS extension for calculating landscape metrics.

The reviewed catalog snapshot records version `0.0.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`, `postgis`.

```sql
CREATE EXTENSION "pg_landmetrics";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_landmetrics';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
