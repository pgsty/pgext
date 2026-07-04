---
title: "CMD: pt"
description: Manage Patroni service and cluster with pig pt subcommand
icon: SquareTerminal
weight: 660
---

The `pig patroni` command (alias `pig pt`) manages Patroni services and PostgreSQL HA clusters. It wraps common `patronictl` and `systemctl` operations for a simplified cluster management experience.

```bash
pig pt - Manage Patroni cluster using patronictl commands.

Cluster Operations (via patronictl):
  pig pt list [cluster]            list cluster members
  pig pt restart [member]          restart PostgreSQL (rolling restart)
  pig pt reload                    reload PostgreSQL config
  pig pt reinit <member>           reinitialize a member
  pig pt pause                     pause automatic failover
  pig pt resume                    resume automatic failover
  pig pt switchover                perform planned switchover
  pig pt failover [candidate]      perform manual failover
  pig pt config <action>           manage cluster config (edit|show|set|pg)

Service Management (via systemctl):
  pig pt status                    show comprehensive patroni status
  pig pt svc start (pig pt start)  start patroni service
  pig pt svc stop  (pig pt stop)   stop patroni service
  pig pt svc restart               restart patroni service
  pig pt svc reload                reload patroni service
  pig pt svc status                show patroni service status

Logs:
  pig pt log [-f] [-n 50]          view patroni logs
  pig pt log tail [-n 50]          follow patroni logs
  pig pt log show [-n 50]          show patroni log snapshot
  pig pt log grep <pattern>        search patroni logs
```

`pt start` / `pt stop` are hidden shortcuts for `pt svc start` / `pt svc stop`; `pt svc` is the explicit entry point for Patroni daemon management.

Short option contract: when there is no conflict in a command's scope, `--wait` must provide the `-w` short option. This currently applies to `pt reinit`, `pt pause`, and `pt resume`; `pt list -w` remains reserved for that command's refresh interval.

Switchover preflight contract: before `pt switchover` and `pt failover` execute or ask for confirmation, they read the current topology through the same structured interface used by `pig pt list` / `pig pt config show`, then verify cluster name, current leader, candidate replicas, and Patroni pause state. If the cluster is paused, pig refuses the switch and prompts you to run `pig pt resume` first.


## Command Overview

**Cluster Operations** (`patronictl` wrapper):

These commands manage PostgreSQL clusters through Patroni.

| Command | Alias | Description | Implementation |
|:---|:---|:---|:---|
| `pt list` | `ls` | List cluster members | `patronictl list -e -t` |
| `pt restart` | `rs` | Restart PostgreSQL instance | `patronictl restart` |
| `pt reload` | `rl` | Reload PostgreSQL config | `patronictl reload` |
| `pt reinit` | `ri` | Reinitialize member | `patronictl reinit` |
| `pt switchover` | `so` | Planned switchover | `patronictl switchover` |
| `pt failover` | `fo` | Manual failover | `patronictl failover` |
| `pt pause` | `p` | Pause automatic failover | `patronictl pause` |
| `pt resume` | `r` | Resume automatic failover | `patronictl resume` |
| `pt config` | `c` | View or modify cluster config | `patronictl show-config / edit-config` |
{.full-width}

**Service Subcommands** (`pt service`, `systemctl` wrapper):

These commands manage the Patroni service itself through `systemctl`.

| Command | Alias | Description |
|:---|:---|:---|
| `pt service start` | `pt svc up` | Start Patroni service |
| `pt service stop` | `pt svc dn` | Stop Patroni service |
| `pt service restart` | `pt svc rs` | Restart Patroni service |
| `pt service reload` | `pt svc rl` | Reload Patroni service |
| `pt service status` | `pt svc st` | Show service status |
{.full-width}

**Service Management** (`systemctl` wrapper):

These top-level commands view or manage the Patroni service directly. `pt start` / `pt stop` are hidden shortcuts.

| Command | Alias | Description | Implementation |
|:---|:---|:---|:---|
| `pt start` | `up` | Start Patroni service | `systemctl start patroni` |
| `pt stop` | `dn` | Stop Patroni service | `systemctl stop patroni` |
| `pt status` | `st` | Show comprehensive status | `systemctl status` + `ps` + `patronictl list` |
| `pt log` | `l` | View Patroni logs | Read log files in the Patroni log directory |
{.full-width}


## Quick Start

