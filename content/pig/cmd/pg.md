---
title: "CMD: pg"
description: Manage local PostgreSQL server with pig pg subcommand
icon: fa-solid fa-terminal
weight: 650
---

`pig postgres` (alias `pig pg`) manages local PostgreSQL instances and databases.
Use it to initialize or control an instance, run SQL, maintain tables, and create database copies.
Use [`pig pt`](/pig/cmd/pt/) for Patroni cluster operations and [`pig pitr`](/pig/cmd/pitr/) for recovery with pgBackRest.

Commands run as the PostgreSQL operating-system user, normally `postgres`.
Install PostgreSQL first. For a new data directory, make its parent writable by that user or
create an empty directory with the correct ownership before running `pg init`.

## Command Overview

**Instance control**:

| Command | Alias | Description | Notes |
|:--------|:------|:------------|:------|
| `pg init` | `initdb, i` | Initialize data directory | Wraps initdb |
| `pg start` | `boot, up` | Start PostgreSQL | Wraps pg_ctl start |
| `pg stop` | `halt, down` | Stop PostgreSQL | Wraps pg_ctl stop |
| `pg restart` | `reboot` | Restart PostgreSQL | Wraps pg_ctl restart |
| `pg reload` | `hup` | Reload configuration | Wraps pg_ctl reload |
| `pg status` | `st, stat` | Show service status | Shows processes & related services |
| `pg promote` | `pro` | Promote replica to primary | Wraps pg_ctl promote |
| `pg role` | `r` | Detect instance role | Outputs primary/replica |
{.full-width}

**Connection & Query**:

| Command | Alias | Description | Notes |
|:--------|:------|:------------|:------|
| `pg psql` | `sql, connect` | Connect to database | Wraps psql |
| `pg ps` | `activity, act` | Show current connections | Queries pg_stat_activity |
| `pg kill` | `k` | Terminate connections | Default dry-run mode |
| `pg clone` | | Clone a single database | `CREATE DATABASE ... TEMPLATE ...` |
{.full-width}

**Database Maintenance**:

| Command | Alias | Description | Notes |
|:--------|:------|:------------|:------|
| `pg vacuum` | `vac, vc` | Vacuum tables | Runs SQL `VACUUM` |
| `pg analyze` | `ana, az` | Analyze tables | Runs SQL `ANALYZE` |
| `pg freeze` | | Freeze vacuum | Runs SQL `VACUUM (FREEZE)` |
| `pg repack` | `rp` | Online table repacking | Requires pg_repack extension |
{.full-width}

**Parameter Tuning**:

| Command | Alias | Description | Notes |
|:--------|:------|:------------|:------|
| `pg tune` | `tuning` | Generate PostgreSQL tuning parameters | Auto-detects hardware and supports structured output |
{.full-width}

**Instance Fork**:

| Command | Alias | Description | Notes |
|:--------|:------|:------------|:------|
| `pg fork` | | Shortcut for `fork init` | Creates a managed fork by default, does not start it |
| `pg fork init` | `create` | Create a local one-off physical copy | Default `/pg/data-<name>` |
| `pg fork list` | | List managed forks | Scans `/pg/data-*` |
| `pg fork start` | | Start an existing fork | Supports managed names or unmanaged `--dst-data` directories |
| `pg fork stop` | | Stop an existing fork | Supports shutdown mode |
| `pg fork rm` | `remove, delete` | Remove a fork | Running forks require `--stop` |
{.full-width}

**Log Tools**:

| Command | Alias | Description | Notes |
|:--------|:------|:------------|:------|
| `pg log` | `l` | Log management | Parent command |
| `pg log list` | `ls` | List log files | |
| `pg log tail` | `t, f` | Real-time log viewing | tail -f |
| `pg log show` | `cat, c` | Output log content | |
| `pg log less` | `vi, v` | View with less | |
| `pg log grep` | `g, search` | Search logs | |
{.full-width}

**Service Subcommand** (`pg svc`, also `pg service` or `pg s`):

