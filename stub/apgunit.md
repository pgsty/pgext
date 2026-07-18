## Usage

Sources:

- [Official upstream documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)
- [Official project or provider page](https://aws.amazon.com/rds/aurora/)

`apgunit` — No-longer-supported internal RDS/Aurora PostgreSQL extension listed for older Aurora PostgreSQL 9.6 releases.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "apgunit";
SELECT extversion
FROM pg_extension
WHERE extname = 'apgunit';
```

This is a provider-specific component for `AWS`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `deprecated`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
