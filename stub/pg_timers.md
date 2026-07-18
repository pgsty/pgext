## Usage

Sources:

- [Official extension control file](https://github.com/proventuslabs/pg_timers/blob/21160c1b2d19a274279b02b1f8fed6170f6d2984/pg_timers.control)
- [Official upstream documentation](https://github.com/proventuslabs/pg_timers/blob/21160c1b2d19a274279b02b1f8fed6170f6d2984/README.md)

`pg_timers` — Schedule SQL execution at exact times or after intervals in PostgreSQL

The reviewed catalog snapshot records version `0.0.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_timers";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_timers';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