| Command | Alias | Description |
|:--------|:------|:------------|
| `pg svc start` | `boot, up` | Start postgres service |
| `pg svc stop` | `halt, dn, down` | Stop postgres service |
| `pg svc restart` | `reboot, rt` | Restart postgres service |
| `pg svc reload` | `rl, hup` | Reload postgres service |
| `pg svc status` | `st, stat` | Show service status |
{.full-width}

## Connect to an instance {#quick-start}

For a running instance at `/pg/data`, check its status and connect:

```bash
pig pg status -D /pg/data
pig pg psql -D /pg/data
```

Use the instance's actual data directory. In `psql`, run SQL commands, then type `\q` to exit.
To install PostgreSQL and create a test instance, follow [Getting Started](/pig/start/).

## Global Options

These options apply to all `pig pg` subcommands:

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--version` | `-v` | auto-detect | PostgreSQL major version |
| `--data` | `-D` | `/pg/data` for instance operations | Data directory path |
| `--dbsu` | `-U` | `postgres` | PostgreSQL operating-system user; `$PIG_DBSU` overrides the default |
{.full-width}

Use `pig pg svc ...` to control the `postgres` systemd unit. Use the instance commands
(`start`, `stop`, and so on) to operate on `-D` with `pg_ctl`, defaulting to `/pg/data`.
An explicit empty or whitespace-only `-D` / `--data` is a parameter error; it does not select a default target.

When `-D` is explicitly supplied, `psql`, `ps`, `kill`, `vacuum`, `analyze`, `freeze`, and `repack`
connect to that directory's instance, including database enumeration for `--all`. Before using its recorded
port and socket, PIG verifies that the PID is live and identifies a PostgreSQL process.
Missing, stale, invalid, or mismatched metadata and unknown process identity fail before a SQL
client runs, without falling back to another connection target.
Without `-D`, these SQL commands retain native client connection defaults and any `PGHOST`/`PGPORT`
visible to the client; PIG does not add host or port arguments or bind the connection to `/pg/data`.
To explicitly target a TCP-only instance, use the native client with connection options.
`pg clone` and `pg fork` retain their separate port options.

Plans, confirmation retries, and verification commands for `promote`, `vacuum`, and `repack`
preserve the selected data directory, PostgreSQL version, and operating-system user. This includes
an explicit `-D /pg/data`; omitting `-D` continues to preserve native SQL connection defaults.
`VACUUM FULL` plans report confirmation as required, matching execution.
Generated maintenance commands put the database name after `--`, so a name such as `--all`
cannot become an option when the command is copied and executed.

`PIG_NON_INTERACTIVE=1` and JSON/YAML output disable prompt stdin and use non-interactive sudo.
Native passthrough retains its interaction contract; explicit data input such as `psql -f -`
remains available.

**Version Detection Logic:**

1. If `-v` specified, use that version
2. Otherwise read from `PG_VERSION` file in data directory
3. Otherwise use the default PostgreSQL installation found by PIG.

## Service Control Commands

### pg init

Initialize PostgreSQL data directory. Wraps `initdb`.

- Use an empty directory, or a path that the database operating-system user can create.
- Data checksums are enabled. Use `-K` to disable them.
- Encoding is UTF8. PostgreSQL 17+ uses the built-in `C.UTF-8` locale. Earlier versions use
  the system `C.UTF-8` locale when available, otherwise `C`.
- An initialized directory requires `--force`, which removes its existing data after confirmation.
  A running instance cannot be overwritten. If its running state cannot be verified, initialization
  fails before removal and reports the probe error instead of claiming that PostgreSQL is running.
- Put extra `initdb` arguments after `--`. To choose another locale or encoding, run `initdb` directly.

```bash
pig pg init                       # Initialize with defaults
pig pg init -v 18                 # Specify PostgreSQL 18
pig pg init -D /data/pg18         # Specify data directory
pig pg init -K                    # Disable data checksums
pig pg init -f                    # Force init (remove existing data)
pig pg init -- --waldir=/wal      # Pass extra args to initdb
```

**Options:**

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--no-data-checksums` | `-K` | false | Disable data checksums |
| `--force` | `-f` | false | Remove an initialized directory and initialize it again |
| `--yes` | `-y` | false | Skip overwrite confirmation when used with `--force` |
{.full-width}

### pg start