```bash
# Check cluster member status
pig pt list                    # List default cluster members
pig pt list pg-meta            # List specific cluster members
pig pt list -W                 # Continuous watch mode
pig pt list -w 5               # Refresh every 5 seconds

# View and modify cluster config
pig pt config                  # Show current cluster config (default show)
pig pt config set ttl=60       # Modify one config item (takes effect immediately)
pig pt config set ttl=60 loop_wait=15  # Modify multiple config items
pig pt config pg max_connections=200   # Modify PostgreSQL parameter

# Cluster operations
pig pt restart                 # Restart PostgreSQL on all members (requires confirmation)
pig pt restart pg-test-1       # Restart specific member
pig pt restart --pending       # Apply pending restarts directly
pig pt restart -y              # Cluster-level restart, skip confirmation
pig pt switchover              # Planned switchover
pig pt pause                   # Pause automatic failover
pig pt resume                  # Resume automatic failover

# Manage Patroni service
pig pt status                  # View service status
pig pt start                   # Hidden shortcut: equivalent to pig pt svc start
pig pt stop                    # Hidden shortcut: equivalent to pig pt svc stop
pig pt svc start               # Start service
pig pt svc stop                # Stop service
pig pt log -f                  # Follow logs in real time
pig pt log grep ERROR          # Search logs
```


## Global Options

These options apply to all `pig pt` subcommands:

| Option | Short | Description |
|:---|:---|:---|
| `--dbsu` | `-U` | Database superuser (default: `$PIG_DBSU` or `postgres`) |
{.full-width}


## Cluster Commands

### pt list

List Patroni cluster member status. This command wraps `patronictl list` and adds `-e` (extended output) and `-t` (show timestamp) by default.

```bash
pig pt list                    # List default cluster members
pig pt list pg-meta            # List specific cluster
pig pt list -W                 # Continuous watch mode
pig pt list -w 5               # Refresh every 5 seconds
pig pt list -w 0.5             # Refresh every 0.5 seconds
pig pt list pg-test -W -w 3    # Watch pg-test cluster, refresh every 3 seconds
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--watch` | `-W` | Enable continuous watch mode |
| `--interval` | `-w` | Watch refresh interval in seconds; decimals like 0.5 are supported |
{.full-width}

Watch mode uses live passthrough output from `patronictl` and cannot be combined with structured `-o json` / `-o yaml` output; structured output returns `CodePtWatchModeUnsupported`.


### pt restart

Restart PostgreSQL instances through Patroni. This triggers a PostgreSQL rolling restart, not a restart of the Patroni daemon itself.

```bash
pig pt restart                   # Restart all members (interactive)
pig pt restart pg-test-1         # Restart specific member
pig pt restart -y                # Cluster-level restart, skip confirmation
pig pt restart --role=replica    # Restart replicas only
pig pt restart --pending         # Restart pending members
pig pt restart --plan            # Preview execution plan
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--yes` | `-y` | Skip confirmation |
| `--role` | `-r` | Filter by role (leader/replica/any) |
| `--pending` | `-p` | Restart only pending members |
| `--plan` | | Show execution plan only |
{.full-width}

`pt restart` uses conditional confirmation: restarting a single member or using `--pending` executes directly; a cluster-level rolling restart without a member requires confirmation, and structured output mode requires explicit `--yes`.
The underlying `patronictl restart` is always invoked by pig with `--force`, so `patronictl` does not trigger its own interactive prompt.


### pt reload

Reload PostgreSQL configuration through Patroni. This triggers a config reload on all members.

```bash
pig pt reload
```


### pt reinit

Reinitialize a cluster member. This re-syncs data from the primary.

```bash
pig pt reinit pg-test-1          # Reinitialize specific member
pig pt reinit pg-test-1 -y       # Skip confirmation
pig pt reinit pg-test-1 -w       # Wait for completion
pig pt reinit pg-test-1 --plan   # Preview execution plan
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--yes` | `-y` | Skip confirmation |
| `--wait` | `-w` | Wait for reinit completion |
| `--plan` | | Show execution plan only |
{.full-width}

**Warning:** This operation deletes all data on the target member and re-syncs it. Text mode requires confirmation; JSON/YAML execution mode requires explicit `--yes`.


### pt switchover

Perform a planned primary-replica switchover through Patroni. (Alias: `so`)

