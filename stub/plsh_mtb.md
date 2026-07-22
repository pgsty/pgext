## Usage

Sources:

- [Official German documentation](https://github.com/credativ/plsh-mtb/blob/ec5de1202ba185acb96561d799797811c909a389/README.de.md)
- [Official extension SQL](https://github.com/credativ/plsh-mtb/blob/ec5de1202ba185acb96561d799797811c909a389/sql/plsh_mtb--1.0.sql)
- [Official backup controller script](https://github.com/credativ/plsh-mtb/blob/ec5de1202ba185acb96561d799797811c909a389/src/plsh_mtb)

`plsh_mtb` is a PL/SH and Korn-shell controller that lets database users start, pause, resume, or abort a server-side logical backup without giving them a restore interface. It launches a configurable dump command as the PostgreSQL operating-system account and records controller state in a database table.

### Prerequisites and Configuration

Install the `plsh` extension, `mksh`, the `plsh_mtb` executable, and the database extension. An administrator must define the custom settings before users call it:

```conf
plsh_mtb.dump = 'pg_dump -Z 5 PGDATABASE -f BACKUPFILE.gz'
plsh_mtb.dir = '/srv/postgresql/tenant-backups'
plsh_mtb.log = 'syslog'
```

`PGDATABASE` and `BACKUPFILE` are string placeholders replaced by the controller. The target directory and command must be writable/executable by the PostgreSQL service account.

```sql
CREATE EXTENSION plsh;
CREATE EXTENSION plsh_mtb;
```

### Control and Status

`customer_backup` accepts the enum values `start`, `stop`, `continue`, and `abort`:

```sql
SELECT customer_backup('start');
SELECT * FROM plsh_mtb_backups ORDER BY started DESC;
SELECT customer_backup('stop');
SELECT customer_backup('continue');
SELECT customer_backup('abort');
```

`plsh_mtb_backups` records `filename`, start/end timestamps, state, and the controller PID. States are `running`, `stopped`, `aborted`, `failed`, and `done`; this is controller metadata, not proof that a backup is complete or restorable.

### Privilege and Recovery Boundaries

- The install SQL explicitly grants `SELECT` on `plsh_mtb_backups` to `PUBLIC`, and PostgreSQL grants function execution to `PUBLIC` by default. Revoke both, then grant status and control separately to intended tenant roles.
- The configured dump command is expanded and executed by a shell as the PostgreSQL OS account. Only trusted administrators should set it; quote paths, constrain the backup directory, and prevent tenant-controlled database names or configuration from becoming shell syntax.
- `stop`, `continue`, and `abort` send operating-system signals. The reviewed script only checks whether `/proc/<pid>` exists; its stronger command-line verification is commented out, so stale metadata or PID reuse can target an unrelated process.
- Failed-command handling calls an undefined `mark_backup_failed` helper in the reviewed script. A dump failure can therefore leave incorrect state; patch and test this path before relying on monitoring.
- The external process and filesystem are not transactional with PostgreSQL. A rolled-back SQL call does not roll back file creation or signals, and backend/session failure can leave an orphan process or stale row.
- The extension provides no restore, retention, encryption, off-site copy, checksum, catalog, or restore verification. Add those controls separately and run regular restore drills with the exact artifacts.
