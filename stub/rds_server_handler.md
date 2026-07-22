## Usage

Sources:

- [TencentDB for PostgreSQL kernel release notes](https://cloud.tencent.com/document/product/409/67017)
- [TencentDB supported extension versions](https://cloud.tencent.com/document/product/409/75121)
- [TencentDB for PostgreSQL product page](https://cloud.tencent.com/product/postgres)

`rds_server_handler` 1.0 is a TencentDB for PostgreSQL server-maintenance extension. Tencent's public release notes describe it as providing commands to clear all server caches and, in newer kernels, process-level syscache/relcache eviction. It is provider-internal operational tooling, not a portable community PostgreSQL extension.

### Core Workflow

Confirm that the target TencentDB kernel lists version 1.0, then create and verify the extension with a privileged account:

```sql
CREATE EXTENSION rds_server_handler;

SELECT extname, extversion
FROM pg_extension
WHERE extname = 'rds_server_handler';
```

Tencent's public documentation does not publish the SQL command or function signatures for cache eviction. Discover only the objects actually installed on the target managed instance and obtain the matching invocation from Tencent Cloud support or the console documentation for that kernel:

```sql
SELECT n.nspname AS schema_name,
       p.proname AS function_name,
       pg_get_function_identity_arguments(p.oid) AS arguments
FROM pg_proc AS p
JOIN pg_namespace AS n ON n.oid = p.pronamespace
JOIN pg_depend AS d ON d.objid = p.oid AND d.deptype = 'e'
JOIN pg_extension AS e ON e.oid = d.refobjid
WHERE e.extname = 'rds_server_handler';
```

This inventory query is diagnostic only; it does not authorize calling an undocumented routine.

### Operational Boundary

Cache eviction can cause a sudden loss of warm-cache performance and affect other sessions. Treat it as a coordinated incident or maintenance action, use the narrowest provider-supported scope, and measure the instance before and after the operation. Do not guess function names or run it as a routine application query.

The extension's availability, required privileges, exact object surface, and behavior are tied to the TencentDB kernel release. Tencent's extension matrix reports version 1.0 across supported offerings, while release notes record cache-eviction additions and security-related fixes; keep the managed kernel current and re-check the provider matrix after upgrades.
