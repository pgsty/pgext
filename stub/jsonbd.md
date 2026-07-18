## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/jsonbd/blob/6a8aefcca9999596d75212a0f68c28590b5c3570/jsonbd.control)
- [Official upstream documentation](https://github.com/postgrespro/jsonbd/blob/6a8aefcca9999596d75212a0f68c28590b5c3570/README.md)

`jsonbd` — Compression method for JSONB type

The reviewed catalog snapshot records version `0.1`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "jsonbd";
SELECT extversion
FROM pg_extension
WHERE extname = 'jsonbd';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
