## Usage

Sources:

- [Official extension control file](https://github.com/vpikulik/postgres_osm_pbf_fdw/blob/d46960460ad3652ebe17dab28b50c0ca650522cd/osm_fdw.control)
- [Official upstream documentation](https://github.com/vpikulik/postgres_osm_pbf_fdw/blob/d46960460ad3652ebe17dab28b50c0ca650522cd/doc/osm_fdw.md)
- [Official PGXN distribution page](https://pgxn.org/dist/osm_fdw/)

`osm_fdw` — Foreign data wrapper for reading OpenStreetMap PBF files

The reviewed catalog snapshot records version `4.1.2`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "osm_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'osm_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
