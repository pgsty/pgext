---
title: "CMD: ext"
description: How to manage extensions with pig ext subcommand
icon: SquareTerminal
weight: 620
---

The `pig ext` command is a comprehensive tool for managing PostgreSQL extensions and kernel packages. It provides a unified interface to search, install, remove, update, and manage over 400 PostgreSQL extensions across different PostgreSQL versions and operating systems.

------

## Subcommands

| Command | Description |
|---------|------------|
| `list`   | List & search available extensions |
| `info`   | Get extension information |
| `status` | Show installed extension on active pg |
| `add`    | Install postgres extension |
| `rm`     | Remove postgres extension |
| `update` | Update installed extensions for current pg version |
| `scan`   | Scan installed extensions for active pg |
| `import` | Import extension packages to local repo |
| `link`   | Link postgres to active PATH |
| `reload` | Reload extension catalog to the latest version |

------

## Overview

```bash
pig ext - Manage PostgreSQL Extensions

  Get Started: https://pigsty.io/ext/pig/
  pig repo add -ru             # add all repo and update cache (brute but effective)
  pig ext add pg17             # install optional postgresql 17 package
  pig ext list duck            # search extension in catalog
  pig ext scan -v 17           # scan installed extension for pg 17
  pig ext add pg_duckdb        # install certain postgresql extension

Usage:
  pig ext [command]

Aliases:
  ext, e, ex, pgext, extension

Examples:

  pig ext list    [query]      # list & search extension
  pig ext info    [ext...]     # get information of a specific extension
  pig ext status  [-v]         # show installed extension and pg status
  pig ext add     [ext...]     # install extension for current pg version
  pig ext rm      [ext...]     # remove extension for current pg version
  pig ext update  [ext...]     # update extension to the latest version
  pig ext import  [ext...]     # download extension to local repo
  pig ext link    [ext...]     # link postgres installation to path
  pig ext reload               # reload the latest extension catalog data

Available Commands:
  add         install postgres extension
  import      import extension packages to local repo
  info        get extension information
  link        link postgres to active PATH
  list        list & search available extensions
  reload      reload extension catalog to the latest version
  rm          remove postgres extension
  scan        scan installed extensions for active pg
  status      show installed extension on active pg
  update      update installed extensions for current pg version

Flags:
  -h, --help          help for ext
  -p, --path string   specify a postgres by pg_config path
  -v, --version int   specify a postgres by major version

Global Flags:
      --debug              enable debug mode
  -H, --home string        pigsty home path
  -i, --inventory string   config inventory path
      --log-level string   log level: debug, info, warn, error, fatal, panic (default "info")
      --log-path string    log file path, terminal by default
```

------

## Quick Start

Before managing extensions, ensure repositories are properly configured:

```bash
# Setup repositories (required first time)
sudo pig repo add -ru

# Search for extensions
pig ext list                     # List all available extensions
pig ext list duckdb              # Search for specific extension
pig ext list olap                # List extensions by category

# Install PostgreSQL and extensions
sudo pig ext install pg17        # Install PostgreSQL 17
sudo pig ext add pg_duckdb       # Install DuckDB extension
sudo pig ext add postgis timescaledb pgvector  # Install multiple extensions

# Check installed extensions
pig ext status                   # Show active PostgreSQL and extensions
pig ext scan                     # Scan all installed extensions
```

------

## PostgreSQL Version Detection

The extension manager needs to know which PostgreSQL version to target. This can be specified in three ways:

### 1. Automatic Detection (Default)
```bash
pig ext add pg_duckdb            # Uses pg_config from PATH
```
The tool will search for `pg_config` in your system PATH and use it to detect the active PostgreSQL installation.

### 2. By Major Version (`-v`)
```bash
pig ext add pg_duckdb -v 17      # Install for PostgreSQL 17
pig ext add postgis -v 16        # Install for PostgreSQL 16
```
Uses the default PGDG installation paths:
- EL: `/usr/pgsql-$version/bin/pg_config`
- Debian/Ubuntu: `/usr/lib/postgresql/$version/bin/pg_config`

### 3. By pg_config Path (`-p`)
```bash
pig ext add pg_duckdb -p /opt/pgsql/bin/pg_config
pig ext add postgis -p /usr/local/pgsql/bin/pg_config
```
Explicitly specify the pg_config path for custom installations.

------

## Command Reference

### `ext list`

Lists and searches available PostgreSQL extensions in the catalog.

```bash
pig ext list                     # List all extensions
pig ext list postgis             # Search by name/description
pig ext list olap                # List by category
pig ext list gis -v 16           # List GIS extensions for PG 16
pig ext ls duck                  # Using alias
```

