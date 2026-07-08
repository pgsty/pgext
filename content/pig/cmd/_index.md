---
title: "CMD Docs"
description: Overview of the pig CLI command reference
icon: SquareTerminal
weight: 600
---

The `pig` CLI provides a comprehensive toolkit for managing PostgreSQL installations, extensions, repositories, and extension builds from source. Use `pig help <command>` to view command documentation.

- [**pig repo**](/pig/cmd/repo/): manage software repositories
- [**pig ext**](/pig/cmd/ext/): manage PostgreSQL extensions
- [**pig build**](/pig/cmd/build/): build extensions from source
- **pig install**: install packages with the native package manager and translate PostgreSQL aliases
- [**pig sty**](/pig/cmd/sty/): manage Pigsty installation
- [**pig pg**](/pig/cmd/pg/): manage local PostgreSQL servers
- [**pig pt**](/pig/cmd/pt/): manage Patroni HA clusters
- [**pig pb**](/pig/cmd/pb/): manage pgBackRest backup and restore
- [**pig pitr**](/pig/cmd/pitr/): run the full PITR workflow
- **pig context**: output an environment context snapshot for humans and agents


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
  do          Run admin tasks
  postgres    Manage local PostgreSQL server and databases (alias: pg)
  patroni     Manage Patroni cluster with patronictl (alias: pt)
  pgbackrest  Manage pgBackRest backup and restore (alias: pb)
  pg_exporter Manage pg_exporter and metrics (alias: pe)
  pitr        Orchestrated point-in-time recovery
  sty         Manage Pigsty installation
  context     Show environment context snapshot

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

Manage APT/YUM repositories for PostgreSQL packages. See [`pig repo`](/pig/cmd/repo/) for details.

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

Manage PostgreSQL extensions and kernel packages. See [`pig ext`](/pig/cmd/ext/) for details.

```bash
pig ext list    duck             # search extensions
pig ext info    pg_duckdb        # extension details
pig ext status                   # show installed extensions
pig ext add     pg_duckdb -y     # install extension
pig ext rm      old_extension    # remove extension
pig ext update  pg_duckdb        # update selected extension
pig ext scan                     # scan installed extensions
pig ext import  pg_duckdb        # download for offline use
pig ext link    polar            # link PG/PolarDB installation into PATH
pig ext reload                   # refresh extension catalog
```


## pig build

Build PostgreSQL extensions from source. See [`pig build`](/pig/cmd/build/) for details.

```bash
# Environment setup
pig build spec                   # initialize build specs
pig build repo -m                # configure repositories with mirror/proxy
pig build tool                   # install build tools
pig build rust -m                # install Rust with mirror config
pig build pgrx -b                # install PGRX and include beta PG auto-detect

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

Install the Pigsty distribution. See [`pig sty`](/pig/cmd/sty/) for details.

```bash
pig sty init -m                  # install Pigsty from mirror to ~/pigsty
pig sty boot                     # install Ansible prerequisites
pig sty conf                     # generate configuration
pig sty deploy                   # run deployment playbook
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

Manage the local PostgreSQL server. See [`pig pg`](/pig/cmd/pg/) for details.

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

Manage Patroni HA clusters. See [`pig pt`](/pig/cmd/pt/) for details.

```bash
pig pt list                      # list cluster members
pig pt config show               # show cluster configuration
pig pt config set ttl=60         # modify cluster configuration
pig pt status                    # check service status
pig pt log -f                    # tail logs in real time
```


## pig pb

Manage pgBackRest backup and restore. See [`pig pb`](/pig/cmd/pb/) for details.

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

Run orchestrated point-in-time recovery (PITR). See [`pig pitr`](/pig/cmd/pitr/) for details.

```bash
pig pitr -d                      # recover to latest data
pig pitr -t "2025-01-01 12:00:00+08"  # recover to a specific time
pig pitr -I                      # recover to backup consistency point
pig pitr -d --plan               # show execution plan without running
pig pitr -d -y                   # skip confirmation for automation
```
