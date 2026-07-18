## Usage

Sources:

- [Official extension control file](https://github.com/nuko-yokohama/pg_fraction/blob/812f0599f82a47024b96ed0c745c0464bf1e2f2b/pg_fraction.control)

`pg_fraction` — Adds a fraction data type with arithmetic, comparison, aggregate and B-tree operator-class support.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_fraction";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_fraction';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
