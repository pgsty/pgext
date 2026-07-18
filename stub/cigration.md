## Usage

Sources:

- [Official extension control file](https://github.com/cloud-sn2/cigration/blob/7da7778375544eb0fee789f51e7b92cf3019b94c/cigration.control)

`cigration` — Online migration and rebalancing of colocated Citus shards for cluster expansion, contraction, and worker replacement.

The reviewed catalog snapshot records version `1.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `citus`, `dblink`.
The curated compatibility set is `10,11,12,13`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "cigration";
SELECT extversion
FROM pg_extension
WHERE extname = 'cigration';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
