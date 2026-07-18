## Usage

Sources:

- [Official extension control file](https://github.com/yanliu8/cbtree/blob/871f57358cb297f22be6f419a8e52dc86567ab7c/cbtree.control)

`cbtree` — counted btree access method

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "cbtree";
SELECT extversion
FROM pg_extension
WHERE extname = 'cbtree';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
