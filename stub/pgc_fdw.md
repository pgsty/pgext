## Usage

Sources:

- [Official extension control file](https://github.com/fengttt/pgc_fdw/blob/5993bb88886c0a27d2ee6086f68255ba9a8460c2/pgc_fdw.control)

`pgc_fdw` — Access remote PostgreSQL tables through a postgres_fdw-derived wrapper with a FoundationDB-backed query cache

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `12`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgc_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgc_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
