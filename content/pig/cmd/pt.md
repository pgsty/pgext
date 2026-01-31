---
title: "CMD: pt"
description: Manage Patroni service and cluster with pig pt subcommand
icon: SquareTerminal
weight: 660
---

The `pig patroni` command (alias `pig pt`) manages Patroni service and PostgreSQL HA clusters. It wraps common `patronictl` and `systemctl` operations for simplified cluster management.

```bash
pig pt - Manage Patroni cluster using patronictl commands.

Cluster Operations (via patronictl):
  pig pt list                      list cluster members
  pig pt restart [member]          restart PostgreSQL (rolling restart)
  pig pt reload                    reload PostgreSQL config
  pig pt reinit <member>           reinitialize a member
  pig pt pause                     pause automatic failover
  pig pt resume                    resume automatic failover
  pig pt switchover                perform planned switchover
  pig pt failover                  perform manual failover
  pig pt config <action>           manage cluster config

Service Management (via systemctl):
  pig pt status                    show comprehensive patroni status
  pig pt start                     start patroni service (shortcut)
  pig pt stop                      stop patroni service (shortcut)
  pig pt svc start                 start patroni service
  pig pt svc stop                  stop patroni service
  pig pt svc restart               restart patroni service
  pig pt svc status                show patroni service status

Logs:
  pig pt log [-f] [-n 100]         view patroni logs
```

------

## Command Overview

**Cluster Operations** (patronictl wrapper):

| Command | Alias | Description | Implementation |
|:--------|:------|:------------|:---------------|
| `pt list` | `ls, l` | List cluster members | `patronictl list -e -t` |
| `pt restart` | `reboot, rt` | Restart PostgreSQL instance | `patronictl restart` |
| `pt reload` | `rl, hup` | Reload PostgreSQL config | `patronictl reload` |
| `pt reinit` | `ri` | Reinitialize member | `patronictl reinit` |
| `pt switchover` | `sw` | Planned switchover | `patronictl switchover` |
| `pt failover` | `fo` | Manual failover | `patronictl failover` |
| `pt pause` | `p` | Pause auto-failover | `patronictl pause` |
| `pt resume` | `r` | Resume auto-failover | `patronictl resume` |
| `pt config` | `cfg, c` | Show or modify cluster config | `patronictl show-config / edit-config` |

**Service Management** (systemctl wrapper):

| Command | Alias | Description | Implementation |
|:--------|:------|:------------|:---------------|
| `pt start` | `boot, up` | Start Patroni service | `systemctl start patroni` |
| `pt stop` | `halt, dn, down` | Stop Patroni service | `systemctl stop patroni` |
| `pt status` | `st, stat` | Show service status | `systemctl status patroni` |
| `pt log` | `l, lg` | View Patroni logs | `journalctl -u patroni` |

**Service Subcommand** (`pt svc`):

| Command | Alias | Description |
|:--------|:------|:------------|
| `pt svc start` | `boot, up` | Start Patroni service |
| `pt svc stop` | `halt, dn, down` | Stop Patroni service |
| `pt svc restart` | `reboot, rt` | Restart Patroni service |
| `pt svc reload` | `rl, hup` | Reload Patroni service |
| `pt svc status` | `st, stat` | Show service status |

------

## Quick Start

```bash
# Check cluster member status
pig pt list                    # List default cluster members
pig pt list pg-meta            # List specific cluster
pig pt list -W                 # Continuous watch mode
pig pt list -w 5               # Refresh every 5 seconds

# View and modify cluster config
pig pt config                  # Show current cluster config
pig pt config ttl=60           # Modify single config item (immediate effect)
pig pt config ttl=60 loop_wait=15  # Modify multiple config items

# Cluster operations
pig pt restart                 # Restart all members' PostgreSQL
pig pt restart pg-test-1       # Restart specific member
pig pt switchover              # Planned switchover
pig pt pause                   # Pause auto-failover
pig pt resume                  # Resume auto-failover

# Manage Patroni service
pig pt status                  # Check service status
pig pt start                   # Start service
pig pt stop                    # Stop service
pig pt log -f                  # Real-time log viewing
```

