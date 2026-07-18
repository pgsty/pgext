## Usage

Sources:

- [Official extension control file](https://github.com/xwkuang5/pgfirestore/blob/bc524d1b5200218b973a157361d74fc55031c907/pgfirestore.control)

`pgfirestore` — Implements Firestore-style document values, collection queries, and query operators inside PostgreSQL.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgfirestore";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgfirestore';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
