---
title: "CMD: build"
description: How to setup building infrastructure with pig build subcommand?
icon: SquareTerminal
weight: 640
---

The `pig build` command simplifies the full workflow for building PostgreSQL extensions from source. It provides build infrastructure setup, dependency management, and compilation environments for standard and custom PostgreSQL extensions across supported operating systems.

```bash
pig build - Build Postgres Extension

Environment Setup:
  pig build spec                   # init build spec and directory (~ext)
  pig build repo                   # init build repo (=repo set -ru)
  pig build tool  [mini|full|...]  # init build toolset
  pig build rust  [-y]             # install Rust toolchain
  pig build pgrx  [-v <ver>]       # install & init pgrx (0.19.1)
  pig build proxy [id@host:port ]  # init build proxy (optional)

Package Building:
  pig build pkg   [ext|pkg...]     # complete pipeline: get + dep + ext
  pig build get   [ext|pkg...]     # download extension source tarball
  pig build dep   [ext|pkg...]     # install extension build dependencies
  pig build ext   [ext|pkg...]     # build extension package

Quick Start:
  pig build spec                   # setup build spec and directory
  pig build pkg citus              # build citus extension
```

| Command | Description | Notes |
|:---|:---|:---|
| `build spec` | Initialize build specification directory | |
| `build repo` | Initialize required repositories | Requires sudo or root |
| `build tool` | Initialize build tools | Requires sudo or root |
| `build rust` | Install Rust toolchain | Requires sudo or root |
| `build pgrx` | Install and initialize pgrx | Requires sudo or root |
| `build proxy` | Initialize build proxy | |
| `build get` | Download source tarballs | |
| `build dep` | Install extension build dependencies | Requires sudo or root |
| `build ext` | Build extension packages | Requires sudo or root |
| `build pkg` | Complete pipeline: get, dep, ext | Requires sudo or root |
{.full-width}


## Quick Start

The fastest way to set up a build environment and build an extension:

```bash
# Step 1: initialize build specs
pig build spec

# Step 2: build an extension with the complete pipeline
pig build pkg citus

# Built packages will be placed under:
# - EL: ~/ext/pkg/ (also available via ~/rpmbuild/RPMS/)
# - Debian: ~/ext/pkg/ (also available via ~/debbuild/DEBS/)
```

For finer control:

```bash
# Environment setup
pig build spec                   # initialize build specs
pig build repo                   # configure repositories
pig build tool                   # install build tools

# Build steps
pig build get citus              # download source
pig build dep citus              # install dependencies
pig build ext citus              # build package

# Or run all three steps at once
pig build pkg citus              # get + dep + ext
```


## Build Infrastructure

### Directory Layout

```
~/ext/                           # real working directory
├── pkg/                         # built package output
├── src/                         # source tarball downloads
├── log/                         # build logs
└── tmp/                         # temporary files

~/rpmbuild/                      # EL build directory
├── RPMS -> ~/ext/pkg            # RPM output symlink
├── SOURCES -> ~/ext/src         # source symlink
├── SPECS/
├── BUILD/
├── BUILDROOT/
└── SRPMS/

~/debbuild/                      # Debian / Ubuntu build directory
├── DEBS -> ~/ext/pkg            # DEB output symlink
├── SOURCES -> ~/ext/src         # source symlink
├── SPECS/
└── BUILD/
```

**Build output locations:**

- **EL systems**: `~/ext/pkg/`, also accessible through `~/rpmbuild/RPMS/`
- **Debian systems**: `~/ext/pkg/`, also accessible through `~/debbuild/DEBS/`


## build spec

Set up build specifications and directory layout.

```bash
pig build spec                   # initialize ~/ext at the default location
pig build spec -f                # force re-download of the build spec tarball
pig build spec -m                # prefer the pigsty.cc China mirror
```

**What it does:**

1. Downloads the RPM or DEB build specification tarball.
2. Creates `~/ext/{pkg,src,log,tmp}` and the platform build directory.
3. Links `RPMS`/`DEBS` and `SOURCES` to `~/ext/pkg` and `~/ext/src`.
4. Syncs makefiles, specs, and Debian packaging files with incremental `rsync`.

**Working directory:** `~/ext/` stores sources, packages, logs, and temporary files. The platform packaging directory is `~/rpmbuild/` or `~/debbuild/`.


## build repo

Initialize package repositories required for building extensions.

```bash
pig build repo                   # equivalent to: pig repo set -ru
```

**What it does:** initializes build repositories with `pig repo set -ru`: remove old repositories, add required repositories, and refresh package caches.


## build tool

Install development tools and compilers.

```bash
pig build tool                   # install default toolset
pig build tool mini              # minimal toolset
pig build tool full              # full toolset
pig build tool rust              # add Rust development tools
```

**Toolsets:**

- **Minimal (`mini`)**: GCC/Clang compilers, Make, core build tools, PostgreSQL development headers, and basic libraries.
- **Default**: minimal tools plus extra compilers, development libraries, and packaging tools such as `rpmbuild` and `dpkg-dev`.
- **Full (`full`)**: default tools plus language-specific development tools and advanced debugging or profiling utilities.


## build rust

Install the Rust toolchain required by Rust-based extensions.

```bash
pig build rust                   # install with confirmation
pig build rust -y                # force reinstall Rust toolchain
```

**Installed components:** Rust compiler (`rustc`), Cargo, Rust standard library, and development tools.


## build pgrx

Install and initialize PGRX, the PostgreSQL extension framework for Rust.

```bash
pig build pgrx                   # install latest stable version (0.19.1)
pig build pgrx -v 0.19.1         # install a specific pgrx version
pig build pgrx --pg 18,17,16     # initialize pgrx for selected PG versions
pig build pgrx --pg init         # run cargo pgrx init without PG arguments
```

