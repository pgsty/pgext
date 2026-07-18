## Usage

Sources:

- [Official extension control file](https://github.com/HUUTFJ/pg_follower/blob/c533664508bf7de6e3d9dfb0fd601f90350c0cbf/pg_follower.control)
- [Official upstream documentation](https://github.com/HUUTFJ/pg_follower/blob/c533664508bf7de6e3d9dfb0fd601f90350c0cbf/README.md)

`pg_follower` — Educational logical-replication extension that captures and applies database changes

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_follower";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_follower';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
