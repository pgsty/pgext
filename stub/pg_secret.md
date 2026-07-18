## Usage

Sources:

- [Official upstream documentation](https://github.com/TypedLambda/pg_secret/blob/eef18ef9f504fa55fb19011e9f4844495f89709a/NOTES)

`pg_secret` — Defines order-revealing encrypted domains, comparison operators, and a B-tree operator class.

The reviewed catalog snapshot records version `0.5`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_secret";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_secret';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
