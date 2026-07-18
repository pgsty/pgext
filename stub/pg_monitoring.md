## Usage

Sources:

- [Official extension control file](https://github.com/einhverfr/pgxn-postgres_monitoring/blob/b2b877bcca7ccc129f281c96587248e9ef5398a5/pg_monitoring.control)
- [Official upstream documentation](https://github.com/einhverfr/pgxn-postgres_monitoring/blob/b2b877bcca7ccc129f281c96587248e9ef5398a5/README)
- [Official PGXN distribution page](https://pgxn.org/dist/pg_monitoring/)

`pg_monitoring` — Provides a common SQL interface for collecting PostgreSQL replication, table-load and database-load monitoring statistics.

The reviewed catalog snapshot records version `0.2`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_monitoring";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_monitoring';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
