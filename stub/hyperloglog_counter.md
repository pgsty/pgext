## Usage

Sources:

- [Official upstream documentation](https://pgxn.org/dist/hyperloglog_estimator/README.html)
- [Official PGXN distribution page](https://pgxn.org/dist/hyperloglog_estimator/)
- [Official upstream README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/README.md)

`hyperloglog_counter` — Aggregation functions and data type for distinct estimation based on HyperLogLog.

The reviewed catalog snapshot records version `1.2.6`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "hyperloglog_counter";
SELECT extversion
FROM pg_extension
WHERE extname = 'hyperloglog_counter';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
