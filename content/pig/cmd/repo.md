---
title: "CMD: repo"
description: How to manage repositories with pig repo subcommand?
icon: SquareTerminal
weight: 610
---

The `pig repo` command is a comprehensive tool for managing package repositories on Linux systems. It provides functionality to add, remove, create, and manage software repositories for both RPM-based (RHEL/CentOS/Rocky/Alma) and Debian-based (Debian/Ubuntu) distributions. This command is essential for setting up the required repositories to install PostgreSQL and its extensions.

------

## Subcommands

| Command | Description |
|---------|------------|
| `list`   | Print available repo list |
| `info`   | Get repo detailed information |
| `status` | Show current repo status |
| `add`    | Add new repository |
| `set`    | Wipe, overwrite, and update repository |
| `rm`     | Remove repository |
| `update` | Update repo cache |
| `create` | Create local YUM/APT repository |
| `cache`  | Create offline package from local repo |
| `boot`   | Bootstrap repo from offline package |

------

## Overview

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

Usage:
  pig repo [command]

Aliases:
  repo, r

Available Commands:
  add         add new repository
  boot        bootstrap repo from offline package
  cache       create offline package from local repo
  create      create local YUM/APT repository
  info        get repo detailed information
  list        print available repo list
  rm          remove repository
  set         wipe, overwrite, and update repository
  status      show current repo status
  update      update repo cache

Flags:
  -h, --help   help for repo

Global Flags:
      --debug              enable debug mode
  -H, --home string        pigsty home path
  -i, --inventory string   config inventory path
      --log-level string   log level: debug, info, warn, error, fatal, panic (default "info")
      --log-path string    log file path, terminal by default
```

------

## Quick Start

The fastest way to set up repositories for PostgreSQL installation:

```bash
# Method 1: Add all essential repos and update cache (recommended)
pig repo add -ru                 # Remove old repos, add all essentials, update cache

# Method 2: Gentle approach - only add required repos
pig repo add pgdg pigsty -u      # Add PGDG and Pigsty repos with cache update

# Method 3: Add repos step by step
pig repo add pgdg                # Add PGDG official repo
pig repo add pigsty              # Add Pigsty extension repo
pig repo add node                # Add OS default repos
pig repo update                  # Update package cache
```

------

## Essential Concepts

### Modules

In pig, repositories are organized into **modules** - logical groups of repositories that serve a specific purpose. A module can map to different actual repositories depending on:
- Operating system distribution (RHEL/Rocky/Debian/Ubuntu)
- OS major version (8/9 for EL, 12 for Debian, 22/24 for Ubuntu)
- CPU architecture (x86_64/aarch64)
- Geographic region (default/china)

### Essential Modules

Three core modules are typically required for PostgreSQL installations:

1. **`pgdg`**: Official PostgreSQL Global Development Group repository
   - Contains PostgreSQL kernel packages
   - Includes 100+ official extensions
   - Provides PostgreSQL utilities and tools

2. **`pigsty`**: Pigsty Extension Repository
   - Contains 200+ additional extensions not in PGDG
   - Includes specialized extensions and utilities
   - Provides pre-compiled packages for various extensions

3. **`node`**: Operating System Default Repository
   - Contains system libraries and dependencies
   - Required for building and running PostgreSQL
   - Includes development tools and compilers

### The `all` Alias

The `all` pseudo-module is a convenient alias that includes all three essential modules (pgdg + pigsty + node). This is the recommended starting point for most users.

------

## Command Reference

### `repo list`

Lists all repository modules and individual repositories available for the current system.

```bash
pig repo list                    # List repos filtered for current system
pig repo list all                # List all repos without filtering
```

**Example Output:**
```yaml
repo_modules:   # Available Modules: 19
  - all       : pigsty-infra, pigsty-pgsql, pgdg, baseos, appstream, extras, powertools, crb, epel
  - pigsty    : pigsty-infra, pigsty-pgsql
  - pgdg      : pgdg
  - node      : baseos, appstream, extras, powertools, crb, epel
  - infra     : pigsty-infra, nginx
  - pgsql     : pigsty-pgsql, pgdg-common, pgdg13, pgdg14, pgdg15, pgdg16, pgdg17
  - extra     : pgdg-extras, timescaledb, citus
  - docker    : docker-ce
  - kube      : kubernetes
  # ... more modules
