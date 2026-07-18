## Usage

Sources:

- [Official extension control file](https://github.com/danolivo/pg_repeater/blob/93cabed0ce5cec20d36d1cd56e200a4a0091d55d/pg_repeater.control)

`pg_repeater` — Replicates utility statements and serialized query plans to a remote PostgreSQL node over postgres_fdw.

The reviewed catalog snapshot records version `0.1`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pg_execplan`, `postgres_fdw`.

```sql
CREATE EXTENSION "pg_repeater";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_repeater';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
