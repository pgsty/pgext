---
title: "CMD: sty"
description: How to manage Pigsty installation with pig sty subcommand
icon: SquareTerminal
weight: 630
---

The `pig sty` command manages Pigsty installation - a battery-included PostgreSQL distribution that brings High Availability, Point-In-Time Recovery, comprehensive monitoring, Infrastructure as Code, and 400+ extensions to your PostgreSQL clusters.

------

## Subcommands

| Command | Description |
|---------|------------|
| `init`    | Install Pigsty (download & extract) |
| `boot`    | Bootstrap Pigsty (install Ansible & dependencies) |
| `conf`    | Configure Pigsty (generate config file) |
| `install` | Run pigsty install.yml playbook |
| `get`     | Download pigsty source tarball |
| `list`    | List available pigsty versions |

------

## Overview

```bash
pig sty - Init (Download), Bootstrap, Configure, and Install Pigsty

  pig sty init    [-pfvd]      # install pigsty (~/pigsty by default)
  pig sty boot    [-rpk]       # install ansible and prepare offline pkg
  pig sty conf    [-civrsxn]   # configure pigsty and generate config
  pig sty install              # use pigsty to install & provisioning env (DANGEROUS!)
  pig sty get                  # download pigsty source tarball
  pig sty list                 # list available pigsty versions

Usage:
  pig sty [command]

Aliases:
  sty, s, pigsty

Examples:
  Get Started: https://pigsty.io/docs/setup/install/
  pig sty init                 # extract and init ~/pigsty
  pig sty boot                 # install ansible & other deps
  pig sty conf                 # generate pigsty.yml config file
  pig sty install              # run pigsty/install.yml playbook

Available Commands:
  boot        Bootstrap Pigsty
  conf        Configure Pigsty
  get         download pigsty available versions
  init        Install Pigsty
  install     run pigsty install.yml playbook
  list        list pigsty available versions

Flags:
  -h, --help   help for sty

Global Flags:
      --debug              enable debug mode
  -H, --home string        pigsty home path
  -i, --inventory string   config inventory path
      --log-level string   log level: debug, info, warn, error, fatal, panic (default "info")
      --log-path string    log file path, terminal by default
```

------

## Quick Start

The complete Pigsty installation workflow:

```bash
# Step 1: Install Pigsty to ~/pigsty
pig sty init

# Step 2: Bootstrap with Ansible
cd ~/pigsty
pig sty boot

# Step 3: Generate configuration
pig sty conf

# Step 4: Run installation playbook
pig sty install
```

After completion, you'll have a fully configured PostgreSQL environment with:
- PostgreSQL cluster with high availability
- Monitoring stack (Grafana, Prometheus, Loki)
- Backup and recovery tools
- Web console at http://localhost:3000

------

## What is Pigsty?

Pigsty is a comprehensive PostgreSQL distribution that provides:

### Core Features
- **High Availability**: Patroni-managed automatic failover
- **Point-In-Time Recovery**: pgBackRest for backups and PITR
- **Monitoring Stack**: Grafana dashboards, Prometheus metrics, Loki logs
- **Infrastructure as Code**: Ansible-based deployment and configuration
- **Extension Library**: 400+ pre-compiled PostgreSQL extensions

### Components
- **PostgreSQL**: Core database with multiple major versions
- **Patroni**: HA orchestrator for automatic failover
- **HAProxy**: Load balancer for connection pooling
- **pgBouncer**: Connection pooler
- **pgBackRest**: Backup and recovery solution
- **Monitoring**: Grafana, Prometheus, Loki, AlertManager
- **Management**: Ansible playbooks for all operations

------

## Command Reference

### `sty init`

Downloads and installs Pigsty to your system.

```bash
pig sty init                     # Install to default location ~/pigsty
pig sty init -p /opt/pigsty      # Custom installation path
pig sty init -v 3.0.4            # Specific version
pig sty init -d                  # Download only, don't extract
pig sty init -f                  # Force overwrite existing installation
```

**Options:**
- `-p, --path`: Installation path (default: `~/pigsty`)
- `-v, --version`: Pigsty version to install
- `-d, --download`: Download only without extraction
- `-f, --force`: Force overwrite existing installation

**What it does:**
1. Downloads Pigsty source tarball
2. Extracts to specified directory
3. Sets up initial directory structure
4. Prepares for bootstrap phase

**Example:**
```bash
# Fresh installation
pig sty init

# Check installation
cd ~/pigsty
ls -la
```

