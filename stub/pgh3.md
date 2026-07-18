## Usage

Sources:

- [Official extension control file](https://github.com/dlr-eoc/pgh3/blob/a225ac6da928cb0668780221da64061d3c4b1d54/pgh3.control)
- [Official upstream documentation](https://github.com/dlr-eoc/pgh3/blob/a225ac6da928cb0668780221da64061d3c4b1d54/README.md)

`pgh3` — Expose Uber H3 hierarchical geospatial indexing functions for PostGIS geometries

The reviewed catalog snapshot records version `0.3.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `postgis`.
The curated compatibility set is `10,11,12,13`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgh3";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgh3';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
