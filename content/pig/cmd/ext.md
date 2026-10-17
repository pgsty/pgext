---
title: "CMD: ext"
description: How to manage extensions with pig ext subcommand
icon: fa-solid fa-terminal
weight: 620
---

`pig ext` searches the extension catalog and installs, updates, or removes PostgreSQL extension packages on the host. It also supports PostgreSQL server packages. After installing a package, apply the SQL and configuration required by the extension.

| Command | Description | Notes |
|:---|:---|:---|
| `ext list` | Search extensions | |
| `ext info` | Show extension details | |
| `ext avail` | Show extension availability matrix | |
| `ext status` | Show host-visible extension files/packages | Does not query every database |
| `ext scan` | Scan installed control/library files | Does not query every database |
| `ext add` | Install extensions | Requires sudo or root |
| `ext rm` | Remove extensions | Requires sudo or root |
| `ext update` | Update extensions | Requires sudo or root |
| `ext import` | Download extensions for offline use | Requires sudo or root |
| `ext link` | Link a PG version into PATH | Requires sudo or root |
| `ext reload` | Refresh extension catalog | |
{.full-width}

## Quick Start

This example uses a running PostgreSQL 18 instance. To install and start one, follow [Getting Started](/pig/start/).

Add the PostgreSQL repositories, inspect the extension, then install its package:

```bash
pig repo add pgsql -u
pig ext list vector
pig ext info vector
pig ext add vector -v 18
```

Enable the extension in the target database. For example:

```sql
CREATE EXTENSION vector;
```

