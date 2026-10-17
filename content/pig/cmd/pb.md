---
title: "CMD: pb"
description: Manage pgBackRest backup and PITR with pig pb subcommand
icon: fa-solid fa-terminal
weight: 670
---

`pig pgbackrest` (alias `pig pb`) creates backups, restores data, and manages pgBackRest stanzas.
It runs pgBackRest as the PostgreSQL operating-system user, selected by `--dbsu` or `$PIG_DBSU`
(default `postgres`).

Before the first backup, install pgBackRest and configure the stanza and repository in `pgbackrest.conf`.
A *stanza* groups the backup configuration for one PostgreSQL cluster.
In PostgreSQL, enable `archive_mode` and set `archive_command` to call pgBackRest.

Use [`pig pitr`](/pig/cmd/pitr/) when recovery must also stop Patroni and PostgreSQL, then restart PostgreSQL.
Use `pig pb restore` when you manage those steps yourself.

## Command Overview

**Information Query**:

| Command | Alias | Description | Implementation |
|:---|:---|:---|:---|
| `pb info` | `i` | Show backup repository info | `pgbackrest info` |
| `pb list` | `ls` | List backup sets, repos, and stanzas | `pgbackrest info` and configuration |
{.full-width}

**Backup & Restore**:

| Command | Alias | Description | Implementation |
|:---|:---|:---|:---|
| `pb backup` | `b` | Create backup | `pgbackrest backup` |
| `pb restore` | `r` | Restore backup files | `pgbackrest restore` |
| `pb expire` | `e` | Clean up expired backups | `pgbackrest expire` |
{.full-width}

**Stanza Management**:

| Command | Alias | Description | Implementation |
|:---|:---|:---|:---|
| `pb create` | `c` | Create stanza (first-time setup) | `pgbackrest stanza-create` |
| `pb upgrade` | `u` | Upgrade stanza after PG major upgrade | `pgbackrest stanza-upgrade` |
| `pb delete` | `d` | Delete stanza and its backups | `pgbackrest stanza-delete` |
{.full-width}

**Control Commands**:

| Command | Alias | Description | Implementation |
|:---|:---|:---|:---|
| `pb check` | `ck` | Check stanza configuration and WAL archiving | `pgbackrest check` |
| `pb start` | `up` | Enable pgBackRest operations | `pgbackrest start` |
| `pb stop` | `dw` | Disable pgBackRest operations | `pgbackrest stop` |
| `pb log` | `l` | View logs | latest log snapshot / tail |
{.full-width}

## Create the first backup {#quick-start}

After configuring pgBackRest and WAL archiving, create the stanza and check the configuration:

```bash
pig pb create
pig pb check
```

If the check succeeds, create a full backup and inspect it:

```bash
pig pb backup full
pig pb info
```

## Global Options

These options apply to all `pig pb` subcommands:

| Option | Short | Description |
|:---|:---|:---|
| `--stanza` | `-s` | pgBackRest stanza name (auto-detected) |
| `--config` | `-c` | Config file path |
| `--repo` | `-r` | Repository number (multi-repo scenario) |
| `--dbsu` | `-U` | PostgreSQL operating-system user (default: `$PIG_DBSU` or `postgres`) |
{.full-width}

**Stanza Auto-Detection:**

If `-s` is not specified, PIG detects the stanza name from the config file:

1. Read the config file (default `/etc/pgbackrest/pgbackrest.conf`)
2. Find sections that do not start with `[global*]`
3. Use the first stanza found

If the config file contains multiple stanzas, PIG emits a warning and uses the first one. In that case, specify `--stanza` explicitly.

**Multi-Repo Support:**

pgBackRest supports multiple repositories (repo1, repo2, etc.). Use `-r` to choose the target repository:

```bash
pig pb backup -r 1                   # Backup to repo1
pig pb backup -r 2                   # Backup to repo2
pig pb info -r 2                     # View repo2 backup info
```

## Information Commands

### pb info

Show detailed backup repository information, including all backup sets and WAL archive status.

