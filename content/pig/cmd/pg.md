---
title: "CMD: pg"
description: Manage local PostgreSQL server with pig pg subcommand
icon: SquareTerminal
weight: 650
---

The `pig pg` command (alias `pig postgres`) manages local PostgreSQL server and databases. It wraps native tools like `pg_ctl`, `psql`, `vacuumdb`, providing a simplified server management experience.

```bash
pig pg - Manage local PostgreSQL server and databases.

Server Control (via pg_ctl):
  pig pg init     [-v ver] [-D datadir]     initialize data directory
  pig pg start    [-D datadir]              start PostgreSQL server
  pig pg stop     [-D datadir] [-m fast]    stop PostgreSQL server
  pig pg restart  [-D datadir] [-m fast]    restart PostgreSQL server
  pig pg reload   [-D datadir]              reload configuration
  pig pg status   [-D datadir]              show server status
  pig pg promote  [-D datadir]              promote standby to primary
  pig pg role     [-D datadir] [-V]         detect instance role (primary/replica)

Service Management (via systemctl):
  pig pg svc start                          start postgres systemd service
  pig pg svc stop                           stop postgres systemd service
  pig pg svc restart                        restart postgres systemd service
  pig pg svc reload                         reload postgres systemd service
  pig pg svc status                         show postgres service status

Connection & Query:
  pig pg psql     [db] [-c cmd]             connect to database via psql
  pig pg ps       [-a] [-u user]            show current connections
  pig pg kill     [-x] [-u user]            terminate connections (dry-run by default)
  pig pg clone    <src> [dst]               clone database with FILE_COPY

Database Maintenance:
  pig pg vacuum   [db] [-a] [-t table]      vacuum tables
  pig pg analyze  [db] [-a] [-t table]      analyze tables
  pig pg freeze   [db] [-a] [-t table]      vacuum freeze tables
  pig pg repack   [db] [-a] [-t table]      repack tables (online rebuild)

Tuning:
  pig pg tune     [-p profile]              generate optimized parameters

Instance Fork:
  pig pg fork init <name> [-s]              create local physical fork
  pig pg fork list                          list managed /pg/data-* forks
  pig pg fork start|stop|rm <name>          manage an existing fork

Utilities:
  pig pg log <list|tail|show|less|grep>     view PostgreSQL logs
```

------

## Command Overview

**Service Control** (pg_ctl wrapper):

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

**Connection & Query**:

| Command | Alias | Description | Notes |
|:--------|:------|:------------|:------|
| `pg psql` | `sql, connect` | Connect to database | Wraps psql |
| `pg ps` | `activity, act` | Show current connections | Queries pg_stat_activity |
| `pg kill` | `k` | Terminate connections | Default dry-run mode |
| `pg clone` | | Clone a single database | `CREATE DATABASE ... TEMPLATE ... FILE_COPY` |

**Database Maintenance**:

| Command | Alias | Description | Notes |
|:--------|:------|:------------|:------|
| `pg vacuum` | `vac, vc` | Vacuum tables | Wraps vacuumdb |
| `pg analyze` | `ana, az` | Analyze tables | Wraps vacuumdb --analyze-only |
| `pg freeze` | `frz` | Freeze vacuum | Wraps vacuumdb --freeze |
| `pg repack` | `rp` | Online table repacking | Requires pg_repack extension |

**Tuning**:

| Command | Alias | Description | Notes |
|:--------|:------|:------------|:------|
| `pg tune` | `tuning` | Generate PostgreSQL tuning parameters | Auto-detects hardware; supports structured output |

**Instance Fork**:

| Command | Alias | Description | Notes |
|:--------|:------|:------------|:------|
| `pg fork` | | Shortcut for `fork init` | Creates managed fork by default; does not start it |
| `pg fork init` | `create` | Create local disposable physical copy | Default `/pg/data-<name>` |
| `pg fork list` | | List managed forks | Scans `/pg/data-*` |
| `pg fork start` | | Start existing fork | Supports managed name or unmanaged `-d` directory |
| `pg fork stop` | | Stop existing fork | Supports shutdown mode |
| `pg fork rm` | `remove, delete` | Remove fork | Running forks require `--stop` |

**Log Tools**:

| Command | Alias | Description | Notes |
|:--------|:------|:------------|:------|
| `pg log` | `l` | Log management | Parent command |
| `pg log list` | `ls` | List log files | |
| `pg log tail` | `t, f` | Real-time log viewing | tail -f |
| `pg log show` | `cat, c` | Output log content | |
| `pg log less` | `vi, v` | View with less | |
| `pg log grep` | `g, search` | Search logs | |

