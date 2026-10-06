---
title: "CMD: build"
description: How to setup building infrastructure with pig build subcommand?
icon: fa-solid fa-terminal
weight: 640
---

The `pig build` command simplifies the full workflow for building PostgreSQL extensions from source. It provides build infrastructure setup, dependency management, and compilation environments for standard and custom PostgreSQL extensions across supported operating systems.

```bash
pig build - Build Postgres Extension

Environment Setup:
  pig build spec                   # init build spec and directory (~ext)
  pig build repo                   # init build repo (=repo set -ru)
  pig build repo  [--beta]         # init build repo with PostgreSQL beta repo
  pig build tool  [mini|full|...]  # init build toolset
  pig build rust  [-y] [-m]        # install Rust toolchain
  pig build pgrx  [-v <ver>] [-b]  # install & init pgrx (0.19.3)
  pig build proxy                  # install or verify Xray
  pig build proxy client URI       # setup an HTTP/SOCKS client
  pig build proxy server [flags]   # setup or export a REALITY server

Package Building:
  pig build pkg   [ext|pkg...]     # complete pipeline: get + dep + ext
  pig build get   [ext|pkg...]     # download extension source tarball
  pig build dep   [ext|pkg...]     # install extension build dependencies
  pig build ext   [ext|pkg...]     # build extension package

Quick Start:
  pig build spec                   # setup build spec and directory
  pig build pkg citus              # build citus extension
```

| Command       | Description                              | Notes                 |
|:--------------|:-----------------------------------------|:----------------------|
| `build spec`  | Initialize build specification directory |                       |
| `build repo`  | Initialize required repositories         | Requires sudo or root |
| `build tool`  | Initialize build tools                   | Requires sudo or root |
| `build rust`  | Install Rust toolchain                   | Requires sudo or root |
| `build pgrx`  | Install and initialize pgrx              | Requires sudo or root |
| `build proxy` | Set up Xray clients and servers | Linux: root; macOS client: regular user |
| `build get`   | Download source tarballs                 |                       |
| `build dep`   | Install extension build dependencies     | Requires sudo or root |
| `build ext`   | Build extension packages                 | Requires sudo or root |
| `build pkg`   | Complete pipeline: get, dep, ext         | Requires sudo or root |
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
pig build repo -m                # use China mirrors (Tencent Cloud first for PGDG)
pig build repo -b                # include PostgreSQL beta repository
```

**What it does:** initializes build repositories with `pig repo set -ru`: remove old repositories, add required repositories, and refresh package caches.

**Options:**

- `-m|--mirror`: prefer China mirrors (Tencent Cloud first for PGDG).
- `-b|--beta`: include the PostgreSQL beta repository module, currently for PG19 beta builds.


## build tool

Install development tools and compilers.

```bash
pig build tool                   # install default toolset
pig build tool mini              # minimal toolset
pig build tool full              # full toolset
pig build tool rust              # add Rust development tools
pig build tool -b                # include PostgreSQL beta build packages
```

**Toolsets:**

- **Minimal (`mini`)**: GCC/Clang compilers, Make, core build tools, PostgreSQL development headers, and basic libraries.
- **Default**: minimal tools plus extra compilers, development libraries, and packaging tools such as `rpmbuild` and `dpkg-dev`.
- **Full (`full`)**: default tools plus language-specific development tools and advanced debugging or profiling utilities.

On EL systems, the default build toolset installs explicit PostgreSQL development/server packages for the stable active majors. Use `--beta` when you also need the current beta major.


## build rust

Install the Rust toolchain required by Rust-based extensions.

```bash
pig build rust                   # install with confirmation
pig build rust -y                # force reinstall Rust toolchain
pig build rust -m                # use rsproxy.cn and Cargo mirror config
```

**Installed components:** Rust compiler (`rustc`), Cargo, Rust standard library, and development tools.

**Options:**

- `-y|--yes`: force Rust reinstallation even if Cargo already exists.
- `-m|--mirror`: download `rustup` from the China mirror and write `~/.cargo/config.toml` for `rsproxy.cn`.


## build pgrx

Install and initialize PGRX, the PostgreSQL extension framework for Rust.

```bash
pig build pgrx                   # install default version (0.19.3)
pig build pgrx -v 0.19.3         # install a specific pgrx version
pig build pgrx --pg 18,17,16     # initialize pgrx for selected PG versions
pig build pgrx --pg init         # run cargo pgrx init without PG arguments
pig build pgrx -b                # include beta PostgreSQL pg_config in auto-detect
```

**Prerequisites:** Rust toolchain and PostgreSQL development headers must be installed first.

**Options:**

- `-v|--pgrx`: cargo-pgrx version to install.
- `--pg`: comma-separated PostgreSQL majors, `init` for no PostgreSQL arguments, or empty for auto-detection.
- `-b|--beta`: include the current beta major during auto-detection.


## build proxy

Set up Xray clients and servers for build environments with restricted internet access.
The `x` alias is retained. With no arguments, `pig build proxy` installs or verifies Xray only.
The client and server role commands are available from PIG v1.9.0.
See the [Xray design record](https://pig.pgsty.com/design/xray-client-server/) for the protocol and migration contract.

### Client

One command installs Xray if needed, writes the client configuration, starts its service, and checks
HTTPS through both HTTP and SOCKS. The first listener defaults to **`127.0.0.1:12345`**, with HTTP
and SOCKS sharing that port. An omitted `--listen` preserves the listener of an existing supported
client. The protocol is VLESS over RAW/TCP, REALITY, and `xtls-rprx-vision`, with a Chrome fingerprint
and multiplexing disabled.

```bash
pig build proxy client 'vless://UUID@proxy.example.com:443?encryption=none&security=reality&type=tcp&flow=xtls-rprx-vision&sni=www.sraoss.co.jp&fp=chrome&pbk=PUBLIC_KEY&sid=SHORT_ID&pqv=VERIFY_KEY'
pig build proxy client --server proxy.example.com:443 --id UUID --sni www.sraoss.co.jp --public-key PUBLIC_KEY --short-id SHORT_ID --pqv VERIFY_KEY
pig build proxy client --from ./client.uri
pig build proxy client --from ./client.uri --listen 127.0.0.1:8888
pig build proxy client --from ./client.uri --plan
```

Choose one connection input: a positional URI, direct connection flags, or `--from FILE/-`.
`client.uri` is simply a text file containing one standard `vless://` URI, with an optional trailing
newline; its name is arbitrary. `--from -` reads stdin. Do not mix input forms. `--pqv` supplies
ML-DSA-65 verification when enabled on the server. Unsupported transports, duplicate URI parameters,
and ambiguous inputs are rejected before setup.

