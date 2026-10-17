---
title: "CMD: repo"
description: How to manage repositories with pig repo subcommand?
icon: fa-solid fa-terminal
weight: 610
---

`pig repo` manages Linux package repositories. Use it to add or remove APT/YUM/DNF sources, refresh indexes, and create or extract offline bundles. `repo create` also supports macOS through SOW.

| Command | Description | Notes |
|:---|:---|:---|
| `repo list` | Print available repositories and modules | |
| `repo info` | Show repository details | |
| `repo status` | Show current repository status | |
| `repo add` | Add repositories | Requires sudo or root |
| `repo set` | Clear, overwrite, and update repositories | Requires sudo or root |
| `repo rm` | Remove repositories | Requires sudo or root |
| `repo update` | Update repository cache | Requires sudo or root |
| `repo create` | Create local YUM/APT repository | Linux: sudo/root; macOS: current user |
| `repo cache` | Create offline package from local repo | Requires sudo or root |
| `repo boot` | Bootstrap repository from offline package | Requires sudo or root |
| `repo reload` | Refresh repository catalog | |
{.full-width}

## Quick Start

On a host with working OS repositories, add the PostgreSQL repositories and refresh their indexes:

```bash
pig repo add pgsql -u
```

To use the China-region mirrors, add `-m`:

```bash
pig repo add pgsql -m -u
```

To replace the host's repository configuration, use `set`. It backs up the existing definitions, adds `node`, `infra`, and `pgsql`, then refreshes the indexes:

```bash
pig repo set
```

Use `pig repo info pgsql` to inspect repository definitions. Use `pig repo status` to inspect the host's current configuration.

## Modules

In pig, APT/YUM repositories are organized as **modules**: groups of repositories serving a specific purpose.

| Module | Description | Repository List |
|:---:|:---|:---|
| `all` | All core modules required to install PG | `node` + `infra` + `pgsql` |
| `pgsql` | PGDG + Pigsty PG extensions | `pigsty-pgsql` + `pgdg` |
| `pigsty` | Pigsty Infra + PGSQL repositories | pigsty-infra, pigsty-pgsql |
| `pgdg` | PGDG official repositories | EL: pgdg-common, pgdg14-18; APT: pgdg |
| `node` | Linux system repositories | base, updates, extras, epel, baseos, appstream... |
| `infra` | Infrastructure component repository | pigsty-infra |
| `nginx` | Nginx repository | nginx |
| `docker` | Docker repository | docker-ce |
| `beta` | PostgreSQL 19 beta repositories | pgdg19-beta, pgdg-beta |
| `extra` | PGDG non-free and third-party extensions | pgdg-extras, timescaledb, citus |
| `groonga` | PGroonga repository | groonga |
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

PIG filters each module by the current operating system and architecture. Run `pig repo list` to see the available entries.

If the host already has working OS repositories, add `pgsql` for PGDG and Pigsty extension packages. `infra` supplies infrastructure tools and some PostgreSQL fork packages. `all` combines `node`, `infra`, and `pgsql`. Both `repo add` and `repo set` select `all` when no module is given.

`repo add`, `repo set`, and `repo info` accept comma-separated names and surrounding whitespace.
Explicit empty names, including an empty argument or `node,`, reject the whole request before
repository operations. An empty argument is not equivalent to omitting the arguments.

## Repository Definitions

