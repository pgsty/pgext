## Usage

Sources:

- [Official extension control file](https://github.com/s-hironobu/dont_drop_db/blob/8610f0b4b37af5634135414614dad1b3080bc8de/dont_drop_db.control)
- [Official upstream documentation](https://github.com/s-hironobu/dont_drop_db/blob/8610f0b4b37af5634135414614dad1b3080bc8de/README.md)

`dont_drop_db` — Prevents DROP DATABASE for databases listed in the dont_drop_db.list setting.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "dont_drop_db";
SELECT extversion
FROM pg_extension
WHERE extname = 'dont_drop_db';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