Start PostgreSQL with `pg_ctl start`. For a Patroni-managed instance, prefer `pig pt start` so PostgreSQL starts under Patroni. For a systemd-managed PostgreSQL service, use `pig pg svc start`.

```bash
pig pg start                      # Start with defaults
pig pg start -D /data/pg18        # Specify data directory
pig pg start -l /pg/log/pg.log    # Redirect output to log file
pig pg start -O "-p 5433"         # Pass options to postgres
pig pg start -o json              # Structured JSON output
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--log` | `-l` | Startup log; default `/tmp/pig-pg-start.log` |
| `--timeout` | `-t` | Wait timeout: `$PGCTLTIMEOUT`, then 60 seconds |
| `--no-wait` | | Don't wait for startup completion |
| `--options` | `-O` | Options to pass to postgres |
{.full-width}

If PostgreSQL is already running, the command prints the existing postmaster PID and returns successfully.

### pg stop

Stop PostgreSQL with `pg_ctl stop`.

On a Patroni-managed instance, stopping PostgreSQL directly can make Patroni restart it or initiate failover. Use `pig pt svc stop` to stop the manager first, then check the PostgreSQL state.

```bash
pig pg stop                       # Fast shutdown (default)
pig pg stop -m smart              # Wait for clients to disconnect
pig pg stop -m immediate          # Immediate shutdown
pig pg stop --plan                # Preview stop plan
```

