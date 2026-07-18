## Usage

Sources:

- [Official extension control file](https://github.com/yvazquezo/check_orapg/blob/d150ae5dc822c4e8f0eddf8a29e7b3d7055e562f/check_orapg.control)

`check_orapg` — Validate Oracle-to-PostgreSQL or EDB migrations by comparing database objects, privileges, attributes, comments, and row counts.

The reviewed catalog snapshot records version `3.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `oracle_fdw`, `plpgsql`.
The curated compatibility set is `11,12,13`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "check_orapg";
SELECT extversion
FROM pg_extension
WHERE extname = 'check_orapg';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
