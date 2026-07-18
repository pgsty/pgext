## Usage

Sources:

- [Official upstream documentation](https://pgxn.org/dist/tab_tier/doc/tab_tier.html)
- [Official PGXN distribution page](https://pgxn.org/dist/tab_tier/)

`tab_tier` — PL/pgSQL extension for job-driven retention, tiering, and date-based table partition maintenance.

The reviewed catalog snapshot records version `1.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "tab_tier";
SELECT extversion
FROM pg_extension
WHERE extname = 'tab_tier';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