**Options:**

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--mode` | `-m` | fast | Shutdown mode: smart/fast/immediate |
| `--timeout` | `-t` | 60 | Wait timeout in seconds; `$PGCTLTIMEOUT` overrides the default |
| `--no-wait` | | false | Don't wait for shutdown completion |
| `--plan` | | false | Preview local `pg_ctl stop` plan only |
{.full-width}

**Shutdown Modes:**

| Mode | Description |
|:---|:---|
| `smart` | Wait for all clients to disconnect |
| `fast` | Rollback active transactions, disconnect clients, clean shutdown |
| `immediate` | Terminate all processes immediately, requires recovery on next start |
{.full-width}

### pg restart

Restart PostgreSQL server.

```bash
pig pg restart                    # Fast restart
pig pg restart -m immediate       # Immediate restart
pig pg restart -O "-p 5433"       # Restart with new options
pig pg restart --plan             # Preview restart plan
```

**Options:** Same as `pg stop`, plus `--options` (`-O`) to pass to postgres.

### pg reload

Reload PostgreSQL configuration. Sends SIGHUP signal to server.

```bash
pig pg reload                     # Reload configuration
pig pg reload -D /data/pg18       # Specify data directory
```

### pg status

Show the selected instance's state:

```bash
pig pg status
pig pg status -D /data/pg18
pig pg status -D /data/pg18 -o json
```

Text output includes `pg_ctl status`, a compact `pg_controldata` summary, processes owned by the
database operating-system user, and related services. The `postgres` row reflects the selected
instance's process state. Service rows cover `patroni`, `pgbouncer`, `vip-manager`, and `haproxy`.

Structured output describes the selected instance, including its data directory, version,
PID, port, and available control-file information.

A missing process or a PID independently identified as a different executable is stale.
`EPERM` alone never proves that PostgreSQL has stopped. If process identity cannot be established,
the state remains unknown and cannot authorize forced initialization or PITR restore.
Plans report an indeterminate running-state check as `unknown`, without a usable PID or a passed
precondition. `pg status` returns `130201` (exit `3`) for permission failures and `130808`
(exit `1`) for other status-probe failures, such as invalid PID metadata or unavailable process identity.
These results have `success: false`; their partial data does not establish that PostgreSQL is stopped.
Stale PIDs no longer make PIG report an unrelated process as PostgreSQL; `pg_ctl` and PostgreSQL
retain their own lock-file checks and may still require operator cleanup after a crash.
A PID file copied from another directory does not identify the selected instance.

### pg promote

Timeline reporting uses the selected instance's port and socket. If its PID metadata cannot be
resolved, PIG reads that directory's control data or leaves the timeline unknown. This does not
change promotion execution or its primary-role success check.

Promote a standby or a paused recovery instance to a writable primary.
Use `--plan` to inspect the action. Text mode requests confirmation; structured execution requires `--yes`.

```bash
pig pg promote                    # Promote replica
pig pg promote -D /data/pg18      # Specify data directory
pig pg promote --plan             # Preview promotion plan
pig pg promote -y                 # Skip confirmation prompt
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--timeout` | `-t` | Wait timeout: `$PGCTLTIMEOUT`, then 60 seconds |
| `--no-wait` | | Don't wait for promotion completion |
| `--plan` | | Preview promotion plan only |
| `--yes` | `-y` | Skip confirmation prompt |
{.full-width}

### pg role

Detect PostgreSQL instance role (primary or replica).

```bash
pig pg role                       # Output: primary, replica, or unknown
pig pg role -V                    # Verbose output, show detection process
pig pg role -D /data/pg18         # Specify data directory
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--verbose` | `-V` | Show detailed detection process |
{.full-width}

**Output:**

- `primary`: Current instance is primary
- `replica`: Current instance is replica
- `unknown`: Cannot determine instance role

**Detection Strategy (by priority):**

1. **Target process detection**: Read the selected `postmaster.pid`, then inspect only that postmaster's process tree
2. **Bound SQL query**: Use the selected instance's port/socket and verify `data_directory` before trusting `pg_is_in_recovery()`
3. **Data directory check**: Check for `standby.signal`, `recovery.signal`, `recovery.conf` files

File-based role detection does not establish whether PostgreSQL is running. A stopped instance
can still report its configured role, but a stale PID file never makes it appear alive.
If the running-state probe fails, the command returns an error instead of reporting a role as verified.

## Connection & Query Commands

### pg psql

Connect to PostgreSQL database via psql.

```bash
pig pg psql                       # Connect to postgres database
pig pg psql mydb                  # Connect to specific database
pig pg psql mydb -c "SELECT 1"    # Execute single command
pig pg psql -f script.sql         # Execute SQL script file
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--command` | `-c` | Execute single SQL command |
| `--file` | `-f` | Execute SQL script file |
{.full-width}

With `-D`, PIG reads that directory's `postmaster.pid` as the selected operating-system user.
It verifies the live process before connecting to the recorded port and Unix socket.
A stopped or uncertain target returns an error. Without `-D`, psql retains its native connection
defaults and any `PGHOST`/`PGPORT` visible to it.

For `-o json` or `-o yaml`, specify `-c` or `-f`; interactive psql uses text mode.

### pg ps

Show PostgreSQL current connections. Queries `pg_stat_activity` view.

```bash
pig pg ps                         # Show client connections
pig pg ps -a                      # Show all connections (including system)
pig pg ps -u admin                # Filter by user
pig pg ps -d mydb                 # Filter by database
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--all` | `-a` | Show all connections (including system) |
| `--user` | `-u` | Filter by user |
| `--database` | `-d` | Filter by database |
{.full-width}

### pg kill

Terminate PostgreSQL connections. **Default is dry-run mode**, requires `-x` to execute.

```bash
pig pg kill                       # Show connections to be terminated (dry-run)
pig pg kill -x                    # Actually terminate connections
pig pg kill --pid 12345 -x        # Terminate specific PID
pig pg kill -u admin -x           # Terminate user's connections
pig pg kill -d mydb -x            # Terminate database connections
pig pg kill -s idle -x            # Terminate idle connections
pig pg kill --cancel -x           # Cancel queries instead of terminating
pig pg kill --watch 5 -x          # Repeat every 5 seconds
pig pg kill --plan                # Preview connection termination plan
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--execute` | `-x` | Actually execute (default is dry-run) |
| `--pid` | | Terminate specific PID |
| `--user` | `-u` | Filter by user |
| `--database` | `-d` | Filter by database |
| `--state` | `-s` | Filter by state (idle/active/idle in transaction) |
| `--query` | `-q` | Filter by query pattern |
| `--all` | `-a` | Include non-client background processes |
| `--cancel` | `-c` | Cancel queries instead of terminating |
| `--watch` | | Repeat every N seconds |
| `--plan` | | Preview execution plan without terminating connections |
{.full-width}

`--state` accepts known connection states, including `idle in transaction` (quote values with spaces).
`--query` matches query text without case sensitivity. A positive `--pid` selects that process
and takes precedence over other filters. Without filters, the command selects other client sessions;
`--all` also includes non-client backends.

### pg clone

Create a database copy in the same instance. The command terminates sessions on the source database
before copying, so schedule it when that interruption is acceptable. It uses `CREATE DATABASE ... TEMPLATE`;
PostgreSQL 15+ adds `STRATEGY FILE_COPY`. PostgreSQL 14 is also supported with its default copy strategy.

```bash
pig pg clone meta                       # Clone meta as meta_1/meta_2/...
pig pg clone meta meta_fork            # Clone to a specific database name
pig pg clone meta meta_fork --owner dba # Try to change new database owner
pig pg clone meta meta_fork --port 5433 # Connect to a specific local port
pig pg clone meta meta_fork --plan      # Preview clone plan
```

**Options:**

| Option | Short | Description |
|:-------|:------|:------------|
| `--port` | | PostgreSQL port, default `5432` or `$PG_PORT` |
| `--conn-db` | | Connection database: `postgres`, or `template1` when cloning `postgres` |
| `--owner` | | Try to change the owner of the cloned database |
| `--conn-limit` | | Connection limit for the new database (`-1` unlimited, `0` disallow connections) |
| `--plan` | | Show execution plan only |
| `--yes` | `-y` | Skip confirmation prompt |
{.full-width}

On PostgreSQL 18+ with `file_copy_method=clone` and a compatible filesystem, the copy can use
copy-on-write (CoW). Otherwise it uses an ordinary copy. Use `pg fork` to copy an entire instance.
Filesystem preflight runs as the selected database system user, including for a private `0700` data directory.
Text mode requests confirmation; structured execution requires `--yes`.

## Database Maintenance Commands

The `[dbname]` operand of `vacuum`, `analyze`, `freeze`, and `repack` accepts a database name.
Values containing `=` or starting with `postgres://` or `postgresql://` are rejected as connection
strings before planning, confirmation, or execution; errors do not echo the value.
With `--all`, `vacuum`, `analyze`, and `freeze` validate the full database list before any maintenance runs.
Use the native client when you need a connection string. Spaces and non-ASCII database names remain supported.
For names beginning with `-`, put all options first and use `--` before the name, for example
`pig pg vacuum --full --plan -- --all` previews the single database named `--all`.

