## Usage

Sources:

- [Official extension control file](https://github.com/pgbigm/ludia_funcs/blob/d220429554fbc1f04573f444a58bc5ab8919d4a6/ludia_funcs.control)
- [Official upstream documentation](https://github.com/pgbigm/ludia_funcs/blob/d220429554fbc1f04573f444a58bc5ab8919d4a6/README.md)

`ludia_funcs` — Ludia full-text search helper functions

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "ludia_funcs";
SELECT extversion
FROM pg_extension
WHERE extname = 'ludia_funcs';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
