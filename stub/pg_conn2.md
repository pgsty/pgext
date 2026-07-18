## Usage

Sources:

- [Official extension control file](https://github.com/serpent7776/pg_conn2/blob/master/pg_conn2.control)
- [Official upstream documentation](https://github.com/serpent7776/pg_conn2/blob/master/README.md)

`pg_conn2` — Manage separate PostgreSQL database connections from SQL code

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_conn2";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_conn2';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