```bash
pig pb info                          # Show all backup info
pig pb info --raw --raw-output json  # Raw JSON output
pig pb info --set 20250101-120000F   # Show details for a specific backup set
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--raw` | `-R` | Raw output mode (passes through pgBackRest output) |
| `--raw-output` | | Raw output format: text, json (only in `--raw` mode) |
| `--set` | | Show details for a specific backup set |
{.full-width}

### pb ls

List resources in the backup repository.

```bash
pig pb ls                            # List all backups (default)
pig pb ls backup                     # List all backups (explicit)
pig pb ls repo                       # List configured repos
pig pb ls stanza                     # List all stanzas
```

**Types:**

| Type | Description | Data Source |
|:---|:---|:---|
| backup | List all backup sets (default) | pgbackrest info |
| repo | List configured repos | Parse pgbackrest.conf |
| stanza | List all stanzas | Parse pgbackrest.conf |
{.full-width}

## Backup Commands

### pb backup

Create a physical backup. PIG checks that the local instance is a primary before running pgBackRest.

```bash
pig pb backup                        # Auto mode
pig pb backup full                   # Full backup
pig pb backup diff                   # Differential backup
pig pb backup incr                   # Incremental backup
pig pb backup --force                # Skip primary role check
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--force` | `-f` | Skip primary role check |
{.full-width}

**Backup Types:**

| Type | Description |
|:---|:---|
| (empty) | Use pgBackRest configuration; by default, incremental, or full when no full backup exists |
| full | Full backup: backup all data |
| diff | Differential backup: changes since the last full backup |
| incr | Incremental backup: changes since the last backup of any type |
{.full-width}

A replica, stopped instance, or unknown role fails this check. Use `--force` only when you intend
to delegate role handling to your pgBackRest configuration. It skips the PIG check; pgBackRest still
applies its own backup requirements.

### pb expire

