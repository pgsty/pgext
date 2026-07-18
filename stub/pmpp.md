## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/pmpp/)

`pmpp` — Run manually partitioned SQL work in parallel over asynchronous libpq connections and combine the results

The reviewed catalog snapshot records version `1.2.3`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pmpp";
SELECT extversion
FROM pg_extension
WHERE extname = 'pmpp';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
