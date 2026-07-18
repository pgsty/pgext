## Usage

Sources:

- [Official extension control file](https://github.com/Octonica/data_rig/blob/e099c8f3a7a4a471a448c4574b7d5049ef7617c3/data_rig.control)
- [Official upstream documentation](https://github.com/Octonica/data_rig/blob/e099c8f3a7a4a471a448c4574b7d5049ef7617c3/README.md)

`data_rig` — Multidimensional OLAP fact data type with GiST indexing support.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `cube`.

```sql
CREATE EXTENSION "data_rig";
SELECT extversion
FROM pg_extension
WHERE extname = 'data_rig';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