------

## Global Options

These options apply to all `pig pt` subcommands:

| Option | Short | Description |
|:-------|:------|:------------|
| `--dbsu` | `-U` | Database superuser (default: `$PIG_DBSU` or `postgres`) |

------

## Cluster Commands

### pt list

List Patroni cluster member status. Wraps `patronictl list` with `-e` (extended output) and `-t` (show timestamp) flags by default.

```bash
pig pt list                    # List default cluster members
pig pt list pg-meta            # List specific cluster
pig pt list -W                 # Continuous watch mode
pig pt list -w 5               # Refresh every 5 seconds
pig pt list pg-test -W -w 3    # Watch pg-test cluster, 3s refresh
```

**Options:**

| Option | Short | Description |
|:-------|:------|:------------|
| `--watch` | `-W` | Enable continuous watch mode |
| `--interval` | `-w` | Watch refresh interval (seconds) |


### pt restart

Restart PostgreSQL instance via Patroni. This triggers a rolling restart of PostgreSQL, not the Patroni daemon itself.

```bash
pig pt restart                   # Restart all members (interactive)
pig pt restart pg-test-1         # Restart specific member
pig pt restart -f                # Skip confirmation
pig pt restart --role=replica    # Restart replicas only
pig pt restart --pending         # Restart pending members
```

**Options:**

| Option | Short | Description |
|:-------|:------|:------------|
| `--force` | `-f` | Skip confirmation |
| `--role` | | Filter by role (primary/replica) |
| `--pending` | | Restart only pending members |


### pt reload

Reload PostgreSQL configuration via Patroni. Triggers config reload on all members.

```bash
pig pt reload
```


### pt reinit

Reinitialize cluster member. This re-syncs data from the primary.

```bash
pig pt reinit pg-test-1          # Reinit specific member
pig pt reinit pg-test-1 -f       # Skip confirmation
pig pt reinit pg-test-1 --wait   # Wait for completion
```

**Options:**

| Option | Short | Description |
|:-------|:------|:------------|
| `--force` | `-f` | Skip confirmation |
| `--wait` | `-w` | Wait for reinit completion |

**Warning:** This operation deletes all data on the target member and re-syncs from primary.


### pt switchover

Perform planned primary-replica switchover.

```bash
pig pt switchover                 # Interactive switchover
pig pt switchover -f              # Skip confirmation
pig pt switchover --leader pg-1   # Specify current primary
pig pt switchover --candidate pg-2  # Specify new primary
```

**Options:**

| Option | Short | Description |
|:-------|:------|:------------|
| `--force` | `-f` | Skip confirmation |
| `--leader` | | Specify current primary |
| `--candidate` | | Specify candidate new primary |


### pt failover

Perform manual failover. Used when primary is unavailable.

```bash
pig pt failover                   # Interactive failover
pig pt failover -f                # Skip confirmation
pig pt failover --candidate pg-2  # Specify new primary
```

**Options:**

| Option | Short | Description |
|:-------|:------|:------------|
| `--force` | `-f` | Skip confirmation |
| `--candidate` | | Specify candidate new primary |


### pt pause

Pause Patroni's automatic failover.

```bash
pig pt pause                      # Pause auto-failover
pig pt pause --wait               # Wait for confirmation
```

**Options:**

| Option | Short | Description |
|:-------|:------|:------------|
| `--wait` | `-w` | Wait for operation completion |

**Use case:** Pause auto-failover during maintenance operations (e.g., major version upgrade, storage migration) to prevent accidental triggers.


### pt resume

Resume Patroni's automatic failover.

```bash
pig pt resume                     # Resume auto-failover
pig pt resume --wait              # Wait for confirmation
```

**Options:**

| Option | Short | Description |
|:-------|:------|:------------|
| `--wait` | `-w` | Wait for operation completion |


### pt config

Show or modify cluster configuration. Without parameters shows current config, with `key=value` parameters modifies config.

