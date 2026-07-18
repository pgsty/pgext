## Usage

Sources:

- [Official upstream source](https://support.huaweicloud.com/intl/en-us/usermanual-rds-pg/rds_09_0073.html)

`rds_pg_sql_ccl` — Huawei Cloud RDS for PostgreSQL kernel extension for limiting total and per-query SQL concurrency.

The reviewed catalog snapshot records version `unknown`, kind `standard`, and implementation language `C`.
The curated compatibility set is `11,12,13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "rds_pg_sql_ccl";
SELECT extversion
FROM pg_extension
WHERE extname = 'rds_pg_sql_ccl';
```

This is a provider-specific component for `Huawei Cloud`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
