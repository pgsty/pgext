## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/pg_log_statements/)

`pg_log_statements` — Enables log_statement=all selectively for chosen PostgreSQL backend processes or newly started backends matching connection filters.

The reviewed catalog snapshot records version `0.0.2`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_log_statements";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_log_statements';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
