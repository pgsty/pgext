## Usage

Sources:

- [Official upstream source](https://docs.cloud.google.com/alloydb/docs/elastic-search?hl=en)

`external_search_fdw` — AlloyDB extension that provides a read-only foreign data wrapper for querying Elasticsearch data from PostgreSQL.

The reviewed catalog snapshot records version `unknown`, kind `standard`, and implementation language `C`.
The curated compatibility set is `17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "external_search_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'external_search_fdw';
```

This is a provider-specific component for `Google Cloud`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
