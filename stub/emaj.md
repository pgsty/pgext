## Usage

Sources:

- [E-Maj 5.0.0 README](https://github.com/dalibo/emaj/blob/v5.0.0/README.md)
- [E-Maj 5.0.0 changelog](https://github.com/dalibo/emaj/blob/v5.0.0/CHANGES.md)
- [E-Maj quick start](https://github.com/dalibo/emaj/blob/v5.0.0/docs/en/quickStart.rst)
- [E-Maj upgrade guide](https://github.com/dalibo/emaj/blob/v5.0.0/docs/en/upgrade.rst)
- [E-Maj setup guide](https://github.com/dalibo/emaj/blob/v5.0.0/docs/en/setup.rst)

The canonical extension name is `emaj`; E-Maj records table and sequence changes for a coordinated table group and can roll that group back to a named mark. It is useful for repeatable test runs, batch savepoints, change inspection, and targeted recovery, but an E-Maj rollback is not a replacement for PostgreSQL transaction rollback or backup.

### Core Workflow

```sql
CREATE EXTENSION emaj CASCADE;
GRANT emaj_adm TO app_admin;

SELECT emaj.emaj_create_group('my_group', true);
SELECT emaj.emaj_assign_table('app', 'orders', 'my_group');
SELECT emaj.emaj_assign_sequences('app', '.*', '', 'my_group');

SELECT emaj.emaj_start_group('my_group', 'mark_1');
-- Run application changes.
SELECT emaj.emaj_set_mark_group('my_group', 'mark_2');
-- Run more application changes.

SELECT emaj.emaj_rollback_group('my_group', 'mark_1');
SELECT emaj.emaj_stop_group('my_group');
SELECT emaj.emaj_drop_group('my_group');
```

A rollbackable table group can contain tables and sequences from several schemas, but each table must have a primary key. Audit-only groups can record changes for objects that are not rollbackable. Starting and stopping a group takes locks on its application tables, so plan these operations around concurrent traffic.

### Important Objects

- `emaj_create_group` and assignment functions define table groups.
- `emaj_start_group`, `emaj_set_mark_group`, and `emaj_stop_group` manage logging sessions and marks.
- `emaj_rollback_group` performs an unlogged rollback; `emaj_logged_rollback_group` records the compensating changes.
- Multi-group variants operate on arrays of group names at one common point in time.
- Statistics and change-dump functions inspect changes between marks or generate SQL for replay.
- `emaj_set_param` changes or resets an E-Maj parameter without direct writes to the internal parameter table.
- `emaj_drop_extension()` is the supported full-removal helper.

### Version 5.0 Upgrade

For an E-Maj extension installed at version 2.3.1 or later, install the new package files and run:

```sql
ALTER EXTENSION emaj UPDATE;
```

The documented extension upgrade preserves logs and can run while groups remain in the LOGGING state. Review these 5.0 compatibility changes before cutover:

- PostgreSQL 14 through 19 are supported; PostgreSQL 12 and 13 are no longer supported.
- Direct `INSERT`, `UPDATE`, or `DELETE` against `emaj_param` must be replaced by `emaj_set_param`.
- Idempotent start and stop calls have new allow-already-active or allow-already-idle parameters; named-argument callers must review renamed parameters.
- The PHP command-line clients and `emaj_uninstall.sql` were removed.

Installations made with the standalone SQL script do not have the same in-place extension upgrade path; follow the official delete-and-reinstall procedure.

### Requirements and Caveats

The standard `CREATE EXTENSION` path requires superuser privileges and installs `dblink` plus `btree_gist` through `CASCADE`. E-Maj also supports a limited non-superuser script installation, with capability restrictions tied to the installer role.

`max_prepared_transactions` is required only for the parallel rollback client and must be at least the intended session count; changing it requires a restart. Large groups can also require a higher `max_locks_per_transaction`. Treat E-Maj log tables as operational data: size retention deliberately, monitor their growth, and keep ordinary backups for disaster recovery.
