---
title: "CMD: pitr"
description: Perform orchestrated Point-In-Time Recovery with pig pitr command
icon: SquareTerminal
weight: 680
---

The `pig pitr` command performs **Orchestrated Point-In-Time Recovery** through pgBackRest restore with conservative PostgreSQL stop/start handling. Unlike the lower-level `pig pb restore`, it performs pre-checks, stops Patroni/PostgreSQL when needed, runs restore, then starts PostgreSQL unless `--no-restart` is used.

For the managed default data directory, Patroni is left stopped after PITR. Validate the restored database first, then resume Patroni outside this command. `pig pitr` does not rejoin Patroni, perform failover, or validate cluster membership after restore.

```bash
pig pitr - Perform PITR with pgBackRest restore and conservative PostgreSQL stop/start handling.

For the managed default data directory, this command may:
  1. Stop Patroni only to keep the target PGDATA offline during restore
  2. Ensure PostgreSQL is stopped (with retry and fallback)
  3. Execute pgbackrest restore
  4. Start PostgreSQL unless --no-restart is used
  5. Provide post-restore guidance

Recovery Targets (at least one required):
  --default, -d      Recover to end of WAL stream (latest)
  --immediate, -I    Recover to backup consistency point
  --time, -t         Recover to specific timestamp
  --name, -n         Recover to named restore point
  --lsn, -l          Recover to specific LSN
  --xid, -x          Recover to specific transaction ID

Backup and Target Options:
  --set, -b          Select backup set to start recovery from
  --target-action    Action when target is reached: pause, promote, shutdown
  --target-timeline  Recover along timeline: latest, current, N, or 0xN

Time Format:
  - Full: "2025-01-01 12:00:00+08"
  - Date only: "2025-01-01" (defaults to 00:00:00)
  - Time only: "12:00:00" (defaults to today)

Examples:
  pig pitr -d                      # Recover to latest (most common)
  pig pitr -t "2025-01-01 12:00:00+08"  # Recover to specific time
  pig pitr -I                      # Recover to backup consistency point
  pig pitr -d --plan               # Show execution plan without running
  pig pitr -d -y                   # Skip confirmation (for automation)
  pig pitr -d --no-restart         # Leave PostgreSQL stopped after restore
  pig pitr -d -D /tmp/pg-restore -S -N  # Side restore to custom dir
```


## Overview

`pig pitr` targets Pigsty-managed local restore workflows:

1. Validate exactly one recovery target (`-d/-I/-t/-n/-l/-x`)
2. Resolve pgBackRest config, stanza, repo, and target data directory
3. For managed data-dir restore, stop Patroni when it is running
4. Ensure PostgreSQL is stopped
5. Execute `pgbackrest restore`
6. Start PostgreSQL unless `--no-restart` is used
7. Print verification and Patroni-resume guidance

**Comparison with `pig pb restore`:**

| Feature | `pig pitr` | `pig pb restore` |
|:--------|:-----------|:-----------------|
| Stop Patroni | Automatic for managed data dir | Manual |
| Stop PostgreSQL | Checked and stopped by command | Must be pre-stopped |
| Start PostgreSQL | Default yes; disable with `--no-restart` | Manual |
| Resume Patroni | Manual after validation | Not handled |
| Use case | Production restore orchestration | Low-level restore or scripting |


## Quick Start

```bash
# Most common: recover to latest data
pig pitr -d

# Recover to specific point in time
pig pitr -t "2025-01-01 12:00:00+08"

# Recover to backup consistency point (fastest)
pig pitr -I

# View execution plan
pig pitr -d --plan

# Skip confirmation (for automation)
pig pitr -d -y

# Recover from specific backup set
pig pitr -d -b 20251225-120000F

# Restore managed data dir but leave PostgreSQL stopped
pig pitr -d --no-restart

# Side-restore to a custom data dir without touching Patroni or /pg/data
pig pitr -d -D /tmp/pg-restore --skip-patroni --no-restart

# Pass extra pgBackRest restore args after --
pig pitr -d -- --delta
```


## Parameters

### Recovery Target (choose one)

| Param | Short | Description |
|:------|:------|:------------|
| `--default` | `-d` | Recover to end of WAL stream (latest data) |
| `--immediate` | `-I` | Recover to backup consistency point |
| `--time` | `-t` | Recover to specific timestamp |
| `--name` | `-n` | Recover to named restore point |
| `--lsn` | `-l` | Recover to specific LSN |
| `--xid` | `-x` | Recover to specific transaction ID |

### Backup and Target Options

| Param | Short | Description |
|:------|:------|:------------|
| `--set` | `-b` | Select backup set to start recovery from |
| `--target-action` | | Action when target is reached: pause/promote/shutdown |
| `--target-timeline` | `-T` | Recovery timeline: latest/current/N/0xN |
| `--exclusive` | `-X` | Exclusive mode: stop before target |
| `--promote` | `-P` | Auto-promote after reaching a manual recovery target |

### Flow Control

| Param | Short | Description |
|:------|:------|:------------|
| `--skip-patroni` | `-S` | Skip Patroni stop operation; for standalone PostgreSQL or custom `-D` side restore |
| `--no-restart` | `-N` | Do not start PostgreSQL after restore |
| `--plan` | | Show execution plan only |
| `--yes` | `-y` | Skip destructive confirmation |
| `--timeout` | | PostgreSQL startup/recovery wait timeout, default 120 seconds |
| `--force-stop` | | Allow immediate shutdown and kill fallback if fast stop fails |

### Configuration

| Param | Short | Description |
|:------|:------|:------------|
| `--stanza` | `-s` | pgBackRest stanza name (auto-detected) |
| `--config` | `-c` | pgBackRest config file path |
| `--repo` | `-r` | Repository number (multi-repo scenario) |
| `--dbsu` | `-U` | Database superuser (default: `postgres`) |
| `--data` | `-D` | Target data directory |