```bash
pig pt config                           # Show current cluster config
pig pt config show                      # Show config (explicit)
pig pt config edit                      # Interactive config edit
pig pt config set ttl=60                # Set TTL to 60 seconds
pig pt config set ttl=60 loop_wait=15   # Modify multiple config items
pig pt config pg max_connections=200    # Modify PostgreSQL parameter
```

**Subcommands:**

| Subcommand | Description |
|:-----------|:------------|
| `show` (default) | Show current config |
| `edit` | Interactive config edit |
| `set key=value` | Directly set config item |
| `pg key=value` | Set PostgreSQL parameter |

**Common config items:**

| Config | Description | Default |
|:-------|:------------|:--------|
| `ttl` | Leader lock time-to-live (seconds) | 30 |
| `loop_wait` | Main loop sleep time (seconds) | 10 |
| `retry_timeout` | DCS and PostgreSQL operation timeout (seconds) | 10 |
| `maximum_lag_on_failover` | Maximum lag allowed during failover (bytes) | 1048576 |

**Note:** This command modifies dynamic cluster config stored in DCS (e.g., etcd), not local config file `/etc/patroni/patroni.yml`.

------

## Service Commands

### pt start

Start Patroni service.

```bash
pig pt start                     # Start Patroni service
pig pt up                        # Alias
pig pt boot                      # Alias
```

Equivalent to `sudo systemctl start patroni`.


### pt stop

Stop Patroni service.

```bash
pig pt stop                      # Stop Patroni service
pig pt down                      # Alias
pig pt halt                      # Alias
```

Equivalent to `sudo systemctl stop patroni`.

**Note:** Stopping Patroni service will also stop the PostgreSQL instance on this node (depending on Patroni configuration).


### pt status

Show Patroni service comprehensive status, including:
- systemd service status
- Patroni process info
- Cluster member status

```bash
pig pt status
```


### pt log

View Patroni service logs.

```bash
pig pt log                     # Show last 50 log lines
pig pt log -f                  # Real-time log following
pig pt log -n 100              # Show last 100 log lines
pig pt log -f -n 200           # Show last 200 lines and follow
```

**Options:**

| Option | Short | Default | Description |
|:-------|:------|:--------|:------------|
| `--follow` | `-f` | false | Real-time log following |
| `--lines` | `-n` | 50 | Number of log lines to show |

Equivalent to `journalctl -u patroni [-f] [-n N]`.

------

## pt svc Subcommand

`pt svc` provides the same functionality as top-level service commands, for explicitly operating on the Patroni daemon:

```bash
pig pt svc start                 # Start Patroni service
pig pt svc stop                  # Stop Patroni service
pig pt svc restart               # Restart Patroni service
pig pt svc reload                # Reload Patroni service
pig pt svc status                # Show service status
```

**Alias Reference:**

| Command | Alias |
|:--------|:------|
| `pt svc start` | `boot, up` |
| `pt svc stop` | `halt, dn, down` |
| `pt svc restart` | `reboot, rt` |
| `pt svc reload` | `rl, hup` |
| `pt svc status` | `st, stat` |

------

## Design Notes

**Relationship with patronictl:**

`pig pt` wraps common `patronictl` operations:
- Cluster queries: `list`, `config show`
- Cluster management: `restart`, `reload`, `reinit`, `switchover`, `failover`, `pause`, `resume`
- Config modification: `config set`, `config edit`
- Service commands (start/stop/restart/reload/status) call `systemctl`
- `log` command calls `journalctl`

**Default Config Paths:**

| Config | Default |
|:-------|:--------|
| Patroni config file | `/etc/patroni/patroni.yml` |
| Service name | `patroni` |

**Permission Handling:**

- If current user is DBSU: execute commands directly
- If current user is root: use `su - postgres -c "..."` to execute
- Other users: use `sudo -inu postgres -- ...` to execute

**Platform Support:**

This command is designed for Linux systems, depends on `systemctl` and `journalctl`.
