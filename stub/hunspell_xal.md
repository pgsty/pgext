## Usage

Sources:

- [Official extension control file](https://github.com/ontaev/kalmyk-hunspell/blob/92f10f374186899e658ef06fd823adae089efaee/hunspell_xal.control)
- [Official upstream documentation](https://github.com/ontaev/kalmyk-hunspell/blob/92f10f374186899e658ef06fd823adae089efaee/README.md)

`hunspell_xal` — Kalmyk Hunspell dictionary for PostgreSQL full-text search

The reviewed catalog snapshot records version `0.1`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "hunspell_xal";
SELECT extversion
FROM pg_extension
WHERE extname = 'hunspell_xal';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
