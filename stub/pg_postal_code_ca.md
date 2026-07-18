## Usage

Sources:

- [Official extension control file](https://github.com/rlichtenwalter/pg_postal_code_ca/blob/30db999d5f5e89007882afd3c2ca2a499ac755c7/pg_postal_code_ca.control)

`pg_postal_code_ca` — an efficient, validated, formatted Canadian postal code type

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_postal_code_ca";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_postal_code_ca';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
