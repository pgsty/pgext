## Usage

Sources:

- [Official extension control file](https://github.com/BlueTreble/pg_thud/blob/4d2aba37d20c942f9c05412a6f045f328106cfb3/pg_thud.control)
- [Official PGXN distribution page](https://pgxn.org/dist/pg_thud/)

`pg_thud` — Framework extending pgtap with scaffolding and test-data helpers for database unit tests.

The reviewed catalog snapshot records version `0.1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `pgtap`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_thud";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_thud';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