**Prerequisites:** Rust toolchain and PostgreSQL development headers must be installed first.


## build proxy

Configure a proxy for build environments with restricted internet access.

```bash
pig build proxy                  # interactive setup
pig build proxy user@host:8080   # use default local endpoint 127.0.0.1:12345
pig build proxy user@host:8080 127.0.0.1:1080
```


## build get

Download extension source tarballs.

```bash
pig build get citus              # single extension
pig build get citus pgvector     # multiple extensions
pig build get pdu pgdog          # use built-in source aliases
pig build get citus -f           # re-download even if the file already exists
pig build get citus -m           # prefer the pigsty.cc China mirror
```

Arguments to `pig build get` are extension names, package names, or source filenames. Unknown names are treated as source filenames. It does not expand `all` or `std` into built-in package sets; list the target package names explicitly for batch downloads.

Some source packages do not map directly to extension names, so `pig build get` includes special aliases for direct source downloads.

```bash
pig build get pdu                # download pdu-3.0.25.12.tar.gz
pig build get pgdog              # download pgdog-0.1.32.tar.gz
pig build get pgedge             # download both PostgreSQL and spock sources
pig build get onesparse          # download onesparse, graphblas, and lagraph
```

Common special source aliases include: `babelfishpg` / `babelfish`, `agensgraph` / `agentsgraph`, `oriolepg` / `orioledb`, `cloudberry`, `pgedge`, `pdu`, `pgdog`, `rdkit`, `onesparse`, and `libfepgutils`.


## build dep

Install dependencies required to build extensions.

```bash
pig build dep citus              # single extension
pig build dep citus pgvector     # multiple extensions
pig build dep citus --pg 17,16   # for selected PG versions
```

**Options:**

- `--pg`: specify one or more PostgreSQL major versions. If omitted, pig infers versions from extension metadata or local PostgreSQL installations.


## build ext

Compile extensions and create installation packages.

```bash
pig build ext citus              # build one extension
pig build ext citus pgvector     # build multiple extensions
pig build ext citus --pg 17      # build for a selected PG version
pig build ext citus -s           # include debug symbols (RPM only)
```

**Options:**

- `--pg`: specify one or more PostgreSQL major versions.
- `-s|--symbol`: build debug symbol packages (RPM only).


## build pkg

Run the complete build pipeline: download, dependency installation, and build.

```bash
pig build pkg citus              # build one extension
pig build pkg citus pgvector     # build multiple extensions
pig build pkg citus --pg 17,16   # build for multiple PG versions
pig build pkg citus -s           # include debug symbols
pig build pkg citus -m           # prefer the pigsty.cc China mirror for sources
```

**Options:**

- `--pg`: specify one or more PostgreSQL major versions.
- `-s|--symbol`: build debug symbol packages (RPM only).
- `-m|--mirror`: prefer the `pigsty.cc` mirror when downloading source files.


## Common Workflows

### Workflow 1: Build a Standard Extension

```bash
# 1. Set up the build environment once
pig build spec
pig build repo
pig build tool

# 2. Build the extension
pig build pkg pg_partman

# 3. Install the built package
sudo rpm -ivh ~/ext/pkg/pg_partman*.rpm  # EL
sudo dpkg -i ~/ext/pkg/*partman*.deb     # Debian
```

### Workflow 2: Build a Rust Extension

```bash
# 1. Set up Rust environment
pig build spec
pig build tool
pig build rust                   # add -y only if you need to force reinstall
pig build pgrx

# 2. Build Rust extension
pig build pkg pgmq

# 3. Install
sudo pig ext add pgmq
```

### Workflow 3: Build Multiple Versions

```bash
# Build for multiple PostgreSQL versions
pig build pkg citus --pg 15,16,17

# Results in packages for each version:
# citus_15-*.rpm
# citus_16-*.rpm
# citus_17-*.rpm
```


## Troubleshooting

### Build Tools Not Found

```bash
# Install build tools
pig build tool

# For specific compiler sets
sudo dnf groupinstall "Development Tools"  # EL
sudo apt install build-essential          # Debian
```

### Missing Dependencies

```bash
# Install extension dependencies
pig build dep <extension>

# Check error messages for specific packages
# Install manually if needed
sudo dnf install <package>  # EL
sudo apt install <package>  # Debian
```

### PostgreSQL Headers Not Found

```bash
# Install PostgreSQL development package
sudo pig ext install pg18-devel

# Or specify pg_config path
export PG_CONFIG=/usr/pgsql-18/bin/pg_config
```

### Rust/PGRX Issues

```bash
# Reinstall Rust
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# Update PGRX
cargo install --locked cargo-pgrx@0.19.1

# Reinitialize PGRX
cargo pgrx init
```


## Extension Build Matrix

### Common Extensions to Build

| Extension | Type | Build Time | Complexity | Special Requirements |
|:---|:---:|:---|:---|:---|
| pg_repack | C | Fast | Simple | None |
| pg_partman | SQL/PLPGSQL | Fast | Simple | None |
| citus | C | Medium | Medium | None |
| timescaledb | C | Slow | Complex | CMake |
| postgis | C | Very slow | Complex | GDAL, GEOS, Proj |
| pg_duckdb | C++ | Medium | Medium | C++17 compiler |
| pgroonga | C | Medium | Medium | Groonga libraries |
| pgvector | C | Fast | Simple | None |
| plpython3 | C | Medium | Medium | Python development |
| pgrx extensions | Rust | Slow | Complex | Rust, PGRX |
{.full-width}
