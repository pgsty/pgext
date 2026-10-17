---
title: "CMD: build"
description: How to setup building infrastructure with pig build subcommand?
icon: fa-solid fa-terminal
weight: 640
---

`pig build` downloads sources, installs dependencies, and produces RPM/DEB packages using Pigsty build recipes. Use it on an EL, Debian, or Ubuntu build host. Rust, pgrx, and the Xray client also support their respective macOS workflows.

| Command | Description | Notes |
|:---|:---|:---|
| `build spec` | Initialize build specification directory | |
| `build repo` | Initialize required repositories | Requires sudo or root |
| `build tool` | Initialize build tools | Requires sudo or root |
| `build rust` | Install Rust toolchain | Current user |
| `build pgrx` | Install and initialize pgrx | Current user |
| `build proxy` | Set up Xray clients and servers | Linux: root; macOS client: regular user |
| `build get` | Download source tarballs | |
| `build dep` | Install extension build dependencies | Requires sudo or root |
| `build ext` | Build extension packages | Current user |
| `build pkg` | Complete pipeline: get, dep, ext | Requires sudo or root |
{.full-width}

## Quick Start

If the build host needs package repositories, run `pig build repo`. This backs up and replaces existing repository definitions.
Skip this step if suitable repositories are already configured.

Then install the build tools, download the recipes, and build an extension:

```bash
pig build tool                  # install build tools and PostgreSQL development packages
pig build spec                  # download recipes and prepare directories
pig build pkg citus --pg 18     # download, install dependencies, and build for PG 18
```

Packages are written under `~/ext/pkg/`; RPM outputs may be in architecture subdirectories.

To run each build stage separately:

```bash
pig build get citus
pig build dep citus --pg 18
pig build ext citus --pg 18
```

`build pkg` combines these three stages. It does not install the general build tools or download build recipes.

## Build Infrastructure

### Directory Layout

```filetree {title="PIG build directories"}
- ~/ext/
  - pkg/   {open=false}
  - src/   {open=false}
  - log/   {open=false}
  - tmp/   {open=false}
- ~/rpmbuild/
  - RPMS -> ~/ext/pkg
  - SOURCES -> ~/ext/src
  - SPECS/   {open=false}
  - BUILD/   {open=false}
  - BUILDROOT/   {open=false}
  - SRPMS/   {open=false}
- ~/debbuild/
  - DEBS -> ~/ext/pkg
  - SOURCES -> ~/ext/src
  - SPECS/   {open=false}
  - BUILD/   {open=false}
```

`~/ext/` is the real working directory: `pkg` stores built packages, `src` stores source tarballs, and `log` / `tmp` hold build logs and temporary files. The platform-specific trees link their output and source directories back into it.

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
pig build repo                   # equivalent to: pig repo set
pig build repo -m                # select the bundled China-region mirrors
pig build repo --beta            # also enable the PostgreSQL 19 beta repo module
```

**What it does:** initializes build repositories with `pig repo set`: remove old repositories, add required repositories, and refresh package caches. `--beta/-b` appends the `beta` module for explicit PostgreSQL 19 beta builds; the stable default path still uses PG14-18.

**Options:**

- `-b|--beta`: additionally enable PostgreSQL beta repository modules
- `-m|--mirror`: select the bundled `china` repository definitions

## build tool

Install development tools and compilers.

```bash
pig build tool                   # install default toolset
pig build tool mini              # minimal toolset
pig build tool full              # full toolset
pig build rust                   # install Rust separately
pig build tool --beta            # add PG19 beta build dependencies
```

**Toolsets:**

- **Minimal (`mini`)**: build tools and libraries, with PostgreSQL and PGDG packages removed from the install list.
- **Default / `full`**: Make, CMake, development libraries, packaging tools, and PostgreSQL build dependencies. EL selects PG14-18 server/devel packages; Debian/Ubuntu uses `postgresql-all` and `postgresql-server-dev-all`.
- **`--beta`**: adds explicit PG19 beta server/devel packages. `mini` still excludes PostgreSQL packages.

## build rust

Install the Rust toolchain required by Rust-based extensions.

```bash
pig build rust                   # install if Cargo is absent
pig build rust -y                # force reinstall Rust toolchain
pig build rust -m                # use China mirror mode and write Cargo mirror config
```

PIG installs Rust through rustup for the current user, without an interactive prompt. If `~/.cargo/bin/cargo` exists, it skips installation unless `-y` is set. `-m` uses `rsproxy.cn` for rustup and writes Cargo mirror configuration, including when Rust is already installed.

## build pgrx

Install and initialize PGRX, the PostgreSQL extension framework for Rust.

```bash
pig build pgrx                   # install default version (0.19.3)
pig build pgrx -v 0.19.3         # install a specific pgrx version
pig build pgrx --pg 18,17,16     # initialize pgrx for selected PG versions
pig build pgrx --pg init         # run cargo pgrx init without PG arguments
pig build pgrx -b                # include PostgreSQL 19 beta pg_config during auto-detection
```

**Prerequisites:** Rust toolchain and PostgreSQL development headers must be installed first. Default auto-detection only covers stable PG14-18; use `-b|--beta`, or explicitly pass `--pg 19`, when PG19 beta is required.

## build proxy

A matching macOS LaunchAgent generated by the older source without a management comment can
upgrade without `--replace`. Recognition compares the complete generated plist; a comment alone
does not grant ownership. Foreign or edited service files require `--replace --yes`. After the
one-time comment migration, identical execution leaves the service running without a restart.
This compatibility statement covers the verified old source format, not every released binary.

Set up Xray clients and servers for build environments with restricted internet access.
The `x` alias is retained. With no arguments, `pig build proxy` installs or verifies Xray only.
See the [Xray design record](https://pig.pgsty.com/design/xray-client-server/) for the design rationale.

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

For an existing server, pass its connection URI directly to client setup:

```bash
# Linux client: the remote command only reads and exports the existing server connection.
ssh root@proxy.example.com 'pig build proxy server --host proxy.example.com --export-only --export -' | sudo pig build proxy client --from -
# macOS client: omit sudo.
ssh root@proxy.example.com 'pig build proxy server --host proxy.example.com --export-only --export -' | pig build proxy client --from -
```

Repeated setup reuses matching configuration and service state. To replace a different client configuration, use `--replace --yes`. Use `--plan` to preview the operation.

After setup, PIG requests `https://www.google.com/generate_204` through both HTTP and SOCKS and requires HTTP 204. If the checks fail, PIG restores the previous configuration and service state. Any installed package remains.

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

