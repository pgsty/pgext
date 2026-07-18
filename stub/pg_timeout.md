## Usage

Sources:

- [Official extension control file](https://github.com/pierreforstmann/pg_timeout/blob/cabf38c88b48da64e84d226793da762e5c6e14a1/pg_timeout.control)
- [Official upstream documentation](https://github.com/pierreforstmann/pg_timeout/blob/cabf38c88b48da64e84d226793da762e5c6e14a1/README.md)

`pg_timeout` — PostgreSQL background-worker extension for terminating idle database sessions.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_timeout";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_timeout';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
