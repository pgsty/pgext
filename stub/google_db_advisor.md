## Usage

Sources:

- [Official upstream source](https://docs.cloud.google.com/alloydb/docs/use-index-advisor)

`google_db_advisor` — AlloyDB index advisor extension for workload analysis and index recommendations.

The reviewed catalog snapshot records version `unknown`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "google_db_advisor";
SELECT extversion
FROM pg_extension
WHERE extname = 'google_db_advisor';
```

This is a provider-specific component for `Google Cloud`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
