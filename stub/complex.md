## Usage

Sources:

- [Official extension control file](https://github.com/semenikhind/complex_contrib/blob/5b815cbf896ba3713862c2302202411d27c7c421/complex.control)

`complex` — Complex-number data type for PostgreSQL.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "complex";
SELECT extversion
FROM pg_extension
WHERE extname = 'complex';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
