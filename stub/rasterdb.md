## Usage

Sources:

- [Official extension control file](https://github.com/hexiaoting/rasterdb/blob/7241bd185509371b08651f9c04876144da59772c/rasterdb.control)
- [Official upstream documentation](https://github.com/hexiaoting/rasterdb/blob/7241bd185509371b08651f9c04876144da59772c/README.md)

`rasterdb` — rasterdb: filesystem-backed PostGIS raster foreign data wrapper

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `postgis`.

```sql
CREATE EXTENSION "rasterdb";
SELECT extversion
FROM pg_extension
WHERE extname = 'rasterdb';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
