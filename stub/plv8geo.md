## Usage

Sources:

- [Official extension control file](https://github.com/Geodan/plv8_geo/blob/ee6bd4878ef234aad5dadfeee7d29b19f0525948/plv8geo.control)
- [Official upstream documentation](https://github.com/Geodan/plv8_geo/blob/ee6bd4878ef234aad5dadfeee7d29b19f0525948/README.md)

`plv8geo` — PL/v8 geospatial functions that bundle D3, TopoJSON, JSTS, GeoTIFF and related JavaScript libraries for PostGIS workflows.

The reviewed catalog snapshot records version `0.0.2`, kind `puresql`, and implementation language `JavaScript`.
Install and validate the declared extension dependencies first: `plv8`, `postgis`.

```sql
CREATE EXTENSION "plv8geo";
SELECT extversion
FROM pg_extension
WHERE extname = 'plv8geo';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
