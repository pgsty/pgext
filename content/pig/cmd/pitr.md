---
title: "CMD: pitr"
description: Perform orchestrated Point-In-Time Recovery with pig pitr command
icon: fa-solid fa-terminal
weight: 680
---

Use `pig pitr` to restore a local PostgreSQL instance from pgBackRest backups.
For a managed restore, PIG stops services as needed and restores the data.
By default, it then starts PostgreSQL and leaves Patroni stopped for you to verify the recovered database.

Recovery replaces data in the target directory. Start with `--plan` to check the directory,
backup source, recovery target, and service actions.

## Before recovery

You need an installed pgBackRest, a readable configuration file, a usable backup,
and the archived WAL needed to reach the target. The default config path is
`/etc/pgbackrest/pgbackrest.conf`. Select another file with `-c` and a stanza with `-s`.

Run the commands on the host that contains the target data directory. PIG runs PostgreSQL and
pgBackRest operations as `--dbsu`, then `$PIG_DBSU`, then `postgres`.
Service operations require root privileges, directly or through sudo.

| Task | `pig pitr` | `pig pb restore` |
|:---|:---|:---|
| Stop Patroni for a managed restore | Automatic if active | Operator stops it |
| Stop PostgreSQL | Checks and stops it | Requires it to be stopped |
| Restore backup files | pgBackRest | pgBackRest |
| Start PostgreSQL | Automatic unless `--no-restart` | Operator starts it |
| Return the instance to Patroni | Operator verifies and resumes management | Operator handles it |
{.full-width}

