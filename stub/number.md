## Usage

Sources:

- [Official extension control file](https://github.com/df7cb/postgresql-number/blob/e11fec180898e8c323bcc0bff4f22af6118628d4/number.control)

`number` — Variable-width integer data type for PostgreSQL

The reviewed catalog snapshot records version `1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "number";
SELECT extversion
FROM pg_extension
WHERE extname = 'number';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
