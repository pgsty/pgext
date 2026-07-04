---
title: "CMD: pitr"
description: Perform orchestrated Point-In-Time Recovery with pig pitr command
icon: SquareTerminal
weight: 680
---

The `pig pitr` command performs point-in-time recovery through pgBackRest and conservatively handles the local PostgreSQL and Patroni lifecycle. Unlike the lower-level `pig pb restore`, `pig pitr` runs pre-restore checks, stops Patroni and PostgreSQL when needed, executes restore, then decides whether to start PostgreSQL based on the selected options.

Note: for a managed default data directory, `pig pitr` leaves Patroni stopped after recovery. Validate the restored result first, then manually restore Patroni management. This command does not automatically rejoin a Patroni cluster, perform failover, or validate cluster member state.

```bash
pig pitr - Perform PITR with pgBackRest restore and conservative PostgreSQL stop/start handling.

For the managed default data directory, this command may:
  1. Stop Patroni only to keep the target PGDATA offline during restore
  2. Ensure PostgreSQL is stopped (fast stop with retry; destructive fallback only with --force-stop)
  3. Execute pgbackrest restore
  4. Start PostgreSQL unless --no-restart is used
  5. Leave Patroni stopped; provide post-restore guidance

Recovery Targets (at least one required):
  --default, -d      Recover to end of WAL stream (latest)
  --immediate, -I    Recover to backup consistency point
  --time, -t         Recover to specific timestamp
  --name             Recover to named restore point
  --lsn              Recover to specific LSN
  --xid              Recover to specific transaction ID

Backup and Target Options:
  --set, -b          Select backup set to start recovery from
  --target-action    Action when target is reached: pause, promote, shutdown
  --target-timeline  Recover along timeline: latest, current, N, or 0xN

Use --no-restart with --target-action=shutdown because PostgreSQL exits
after reaching the recovery target.

Custom -D side restores require --no-restart. The custom directory must
already exist, be owned by the configured DBSU (postgres by default), and
be writable by that user. Pig does not create this directory automatically.

Additional pgBackRest arguments:
  Put raw pgBackRest restore arguments after -- so Cobra stops parsing them.
  Example: pig pitr -d -- --delta

Time Format:
  - Full: "2025-01-01 12:00:00+08"
  - Date only: "2025-01-01" (defaults to 00:00:00)
  - Time only: "12:00:00" (defaults to today)

Examples:
  pig pitr -d                                      # Recover to latest
  pig pitr -t "2025-01-01 12:00:00+08"            # Recover to time
  pig pitr -d --plan                               # Preview plan
  pig pitr -d -y                                   # Skip y/yes confirmation
  pig pitr -d --no-restart                         # Leave PostgreSQL stopped
  pig pitr -d -D /tmp/pg-restore --no-restart      # Side restore
```


## Overview

The default target of `pig pitr` is the Pigsty-managed primary data directory. A typical workflow:

1. Validate recovery target parameters. One of `-d/-I/-t/--name/--lsn/--xid` is required.
2. Resolve pgBackRest configuration and target data directory.
3. For default data-directory recovery, stop Patroni if it is running.
4. Ensure PostgreSQL is stopped.
5. Run `pgbackrest restore`.
6. Start PostgreSQL unless `--no-restart` is specified.
7. Print post-recovery validation and Patroni recovery guidance.

**Comparison with `pig pb restore`:**

| Feature | `pig pitr` | `pig pb restore` |
|:---|:---|:---|
| Stop Patroni | Automatic for default data-directory recovery | Manual |
| Stop PostgreSQL | Checks and stops automatically | Must be pre-stopped |
| Start PostgreSQL | Automatic by default, can be disabled with `--no-restart` | Manual |
| Patroni recovery | Not automatic; restore manually after verification | Not handled |
| Use case | Production recovery orchestration | Low-level restore or scripting |
{.full-width}


## Quick Start

