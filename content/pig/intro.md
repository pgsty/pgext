---
title: Introduction
description: Why would we need yet another package manager? especially for Postgres extensions?
weight: 100
icon: fa-solid fa-circle-question
breadcrumb: false
---

PIG is a command-line tool for installing and managing PostgreSQL and its extensions. It maps extension names to RPM or DEB packages and uses `apt`, `dnf`, or `yum` to install them.

The catalog contains [{{< param pig_catalog_count >}} extension entries](https://pigsty.io/ext/list/). PIG selects package names for your PostgreSQL major version, operating system, and CPU architecture. It also provides commands for source builds, local PostgreSQL operations, backups, and Pigsty administration.

## What you can do

| Task | Command | Reference |
|:---|:---|:---|
| Find an extension and check package availability | `pig ext list`, `pig ext avail` | [Extensions](/pig/cmd/ext/) |
| Add PGDG and Pigsty repositories | `pig repo add pgsql -u` | [Repositories](/pig/cmd/repo/) |
| Install PostgreSQL 18 and pgvector | `pig install -v 18 pg18 vector` | [Getting started](/pig/start/) |
| Build an extension package from source | `pig build pkg NAME` | [Build](/pig/cmd/build/) |
| Start, inspect, or stop a local PostgreSQL instance | `pig pg` | [PostgreSQL](/pig/cmd/pg/) |
| Manage a Patroni cluster | `pig pt` | [Patroni](/pig/cmd/pt/) |
| Back up or restore an instance | `pig pb`, `pig pitr` | [pgBackRest](/pig/cmd/pb/), [PITR](/pig/cmd/pitr/) |
| Configure Pigsty and run its playbooks | `pig sty`, `pig inventory`, `pig do` | [Pigsty](/pig/cmd/sty/) |

Start with the [tutorial](/pig/start/) to install PostgreSQL and enable your first extension.

## How package installation works

For example, this command installs pgvector for PostgreSQL 18:

```bash
sudo pig install vector -v 18
```

PIG resolves `vector` to `pgvector_18` on EL or `postgresql-18-pgvector` on Debian and Ubuntu. The system package manager resolves dependencies and installs the files.

To enable the extension, connect to the intended database and run:

```sql
CREATE EXTENSION vector;
```

Some extensions also need a setting in `shared_preload_libraries` and a PostgreSQL restart. Use `pig ext info NAME` to check these requirements. See [package and database operations](/pig/cmd/ext/#package-layer-vs-database-layer) for updates and removal.

PIG is needed only when you run its commands. PostgreSQL and installed extensions continue to work without it. You can also use the [Pigsty repositories](https://pigsty.io/docs/repo/pgsql/) directly with `apt` or `dnf`.

## Linux compatibility

The package repositories cover PostgreSQL **14–18** on these Linux releases. Each release has packages for **x86_64** and **aarch64**.

| OS code | Distribution | Release |
|:---|:---|:---|
| `el8` | RHEL, Rocky Linux, AlmaLinux | 8 |
| `el9` | RHEL, Rocky Linux, AlmaLinux | 9 |
| `el10` | RHEL, Rocky Linux, AlmaLinux | 10 |
| `d12` | Debian | 12 |
| `d13` | Debian | 13 |
| `u22` | Ubuntu | 22.04 LTS |
| `u24` | Ubuntu | 24.04 LTS |
| `u26` | Ubuntu | 26.04 LTS |

Package availability varies by extension and target. On the target host, run:

```bash
pig ext avail vector
```

The {{< param pig_catalog_count >}} catalog entries include packaged SQL utilities and retained records for removed extensions. The website and CLI use this same [catalog-entry count](/pig/cmd/ext/#catalog-count). Check the entry's status and availability before selecting a package. EL 7 and Debian 11 are outside the current package matrix.

The macOS binary supports catalog queries and commands such as inventory editing. PostgreSQL package installation uses Linux package managers. See [Installation](/pig/install/) for the binary downloads.

## Choose a workflow

- **Use an existing host:** add repositories with `pig repo add pgsql -u`, then install the required packages.
- **Configure all repositories on a new host:** use `pig repo set`. It backs up and replaces the existing repository definitions.
- **Build a package:** prepare a build host and use the maintained specifications with [`pig build`](/pig/cmd/build/).
- **Prepare an offline installation:** collect packages with `pig ext import`, index them with `pig repo create`, and archive them with `pig repo cache`. See [offline repositories](/pig/cmd/repo/).
- **Write a script or agent workflow:** inspect `pig -o json --help` and use the supported output and plan options in the [command reference](/pig/cmd/).

The [Design Records](https://pig.pgsty.com/design/) explain the main implementation decisions. The [release notes](/release/pig/) describe changes in each published version.
