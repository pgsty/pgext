## Usage

Sources:

- [Official extension control file](https://github.com/ossc-db/pg_dbms_stats/blob/REL14_0/pg_dbms_stats.control)
- [Official upstream documentation](https://github.com/ossc-db/pg_dbms_stats/blob/REL14_0/doc/pg_dbms_stats-en.md)
- [Official upstream README](https://github.com/ossc-db/pg_dbms_stats/blob/e041b7cf92f105e41746e8e06b106f1246cc51b0/README.md)

`pg_dbms_stats` — Statistics management extension for stabilizing PostgreSQL execution plans using fixed statistics.

The reviewed catalog snapshot records version `14.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `14`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_dbms_stats";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_dbms_stats';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
