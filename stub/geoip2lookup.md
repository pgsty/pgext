## Usage

Sources:

- [Official extension control file](https://github.com/adjust/pg-geoip2lookup/blob/5c8144578f1c1461cc6ec0a12e18e9708de77220/geoip2lookup.control)
- [Official upstream documentation](https://pgxn.org/dist/geoip2lookup/0.0.3/)
- [Official PGXN distribution page](https://pgxn.org/dist/geoip2lookup/)

`geoip2lookup` — Look up IP-address info in MaxMind GeoIP2 MMDB files

The reviewed catalog snapshot records version `0.0.3`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plperlu`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "geoip2lookup";
SELECT extversion
FROM pg_extension
WHERE extname = 'geoip2lookup';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
