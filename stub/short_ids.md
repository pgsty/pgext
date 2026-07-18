## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/short_ids/)
- [Official upstream source](https://gitlab.com/depesz/short_ids)

`short_ids` — PL/pgSQL generator for short random text IDs with collision checks and transaction advisory locks.

The reviewed catalog snapshot records version `0.1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "short_ids";
SELECT extversion
FROM pg_extension
WHERE extname = 'short_ids';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
