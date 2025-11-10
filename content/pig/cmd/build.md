---
title: "CMD: build"
description: How to setup building infrastructure with pig build subcommand?
icon: SquareTerminal
weight: 640
---

The `pig build` command is a powerful tool that simplifies the entire workflow of building PostgreSQL extensions from source. It provides a complete build infrastructure setup, dependency management, and compilation environment for both standard and custom PostgreSQL extensions across different operating systems.

------

## Subcommands

| Command | Description |
|---------|------------|
| `spec`  | Initialize building spec repo |
| `repo`  | Initialize required repos |
| `tool`  | Initialize build tools |
| `rust`  | Install Rust toolchain |
| `pgrx`  | Install and initialize pgrx (requires Rust) |
| `proxy` | Initialize build proxy |
| `get`   | Download source code tarball |
| `dep`   | Install extension build dependencies |
| `ext`   | Build extension package |
| `pkg`   | Complete build pipeline: get, dep, ext |

------

## Overview

```bash
Build Postgres Extension

Usage:
  pig build [command]

Aliases:
  build, b

Examples:
pig build - Build Postgres Extension

Environment Setup:
  pig build spec                   # init build spec and directory (~ext)
  pig build repo                   # init build repo (=repo set -ru)
  pig build tool  [mini|full|...]  # init build toolset
  pig build rust  [-y]             # install Rust toolchain
  pig build pgrx  [-v <ver>]       # install & init pgrx (0.16.1)
  pig build proxy [id@host:port ]  # init build proxy (optional)

Package Building:
  pig build pkg   [ext|pkg...]     # complete pipeline: get + dep + ext
  pig build get   [ext|pkg...]     # download extension source tarball
  pig build dep   [ext|pkg...]     # install extension build dependencies
  pig build ext   [ext|pkg...]     # build extension package

Quick Start:
  pig build spec                   # setup build spec and directory
  pig build pkg citus              # build citus extension

Available Commands:
  dep         Install extension build dependencies
  ext         Build extension package
  get         Download source code tarball
  pgrx        Install and initialize pgrx (requires Rust)
  pkg         Complete build pipeline: get, dep, ext
  proxy       Initialize build proxy
  repo        Initialize required repos
  rust        Install Rust toolchain
  spec        Initialize building spec repo
  tool        Initialize build tools

Flags:
  -h, --help   help for build

Global Flags:
      --debug              enable debug mode
  -H, --home string        pigsty home path
  -i, --inventory string   config inventory path
      --log-level string   log level: debug, info, warn, error, fatal, panic (default "info")
      --log-path string    log file path, terminal by default
```

------

## Quick Start

The fastest way to set up a build environment and build an extension:

```bash
# Step 1: Initialize build specifications
pig build spec

# Step 2: Build an extension (complete pipeline)
pig build pkg citus

# The built package will be available in:
# - EL: ~/rpmbuild/RPMS/
# - Debian: ~/
```

For a more controlled approach:

```bash
# Setup environment
pig build spec                   # Initialize build specs
pig build repo                   # Setup repositories
pig build tool                   # Install build tools

# Build process
pig build get citus              # Download source
pig build dep citus              # Install dependencies
pig build ext citus              # Build package

# Or do all three steps at once
pig build pkg citus              # get + dep + ext
```

------

## Build Infrastructure

### Build Specifications

The build system uses specification files that define how each extension should be built. These specs include:
- Source code location and version
- Build dependencies
- Compilation flags
- PostgreSQL version compatibility
- Platform-specific build instructions

### Directory Structure

```
~/ext/                           # Default build spec directory
├── Makefile                     # Master build makefile
├── <extension>/                 # Per-extension directory
│   ├── Makefile                # Extension-specific makefile
│   ├── <extension>.spec        # RPM spec file (EL)
│   └── debian/                 # Debian packaging files
│       ├── control
│       ├── rules
│       └── ...
```

Build output locations:
- **EL Systems**: `~/rpmbuild/RPMS/<arch>/`
- **Debian Systems**: `~/` (deb files)

------

## Command Reference

### `build spec`

Sets up the build specification repository and directory structure.

```bash
pig build spec                   # Initialize at default location ~/ext
```

**What it does:**
1. Clones or updates the extension build spec repository
2. Creates necessary directory structure
3. Sets up makefiles and build scripts
4. Prepares platform-specific packaging files

**Repository Location:**
- Default: `~/ext/`
- Contains build specifications for 100+ extensions

