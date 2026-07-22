## Usage

Sources:

- [Huawei Cloud RDS for PostgreSQL rds_hwdrs_privs documentation](https://support.huaweicloud.com/intl/en-us/usermanual-rds-pg/rds_09_0068.html)

`rds_hwdrs_privs` is a Huawei Cloud RDS for PostgreSQL managed extension that delegates narrowly defined replication and catalog privileges from the service-managed `root` role. It exists for older RDS engines where `root` could not perform these operations directly; it is not an installable community extension.

### Check and Install

Huawei documents it for PostgreSQL 9.5, 9.6, 10, and 11.5 or earlier. Later versions allow `root` to perform the covered operations directly. Check the exact instance and use the provider control function rather than stock `CREATE EXTENSION`:

```sql
SELECT *
FROM pg_available_extension_versions
WHERE name = 'rds_hwdrs_privs';

SELECT control_extension('create', 'rds_hwdrs_privs');
```

Only `root` or a member of `root` can use the extension functions.

### Delegated Operations

```sql
SELECT control_select_on_pg_authid('grant', 'drs_sync');
SELECT control_user_privilege('replication', 'drs_sync');
SELECT control_user_privilege('bypassrls', 'drs_sync');

SELECT create_publication_for_all_tables(
  'drs_all_tables',
  'insert, update, delete'
);

SELECT exec_pg_replication_origin_func(
  'pg_replication_origin_create',
  'drs_origin'
);
```

`control_select_on_pg_authid` accepts `grant` or `revoke`. `control_user_privilege` accepts `BYPASSRLS`, `NOBYPASSRLS`, `REPLICATION`, or `NOREPLICATION`. The publication helper creates a `FOR ALL TABLES` publication owned by `root`; PostgreSQL 10 supports insert/update/delete publication actions, while PostgreSQL 11 also supports truncate. The replication-origin wrapper allows only the provider-listed `pg_replication_origin_*` operations and takes a second argument only when the selected function requires one.

### Security and Lifecycle Boundaries

These privileges expose password-role metadata, bypass row-level security, permit replication connections, and affect all-table logical publication. Grant only the capability required by the migration/DRS account, use a dedicated role, and revoke it after the job. Granting membership in `root` is much broader than the helper calls.

Dropping the extension uses `SELECT control_extension('drop', 'rds_hwdrs_privs')`. First remove delegated privileges and replication objects deliberately; do not assume uninstall reverses operational state. Availability and semantics belong to the Huawei managed engine, so coordinate use with the provider and the target DRS workflow.
