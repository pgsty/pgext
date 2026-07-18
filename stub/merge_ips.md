## Usage

Sources:

- [Official extension control file](https://github.com/cybertec-postgresql/merge_ips/blob/3cf654cef7c7ce7fafebc9b0799f15da8725da52/merge_ips.control)
- [Official upstream documentation](https://github.com/cybertec-postgresql/merge_ips/blob/3cf654cef7c7ce7fafebc9b0799f15da8725da52/README)

`merge_ips` — Functions to merge IP addresses, ignoring gaps

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "merge_ips";
SELECT extversion
FROM pg_extension
WHERE extname = 'merge_ips';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
