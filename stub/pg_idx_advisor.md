## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/pg_idx_advisor/)

`pg_idx_advisor` — PostgreSQL index advisor

The reviewed catalog snapshot records version `0.1.2`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_idx_advisor";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_idx_advisor';
```

The curated lifecycle is `deprecated`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