------

### `sty boot`

Bootstraps Pigsty by installing Ansible and system dependencies.

```bash
pig sty boot                     # Bootstrap with online packages
pig sty boot -r                  # Also setup repositories
pig sty boot -p                  # Use offline package if available
pig sty boot -k                  # Keep packages after installation
```

**Options:**
- `-r, --repo`: Setup system repositories
- `-p, --pkg`: Use offline package (`/tmp/pkg.tgz`)
- `-k, --keep`: Keep downloaded packages

**What it installs:**
- Ansible and dependencies
- Python modules
- System utilities
- PostgreSQL client tools

**Prerequisites:**
- Pigsty must be initialized first (`pig sty init`)
- Root or sudo access for system packages

**Example:**
```bash
# Standard bootstrap
cd ~/pigsty
pig sty boot

# With offline package
pig sty boot -p

# Verify Ansible installation
ansible --version
```

------

### `sty conf`

Generates Pigsty configuration file based on system detection.

```bash
pig sty conf                     # Generate default config
pig sty conf -c                  # Clean mode (minimal config)
pig sty conf -i 192.168.1.10    # Specify IP address
pig sty conf -v 17               # PostgreSQL version
pig sty conf -r                  # Include region settings
pig sty conf -s                  # Single node setup
pig sty conf -x                  # Expert mode (all parameters)
pig sty conf -n mycluster       # Custom cluster name
```

**Options:**
- `-c, --clean`: Generate minimal configuration
- `-i, --ip`: Primary IP address
- `-v, --version`: PostgreSQL major version (default: 17)
- `-r, --region`: Include region-specific settings
- `-s, --single`: Single-node configuration
- `-x, --expert`: Include all configuration parameters
- `-n, --name`: Cluster name

**Generated file:** `~/pigsty/pigsty.yml`

**Configuration includes:**
- Cluster topology
- PostgreSQL settings
- Monitoring configuration
- Network parameters
- Extension list

**Example:**
```bash
# Generate config for single node
cd ~/pigsty
pig sty conf -s

# Custom cluster configuration
pig sty conf -n prod-cluster -v 16 -i 10.10.10.10

# Review generated config
cat pigsty.yml
```

------

### `sty install`

Runs the Pigsty installation playbook to provision the environment.

```bash
pig sty install                  # Run complete installation
```

**⚠️ WARNING**: This command will make significant changes to your system:
- Installs PostgreSQL clusters
- Configures monitoring stack
- Sets up network services
- Modifies system settings

**Prerequisites:**
- Pigsty bootstrapped (`pig sty boot`)
- Configuration generated (`pig sty conf`)
- Sufficient system resources
- Root or sudo access

**What it does:**
1. Validates configuration
2. Installs infrastructure components
3. Deploys PostgreSQL clusters
4. Sets up monitoring stack
5. Configures backup solutions
6. Initializes web console

**Access after installation:**
- Grafana: http://localhost:3000 (admin/pigsty)
- Prometheus: http://localhost:9090
- AlertManager: http://localhost:9093
- PostgreSQL: localhost:5432

**Example:**
```bash
cd ~/pigsty
pig sty install

# Check service status
systemctl status patroni
systemctl status grafana-server

# Access web console
open http://localhost:3000
```

------

### `sty get`

Downloads Pigsty source tarball without installation.

```bash
pig sty get                      # Download latest version
pig sty get -v 3.0.4             # Specific version
pig sty get -o /tmp/pigsty.tgz  # Custom output path
```

**Options:**
- `-v, --version`: Version to download
- `-o, --output`: Output file path

**Example:**
```bash
# Download latest
pig sty get

# Download specific version
pig sty get -v 3.0.3

# Check downloaded file
ls -lh pigsty*.tgz
```

------

### `sty list`

Lists all available Pigsty versions.

```bash
pig sty list                     # List all versions
```

**Output includes:**
- Version number
- Release date
- Download size
- Major features

**Example:**
```bash
pig sty list

# Output:
# Available Pigsty Versions:
# - 3.0.4 (latest) - 2024-10-28
# - 3.0.3 - 2024-09-15
# - 3.0.2 - 2024-08-30
# ...
```

------

## Common Workflows

### Workflow 1: Single-Node Installation

```bash
# Complete single-node setup
pig sty init
cd ~/pigsty
pig sty boot
pig sty conf -s
pig sty install

# Access services
open http://localhost:3000
```

