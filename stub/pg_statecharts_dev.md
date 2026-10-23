## Usage

Sources:

- [Official source](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts_dev/README.md)
- [Official source](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts_dev/pg_statecharts_dev.control)
- [Official source](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts/README.md)

`pg_statecharts_dev` 0.1.0 is the development companion to `pg_statecharts`. It imports SCXML definitions and generates Sqitch migrations. The production state-machine runtime does not need this file-access tooling.

### Core Workflow

After installing the runtime and its dependency, enable the companion:

```sql
CREATE EXTENSION pg_statecharts_dev;
SELECT * FROM fsm.import_scxml_files(
  source_path => '/repo/statecharts',
  recursive => false,
  on_conflict_do_nothing => false
);
```

The source directory is on the database host. Statechart callbacks must exist with the expected signatures. Parse or validation failures roll back the import transaction instead of accepting part of a chart.

### Migration Generation

```sql
SELECT fsm.gen_statechart_sqitch_migrations(
  source_path => '/repo/statecharts',
  sqitch_plan_file_path => '/repo/sqitch/sqitch.plan',
  recursive => false
);
```

Create the destination directories before generation and make them accessible to the database OS user. The generator writes deploy, revert and verify scripts, updates the plan, and can overwrite existing generated chart files. Review the changes before using them as application migrations.

### Permissions and Boundaries

Server file reads require `pg_read_server_files`; writes require `pg_write_server_files`. These are broad privileges, not a sandbox restricted to the supplied path. Restrict the extension and roles to a controlled development database. It shares the `fsm` schema, is pure SQL and requires no preload; PostgreSQL 16–18 are the documented test targets.

Use the runtime's current semantic-version helpers in generated migrations. For old migrations that cast to the removed semver type, follow the upstream migration guide; regenerating every historical chart can introduce callback-order failures. Deploy reviewed generated SQL to production without installing this companion there.
