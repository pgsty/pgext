## Usage

Sources:

- [Official upstream documentation](https://github.com/dgillis/pg_vartype/blob/536d1581db6c9e6651f014797bf6391a5231e9be/README.md)

`pg_vartype` — Multiple scalar types in a single column

The reviewed catalog snapshot records version `0.4`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_vartype";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_vartype';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
