## Usage

Sources:

- [Official extension control file](https://github.com/gciolli/pg2podg/blob/039dedf753e8c4aeb05564b4bf9082d80aa8b417/pg2podg.control)
- [Official upstream documentation](https://pgxn.org/dist/pg2podg/)
- [Official upstream README](https://github.com/gciolli/pg2podg/blob/039dedf753e8c4aeb05564b4bf9082d80aa8b417/README.md)

`pg2podg` — PostgreSQL extension for two-player open deterministic games.

The reviewed catalog snapshot records version `0.1.3`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg2podg";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg2podg';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
