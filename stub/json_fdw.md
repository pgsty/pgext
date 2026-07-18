## Usage

Sources:

- [Official extension control file](https://github.com/nkhorman/json_fdw/blob/65b8e7d4dc39bb2844879bb3c0199588e9cfa8a2/json_fdw.control)
- [Official upstream documentation](https://github.com/nkhorman/json_fdw/blob/65b8e7d4dc39bb2844879bb3c0199588e9cfa8a2/README.md)

`json_fdw` — Foreign data wrapper for local and remote JSON document access.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "json_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'json_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