Remove backups and WAL archives according to the configured retention policy.
Use an exact backup label with `--set`; this also expires backups that depend on that set.
WAL retention still applies during this operation. See the [pgBackRest expire reference](https://pgbackrest.org/command.html#command-expire).

```bash
pig pb expire                        # Clean up per policy
pig pb expire --set 20250101-120000F       # Delete a specific backup set
pig pb expire --set 20250101-120000F -y    # Skip confirmation when deleting backup set
pig pb expire --plan                 # Preview cleanup plan only, do not delete
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--set` | | Delete a specific backup set |
| `--plan` | | Preview cleanup plan only, do not delete backups |
| `--yes` | `-y` | Skip confirmation prompt when used with `--set` |
{.full-width}

`--plan` runs a pgBackRest dry run. With structured output, the PIG plan includes the dry-run
output when it is available. Review that output to see the affected backups.

**Retention Policy:**

Retention policy is configured in `pgbackrest.conf`:

```ini
[global]
repo1-retention-full=2               # Full backups to retain
repo1-retention-diff=4               # Differential backups to retain
repo1-retention-archive=2            # WAL archive retention policy
```

## Restore Commands

### pb restore

Restore backup files into the target data directory. PostgreSQL must be stopped.
Select exactly one recovery target: `-d`, `-I`, `-t`, `--name`, `--lsn`, or `--xid`.
A missing target returns an error; text mode also shows help.

```bash
# Recovery target (mutually exclusive options)
pig pb restore -d                    # Restore to latest (explicit)
pig pb restore -I                    # Restore to backup consistency point
pig pb restore -t "2025-01-01 12:00:00+08"  # Restore to specific time
pig pb restore -t "2025-01-01"       # Restore to date (00:00:00 that day)
pig pb restore -t "12:00:00"         # Restore to time (today)
pig pb restore --name my-savepoint   # Restore to named restore point
pig pb restore --lsn "0/7C82CB8"     # Restore to LSN
pig pb restore --xid 12345           # Restore to transaction ID

# Backup set selection (must combine with a recovery target)
pig pb restore -b 20251225-120000F -d  # Restore from specific backup set to latest

# Other options
pig pb restore -t "2025-01-01 12:00:00+08" -X           # Exclusive mode (stop before target)
pig pb restore -t "2025-01-01 12:00:00+08" --target-action=promote  # Promote after reaching target
pig pb restore -t "2025-01-01 12:00:00+08" -T current   # Recover along current timeline
pig pb restore -d --plan             # Preview restore plan only
pig pb restore -d -y                 # Skip interactive y/yes confirmation
pig pb restore -d -- --delta         # Pass native pgBackRest restore args after --
```

**Recovery Target Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--default` | `-d` | Restore to end of WAL stream (latest data) |
| `--immediate` | `-I` | Restore to backup consistency point |
| `--time` | `-t` | Restore to specific timestamp |
| `--name` | | Restore to named restore point |
| `--lsn` | | Restore to specific LSN |
| `--xid` | | Restore to specific transaction ID |
{.full-width}

**Backup Set and Other Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--set` | `-b` | Restore from a specific backup set (can combine with target) |
| `--data` | `-D` | Target data directory |
| `--exclusive` | `-X` | Exclusive mode: stop before target |
| `--target-action` | | Action after reaching recovery target: pause/promote/shutdown |
| `--target-timeline` | `-T` | Recovery timeline: latest/current/N/0xN |
| `--plan` | | Preview restore plan only, do not execute |
| `--yes` | `-y` | Skip interactive y/yes confirmation |
| Native args after `--` | | Pass through pgBackRest restore args, for example `-- --delta` |
{.full-width}

**Combination Rules:** `--target-action` cannot be used with `--default`, because `--default` already means recovery to the end of the WAL stream. `--exclusive/-X` must be used with an explicit stop target: `--time`, `--lsn`, or `--xid`.

Set the target, data directory, stanza, repository, and config through PIG options.
PIG rejects native arguments that would override them, including `--repoN-*` and `--recovery-option`.
Pass other restore options after `--`, for example `--delta`, `--tablespace-map`, `--link-map`, or `--link-all`.
Use `-- --delta` to restore over an existing data directory with checksum-based file reuse.

**Time Formats:**

Supports multiple time formats and auto-completes timezones (including non-integer-hour zones such as +05:30):

| Format | Example | Description |
|:---|:---|:---|
| Full format | `2025-01-01 12:00:00+08` | Complete timestamp with timezone |
| Date only | `2025-01-01` | Auto-completes to 00:00:00 that day (current timezone) |
| Time only | `12:00:00` | Auto-completes to today (current timezone) |
{.full-width}

**Restore Flow:**

1. Validate parameters and environment
2. Check that PostgreSQL is stopped
3. Show restore plan and wait for interactive `y`/`yes` confirmation
4. Execute pgbackrest restore
5. Provide post-restore guidance

For a standalone instance at `/pg/data`, inspect the plan, stop the instance, then restore:

```bash
pig pb restore -D /pg/data -t "2025-01-01 12:00:00+08" --plan
pig pg stop -D /pg/data
pig pb restore -D /pg/data -t "2025-01-01 12:00:00+08" -- --delta
pig pg start -D /pg/data
pig pg psql -D /pg/data
```

For another path, pass the same `-D` to each command. Verify the recovered data before promotion.
When Patroni is active for the managed directory, `pb restore` refuses the operation;
use [`pig pitr`](/pig/cmd/pitr/).

Text mode asks for `y` or `yes` before restore. In scripts with `-o json` or `-o yaml`,
add `--yes` to execute, or use `--plan` to inspect the operation.

## Stanza Management Commands

### pb create

Initialize a new stanza. Must run before the first backup.

```bash
pig pb create                        # Create stanza
pig pb create --no-online            # Create when PostgreSQL is not running
pig pb create --force                # Force create
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--no-online` | | Create when PostgreSQL is not running |
| `--force` | `-f` | Force create |
{.full-width}

### pb upgrade

Update stanza after a PostgreSQL major version upgrade.

```bash
pig pb upgrade                       # Upgrade stanza
pig pb upgrade --no-online           # Upgrade when PostgreSQL is not running
```

**Options:**

| Option | Description |
|:---|:---|
| `--no-online` | Upgrade when PostgreSQL is not running |
{.full-width}

**Use Case:**

After a PostgreSQL major version upgrade (for example, 16 -> 17), run this command to update stanza metadata.

### pb delete

Delete a stanza and all its backups.

```bash
pig pb delete                        # Delete stanza (interactive y/N confirmation)
pig pb delete --yes                  # Skip interactive y/yes confirmation
pig pb delete --plan                 # Preview delete plan
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--plan` | | Preview delete plan only, do not execute |
| `--yes` | `-y` | Skip interactive y/yes confirmation |
{.full-width}

Deletion permanently removes the selected stanza's backups. In text mode, enter `y` or `yes`
to confirm, or use `--yes`. Structured execution requires `--yes`.
When the configuration has several stanzas, select one with `--stanza`.

## Control Commands

### pb check

Check stanza configuration, repository access, and WAL archiving.

```bash
pig pb check                         # Check configuration and archiving
```

This command checks:

- WAL archive configuration correctness
- Repository accessibility
- Stanza configuration validity

`check` does not read every backup file to verify its contents. Use pgBackRest
[`verify`](https://pgbackrest.org/command.html#command-verify) when you need backup and archive checksum verification.

### pb start

Enable pgBackRest operations.

```bash
pig pb start                         # Enable operations
```

Use this command after `pb stop` to resume normal operations.

### pb stop

Disable pgBackRest operations for maintenance.

```bash
pig pb stop                          # Disable operations
pig pb stop --force                  # Terminate running operations
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--force` | `-f` | Terminate running operations |
{.full-width}

**Use Case:**

Use this command during system maintenance to prevent new backup operations from starting.

## Log Commands

### pb log

View pgBackRest log files. The log directory is read from pgBackRest config `log-path` first, and falls back to `/pg/log/pgbackrest/` when not configured. The parent command shows the latest log snapshot by default; use `tail` or `-f` for live follow. Only `pb log` and `pb log show` support `-o json` JSONL output; log snapshots do not support `yaml` or `json-pretty`, and follow/tail does not support structured output.

```bash
pig pb log                           # Show latest log lines
pig pb log list                      # List log files
pig pb log tail                      # Real-time view latest log
pig pb log tail -n 100               # Show last 100 lines and follow
pig pb log show                      # Show latest log content
pig pb log show -n 50                # Show last 50 lines
pig pb log cat                       # Alias for show
```

**Subcommands:**

| Subcommand | Aliases | Description |
|:---|:---|:---|
| list | ls | List log files |
| show | cat, c | Show latest log content |
| tail | t, f, follow | Real-time follow latest log |
{.full-width}

**Options:**

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--lines` | `-n` | 50 | Number of lines to show |
| `--follow` | `-f` | false | Follow with `pb log`; `pb log tail` always follows |
{.full-width}

**Permission Handling:**

If the current user does not have permission to read the log directory, the command automatically retries with `sudo`.

## Paths and privileges

| Setting | Default |
|:---|:---|
| Configuration file | `/etc/pgbackrest/pgbackrest.conf` |
| Log directory | Configured `log-path`, then `/pg/log/pgbackrest` |
| Restore data directory | `--data`, configured `pg1-path`, `$PGDATA`, then `/pg/data` |
{.full-width}

PIG runs pgBackRest directly when the current user is the selected database operating-system user.
From root, it uses `su`; from another account, it uses `sudo -H -u`.
The user must be able to read the configuration and access the data and repository paths.

These commands use Linux and Pigsty defaults. Set the configuration path and stanza explicitly
for other layouts. For pgBackRest operations without a PIG command, run `pgbackrest` directly.

## Design record

[PIG's Operations CLI Safety Contract](https://pig.pgsty.com/design/ops-cli-safety/) explains the boundary between
the low-level `pb restore` primitive and the orchestrated `pig pitr` workflow.
