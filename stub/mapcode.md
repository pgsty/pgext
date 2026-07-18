## Usage

Sources:

- [Official extension control file](https://github.com/matlads/mapcode-postgres/blob/c74956e43492bcf9a4761af641753ede07ce122b/mapcode.control)
- [Official upstream documentation](https://github.com/matlads/mapcode-postgres/blob/c74956e43492bcf9a4761af641753ede07ce122b/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/mapcode/)

`mapcode` — Mapcodes in PostgreSQL

The reviewed catalog snapshot records version `0.0.6`, kind `standard`, and implementation language `C++`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "mapcode";
SELECT extversion
FROM pg_extension
WHERE extname = 'mapcode';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
