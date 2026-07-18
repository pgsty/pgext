## Usage

Sources:

- [Official extension control file](https://github.com/zilder/tuple_fdw/blob/9c737a2043bfd78fe085f0966e3a58b62e38fd70/tuple_fdw.control)
- [Official upstream documentation](https://github.com/zilder/tuple_fdw/blob/9c737a2043bfd78fe085f0966e3a58b62e38fd70/README.md)

`tuple_fdw` — Proof-of-concept C foreign data wrapper for append-only, LZ4-compressed binary tuple files used as cold storage.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "tuple_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'tuple_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