**Output Format:**
```
Name            Version    Description                                     Category
----            -------    -----------                                     --------
pg_duckdb       0.1.0      DuckDB embedded in PostgreSQL                  OLAP
duckdb_fdw      1.1.2      DuckDB Foreign Data Wrapper                    FDW
pgvector        0.7.4      Vector similarity search for PostgreSQL        AI/ML
```

**Categories:**
- `OLAP`: Analytical processing extensions
- `GIS`: Geographic and spatial extensions
- `AI/ML`: Machine learning and AI extensions
- `FDW`: Foreign Data Wrappers
- `STAT`: Statistics and monitoring
- `TIME`: Time-series extensions
- `SEARCH`: Full-text search extensions
- `ADMIN`: Administration tools
- `SEC`: Security extensions
- `SHARD`: Sharding and distribution
- `JSON`: JSON/JSONB utilities

**Search Tips:**
- Search is case-insensitive
- Searches both name and description fields
- Category names can be used as search terms
- Use `-v` flag to filter by PostgreSQL version

------

### `ext info`

Displays detailed information about specific extensions.

```bash
pig ext info pg_duckdb           # Single extension
pig ext info postgis pgvector    # Multiple extensions
```

**Example Output:**
```yaml
Name        : pg_duckdb
Version     : 0.1.0
Category    : OLAP
License     : MIT
Homepage    : https://github.com/duckdb/pg_duckdb
Description : DuckDB embedded in PostgreSQL
Maintainer  : DuckDB Labs
Repository  : pigsty
PostgreSQL  : 12,13,14,15,16,17
Platforms   : el8,el9,d12,u22,u24
Package Names:
  - EL: pg_duckdb_$v
  - DEB: postgresql-$v-pg-duckdb
Dependencies:
  - duckdb (runtime)
  - cmake (build)
```

------

### `ext status`

Shows the current PostgreSQL installation status and installed extensions.

```bash
pig ext status                   # Show active PG and extensions
pig ext status -v 17             # Status for specific PG version
```

**Example Output:**
```
PostgreSQL Status:
  Active Version: 17
  Installation: /usr/pgsql-17
  Config Path: /usr/pgsql-17/bin/pg_config
  Data Directory: /var/lib/pgsql/17/data

Installed Extensions (17):
  - postgis (3.4.2)
  - pgvector (0.7.4)
  - pg_duckdb (0.1.0)
  - timescaledb (2.16.1)
  - pg_cron (1.6.4)

Available PostgreSQL Versions:
  - 14: /usr/pgsql-14
  - 15: /usr/pgsql-15
  - 16: /usr/pgsql-16
  - 17: /usr/pgsql-17 [active]
```

------

### `ext add`

Installs PostgreSQL extensions or kernel packages.

```bash
# Install extensions
sudo pig ext add pg_duckdb                      # Single extension
sudo pig ext add postgis pgvector timescaledb   # Multiple extensions
sudo pig ext add pg_duckdb -y                   # Auto-confirm
sudo pig ext add pg_duckdb -v 16                # For specific PG version

# Install PostgreSQL kernel packages
sudo pig ext install pg17                       # Full PostgreSQL 17
sudo pig ext install pg16-core                  # Core packages only
sudo pig ext install pg15-main                  # Core + essential extensions
sudo pig ext install pg14-devel                 # With development packages
sudo pig ext install pgsql-common               # Common utilities
```

**PostgreSQL Package Sets:**
- `pg$v` or `pgsql`: Latest PostgreSQL version
- `pg$v-core`: Minimal installation (server, client)
- `pg$v-main`: Core + essential extensions (pgvector, pg_repack, wal2json)
- `pg$v-mini`: Minimal viable installation
- `pg$v-devel`: Include development packages
- `pgsql-common`: Utilities (patroni, pgbouncer, pgbackrest, pg_exporter)

**Options:**
- `-y, --yes`: Auto-confirm installation
- `-v, --version`: Target PostgreSQL version
- `-p, --path`: Target PostgreSQL by pg_config path

**Common Extension Sets:**

```bash
# OLAP Stack
sudo pig ext add pg_duckdb parquet_fdw

# GIS Stack
sudo pig ext add postgis pgrouting h3 ogr_fdw

# AI/ML Stack
sudo pig ext add pgvector pgvectorscale pg_similarity pgml

# Time-series Stack
sudo pig ext add timescaledb pg_partman pg_cron

# Monitoring Stack
sudo pig ext add pg_stat_statements pg_stat_kcache pg_qualstats pg_wait_sampling

# Development Stack
sudo pig ext add pgtap plpgsql_check pg_hint_plan
```

