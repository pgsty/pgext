## Usage

Sources:

- [Official extension control file](https://github.com/m-martinez/pg-audit-json/blob/68bbc4417c44dac145084ccb703a24e2a199977c/pg-audit-json.control)
- [Official upstream documentation](https://github.com/m-martinez/pg-audit-json/blob/68bbc4417c44dac145084ccb703a24e2a199977c/README.md)

`pg-audit-json` — Trigger-based PostgreSQL table auditing with JSONB row snapshots and recursive change diffs.

The reviewed catalog snapshot records version `1.0.2`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg-audit-json";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg-audit-json';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