The full repository definition bundled with Pigsty is in [`cli/repo/assets/repo.yml`](https://github.com/pgsty/pig/blob/main/cli/repo/assets/repo.yml).

Create `~/.pig/repo.yml` to override the embedded catalog. Use the source file above as the format reference. `repo reload` updates this file, so save custom changes before a refresh.

`--region` selects a regional URL from `baseurl`. If that region is absent, PIG uses `default`. `-m` selects `china`, including the Pigsty China mirror and other mirrors defined in the catalog.

### Trust and configuration ownership

`repo add` writes the selected module files and preserves other files. `repo set` and `repo add --remove` back up and replace the existing definitions. If Ansible or another tool manages these files, keep the definitions in that tool's configuration so its next run preserves your changes.

### Pigsty signing key {#repo-key}

When a module includes `pigsty-infra` or `pigsty-pgsql`, PIG reuses the existing public-key file or installs its embedded key with mode `0644`:

| Platform | Public-key path |
|:---|:---|
| EL | `/etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty` |
| Debian / Ubuntu | `/etc/apt/keyrings/pigsty.asc` |

This step works offline. Selecting only other repositories, such as `pgdg` or `node`, does not prepare the Pigsty key.

**Built-in repository definitions do not enforce package signatures by default**: EL uses `gpgcheck=0`; Debian/Ubuntu uses `trusted=yes`. Ordinary `repo add/set` preserves these settings. [`pig sty boot`](/pig/cmd/sty/#sty-boot) enables Pigsty signature checks when key preparation succeeds. To require signatures, configure a trusted key and enable verification in the repository metadata.

If implicit key installation fails, the command warns and continues. The selected Pigsty definitions use `trusted=yes`, or `gpgcheck=0` with `repo_gpgcheck=0`, and retain their default key reference. Structured results include the warning in `data.warnings`. APT may retain old indexes in this case; check the cache-update output.

An explicit key reference makes preparation mandatory:

- On Debian/Ubuntu, `signed-by` names existing regular keyring files by absolute path. Separate multiple paths with commas.
- On EL, `gpgkey` contains paths or literal URLs, which PIG passes to `rpm --import`.
- PIG can install an explicitly referenced default Pigsty key from its embedded copy.

If preparation of an explicit key fails, repository configuration stops. Implicit installation prepares only the file and repository reference; explicit EL `gpgkey` also imports the key into RPM. See the [design record](https://pig.pgsty.com/design/automatic-pigsty-repo-key/) for the rationale.

EL repositories retain DNF module filtering. Definitions with `module_hotfixes=1`, including Pigsty and PGDG, bypass module streams. PIG omits this key in EL7 configuration.

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
pig repo add pgsql -m            # explicitly select the China-region mirrors
```

**Options:**

- `-r|--remove`: remove existing repositories before adding new ones
- `-u|--update`: run package cache update after adding repositories
- `-m|--mirror`: explicitly select the bundled `china` repository definitions
- `--region <region>`: use regional mirror repositories (`default` / `china` / `europe`)

| Platform | Module Location |
|:---:|:---|
| EL | `/etc/yum.repos.d/<module>.repo` |
| Debian | `/etc/apt/sources.list.d/<module>.list` |
{.full-width}

## repo set

Back up and remove existing repository definitions, write the selected modules, then refresh package metadata.
This is equivalent to `repo add --remove --update`. If you omit the modules, PIG uses `all`.

```bash
pig repo set                     # replace with default repositories
pig repo set pgdg pigsty         # replace with selected repositories and update
pig repo set all --region=china  # use China mirror
pig repo set -m                  # explicitly select the China-region mirrors
```

`repo set` supports the same `--region` and `-m|--mirror` options as `repo add`.

## repo rm

With truly no operands, move all repository definitions to the existing backup directory.
Explicit module operands remove only those module files; `all` expands to `node`, `pgsql`, and
`infra`, and does not mean every file. Empty strings, whitespace-only operands, empty comma
items, and unsafe names fail before any repository or cache operation, even in a mixed list.
`-u` updates the cache after either a successful full backup or module removal. Backup and
restore behavior is unchanged.
Removing an already absent module remains successful. The summary reports how many module files
are now absent, without claiming that every request deleted an existing file. Custom module names
remain supported within the repository directory.
If a full backup finds no repository files, the summary reports that there is nothing to back up.
Structured results retain the original operands in `requested_modules`, including `all` and comma-separated
arguments; `removed_repos` lists the expanded modules and their individual outcomes.

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
| EL | `dnf makecache` (`yum makecache` on EL7) |
| Debian | `apt-get update` |
{.full-width}

## repo create

Create a local package repository for offline installation.

```bash
pig repo create                  # Linux: /www/pigsty; macOS: current directory
pig repo create /srv/repo        # create at custom location
```

The current implementation prefers [`sow`](https://sow.pgsty.com/docs/reference/cli/create/) when it is available in `PATH`. On Linux, the SOW backend runs the equivalent of the following with sudo for every target directory:

```bash
sow create --pigsty --timeout 10m -- /absolute/repository/path
```

On macOS, install SOW first. It runs without `sudo`, and the default target is the current directory.

On Linux, if SOW is absent, PIG uses `createrepo_c` on EL or `dpkg-scanpackages` from `dpkg-dev` on Debian/Ubuntu.
If SOW runs and fails, PIG returns the error. The `10m` timeout limits the wait for SOW's directory lock.

SOW's `--pigsty` mode:

1. Scans only top-level regular `.rpm` and `.deb` files—no recursion and no symlink following.
2. Removes DEB packages for `i386`, and RPM/DEB packages named `patroni` with upstream version `3.0.4`.
3. Generates RPM/DEB metadata and writes `repo_complete` last. The marker contains SHA-256 hashes for remaining top-level packages, sorted by basename.

Documentation packages such as `annobin-docs` are retained unless another cleanup rule applies; a filename containing `docs` is not a reason to delete a package.

SOW preserves non-package files and directories. It stops if a package cannot be parsed or package coordinates conflict. The legacy backend uses different cleanup rules and records MD5 hashes in `repo_complete`. PIG checks that this file exists; it does not check its hashes. Use the list to verify package integrity before distributing an offline bundle.

## repo cache

Create a compressed tarball of repository contents for offline distribution.

```bash
pig repo cache                   # default: /www to /tmp/pkg.tgz
pig repo cache -d /srv           # custom source directory
pig repo cache pigsty docker     # include multiple top-level repositories
```

**Options:**

- `-d, --dir`: source directory, default `/www/`
- `-p, --path`: output path, default `/tmp/pkg.tgz`

An offline bundle must include `pigsty` and a regular `pigsty/repo_complete` marker. Missing
requested repositories are skipped and reported separately from the repositories actually included.
Symbolic links are rejected because the native boot extractor accepts only directories and regular
files. A failed cache leaves an existing output package unchanged.

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

The command extracts into a staging directory, checks target paths, then moves the repositories into place. It moves additional repositories first and `pigsty` last.

- If a completed `pigsty/repo_complete` already exists at the target, PIG reuses that repository. It does not inspect the archive or import its additional repositories. Structured output sets `reused: true`.
- A conflicting target path stops extraction. After an interrupted run, inspect and move any conflicting directories before retrying.
- The archive must contain a regular `pigsty/repo_complete` file. PIG checks its presence, not its contents.

## repo reload

Download the latest repository catalog from the Pigsty repository sites.

```bash
pig repo reload                  # refresh repository catalog
```

The updated file is placed in `~/.pig/repo.yml`.