```bash
# Most common: recover to the end of the WAL stream
pig pitr -d

# Recover to a specific point in time
pig pitr -t "2025-01-01 12:00:00+08"

# Recover to the backup consistency point
pig pitr -I

# Preview the execution plan without running
pig pitr -d --plan

# Skip confirmation, suitable for automation
pig pitr -d -y

# Recover from a specific backup set
pig pitr -d -b 20251225-120000F

# Restore the default data directory but do not start PostgreSQL afterward
pig pitr -d --no-restart

# Side restore: restore to a custom directory without touching Patroni or /pg/data
install -d -m 700 -o postgres -g postgres /tmp/pg-restore
pig pitr -d -D /tmp/pg-restore --no-restart

# Extra native pgBackRest restore arguments go after --
pig pitr -d -- --delta
```


## Parameters

### Recovery Target

At least one recovery target is required.

| Parameter | Short | Description |
|:---|:---|:---|
| `--default` | `-d` | Recover to the end of the WAL stream, the latest data |
| `--immediate` | `-I` | Recover to the backup consistency point |
| `--time` | `-t` | Recover to a specific timestamp |
| `--name` | | Recover to a named restore point |
| `--lsn` | | Recover to a specific LSN |
| `--xid` | | Recover to a specific transaction ID |
{.full-width}

### Backup And Target Options

| Parameter | Short | Description |
|:---|:---|:---|
| `--set` | `-b` | Start recovery from a specific backup set |
| `--target-action` | | Action when the recovery target is reached: pause/promote/shutdown |
| `--target-timeline` | `-T` | Recovery timeline: latest/current/N/0xN |
| `--exclusive` | `-X` | Exclusive mode: stop before target |
{.full-width}

Use `--target-action=shutdown` with `--no-restart`, because PostgreSQL exits after reaching that recovery target. `--target-action` cannot be used with `--default`, because `--default` already means recovery to the end of WAL. `--exclusive/-X` requires a precise stop-before target: `--time`, `--lsn`, or `--xid`.

Native pgBackRest restore arguments after `--` cannot override recovery target, lifecycle, data directory, repository, config, or selection parameters managed by Pig. Use Pig's first-class parameters for these semantics. This restriction is consistent with the `pig pb restore` passthrough blocklist.

### Flow Control

| Parameter | Short | Description |
|:---|:---|:---|
| `--no-restart` | | Do not start PostgreSQL after restore |
| `--plan` | | Show execution plan only, do not execute |
| `--yes` | `-y` | Skip interactive y/yes confirmation |
| `--timeout` | | PostgreSQL startup/recovery wait timeout, default 120 seconds |
| `--force-stop` | | Allow immediate shutdown and kill fallback if fast stop fails |
{.full-width}

### Configuration

| Parameter | Short | Description |
|:---|:---|:---|
| `--stanza` | `-s` | pgBackRest stanza name |
| `--config` | `-c` | pgBackRest config file path |
| `--repo` | `-r` | Repository number |
| `--dbsu` | `-U` | Database superuser, default `postgres` |
| `--data` | `-D` | Target data directory |
{.full-width}


## Time Format

The `--time` parameter supports multiple formats and completes missing parts using the current timezone:

| Format | Example | Description |
|:---|:---|:---|
| Full | `2025-01-01 12:00:00+08` | Complete timestamp with timezone |
| Date time without timezone | `2025-01-01 12:00:00` | Adds the current local timezone automatically; `T` separator is also accepted |
| Date only | `2025-01-01` | Completes to 00:00:00 on that date |
| Time only | `12:00:00` | Completes to that time today |
{.full-width}

Plan output and replayable next-action commands normalize date-only and time-only targets into deterministic timestamps with timezone. This rule is consistent with `pig pb restore --plan`.


## Managed Directory And Side Restore

The managed PostgreSQL data directory comes from the effective pgBackRest `pg1-path` and command parameters. It is not hardcoded to `/pg/data`. For example, if managed PGDATA is `/var/lib/pgsql/18/data`, the command still treats it as a managed restore. Path comparison resolves symlinks as the database superuser when needed, so a symlink to managed PGDATA is not mistaken for a side restore.

