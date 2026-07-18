## Usage

Sources:

- [Official extension control file](https://github.com/lawrencejones/pg_frozen/blob/68004b4f21b01041487e2a080658565dc4cdc937/pg_frozen.control)
- [Official upstream documentation](https://github.com/lawrencejones/pg_frozen/blob/68004b4f21b01041487e2a080658565dc4cdc937/README.md)

`pg_frozen` — Exposes a frozen(tid) function that explains if a tuple has been frozen

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_frozen";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_frozen';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
