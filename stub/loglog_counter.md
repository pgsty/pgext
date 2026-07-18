## Usage

Sources:

- [Official extension control file](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/loglog/loglog_counter.control)
- [Official upstream documentation](https://pgxn.org/dist/loglog_estimator/README.html)
- [Official PGXN distribution page](https://pgxn.org/dist/loglog_estimator/)

`loglog_counter` — Aggregation functions and data type for distinct estimation based on LogLog.

The reviewed catalog snapshot records version `1.2.4`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "loglog_counter";
SELECT extversion
FROM pg_extension
WHERE extname = 'loglog_counter';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
