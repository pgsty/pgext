## Usage

Sources:

- [Official upstream source](https://docs.cloud.google.com/alloydb/docs/columnar-engine/about?hl=en)

`google_columnar_engine` — AlloyDB column store, planner, and execution engine for accelerating HTAP and OLAP scans, joins, and aggregates.

The reviewed catalog snapshot records version `unknown`, kind `standard`, and implementation language `C`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "google_columnar_engine";
SELECT extversion
FROM pg_extension
WHERE extname = 'google_columnar_engine';
```

This is a provider-specific component for `Google Cloud`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
