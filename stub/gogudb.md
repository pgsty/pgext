## Usage

Sources:

- [Official extension control file](https://github.com/csudata/gogudb/blob/490294508c8df14e29f312a1720e53690ef24de4/gogudb.control)
- [Official upstream documentation](https://github.com/csudata/gogudb/blob/490294508c8df14e29f312a1720e53690ef24de4/Install.md)
- [Official upstream README](https://github.com/csudata/gogudb/blob/490294508c8df14e29f312a1720e53690ef24de4/README.md)

`gogudb` — PostgreSQL extension for distributed hash and range partitioning across foreign servers.

The reviewed catalog snapshot records version `1.1`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10,11`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "gogudb";
SELECT extversion
FROM pg_extension
WHERE extname = 'gogudb';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
