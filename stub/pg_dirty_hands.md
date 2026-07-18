## Usage

Sources:

- [Official extension control file](https://github.com/dsarafan/pg_dirty_hands/blob/3d5bb01c39e987bc394fc0f093c37d673744ba7e/pg_dirty_hands.control)
- [Official upstream documentation](https://github.com/dsarafan/pg_dirty_hands/blob/3d5bb01c39e987bc394fc0f093c37d673744ba7e/README.md)

`pg_dirty_hands` — Provides functions to freeze individual heap tuples, including an unlogged variant.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_dirty_hands";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_dirty_hands';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
