## Usage

Sources:

- [Official extension control file](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/pathman_sharding.control)

`pathman_sharding` — Improved sharding for pg_pathman.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pg_pathman`, `plpgsql`, `postgres_fdw`.
The curated compatibility set is `10`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pathman_sharding";
SELECT extversion
FROM pg_extension
WHERE extname = 'pathman_sharding';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
