## Usage

Sources:

- [Official extension control file](https://github.com/rlichtenwalter/pg_array_multi_index/blob/d315364e110db6ec50e822e2eb5957d988d5a5a6/pg_array_multi_index.control)

`pg_array_multi_index` — simultaneous multiple-indexing of PostgreSQL arrays

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_array_multi_index";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_array_multi_index';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
