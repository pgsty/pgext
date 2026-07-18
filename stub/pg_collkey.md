## Usage

Sources:

- [Official extension control file](https://api.pgxn.org/src/pg_collkey/pg_collkey-0.5.1/pg_collkey.control)
- [Official upstream documentation](https://pgxn.org/dist/pg_collkey/0.5.1/)
- [Official PGXN distribution page](https://pgxn.org/dist/pg_collkey/)

`pg_collkey` — ICU collation function wrapper

The reviewed catalog snapshot records version `0.5.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_collkey";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_collkey';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
