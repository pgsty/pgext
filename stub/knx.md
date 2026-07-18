## Usage

Sources:

- [Official extension control file](https://github.com/mike-kfed/pgknx/blob/c485a440f7347a0bd840018054466735d8f5f9ce/knx.control)
- [Official upstream documentation](https://github.com/mike-kfed/pgknx/blob/c485a440f7347a0bd840018054466735d8f5f9ce/README.md)

`knx` — data types for KNX Addresses

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "knx";
SELECT extversion
FROM pg_extension
WHERE extname = 'knx';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
