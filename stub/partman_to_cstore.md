## Usage

Sources:

- [Official extension control file](https://github.com/dturon/partman_to_cstore/blob/80c6411318494837e5ef92825dd7a6c80b61fb66/partman_to_cstore.control)

`partman_to_cstore` — Move old pg_partman partitions to cstore columnar storage.

The reviewed catalog snapshot records version `1.1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `cstore_fdw`, `pg_partman`, `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "partman_to_cstore";
SELECT extversion
FROM pg_extension
WHERE extname = 'partman_to_cstore';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
