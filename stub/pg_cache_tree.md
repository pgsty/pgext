## Usage

Sources:

- [Official extension control file](https://github.com/ozum/pg_cache_tree/blob/95591ed10a1eebfaf2736a9863bd71fe550d7b1b/pg_cache_tree.control)
- [Official upstream documentation](https://github.com/ozum/pg_cache_tree/blob/95591ed10a1eebfaf2736a9863bd71fe550d7b1b/doc/pg_cache_tree.md)
- [Official upstream README](https://github.com/ozum/pg_cache_tree/blob/95591ed10a1eebfaf2736a9863bd71fe550d7b1b/README.md)

`pg_cache_tree` — Caches recursive parent and child IDs in array columns using trigger and helper functions.

The reviewed catalog snapshot records version `1.0.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_cache_tree";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_cache_tree';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