**Example:**
```bash
# Initialize specs
pig build spec

# Check available extensions
ls ~/ext/

# View specific extension spec
cat ~/ext/citus/Makefile
```

------

### `build repo`

Initializes required package repositories for building extensions.

```bash
pig build repo                   # Equivalent to: pig repo set -ru
```

**What it does:**
- Removes existing repositories
- Adds all required repositories (pgdg, pigsty, node)
- Updates package cache

This ensures you have access to all necessary packages and dependencies for building.

------

### `build tool`

Installs the necessary development tools and compilers.

```bash
pig build tool                   # Install default toolset
pig build tool mini              # Minimal toolset
pig build tool full              # Complete toolset
pig build tool rust              # Add Rust development tools
```

**Tool Packages:**

**Minimal (`mini`):**
- GCC/Clang compiler
- Make and build essentials
- PostgreSQL development headers
- Basic libraries

**Default:**
- All minimal tools
- Additional compilers (g++, clang++)
- Development libraries
- Packaging tools (rpmbuild, dpkg-dev)

**Full (`full`):**
- All default tools
- Language-specific tools (Python, Perl, Ruby dev)
- Advanced debugging tools
- Performance analysis tools

**Common packages installed:**
```bash
# EL Systems
gcc gcc-c++ make cmake
postgresql17-devel
rpm-build rpmdevtools
git wget curl

# Debian Systems
build-essential
postgresql-server-dev-17
dpkg-dev debhelper
git wget curl
```

------

### `build rust`

Installs the Rust programming language toolchain, required for Rust-based extensions.

```bash
pig build rust                   # Install with confirmation
pig build rust -y                # Auto-confirm installation
```

**What it installs:**
- Rust compiler (rustc)
- Cargo package manager
- Rust standard library
- Development tools

**Installation method:**
- Uses rustup installer
- Installs to user home directory
- Automatically configures PATH

**Example:**
```bash
# Install Rust
pig build rust -y

# Verify installation
rustc --version
cargo --version
```

------

### `build pgrx`

Installs and initializes PGRX (PostgreSQL extension framework for Rust).

```bash
pig build pgrx                   # Install latest stable (0.16.1)
pig build pgrx -v 0.15.0         # Install specific version
```

**Prerequisites:**
- Rust toolchain must be installed first
- PostgreSQL development headers

**What it does:**
1. Installs cargo-pgrx tool
2. Initializes PGRX for installed PostgreSQL versions
3. Sets up development environment

**Example:**
```bash
# Install Rust first
pig build rust -y

# Install PGRX
pig build pgrx

# Initialize for specific PG versions
cargo pgrx init --pg17=/usr/pgsql-17/bin/pg_config
```

------

### `build proxy`

Sets up proxy configuration for build environments with restricted internet access.

```bash
pig build proxy                  # Interactive setup
pig build proxy user@host:8080   # Direct configuration
pig build proxy http://proxy.company.com:3128
```

**Configures proxy for:**
- System environment (HTTP_PROXY, HTTPS_PROXY)
- Package managers (yum/dnf, apt)
- Build tools (cargo, npm, pip)
- Git configuration

**Example:**
```bash
# Setup corporate proxy
pig build proxy proxy.corp.com:8080

# With authentication
pig build proxy user:pass@proxy.corp.com:8080
```

------

### `build get`

Downloads extension source code tarballs.

```bash
pig build get citus              # Single extension
pig build get citus pgvector     # Multiple extensions
pig build get all                # All available extensions
pig build get std                # Standard extensions
```

**Source locations:**
- Downloads to: `~/ext/<extension>/`
- Source types:
  - Release tarballs from GitHub
  - Official project downloads
  - Source repositories

**Example:**
```bash
# Download specific version
pig build get citus

# Check downloaded source
ls ~/ext/citus/*.tar.gz
```

------

### `build dep`

Installs required dependencies for building extensions.

```bash
pig build dep citus              # Single extension
pig build dep citus pgvector     # Multiple extensions
pig build dep citus --pg 17,16   # For specific PG versions
```

**Options:**
- `--pg`: Comma-separated PostgreSQL versions (e.g., '17,16')

**Dependency types:**
- Build tools (compilers, make)
- Development libraries
- PostgreSQL headers
- Extension-specific requirements

**Example:**
```bash
# Install dependencies for Citus
pig build dep citus

# For multiple PostgreSQL versions
pig build dep citus --pg 17,16,15
```