### pg vacuum

Run SQL `VACUUM` through psql. The default database is `postgres`.

```bash
pig pg vacuum                     # Vacuum the postgres database
pig pg vacuum mydb                # Vacuum specific database
pig pg vacuum -a                  # Vacuum all databases
pig pg vacuum mydb -t mytable     # Vacuum specific table
pig pg vacuum mydb --schema myschema -t mytable  # Vacuum one table in a schema
pig pg vacuum mydb --full         # VACUUM FULL (requires exclusive lock)
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--all` | `-a` | Process all databases |
| `--schema` | | Specify schema |
| `--table` | `-t` | Specify table |
| `--verbose` | `-V` | Verbose output |
| `--full` | `-F` | VACUUM FULL (requires exclusive lock) |
| `--plan` | | Show the operation without executing it |
| `--yes` | `-y` | Skip VACUUM FULL confirmation |
{.full-width}

For `vacuum` and `freeze`, combine `--schema` with `--table` to select a table.
To process a whole schema, use native `vacuumdb --schema SCHEMA`.

`--full` rewrites tables and needs an exclusive lock. Text mode asks for confirmation;
structured execution requires `--yes`.

### pg analyze

Run SQL `ANALYZE` through psql to update planner statistics. The default database is `postgres`.

```bash
pig pg analyze                    # Analyze the postgres database
pig pg ana                        # Alias
pig pg analyze mydb               # Analyze specific database
pig pg analyze -a                 # Analyze all databases
pig pg analyze mydb -t mytable    # Analyze specific table
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--all` | `-a` | Process all databases |
| `--schema` | | Specify schema |
| `--table` | `-t` | Specify table |
| `--verbose` | `-V` | Verbose output |
{.full-width}

### pg freeze

