---
title: "CMD: repo"
description: How to manage repositories with pig repo subcommand?
icon: SquareTerminal
weight: 610
---

The `pig repo` command is a comprehensive package repository manager. It can add, remove, create, and manage repositories on RPM systems (RHEL/CentOS/Rocky/Alma) and Debian systems (Debian/Ubuntu).

```bash
pig repo - Manage Linux APT/YUM Repo

  pig repo list                    # available repo list             (info)
  pig repo info   [repo|module...] # show repo info                  (info)
  pig repo status                  # show current repo status        (info)
  pig repo add    [repo|module...] # add repo and modules            (root)
  pig repo rm     [repo|module...] # remove repo & modules           (root)
  pig repo update                  # update repo pkg cache           (root)
  pig repo create                  # create repo on current system   (root)
  pig repo boot                    # boot repo from offline package  (root)
  pig repo cache                   # cache repo as offline package   (root)

Examples:
  pig repo add -ru                 # add all repo and update cache (brute but effective)
  pig repo add pigsty -u           # gentle version, only add pigsty repo and update cache
  pig repo add node pgdg pigsty    # essential repo to install postgres packages
  pig repo add all                 # all = node + pgdg + pigsty
  pig repo add all extra           # extra module has non-free and some 3rd repo for certain extensions
  pig repo update                  # update repo cache
  pig repo create                  # update local repo /www/pigsty meta
  pig repo boot                    # extract /tmp/pkg.tgz to /www/pigsty
  pig repo cache                   # cache /www/pigsty into /tmp/pkg.tgz
```

| Command | Description | Notes |
|:---|:---|:---|
| `repo list` | Print available repositories and modules | |
| `repo info` | Show repository details | |
| `repo status` | Show current repository status | |
| `repo add` | Add repositories | Requires sudo or root |
| `repo set` | Clear, overwrite, and update repositories | Requires sudo or root |
| `repo rm` | Remove repositories | Requires sudo or root |
| `repo update` | Update repository cache | Requires sudo or root |
| `repo create` | Create local YUM/APT repository | Requires sudo or root |
| `repo cache` | Create offline package from local repo | Requires sudo or root |
| `repo boot` | Bootstrap repository from offline package | Requires sudo or root |
| `repo reload` | Refresh repository catalog | |
{.full-width}


## Quick Start

```bash
# Method 1: clean existing repositories, add all required repos, and update cache (recommended)
pig repo add all --remove --update    # remove old repos, add all required repos, update cache

# Method 1 variant: one step
pig repo set                          # = pig repo add all --remove --update

# Method 2: gentle mode - only add required repos and keep current repo configuration
pig repo add pgsql                    # add PGDG and Pigsty PGSQL repositories
pig repo add pigsty --region=china    # add Pigsty repositories with China region
pig repo add pgdg   --region=europe   # add PGDG repositories with Europe region
pig repo add infra  --region=default  # add INFRA repositories with default region

# If the commands above did not use -u|--update, run this as an extra step
pig repo update                       # update system package cache
```


## Modules

In pig, APT/YUM repositories are organized as **modules**: groups of repositories serving a specific purpose.

| Module | Description | Repository List |
|:---:|:---|:---|
| `all` | All core modules required to install PG | `node` + `infra` + `pgsql` |
| `pgsql` | PGDG + Pigsty PG extensions | `pigsty-pgsql` + `pgdg` |
| `pigsty` | Pigsty Infra + PGSQL repositories | pigsty-infra, pigsty-pgsql |
| `pgdg` | PGDG official repositories | pgdg-common, pgdg14-18 |
| `node` | Linux system repositories | base, updates, extras, epel, baseos, appstream... |
| `infra` | Infrastructure component repositories | pigsty-infra, nginx, docker-ce |
| `docker` | Docker repository | docker-ce |
| `beta` | PostgreSQL 19 beta repositories | pgdg19-beta, pgdg-beta |
| `extra` | PGDG non-free and third-party extensions | pgdg-extras, timescaledb, citus |
| `groonga` | PGroonga repository | groonga |
| `mssql` | WiltonDB repository (deprecated) | babelfish |
| `percona` | Percona PG + PG_TDE | percona |
| `llvm` | LLVM toolchain repository | llvm |
| `kube` | Kubernetes repository | kubernetes |
| `grafana` | Grafana repository | grafana |
| `haproxy` | HAProxy repositories | haproxyd, haproxyu |
| `redis` | Redis repository | redis |
| `mongo` | MongoDB repository | mongo |
| `mysql` | MySQL repository | mysql |
| `click` | ClickHouse repository | clickhouse |
| `gitlab` | GitLab repository | gitlab-ce, gitlab-ee |
{.full-width}

