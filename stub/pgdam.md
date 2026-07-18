## Usage

Sources:

- [Official upstream documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`pgdam` — AWS Aurora PostgreSQL provider extension listed as pgdam version 1.7 for Aurora PostgreSQL 16 and 17.

The reviewed catalog snapshot records version `1.7`, kind `puresql`, and implementation language `C`.
The curated compatibility set is `16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgdam";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgdam';
```

This is a provider-specific component for `AWS`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
