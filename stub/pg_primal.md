## Usage

Sources:

- [Official extension control file](https://github.com/PrimalHQ/primal-server/blob/4b6f384c00fb68d7df10a38bdc33e0e951d2d446/pg_primal/pg_primal.control)
- [Official upstream README](https://github.com/PrimalHQ/primal-server/blob/4b6f384c00fb68d7df10a38bdc33e0e951d2d446/README.md)

`pg_primal` — Primal database extension used by the Primal Nostr cache server.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_primal";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_primal';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
