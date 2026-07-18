## Usage

Sources:

- [Official project or provider page](https://anthonysotolongo.wordpress.com/)

`rep_fdw` — Replicate table changes to another PostgreSQL server through postgres_fdw

The reviewed catalog snapshot records version `0.1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`, `postgres_fdw`.

```sql
CREATE EXTENSION "rep_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'rep_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
