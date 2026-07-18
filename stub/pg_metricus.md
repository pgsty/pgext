## Usage

Sources:

- [Official extension control file](https://github.com/nvorobev/pg_metricus_c/blob/085cf982da50211938cf014551cfe194ebc48b27/pg_metricus.control)
- [Official PGXN distribution page](https://pgxn.org/dist/pg_metricus/)

`pg_metricus` — Sends custom metrics from PL/pgSQL code to socket-based aggregators such as Brubeck or Graphite.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_metricus";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_metricus';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
