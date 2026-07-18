## Usage

Sources:

- [Official extension control file](https://github.com/dylex/pguri/blob/6b80c2f61b1de04faa3f2e6c1f768e9a8997db30/pguri.control)
- [Official upstream documentation](https://github.com/dylex/pguri/blob/6b80c2f61b1de04faa3f2e6c1f768e9a8997db30/README)

`pguri` — provide indexable URI and domain-name types and URI text search

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pguri";
SELECT extversion
FROM pg_extension
WHERE extname = 'pguri';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
