## Usage

Sources:

- [Official upstream documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.PostgreSQL.rds.version.html)

`rds_tools` — AWS RDS for PostgreSQL administrative functions for engine-version inspection, password-encryption audits, autovacuum diagnostics, and storage monitoring

The reviewed catalog snapshot records version `1.9`, kind `standard`, and implementation language `C`.
The curated compatibility set is `11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "rds_tools";
SELECT extversion
FROM pg_extension
WHERE extname = 'rds_tools';
```

This is a provider-specific component for `AWS`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
