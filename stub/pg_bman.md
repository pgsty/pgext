## Usage

Sources:

- [Official extension control file](https://github.com/s-hironobu/pg_bman/blob/master/pg_bman.control)

`pg_bman` — PostgreSQL hot-backup tooling with remote full/incremental backup and archive-log helpers.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_bman";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_bman';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
