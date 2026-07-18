## Usage

Sources:

- [Official extension control file](https://github.com/johto/pg_metrics/blob/ce7dcbf8ad5bec4bf27906c28f404f702f89cbaf/pg_metrics.control)

`pg_metrics` — Collects named counters from PostgreSQL backends in shared memory and exposes them through SQL metric functions.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_metrics";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_metrics';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