First setup generates a UUID, X25519 key pair, short ID, and ML-DSA-65 seed and verification key. SNI defaults to the `--target` hostname. Repeated setup preserves these credentials. If you omit `--listen`, it also preserves the listener; specify `--listen` to change it. Matching configuration does not restart the service. Conflicting targets or unsupported existing configurations produce an error.

Connection export options:

| Option | Purpose |
|:---|:---|
| `--export FILE` | Write a client URI with mode `0600`; identical content can be exported again, but different content is not overwritten |
| `--export -` | Print the complete client URI on stdout and diagnostics on stderr |
| `--export-only` | Read and export existing configuration without changing files or services; requires `--export` |
| `--plan` | Preview setup |

Export requires text output. The URI contains client credentials, so pass it through a file or stdin. It does not contain the server's private key or seed. When exporting an existing server, `--host` and `--port` specify the public endpoint the client reaches, not the backend listener.

After setup, test connectivity from a real client. Open the required firewall or cloud security-group rules. `--proxy-protocol` requires an existing trusted frontend and a loopback listener. PIG does not stop another service that occupies the port or remove an existing V2Ray installation.

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
pig build get pgedge             # download both PostgreSQL and spock sources
pig build get onesparse          # download onesparse, graphblas, and lagraph
```

Common special source aliases include: `babelfish`, `agensgraph`, `orioledb`, `cloudberry`, `pgedge`, `pdu`, `pgdog`, `rdkit`, `onesparse`, and `libpgfeutils`.

## build dep

Install dependencies required to build extensions.

```bash
pig build dep citus              # single extension
pig build dep citus pgvector     # multiple extensions
pig build dep citus --pg 17,16   # for selected PG versions
```

**Options:**

- `--pg`: specify one or more PostgreSQL major versions. If omitted, pig infers versions from extension metadata or local PostgreSQL installations.

For a list such as `--pg 17,16`, version-specific dependency names are expanded for every
selected major version and duplicates are removed before installation. The same dependency
step is used by `build pkg`.

## build ext

Compile extensions and create installation packages.

For `build ext` and `build pkg`, omitting `--pg` retains version selection from extension
metadata and platform exceptions; Debian/Ubuntu also consider locally installed PG versions.
Package-matrix gaps do not remove build targets. An explicit `--pg` selects the requested versions.
Debian/Ubuntu builds preserve the inherited `PATH` order and append the current user's
`~/.cargo/bin` when that directory exists, so an explicitly selected toolchain keeps precedence.

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

Download sources, install dependencies, then build each requested package. A source-download failure stops that package. A dependency-install failure produces a warning and the build still runs, because the dependencies may already be present.

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

### Build and install a local package

After preparing the environment in Quick Start, build for the target PostgreSQL version:

```bash
pig build pkg pg_partman --pg 18
find ~/ext/pkg -type f -name '*partman*'
```

Select the package that matches the target PostgreSQL version and architecture, then install it with the system package manager. Replace these placeholder paths with the actual package path:

```bash
sudo dnf install /absolute/path/to/package.rpm   # EL
sudo apt install /absolute/path/to/package.deb   # Debian / Ubuntu
```

`pig ext add NAME` installs from configured repositories. Use a file path as shown above to install the package you just built.

### Build a Rust extension

In a prepared build environment, install Rust and the pgrx version required by the recipe for the current user:

```bash
pig build rust
pig build pgrx --pg 18
pig build pkg pgmq --pg 18
```

Add `-m` to `build rust` to use China mirrors. If the recipe requires a specific pgrx version, run `pig build pgrx -v VERSION --pg 18`. Install the local package as described above.

### Build for multiple PostgreSQL versions

```bash
pig build pkg citus --pg 15,16,17
```

PIG produces packages for each major version. Each version needs matching development packages and recipe support.

### Build for PostgreSQL 19 beta

```bash
pig build repo --beta
pig build tool --beta
pig build pgrx -b                # only for Rust/pgrx extensions
pig build pkg EXTENSION --pg 19
```

Replace `EXTENSION` with one that supports the selected version. The beta version comes from PIG's current embedded configuration.

## Troubleshooting

| Problem | Action |
|:---|:---|
| Missing build tools | Run `pig build tool`; install any additional compiler required by the recipe |
| Missing extension dependencies | Run `pig build dep NAME --pg MAJOR`; check the reported packages and repositories |
| PostgreSQL headers not found | Run `pig ext add pg18-devel`, then select the matching target with `--pg 18` |
| Wrong pgrx version | Run `pig build pgrx -v VERSION --pg MAJOR` with the version required by the recipe |
| Stale source archive | Run `pig build get NAME -f`; first check that the catalog's source version matches the recipe |

Build logs are in `~/ext/log/`. Build time depends on the recipe, CPU, memory, parallelism, and caches. PostGIS, pg_duckdb, and Rust extensions usually need more resources.
