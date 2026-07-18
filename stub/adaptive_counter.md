## Usage

Sources:

- [Official extension control file](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/adaptive/adaptive_counter.control)
- [Official upstream documentation](https://pgxn.org/dist/adaptive_estimator/README.html)
- [Official PGXN distribution page](https://pgxn.org/dist/adaptive_estimator/)

`adaptive_counter` — Aggregation functions and data type for distinct estimation based on adaptive sampling.

The reviewed catalog snapshot records version `1.3.3`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "adaptive_counter";
SELECT extversion
FROM pg_extension
WHERE extname = 'adaptive_counter';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
