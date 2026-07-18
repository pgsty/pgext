## Usage

Sources:

- [Official extension control file](https://github.com/tongxin/mypg_sharding/blob/567526b4292bc4b613b7d7c18ad0b423da1985d6/mypg_sharding.control)
- [Official upstream documentation](https://github.com/tongxin/mypg_sharding/blob/567526b4292bc4b613b7d7c18ad0b423da1985d6/README.md)

`mypg_sharding` — mypg sharding utilities for PostgreSQL

The reviewed catalog snapshot records version `0.0.1`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`, `postgres_fdw`.

```sql
CREATE EXTENSION "mypg_sharding";
SELECT extversion
FROM pg_extension
WHERE extname = 'mypg_sharding';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
