# pgext

[![Website: pgext.cloud](https://img.shields.io/badge/website-pgext.cloud-336791?logo=postgresql&logoColor=white)](https://pgext.cloud)
[![Docs: doc.pgext.cloud](https://img.shields.io/badge/docs-doc.pgext.cloud-4169E1)](https://doc.pgext.cloud)
[![Release](https://img.shields.io/github/v/release/pgsty/pgext?color=teal)](https://github.com/pgsty/pgext/releases)
[![License: Apache-2.0](https://img.shields.io/badge/license-Apache--2.0-green)](LICENSE)

**pgext** catalogs the PostgreSQL extension universe: 1,600+ extensions with metadata,
dependencies, categories, and exact package availability across PostgreSQL versions,
operating systems, and repositories (PGDG / Pigsty, APT / YUM).

It is a single static binary that does two things:

- **Build the catalog**: fetch repository metadata, parse APT/RPM packages, and
  maintain the curated extension knowledge base — all stored in PostgreSQL.
- **Serve the catalog**: a self-contained web app + JSON API, the same software
  that powers [pgext.cloud](https://pgext.cloud).

## Quick Start

All you need is a PostgreSQL database to hold the catalog:

```bash
export PGURL="postgres:///data"   # where the catalog lives (default: postgres:///data)

pgext init     # one-time setup: create the pgext schema and load catalog data
pgext serve    # web UI + JSON API at http://localhost:8432
```

That's it. Point your browser at [http://localhost:8432](http://localhost:8432).

```bash
pgext serve --db meta --listen :8080   # database name shorthand, custom port
curl 'localhost:8432/api/v1/ext?q=vector'   # query the JSON API
```

## Installation

Download a single-binary release for Linux / macOS (amd64 / arm64) from the
[release page](https://github.com/pgsty/pgext/releases), or `deb` / `rpm` packages
if you prefer a package manager.

Build from source (Go 1.26+, no CGO required):

```bash
git clone https://github.com/pgsty/pgext.git
cd pgext && go build -o pgext .
```

## Commands

```text
pgext init                    # setup everything (schema + reload)
pgext serve                   # serve web UI + JSON API (:8432)
pgext status                  # show metadata status

pgext reload                  # reload data: fetch + complete parse
pgext fetch                   # get repo metadata from upstream
pgext scan                    # scan local Pigsty repo metadata
pgext parse                   # atomically populate apt, dnf, bin, pkg
pgext recap                   # rebuild pkg from current catalog data

pgext ext <name>              # show extension information
pgext pkg <name>              # show package availability matrix
pgext bin <name> -p 17 -o el9 # show binary packages with URLs

pgext gen all                 # run every content generator
pgext pgxn                    # crawl PGXN dist metadata
pgext purge                   # drop the pgext schema
```

Run `pgext help <command>` for details on any command.

## Links

- Catalog website: [pgext.cloud](https://pgext.cloud)
- Documentation: [doc.pgext.cloud](https://doc.pgext.cloud)
- Pigsty — battery-included PostgreSQL distribution: [pigsty.io](https://pigsty.io)
- `pig` — the PostgreSQL extension package manager: [github.com/pgsty/pig](https://github.com/pgsty/pig)

## License

[Apache License 2.0](LICENSE) © [Ruohang Feng](https://vonng.com)
