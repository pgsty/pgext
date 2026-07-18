## Usage

Sources:

- [Official extension control file](https://github.com/GeoHistoricalData/geohistorical_objects/blob/5e7ec83dc568a509a88a0a66e3add8310bba13b3/geohistorical_objects.control)
- [Official upstream documentation](https://github.com/GeoHistoricalData/geohistorical_objects/blob/5e7ec83dc568a509a88a0a66e3add8310bba13b3/README.md)

`geohistorical_objects` — Geohistorical objects support

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `pg_trgm`, `pgsfti`, `postgis`, `unaccent`.

```sql
CREATE EXTENSION "geohistorical_objects";
SELECT extversion
FROM pg_extension
WHERE extname = 'geohistorical_objects';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
