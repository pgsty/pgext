## Usage

Sources:

- [Official extension control file](https://api.pgxn.org/src/temporal/temporal-0.7.1/temporal.control)
- [Official upstream documentation](https://pgxn.org/dist/temporal/doc/html/reference.html)
- [Official PGXN distribution page](https://pgxn.org/dist/temporal/)

`temporal` — C extension providing a PERIOD type over timestamp-with-time-zone values, temporal relation operators and functions, plus GiST indexing support.

The reviewed catalog snapshot records version `0.7.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "temporal";
SELECT extversion
FROM pg_extension
WHERE extname = 'temporal';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
