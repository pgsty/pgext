# pgext

[![Website: pgext.cloud](https://img.shields.io/badge/website-pgext.cloud-336791?logo=postgresql&logoColor=white)](https://pgext.cloud)
[![Release](https://img.shields.io/github/v/release/pgsty/pgext?color=teal)](https://github.com/pgsty/pgext/releases)
[![License: Apache-2.0](https://img.shields.io/badge/license-Apache--2.0-green)](LICENSE)

**PGEXT** helps you discover PostgreSQL extensions, read their usage guides,
and find packages for your PostgreSQL version and Linux system.
[Browse the catalog at pgext.cloud](https://pgext.cloud).

This repository contains the catalog data and the Go application behind
PGEXT.CLOUD. The `pgext` CLI collects APT/RPM repository metadata and builds
package availability matrices. `pgext serve` exposes the catalog as a web
application and JSON query API, with web assets embedded in the binary and
data stored in PostgreSQL.

## The catalog today

Catalog snapshot, **2026-10-29**:

| Layer | Current scope |
| --- | --- |
| Extension directory | **2,703** entries, including packaged, source, and provider-specific extensions |
| Packaged catalog | **620** records: 617 marked available and 3 retained as removed |
| Categories | **16**, covering spatial data, vectors, search, analytics, operations, security, and more |
| PostgreSQL targets | **14, 15, 16, 17, 18** |
| Linux targets | **16** OS/architecture combinations: EL 8/9/10, Debian 12/13, Ubuntu 22.04/24.04/26.04; `x86_64` and `aarch64` |

The broad directory lives in [`pgext.universe`](db/universe.csv); the packaged
baseline lives in [`pgext.extension`](db/extension.csv). The latter groups its
620 records under 512 package names, including PostgreSQL contrib modules.
A package can provide several extensions. Availability varies by PostgreSQL
major version, OS, architecture, and repository; consult the
[package matrix](https://pgext.cloud/matrix) for a specific target.

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
`pgext` on your `PATH`. To build from source, use Go 1.27.3 or later:

```bash
git clone https://github.com/pgsty/pgext.git
cd pgext
go build -o pgext .
```

For a source build, use `./pgext` in the examples below.

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

pgext reload                  # fetch repository metadata and rebuild the matrix
pgext rescan                  # scan local Pigsty repositories and rebuild it
pgext fetch                   # fetch APT/RPM repository metadata
pgext scan                    # scan local repository metadata
pgext parse                   # parse cached metadata and rebuild the matrix
pgext recap                   # rebuild package availability from parsed data

pgext ext <name>              # show extension metadata
pgext pkg <name>              # show a package availability matrix
pgext bin <name> -p 18 -o el9  # show binary packages and download URLs

pgext gen all                 # run content generators
pgext gen lint                # lint bilingual usage stubs
pgext pgxn                    # export PGXN metadata to standalone reports
```

Run `pgext help <command>` for command-specific options. `pgext reload` reads
the current repository manifests, validates and parses their metadata, and
publishes the parsed package tables and availability matrix atomically.

## Development

```bash
go test ./...      # Go unit tests
make check-stubs   # validate usage Markdown
make build         # build the Hugo site
```

The Hugo site uses **Hugo Extended 0.168.0** and **OINK v1.4.0**.
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