```bash
pig pt switchover                 # Interactive switchover
pig pt switchover -y              # Skip confirmation
pig pt switchover -l pg-1 -c pg-2 # Specify current leader and new leader
pig pt switchover -s "2026-07-01T12:00:00"  # Scheduled switchover
pig pt switchover --plan          # Preview execution plan
pig pt so -c pg-test-1 -y         # Switch directly to pg-test-1 without confirmation
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--yes` | `-y` | Skip confirmation |
| `--leader` | `-l` | Specify current leader |
| `--candidate` | `-c` | Specify candidate new leader |
| `--scheduled` | `-s` | Scheduled switchover time |
| `--plan` | | Show execution plan only |
{.full-width}

Compared with the raw `patronictl` CLI, pig parses and fills the cluster scope from `/etc/patroni/patroni.yml`, so users do not need to enter the cluster name manually.
Before execution or confirmation, pig reads the current topology: cluster name, current leader, candidate replicas, and pause state. If the cluster is already paused, the command refuses to run and asks you to run `pig pt resume` first.

If `--candidate` is not specified, pig does not choose an instance itself; candidate selection is delegated to `patronictl` / Patroni. The confirmation prompt explains that leadership will be transferred to the best eligible replica selected by Patroni and lists currently observed candidate members. Use `--candidate/-c` when you need to choose the new leader explicitly.


### pt failover

Perform manual failover. Use this when the primary is unavailable. (Alias: `fo`)

Unlike `switchover`, `failover` does not require the current leader to be available, but you **must** specify a candidate new leader. The candidate can be specified with `--candidate/-c`, or as the only positional argument: `pig pt failover <member>`.

```bash
pig pt failover --candidate pg-2          # Interactive failover
pig pt failover pg-2                      # Positional form, equivalent to -c pg-2
pig pt failover -c pg-2 -y                # Skip confirmation
pig pt failover -c pg-2 --plan            # Preview execution plan
pig pt fo pg-test-2 -y                    # Alias + confirmation skip
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--yes` | `-y` | Skip confirmation |
| `--candidate` | `-c` | Specify candidate new leader; positional argument is also supported |
| `--plan` | | Show execution plan only |
{.full-width}

Before execution or confirmation, pig reads the current topology and checks pause state. If the cluster is already paused, the command refuses to run and asks you to run `pig pt resume` first. The confirmation prompt includes cluster name, current leader, specified candidate new leader, currently observed candidate members, and keeps the warning that failover may lose data.


### pt pause

Pause Patroni automatic failover and enter maintenance mode to prevent failover during maintenance.
If the cluster is already in maintenance mode, the command returns an error.

```bash
pig pt pause                      # Pause automatic failover
pig pt pause -w                   # Wait for confirmation
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--wait` | `-w` | Wait for operation completion |
{.full-width}

**Use Case:** Pause automatic failover during maintenance operations such as major version upgrades or storage migrations to prevent accidental triggers.


### pt resume

Resume Patroni automatic failover and exit maintenance mode.
If the cluster is not in maintenance mode, the command returns an error.

```bash
pig pt resume                     # Resume automatic failover
pig pt resume -w                  # Wait for confirmation
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--wait` | `-w` | Wait for operation completion |
{.full-width}


### pt config

Show or modify cluster configuration. `show` displays current config, `set` modifies Patroni dynamic config, and `pg` modifies PostgreSQL parameters.

```bash
pig pt config show                      # Show current cluster config
pig pt config edit                      # Interactive config edit
pig pt config set ttl=60                # Set TTL to 60 seconds
pig pt config set ttl=60 loop_wait=15   # Modify multiple config items
pig pt config pg max_connections=200    # Modify PostgreSQL parameter
pig pt config set ttl=60 --plan         # Preview config change
```

**Subcommands:**

| Subcommand | Description |
|:---|:---|
| `show` | Show current config |
| `edit` | Interactive config edit |
| `set key=value` | Directly set config item |
| `pg key=value` | Set PostgreSQL parameter |
{.full-width}

**Options:**

| Option | Description |
|:---|:---|
| `--plan` | For `set` / `pg`, show execution plan only |
{.full-width}

Use `pig pt config set k=v` to modify Patroni configuration parameters. Both `set` and `pg` require arguments in `key=value` form. Non-`key=value` tokens are rejected instead of being silently ignored. `edit` is interactive and does not support structured output.

| Config | Description | Default |
|:---|:---|:---|
| `ttl` | Leader lock time-to-live (seconds) | 30 |
| `loop_wait` | Main loop sleep interval (seconds) | 10 |
| `retry_timeout` | DCS and PostgreSQL operation timeout (seconds) | 10 |
| `maximum_lag_on_failover` | Maximum lag allowed during failover (bytes) | 1048576 |
{.full-width}