```

Repository definitions are stored in `~/.pig/repo.yml` or use the built-in defaults.

------

### `repo info`

Displays detailed information about specific repositories or modules, including URLs, metadata, and regional mirrors.

```bash
pig repo info pgdg               # Show info about pgdg module
pig repo info pigsty pgdg        # Show info about multiple modules
pig repo info all                # Show info about all modules
```

**Example Output:**
```bash
#-------------------------------------------------
Name       : pgdg
Summary    : PostgreSQL Global Development Group Official Repository
Available  : Yes (debian d12 amd64)
Module     : pgsql
OS Arch    : [x86_64, aarch64]
OS Distro  : deb [11,12,20,22,24], el [8,9]
Meta       : trusted=yes
Base URL   : http://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg main
     china : https://mirrors.tuna.tsinghua.edu.cn/postgresql/repos/apt/ ${distro_codename}-pgdg main

# Repository content for current system:
deb [trusted=yes] http://apt.postgresql.org/pub/repos/apt/ bookworm-pgdg main

# China mirror (if using --region china):
deb [trusted=yes] https://mirrors.tuna.tsinghua.edu.cn/postgresql/repos/apt/ bookworm-pgdg main
```

------

### `repo status`

Shows the current repository configuration on the system.

```bash
pig repo status
```

**Example Output:**
```bash
# For EL systems:
Repository Directory: /etc/yum.repos.d
Active Repositories:
  - pgdg.repo
  - pigsty.repo
  - rocky.repo

Total: 15 repositories enabled

# For Debian/Ubuntu systems:
Repository Directory: /etc/apt/sources.list.d
Active Repositories:
  - pgdg.list
  - pigsty.list
  - debian.list

Total: 8 repositories enabled
```

------

### `repo add`

Adds repository configuration files to the system. Requires root/sudo privileges.

```bash
# Basic usage
pig repo add pgdg                # Add PGDG repository
pig repo add pgdg pigsty         # Add multiple repositories
pig repo add all                 # Add all essential repos (pgdg + pigsty + node)

# With options
pig repo add pigsty -u           # Add and update cache
pig repo add all -r              # Remove existing repos before adding
pig repo add all -ru             # Remove, add, and update (complete reset)

# Regional mirrors
pig repo add pgdg --region china  # Use China mirrors
```

**Options:**
- `-u, --update`: Run package cache update after adding repos
- `-r, --remove`: Remove existing repos before adding new ones
- `--region <region>`: Use regional mirrors (default/china)

**File Locations:**
- EL Systems: `/etc/yum.repos.d/<module>.repo`
- Debian Systems: `/etc/apt/sources.list.d/<module>.list`
- Backups: `backup/` subdirectory in repo directory

**Common Workflows:**

```bash
# Fresh installation - recommended approach
sudo pig repo add -ru            # Clean slate approach

# Adding specific extension repos
sudo pig repo add pgdg -u        # Just PGDG
sudo pig repo add extra -u       # Additional third-party repos

# Setting up for air-gapped environment
sudo pig repo add local          # Use local repository
```

------

### `repo set`

Equivalent to `repo add --remove`. Wipes existing repositories and sets up new ones.

```bash
pig repo set                     # Replace with default repos
pig repo set pgdg pigsty -u      # Replace with specific repos and update
pig repo set all --region china  # Use China mirrors
```

This is useful when you want to ensure a clean repository configuration without old or conflicting repos.

------

### `repo rm`

Removes repository configuration files and backs them up.

```bash
pig repo rm                      # Remove all repos
pig repo rm pgdg                 # Remove specific repo
pig repo rm pgdg pigsty -u       # Remove and update cache
```

**Backup Location:**
- EL: `/etc/yum.repos.d/backup/`
- Debian: `/etc/apt/sources.list.d/backup/`

------

### `repo update`

Updates the package manager cache to reflect repository changes.

```bash
pig repo update                  # Update package cache
```

**Equivalent Commands:**
- EL Systems: `yum makecache` or `dnf makecache`
- Debian Systems: `apt update`

------

### `repo create`

Creates a local package repository for offline installations or caching.

```bash
pig repo create                  # Create at default location (/www/pigsty)
pig repo create /srv/repo        # Create at custom location
pig repo create /www/pigsty /www/backup  # Multiple locations
```

**Requirements:**
- EL Systems: `createrepo_c` package
- Debian Systems: `dpkg-dev` package

**Example Workflow:**
```bash
# Create local repo structure
sudo pig repo create /www/pigsty

# Copy packages to repo
cp *.rpm /www/pigsty/     # For EL
cp *.deb /www/pigsty/     # For Debian

# Recreate repo metadata
sudo pig repo create /www/pigsty

