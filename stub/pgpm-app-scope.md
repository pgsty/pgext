## Usage

Sources:

- [packages/app-scope/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/app-scope/README.md)
- [packages/app-scope/sql/pgpm-app-scope--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/app-scope/sql/pgpm-app-scope--0.47.0.sql)
- [packages/app-scope/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/app-scope/Makefile)
- [packages/app-scope/pgpm-app-scope.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/app-scope/pgpm-app-scope.control)

`pgpm-app-scope` produces ordered scope frames for nearest-match lookups in a provisioned Constructive catalog. `app_scope.frames` returns scope, lookup database and key; `app_scope.local_frames` stays within one database.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-app-scope";
SELECT scope, lookup_database_id, key_value
FROM app_scope.frames(app_scope.platform_database_id(), 'platform');
```

### Operational Boundaries

`app_scope.platform_database_id` raises if the platform database is not registered. A NULL execution scope is rejected. The lookup depends on populated metaschema and scope projections. Upstream explicitly excludes calling the frame walker from row-security policies. Several helpers are `SECURITY DEFINER`; restrict ownership, execution privileges and the surrounding catalog. Dependencies are `plpgsql`, `pgpm-verify`, `metaschema-schema`, `metaschema-modules`. The installation SQL grants privileges to the pre-existing roles `administrator`, `authenticated`. This SQL-only module needs no library preload or restart; no PostgreSQL-major support matrix is published for this revision. Its fixed schema layout is not relocatable. Install the matching dependency closure before enabling it.