An explicit `-D/--data` that resolves to a path different from the managed directory is a side restore. Side restores must use `--no-restart`; they do not stop Patroni and do not manage the default PostgreSQL service. After restore, handle the side instance manually with commands such as `pg_ctl -D <dir> -o "-p 5433" start`, `pg_ctl -D <dir> status`, and `pgbackrest --pg1-path=<dir> stanza-create`. The side-restore directory must already exist and be owned by the DBSU. Unlike managed PGDATA, it does not need a pre-existing `PG_VERSION` marker. Pig does not create this directory automatically because it needs a concrete path for destructive-restore classification, owner checks, and safe guidance before restore.

```bash
install -d -m 700 -o postgres -g postgres /tmp/pg-restore
pig pitr -d -D /tmp/pg-restore --no-restart
```

For managed PGDATA outside `/pg/data`, post-restore runbook commands explicitly include the effective data directory:

```bash
pig pg start -D /var/lib/pgsql/18/data
pig pg psql -D /var/lib/pgsql/18/data
pig pg promote -D /var/lib/pgsql/18/data
```

`pig pg psql -D <dir>` reads `postmaster.pid` under that directory and uses the recorded port and socket directory to connect to the restored instance. If postmaster information cannot be resolved, it does not silently fall back to the default connection target.


## Execution Flow

### Phase 1: Pre-Checks

- Validate recovery target parameters; missing targets only print help and return an error.
- Resolve effective pgBackRest config, stanza, repository, and managed `pg1-path`.
- Check that the managed data directory exists and is initialized.
- For side restore, check that the custom data directory exists and is owned by the DBSU.
- Verify the selected stanza is healthy and has backups; if `--set` is specified, verify that backup set exists.
- Detect Patroni service state and PostgreSQL runtime state.

### Phase 2: Patroni Handling

For managed data-directory recovery, if Patroni is running, the command stops Patroni so the target PGDATA stays offline during restore. Patroni remains stopped after recovery. A custom `-D` side restore does not touch the managed data directory and therefore does not stop Patroni.

### Phase 3: Ensure PostgreSQL Is Stopped

The command first waits for PostgreSQL to exit after Patroni stops, then retries `pg_ctl stop -m fast`. If PostgreSQL still cannot stop, it does not use more aggressive methods by default. Only explicit `--force-stop` allows immediate shutdown and final kill fallback.

### Phase 4: Restore

The command runs pgBackRest restore and maps recovery target, backup set, timeline, target action, and related options into `pgbackrest restore`. Native pgBackRest arguments can be placed after `--`:

```bash
pig pitr -d -- --delta
```

### Phase 5: Start Or Stay Stopped

Unless `--no-restart` is specified, the command starts PostgreSQL after restore and waits for recovery completion. For `--default` and `--target-action=promote`, it waits until `pg_is_in_recovery()` becomes false on the recovered instance. Recovery and subsequent SQL probes bind to the port recorded in `postmaster.pid` for the restored data directory and use the socket directory there when present. Use `--no-restart` when:

- Running a custom `-D` side restore, because the restored config still keeps the original port and must be started manually on a free port.
- Using `--target-action=shutdown`, because PostgreSQL exits after reaching the recovery target.
- You need to inspect the restored directory before deciding whether to start.


## Examples

### Scenario 1: Recover Dropped Data

```bash
# 1. View available backups
pig pb info

# 2. Preview recovery plan
pig pitr -t "2025-01-15 09:30:00+08" --plan

# 3. Execute recovery
pig pitr -t "2025-01-15 09:30:00+08"

# 4. Verify data
pig pg psql
SELECT * FROM important_table;
```

### Scenario 2: Recover To Latest

```bash
pig pitr -d
```

### Scenario 3: Recover To Backup Consistency Point

```bash
pig pitr -I
```

### Scenario 4: Keep Stopped After Restore