Use `pig pt config pg k=v` to modify PostgreSQL parameters. The command recognizes known postmaster-context parameters. If you modify a PostgreSQL parameter that requires restart, the plan and structured result include follow-up actions such as `pig pt list` and `pig pt restart --pending`.

**Note:** This command modifies dynamic cluster config stored in DCS (such as etcd), not the local config file `/etc/patroni/patroni.yml`.

For example:

```bash
pig pt config set ttl=60 --plan
pig pt config pg shared_buffers=4GB --plan
```


## Service Commands

### pt start

Start Patroni service.

```bash
pig pt start                     # Start Patroni service
pig pt up                        # Alias
```

Equivalent to `sudo systemctl start patroni`.


### pt stop

Stop Patroni service.

```bash
pig pt stop                      # Stop Patroni service
pig pt dn                        # Alias
```

Equivalent to `sudo systemctl stop patroni`.

**Note:** Stopping the Patroni service can also stop the PostgreSQL instance on that node, depending on Patroni configuration.


### pt status

Show comprehensive Patroni service status, including:

- systemd service status
- Patroni process info
- Cluster member status

```bash
pig pt status
```


### pt log

View Patroni service logs. The log directory is read from `log.dir` in `/etc/patroni/patroni.yml` by default and falls back to `/pg/log/patroni`; you can override it with `--log-dir`. Only `pt log` and `pt log show` support `-o json` JSONL output. Log snapshots do not support `yaml` or `json-pretty`, and follow/tail/grep do not support structured output.

```bash
pig pt log                     # Show the last 50 log lines
pig pt log -f                  # Follow log output in real time
pig pt log show                # Show recent logs
pig pt log tail                # Follow logs
pig pt log grep ERROR          # Search logs
pig pt log -n 100              # Show the last 100 log lines
pig pt log -f -n 200           # Show the last 200 lines and keep following
```

**Subcommands:**

| Subcommand | Aliases | Description |
|:---|:---|:---|
| `show` | `cat, c, s` | Output recent Patroni logs |
| `tail` | `t, f, follow` | Keep following Patroni logs |
| `grep` | `g, search` | Search Patroni logs |
{.full-width}

**Options:**

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--follow` | `-f` | false | Follow log output in real time |
| `--lines` | `-n` | 50 | Number of log lines to show |
| `--log-dir` | | auto-detect | Log directory |
{.full-width}


## pt svc Subcommand

`pt svc` (also `pt service`) provides the same functionality as the top-level service commands, for explicitly operating on the Patroni daemon:

```bash
pig pt svc start                 # Start Patroni service
pig pt svc stop                  # Stop Patroni service
pig pt svc restart               # Restart Patroni service
pig pt svc reload                # Reload Patroni service
pig pt svc status                # Show service status
```

**Alias Reference:**

| Command | Alias |
|:---|:---|
| `pt svc start` | `up` |
| `pt svc stop` | `dn` |
| `pt svc restart` | `rs` |
| `pt svc reload` | `rl` |
| `pt svc status` | `st` |
{.full-width}

`start` and `stop` have dedicated `pt start` and `pt stop` shortcuts, making direct Patroni service management convenient.
Note that `pt restart` is not a shortcut for `pt svc restart`; it is the command for restarting PostgreSQL through Patroni, so the two have different semantics.


## Design Notes

**Relationship with patronictl:**

`pig pt` wraps common `patronictl` operations:

- Cluster queries: `list`, `config show`
- Cluster management: `restart`, `reload`, `reinit`, `switchover`, `failover`, `pause`, `resume`
- Config modification: `config set`, `config pg`, `config edit`
- Service management commands (start/stop/restart/reload/status) call `systemctl`
- The `log` command reads log files from the Patroni log directory

**Default Config Paths:**

| Config | Default |
|:---|:---|
| Patroni config file | `/etc/patroni/patroni.yml` |
| Log directory | config `log.dir`, fallback `/pg/log/patroni` |
| Service name | `patroni` |
{.full-width}

**Permission Handling:**

- If the current user is DBSU: execute commands directly
- If the current user is root: use `su - postgres -c "..."` to execute
- Other users: use `sudo -inu postgres -- ...` to execute

**Platform Support:**

This command is designed for Linux systems. Service management depends on `systemctl`, and log functionality depends on readable Patroni log files.
