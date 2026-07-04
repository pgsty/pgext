---
title: "CMD: ext"
description: How to manage extensions with pig ext subcommand
icon: SquareTerminal
weight: 620
---

The `pig ext` command is an all-in-one tool for managing PostgreSQL extensions. It lets you search, install, remove, update, and manage PostgreSQL extensions, and it can also handle PostgreSQL kernel packages.

```bash
pig ext - Manage PostgreSQL Extensions

  pig repo add -ru             # add all repo and update cache (brute but effective)
  pig ext add pg18             # install optional postgresql 18 package
  pig ext list duck            # search extension in catalog
  pig ext scan -v 18           # scan installed extension for pg 18
  pig ext add pg_duckdb        # install certain postgresql extension

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
```

| Command | Description | Notes |
|:---|:---|:---|
| `ext list` | Search extensions | |
| `ext info` | Show extension details | |
| `ext avail` | Show extension availability matrix | |
| `ext status` | Show installed extensions | |
| `ext scan` | Scan installed extensions | |
| `ext add` | Install extensions | Requires sudo or root |
| `ext rm` | Remove extensions | Requires sudo or root |
| `ext update` | Update extensions | Requires sudo or root |
| `ext import` | Download extensions for offline use | Requires sudo or root |
| `ext link` | Link a PG version into PATH | Requires sudo or root |
| `ext reload` | Refresh extension catalog | |
{.full-width}


## Quick Start

Before installing PostgreSQL extensions, add the required repositories with [`pig repo add`](/pig/cmd/repo/):

```bash
pig repo add pgdg pigsty -u    # gentle way to add pgdg and pigsty repos
pig repo set                   # brute-force way to remove and add all required repos
```

Then search and install PostgreSQL extensions:

```bash
pig ext install pg_duckdb
pig ext install pg_partman
pig ext install pg_cron
pig ext install pg_repack
pig ext install pg_stat_statements
pig ext install pg_stat_kcache
```

See the [**extension list**](/list/) for available extensions and package names.

**Notes:**

1. If no PostgreSQL version is specified, pig tries to detect the active PostgreSQL installation from `pg_config` in `PATH`.
2. PostgreSQL can be selected by major version (`-v`) or by pg_config path (`-p`).
   - With `-v`, pig uses the default PGDG kernel package path for that major version.
     - EL distros: `/usr/pgsql-$v/bin/pg_config`
     - DEB distros: `/usr/lib/postgresql/$v/bin/pg_config`
   - With `-p`, pig locates PostgreSQL directly from that path.
3. The extension manager automatically adapts to the OS package format:
   - RPM packages for RHEL/CentOS/Rocky Linux/AlmaLinux
   - DEB packages for Debian/Ubuntu
4. Extension dependencies are resolved automatically when possible.
5. Use `-y` carefully because it auto-confirms prompts.

Pigsty assumes official PGDG kernel packages are installed. If not, install them with:

```bash
pig ext install pg18          # install PostgreSQL 18 kernel packages except devel
```


## ext list

List or search extensions in the extension catalog.

```bash
pig ext list                     # list all extensions
pig ext list duck                # search extensions containing "duck"
pig ext list -v 18               # filter by PG version
pig ext ls olap                  # list OLAP category extensions
pig ext ls gis -v 16             # list GIS extensions for PG 16
pig ext ls rag                   # list RAG category extensions
```

Category filtering is done by passing the category name as the query. Supported categories include: `time`, `gis`, `rag`, `fts`, `olap`, `feat`, `lang`, `type`, `func`, `util`, `admin`, `stat`, `sec`, `fdw`, `sim`, and `etl`.

**Options:**

- `-v|--version`: filter by PG version
- `--pkg`: show package names instead of extension names, listing only leading extensions

**Status column:**

- `installed`: extension is installed
- `available`: extension is available but not installed
- `not avail`: extension is not available on the current system