Run SQL `VACUUM (FREEZE)` through psql to freeze old transaction IDs. The default database is `postgres`.

```bash
pig pg freeze                     # Freeze the postgres database
pig pg freeze mydb                # Freeze specific database
pig pg freeze -a                  # Freeze all databases
```

**Options:** Same as `pg analyze`.

### pg repack

Rebuild tables with `pg_repack`. Install the executable and the extension in each target database.
The default database is `postgres`.

```bash
pig pg repack mydb                # Repack all tables in database
pig pg repack -a                  # Repack all databases
pig pg repack mydb -t mytable     # Repack specific table
pig pg repack mydb --schema myschema  # Repack tables in schema
pig pg repack mydb -j 4           # Use 4 parallel jobs
pig pg repack mydb --plan         # Preview target and command without running pg_repack
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--all` | `-a` | Process all databases |
| `--schema` | | Specify schema |
| `--table` | `-t` | Specify table |
| `--verbose` | `-V` | Verbose output |
| `--jobs` | `-j` | Number of parallel jobs (default 1) |
| `--plan` | | Show tables to be repacked |
{.full-width}

## Parameter Tuning Commands

### pg tune

Generate starting values for memory, parallelism, and WAL settings. PIG uses the PostgreSQL version,
CPU count, memory, disk capacity, and selected workload profile. The storage formulas assume SSDs.
The command prints recommendations; it does not edit configuration files.

```bash
pig pg tune                       # auto-detect hardware, use oltp profile
pig pg tuning                     # alias
pig pg tune -p olap               # use OLAP profile
pig pg tune -p tiny               # for small instances
pig pg tune -c 8 -m 32768 -d 500  # override hardware detection
pig pg tune -C 500                # override max_connections
pig pg tune -R 0.30               # adjust shared_buffers ratio
pig pg tune -o json               # structured JSON output
pig pg tune -o yaml               # structured YAML output
```

