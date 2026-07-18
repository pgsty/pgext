## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/psl/)

`psl` — Returns the registrable domain for a hostname using a compiled-in Public Suffix List.

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `11,12,13,14,15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "psl";
SELECT extversion
FROM pg_extension
WHERE extname = 'psl';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