See the [extension list](https://pigsty.io/ext/list/) for other extensions.
If an extension requires preloading or other configuration, complete the steps from `pig ext info NAME` before enabling it.

### Select a PostgreSQL installation

- `-v, --version 18`: select a PostgreSQL major version.
- `-p, --path /path/to/pg_config`: select an installation by its `pg_config` path.
- Use one option at a time. When neither is supplied, PIG attempts local PostgreSQL detection. Set `-v` explicitly in automation.

The default PGDG paths are `/usr/pgsql-18/bin/pg_config` on EL and `/usr/lib/postgresql/18/bin/pg_config` on Debian/Ubuntu. The package manager resolves RPM/DEB dependencies. `-y` confirms package operations automatically.

## Package and database operations {#package-layer-vs-database-layer}

| Task | PIG operation | Database operation |
|:---|:---|:---|
| Install | `ext add NAME` | Configure preloading if required, then run `CREATE EXTENSION` |
| Upgrade | `ext update NAME` | Follow the extension's upgrade instructions and run `ALTER EXTENSION ... UPDATE` |
| Remove | `ext rm NAME` | Remove or migrate dependencies in affected databases before uninstalling the package |
| Inspect | `ext status` / `ext scan` inspect host files | Query `pg_extension` for enabled database extensions |
| Refresh catalog | `ext reload` updates extension metadata | None |

`pg_available_extensions` lists extensions available to the server. `pg_extension` lists extensions enabled in the current database.

## ext list

List or search extensions in the extension catalog.

```bash
pig ext list                     # list all extensions
pig ext list duck                # search extensions containing "duck"
pig ext list -v 18               # evaluate availability for PG 18
pig ext ls olap                  # list OLAP category extensions
pig ext ls gis -v 16             # list GIS entries with PG 16 availability
pig ext ls rag                   # list RAG category extensions
```

Category filtering is done by passing the category name as the query. Supported categories include: `time`, `gis`, `rag`, `fts`, `olap`, `feat`, `lang`, `type`, `func`, `util`, `admin`, `stat`, `sec`, `fdw`, `sim`, and `etl`.

**Options:**

- `-v|--version`: select the PG version used to evaluate availability; unavailable entries remain in the list
- `--pkg`: show package names instead of extension names, listing only leading extensions

**Status column:**

- `installed`: extension is installed
- `available`: extension is available but not installed
- `not avail`: extension is not available on the current system

Availability and text/JSON/YAML summaries use the same platform, architecture, and PostgreSQL
version. For platforms covered by a package matrix, that matrix takes precedence; a missing cell
is unavailable. Catalog entries without a matrix retain RPM/DEB metadata and platform-exception
fallbacks. `amd64`/`x86_64` and `arm64`/`aarch64` are equivalent. On macOS, this status describes
declared PostgreSQL version support, not the availability of an RPM/DEB package.

### Catalog count {#catalog-count}

The website and CLI count every catalog record, including packaged SQL utilities such as `pg_relation_sql` and retained records for removed extensions. The bundled catalog currently contains {{< param pig_catalog_count >}} entries. This total does not mean that every entry supports `CREATE EXTENSION` or is available on every platform.

The unfiltered list includes every record. Its text row count and JSON/YAML `data.count` describe the rows returned; a search counts its matches, and `--pkg` counts leading package records. The global `ext avail` text summary uses the full catalog total. `ext reload` reports the number of records in the downloaded catalog through `data.extension_count`.

The [counting decision](https://pig.pgsty.com/design/catalog-entry-count/) explains this shared definition. The default extension catalog is defined in [**`cli/ext/assets/extension.csv`**](https://github.com/pgsty/pig/blob/main/cli/ext/assets/extension.csv).

Use `pig ext reload` to refresh to the latest catalog. The downloaded catalog is stored at `~/.pig/extension.csv`; the latest online catalog is also published at [**repo.pigsty.io/ext/data/extension.csv**](https://repo.pigsty.io/ext/data/extension.csv).

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

Show extension files/packages visible to the selected PostgreSQL installation on this host.

```bash
pig ext status              # show installed extensions
pig ext status -c           # include contrib extensions
pig ext status -v 16        # show installed extensions for PG 16
```

**Options:**

- `-c|--contrib`: include contrib extensions in the result

## ext scan

Scan the selected PostgreSQL installation's extension directories.

```bash
pig ext scan [-v version]
```

The command reads control files and related installation paths for the selected PostgreSQL installation.
It also matches shared-library names against catalog metadata, so a catalog entry without a
control file can still be reported as installed. This does not imply `CREATE EXTENSION` support.

## ext add

Install one or more PostgreSQL extensions. Aliases for `pig ext add` include `pig ext install`, `pig ext ins`, and `pig ext a`. The top-level [`pig install`](/pig/cmd/#pig-install) command also accepts extension aliases and general OS packages.

```bash
pig ext add pg_duckdb            # install pg_duckdb
pig ext add pg_duckdb -v 18      # install for PG 18
pig ext add pg_duckdb -y         # auto-confirm installation
pig ext add vector postgis       # install multiple extensions
pig ext add postgis --plan       # preview install plan without executing

# Top-level package installer
pig install pg_duckdb
pig install pg_duckdb -v 18 -y

# Install PostgreSQL packages
pig ext add pg18                 # install PostgreSQL 18
pig ext add pg18-client          # install client tools only
```

Use these aliases to select a PostgreSQL package group:

| Alias | Packages |
|:---|:---|
| `pg18` | Server, client, contrib modules, and procedural languages |
| `pg18-server` | Server and contrib modules |
| `pg18-client` | Client tools |
| `pg18-devel` | Development files |
| `pg18-basic` | pg_repack, wal2json, and pgvector |

Replace `18` with the required major version. You can also use `pgsql` with `-v 18`.
To select a specific package version, use `NAME=VERSION` with a version available in the configured repository.

**Options:**

- `-v|--version`: specify PG major version
- `-y|--yes`: auto-confirm installation
- `--plan`: preview the install plan without running package-manager commands

Successful installation only makes the extension files available on the host. Follow the extension's `ext info` hints for preload, restart, and `CREATE EXTENSION` steps.

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

This removes an OS package, not database objects. Drop or migrate dependent extensions in every affected database before removing the shared files.

## ext update

Update the named packages to the versions available from the configured repositories. With no arguments, the command makes no changes.

```bash
pig ext update                   # no-op: explicit targets are required
pig ext update pg_duckdb         # update one extension
pig ext update postgis timescaledb  # update multiple extensions
pig ext update pg_duckdb -y      # auto-confirm update
pig ext update pg_duckdb -m      # refresh the catalog from the China mirror first
```

**Options:**

- `-v|--version`: specify PG major version
- `-y|--yes`: auto-confirm update
- `-m|--mirror`: refresh the extension catalog from the `pigsty.cc` mirror before updating packages; package downloads still use the configured OS repositories

The command does not run SQL migrations. After the package upgrade, follow the extension's release notes and run `ALTER EXTENSION ... UPDATE` in each database where appropriate.

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

If some names cannot be resolved, successfully resolved packages are still downloaded, but the
command exits with failure. Structured output preserves `requested`, resolved `packages`,
confirmed `downloaded` packages, and unresolved `failed` names. A native download failure can
also leave files in the destination; it does not imply rollback or a complete download.

On Debian/Ubuntu, virtual dependency markers such as `<virtual-package>` are excluded from
the package download list; real package dependencies remain eligible.

JSON/YAML output and `PIG_NON_INTERACTIVE=1` use non-interactive sudo. If privilege escalation
requires a password, the command fails instead of waiting for input; configure suitable sudo access first.

## ext link

Link a selected PG version into the system PATH.

```bash
pig ext link 18                  # link PG 18 to PATH
pig ext link pg17                # strip the pg prefix and link PG 17
pig ext link 16                  # link PG 16 to /usr/pgsql
pig ext link /usr/pgsql-16       # link from a selected path to /usr/pgsql
pig ext link polar               # link PolarDB / PolarPG installation
pig ext link /usr/polar-17       # link from a selected PolarDB path
pig ext link null                # remove current PostgreSQL link
pig ext link none                # null / none / nil / nop / no all remove the link
```

This command creates the `/usr/pgsql` symlink and writes `/etc/profile.d/pgsql.sh`. Load the profile in your current shell:

```bash
. /etc/profile.d/pgsql.sh
```

## ext reload

Refresh extension metadata.

```bash
pig ext reload                   # refresh extension catalog
```

The updated catalog is stored at `~/.pig/extension.csv`.
