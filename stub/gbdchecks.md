## Usage

Sources:

- [Official extension control file](https://github.com/huseynsnmz/gbdchecks/blob/ff7090d1a2a7064df6d43deeab4f10cf7c09ff0d/gbdchecks.control)
- [Official upstream documentation](https://github.com/huseynsnmz/gbdchecks/blob/ff7090d1a2a7064df6d43deeab4f10cf7c09ff0d/README.md)

`gbdchecks` — Database health-check functions and views for bloat, locks, query statistics, and privileges.

The reviewed catalog snapshot records version `1.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `pg_stat_statements`.
The curated compatibility set is `11`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "gbdchecks";
SELECT extversion
FROM pg_extension
WHERE extname = 'gbdchecks';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