On **Linux**, run setup as root or with `sudo`, with systemd running and a repository providing the
Pigsty `xray` package configured first, for example `sudo pig repo add infra -u`. Setup uses
`/etc/xray.json` with mode 0640 and ownership `root:xray`, the `xray` service account, and a PIG-owned
systemd drop-in. On **macOS**, run as your regular user with Homebrew available. Setup uses
`~/.config/xray/pig-proxy.json` with mode 0600 and its own `com.pigsty.xray-proxy` LaunchAgent.

For an already configured server, stream the connection directly into client setup. Both ends must
use a PIG build containing these commands:

```bash
# Linux client: the remote command only reads and exports the existing server connection.
ssh root@proxy.example.com 'pig build proxy server --host proxy.example.com --export-only --export -' | sudo pig build proxy client --from -
# macOS client: omit sudo.
ssh root@proxy.example.com 'pig build proxy server --host proxy.example.com --export-only --export -' | pig build proxy client --from -
```

Repeated setup with matching input, ownership, permissions, and service state leaves files and the
running service unchanged. Incorrect managed-file ownership is repaired without changing connection
credentials. A different or unsupported existing client configuration requires `--replace --yes`; `--plan` previews the
operation without installation, writes, or service changes. Setup checks both proxy protocols using
`https://www.google.com/generate_204` and requires HTTP 204, so the server needs outbound access to
that endpoint. A failed check returns nonzero and restores prior managed files and service state;
an installed package may remain.

The generated shell file defines `po`, `px`, and `pck`. To enable proxy variables in the calling shell:

```bash
# Linux
source /etc/profile.d/proxy.sh
po
# macOS
source ~/.config/xray/proxy.sh
po
```

### Server

Server setup supports Linux with the same package, root, and systemd requirements. Require the
advertised public `--host` and REALITY camouflage `--target host:port`. The first direct listener
defaults to `0.0.0.0:443`; `--port` changes the advertised public port and first direct listener.
`--listen` sets an independent bind endpoint. The target must support TLS 1.3 and HTTP/2.

```bash
# Direct public listener; explicitly export a protected client URI file.
sudo pig build proxy server --host proxy.example.com --target www.sraoss.co.jp:443 --export ./client.uri
# Backend behind an already configured trusted PROXY-protocol ingress: 127.0.0.1:9443.
sudo pig build proxy server --host proxy.example.com --target www.sraoss.co.jp:443 --proxy-protocol --export ./client.uri
# Read an existing supported server and print its connection without changing the deployment.
sudo pig build proxy server --host proxy.example.com --export-only --export -
```

First setup generates a UUID, X25519 key pair, short ID, and ML-DSA-65 seed and verification key.
SNI defaults to the target hostname. Repeating setup retains all authentication fields, SNI, and
protocol settings; an omitted `--listen` retains the existing listener. Matching configuration, file
ownership and permissions, and service state require no rewrite or restart. With the same `--host` and `--port`, export returns the
same canonical URI. Conflicting target/SNI or malformed and ambiguous existing material is rejected
instead of silently rotating credentials. An existing unsupported server configuration is refused.

