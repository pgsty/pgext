---
title: "CMD Docs"
description: Overview of the pig CLI command reference
icon: fa-solid fa-terminal
weight: 600
---

The `pig` CLI provides a comprehensive toolkit for managing PostgreSQL installations, extensions, repositories, and extension builds from source. Use `pig help <command>` to view command documentation.

- [**pig repo**](https://pig.pgsty.com/repo/): manage software repositories
- [**pig ext**](https://pig.pgsty.com/ext/): manage PostgreSQL extensions
- [**pig build**](https://pig.pgsty.com/build/): build extensions from source
- **pig install**: install packages with the native package manager and translate PostgreSQL aliases
- [**pig sty**](https://pig.pgsty.com/sty/): manage Pigsty installation and Grafana dashboards
- [**pig inventory**](https://pig.pgsty.com/inventory/): inspect, edit, validate, and exchange the Pigsty inventory (new in v1.6.0)
- **pig do**: run Pigsty administration playbook tasks
- **pig pe**: access pg_exporter metrics and configuration
- [**pig pg**](https://pig.pgsty.com/pg/): manage local PostgreSQL servers
- [**pig pt**](https://pig.pgsty.com/pt/): run patronictl transparently with service and config helpers
- [**pig pb**](https://pig.pgsty.com/pb/): manage pgBackRest backup and restore
- [**pig pitr**](https://pig.pgsty.com/pitr/): run the full PITR workflow
- **pig context**: output an environment context snapshot for humans and agents
- **pig status / update / version**: inspect environment status, upgrade pig, and print version information


## Overview

```bash
pig - The Linux package manager for PostgreSQL

Usage:
  pig [command]

Examples:

  pig repo add -ru            # overwrite existing repo and update cache
  pig install pg18            # install PostgreSQL 18 PGDG packages
  pig install pg_duckdb       # install a specific PostgreSQL extension
  pig install pgactive -v 18  # install extension for a specific PG version

  visit https://pigsty.io/ext/ for details

PostgreSQL Extension Manager
  build       Build Postgres extensions
  ext         Manage PostgreSQL extensions (pgext)
  repo        Manage Linux software repo (apt/dnf)

Pigsty Management Commands
  context     Show environment context snapshot
  do          Run admin tasks
  inventory   Inspect, edit, validate, check, and exchange Pigsty Inventory
  patroni     Run patronictl with Pigsty service and config helpers (alias: pt)
  pg_exporter Manage pg_exporter and metrics (alias: pe)
  pgbackrest  Manage pgBackRest backup and restore (alias: pb)
  pitr        Point-in-time recovery using pgBackRest
  postgres    Manage local PostgreSQL server and databases (alias: pg)
  sty         Manage Pigsty installation and controller services

Additional Commands:
  completion  Generate shell completion scripts
  help        Help about any command
  install     Install packages using native package manager
  status      Show environment status
  update      Upgrade pig itself
  version     Show pig version info

Flags:
      --debug              enable debug mode
  -h, --help               help for pig
  -H, --home string        Pigsty home path
  -i, --inventory string   config inventory path
      --log-level string   log level: debug, info, warn, error, fatal, panic (default "info")
      --log-path string    log file path, terminal by default
  -o, --output string      output format: text, yaml, json, json-pretty (default "text")

Use "pig [command] --help" for more information about a command.
```


## pig repo

Manage APT/YUM repositories for PostgreSQL packages. See [`pig repo`](https://pig.pgsty.com/repo/) for details.

```bash
pig repo list                    # list available repositories
pig repo info   pgdg             # show repository details
pig repo status                  # check current repository status
pig repo add    pgdg pigsty -u   # add repositories
pig repo rm     old-repo         # remove repository
pig repo update                  # update package cache
pig repo create /www/pigsty      # create local repository
pig repo cache                   # create offline package
pig repo boot                    # bootstrap from offline package
```


## pig ext

Manage PostgreSQL extensions and kernel packages. See [`pig ext`](https://pig.pgsty.com/ext/) for details.

```bash
pig ext list    duck             # search extensions
pig ext info    pg_duckdb        # extension details
pig ext status                   # show installed extensions
pig ext add     pg_duckdb -y     # install extension
pig ext rm      old_extension    # remove extension
pig ext update                   # update extension
pig ext scan                     # scan installed extensions
pig ext import  pg_duckdb        # download for offline use
pig ext link    17               # link PG version into PATH
pig ext reload                   # refresh extension catalog
```


## pig build

Build PostgreSQL extensions from source. See [`pig build`](https://pig.pgsty.com/build/) for details.

```bash
# Environment setup
pig build spec                   # initialize build specs
pig build repo                   # configure repositories
pig build repo --beta            # configure repositories and add PostgreSQL 19 beta repo
pig build tool                   # install build tools
pig build tool --beta            # install build tools plus PG19 beta build packages
pig build rust -y                # force reinstall Rust (default does not reinstall)
pig build rust -m                # use China mirror mode and write Cargo mirror config
pig build pgrx                   # install PGRX framework
pig build pgrx -b                # include PostgreSQL 19 beta pg_config during auto-detection

# Build extensions
pig build pkg citus              # complete build pipeline = get + dep + ext
pig build get citus              # download source
pig build dep citus              # install dependencies
pig build ext citus              # build package
```


## pig install

Install packages through the system's native package manager and translate PostgreSQL kernel, extension, and common aliases into package names. Use `-n/--no-translation` when you need to pass raw system package names directly.

```bash
pig install pg_duckdb            # install extension and translate package name
pig install pg18                 # install PostgreSQL 18 kernel package group
pig install nginx htop vim       # install multiple system packages
pig install unknown-package -n   # disable translation and use raw package name
pig install pg18 --plan          # preview installation plan
pig install pg_vector -y         # auto-confirm installation
```


## pig sty

Install the Pigsty distribution. See [`pig sty`](https://pig.pgsty.com/sty/) for details.

```bash
pig sty init                     # install Pigsty to ~/pigsty
pig sty boot                     # install Ansible prerequisites
pig sty conf                     # generate configuration
pig sty deploy                   # run deployment playbook
pig sty list                     # list available Pigsty releases
pig sty get 4.5.0                # download a Pigsty release
pig sty grafana list             # manage Grafana dashboards (info/list/boot/load/init/dump/clean/lang/style)
```


## pig inventory

Inspect, edit, validate, check, and exchange the Pigsty inventory (`pigsty.yml`) with the
CMDB. Root-level command group with alias `inv`; see [`pig inventory`](https://pig.pgsty.com/inventory/)
for details. (New in v1.6.0)

```bash
pig inventory status             # inspect the active inventory source
pig inventory list               # list inventory topology and value kinds
pig inventory show               # show verbatim inventory YAML (may contain secrets)
pig inventory edit               # edit the inventory (or one fragment) in $EDITOR
pig inventory validate           # validate one complete static Pigsty inventory
pig inventory check              # check inventory, controller, and target readiness
pig inventory diff other.yml     # compare declarations without emitting values
pig inventory cmdb check         # exchange with the PostgreSQL CMDB (experimental)
```


## pig do

Run Pigsty administration tasks through the corresponding Ansible playbooks.

```bash
pig do pgsql-add  <cls> [ip...]       # add cluster or instances
pig do pgsql-rm   <cls> [ip...]       # remove cluster or instances
pig do pgsql-db   <cls> <dbname>      # create or update database
pig do pgsql-user <cls> <username>    # create or update user
pig do pgsql-ext  <cls> [ext...]      # install extensions
pig do node-pkg   <sel> [pkg...]      # install node packages
```


## pig pe

Access PostgreSQL monitoring metrics exposed by pg_exporter. The default endpoint is `127.0.0.1:9630`.

```bash
pig pe list                    # list available metric types
pig pe get                     # get all pg_ prefixed metrics
pig pe stat                    # show exporter statistics
pig pe reload                  # reload pg_exporter configuration
pig pe --host 127.0.0.1 -p 9630 get
```


## pig context

Output an environment context snapshot covering host, PostgreSQL, Patroni, pgBackRest, and installed extensions. This command is useful for troubleshooting and for automation scripts that need a quick view of the current node.

```bash
pig context                      # text output
pig context -o json              # JSON output
pig context -m postgres          # only output PostgreSQL module (host included by default)
pig context -m postgres,!host    # exclude host module
```


## pig pg

Manage the local PostgreSQL server. See [`pig pg`](https://pig.pgsty.com/pg/) for details.

```bash
pig pg init                      # initialize data directory
pig pg start                     # start PostgreSQL
pig pg stop                      # stop PostgreSQL
pig pg status                    # check status
pig pg psql mydb                 # connect to database
pig pg ps                        # show current connections
pig pg vacuum mydb               # vacuum database
pig pg tune -p olap              # generate tuned parameters
pig pg fork dev                  # create a local one-off physical fork
pig pg fork list                 # list local forks
pig pg log tail                  # tail logs in real time
```


## pig pt

Run `patronictl` transparently to manage Patroni HA clusters. See [`pig pt`](https://pig.pgsty.com/pt/) for details.

```bash
pig pt list pg-meta              # list cluster members (native passthrough)
pig pt show-config pg-meta       # show cluster dynamic config (native passthrough)
pig pt set ttl=60                # local sugar: update Patroni/PG settings
pig pt restart pg-test --pending # apply pending restarts (native passthrough)
pig pt status                    # local: combined service and cluster status
pig pt log -f                    # local: tail logs in real time
```


## pig pb

Manage pgBackRest backup and restore. See [`pig pb`](https://pig.pgsty.com/pb/) for details.

```bash
pig pb info                      # show backup information
pig pb ls                        # list all backups
pig pb backup                    # create backup
pig pb backup full               # full backup
pig pb restore -d                # restore to latest
pig pb restore -t "2025-01-01"   # restore to a specific time
pig pb log tail                  # tail logs in real time
```


## pig pitr

Run orchestrated point-in-time recovery (PITR). See [`pig pitr`](https://pig.pgsty.com/pitr/) for details.

```bash
pig pitr -d                      # recover to latest data
pig pitr -t "2025-01-01 12:00:00+08"  # recover to a specific time
pig pitr -I                      # recover to backup consistency point
pig pitr -d --plan               # show execution plan without running
pig pitr -d -y                   # skip confirmation for automation
```


## Helper Commands

```bash
pig status                       # show current environment status
pig status -o json               # structured status output
pig update                       # upgrade pig itself to the latest version
pig update -m                    # upgrade using the pigsty.cc mirror
pig update -v X.Y.Z              # upgrade to a version published by the configured repository
pig version                      # show pig version information
pig version -o json              # structured version output
```
