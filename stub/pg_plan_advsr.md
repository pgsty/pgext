## Usage

Sources:

- [Official extension control file](https://github.com/ossc-db/pg_plan_advsr/blob/6b41a96c57c3785ef5315d4368f92d17a3488327/pg_plan_advsr.control)

`pg_plan_advsr` — Automated execution plan tuning extension

The reviewed catalog snapshot records version `0.1`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pg_hint_plan`, `pg_store_plans`.
The curated compatibility set is `12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_plan_advsr";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_plan_advsr';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