The default extension catalog is defined in [**`cli/ext/assets/extension.csv`**](https://github.com/pgsty/pig/blob/main/cli/ext/assets/extension.csv).

Use `pig ext reload` to refresh to the latest catalog. The downloaded catalog is stored at `~/.pig/extension.csv`; the latest online catalog is also published at [**pigsty.io/ext/data/extension.csv**](https://pigsty.io/ext/data/extension.csv).


## ext info

Show detailed information for selected extensions.

```bash
pig ext info postgis        # show PostGIS details
pig ext info timescaledb    # show TimescaleDB details
pig ext info vector postgis # show multiple extensions
```


## ext avail

Show the extension availability matrix across operating systems, architectures, and PostgreSQL versions.

```bash
pig ext avail                     # show package availability on current system
pig ext avail timescaledb         # show timescaledb availability matrix
pig ext avail postgis pg_duckdb   # show multiple extensions
pig ext av pgvector               # show pgvector availability
pig ext matrix citus              # alias for avail
```

The matrix shows availability across operating systems (EL8/9/10, Debian 12/13, Ubuntu 22/24/26), architectures (x86_64/aarch64), and PostgreSQL versions (14-18).


## ext status

Show installed extension status for the current PostgreSQL instance.

```bash
pig ext status              # show installed extensions
pig ext status -c           # include contrib extensions
pig ext status -v 16        # show installed extensions for PG 16
```

**Options:**

- `-c|--contrib`: include contrib extensions in the result


## ext scan

Scan installed extensions for the current PostgreSQL instance.

```bash
pig ext scan [-v version]
```

This command scans the PostgreSQL extension directory and finds extensions that are actually installed.


## ext add

Install one or more PostgreSQL extensions. Same-level aliases for `pig ext add` include `pig ext install`, `pig ext ins`, and `pig ext a`. The top-level [`pig install`](/pig/cmd/#install) command is a separate native package-manager wrapper that also supports PostgreSQL and extension alias translation.

```bash
pig ext add pg_duckdb            # install pg_duckdb
pig ext add pg_duckdb -v 18      # install for PG 18
pig ext add pg_duckdb -y         # auto-confirm installation
pig ext add vector postgis       # install multiple extensions
pig ext add postgis --plan       # preview install plan without executing

# Using alias
pig install pg_duckdb
pig install pg_duckdb -v 18 -y

# Install PostgreSQL kernel packages
pig ext install pgsql            # install latest PostgreSQL kernel
pig ext a pg18                   # install PostgreSQL 18 kernel packages
pig ext ins pg16                 # install PostgreSQL 16 kernel packages
pig ext install pg15-core        # install PostgreSQL 15 core packages
pig ext install pg14-main -y     # install PG 14 + common extensions (vector, repack, wal2json)
```

**Options:**

- `-v|--version`: specify PG major version
- `-y|--yes`: auto-confirm installation
- `--plan`: preview the install plan without running package-manager commands


## ext rm

Remove one or more PostgreSQL extensions.

```bash
pig ext rm pg_duckdb             # remove pg_duckdb
pig ext rm pg_duckdb -v 18       # remove PG 18 package
pig ext rm pgvector -y           # auto-confirm removal
pig ext rm pgvector --plan       # preview remove plan without executing
```

**Options:**

- `-v|--version`: specify PG major version
- `-y|--yes`: auto-confirm removal
- `--plan`: preview the remove plan without running package-manager commands


## ext update

Update explicitly selected installed extensions to the latest version. For safety, `pig ext update` with no arguments is a no-op; you must name the targets explicitly.

```bash
pig ext update                   # no-op: explicit targets are required
pig ext update pg_duckdb         # update one extension
pig ext update postgis timescaledb  # update multiple extensions
pig ext update pg_duckdb -y      # auto-confirm update
```

**Options:**

- `-v|--version`: specify PG major version
- `-y|--yes`: auto-confirm update


## ext import

Download extension packages into a local repository for offline installation.

```bash
pig ext import postgis                # import PostGIS packages
pig ext import timescaledb pg_cron    # import multiple extension packages
pig ext import pg16                   # import PostgreSQL 16 packages
pig ext import pgsql-common           # import common utility packages
pig ext import -d /www/pigsty postgis # import to a selected path
```

**Options:**

- `-d|--repo`: local repository directory (default: `/www/pigsty`)


## ext link

Link a selected PG version into the system PATH.

```bash
pig ext link 18                  # link PG 18 to PATH
pig ext link 16                  # link PG 16 to /usr/pgsql
pig ext link /usr/pgsql-16       # link from a selected path to /usr/pgsql
pig ext link null                # remove current PostgreSQL link
pig ext link none                # null / none / nil / nop / no all remove the link
```

This command creates the `/usr/pgsql` symlink and writes `/etc/profile.d/pgsql.sh`.


## ext reload

Refresh extension metadata.

```bash
pig ext reload                   # refresh extension catalog
```

The updated catalog is stored at `~/.pig/extension.csv`.