------

### `ext rm`

Removes installed PostgreSQL extensions.

```bash
sudo pig ext rm pg_duckdb                       # Remove single extension
sudo pig ext rm postgis pgvector                # Remove multiple
sudo pig ext rm pg_duckdb -v 16                 # From specific PG version
```

**Note:** Some extensions may have dependencies that prevent removal. Use system package manager directly for force removal if needed.

------

### `ext update`

Updates installed extensions to the latest available version.

```bash
sudo pig ext update                             # Update all extensions
sudo pig ext update pg_duckdb                   # Update specific extension
sudo pig ext update postgis pgvector -y         # Auto-confirm updates
```

**Options:**
- `-y, --yes`: Auto-confirm updates
- `-v, --version`: Target PostgreSQL version

------

### `ext scan`

Scans and lists all installed extensions for the active or specified PostgreSQL.

```bash
pig ext scan                     # Scan active PostgreSQL
pig ext scan -v 17               # Scan PostgreSQL 17
pig ext scan -p /opt/pgsql/bin/pg_config  # Scan custom installation
```

**Output includes:**
- Extension name and version
- Installation path
- Control file location
- Shared library status
- Database availability

------

### `ext import`

Downloads extension packages to local repository for offline installation.

```bash
pig ext import pg_duckdb                        # Import single extension
pig ext import postgis pgvector timescaledb     # Import multiple
pig ext import pg_duckdb -v 16,17               # For multiple PG versions
```

**Use Cases:**
- Preparing offline installation packages
- Creating local mirror repositories
- Pre-downloading for air-gapped environments

**Workflow:**
```bash
# On internet-connected machine
pig ext import pg_duckdb postgis pgvector
pig repo create                  # Create local repo
pig repo cache                   # Create offline package

# Transfer to offline machine
# On offline machine
pig repo boot                    # Extract packages
pig repo add local               # Use local repo
pig ext add pg_duckdb postgis pgvector
```

------

### `ext link`

Creates symbolic links to make a specific PostgreSQL version active in system PATH.

```bash
sudo pig ext link 17             # Link PostgreSQL 17
sudo pig ext link pg17           # Alternative syntax
sudo pig ext link -p /usr/pgsql-16/bin/pg_config  # Link by path
```

**What it does:**
1. Creates symlinks in `/usr/bin/` for PostgreSQL binaries
2. Updates alternatives system (if available)
3. Makes the specified version the default `psql`, `pg_dump`, etc.

**Example:**
```bash
# Before linking
which psql                       # /usr/bin/psql -> /etc/alternatives/pgsql-psql

# Link PostgreSQL 17
sudo pig ext link 17

# After linking
which psql                       # /usr/bin/psql -> /usr/pgsql-17/bin/psql
psql --version                   # psql (PostgreSQL) 17.0
```

------

### `ext reload`

Updates the extension catalog to the latest version.

```bash
pig ext reload                   # Reload catalog
```

**What it does:**
- Downloads the latest extension metadata
- Updates local catalog cache
- Refreshes available extension list

**When to use:**
- After major PostgreSQL releases
- When new extensions become available
- If extension metadata seems outdated

------

## Common Workflows

### Workflow 1: Setting Up New PostgreSQL Installation

```bash
# 1. Setup repositories
sudo pig repo add -ru

# 2. Install PostgreSQL 17
sudo pig ext install pg17

# 3. Link to PATH
sudo pig ext link 17

# 4. Install essential extensions
sudo pig ext add pg_stat_statements pg_repack wal2json

# 5. Install application-specific extensions
sudo pig ext add pg_duckdb postgis pgvector timescaledb

# 6. Verify installation
pig ext status
```

### Workflow 2: Migrating Between PostgreSQL Versions

```bash
# Install new version alongside old
sudo pig ext install pg17

# Install same extensions for new version
pig ext scan -v 16 > pg16_extensions.txt
sudo pig ext add $(cat pg16_extensions.txt | awk '{print $1}') -v 17

# Switch active version
sudo pig ext link 17
```

### Workflow 3: Building Development Environment

```bash
# Install multiple PostgreSQL versions
sudo pig ext install pg15-devel pg16-devel pg17-devel

# Install development extensions
sudo pig ext add pgtap plpgsql_check pg_hint_plan -v 17
sudo pig ext add pgbench_extended pg_stat_kcache -v 17

# Install language extensions
sudo pig ext add plpython3 plperl pltcl -v 17
```

### Workflow 4: Setting Up Analytics Stack