**Service Subcommand** (`pg svc`):

| Command | Alias | Description |
|:--------|:------|:------------|
| `pg svc start` | `boot, up` | Start postgres service |
| `pg svc stop` | `halt, dn, down` | Stop postgres service |
| `pg svc restart` | `reboot, rt` | Restart postgres service |
| `pg svc reload` | `rl, hup` | Reload postgres service |
| `pg svc status` | `st, stat` | Show service status |

------

## Quick Start

```bash
# Service control
pig pg init                       # Initialize data directory
pig pg start                      # Start PostgreSQL
pig pg status                     # Check status
pig pg stop                       # Stop PostgreSQL
pig pg restart                    # Restart PostgreSQL
pig pg reload                     # Reload configuration

# Connection & query
pig pg psql                       # Connect to postgres database
pig pg psql mydb                  # Connect to specific database
pig pg ps                         # View current connections
pig pg kill -x                    # Terminate connections (requires -x to execute)
pig pg clone meta meta_fork       # Clone a single database

# Database maintenance
pig pg vacuum mydb                # Vacuum specific database
pig pg analyze mydb               # Analyze specific database
pig pg repack mydb                # Online repack database

# Tuning
pig pg tune                       # Auto-detect hardware and generate parameters
pig pg tune -p olap               # Use OLAP workload profile

# Instance fork
pig pg fork dev                   # Create /pg/data-dev
pig pg fork init dev --start      # Create and start fork on an auto-selected high port
pig pg fork list                  # List /pg/data-* forks

# Log viewing
pig pg log tail                   # Real-time view latest log
pig pg log grep ERROR             # Search error logs
pig pg log list --log-dir /var/log/pg  # Custom log directory
```

------

## Global Options

These options apply to all `pig pg` subcommands:

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--version` | `-v` | auto-detect | PostgreSQL major version |
| `--data` | `-D` | `/pg/data` | Data directory path |
| `--dbsu` | `-U` | `postgres` | Database superuser (or `$PIG_DBSU` env) |
**Version Detection Logic:**

1. If `-v` specified, use that version
2. Otherwise read from `PG_VERSION` file in data directory
3. If neither available, use default PostgreSQL in PATH

------

## Service Control Commands

### pg init

Initialize PostgreSQL data directory. Wraps `initdb`.

```bash
pig pg init                       # Initialize with defaults
pig pg init -v 17                 # Specify PostgreSQL 17
pig pg init -D /data/pg17         # Specify data directory
pig pg init -k                    # Enable data checksums
pig pg init -f                    # Force init (remove existing data)
pig pg init -- --waldir=/wal      # Pass extra args to initdb
```

**Options:**

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--encoding` | `-E` | UTF8 | Database encoding |
| `--locale` | | C | Locale setting |
| `--data-checksum` | `-k` | false | Enable data checksums |
| `--force` | `-f` | false | Force init, remove existing data (dangerous!) |

**Safety:** Even with `--force`, command refuses to run if PostgreSQL is running.


### pg start

Start PostgreSQL server.

```bash
pig pg start                      # Start with defaults
pig pg start -D /data/pg17        # Specify data directory
pig pg start -l /pg/log/pg.log    # Redirect output to log file
pig pg start -o "-p 5433"         # Pass options to postgres
pig pg start -y                   # Force start (skip running check)
pig pg start -S                   # Use systemctl to start
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--log` | `-l` | Redirect stdout/stderr to log file |
| `--timeout` | `-t` | Wait timeout (seconds) |
| `--no-wait` | `-W` | Don't wait for startup completion |
| `--options` | `-o` | Options to pass to postgres |
| `--yes` | `-y` | Force start (even if already running) |


### pg stop

Stop PostgreSQL server.

```bash
pig pg stop                       # Fast shutdown (default)
pig pg stop -m smart              # Wait for clients to disconnect
pig pg stop -m immediate          # Immediate shutdown
pig pg stop -S                    # Use systemctl to stop
```

