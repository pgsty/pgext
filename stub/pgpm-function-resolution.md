## Usage

Sources:

- [packages/function-resolution/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/function-resolution/README.md)
- [packages/function-resolution/sql/pgpm-function-resolution--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/function-resolution/sql/pgpm-function-resolution--0.47.0.sql)
- [packages/function-resolution/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/function-resolution/Makefile)
- [packages/function-resolution/pgpm-function-resolution.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/function-resolution/pgpm-function-resolution.control)

`pgpm-function-resolution` chooses function definitions through the PGPM scope chain and can enqueue the selected task. `function_resolution.resolve` reports a winning definition, while `function_resolution.enqueue` resolves routing and calls the application job queue.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-function-resolution";
SELECT function_definition_id, resolved_scope
FROM function_resolution.resolve(
  app_scope.platform_database_id(), 'platform', NULL, 'my:task', false
);
```

### Operational Boundaries

Provide populated metaschema and function catalogs, JWT execution context and a non-NULL scope. Definitions missing from the catalog and catalogs that cannot answer a lookup are different conditions. The pinned API also includes `function_resolution.resolve_invocation`, `function_resolution.resolve_bucket`, `function_resolution.resolve_api` and capability resolution. Enqueueing writes queue rows; it does not run a worker. Many helpers use `SECURITY DEFINER`, so application roles and execution permissions are part of the deployment boundary. Dependencies are `plpgsql`, `errors`, `pgpm-verify`, `metaschema-schema`, `metaschema-modules`, `pgpm-app-scope`, `pgpm-database-jobs`, `pgpm-jwt-claims`. The installation SQL grants privileges to the pre-existing roles `administrator`, `authenticated`, `anonymous`. This SQL-only module needs no library preload or restart; no PostgreSQL-major support matrix is published for this revision. Its fixed schema layout is not relocatable. Install the matching dependency closure before enabling it.