```bash
# Install PostgreSQL with analytics focus
sudo pig ext install pg17

# OLAP extensions
sudo pig ext add pg_duckdb parquet_fdw apache_arrow_fdw

# Time-series
sudo pig ext add timescaledb pg_partman pg_cron

# Statistics
sudo pig ext add pg_stat_monitor pg_qualstats pg_wait_sampling

# ML/AI
sudo pig ext add pgvector pgml
```

### Workflow 5: Offline Installation

```bash
# On online machine: Prepare packages
sudo pig repo add -ru
pig ext import pg17 pg_duckdb postgis pgvector
sudo pig repo create /www/pigsty
sudo pig repo cache

# Transfer /tmp/pkg.tgz to offline machine

# On offline machine: Install from cache
sudo pig repo boot
sudo pig repo add local
sudo pig ext install pg17
sudo pig ext add pg_duckdb postgis pgvector
```

------

## Extension Categories Guide

### OLAP & Analytics
Essential for analytical workloads:
- `pg_duckdb`: Embedded DuckDB for OLAP queries
- `parquet_fdw`: Query Parquet files directly
- `citus`: Distributed PostgreSQL
- `columnar`: Columnar storage engine

### GIS & Spatial
Geographic and spatial data processing:
- `postgis`: Spatial database extender
- `pgrouting`: Routing functionality
- `h3`: Uber's H3 hierarchical spatial index
- `ogr_fdw`: Access OGR datasources

### AI & Machine Learning
ML capabilities within PostgreSQL:
- `pgvector`: Vector similarity search
- `pgvectorscale`: Scalable vector search
- `pgml`: PostgresML for machine learning
- `pg_similarity`: Similarity search functions

### Time Series
Time-series data management:
- `timescaledb`: Time-series database
- `pg_partman`: Partition management
- `pg_cron`: Job scheduler

### Full-Text Search
Advanced text search capabilities:
- `pgroonga`: Full-text search using Groonga
- `pg_bigm`: Bi-gram search
- `zhparser`: Chinese text search
- `pg_jieba`: Chinese word segmentation

### Monitoring & Statistics
Performance monitoring and analysis:
- `pg_stat_statements`: Query performance tracking
- `pg_stat_kcache`: Kernel cache statistics
- `pg_stat_monitor`: Advanced monitoring
- `pg_wait_sampling`: Wait event sampling
- `pg_qualstats`: Query quality statistics

### Security
Security and auditing extensions:
- `pgcrypto`: Cryptographic functions
- `pg_audit`: Audit logging
- `set_user`: Privilege escalation control
- `supautils`: Supabase utilities

### Development Tools
Development and debugging:
- `pgtap`: Unit testing framework
- `plpgsql_check`: PL/pgSQL code checker
- `pg_hint_plan`: Query plan hints
- `pgbench`: Benchmarking tool

------

## Troubleshooting

### Extension Not Found

```bash
# Update catalog
pig ext reload

# Search with different terms
pig ext list <partial_name>

# Check available repos
pig repo status
```

### Version Compatibility Issues

```bash
# Check extension compatibility
pig ext info <extension_name>

# Specify correct PostgreSQL version
sudo pig ext add <extension> -v <version>
```

### Missing Dependencies

```bash
# Some extensions require OS packages
sudo dnf install <required_package>  # For EL
sudo apt install <required_package>  # For Debian/Ubuntu

# Then retry extension installation
sudo pig ext add <extension>
```

### pg_config Not Found

```bash
# Install PostgreSQL first
sudo pig ext install pg17

# Or specify path explicitly
pig ext add <extension> -p /usr/pgsql-17/bin/pg_config

# Or link PostgreSQL to PATH
sudo pig ext link 17
```

### Permission Denied

```bash
# Installation requires sudo
sudo pig ext add <extension>

# For system-wide operations
sudo pig ext link 17
```

------

## Best Practices

1. **Always Setup Repos First**: Run `pig repo add -ru` before installing extensions

2. **Use Version Flags**: Explicitly specify PostgreSQL version when managing multiple installations

3. **Check Compatibility**: Use `pig ext info` to verify extension supports your PostgreSQL version

4. **Group Related Extensions**: Install related extensions together for better dependency resolution

5. **Regular Updates**: Periodically run `pig ext update` to get security and bug fixes

6. **Catalog Maintenance**: Run `pig ext reload` occasionally to get new extension listings

7. **Backup Before Major Changes**: Create database backups before updating critical extensions

8. **Test in Development**: Always test extension updates in development before production

9. **Monitor After Installation**: Check logs after installing new extensions for any issues

10. **Document Your Stack**: Keep a record of installed extensions for disaster recovery