**Options:**

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--mode` | `-m` | fast | Shutdown mode: smart/fast/immediate |
| `--timeout` | `-t` | 60 | Wait timeout (seconds) |
| `--no-wait` | `-W` | false | Don't wait for shutdown completion |

**Shutdown Modes:**

| Mode | Description |
|:---|:---|
| `smart` | Wait for all clients to disconnect |
| `fast` | Rollback active transactions, disconnect clients, clean shutdown |
| `immediate` | Terminate all processes immediately, requires recovery on next start |


### pg restart

Restart PostgreSQL server.

```bash
pig pg restart                    # Fast restart
pig pg restart -m immediate       # Immediate restart
pig pg restart -o "-p 5433"       # Restart with new options
pig pg restart -S                 # Use systemctl to restart
```

**Options:** Same as `pg stop`, plus `--options` (`-o`) to pass to postgres.


### pg reload

Reload PostgreSQL configuration. Sends SIGHUP signal to server.

```bash
pig pg reload                     # Reload configuration
pig pg reload -D /data/pg17       # Specify data directory
pig pg reload -S                  # Use systemctl reload
```


### pg status

Show PostgreSQL server status. Displays not only `pg_ctl status` output, but also postgres processes and Pigsty-related service status.

```bash
pig pg status                     # Check service status
pig pg status -D /data/pg17       # Specify data directory
```

**Output includes:**

1. `pg_ctl status` output (running status, PID, etc.)
2. PostgreSQL process list (`ps -u postgres`)
3. Related service status:
   - `postgres`: PostgreSQL systemd service
   - `patroni`: Patroni HA manager
   - `pgbouncer`: Connection pooler
   - `pgbackrest`: Backup service
   - `vip-manager`: VIP manager
   - `haproxy`: Load balancer


### pg promote

Promote replica to primary.

```bash
pig pg promote                    # Promote replica
pig pg promote -D /data/pg17      # Specify data directory
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--timeout` | `-t` | Wait timeout (seconds) |
| `--no-wait` | `-W` | Don't wait for promotion completion |


### pg role

Detect PostgreSQL instance role (primary or replica).

```bash
pig pg role                       # Output: primary, replica, or unknown
pig pg role -V                    # Verbose output, show detection process
pig pg role -D /data/pg17         # Specify data directory
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--verbose` | `-V` | Show detailed detection process |

**Output:**

- `primary`: Current instance is primary
- `replica`: Current instance is replica
- `unknown`: Cannot determine instance role

**Detection Strategy (by priority):**

1. **Process detection**: Check for `walreceiver`, `recovery` processes
2. **SQL query**: Execute `pg_is_in_recovery()` (requires PostgreSQL running)
3. **Data directory check**: Check for `standby.signal`, `recovery.signal`, `recovery.conf` files

------

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
pig pg kill -w 5 -x               # Repeat every 5 seconds
pig pg kill --plan                # Preview kill plan
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
| `--all` | `-a` | Include replication connections |
| `--cancel` | `-c` | Cancel queries instead of terminating |
| `--watch` | `-w` | Repeat every N seconds |
| `--plan` | | Preview execution plan without terminating connections |

**Security:** `--state` and `--query` parameters are validated to accept only simple alphanumeric patterns, preventing SQL injection.

------

### pg clone

Clone a database inside the current PostgreSQL instance. This wraps `CREATE DATABASE ... TEMPLATE ... STRATEGY FILE_COPY` and terminates existing sessions on the source database before cloning.

```bash
pig pg clone meta                       # Clone meta as meta_1/meta_2/...
pig pg clone meta meta_fork             # Clone to a specific database name
pig pg clone meta meta_fork --owner dba # Try to change owner after clone
pig pg clone meta meta_fork -p 5433     # Connect through a specific local port
pig pg clone meta meta_fork --plan      # Preview clone plan
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--port` | `-p` | PostgreSQL port (default `5432` or `$PG_PORT`) |
| `--conn-db` | | Database used to run `CREATE DATABASE`; defaults to `template1` when cloning `postgres` |
| `--owner` | | Try to alter owner after clone |
| `--conn-limit` | | Connection limit for new database (`-1` unlimited, `0` disallow connections) |
| `--plan` | | Show plan only |
| `--yes` | `-y` | Skip confirmation prompt |

**Notes:** PostgreSQL 18+ can use CoW semantics when `file_copy_method=clone` is available; otherwise it falls back to normal file copy. This command clones one database, not a new PostgreSQL instance.


## Database Maintenance Commands

### pg vacuum

Vacuum database tables. Wraps `vacuumdb`.

```bash
pig pg vacuum                     # Vacuum current database
pig pg vacuum mydb                # Vacuum specific database
pig pg vacuum -a                  # Vacuum all databases
pig pg vacuum mydb -t mytable     # Vacuum specific table
pig pg vacuum mydb -n myschema    # Vacuum tables in schema
pig pg vacuum mydb --full         # VACUUM FULL (requires exclusive lock)
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--all` | `-a` | Process all databases |
| `--schema` | `-n` | Specify schema |
| `--table` | `-t` | Specify table |
| `--verbose` | `-V` | Verbose output |
| `--full` | `-F` | VACUUM FULL (requires exclusive lock) |

**Security:** `--schema` and `--table` parameters are validated for proper PostgreSQL identifier format.


### pg analyze

Analyze database tables to update statistics.

```bash
pig pg analyze                    # Analyze current database
pig pg analyze mydb               # Analyze specific database
pig pg analyze -a                 # Analyze all databases
pig pg analyze mydb -t mytable    # Analyze specific table
```

**Options:** Same as `pg vacuum` (without `--full`).


### pg freeze

Freeze vacuum database to prevent transaction ID wraparound.

```bash
pig pg freeze                     # Freeze current database
pig pg freeze mydb                # Freeze specific database
pig pg freeze -a                  # Freeze all databases
```

**Options:** Same as `pg analyze`.


### pg repack

Online table repacking. Requires `pg_repack` extension.

```bash
pig pg repack mydb                # Repack all tables in database
pig pg repack -a                  # Repack all databases
pig pg repack mydb -t mytable     # Repack specific table
pig pg repack mydb -n myschema    # Repack tables in schema
pig pg repack mydb -j 4           # Use 4 parallel jobs
pig pg repack mydb --plan         # Show tables to be repacked
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--all` | `-a` | Process all databases |
| `--schema` | `-n` | Specify schema |
| `--table` | `-t` | Specify table |
| `--verbose` | `-V` | Verbose output |
| `--jobs` | `-j` | Number of parallel jobs (default 1) |
| `--plan` | `-N` | Show tables to be repacked |

------

## Tuning Commands

### pg tune

Generate recommended PostgreSQL parameters from the current PostgreSQL major version, host resources, and workload profile.

```bash
pig pg tune                       # Auto-detect hardware, use oltp profile
pig pg tune -p olap               # Use OLAP workload profile
pig pg tune -p tiny               # Small instance profile
pig pg tune -c 8 -m 32768 -d 500  # Override CPU, memory, and disk
pig pg tune -C 500                # Override max_connections
pig pg tune -R 0.30               # Adjust shared_buffers ratio
pig pg tune -o json               # Structured JSON output
pig pg tune -o yaml               # Structured YAML output
```

**Options:**

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--profile` | `-p` | oltp | Tuning profile: `oltp` / `olap` / `tiny` / `crit` |
| `--cpu` | `-c` | 0 | CPU cores; 0 means auto-detect |
| `--mem` | `-m` | 0 | Memory in MB; 0 means auto-detect |
| `--disk` | `-d` | 0 | Data disk size in GB; 0 means auto-detect |
| `--max-conn` | `-C` | 0 | Override `max_connections`; 0 means profile default |
| `--shmem-ratio` | `-R` | 0.25 | `shared_buffers` memory ratio, range `0.1 ~ 0.4` |