**Options:**

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--profile` | `-p` | oltp | Tuning profile: `oltp` / `olap` / `tiny` / `crit` |
| `--cpu` | `-c` | 0 | CPU cores, `0` means auto-detect |
| `--mem` | `-m` | 0 | Total memory in MB, `0` means auto-detect |
| `--disk` | `-d` | 0 | Data disk size in GB, `0` means auto-detect |
| `--max-conn` | `-C` | 0 | Override `max_connections`, `0` uses profile default |
| `--shmem-ratio` | `-R` | 0.25 | Fraction of memory used for `shared_buffers`, range `0.1 ~ 0.4` |
{.full-width}

**Profiles:**

| Profile | Best for | Characteristics |
|:---|:---|:---|
| `oltp` | General transactional workloads | Balanced connection count, cache, and parallelism |
| `olap` | Analytical workloads | More aggressive parallelism and work memory |
| `tiny` | Small instances | Constrained memory footprint and parallelism |
| `crit` | Latency-sensitive workloads | Restricts parallel gather and favors stable response time |
{.full-width}

**Notes:**

- Generated parameters are automatically gated by PostgreSQL major version. For example, `io_workers` is only emitted for PG 18+.
- Text output includes a success message. Copy only the parameter assignments into your configuration;
  do not redirect the complete output into a PostgreSQL config file.
- For scripts, use `-o json` and read `data.parameters`. Review the values, apply them through your
  configuration workflow, then reload or restart PostgreSQL as required.

## Instance Fork

### pg fork

Create a local one-off PostgreSQL physical copy for temporary analysis, troubleshooting, recovery validation, and development testing. Managed forks are written to `/pg/data-<name>` by default and are not registered with Pigsty, systemd, or Patroni. When `--dst-data` is specified explicitly, the command creates an unmanaged fork that is not enumerated by `fork list`.

```bash
pig pg fork dev                       # Create /pg/data-dev, do not start
pig pg fork init dev --start          # Create and start, probing ports from 15432
pig pg fork init dev -s --dst-port 15433    # Create and start on specified port
pig pg fork init dev -D /pg/data2 --src-port 15431  # Specify source dir and source port
pig pg fork init dev --dst-data /tmp/dev    # Create unmanaged fork
pig pg fork list                      # List managed forks
pig pg fork start dev                 # Start existing managed fork
pig pg fork stop dev                  # Stop existing managed fork
pig pg fork rm dev --stop             # Stop and remove a running fork
pig pg fork init dev --plan           # Show execution plan only
```

**Create Options:**

| Option | Short | Default | Description |
|:-------|:------|:--------|:------------|
| `--dst-data` | | `/pg/data-<name>` | Unmanaged target data directory |
| `--dst-port` | | auto-detect | Target port, probes free ports starting at 15432 |
| `--src-data` | | `/pg/data` or `$PG_DATA` | Source data directory; can also be set globally with `pg -D/--data` |
| `--src-port` | | `5432` or `$PG_PORT` | Source port |
| `--start` | `-s` | false | Start the fork after creation |
| `--force` | `-f` | false | Overwrite an existing stopped target directory and skip confirmation |
| `--timeout` | `-t` | 60 | Startup wait timeout in seconds |
| `--yes` | `-y` | false | Skip confirmation prompt |
| `--plan` | | false | Show execution plan only |
{.full-width}

**Management Commands:**

| Command | Common Options | Description |
|:--------|:---------------|:------------|
| `pig pg fork list` | `-o json/yaml` | List managed forks with optional structured output |
| `pig pg fork start <name> or --dst-data <dir>` | `--dst-data`, `--dst-port`, `-t/--timeout`, `--plan` | Start existing fork |
| `pig pg fork stop <name> or --dst-data <dir>` | `--dst-data`, `-m/--mode`, `-t/--timeout`, `--plan` | Stop existing fork |
| `pig pg fork rm <name> or --dst-data <dir>` | `--dst-data`, `--stop`, `-m/--mode`, `-t/--timeout`, `-f/--force`, `-y/--yes`, `--plan` | Remove fork; running forks require `--stop` |
{.full-width}

**Behavior Notes:**

- When the source instance is running, the command uses PostgreSQL low-level backup APIs to create a consistent physical copy; when the source is stopped, it can perform a cold copy.
- The command prefers CoW/reflink. If only ordinary copy is available, interactive mode warns about disk-space risk and waits for confirmation.
- The target cannot be `/`, `/pg`, source PGDATA, or a parent/child of source PGDATA. Symlinks are resolved before checking the source/destination relationship. `/pg/data` follows the same ownership, relationship, and runtime checks as other data directories.
- Replacing a nonempty unmanaged target requires valid unmanaged `pg_fork` metadata in `fork.json` with an explicitly recorded `target.data` matching the directory. An arbitrary directory without `PG_VERSION` is not eligible for deletion. Missing, damaged, incomplete, or mismatched metadata is refused. An empty directory remains usable with `--force`. Managed forks must be replaced by name. A running or uncertain destination cannot be replaced. Forced removal by name still permits cleanup of a managed orphan under `/pg/data-<name>`.
- Path validation resolves private directories as DBSU when the caller cannot traverse them; unresolved paths and source/destination aliases are refused.
- Unmanaged fork removal checks an inaccessible target as the selected DBSU. The target must be a real directory; a symlink or failed check still prevents removal.
- After copy, runtime and replication state is cleaned from the fork and `fork.json` is written. The new instance starts only when `-s|--start` is specified.
- Managed forks must be managed by name. Unmanaged forks require `--dst-data` when starting, stopping, or removing.

**List Forks:**

`pig pg fork list` scans `/pg/data-*`, reads `fork.json`, and refreshes each managed fork's runtime state. Text output shows `NAME`, `PORT`, `STATE`, `PID`, `AGE`, `SOURCE`, `COPY`, and `DATA`; state is reported as `running`, `stopped`, `orphan`, or `unknown`. Unknown state sets `runtime_unknown: true` in structured output without claiming that the fork has started. Lifecycle operations fail without rewriting metadata in that state.

JSON and YAML put fork records in the result's `data` array. Records include source and target
endpoints, copy and backup metadata, management commands, and PIG build metadata.

**Structured Output:**

```bash
pig pg fork init dev --plan -o yaml
pig pg fork list -o json
```

## Log Commands

Log commands view PostgreSQL log files. Default log directory is `/pg/log/postgres`, can be changed via `--log-dir`.

**Log Command Global Options:**

| Option | Description |
|:---|:---|
| `--log-dir` | Log directory path (default: `/pg/log/postgres`) |
| `--lines` / `-n` | Number of lines to show, default 50 |
| `--follow` / `-f` | Follow latest log, only on parent `pg log` |
{.full-width}

If the log directory is not readable, PIG retries through `sudo`.
Use `pg log -o json` or `pg log show -o json` for JSONL snapshots.
Snapshots do not support YAML or `json-pretty`; `tail`, `less`, and `grep` use text output.

### pg log

Show the latest log snapshot; with `-f`, follow the latest log.

```bash
pig pg log                        # Show latest 50 lines
pig pg log -n 100                 # Show latest 100 lines
pig pg log -f                     # Follow latest log
```

### pg log list

List log files in log directory.

```bash
pig pg log list                              # List logs in default directory
pig pg log list --log-dir /var/log/postgres  # List logs in specified directory
```

### pg log tail

Real-time log viewing (like `tail -f`). Default views latest CSV log file.

```bash
pig pg log tail                   # View latest log
pig pg log t                      # Alias
pig pg log f                      # Alias
pig pg log tail postgresql.csv    # View specific log file
pig pg log tail -n 100            # Show last 100 lines then follow
pig pg log tail --log-dir /var/log/postgres  # Use custom directory
```

**Options:**

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--lines` | `-n` | 50 | Number of lines to show |
| `--follow` | `-f` | false | Accepted for compatibility; `tail` always follows |
{.full-width}

