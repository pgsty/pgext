## Usage

Sources:

- [Official project or provider page](https://enquo.org/)

`pg_enquo` — PostgreSQL extension for encrypted query operations

The reviewed catalog snapshot records version `0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `12,13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_enquo";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_enquo';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