```bash
pig pitr -d --no-restart

# PostgreSQL and Patroni both remain stopped; inspect directory or logs before starting
pig pg start
```

### Scenario 5: Custom Directory Side Restore

```bash
install -d -m 700 -o postgres -g postgres /tmp/pg-restore
pig pitr -d -D /tmp/pg-restore --no-restart

# Start the side restore manually on a free port
pg_ctl -D /tmp/pg-restore -o "-p 5433" start
```


## Execution Plan Example

Running `pig pitr -d --plan` shows a plan like:

```text
══════════════════════════════════════════════════════════════════
 PITR Execution Plan
══════════════════════════════════════════════════════════════════

Current State:
  Data Directory:  /pg/data
  Database User:   postgres
  Patroni Service: active
  PostgreSQL:      running (PID: 12345)

Recovery Target:
  Latest (end of WAL stream)

Execution Steps:
  [1] Stop Patroni service
  [2] Ensure PostgreSQL is stopped
  [3] Execute pgBackRest restore
  [4] Start PostgreSQL
  [5] Print post-restore guidance

══════════════════════════════════════════════════════════════════
```


## Post-Recovery Actions

After successful recovery, verify data before restoring orchestration:

```bash
# Verify data
pig pg psql

# If recovered to a manual target and promotion is needed
pig pg promote

# Restore Patroni management only after validation
systemctl start patroni

# Re-create pgBackRest stanza if needed
pig pb create
```

Leaving Patroni stopped after managed data-directory recovery is intentional. It prevents restored old state from re-entering HA orchestration before validation.


## Safety Mechanisms

**Recovery target required:** Without `-d/-I/-t/--name/--lsn/--xid`, the command only shows help and does not run restore.

**Confirmation:** In text mode, destructive recovery asks for interactive `y`/`yes` confirmation before execution. Automation can use `-y|--yes`. Structured output mode does not prompt interactively; it must use `--yes` to execute or `--plan` to preview.

**Patroni boundary:** For managed data-directory recovery, the command stops Patroni when needed to prevent Patroni from restarting PostgreSQL during restore. It does not automatically rejoin Patroni after recovery.

**Side restore boundary:** A custom `-D` side restore must use `--no-restart` because the restored PostgreSQL config still uses the original port. Side restore does not manage Patroni or the default PostgreSQL service.

**Failure boundary:** If restore fails after Patroni has been stopped, Patroni remains stopped and the target data directory may be partially restored. Fix the underlying issue and rerun PITR, or validate the recovery state first; do not start Patroni before confirmation. If restore has run but PostgreSQL startup fails, inspect PostgreSQL and pgBackRest logs and validate the data directory before deciding whether to restore Patroni management.

**Structured output:** Structured execution requires `--yes`; `--plan` is the preview path. After successful execution, structured PITR results put post-restore actions in the Result envelope's `next_actions`, not inside `data`. `data` includes `requested_data_dir`, `effective_data_dir`, `managed_data_dir`, and `side_restore` so automation can distinguish user input from the actual recovery target.


## Design Notes

- `pig pitr` calls pgBackRest restore and handles local Patroni/PostgreSQL stop and optional start.
- `pig pitr` is not a cluster recovery controller. It does not handle Patroni failover, rejoin, VIP, or application traffic switching.
- Use [`pig pb restore`](/pig/cmd/pb/#pb-restore/) when you need lower-level restore semantics or fine-grained scripting control.
- Use [`pig pt switchover`](/pig/cmd/pt/#pt-switchover/) or [`pig pt failover`](/pig/cmd/pt/#pt-failover/) when you need manual Patroni cluster switching.

**Privilege execution:**

- If the current user is the DBSU: execute directly.
- If the current user is root: execute with `su - postgres -c`.
- Other users: execute with `sudo -inu postgres --`.

**Platform support:**

This command is designed for Linux and depends on pgBackRest, systemd for managed service scenarios, and data/log paths accessible by the DBSU.
