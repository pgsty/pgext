## Usage

Sources:

- [Official extension control file](https://github.com/philopon/cas/blob/e23971e2a65e5c9e50fa238cce9bb4209f041b36/cas.control)

`cas` — CAS Registry Number data type for PostgreSQL.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "cas";
SELECT extversion
FROM pg_extension
WHERE extname = 'cas';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
