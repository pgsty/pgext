---
title: "CMD Docs"
description: Overview of the pig CLI command reference
icon: SquareTerminal
weight: 600
---

The `pig` CLI provides comprehensive tools for managing PostgreSQL installations, extensions, repositories, and building extensions from source. Check command documentation with `pig help <command>`.

{{< cards cols="4" >}}
{{< card link="/pig/cmd/repo"  title="pig repo"  subtitle="Manage software repositories" icon="library" >}}
{{< card link="/pig/cmd/ext"   title="pig ext"   subtitle="Manage postgres extensions"   icon="cube" >}}
{{< card link="/pig/cmd/build" title="pig build" subtitle="Build extension from source"  icon="view-grid" >}}
{{< card link="/pig/cmd/sty"   title="pig sty"   subtitle="Manage pigsty installation"   icon="cloud-download" >}}
{{< /cards >}}

{{< cards cols="4" >}}
{{< card link="/pig/cmd/pg"    title="pig pg"    subtitle="Manage local PostgreSQL"      icon="database" >}}
{{< card link="/pig/cmd/pt"    title="pig pt"    subtitle="Manage Patroni HA cluster"    icon="refresh" >}}
{{< card link="/pig/cmd/pb"    title="pig pb"    subtitle="Manage pgBackRest backup"     icon="archive" >}}
{{< card link="/pig/cmd/pitr"  title="pig pitr"  subtitle="Orchestrated PITR recovery"   icon="clock" >}}
{{< /cards >}}

## Overview

```bash
pig - the Linux Package Manager for PostgreSQL

Usage:
  pig [command]

Examples:

  pig repo add -ru            # overwrite existing repo & update cache
  pig install pg17            # install postgresql 17 PGDG package
  pig install pg_duckdb       # install certain postgresql extension
  pig install pgactive -v 18  # install extension for specifc pg major

  check https://pigsty.io/docs/pig for details

PostgreSQL Extension Manager
  ext         Manage PostgreSQL Extensions (pgext)
  repo        Manage Linux Software Repo (apt/dnf)
  build       Build Postgres Extension

PostgreSQL Management Commands
  pg          Manage local PostgreSQL server
  pt          Manage Patroni HA cluster
  pb          Manage pgBackRest backup & recovery
  pitr        Orchestrated Point-In-Time Recovery

Pigsty Management Commands
  do          Run admin tasks
  sty         Manage Pigsty installation

Additional Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  status      Show Environment Status
  update      Upgrade pig itself
  version     Show pig version info

Flags:
      --debug              enable debug mode
  -h, --help               help for pig
  -H, --home string        pigsty home path
  -i, --inventory string   config inventory path
      --log-level string   log level: debug, info, warn, error, fatal, panic (default "info")
      --log-path string    log file path, terminal by default

Use "pig [command] --help" for more information about a command.
```

------

## Quick Start Guide

### 1. Setup Repositories

Before installing PostgreSQL or extensions, you need to configure package repositories:

```bash
# Quick setup - add all essential repos (pgdg + pigsty + node)
sudo pig repo add -ru

# Alternative: add specific repos
sudo pig repo add pgdg pigsty -u
```

### 2. Install PostgreSQL

```bash
# Install PostgreSQL 17
sudo pig ext install pg17

# Link to system PATH
sudo pig ext link 17
```

### 3. Install Extensions

```bash
# Search for extensions
pig ext list                    # List all
pig ext list duckdb             # Search specific

# Install extensions
sudo pig ext add pg_duckdb postgis pgvector
```

### 4. Build Extensions (Optional)

```bash
# Setup build environment
pig build spec

# Build an extension
pig build pkg citus
```

------

## Core Commands

### [`pig repo`](/pig/cmd/repo/) - Repository Management

Manage APT/YUM repositories for PostgreSQL packages:

```bash
pig repo list                    # List available repositories
pig repo info   pgdg             # Show repository details
pig repo status                  # Check current repo status
pig repo add    pgdg pigsty -u   # Add repositories
pig repo rm     old-repo         # Remove repositories
pig repo update                  # Update package cache
pig repo create /www/pigsty      # Create local repository
pig repo cache                   # Create offline package
pig repo boot                    # Bootstrap from offline package
```

**Key Features:**
- Support for both RPM (EL) and DEB (Debian/Ubuntu) systems
- Regional mirrors (China, Europe, etc.)
- Offline package support
- Module-based repository organization

------

### [`pig ext`](/pig/cmd/ext/) - Extension Management

Manage PostgreSQL extensions and kernel packages:

```bash
pig ext list    duck             # Search extensions
pig ext info    pg_duckdb        # Extension details
pig ext status                   # Show installed extensions
pig ext add     pg_duckdb -y     # Install extension
pig ext rm      old_extension    # Remove extension
pig ext update                   # Update extensions
pig ext scan                     # Scan installed extensions
pig ext import  pg_duckdb        # Download for offline use
pig ext link    17               # Link PG version to PATH
pig ext reload                   # Refresh extension catalog
```

**Key Features:**
- 451 PostgreSQL extensions
- Multi-version PostgreSQL support
- Automatic dependency resolution
- Category-based browsing
- Offline installation support

------

### [`pig build`](/pig/cmd/build/) - Build Extensions

Build PostgreSQL extensions from source:

```bash
# Environment setup
pig build spec                   # Initialize build specs
pig build repo                   # Setup repositories
pig build tool                   # Install build tools
pig build rust -y                # Install Rust (for Rust extensions)
pig build pgrx                   # Install PGRX framework

# Build extensions
pig build pkg citus              # Complete build pipeline
pig build get citus              # Download source
pig build dep citus              # Install dependencies
pig build ext citus              # Build package
```

