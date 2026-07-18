## Usage

Sources:

- [Official extension control file](https://github.com/sangli00/hash_fdw/blob/9c685317f6a08381b426706c235ac515505cac51/hash_fdw.control)
- [Official upstream documentation](https://github.com/sangli00/hash_fdw/blob/9c685317f6a08381b426706c235ac515505cac51/README)

`hash_fdw` — Shared-memory hash-table foreign data wrapper for storing and querying rows by a configured key.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "hash_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'hash_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