# Add local repo to system
sudo pig repo add local
```

------

### `repo cache`

Creates a compressed tarball of repository contents for offline distribution.

```bash
pig repo cache                   # Default: /www to /tmp/pkg.tgz
pig repo cache -f                # Force overwrite existing
pig repo cache -d /srv           # Custom source directory
pig repo cache pigsty mssql      # Specific repos only
```

**Options:**
- `-d, --dir`: Source directory (default: `/www/`)
- `-p, --path`: Output path (default: `/tmp/pkg.tgz`)
- `-f`: Force overwrite existing package

**Example Workflow:**
```bash
# On online system: Create offline package
pig repo add -ru                 # Setup repos
pig install pg17 pg_duckdb       # Install packages
pig repo create                  # Create local repo
pig repo cache                   # Create offline package

# Transfer pkg.tgz to offline system

# On offline system: Bootstrap from package
pig repo boot                    # Extract and setup
pig repo add local              # Use local repo
pig install pg17 pg_duckdb      # Install from local
```

------

### `repo boot`

Extracts and sets up a local repository from an offline package.

```bash
pig repo boot                    # Default: /tmp/pkg.tgz to /www
pig repo boot -p /mnt/pkg.tgz   # Custom package path
pig repo boot -d /srv           # Custom target directory
```

**Options:**
- `-p, --path`: Package path (default: `/tmp/pkg.tgz`)
- `-d, --dir`: Target directory (default: `/www/`)

------

## Common Scenarios

### Scenario 1: Fresh PostgreSQL Installation

```bash
# Setup repositories
sudo pig repo add -ru

# Install PostgreSQL 17
sudo pig ext install pg17

# Install popular extensions
sudo pig ext add pg_duckdb postgis timescaledb
```

### Scenario 2: Air-gapped Environment

```bash
# On internet-connected machine:
sudo pig repo add -ru
sudo pig ext install pg17
sudo pig ext add pg_duckdb postgis
sudo pig repo create
sudo pig repo cache

# Transfer /tmp/pkg.tgz to air-gapped machine

# On air-gapped machine:
sudo pig repo boot
sudo pig repo add local
sudo pig ext install pg17
sudo pig ext add pg_duckdb postgis
```

### Scenario 3: Using Regional Mirrors

```bash
# For users in China
sudo pig repo add all --region china -u

# Check mirror URLs
pig repo info pgdg
```

### Scenario 4: Troubleshooting Repository Issues

```bash
# Check current status
pig repo status

# Backup and clean existing repos
sudo pig repo rm

# Add fresh repos
sudo pig repo add all -u

# Verify
pig repo status
```

------

## Advanced Module Reference

### Infrastructure Modules

- **`infra`**: Monitoring and observability stack (Grafana, Prometheus)
- **`docker`**: Docker CE repository
- **`kube`**: Kubernetes repository
- **`grafana`**: Grafana official repository

### Database Modules

- **`mssql`**: WiltonDB (SQL Server compatibility)
- **`mysql`**: MySQL/MariaDB repository
- **`mongo`**: MongoDB repository
- **`redis`**: Redis repository

### PostgreSQL Extension Modules

- **`extra`**: Additional third-party extensions
- **`pgml`**: PostgresML for machine learning
- **`groonga`**: PGroonga full-text search
- **`ivory`**: IvorySQL Oracle compatibility

### Special Modules

- **`local`**: Local repository at `/www/pigsty`
- **`haproxy`**: HAProxy load balancer

------

## Best Practices

1. **Always Update After Adding Repos**: Use the `-u` flag or run `pig repo update` after adding repositories to ensure package lists are current.

2. **Use Module Groups**: Instead of adding individual repos, use module groups like `all` for simpler management.

3. **Backup Before Removal**: The `rm` command automatically backs up repos, but consider manual backups for critical systems.

4. **Check Compatibility**: Use `pig repo info` to verify repository availability for your OS before adding.

5. **Regional Mirrors**: Use `--region` flag for faster downloads if regional mirrors are available.

6. **Offline Preparation**: For air-gapped environments, prepare packages on an internet-connected system first.

------

## Troubleshooting

### Permission Denied

Most repo operations require root/sudo privileges:
```bash
sudo pig repo add all -u
```

### Repository Not Available

Check OS compatibility:
```bash
pig repo info <repo_name>
pig status  # Check detected OS
```

### Package Not Found After Adding Repo

Update package cache:
```bash
sudo pig repo update
```

### Conflicting Repositories

Clean and reset:
```bash
sudo pig repo set all -u
```

### Slow Downloads

Use regional mirrors:
```bash
sudo pig repo add all --region china -u
```