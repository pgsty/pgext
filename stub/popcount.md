## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/popcount/)

`popcount` — provide population-count functions for PostgreSQL bit(n) values

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "popcount";
SELECT extversion
FROM pg_extension
WHERE extname = 'popcount';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
