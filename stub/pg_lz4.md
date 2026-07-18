## Usage

Sources:

- [Official extension control file](https://github.com/zilder/pg_lz4/blob/f3259686848a61c544dbe58df5627ed3d77a2144/pg_lz4.control)

`pg_lz4` — LZ4 compression access method for a patched PostgreSQL development branch.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_lz4";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_lz4';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
