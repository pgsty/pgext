## Usage

Sources:

- [Official upstream documentation](https://github.com/MasahikoSawada/pg_simula/blob/532efbc57ebd57de8308af8278e77581bb6be885/README.md)

`pg_simula` — Inject WAIT, ERROR, FATAL, or PANIC failures into PostgreSQL operations

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_simula";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_simula';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
