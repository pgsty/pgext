## Usage

Sources:

- [Official extension control file](https://api.pgxn.org/src/pair/pair-0.1.8/pair.control)
- [Official upstream documentation](https://pgxn.org/dist/pair/0.1.8/README.html)
- [Official PGXN distribution page](https://pgxn.org/dist/pair/)

`pair` — A key/value pair data type for PostgreSQL

The reviewed catalog snapshot records version `0.1.2`, kind `puresql`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pair";
SELECT extversion
FROM pg_extension
WHERE extname = 'pair';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
