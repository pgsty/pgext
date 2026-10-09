## Usage

Sources:

- [packages/partman/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/partman/README.md)
- [packages/partman/sql/pgpm-partman--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/partman/sql/pgpm-partman--0.47.0.sql)
- [packages/partman/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/partman/Makefile)
- [packages/partman/pgpm-partman.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/partman/pgpm-partman.control)

`pgpm-partman` adds partition helpers and metaschema triggers around an existing pg_partman installation. `partman.create_parent_with_retention` configures a parent and retention; ID-based helpers resolve a table through the application catalog.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-partman";
SELECT parent_table, retention, retention_keep_table
FROM partman.part_config;
```

### Operational Boundaries

Install pg_partman in the `partman` schema first; the pinned wrapper SQL does not create it. `partman.verify_parent_by_id` checks registration and `partman.run_maintenance_by_id` runs maintenance. Installation attaches INSERT and DELETE triggers to `metaschema_public.partition`. Deleting a partition record removes its configuration and can drop its recorded template table. Several maintenance helpers use `SECURITY DEFINER`; restrict execution and catalog writes. Compatibility follows the installed pg_partman API and matching metaschema layout. Dependencies are `plpgsql`, `pg_partman`, `metaschema-schema`. The installation SQL grants privileges to the pre-existing roles `authenticated`. This SQL-only module needs no library preload or restart; no PostgreSQL-major support matrix is published for this revision. Its fixed schema layout is not relocatable. Install the matching dependency closure before enabling it.
