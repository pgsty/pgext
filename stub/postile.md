## Usage

Sources:

- [Official extension control file](https://github.com/nyurik/postile/blob/a11e6479aa435fc17379a7c23ad1f9501c43aca8/postile.control)

`postile` — A pgrx PostgreSQL extension for map-tile generation helpers and gzip/Brotli bytea compression.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "postile";
SELECT extversion
FROM pg_extension
WHERE extname = 'postile';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
