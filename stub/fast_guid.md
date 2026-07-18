## Usage

Sources:

- [Official extension control file](https://github.com/dyninc/postgresql-fast-guid/blob/1fc31f4f9177be1430411a82da1da23834a83ac7/fast_guid.control)
- [Official upstream documentation](https://github.com/dyninc/postgresql-fast-guid/blob/1fc31f4f9177be1430411a82da1da23834a83ac7/README.rst)

`fast_guid` — Fast GUID generator function for PostgreSQL.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "fast_guid";
SELECT extversion
FROM pg_extension
WHERE extname = 'fast_guid';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
