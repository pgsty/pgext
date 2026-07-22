## Usage

Sources:

- [Huawei Cloud RDS for PostgreSQL documentation](https://support.huaweicloud.com/intl/en-us/usermanual-rds-pg/rds_09_0067.html)

`rds_hwdrs_ddl` is a Huawei Cloud RDS for PostgreSQL provider extension used during incremental DDL synchronization. It creates the table, function, event trigger, and grants needed by Huawei's synchronization service when the database user cannot create those objects directly. It is a managed-service facility, not a portable open-source extension for self-hosted PostgreSQL.

### Availability and Enablement

First check the target RDS instance's live extension inventory; supported versions depend on the service instance and are not given as a general PostgreSQL compatibility matrix:

```sql
SELECT *
FROM pg_available_extension_versions
WHERE name = 'rds_hwdrs_ddl';
```

Huawei's documented SQL installation path uses the provider function rather than direct `CREATE EXTENSION`:

```sql
SELECT control_extension('create', 'rds_hwdrs_ddl');
```

After enablement, confirm the provider-created objects before starting a synchronization task:

```sql
SELECT relname, relowner::regrole, relacl
FROM pg_class
WHERE relname = 'hwdrs_ddl_info';

SELECT proname, proowner::regrole
FROM pg_proc
WHERE proname = 'hwdrs_ddl_function';

SELECT evtname, evtevent
FROM pg_event_trigger
WHERE evtname = 'hwdrs_ddl_event';
```

### Synchronization Lifecycle

The documented workflow is intentionally temporary:

1. enable `rds_hwdrs_ddl` on the source or target RDS for PostgreSQL instance as required by the Huawei synchronization workflow;
2. create and run the PostgreSQL-to-RDS synchronization task;
3. after synchronization completes, remove the extension through the provider-supported console or SQL management path so its tracking table, function, and event trigger are deleted.

Follow the current console instructions for the exact topology and removal command. The provider page's displayed removal identifier differs from the extension name used everywhere else on that page, so verify the command against the instance's available extensions or use the RDS console rather than copying it blindly.

### Operational Boundaries

The `hwdrs_ddl_event` event trigger fires at `ddl_command_end`, and `hwdrs_ddl_info` stores information used by synchronization. Installing the extension therefore changes cluster-side DDL behavior and grants access needed by the service. Enable it only for the duration of the approved migration, monitor the synchronization task, and confirm removal afterward.

Do not substitute hand-created lookalike tables, functions, or triggers: object definitions and permissions are controlled by Huawei Cloud and may vary with the service. The documentation page was updated on 2025-11-10; for a current production task, re-check the live provider documentation and instance inventory before use.
