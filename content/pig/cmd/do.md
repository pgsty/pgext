---
title: "pig do"
description: "Run Pigsty administration tasks with pig do"
weight: 620
icon: fas fa-list-check
---

`pig do` runs administration playbooks from the selected Pigsty directory. It uses the inventory to add clusters, create databases and users, install packages, or remove services.

Define the target configuration in the inventory first, then run the corresponding command. The commands invoke Ansible directly. This command group has no general `--plan` mode.

```bash
pig do pgsql-add  pg-test 10.10.10.12     # add one PostgreSQL instance
pig do pgsql-rm   pg-test 10.10.10.12     # remove one PostgreSQL instance
pig do pgsql-db   pg-meta app             # create or update a database
pig do pgsql-user pg-meta dbuser_app      # create or update a user
pig do node-pkg   pg-meta restic          # install a package on selected nodes
pig do redis-add  redis-meta              # initialize a Redis cluster
```

Inventory selection is shared with [`sty deploy`](/pig/cmd/sty/#sty-deploy). PIG passes the resolved
source to Ansible without interpreting it again relative to Pigsty home. Missing explicit,
environment, or default sources fail before playbook execution. Existing directory and executable
Inventory sources remain Ansible inputs; static Inventory commands do not execute them.

## Command overview

| Command | Purpose | Required arguments |
|:---|:---|:---|
| `pgsql-add` | Add a PostgreSQL cluster or instances | cluster, optional IPs |
| `pgsql-rm` | Remove a PostgreSQL cluster or instances | cluster, optional IPs |
| `pgsql-db` | Apply a database declaration | cluster, database |
| `pgsql-user` | Apply a user declaration | cluster, user |
| `pgsql-ext` | Install extensions through Pigsty | cluster, optional extensions |
| `pgsql-hba` | Refresh PostgreSQL HBA rules | cluster |
| `pgsql-svc` | Refresh PostgreSQL service definitions | cluster |
| `pgmon-add` / `pgmon-rm` | Add or remove a remote monitoring target | cluster |
| `node-add` / `node-rm` | Add or remove selected nodes | one or more selectors |
| `node-repo` | Configure repository modules on selected nodes | optional selector and modules |
| `node-pkg` | Install or update packages on selected nodes | selector, optional packages |
| `repo-build` | Rebuild the Pigsty infrastructure repository | none |
| `redis-add` / `redis-rm` | Add or remove Redis clusters, nodes, or instances | selector, optional ports |
{.full-width}

Run `pig do COMMAND --help` for the aliases and exact examples of one operation.

## PostgreSQL operations

```bash
pig do pgsql-add pg-meta                  # initialize the declared cluster
pig do pgsql-add pg-test 10.10.10.12      # add one declared replica
pig do pgsql-rm pg-test 10.10.10.13       # remove one instance
pig do pgsql-rm pg-test --uninstall       # also uninstall packages during removal
pig do pgsql-db pg-meta meta              # create or update database meta
pig do pgsql-user pg-meta dbuser_view     # create or update a user
pig do pgsql-ext pg-meta postgis vector   # install extension packages
pig do pgsql-hba pg-meta                  # refresh pg_hba rules
pig do pgsql-svc pg-meta                  # refresh PostgreSQL services
```

`pgsql-db` and `pgsql-user` apply existing `pg_databases` and `pg_users` declarations. They do not edit the inventory. `pgsql-rm --uninstall` also removes packages. Omit this option if the host still needs those packages.

### Name rules

| Object | Accepted names |
|:---|:---|
| PostgreSQL cluster | Letters, digits, and hyphens; start with a letter or digit. `root`, `all`, and `ungrouped` are reserved. |
| PostgreSQL user | Up to 63 characters; start with a lowercase letter or `_`, then use lowercase letters, digits, `_`, `@`, `.`, or `-`. `postgres` is reserved. |
| Database | Up to 63 characters; start with a letter or `_`, then use letters, digits, `_`, `@`, `.`, or `-`. |

These rules apply to arguments that require a cluster, user, or database name. Selector arguments support the target-selection forms below.

## Node and repository operations

Selectors may be cluster names, host names, IP addresses, or another selector form supported by
the Pigsty playbook. Their meaning comes from the active Inventory.

```bash
pig do node-add pg-test                    # add nodes selected by cluster
pig do node-rm 10.10.10.13                 # remove one selected node
pig do node-repo pg-meta node,infra        # configure selected repo modules
pig do node-pkg pg-meta openssh restic     # install or update packages
pig do repo-build                          # rebuild the infra repository
```

With no selector, `node-repo` applies to every host targeted by `node.yml`. With no module, it uses the Inventory setting. Repository modules come from the installed Pigsty release, such as `local`,
`infra`, `pgsql`, `node`, and `extra`. Check that the required repositories are available to the
target nodes before installing packages.

## Redis operations

```bash
pig do redis-add redis-meta                # initialize a declared Redis cluster
pig do redis-add 10.10.10.11 6379 6380     # add selected instances
pig do redis-rm 10.10.10.11 6379           # remove one selected instance
pig do redis-rm redis-test --uninstall     # remove and uninstall packages
```

`redis-rm --uninstall` also removes packages during node or cluster removal.

## Configuration and output

`pig do` uses the global PIG options:

| Option | Purpose |
|:---|:---|
| `-H, --home` | Select the Pigsty home directory |
| `-i, --inventory` | Select the Pigsty Inventory |
| `-o, --output` | Choose text, JSON, YAML, or pretty JSON result wrapping |
| `--log-level`, `--log-path` | Configure PIG diagnostics |
{.full-width}

Structured mode reports the requested operation and captured execution result. The underlying
Ansible and playbook behavior is still determined by the installed Pigsty release.

See the [Pigsty documentation](https://pigsty.io/docs/) for full playbook variables and procedures.
