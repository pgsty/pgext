## Usage

Sources:

- [Official extension control file](https://github.com/akorotkov/pg_aa/blob/4ae0a182dfd66e65e18ef8faf460973c8706ac44/pg_aa.control)
- [Official upstream documentation](https://github.com/akorotkov/pg_aa)

`pg_aa` — ASCII art extension for PostgreSQL

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_aa";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_aa';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
