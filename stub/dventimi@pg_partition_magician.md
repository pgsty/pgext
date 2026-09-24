## Usage

Sources:

- [database.dev package](https://database.dev/dventimi/pg_partition_magician)
- [Version 0.6.0 release SQL](https://github.com/dventimisupabase/pg_partition_magician/releases/download/v0.6.0/pg_partition_magician--0.6.0.sql)
- [Version 0.6.0 control file](https://github.com/dventimisupabase/pg_partition_magician/blob/v0.6.0/pgpm_core/extension.control)
- [Version 0.6.0 README](https://github.com/dventimisupabase/pg_partition_magician/blob/v0.6.0/README.md)
- [Version 0.6.0 guide](https://github.com/dventimisupabase/pg_partition_magician/blob/v0.6.0/docs/guide.md)
- [Version 0.6.0 API reference](https://github.com/dventimisupabase/pg_partition_magician/blob/v0.6.0/docs/reference.md)
- [Version 0.6.0 changelog](https://github.com/dventimisupabase/pg_partition_magician/blob/v0.6.0/CHANGELOG.md)

`dventimi@pg_partition_magician` manages native RANGE partitions using SQL and PL/pgSQL. It supports time, numeric, UUIDv7, and suitably encoded text keys. Conversion keeps the original table as a bounded historical partition; splitting that history into smaller partitions is a separate operation.

### Enablement and Core Workflow

This entry covers the database.dev package, whose registered name remains `dventimi@pg_partition_magician` even though the GitHub owner is `dventimisupabase`. Its version `0.6.0` registry payload matches the official release SQL. With the database.dev installer and `pg_tle` already configured, register it and use the quoted extension name. The required `pg_cron` extension must be enabled in this database first.

```sql
SELECT dbdev.install('dventimi@pg_partition_magician');
CREATE EXTENSION "dventimi@pg_partition_magician" VERSION '0.6.0';
SELECT pgpm.version();
```

The following example assumes an existing `public.events` table with a monotonic, non-null `created_at` column. Any primary key or unique constraint must include that column; keyless tables are also supported. Run the conversion as a top-level statement with autocommit: `pgpm.transmute` commits between phases and cannot run inside a surrounding transaction.

```sql
CALL pgpm.transmute(
  p_parent   => 'public.events',
  p_control  => 'created_at',
  p_interval => interval '1 month',
  p_obtain   => 7,
  p_retain   => NULL
);
SELECT pgpm.schedule();
SELECT * FROM pgpm.status();
SELECT pgpm.resume('public.events');
```

The converted table initially remains paused for scheduled maintenance. Inspect its bounds before resuming. This example keeps historical partitions indefinitely. `pgpm.schedule()` creates two independently scheduled maintenance jobs; by default both run every minute.

### Important Objects

| Object | Purpose |
| --- | --- |
| `pgpm.transmute` | Convert an ordinary table and register its partitioning policy. |
| `pgpm.obtain`, `pgpm.extend_to` | Create forward partitions; explicitly extend coverage before a known large key jump. |
| `pgpm.set_regrain`, `pgpm.regrain_history` | Enable paced historical splitting or drive historical splitting manually. |
| `pgpm.set_retain`, `pgpm.retain` | Configure retention or apply it; eligible old partitions are dropped. |
| `pgpm.pause`, `pgpm.resume` | Disable or enable scheduled maintenance for a table. |
| `pgpm.schedule`, `pgpm.unschedule` | Manage the maintenance jobs in the database containing the scheduler. |
| `pgpm.status()`, `pgpm.config`, `pgpm.log` | Inspect partition coverage, policies, progress, and failures. |
| `pgpm.version()`, `pgpm.installed` | Inspect the installed implementation version and SQL installation history. |

### Requirements and Operational Boundaries

- Upstream tests PostgreSQL 15–18. Objects live in the fixed `pgpm` schema. The extension uses SQL and PL/pgSQL and has no shared library of its own; scheduled operation depends on the server's existing `pg_cron` configuration.
- Installation and table conversion need the relevant schema/DDL permissions and table ownership. Run scheduling and maintenance under a role able to perform those table operations and use the scheduler; the SQL routines do not provide a privilege-escalation wrapper.
- Conversion scans the original table to certify its bounds and takes brief catalog locks. During conversion, writes outside those bounds fail. Size bound headroom for concurrent inserts and inspect incoming foreign keys: they are rejected by default unless the documented preservation mode is selected.
- There is no DEFAULT partition. Late or far-ahead values outside the available ranges fail; monitor coverage and use `pgpm.extend_to` before an intentional key jump. Arbitrarily backdated keys do not satisfy the monotonic-key design.
- Regraining copies rows, generates WAL, and temporarily needs extra disk space. The original coarse partition can remain in place, but fine-grained pruning and retention within that span wait for splitting. Retention drops data; configure it deliberately and maintain suitable backups. Archiving is a separate optional module.
- When upgrading from a version before 0.5.0, re-run `pgpm.schedule()`. Forward partition creation moved to a separate job; merely installing new SQL does not register it. The standalone SQL channel supports re-running its installation file, while the database.dev channel requires its registered package installation/update path.

Version 0.6.0 anchors partition operations to relation OIDs and backfills `pgpm.part.child_oid` during installation. Archive, write-block, retention, and hypertable cutover operations stop when a name resolves to the wrong object. The `fail_archive_identity`, `fail_write_block_identity`, and `fail_retain_identity` events count toward `status().retain_drop_failures` and do not clear themselves. Investigate the recorded identity mismatch; upgrading from 0.5.0 requires no additional manual migration step.