Use [`pig pb restore`](/pig/cmd/pb/#pb-restore) when you want to manage service shutdown and startup yourself.

## Quick start

This example restores an existing managed data directory to a chosen time.
Replace the timestamp with your recovery target:

```bash
# Inspect available backups
pig pb info

# Check the recovery plan
pig pitr -t "2025-01-15 09:30:00+08" --plan -- --delta

# Restore the data and start PostgreSQL
pig pitr -t "2025-01-15 09:30:00+08" -- --delta

# Connect to the recovered instance; use its actual data directory
pig pg psql -D /pg/data
```

`--delta` lets pgBackRest restore over an existing directory by comparing file checksums.
PIG asks for `y` or `yes` before recovery. Use `--yes` for execution without a prompt.

For a time, name, LSN, transaction ID, or immediate target, PostgreSQL normally pauses when it
reaches the target. Inspect the data before promotion. See [post-recovery actions](#post-recovery-actions).

## Parameters

### Recovery target

Choose exactly one target. Without a target, PIG returns an error and shows help in text mode.

| Option | Short | Description |
|:---|:---|:---|
| `--default` | `-d` | Recover to the end of available WAL |
| `--immediate` | `-I` | Recover to the backup consistency point |
| `--time` | `-t` | Recover to a timestamp |
| `--name` | | Recover to a named restore point |
| `--lsn` | | Recover to an LSN, such as `0/7C82CB8` |
| `--xid` | | Recover to a positive transaction ID |
{.full-width}

`--set` selects a backup; it does not select a recovery target.

### Backup and target options

| Option | Short | Description |
|:---|:---|:---|
| `--set` | `-b` | Backup label from `pig pb info` |
| `--target-action` | | Action at the target: `pause`, `promote`, or `shutdown` |
| `--target-timeline` | `-T` | Timeline: `latest`, `current`, a positive integer, or `0xHEX` |
| `--exclusive` | `-X` | Stop before a time, LSN, or transaction ID target |
{.full-width}

Use `--no-restart` with `--target-action=shutdown`. Start PostgreSQL manually when you are ready
for WAL replay; it will shut down after reaching the target.

`--target-action` cannot be combined with `--default`. `--exclusive` requires `--time`, `--lsn`, or `--xid`.

### Flow control

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--no-restart` | | false | Leave PostgreSQL stopped after restoring files |
| `--plan` | | false | Inspect the plan without changing data or services |
| `--yes` | `-y` | false | Skip confirmation |
| `--timeout` | | 120 | Startup and recovery wait timeout, in seconds |
| `--force-stop` | | false | Allow immediate shutdown and process termination if fast shutdown fails |
{.full-width}

### Configuration

| Option | Short | Default or source |
|:---|:---|:---|
| `--stanza` | `-s` | First stanza in the config file; specify it when several exist |
| `--config` | `-c` | `/etc/pgbackrest/pgbackrest.conf` |
| `--repo` | `-r` | Repository selected by pgBackRest |
| `--dbsu` | `-U` | PostgreSQL operating-system user: `$PIG_DBSU`, then `postgres` |
| `--data` | `-D` | Target directory; configured `pg1-path`, then `$PGDATA`, then `/pg/data` |
{.full-width}

### Native restore options

Put extra pgBackRest arguments after `--`:

```bash
pig pitr -d -- --delta
```

Use PIG flags for the recovery target, data directory, config, stanza, and repository.
Arguments that override these settings are rejected, including `--repoN-*` and `--recovery-option`.
Other restore options, such as `--delta`, `--tablespace-map`, `--link-map`, and `--link-all`, can be forwarded.

## Time format

PIG fills in missing date or timezone fields using the local date and timezone:

| Format | Example | Interpretation |
|:---|:---|:---|
| Full timestamp | `2025-01-01 12:00:00+08` | Explicit date, time, and timezone |
| No timezone | `2025-01-01 12:00:00` | Local timezone; a `T` separator also works |
| Date only | `2025-01-01` | Midnight on that date |
| Time only | `12:00:00` | That time today |
{.full-width}

Use a full timestamp with timezone in scripts. Plans and suggested follow-up commands expand
short forms to a full timestamp so that rerunning the command keeps the same target.

## Managed directory and side restore

The managed directory comes from `pg1-path` in the selected stanza, then `$PGDATA`, then `/pg/data`.
It must already contain an initialized PostgreSQL instance. An explicit `-D` that resolves to this
same path is still a managed restore, including a symlink to it.

An explicit `-D` that resolves to a different path selects a **side restore**.
PIG leaves Patroni and the managed instance alone. The side directory must exist, belong to the
selected operating-system user, and be writable by that user. It can be empty.

The `-D` option selects only the data directory. If the backup contains tablespaces or external links, map them to separate paths.
Add pgBackRest's `--tablespace-map` and `--link-map` options after `--` in the restore command.

A side restore requires `--no-restart`. The restored configuration retains the original port;
start it manually with a free port:

```bash
# Prepare the directory as root or through sudo
sudo install -d -m 700 -o postgres -g postgres /tmp/pg-restore

# Restore backup files into the side directory
pig pitr -d -D /tmp/pg-restore --no-restart

# Start the side instance on a free port
pig pg start -D /tmp/pg-restore -O "-p 5433"
pig pg psql -D /tmp/pg-restore
```

`pig pg psql -D` reads the restored directory's `postmaster.pid` to select its port and socket.
It returns an error if it cannot resolve that instance.

## Execution flow

1. Validate options, configuration, target directory, stanza health, and available backups.
   If `--set` is given, check that the backup label exists.
2. Show the operation and request confirmation, unless `--yes` is given.
3. For a managed restore, stop an active Patroni service.
4. Wait for PostgreSQL to stop, then retry fast shutdown if needed. Use `--force-stop` to allow
   immediate shutdown and a final process-termination attempt.
5. Run `pgbackrest restore`.
6. Start PostgreSQL unless `--no-restart` is given. For `--default` or `--target-action=promote`,
   wait until `pg_is_in_recovery()` is false. For other targets, verify recovery progress yourself.
7. Print commands for data verification and the remaining recovery steps.

Declining confirmation or reaching EOF returns `CodePITRConfirmationRequired` (`161001`, exit
status `7`) before stopping Patroni or PostgreSQL, restoring data, or starting PostgreSQL.

An indeterminate PostgreSQL runtime state is an error, not proof that the instance is stopped.
PIG will not restore over an unverified instance or use an unverified PID for forced termination.
After restoration, an unknown runtime state cannot count as successful startup and prevents
follow-up SQL checks against an unverified connection target.

`--plan` runs the prechecks and returns the planned actions. It needs access to the configuration,
data directory, and backup repository, but does not stop services or restore data.

## More examples

Run one recovery command for the target you need:

```bash
# Latest available state, restoring over an existing directory
pig pitr -d -- --delta

# Backup consistency point
pig pitr -I -- --delta

# Use a specific backup label
pig pitr -d -b 20251225-120000F -- --delta

# Restore files, then inspect them before starting PostgreSQL
pig pitr -d --no-restart -- --delta

# Promote automatically after a time target
pig pitr -t "2025-01-15 09:30:00+08" --target-action=promote -- --delta
```

## Post-recovery actions {#post-recovery-actions}

1. Connect to the restored instance. Use its actual data directory:

   ```bash
   pig pg psql -D /pg/data
   ```

2. Check recovery state, then verify the tables and records you need:

   ```sql
   SELECT pg_is_in_recovery() AS recovering,
          CASE WHEN pg_is_in_recovery() THEN pg_is_wal_replay_paused() END AS paused;
   ```

3. If recovery is paused at the intended target and you want a writable primary, promote it:

   ```bash
   pig pg promote -D /pg/data
   ```

4. For a Patroni-managed cluster, reconcile the restored instance with the cluster's recovery plan,
   then resume Patroni with `pig pt svc start`. Handle other members and application routing
   as part of that plan.

If a managed restore fails after stopping Patroni, Patroni stays stopped. The target directory can contain a partial restore. Inspect PostgreSQL and pgBackRest logs, fix the failure, and complete recovery before
starting Patroni. A startup timeout also needs investigation; it does not undo the restored files.

## Structured output

Use `--plan -o json` to inspect a structured plan. Use `--yes -o json` to execute without interaction.
YAML is also available.

Successful results include `requested_data_dir`, `effective_data_dir`, `managed_data_dir`, and
`side_restore` in `data`. Follow-up commands are in the result's top-level `next_actions` field.

## Design record

[PIG's Operations CLI Safety Contract](https://pig.pgsty.com/design/ops-cli-safety/) explains the separate roles of
`pig pb restore` and `pig pitr`.
