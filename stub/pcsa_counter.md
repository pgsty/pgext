## Usage

Sources:

- [Official extension control file](https://api.pgxn.org/src/pcsa_estimator/pcsa_estimator-1.3.3/pcsa_counter.control)
- [Official upstream documentation](https://pgxn.org/dist/pcsa_estimator/1.3.3/README.html)
- [Official PGXN distribution page](https://pgxn.org/dist/pcsa_estimator/)

`pcsa_counter` — Aggregation functions and data type for distinct estimation based on PCSA.

The reviewed catalog snapshot records version `1.3.3`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pcsa_counter";
SELECT extversion
FROM pg_extension
WHERE extname = 'pcsa_counter';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
