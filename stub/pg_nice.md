## Usage

Sources:

- [Official extension control file](https://github.com/roa/pg_nice/blob/9394ae17ae7b577daba2f44b146a69b8ca41e100/pg_nice.control)
- [Official upstream documentation](https://github.com/roa/pg_nice/blob/9394ae17ae7b577daba2f44b146a69b8ca41e100/README.md)

`pg_nice` — Lower the I/O scheduling priority of the current PostgreSQL backend with Linux ionice

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_nice";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_nice';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
