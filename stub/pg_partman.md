## Usage

Sources:

- [pg_partman 5.5.0 README](https://github.com/pgpartman/pg_partman/blob/v5.5.0/README.md)
- [pg_partman 5.5.0 changelog](https://github.com/pgpartman/pg_partman/blob/v5.5.0/CHANGELOG.md)
- [pg_partman usage guide](https://github.com/pgpartman/pg_partman/blob/v5.5.0/doc/pg_partman_howto.md)
- [pg_partman reference](https://github.com/pgpartman/pg_partman/blob/v5.5.0/doc/pg_partman.md)
- [pg_partman 5.5.0 control file](https://github.com/pgpartman/pg_partman/blob/v5.5.0/pg_partman.control)

`pg_partman` automates PostgreSQL declarative partition sets by time or integer ID. It creates future partitions, applies retention, moves existing data, and can run maintenance through either SQL scheduling or an optional background worker. PostgreSQL tables remain ordinary native partitioned tables.

### Core Workflow

```sql
CREATE SCHEMA partman;
CREATE EXTENSION pg_partman SCHEMA partman;

CREATE TABLE public.measurements (
    id bigint GENERATED ALWAYS AS IDENTITY,
    created_at timestamptz NOT NULL,
    value numeric
) PARTITION BY RANGE (created_at);

SELECT partman.create_partition(
    p_parent_table := 'public.measurements',
    p_control := 'created_at',
    p_interval := '1 day'
);

CALL partman.run_maintenance_proc();
SELECT * FROM partman.show_partitions('public.measurements');
```

`create_partition()` is the current name for creating a managed set. The older `create_parent()` remains available for backward compatibility in the 5.x line. Template tables carry properties that PostgreSQL does not automatically propagate; changes made to a template after children exist apply only to future children unless old partitions are adjusted separately.

### Retention and Data Movement

```sql
UPDATE partman.part_config
SET retention = '30 days',
    retention_keep_table = false
WHERE parent_table = 'public.measurements';

CALL partman.partition_data_proc('public.measurements');
CALL partman.undo_partition_proc('public.measurements');
```

Retention is destructive when child tables are configured to be dropped. If another table references the partition set with a foreign key, set `detach_before_drop` only after ensuring referencing rows no longer block detach or drop. When using `retention_schema`, version 5.5 requires that schema and each moved child table to have the same owner.

### Background Worker

Add the worker library before server start:

```conf
shared_preload_libraries = 'pg_partman_bgw'
pg_partman_bgw.interval = 3600
pg_partman_bgw.dbname = 'mydb'
pg_partman_bgw.role = 'partman_maintainer'
```

Changing `shared_preload_libraries` requires a restart; the other worker settings can be reloaded. The worker role needs full access to the pg_partman schema and every managed partition set. Use a dedicated non-superuser role and grant it membership in the roles that own those tables:

```sql
CREATE ROLE partman_maintainer WITH LOGIN;
GRANT table_owner TO partman_maintainer;
```

The 5.5 default for `pg_partman_bgw.role` is `partman_maintainer`. An upgrade can therefore stop a previously implicit worker configuration from succeeding until that role exists and has the required privileges.

### Version 5.5 Upgrade

```sql
ALTER EXTENSION pg_partman UPDATE TO '5.5.0';
```

Version 5.5 fixes several SQL-injection and privilege-escalation paths, adds `maintenance_role` columns for RLS policies on configuration rows, and lets maintenance continue with other partition sets after one set fails. A failed set gets a warning and a null last-run marker, so monitoring must alert on both PostgreSQL logs and configuration status.

The release also adds `detach_before_drop`, inherits per-column statistics targets, and changes the retention-schema ownership rule. Review PUBLIC grants after extension updates because some update scripts recreate extension functions or procedures.

### Operational Boundaries

- PostgreSQL 14 or newer is required; version 5 uses only native declarative partitioning.
- `pg_jobmon` is optional. Installing it adds job monitoring but also another privilege boundary.
- pg_partman can be installed and run without superuser privileges when the owner, schema, table, procedure, function, temporary-table, and optional RLS grants are configured as documented.
- Only one scheduler should own routine maintenance. Do not run the background worker and an external scheduler concurrently without deliberate coordination.
- A large maintenance run can hold many locks and move substantial data. Test retention and migration on representative data, monitor the default partition, and keep backups independent of partition retention.
