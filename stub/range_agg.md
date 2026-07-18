## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/range_agg/)

`range_agg` — range_agg: merge adjacent or overlapping PostgreSQL range values

The reviewed catalog snapshot records version `1.2.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "range_agg";
SELECT extversion
FROM pg_extension
WHERE extname = 'range_agg';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
