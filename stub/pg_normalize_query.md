## Usage

Sources:

- [Official upstream documentation](https://github.com/fabriziomello/pg_normalize_query/blob/bd73d4979dbdb89e152b2a776390369992c8aa42/README.md)

`pg_normalize_query` — Normalize constants in SQL statements using PostgreSQL parser logic

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_normalize_query";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_normalize_query';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