### pg log show

Output log file content.

```bash
pig pg log show                   # Output latest log
pig pg log cat                    # Alias for show
pig pg log c                      # Alias for show
pig pg log show -n 100            # Output last 100 lines
pig pg log show postgresql.csv    # Output specific log file
```

**Options:**

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--lines` | `-n` | 50 | Number of lines to show |
{.full-width}

### pg log less

Open log file with less. Defaults to end of file (`+G`).

```bash
pig pg log less                   # Open latest log with less
pig pg log vi                     # Alias
pig pg log v                      # Alias
pig pg log less postgresql.csv    # Open specific log file
```

### pg log grep

Search log file content.

```bash
pig pg log grep ERROR             # Search for ERROR
pig pg log grep --ignore-case error  # Ignore case
pig pg log grep -C 3 ERROR        # Show context
pig pg log grep ERROR pg.csv      # Search specific log file
```

**Options:**

| Option | Short | Description |
|:-------|:------|:------------|
| `--ignore-case` | | Ignore case |
| `--context` | `-C` | Show context lines |
{.full-width}

## pg svc Subcommand

`pg svc` (also `pg service` or `pg s`) controls the `postgres` systemd unit used by Pigsty.
It requires that unit to exist. For distribution units such as `postgresql-18` or
`postgresql@18-main`, use `systemctl` with the actual unit name.

```bash
pig pg svc start                 # Start postgres service
pig pg svc stop                  # Stop postgres service
pig pg svc restart               # Restart postgres service
pig pg svc reload                # Reload postgres service
pig pg svc status                # Show service status
```

**Alias Reference:**

| Command | Alias |
|:--------|:------|
| `pg svc start` | `boot, up` |
| `pg svc stop` | `halt, dn, down` |
| `pg svc restart` | `reboot, rt` |
| `pg svc reload` | `rl, hup` |
| `pg svc status` | `st, stat` |
{.full-width}

## Tools and privileges

PIG uses `initdb` to initialize data, `pg_ctl` to control instances, psql for SQL and maintenance,
`pg_repack` to rebuild tables, and system tools to read logs. Use those tools directly for options
that this command does not expose.

PostgreSQL commands run directly when you are the selected database operating-system user.
From root, PIG uses `su`; from another account, it uses `sudo -H -u`. Service management uses root privileges.
These commands target Linux; service management requires systemd.

## Design records

- [The 70% Tuner: Defining the Boundary of pig pg tune](https://pig.pgsty.com/design/pg-tune/)
- [PIG's Operations CLI Safety Contract](https://pig.pgsty.com/design/ops-cli-safety/)
