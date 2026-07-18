## Usage

Sources:

- [Official extension control file](https://github.com/segasai/pg_hist/blob/3e67428bcc5017c8fda9a8635ca6a076417f30d7/pg_hist.control)
- [Official upstream documentation](https://github.com/segasai/pg_hist/blob/3e67428bcc5017c8fda9a8635ca6a076417f30d7/README.md)

`pg_hist` — Compute fast weighted or unweighted histograms in one or more dimensions.

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_hist";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_hist';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
