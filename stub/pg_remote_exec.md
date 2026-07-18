## Usage

Sources:

- [Official extension control file](https://github.com/cybertec-postgresql/pg_remote_exec/blob/1ac5d03725b86a1bebcf754b2e7f0212a46218c5/pg_remote_exec.control)
- [Official upstream documentation](https://github.com/cybertec-postgresql/pg_remote_exec/blob/1ac5d03725b86a1bebcf754b2e7f0212a46218c5/README.md)

`pg_remote_exec` — Executes shell commands from SQL and optionally returns their standard output.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_remote_exec";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_remote_exec';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
