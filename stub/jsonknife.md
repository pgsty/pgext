## Usage

Sources:

- [Official extension control file](https://github.com/niquola/jsonknife/blob/ce6e65ad39a0344feab49754636daf8bd06c5c92/jsonknife.control)
- [Official upstream documentation](https://github.com/niquola/jsonknife/blob/ce6e65ad39a0344feab49754636daf8bd06c5c92/README.md)

`jsonknife` — Useful jsonb inspection and transformation functions

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `11`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "jsonknife";
SELECT extversion
FROM pg_extension
WHERE extname = 'jsonknife';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
