## Usage

Sources:

- [Official extension control file](https://github.com/vibhorkum/efm_extension/blob/ecb305cc185b05eaf0e766d6a9c5d23b3072e2fb/efm_extension.control)
- [Official upstream documentation](https://github.com/vibhorkum/efm_extension/blob/ecb305cc185b05eaf0e766d6a9c5d23b3072e2fb/README.md)

`efm_extension` — SQL interface for EDB Failover Manager cluster management

The reviewed catalog snapshot records version `1.1`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `dblink`, `pgcrypto`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "efm_extension";
SELECT extversion
FROM pg_extension
WHERE extname = 'efm_extension';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
