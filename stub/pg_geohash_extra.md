## Usage

Sources:

- [Official extension control file](https://github.com/siose-innova/pg_geohash_extra/blob/ea5fcbb319503599ef26f67bb165f99b39add763/pg_geohash_extra.control)
- [Official upstream documentation](https://github.com/siose-innova/pg_geohash_extra/blob/ea5fcbb319503599ef26f67bb165f99b39add763/README.md)
- [Official project or provider page](http://www.siose-innova.es/)

`pg_geohash_extra` — Additional geohash functions for PostGIS

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `postgis`.

```sql
CREATE EXTENSION "pg_geohash_extra";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_geohash_extra';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
