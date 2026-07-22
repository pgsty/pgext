## Usage

Sources:

- [Official TencentDB database resource-isolation guide](https://cloud.tencent.com/document/product/409/105554)
- [Official TencentDB for PostgreSQL extension-version matrix](https://cloud.tencent.com/document/product/409/75121)

`tencentdb_serverless` is a TencentDB for PostgreSQL provider extension for per-database CPU resource limits in multi-tenant instances. It is installed and configured by Tencent Cloud support as part of the managed service; it is not a downloadable, portable PostgreSQL extension.

### Enablement and Core Workflow

The official feature guide says to submit a support ticket to enable database resource isolation, install `tencentdb_serverless`, and set its service-managed parameters. It documents qualifying PostgreSQL 14, 15, and 16 kernel versions; consult the current service matrix for other versions. The feature currently applies to primary instances.

After Tencent Cloud confirms enablement, inspect the managed bounds and assign limits to every non-template database:

```sql
SHOW tencentdb_serverless.min_cpu_cores;
SHOW tencentdb_serverless.max_cpu_cores;

SELECT tencentdb_serverless.set_database_cpu_limit(
    'tenant_001', 2.0, 2.5
);

SELECT *
FROM tencentdb_serverless.resource_limit_view;
```

`set_database_cpu_limit(database_name, min_cpu_cores, max_cpu_cores)` sets the relative guaranteed minimum and absolute maximum. A maximum of `-1` means no configured cap. The provider-managed GUCs describe the instance-wide allowable range and are for inspection, not direct user modification.

Reset one database or the whole instance configuration with:

```sql
SELECT tencentdb_serverless.reset_database_limit('tenant_001');
SELECT tencentdb_serverless.reset_all_database_limit();
```

### Operational Notes

Limits take effect in real time, but the guide requires configuring every database for isolation to work as intended. Minimum CPU values are weights when total resources are constrained, not permanently dedicated cores. `resource_limit_view` includes memory columns that the guide marks reserved and currently unused.

Deleting a database removes its configuration automatically. Disabling the feature also requires a Tencent Cloud support ticket to reset parameters, clear limits, and remove the extension. Validate tenant SLOs under contention and follow the provider's current kernel, privilege, failover, billing, and support rules.
