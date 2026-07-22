## Usage

Sources:

- [Official README](https://github.com/cybertec-postgresql/ora_migrator/blob/c1f90bb8f1b1d50929b3e2211033467f39535980/README.md)
- [Official migration example](https://github.com/cybertec-postgresql/ora_migrator/blob/c1f90bb8f1b1d50929b3e2211033467f39535980/sql/migrate.sql)
- [Official extension control file](https://github.com/cybertec-postgresql/ora_migrator/blob/c1f90bb8f1b1d50929b3e2211033467f39535980/ora_migrator.control)

`ora_migrator` is the Oracle plugin for `db_migrator`. It uses `oracle_fdw` to inspect an Oracle database, stage editable migration metadata, copy data and objects, and optionally run trigger-based change capture for a low-downtime cutover.

### Core Workflow

Install the declared dependencies, create an `oracle_fdw` server and user mapping, then prepare and review the staging schemas before copying anything.

```sql
CREATE EXTENSION oracle_fdw;
CREATE EXTENSION db_migrator;
CREATE EXTENSION ora_migrator;

SELECT db_migrate_prepare(
  plugin => 'ora_migrator',
  server => 'oracle_src',
  only_schemas => ARRAY['APP'],
  options => '{"max_long": 32767}'::jsonb
);

SELECT schema, table_name, migrate
FROM pgsql_stage.tables
ORDER BY schema, table_name;

SELECT db_migrate_mkforeign(
  plugin => 'ora_migrator',
  server => 'oracle_src',
  options => '{"max_long": 32767}'::jsonb
);

SELECT db_migrate_tables(plugin => 'ora_migrator');
SELECT db_migrate_indexes(plugin => 'ora_migrator');
SELECT db_migrate_constraints(plugin => 'ora_migrator');

SELECT db_migrate_finish();
```

Oracle schema names are normally uppercase, so values passed through `only_schemas` must use the Oracle spelling. `max_long` controls the `oracle_fdw` limit applied to LONG, LONG RAW, and XMLTYPE columns.

### Migration and Replication Objects

- `create_oraviews` builds the Oracle dictionary foreign tables used by the plugin.
- `oracle_test_table` and `oracle_migrate_test_data` check for zero bytes, encoding failures, and other data-conversion problems after preparation.
- `db_migrator_callback` integrates Oracle-specific discovery and translation with the generic migration pipeline.
- `oracle_replication_start`, `oracle_replication_catchup`, and `oracle_replication_finish` manage temporary Oracle log tables and triggers plus PostgreSQL staging objects.

### Low-Downtime Boundary

For replication, pause Oracle writes while `oracle_replication_start` and the initial `db_migrate_tables` transition are established. Catch-up runs at Oracle `SERIALIZABLE` isolation; perform the final catch-up with writes stopped, then cut the application over. The source user needs dictionary access, and replication needs substantial table and trigger privileges on Oracle. `session_replication_role` may need to be `replica` during catch-up to avoid firing target-side triggers and foreign-key checks.

Not every Oracle partitioning feature is supported. Ensure Oracle has enough UNDO for the consistent snapshot, review `pgsql_stage.migrate_log`, and remember that `oracle_replication_finish` removes objects on both systems. PostgreSQL `9.5+`, `oracle_fdw`, and `db_migrator` are upstream prerequisites.
