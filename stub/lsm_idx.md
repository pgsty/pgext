## Usage

Sources:

- [Official extension control file](https://github.com/topretejal/lsm_idx/blob/8e71df09b49b4413dd18233705ccac223ccd75ce/lsm_idx.control)

`lsm_idx` — Experimental LSM-style index access method built on PostgreSQL B-tree internals

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "lsm_idx";
SELECT extversion
FROM pg_extension
WHERE extname = 'lsm_idx';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
