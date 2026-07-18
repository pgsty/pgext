## Usage

Sources:

- [Official extension control file](https://github.com/pgstuff/sys_syn_dblink/blob/0c7a18da219a2e71b83ce374e589cfa383a72518/sys_syn_dblink.control)

`sys_syn_dblink` — PL/pgSQL synchronization processor that pulls sys_syn queues over named dblink connections, transforms batches, writes local tables, and returns processing status to the source database.

The reviewed catalog snapshot records version `0.0.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `dblink`, `hstore`, `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "sys_syn_dblink";
SELECT extversion
FROM pg_extension
WHERE extname = 'sys_syn_dblink';
```

The curated lifecycle is `preview`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