------

### `build ext`

Compiles the extension and creates installation packages.

```bash
pig build ext citus              # Build single extension
pig build ext citus pgvector     # Build multiple
pig build ext citus --pg 17      # For specific PG version
pig build ext citus -s           # With debug symbols (RPM only)
```

**Options:**
- `--pg`: Target PostgreSQL versions
- `-s, --symbol`: Include debug symbols

**Build process:**
1. Extracts source code
2. Configures build environment
3. Compiles extension
4. Creates platform package (RPM/DEB)
5. Signs package (if configured)

**Output locations:**
- EL: `~/rpmbuild/RPMS/<arch>/*.rpm`
- Debian: `~/*.deb`

**Example:**
```bash
# Build Citus for PostgreSQL 17
pig build ext citus --pg 17

# Check built package
ls ~/rpmbuild/RPMS/x86_64/citus*.rpm  # EL
ls ~/citus*.deb                        # Debian
```

------

### `build pkg`

Executes the complete build pipeline: download, dependencies, and build.

```bash
pig build pkg citus              # Build single extension
pig build pkg citus pgvector     # Build multiple
pig build pkg citus --pg 17,16   # For multiple PG versions
pig build pkg citus -s           # With debug symbols
```

**Equivalent to running:**
```bash
pig build get <extension>
pig build dep <extension>
pig build ext <extension>
```

**Options:**
- `--pg`: Target PostgreSQL versions
- `-s, --symbol`: Include debug symbols

**Example:**
```bash
# Complete build pipeline
pig build pkg citus --pg 17

# Build multiple extensions
pig build pkg citus pg_partman timescaledb
```

------

## Common Workflows

### Workflow 1: Building Standard Extension

```bash
# 1. Setup build environment (once)
pig build spec
pig build repo
pig build tool

# 2. Build extension
pig build pkg pg_partman

# 3. Install built package
sudo rpm -ivh ~/rpmbuild/RPMS/x86_64/pg_partman*.rpm  # EL
sudo dpkg -i ~/pg_partman*.deb                         # Debian
```

### Workflow 2: Building Rust Extension

```bash
# 1. Setup Rust environment
pig build spec
pig build tool
pig build rust -y
pig build pgrx

# 2. Build Rust extension
pig build pkg pgmq

# 3. Install
sudo pig ext add pgmq
```

### Workflow 3: Building Multiple Versions

```bash
# Build extension for multiple PostgreSQL versions
pig build pkg citus --pg 15,16,17

# Results in packages for each version:
# citus_15-*.rpm
# citus_16-*.rpm
# citus_17-*.rpm
```

### Workflow 4: Custom Extension Build

```bash
# 1. Create custom spec
mkdir ~/ext/myextension
cd ~/ext/myextension

# 2. Create Makefile
cat > Makefile << 'EOF'
EXTENSION = myextension
VERSION = 1.0.0
REPO_URL = https://github.com/myorg/myextension
include ../Makefile
EOF

# 3. Build
pig build pkg myextension
```

### Workflow 5: Debugging Build Issues

```bash
# Build with debug output
pig build pkg citus --debug

# Build with debug symbols
pig build pkg citus -s

# Step-by-step debugging
pig build get citus              # Check source download
pig build dep citus              # Verify dependencies
pig build ext citus --debug     # Debug compilation
```

### Workflow 6: Building for Distribution

```bash
# 1. Build multiple extensions
extensions="citus pg_partman timescaledb pgvector"
for ext in $extensions; do
    pig build pkg $ext --pg 16,17
done

# 2. Collect packages
mkdir ~/packages
cp ~/rpmbuild/RPMS/x86_64/*.rpm ~/packages/  # EL
cp ~/*.deb ~/packages/                        # Debian

# 3. Create repository
pig repo create ~/packages

# 4. Distribute
tar czf extensions.tar.gz ~/packages
```

------

## Extension Build Specifications

### Standard Extensions

Extensions with official build support:
- **Distributed**: citus, pg_partman
- **Time-series**: timescaledb, pg_cron
- **Analytics**: pg_duckdb, parquet_fdw
- **Search**: pgroonga, pg_bigm
- **Languages**: plpython3, plperl, pltcl

### Build Requirements

Different extensions have different requirements:

**Simple C Extensions:**
- PostgreSQL development headers
- C compiler (gcc/clang)
- Make

**Complex Extensions:**
- Additional libraries (e.g., GDAL for PostGIS)
- Language runtimes (Python, Perl, etc.)
- Special compilers (Rust, Go)

