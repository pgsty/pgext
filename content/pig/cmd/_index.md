---
title: "CMD Docs"
description: Overview of the pig CLI command reference
icon: fa-solid fa-terminal
weight: 600
---

`pig` manages PostgreSQL packages, local database services, and Pigsty deployments. Choose a command below for your task. Use `pig COMMAND --help` to see the options supported by your installed program.

- [**pig repo**](/pig/cmd/repo/): manage software repositories
- [**pig ext**](/pig/cmd/ext/): manage PostgreSQL extensions
- [**pig build**](/pig/cmd/build/): build extensions from source
- **pig install**: install packages with the native package manager and translate PostgreSQL aliases
- [**pig sty**](/pig/cmd/sty/): manage Pigsty installation and Grafana dashboards
- [**pig inventory**](/pig/cmd/inventory/): inspect, edit, validate, and exchange the Pigsty inventory
- [**pig do**](/pig/cmd/do/): run Pigsty administration playbook tasks
- [**pig pe**](/pig/cmd/pe/): access pg_exporter metrics and configuration
- [**pig pg**](/pig/cmd/pg/): manage local PostgreSQL servers
- [**pig pt**](/pig/cmd/pt/): run patronictl transparently with service and config helpers
- [**pig pb**](/pig/cmd/pb/): manage pgBackRest backup and restore
- [**pig pitr**](/pig/cmd/pitr/): run the full PITR workflow
- **pig context**: output an environment context snapshot for humans and agents
- **pig license**: issue and verify Pigsty licenses
- **pig status / update / version**: inspect environment status, upgrade pig, and print version information

## Global options

| Option | Purpose |
|:---|:---|
| `-h, --help` | Show help |
| `-H, --home PATH` | Select the Pigsty installation directory |
| `-i, --inventory PATH` | Select the inventory |
| `-o, --output FORMAT` | Output format: `text` (default), `json`, `yaml`, or `json-pretty` |
| `--debug` | Enable debug logging |
| `--log-level LEVEL` | Log level: `debug`, `info` (default), `warn`, `error`, `fatal`, or `panic` |
| `--log-path PATH` | Write logs to a file; the default is the terminal |

PIG-owned results support structured output. Raw-content commands such as `inventory show` and passthrough commands such as `pt list` retain their native output. See the command reference for details.

Structured results keep diagnostics and wrapped-tool chatter on stderr. JSON/YAML output and
`PIG_NON_INTERACTIVE=1` disable prompt input in PIG-owned command adapters and use `sudo -n`;
they do not supply destructive-operation confirmation. Pass `--yes` where required.
Documented data input, such as `pig pg psql -f -`, remains available. Native Patronictl
forwarding retains its own interaction and output contract and rejects PIG structured-output modes.

```bash
pig repo list
pig ext list vector -o json
pig help inventory edit
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
pig ext update  pg_duckdb        # update the named extension package
pig ext scan                     # scan installed extensions
pig ext import  pg_duckdb        # download for offline use
pig ext link    17               # link PG version into PATH
pig ext reload                   # refresh extension catalog
```

## pig build

Build PostgreSQL extensions from source. See [`pig build`](/pig/cmd/build/) for details.

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
pig build proxy                  # install or verify Xray
pig build proxy client --from ./client.uri
pig build proxy server --host proxy.example.com --target www.sraoss.co.jp:443 --export ./client.uri

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
pig install vector -y         # auto-confirm installation
```

## pig sty

Install the Pigsty distribution. See [`pig sty`](/pig/cmd/sty/) for details.

```bash
pig sty init                     # install Pigsty to ~/pigsty
pig sty boot                     # install Ansible prerequisites
pig sty conf -g                  # generate configuration with random passwords
pig sty deploy                   # run deployment playbook
pig sty list                     # list available Pigsty releases
pig sty get {{< param pigsty_version >}}                # download the embedded Pigsty release
pig sty grafana list             # manage Grafana dashboards (info/list/boot/load/init/dump/clean/lang/style)
```

## pig inventory

Inspect, edit, validate, check, and exchange the Pigsty inventory (`pigsty.yml`) with the
CMDB. Root-level command group with alias `inv`; see [`pig inventory`](/pig/cmd/inventory/)
for details.

```bash
pig inventory status             # inspect the active inventory source
pig inventory list               # list inventory topology and value kinds
pig inventory show               # show verbatim inventory YAML (may contain secrets)
pig inventory edit               # edit the inventory (or one fragment) in $EDITOR
pig inventory validate           # validate one complete static Pigsty inventory
pig inventory check              # check inventory and local controller readiness
pig inventory diff other.yml     # compare declarations without emitting values
pig inventory cmdb check         # check the PostgreSQL CMDB (experimental)
```

## pig do

Run Pigsty administration tasks through the corresponding Ansible playbooks. See
[`pig do`](/pig/cmd/do/) for the command reference.

```bash
pig do pgsql-add  <cls> [ip...]       # add cluster or instances
pig do pgsql-rm   <cls> [ip...]       # remove cluster or instances
pig do pgsql-db   <cls> <dbname>      # create or update database
pig do pgsql-user <cls> <username>    # create or update user
pig do pgsql-ext  <cls> [ext...]      # install extensions
pig do node-pkg   <sel> [pkg...]      # install node packages
```

## pig pe

Access PostgreSQL monitoring metrics exposed by pg_exporter. The default endpoint is
`127.0.0.1:9630`; see [`pig pe`](/pig/cmd/pe/) for endpoint and operation details.

```bash
pig pe list                    # list available metric types
pig pe get                     # get all pg_ prefixed metrics
pig pe stat                    # show exporter statistics
pig pe reload                  # reload pg_exporter configuration
pig pe --host 127.0.0.1 -p 9630 get
```

## pig license {#license}

`pig license issue <name>` delivers the newly issued license to its caller. Text output includes
the token; JSON/YAML preserve it in `data.captured_output`. Successful issuance is an explicit
credential-delivery operation, so that result must be stored as a credential. The signing key is
never included in arguments, parameters, or diagnostics returned in the result. Raw key-loading
errors are replaced with fixed messages. Parameter errors and safe signing-key failure reasons
remain available in text, JSON, and YAML. `license status` and `license verify` do not reveal tokens.

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

Run `patronictl` transparently to manage Patroni HA clusters. See [`pig pt`](/pig/cmd/pt/) for details.

```bash
pig pt list pg-meta              # list cluster members (native passthrough)
pig pt show-config pg-meta       # show cluster dynamic config (native passthrough)
pig pt set ttl=60                # local sugar: update Patroni/PG settings
pig pt restart pg-test --pending # apply pending restarts (native passthrough)
pig pt status                    # local: combined service and cluster status
pig pt log -f                    # local: tail logs in real time
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

`pig update` uses the native package manager on supported Linux distributions. On macOS,
download the new archive and replace the executable as described in [Installation](/pig/install/).

If `~/.pig/config.yml` has a syntax error, fix it before running commands that read configuration, such as `pig status`. `pig version` and the `pig pe` HTTP commands remain available because they do not read this file.

## Design records

- [From Human-Friendly to Agent-Native](https://pig.pgsty.com/design/agent-native-cli/)
- [Why PIG Keeps the Cobra Command Layer Flat](https://pig.pgsty.com/design/flat-cobra-command-layer/)
- [PIG's Operations CLI Safety Contract](https://pig.pgsty.com/design/ops-cli-safety/)
