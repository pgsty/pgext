## Usage

Sources:

- [Official extension control file](https://github.com/hlinnaka/pg_quota/blob/0e86613bc39c1653f4c562b2cfc2b79aa3eeebe6/pg_quota.control)

`pg_quota` — Enforce soft per-role disk-space quotas with background workers

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_quota";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_quota';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