The command prints recommended settings only; it does not modify PostgreSQL configuration files.


## Instance Fork

### pg fork

Create a local disposable PostgreSQL physical copy for temporary analysis, troubleshooting, recovery checks, and development tests. Managed forks default to `/pg/data-<name>` and are not registered with Pigsty, systemd, or Patroni. Using `-d|--dst-data` creates an unmanaged fork that is not listed by `fork list`.

```bash
pig pg fork dev                       # Create /pg/data-dev, do not start
pig pg fork init dev --start          # Create and start, probing ports from 15432
pig pg fork init dev -s -p 15433      # Create and start on a specific port
pig pg fork init dev -D /pg/data2 -P 15431  # Specify source data dir and port
pig pg fork init dev -d /tmp/dev      # Create unmanaged fork
pig pg fork list                      # List managed forks
pig pg fork start dev                 # Start existing managed fork
pig pg fork stop dev                  # Stop existing managed fork
pig pg fork rm dev --stop             # Stop and remove a running fork
pig pg fork init dev --plan           # Show plan only
```

**Create Options:**

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--dst-data` | `-d` | `/pg/data-<name>` | Unmanaged target data directory |
| `--dst-port` | `-p` | auto | Target port; probes from 15432 |
| `--src-data` | | `/pg/data` or `$PG_DATA` | Source data directory; can also use global `pg -D/--data` |
| `--src-port` | `-P` | `5432` or `$PG_PORT` | Source port |
| `--start` | `-s` | false | Start fork after creation |
| `--force` | `-f` | false | Replace an existing stopped target directory and skip confirmation |
| `--list` | | false | List `/pg/data-*` forks |
| `--timeout` | `-t` | 60 | Startup wait timeout in seconds |
| `--yes` | `-y` | false | Skip confirmation prompt |
| `--plan` | | false | Show plan only |

**Management Commands:**

| Command | Common Options | Description |
|:---|:---|:---|
| `pig pg fork list` | `--plan`, `-o json/yaml` | List managed forks |
| `pig pg fork start <name> or -d <dir>` | `-d/--dst-data`, `-p/--dst-port`, `-t/--timeout`, `--plan` | Start existing fork |
| `pig pg fork stop <name> or -d <dir>` | `-d/--dst-data`, `-m/--mode`, `-t/--timeout`, `--plan` | Stop existing fork |
| `pig pg fork rm <name> or -d <dir>` | `-d/--dst-data`, `--stop`, `-m/--mode`, `-t/--timeout`, `-f/--force`, `-y/--yes`, `--plan` | Remove fork; running forks require `--stop` |

**Behavior Notes:**

- When the source is running, the command uses PostgreSQL low-level backup APIs to create a consistent physical copy; when the source is stopped, it can perform a cold copy.
- The command prefers CoW/reflink and warns before falling back to regular copy in interactive mode.
- To avoid deleting source data, target paths cannot be `/`, `/pg`, source PGDATA, or a parent/child of the source; symlinks are resolved before checks.
- The fork copy is cleaned of runtime and replication state, and `fork.json` is written. The new instance starts only with `-s|--start`.
- Managed forks are addressed by name. Unmanaged forks must be managed with `-d|--dst-data`.


## Log Commands

Log commands view PostgreSQL log files. Default log directory is `/pg/log/postgres`, can be changed via `--log-dir`.

**Log Command Global Options:**

| Option | Description |
|:---|:---|
| `--log-dir` | Log directory path (default: `/pg/log/postgres`) |
| `--lines` / `-n` | Number of lines to show (default 50) |
| `--follow` / `-f` | Follow latest log (parent `pg log` command only) |

**Permission Handling:** If current user lacks permission to read log directory, command automatically retries with `sudo`. `-o json` outputs JSONL log records; log snapshots do not support `yaml` or `json-pretty`.


### pg log

Show a snapshot from the latest log; with `-f`, follow the latest log.

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
pig pg log tail postgresql.csv    # View specific log file
pig pg log tail -n 100            # Show last 100 lines then follow
pig pg log tail --log-dir /var/log/postgres  # Use custom directory
```

