## Usage

Sources:

- [Official extension control file](https://github.com/mikeizbicki/pgrollup/blob/904f19f5d02fb0901fdb6a7297eceefab549d7ae/pgrollup.control)
- [Official upstream documentation](https://github.com/mikeizbicki/pgrollup/blob/904f19f5d02fb0901fdb6a7297eceefab549d7ae/README.md)

`pgrollup` — create rollup tables for fast aggregate queries

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`, `plpython3u`.
The curated compatibility set is `12,13`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgrollup";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgrollup';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