**Rust Extensions (via PGRX):**
- Rust toolchain
- PGRX framework
- Cargo build system

### Platform Differences

**EL Systems (RHEL/Rocky/Alma):**
- Uses RPM spec files
- Builds with rpmbuild
- Outputs to ~/rpmbuild/RPMS/

**Debian/Ubuntu Systems:**
- Uses debian/ directory structure
- Builds with dpkg-buildpackage
- Outputs to ~/

------

## Troubleshooting

### Build Tools Not Found

```bash
# Install build tools
pig build tool

# For specific compiler
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
sudo pig ext install pg17-devel

# Or specify pg_config path
export PG_CONFIG=/usr/pgsql-17/bin/pg_config
```

### Rust/PGRX Issues

```bash
# Reinstall Rust
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# Update PGRX
cargo install cargo-pgrx --force

# Reinitialize PGRX
cargo pgrx init
```

### Source Download Failures

```bash
# Check network/proxy
pig build proxy

# Download manually
cd ~/ext/<extension>
wget <source_url>

# Continue build
pig build dep <extension>
pig build ext <extension>
```

### Build Failures

```bash
# Check build logs
tail -f ~/rpmbuild/BUILD/*/config.log  # EL

# Common fixes:
# 1. Update build tools
pig build tool full

# 2. Clear build cache
rm -rf ~/rpmbuild/BUILD/*

# 3. Try with different PG version
pig build pkg <extension> --pg 16
```

------

## Best Practices

1. **Setup Once, Build Many**: Initialize the build environment once, then build multiple extensions

2. **Use pkg Command**: For standard builds, use `pig build pkg` for the complete pipeline

3. **Check Compatibility**: Verify extension supports your PostgreSQL version before building

4. **Keep Specs Updated**: Regularly update build specs with `pig build spec`

5. **Test Builds**: Test built packages in development before production deployment

6. **Document Custom Builds**: Keep notes on custom build flags or modifications

7. **Use Version Control**: Track custom build specs in version control

8. **Parallel Builds**: Build independent extensions in parallel for efficiency

9. **Cache Dependencies**: Keep commonly used dependencies installed to speed up builds

10. **Sign Packages**: For distribution, sign packages with GPG for security

------

## Advanced Topics

### Custom Build Flags

```bash
# Set custom CFLAGS
export CFLAGS="-O3 -march=native"
pig build ext <extension>

# PostgreSQL-specific flags
export PG_CPPFLAGS="-DUSE_SPECIAL_FEATURE"
pig build ext <extension>
```

### Cross-Compilation

```bash
# For different architecture
export CC=aarch64-linux-gnu-gcc
export ARCH=arm64
pig build ext <extension>
```

### Debugging Symbols

```bash
# Build with debug symbols
pig build ext <extension> -s

# For debugging
export CFLAGS="-g -O0"
pig build ext <extension>
```

### Custom PostgreSQL Builds

```bash
# For custom PostgreSQL installation
export PG_CONFIG=/opt/postgresql/bin/pg_config
export PKG_CONFIG_PATH=/opt/postgresql/lib/pkgconfig
pig build ext <extension>
```

### Continuous Integration

```bash
#!/bin/bash
# CI build script
set -e

# Setup
pig build spec
pig build tool

# Build matrix
for pg_version in 15 16 17; do
    for extension in citus pgvector timescaledb; do
        pig build pkg $extension --pg $pg_version
    done
done

# Test installations
for rpm in ~/rpmbuild/RPMS/x86_64/*.rpm; do
    sudo rpm -Uvh --test $rpm
done
```

------

## Extension Build Matrix

### Commonly Built Extensions

| Extension | Type | Build Time | Complexity | Special Requirements |
|-----------|------|------------|------------|---------------------|
| pg_repack | C | Fast | Simple | None |
| pg_partman | SQL/PLPGSQL | Fast | Simple | None |
| citus | C | Medium | Medium | None |
| timescaledb | C | Slow | Complex | CMake |
| postgis | C | Very Slow | Complex | GDAL, GEOS, Proj |
| pg_duckdb | C++ | Medium | Medium | C++17 compiler |
| pgroonga | C | Medium | Medium | Groonga library |
| pgvector | C | Fast | Simple | None |
| plpython3 | C | Medium | Medium | Python dev |
| pgrx extensions | Rust | Slow | Complex | Rust, PGRX |