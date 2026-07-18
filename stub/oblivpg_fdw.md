## Usage

Sources:

- [Official extension control file](https://github.com/rogerioacp/oblivpg_ftw/blob/5a26f0ffa02242aa187d22bde8ca8dc3bc05dde0/oblivpg_fdw.control)
- [Official upstream documentation](https://github.com/rogerioacp/oblivpg_ftw/blob/5a26f0ffa02242aa187d22bde8ca8dc3bc05dde0/README.md)

`oblivpg_fdw` — Oblivious foreign data wrapper for table access backed by trusted hardware and ORAM

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "oblivpg_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'oblivpg_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
