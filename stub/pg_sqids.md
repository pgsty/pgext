## Usage

Sources:

- [Official upstream documentation](https://github.com/sqids/sqids-postgresql/blob/709167f4c3c97f7af096a94c854e8e5071ddeaf5/README.md)
- [Official project or provider page](https://sqids.org/postgresql)

`pg_sqids` — Encode numbers as short URL-safe Sqids and decode them in PostgreSQL

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_sqids";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_sqids';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