Pig also includes APT/DNF repositories for other databases and systems such as `redis`, `kubernetes`, `grafana`, `clickhouse`, `gitlab`, `haproxy`, `mongodb`, and `mysql`.

In general, `node` (Linux system repositories) and `pgsql` (PGDG + Pigsty) are required for PostgreSQL installation. The `infra` repository is optional and contains tools, IvorySQL kernel packages, and similar components. The special `all` module adds all required repositories at once and is a suitable starting point for most users.

```bash
pig repo add all      # add node, pgsql, and infra repositories
pig repo add          # default to the all module when no argument is given
pig repo set          # set clears/backups existing definitions and overwrites with new repo definitions
```


## Repository Definitions

The full repository definition bundled with Pigsty is in [`cli/repo/assets/repo.yml`](https://github.com/pgsty/pig/blob/main/cli/repo/assets/repo.yml).

You can create `~/.pig/repo.yml` to explicitly modify and override pig's repository definitions. When editing repository definitions, you can add extra regional mirror URLs under `baseurl`, such as China or Europe mirrors. When `--region` is specified, pig first looks for the matching regional URL and falls back to the `default` URL if the region is unavailable.


## repo list

`pig repo list` lists all repository modules available on the current system.

```bash
pig repo list                # list repositories available on current system
pig repo list all            # list all repositories without filtering
```


## repo info

Show detailed information for specific repositories or modules, including URLs, metadata, regional mirrors, and `.repo` / `.list` repository file content.

```bash
pig repo info pgdg               # show pgdg module information
pig repo info pigsty pgdg        # show multiple modules
pig repo info all                # show all modules
```


## repo status

Show current repository configuration on the system.

```bash
pig repo status
```


## repo add

Add repository configuration files to the system. Requires root/sudo privileges.

```bash
pig repo add pgdg                # add PGDG repository
pig repo add pgdg pigsty         # add multiple repositories
pig repo add all                 # add all required repositories (pgdg + pigsty + node)
pig repo add pigsty -u           # add and update cache
pig repo add all -r              # remove existing repos before adding
pig repo add all -ru             # remove, add, and update (full reset)
pig repo add pgdg --region=china # use China mirror
```

**Options:**

- `-r|--remove`: remove existing repositories before adding new ones
- `-u|--update`: run package cache update after adding repositories
- `--region <region>`: use regional mirror repositories (`default` / `china` / `europe`)

| Platform | Module Location |
|:---:|:---|
| EL | `/etc/yum.repos.d/<module>.repo` |
| Debian | `/etc/apt/sources.list.d/<module>.list` |
{.full-width}


## repo set

Equivalent to `repo add --remove --update`. It clears existing repositories, sets up new ones, then updates cache.

```bash
pig repo set                     # replace with default repositories
pig repo set pgdg pigsty         # replace with selected repositories and update
pig repo set all --region=china  # use China mirror
```


## repo rm

Remove repository configuration files and back them up.

```bash
pig repo rm                      # remove all repositories
pig repo rm pgdg                 # remove selected repository
pig repo rm pgdg pigsty -u       # remove and update cache
```

| Platform | Backup Location |
|:---:|:---|
| EL | `/etc/yum.repos.d/backup/` |
| Debian | `/etc/apt/sources.list.d/backup/` |
{.full-width}


## repo update

Update package-manager cache to reflect repository changes.

```bash
pig repo update                  # update package cache
```

| Platform | Equivalent Command |
|:---:|:---|
| EL | `dnf makecache` |
| Debian | `apt update` |
{.full-width}


## repo create

Create a local package repository for offline installation.

```bash
pig repo create                  # create at default location (/www/pigsty)
pig repo create /srv/repo        # create at custom location
```

| Platform | Dependency |
|:---:|:---|
| EL | `createrepo_c` |
| Debian | `dpkg-dev` |
{.full-width}


## repo cache

Create a compressed tarball of repository contents for offline distribution.

```bash
pig repo cache                   # default: /www to /tmp/pkg.tgz
pig repo cache -d /srv           # custom source directory
```

**Options:**

- `-d, --dir`: source directory, default `/www/`
- `-p, --path`: output path, default `/tmp/pkg.tgz`


## repo boot

Extract and set up a local repository from an offline package.

```bash
pig repo boot                    # default: /tmp/pkg.tgz to /www
pig repo boot -p /mnt/pkg.tgz    # custom package path
pig repo boot -d /srv            # custom target directory
```

**Options:**

- `-p, --path`: package path, default `/tmp/pkg.tgz`
- `-d, --dir`: target directory, default `/www/`


## repo reload

Refresh repository metadata from GitHub to the latest version.

```bash
pig repo reload                  # refresh repository catalog
```

The updated file is placed in `~/.pig/repo.yml`.