**Key Features:**
- Support for 100+ extensions
- RPM and DEB package generation
- Rust/PGRX extension support
- Multi-PostgreSQL version builds
- Custom extension support

------

### [`pig sty`](/pig/cmd/sty/) - Pigsty Management

Manage Pigsty PostgreSQL distribution:

```bash
pig sty init                     # Install Pigsty to ~/pigsty
pig sty boot                     # Install Ansible prerequisites
pig sty conf                     # Generate configuration
pig sty deploy                   # Run deployment playbook
```

**Pigsty Features:**
- Battery-included PostgreSQL distribution
- High Availability with Patroni
- Point-In-Time Recovery (PITR)
- Comprehensive monitoring stack
- Infrastructure as Code (IaC)

------

### [`pig pg`](/pig/cmd/pg/) - PostgreSQL Management

Manage local PostgreSQL server and databases:

```bash
# Service control
pig pg init                      # Initialize data directory
pig pg start                     # Start PostgreSQL
pig pg stop                      # Stop PostgreSQL
pig pg status                    # Check status

# Connection & query
pig pg psql mydb                 # Connect to database
pig pg ps                        # Show current connections
pig pg kill -x                   # Terminate connections

# Maintenance
pig pg vacuum mydb               # Vacuum database
pig pg analyze mydb              # Analyze database
pig pg log tail                  # Real-time log viewing
```

**Key Features:**
- Service control (pg_ctl/systemctl wrapper)
- Connection management
- Database maintenance utilities
- Log viewing tools

------

### [`pig pt`](/pig/cmd/pt/) - Patroni Management

Manage Patroni HA cluster:

```bash
pig pt list                      # List cluster members
pig pt config                    # Show cluster config
pig pt config ttl=60             # Modify cluster config
pig pt status                    # Check service status
pig pt restart                   # Restart PostgreSQL via Patroni
pig pt switchover                # Perform planned switchover
pig pt log -f                    # Real-time log viewing
```

**Key Features:**
- Cluster member management
- Configuration management
- Service control
- Switchover/failover operations

------

### [`pig pb`](/pig/cmd/pb/) - pgBackRest Management

Manage pgBackRest backup and PITR:

```bash
pig pb info                      # Show backup info
pig pb ls                        # List all backups
pig pb backup                    # Create backup
pig pb backup full               # Full backup
pig pb restore                   # Restore to latest
pig pb restore -t "2025-01-01"   # Restore to specific time
pig pb log tail                  # Real-time log viewing
```

**Key Features:**
- Backup management
- Point-in-time recovery
- Stanza management
- Multi-repository support

------

### [`pig pitr`](/pig/cmd/pitr/) - Orchestrated PITR

Orchestrated Point-In-Time Recovery that coordinates Patroni, PostgreSQL, and pgBackRest:

```bash
pig pitr -d                      # Restore to latest (end of WAL stream)
pig pitr -t "2025-01-01 12:00"   # Restore to specific time
pig pitr -I                      # Restore to backup consistency point
pig pitr -d --dry-run            # Show execution plan
pig pitr -d -y                   # Skip confirmation (automation)
pig pitr -d --skip-patroni       # Standalone PostgreSQL (non-Patroni)
```

**Key Features:**
- Automatic Patroni/PostgreSQL stop
- Progressive shutdown strategy
- Automatic PostgreSQL restart
- Post-restore guidance

------

## Common Workflows

### Installing PostgreSQL with Extensions

```bash
# 1. Setup repositories
sudo pig repo add -ru

# 2. Install PostgreSQL 17
sudo pig ext install pg17
sudo pig ext link 17

# 3. Install popular extensions
sudo pig ext add pg_duckdb postgis pgvector timescaledb

# 4. Verify installation
pig ext status
```

### Building Custom Extensions

```bash
# 1. Setup build environment
pig build spec
pig build tool

# 2. Build extension
pig build pkg my_extension

# 3. Install built package
sudo rpm -ivh ~/rpmbuild/RPMS/x86_64/my_extension*.rpm  # EL
sudo dpkg -i ~/my_extension*.deb                         # Debian
```

### Offline Installation

```bash
# On online machine:
sudo pig repo add -ru
pig ext import pg17 pg_duckdb postgis
pig repo create
pig repo cache

# Transfer /tmp/pkg.tgz to offline machine

# On offline machine:
pig repo boot
sudo pig repo add local
sudo pig ext install pg17
sudo pig ext add pg_duckdb postgis
```

------

## Environment Detection

The `pig` tool automatically detects your environment:

```bash
pig status                       # Show comprehensive status
```

**Detected Information:**
- Operating System (EL/Debian/Ubuntu)
- OS Version and Architecture
- PostgreSQL installations
- Active PostgreSQL version
- Installed extensions
- Repository configuration

------

## Best Practices

1. **Always Update Repos First**: Run `pig repo add -ru` before installing
2. **Use Version Flags**: Specify PostgreSQL version with `-v` when needed
3. **Check Compatibility**: Use `pig ext info` before installation
4. **Backup Before Changes**: Create backups before major updates
5. **Test in Development**: Test extensions in dev before production
6. **Keep Catalog Updated**: Run `pig ext reload` periodically
7. **Document Your Stack**: Keep records of installed extensions
8. **Use Offline Packages**: Prepare offline packages for air-gapped systems
9. **Build When Needed**: Use build commands for custom or newer versions
10. **Monitor After Changes**: Check logs after installations

------

## Getting Help

```bash
# General help
pig help

# Command-specific help
pig repo --help
pig ext --help
pig build --help

# Subcommand help
pig ext add --help
pig repo info --help
pig build pkg --help
```

For issues and feedback, visit: https://github.com/pgsty/pig/issues
