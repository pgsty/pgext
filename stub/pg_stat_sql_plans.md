## Usage

Sources:

- [Official extension control file](https://github.com/legrandlegrand/pg_stat_sql_plans/blob/d0401cdf5ff3f01ceb0e0b7a0d8bc3edbb620d24/pg_stat_sql_plans.control)
- [Official upstream documentation](https://github.com/legrandlegrand/pg_stat_sql_plans/blob/d0401cdf5ff3f01ceb0e0b7a0d8bc3edbb620d24/README.md)

`pg_stat_sql_plans` — Track execution statistics by normalized SQL query and plan ID.

The reviewed catalog snapshot records version `0.2`, kind `preload`, and implementation language `C`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_stat_sql_plans";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_stat_sql_plans';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