`--export-only` requires `--export` and does not install, modify configuration, start, restart, or
enable the service. It derives client verification material from the existing private material in
process. One supported VLESS/REALITY inbound, account, SNI, and short ID must be unambiguous.
`--host` and `--port` describe the external endpoint; a backend port such as 9443 is not inferred as
the public port. Do not combine this read-only mode with setup options.

`--export FILE` writes mode 0600 and never includes the server private key or seed. Re-exporting
identical content succeeds without rewriting; a different existing file is refused. `--export -`
explicitly prints one complete credential URI on stdout; diagnostics stay on stderr. Export requires
text output and cannot be combined with JSON/YAML. Ordinary output and plans contain no credentials.
Treat the whole URI and direct authentication flags as credentials; file/stdin input avoids retaining
them as shell arguments.

Server success proves its local listener and service, with public reachability pending a real client
request. PIG does not configure Nginx, firewall rules, or cloud security groups. `--proxy-protocol`
requires loopback binding and an existing trusted frontend. Port conflicts fail without stopping
another service. Existing V2Ray is not automatically stopped or removed.

### Historical VMess form

The historical positional grammar keeps its V2Ray/VMess behavior and first client port 12345:

```bash
pig build proxy user@host:8080
pig build proxy user@host:8080 127.0.0.1:1080
```

This Linux compatibility path still needs a repository providing `vray`, writes `/etc/v2ray.json`
and `/etc/profile.d/proxy.sh`, and restarts `v2ray`. It requires root or `sudo`; a failed HTTPS check
returns nonzero. The remote user ID is a credential and is redacted from ordinary results.

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
pig build get pgedge-17          # download PostgreSQL 17 plus pgEdge suite sources
pig build get babelfish-18       # download Babelfish PG18 source bundle
pig build get orioledb-16        # download OrioleDB plus PG16 patch bundle
pig build get onesparse          # download onesparse, graphblas, and lagraph
```

Common special source aliases include: `babelfish` / `babelfish-17` / `babelfish-18`, `ivorysql`, `orioledb` / `orioledb-16` / `orioledb-17` / `orioledb-18`, `cloudberry` / `cloudberry-backup` / `cloudberry-pxf`, `pgedge` / `pgedge-15` / `pgedge-16` / `pgedge-17` / `pgedge-18`, `polardb`, `polarstore`, `zlog`, `pdu`, `pgdog`, `rdkit`, `onesparse`, `libpgfeutils`, and `libfq`.


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

Debug packages are enabled by default. Compiled RPM builds normally add `debuginfo` and
`debugsource` packages, while Debhelper-based builds normally add a `dbgsym` package. Pure SQL,
architecture-independent, or recipe-specific builds may have no debug payload to split. A package
recipe can also make a narrower explicit choice; PIG does not rewrite spec files or `debian/rules`.

```bash
pig build ext citus              # build one extension
pig build ext citus pgvector     # build multiple extensions
pig build ext citus --pg 17      # build for a selected PG version
pig build ext citus --nodbg      # explicitly omit automatic debug packages
```

**Options:**

- `--pg`: specify one or more PostgreSQL major versions.
- `--nodbg`: disable automatic `debuginfo` / `debugsource` packages on RPM builds and `dbgsym`
  packages on DEB builds.

## build pkg

Run the complete build pipeline: download, dependency installation, and build.

```bash
pig build pkg citus              # build one extension
pig build pkg citus pgvector     # build multiple extensions
pig build pkg citus --pg 17,16   # build for multiple PG versions
pig build pkg citus --nodbg      # explicitly omit automatic debug packages
pig build pkg citus -m           # prefer the pigsty.cc China mirror for sources
```

**Options:**

- `--pg`: specify one or more PostgreSQL major versions.
- `--nodbg`: disable automatic `debuginfo` / `debugsource` packages on RPM builds and `dbgsym`
  packages on DEB builds.
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
cargo install --locked cargo-pgrx@0.19.3

# Reinitialize PGRX
cargo pgrx init
```


## Extension Build Matrix

### Common Extensions to Build

| Extension       |    Type     | Build Time | Complexity | Special Requirements |
|:----------------|:-----------:|:-----------|:-----------|:---------------------|
| pg_repack       |      C      | Fast       | Simple     | None                 |
| pg_partman      | SQL/PLPGSQL | Fast       | Simple     | None                 |
| citus           |      C      | Medium     | Medium     | None                 |
| timescaledb     |      C      | Slow       | Complex    | CMake                |
| postgis         |      C      | Very slow  | Complex    | GDAL, GEOS, Proj     |
| pg_duckdb       |     C++     | Medium     | Medium     | C++17 compiler       |
| pgroonga        |      C      | Medium     | Medium     | Groonga libraries    |
| pgvector        |      C      | Fast       | Simple     | None                 |
| plpython3       |      C      | Medium     | Medium     | Python development   |
| pgrx extensions |    Rust     | Slow       | Complex    | Rust, PGRX           |
{.full-width}
