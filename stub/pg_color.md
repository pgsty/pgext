## Usage

Sources:

- [Official extension control file](https://github.com/byucesoy/pg_color/blob/master/pg_color.control)
- [Official upstream documentation](https://github.com/byucesoy/pg_color/blob/master/README.md)

`pg_color` — Color data type for PostgreSQL

The reviewed catalog snapshot records version `1.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_color";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_color';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