## Time Format

The `--time` parameter supports multiple formats with automatic timezone completion:

| Format | Example | Description |
|:-------|:--------|:------------|
| Full format | `2025-01-01 12:00:00+08` | Complete timestamp with timezone |
| Date only | `2025-01-01` | Auto-complete to 00:00:00 (current timezone) |
| Time only | `12:00:00` | Auto-complete to today (current timezone) |


## Execution Flow

### Phase 1: Pre-check

- Validate recovery target parameters; missing target shows help and returns an error
- Resolve pgBackRest stanza, repo, and target data directory
- Determine whether `-D` is a custom side restore
- Check target directory initialization and DBSU ownership
- Detect Patroni service and PostgreSQL status

### Phase 2: Handle Patroni

For managed data-dir restore, if Patroni is running, the command stops Patroni so the target PGDATA stays offline during restore. Patroni remains stopped after restore.

If Patroni is running and the target is the managed data directory, `--skip-patroni` is rejected because Patroni may restart PostgreSQL during restore. Custom `-D` side restores do not touch `/pg/data` and can use `--skip-patroni --no-restart`.

### Phase 3: Ensure PostgreSQL Stopped

The command tries fast stop first. If PostgreSQL cannot be stopped, it does not use more aggressive methods by default. Use `--force-stop` to allow immediate shutdown and kill fallback.

### Phase 4: Execute Recovery

Call pgBackRest for actual data recovery, mapping target, backup set, timeline, and target action to `pgbackrest restore`. Raw pgBackRest restore arguments go after `--`:

```bash
pig pitr -d -- --delta
```

### Phase 5: Start or Stay Stopped

Unless `--no-restart` is specified, start PostgreSQL after restore and wait for recovery. Use `--no-restart` for:

- custom `-D` side restore, because restored config keeps the original port
- `--target-action=shutdown`, because PostgreSQL exits after reaching the target
- cases where an operator wants to inspect the restored directory before startup

Start side restores manually with a free port, for example:

```bash
pg_ctl -D /tmp/pg-restore -o "-p 15432" start
```


## Usage Examples

### Scenario 1: Accidental Data Deletion Recovery

```bash
# 1. View available backups
pig pb info

# 2. Preview recovery plan
pig pitr -t "2025-01-15 09:30:00+08" --plan

# 3. Recover to time before deletion
pig pitr -t "2025-01-15 09:30:00+08"

# 4. Verify data
pig pg psql
SELECT * FROM important_table;
```

### Scenario 2: Recover to Latest State

```bash
# Recover to latest data after server failure
pig pitr -d
```

### Scenario 3: Fast Recovery to Backup Point

```bash
# Recover to backup consistency point (no WAL replay needed)
pig pitr -I
```

### Scenario 4: Automation Scripts

```bash
# Skip all confirmations, suitable for automation
pig pitr -d -y
```

### Scenario 5: Side Restore

```bash
pig pitr -d -D /tmp/pg-restore --skip-patroni --no-restart

# Start manually on a free port
pg_ctl -D /tmp/pg-restore -o "-p 15432" start
```

### Scenario 6: Recover Without Starting Managed Data Dir

```bash
# Recover then manually inspect before deciding to start
pig pitr -d --no-restart

# Check recovered data directory
ls -la /pg/data/

# Manual start
pig pg start
```


## Execution Plan Example

Running `pig pitr -d --plan` shows a plan like:

```
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

[Plan mode] No changes made.
```


## Post-Recovery Operations

After successful recovery, detailed follow-up instructions are displayed:

```
══════════════════════════════════════════════════════════════════
 PITR Complete
══════════════════════════════════════════════════════════════════

[1] Verify recovered data:
   pig pg psql

[2] If satisfied, promote to primary:
   pig pg promote

[3] To resume Patroni cluster management:
   WARNING: Ensure data is correct before starting Patroni!
   systemctl start patroni

   Or if you want this node to be the leader:
   1. Promote PostgreSQL first: pig pg promote
   2. Then start Patroni: systemctl start patroni

[4] Re-create pgBackRest stanza if needed:
   pig pb create

══════════════════════════════════════════════════════════════════
```


## Safety Mechanisms

### Confirmation Countdown

Unless `--yes` is used, a 5-second countdown is displayed before execution:

```
WARNING: This will overwrite the current database!
Press Ctrl+C to cancel, or wait for countdown...
Starting PITR in 5 seconds...
```

### Progressive Stop Strategy

To keep the default path conservative, PostgreSQL is first stopped with fast shutdown. More aggressive immediate shutdown and kill fallback are allowed only when `--force-stop` is explicitly set.

### Recovery Verification

When PostgreSQL is started after restore, the command waits for startup/recovery and prompts for log checks on failure.


## Design Notes

**Relationship with other commands:**

- `pig pitr` calls pgBackRest restore and handles local Patroni/PostgreSQL stop plus optional PostgreSQL start.
- `pig pitr` is not a cluster recovery controller; it does not perform Patroni failover, rejoin, VIP handling, or application traffic switching.
- Use `pig pb restore` when you need lower-level restore semantics or finer scripting control.
- Use `pig pt switchover` or `pig pt failover` for manual Patroni role changes.

**Error Handling:**

- Detailed error messages at each phase
- Prompts relevant log locations on failure
- Supports manual continuation after interruption

**Permission Execution:**

- If current user is DBSU: execute commands directly
- If current user is root: use `su - postgres -c "..."` to execute
- Other users: use `sudo -inu postgres -- ...` to execute

**Platform Support:**

This command is designed for Linux systems, depends on Pigsty's default directory structure.
