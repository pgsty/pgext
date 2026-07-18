## Usage

Sources:

- [Official upstream documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AD.Security.Groups.html)
- [Official primary documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`pg_ad_mapping` — Maps Active Directory security groups to database roles on Aurora PostgreSQL.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_ad_mapping";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ad_mapping';
```

This is a provider-specific component for `AWS`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