**Options:**

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--lines` | `-n` | 50 | Number of lines to show |


### pg log show

Output log file content.

```bash
pig pg log show                   # Output latest log
pig pg log cat                    # Alias of show
pig pg log c                      # Alias of show
pig pg log show -n 100            # Output last 100 lines
pig pg log show postgresql.csv    # Output specific log file
```

**Options:**

| Option | Short | Default | Description |
|:---|:---|:---|:---|
| `--lines` | `-n` | 50 | Number of lines to show |


### pg log less

Open log file with less. Defaults to end of file (`+G`).

```bash
pig pg log less                   # Open latest log with less
pig pg log less postgresql.csv    # Open specific log file
```


### pg log grep

Search log files.

```bash
pig pg log grep ERROR             # Search for ERROR lines
pig pg log grep --ignore-case error  # Case insensitive
pig pg log grep -C 3 ERROR        # Show 3 lines context
pig pg log grep ERROR pg.csv      # Search specific log file
```

**Options:**

| Option | Short | Description |
|:---|:---|:---|
| `--ignore-case` | | Case insensitive |
| `--context` | `-C` | Show N lines of context |

------

## pg svc Subcommand

`pg svc` provides systemctl-based PostgreSQL service management:

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

------

## Design Notes

**Relationship with Native Tools:**

`pig pg` is not a simple wrapper of PostgreSQL native tools, but a higher-level abstraction for common operations:

- Service control commands (init/start/stop/restart/reload/promote) call `pg_ctl` or `systemctl`
- `status` command shows process and related service status beyond `pg_ctl status`
- Connection management commands (psql/ps/kill) call `psql`
- Maintenance commands (vacuum/analyze/freeze) call `vacuumdb`
- repack command calls `pg_repack`
- Log commands call system tools like `tail`, `less`, `grep`

For full native tool functionality, call the respective commands directly.

**Security Considerations:**

- `--state`, `--query`, `--schema`, `--table` parameters are validated to prevent SQL injection
- `pg kill` defaults to dry-run mode to prevent accidents
- Log commands auto-retry with sudo when permissions insufficient

**Platform Support:**

This command is designed for Linux systems, some features depend on `systemctl` and `journalctl`.