### Workflow 2: Production Cluster Setup

```bash
# 1. Initialize on admin node
pig sty init
cd ~/pigsty

# 2. Bootstrap
pig sty boot -r

# 3. Configure for production
pig sty conf -x -n prod-db

# 4. Edit configuration
vim pigsty.yml
# Add cluster nodes, adjust parameters

# 5. Deploy
pig sty install
```

### Workflow 3: Offline Installation

```bash
# On online machine:
pig sty get
pig repo cache
# Transfer pigsty.tgz and pkg.tgz to offline machine

# On offline machine:
pig sty init -p ~/pigsty
cd ~/pigsty
pig sty boot -p
pig sty conf
pig sty install
```

### Workflow 4: Development Environment

```bash
# Quick dev setup
pig sty init
cd ~/pigsty
pig sty boot
pig sty conf -c -s  # Minimal single-node config
pig sty install

# Install additional extensions
pig ext add pg_duckdb postgis pgvector
```

------

## Configuration Examples

### Single Node Configuration

```yaml
# Generated by: pig sty conf -s
all:
  children:
    infra:
      hosts:
        10.10.10.10: { infra_seq: 1 }
    etcd:
      hosts:
        10.10.10.10: { etcd_seq: 1 }
    pg-meta:
      hosts:
        10.10.10.10: { pg_seq: 1, pg_role: primary }
      vars:
        pg_cluster: pg-meta
        pg_databases:
          - { name: meta }
        pg_users:
          - { name: dbuser_meta, password: DBUser.Meta }
```

### Multi-Node HA Cluster

```yaml
# Three-node HA cluster
all:
  children:
    pg-prod:
      hosts:
        10.10.10.11: { pg_seq: 1, pg_role: primary }
        10.10.10.12: { pg_seq: 2, pg_role: replica }
        10.10.10.13: { pg_seq: 3, pg_role: replica }
      vars:
        pg_cluster: pg-prod
        pg_version: 17
        pg_vip_enabled: true
        pg_vip_address: 10.10.10.10/24
        pg_vip_interface: eth0
```

------

## Post-Installation Tasks

### Accessing Services

```bash
# Grafana dashboard
open http://localhost:3000
# Default: admin/pigsty

# PostgreSQL
psql -h localhost -U postgres

# Check cluster status
patronictl -c /etc/patroni/pg-meta.yml list
```

### Managing Clusters

```bash
# Add new database
cd ~/pigsty
./pgsql-db.yml -l pg-meta

# Add new user
./pgsql-user.yml -l pg-meta

# Create backup
pg_basebackup -h localhost -U replicator -D /tmp/backup
```

### Monitoring

```bash
# View metrics
open http://localhost:3000/d/pgsql-overview

# Check alerts
open http://localhost:9093

# View logs
open http://localhost:3000/d/pgsql-logs
```

------

## Troubleshooting

### Installation Fails

```bash
# Check prerequisites
cd ~/pigsty
./configure -n

# Verify Ansible
ansible --version

# Check logs
tail -f /tmp/pigsty-install.log
```

### Connection Issues

```bash
# Check services
systemctl status postgresql
systemctl status patroni
systemctl status haproxy

# Test connectivity
pg_isready -h localhost
```

### Configuration Problems

```bash
# Validate configuration
cd ~/pigsty
ansible-playbook install.yml --check

# Regenerate config
pig sty conf -x
```

### Service Recovery

```bash
# Restart Patroni cluster
systemctl restart patroni

# Reload HAProxy
systemctl reload haproxy

# Restart monitoring
systemctl restart prometheus grafana-server
```

------

## Best Practices

1. **Test First**: Always test in development before production deployment

2. **Backup Configuration**: Keep copies of `pigsty.yml` and customizations

3. **Monitor Resources**: Ensure adequate CPU, memory, and disk space

4. **Use Version Control**: Track configuration changes in Git

5. **Regular Updates**: Keep Pigsty and extensions updated

6. **Security First**: Change default passwords immediately after installation

7. **Plan Topology**: Design cluster topology before deployment

8. **Document Changes**: Keep records of configuration modifications

9. **Practice Recovery**: Test backup and recovery procedures regularly

10. **Join Community**: Participate in Pigsty community for support

------

## Resources

- **Documentation**: https://pigsty.io/docs/
- **GitHub**: https://github.com/Vonng/pigsty
- **Demos**: https://demo.pigsty.cc
- **Community**: https://github.com/Vonng/pigsty/discussions