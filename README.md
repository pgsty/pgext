# pgext

[![Website: pgext.cloud](https://img.shields.io/badge/website-pgext.cloud-336791?logo=postgresql&logoColor=white)](https://pgext.cloud)
[![Release](https://img.shields.io/github/v/release/pgsty/pgext?color=teal)](https://github.com/pgsty/pgext/releases)
[![License: Apache-2.0](https://img.shields.io/badge/license-Apache--2.0-green)](LICENSE)

**PGEXT** is a PostgreSQL extension builder and catalog. Discover extensions,
read their usage guides, find packages for your PostgreSQL version and Linux
system, or build RPM/DEB packages from source.
[Browse the catalog at pgext.cloud](https://pgext.cloud).

This repository contains the catalog data and the Go application behind
PGEXT.CLOUD. The `pgext` CLI collects APT/RPM repository metadata and builds
package availability matrices. It also builds extension RPM/DEB packages
directly from the Pigsty recipes, without PIG or a metadata database.
`pgext serve` exposes the catalog as a web application and JSON query API,
with web assets embedded in the binary and data stored in PostgreSQL.

## The catalog today

Local catalog snapshot, **2026-11-09**:

| Layer | Current scope |
| --- | --- |
| Extension directory | **2,733** entries, including packaged, source, and provider-specific extensions |
| Packaged directory entries | **625** marked `packaged`, including PostgreSQL contrib and retained entries |
| Package matrix | **457** distinct package families across **36,560** recorded target slots: 33,544 available and 3,016 not applicable |
| Categories | **16**, covering spatial data, vectors, search, analytics, operations, security, and more |
| PostgreSQL targets | **14, 15, 16, 17, 18** |
| Linux targets | **16** OS/architecture combinations: EL 8/9/10, Debian 12/13, Ubuntu 22.04/24.04/26.04; `x86_64` and `aarch64` |

The broad directory lives in [`pgext.universe`](db/universe.csv); the packaged
baseline lives in [`pgext.extension`](db/extension.csv). The latter has
626 records (624 available, 2 removed), grouped under 518 package names,
including PostgreSQL contrib modules. These are separate counts: the homepage
uses Universe's `packaged` flag, while its package and slot totals come from
the active `pgext.pkg` matrix. A package can provide several extensions;
matrix slots include missing and unsupported combinations, not just successful
builds. Availability varies by PostgreSQL major version, OS, architecture,
and repository; consult the [package matrix](https://pgext.cloud/matrix) for
a specific target or [`/api/v1/meta`](https://pgext.cloud/api/v1/meta) for the
running service's current totals.

## Explore the catalog

Use the [catalog](https://pgext.cloud) to:

- **Find extensions** by category, license, language, or PostgreSQL version.
- **Plan installation** with dependency, preload, and compatibility information.
- **Choose packages** from PGDG and Pigsty, with versions, checksums, and download links.
- **Read usage guides** in English and Chinese, with links to upstream sources.

The JSON query API is available without running a local instance:

```bash
curl 'https://pgext.cloud/api/v1/ext?q=vector'
curl 'https://pgext.cloud/api/v1/ext/vector/files?pg=18&os=el9.x86_64'
curl 'https://pgext.cloud/api/v1/matrix?repo=pgdg'
```

See the [server reference](server/README.md) for endpoints,
filters, and server options.

## Install

Download a binary for Linux or macOS (`amd64`/`arm64`), or a Linux `deb`/`rpm`
package, from the [releases](https://github.com/pgsty/pgext/releases), and put
`pgext` on your `PATH`. To build from source, use Go 1.27.3:

```bash
git clone https://github.com/pgsty/pgext.git
cd pgext
go build -o pgext .
```

For a source build, use `./pgext` in the examples below.

## Build extensions

Build natively on EL 8/9/10, Debian 12/13, or Ubuntu 22.04/24.04/26.04,
on `amd64` or `arm64`. Inside the Linux host or container, prepare the
environment once (root or sudo is needed for system package operations):

```bash
pgext env
pgext env repo
pgext env tool --pg 18
pgext env spec

# For Rust/pgrx extensions; select the recipe's cargo-pgrx version:
pgext env rust
pgext env pgrx --pgrx 0.19.3 --pg 18
```

Then run the complete pipeline or its individual stages:

```bash
pgext pkg pgvector --pg 18

pgext get pgvector --pg 18
pgext dep pgvector --pg 18
pgext mak pgvector --pg 18
```

`make` is the command name; `m`, `mk`, and `mak` are aliases. Examples use
`mak`. It builds and verifies packages from prepared inputs; `pkg` runs
`get → dep → mak` and stops on failure. Both preserve debug packages by default.
Use `--pg 17,18` for multiple majors; omitting it selects installed majors
supported by the recipe within PG14–18. Extension names such as `vector`
also resolve to their source package (`pgvector`).

Build commands choose metadata from `--catalog`, then
`$XDG_CONFIG_HOME/pgext/extension.csv` (normally `~/.config/pgext/extension.csv`),
then the embedded `db/extension.csv`. They run without PIG or `PGURL`.
Recipes default to `~/rpmbuild` or `~/debbuild`; use `--spec-dir` to select a
checkout, or `pgext env spec --from /path/to/recipes` to install one.

Preview a build without downloading, installing, or compiling. This also
works on macOS with a local recipe checkout:

```bash
pgext pkg vector --pg 18 --plan --platform u24.arm64 \
  --spec-dir ../deb/debbuild --json
```

The workspace defaults to `~/ext` (`--work-dir` overrides it): source archives
go to `src/`, isolated build trees to `tmp/`, verified packages to `pkg/`, and
stage logs plus `result.json` manifests to `log/run-*/`. Use `--json` for a
machine-readable result, `--mirror` on source/environment commands that support
it to prefer the China mirror, and `--checksums` for a filename-to-SHA256 source
manifest. Building produces local artifacts; installation and publication are
separate operations.

See [the build workflow](docs/build-workflow.md) for Docker integration,
local recipes, source verification, outputs, and migration from PIG.

## Manage repositories

`pgext repo` manages Linux package sources using the same repository definitions
as `confgen` and Pigsty. It supports repository names, modules, regional mirrors,
configuration previews, backups, and guarded restores without a database.

```bash
pgext repo list
pgext repo show pgsql --platform u24.arm64 --mirror
pgext repo add pgdg pigsty --dry-run
sudo pgext repo add pgdg pigsty --update
sudo pgext repo set all --replace --mirror
pgext repo status
sudo pgext repo rm docker
sudo pgext repo restore <backup-id>
```

`all` selects `node,pgsql,infra`. `add`, `set`, and `rm` normally manage only
PGEXT-owned files; `set --replace` backs up and replaces all existing system
source definitions. `set` also updates the cache by default. Use `--dry-run`
to inspect changes and `--json` for scripts. The old database-backed repository
summary is available as `pgext repo summary`.

See [repository management](docs/repository-management.md) for custom Pigsty
catalogs, platform overrides, trust settings, and isolated staging.

## Run it locally

With PostgreSQL running, create a database and initialize the catalog using
a role that can create the `pgext` schema. Initialization loads the catalog
shipped with the binary, downloads repository indexes, and builds the package
matrix.

```bash
createdb data
export PGURL='postgres:///data'

pgext init
pgext serve
```

Open [http://localhost:8432](http://localhost:8432). `PGURL` selects the database;
`pgext serve --db` overrides it. For example, to use an existing database on
the local `/tmp` socket and listen on port 8080:

```bash
pgext serve --db 'postgresql://postgres@%2Ftmp:5432/data' --listen :8080
```

## Commands

```text
pgext init                    # load the bundled catalog and fetch package data
pgext serve                   # serve the web UI and JSON API
pgext status                  # show catalog and repository status
pgext repo list               # list available Linux package repositories
pgext repo summary            # show fetched repository metadata in the database

pgext reload                  # fetch repository metadata and rebuild the matrix
pgext rescan                  # scan local Pigsty repositories and rebuild it
pgext fetch                   # fetch APT/RPM repository metadata
pgext scan                    # scan local repository metadata
pgext parse                   # parse cached metadata and rebuild the matrix
pgext recap                   # rebuild package availability from parsed data

pgext info <name>             # show extension metadata (alias: ext)
pgext matrix <name>           # show a package availability matrix
pgext bin <name> -p 18 -o el9  # show binary packages and download URLs

pgext env                    # inspect the build environment (no database)
pgext env repo               # prepare builder package repositories
pgext env tool --pg 18       # install build tools and PostgreSQL development files
pgext env spec               # install or update RPM/DEB recipes
pgext get <name>             # download source inputs
pgext dep <name> --pg 18     # install recipe dependencies
pgext mak <name> --pg 18     # build and verify packages from prepared inputs
pgext pkg <name> --pg 18     # get → dep → mak

pgext gen all                 # run content generators
pgext gen lint                # lint bilingual usage stubs
pgext pgxn                    # export PGXN metadata to standalone reports
```

Run `pgext help <command>` for command-specific options. `pgext reload` reads
the current repository manifests, validates and parses their metadata, and
publishes the parsed package tables and availability matrix atomically.
The former read-only `pgext pkg` query is now `pgext matrix`; `pkg` runs
the extension build pipeline.

## Development

```bash
go test ./...                 # Go unit tests
node --test server/*_test.cjs  # frontend behavior tests
make check-stubs              # validate usage Markdown
make build                    # build the Hugo site
```

The Hugo site uses **Hugo Extended 0.168.0** and **OINK v1.5.1**.
`make dev` previews it at [http://localhost:1313](http://localhost:1313) using a
sibling `../oink` checkout; set `THEME_DIR` to use another local theme checkout.

Catalog snapshots and schema definitions live in `db/`. English and Chinese
usage sources live in `stub/` and `stub-zh/`, and are loaded separately into
`pgext.doc` with `bin/load-stub`. Follow the
[stub documentation SOP](docs/stub-documentation-sop.md) for source review,
validation, and loading.

## Related projects

- [PGSTY](https://pgsty.com) — open-source PostgreSQL software, infrastructure, and services.
- [Pigsty](https://pigsty.io) — a PostgreSQL distribution with monitoring, high availability, and backups.
- [PIG](https://pig.pgsty.com) — install PostgreSQL and extensions through APT/DNF.
- [PG.CENTER](https://pg.center) — PostgreSQL manuals, reference material, and a knowledge graph.

## License

[Apache License 2.0](LICENSE) © [Ruohang Feng](https://vonng.com)
