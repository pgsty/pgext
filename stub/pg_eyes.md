## Usage

Sources:

- [Official extension control file](https://github.com/alexandersamoylov/pg_eyes/blob/31c21bb1a307235bfda250727254491d505e0ee3/pg_eyes.control)
- [Official upstream documentation](https://github.com/alexandersamoylov/pg_eyes/blob/31c21bb1a307235bfda250727254491d505e0ee3/README.md)

`pg_eyes` — SQL monitoring functions and views for PostgreSQL database activity

The reviewed catalog snapshot records version `1.4`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `pg_stat_statements`, `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_eyes";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_eyes';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
