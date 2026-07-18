## Usage

Sources:

- [Official extension control file](https://github.com/brahmlower/pgzan/blob/b2e45f17059639863dae8fdadff805fe51ca55ac/pgzan.control)

`pgzan` — Proof-of-concept pgrx extension for evaluating Zanzibar-style authorization inside PostgreSQL row-level security policies.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgzan";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgzan';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
