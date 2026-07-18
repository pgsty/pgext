## Usage

Sources:

- [Official upstream documentation](https://github.com/adjust/pg_querylog/blob/f639adb6dc1686044b76d82139df1c3a49076053/README.md)

`pg_querylog` — Captures current and recently completed backend queries in shared memory for SQL inspection.

The reviewed catalog snapshot records version `0.1`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11,12`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_querylog";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_querylog';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
