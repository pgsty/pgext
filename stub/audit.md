## Usage

Sources:

- [Official extension control file](https://github.com/ttfkam/pg_audit/blob/c7f3069058215873ae8e497e1476c1e195309479/audit.control)
- [Official upstream documentation](https://github.com/ttfkam/pg_audit/blob/c7f3069058215873ae8e497e1476c1e195309479/README.md)

`audit` — Audits insert, update, and delete operations on PostgreSQL tables.

The reviewed catalog snapshot records version `1.0.0`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "audit";
SELECT extversion
FROM pg_extension
WHERE extname = 'audit';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
