## Usage

Sources:

- [Official extension control file](https://github.com/asotolongo/vacuum_utils/blob/8de352f1960ad0d32c9d0d8cf24fa7417ef6c057/vacuum_utils.control)
- [Official PGXN distribution page](https://pgxn.org/dist/vacuum_utils/)

`vacuum_utils` — SQL functions for inspecting VACUUM and ANALYZE thresholds and activity and for invoking table maintenance.

The reviewed catalog snapshot records version `0.1.0`, kind `puresql`, and implementation language `SQL`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "vacuum_utils";
SELECT extversion
FROM pg_extension
WHERE extname = 'vacuum_utils';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
