## Usage

Sources:

- [Official extension control file](https://github.com/powa-team/pg_sortstats/blob/b9228cca5947b010c8114ecd5091e76d57ef28c3/pg_sortstats.control)
- [Official upstream documentation](https://github.com/powa-team/pg_sortstats/blob/b9228cca5947b010c8114ecd5091e76d57ef28c3/README.md)

`pg_sortstats` — Collect per-query sort statistics and estimate required work_mem

The reviewed catalog snapshot records version `0.0.1`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pg_stat_statements`.

```sql
CREATE EXTENSION "pg_sortstats";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_sortstats';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
