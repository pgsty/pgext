## Usage

Sources:

- [Official upstream documentation](https://github.com/BRupireddy2/pg_synthesize_wal/blob/5eb6e95f7f3e615626cfc5fede3c789f8b2de852/README.md)

`pg_synthesize_wal` — Generate PostgreSQL WAL records of various sizes

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pg_walinspect`.
The curated compatibility set is `15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_synthesize_wal";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_synthesize_wal';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
