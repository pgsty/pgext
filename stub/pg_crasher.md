## Usage

Sources:

- [Official extension control file](https://github.com/davidcrawford/pg_crasher/blob/master/pg_crasher.control)
- [Official upstream documentation](https://github.com/davidcrawford/pg_crasher/blob/master/README)

`pg_crasher` — Provides a function that intentionally crashes a PostgreSQL backend for client error-handling tests.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_crasher";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_crasher';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
