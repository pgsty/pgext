## Usage

Sources:

- [Official extension control file](https://github.com/bigsql/pgtsql/blob/3240980dff20736a38101eb0806c8580ff169d9c/pgtsql.control)
- [Official upstream documentation](https://github.com/bigsql/pgtsql/blob/3240980dff20736a38101eb0806c8580ff169d9c/README.md)

`pgtsql` — provide Transact-SQL compatibility for PostgreSQL

The reviewed catalog snapshot records version `3.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `11`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgtsql";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgtsql';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
