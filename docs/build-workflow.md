# Building extensions with PGEXT

PGEXT prepares a Linux builder, obtains source archives, installs native build
dependencies, builds extension RPM/DEB packages, and verifies the resulting
packages. The build commands require neither PIG nor a metadata database.

## Command contract

| PIG operation | PGEXT operation | Effect |
| --- | --- | --- |
| `pig build repo` | `pgext env repo` | Configure builder repositories and refresh metadata |
| `pig build tool` | `pgext env tool` | Install the compiler, packaging tools, and PostgreSQL development packages |
| `pig build spec` | `pgext env spec` | Install or update native recipes |
| `pig build rust` | `pgext env rust` | Install Rust and optionally configure the China mirror |
| `pig build pgrx` | `pgext env pgrx` | Install cargo-pgrx and register installed PostgreSQL versions |
| `pig build proxy` | `pgext env proxy` | Install/configure the build proxy |
| `pig build get` | `pgext get` | Download source inputs |
| `pig build dep` | `pgext dep` | Install recipe dependencies |
| `pig build ext` | `pgext mak` | Build and verify packages from prepared inputs |
| `pig build pkg` | `pgext pkg` | Run get, dep, and mak in order |

The canonical build command is `make`, with aliases `m`, `mk`, and `mak`.
Use `mak` in examples and maintained scripts. Here, “build only” still includes
the native recipe's packaging and output verification. It does not install
the built extension into the host PostgreSQL instance.

`pkg` stops at the first failed stage or package. `get`, `dep`, and `mak` do
not invoke one another. Literal APT installation steps found in older DEB
Makefiles run during `dep` and become no-ops in the staged recipe used by
`mak`. Native recipes and their build scripts remain executable code; this
is not an offline sandbox for arbitrary upstream build scripts.

This API changes the old read-only `pgext pkg <name>` to
`pgext matrix <name>`. Extension metadata is available as `pgext info <name>`;
`ext` remains an alias for that query. Query commands retain their existing
database behavior.

## Prepare a builder

Supported systems are EL8/9/10, Debian 12/13, and Ubuntu 22.04/24.04/26.04,
on amd64 or arm64. Package builds run natively inside the selected Linux
system or container. macOS can download sources and preview build plans.

Run the following in a disposable builder as root, or use sudo for package
installation and system repository operations:

```bash
pgext env repo --mirror
pgext env tool --pg 17,18
pgext env spec --mirror
pgext env rust --mirror
pgext env pgrx --pgrx 0.19.3 --pg 17,18
```

Omit `--mirror` to prefer the global source service. `env tool` defaults to
PG14–18; `env tool mini` installs packaging tools without PG server/devel
packages. `env pgrx` registers installed PG14–18 installations, or requires
every version explicitly selected by `--pg`. Choose the cargo-pgrx version
required by the extension recipe. `env rust` preserves existing Cargo
configuration; `--force` backs up and replaces a differing mirror config.

Use `pgext env` for the detected platform and configured paths. Environment
commands accept `--plan` where shown by their help. Proxy client/server
subcommands expose their own `--plan` option. Environment preparation does not
load the extension catalog, so an invalid local CSV cannot block repairs.

Repository maintenance is deliberately limited to builder sources:

```bash
pgext env repo list
pgext env repo add pgdg pigsty --mirror
pgext env repo update
pgext env repo rm pigsty
```

`env repo` defaults to `add all`; modules are `node`, `pgdg`, and `pigsty`.
Node setup retains distribution sources and enables the required EL
CRB/PowerTools/EPEL repositories. PGDG/Pigsty definitions use signature
verification and separate `pgext-*` filenames. `rm` removes only PGEXT-owned
PGDG/Pigsty definitions, not OS repositories. Managed definition changes
are backed up, and restored if their cache refresh fails; installing EPEL
or changing an OS module is not undone by that file rollback.

For an existing PIG builder with `infra`/`pgsql` source files, migrate both
modules together:

```bash
pgext env repo add pgdg pigsty --replace --mirror
```

Only recognized legacy files qualify. Mixed/custom definitions are rejected
for explicit migration. The result records the backup directory. New Docker
build stages start from base images and need no legacy migration.

## Catalog and recipes

Catalog selection is, in priority order:

1. `--catalog /path/to/extension.csv`.
2. `$XDG_CONFIG_HOME/pgext/extension.csv`, or `~/.config/pgext/extension.csv`.
3. The existing `db/extension.csv` embedded in the PGEXT binary.

An invalid explicitly selected or existing local catalog is an error, not a
reason to silently use a different catalog. Canonical extension names and
package names resolve to the same source package; family members are deduplicated.
Native recipe aliases cover H3, OneSparse, address_standardizer and
postgis_tiger_geocoder. Without a recipe, the catalog's source list takes
precedence over the built-in fallback bundles.

No encoded availability matrix is needed. The catalog supplies source names,
family identities and package naming; the native recipe declares build
inputs and targets. `rpm_pg`, `deb_pg` and PIG's `extra.matrix` describe
published package availability and do not restrict building missing packages.
Publication selection remains a separate use of the live package matrix and
the existing supplier/gap policy.

Default recipe roots are `~/rpmbuild` and `~/debbuild`. Install a local
checkout or a recipe archive without downloading a published snapshot:

```bash
pgext env spec --from /inputs/rpmbuild
# On Debian/Ubuntu:
pgext env spec --from /inputs/debbuild
```

Use `--spec-dir` for a different recipe root. `env spec` excludes source
caches and old packages, updates recipes without deleting local files, and
creates `SOURCES` and output links. Existing nonempty directories at those
link locations are preserved and reported for migration.

## Download, build, and inspect

```bash
pgext pkg pgvector --pg 18 --mirror --json > result.json

pgext get citus --pg 15,18 --mirror
pgext dep citus --pg 15,18
pgext mak citus --pg 15,18 --jobs 4

pgext pkg pgmq --pg 17,18
pgext pkg pg_eviltransform --pg 18 --jobs 4
```

Without `--pg`, builds select installed majors supported by the recipe,
limited to PG14–18. An explicit target outside declared support fails before
building. For DEB recipes, `debian/pgversions` declares support; a Makefile's
`PG_VERSIONS` default is used only when that declaration is absent or empty.
Packages independent of PostgreSQL majors do not accept `--pg`.
RPM receives `pgmajorversion`; DEB receives selected `pgversions`,
`PG_VERSIONS`, `PG_MAJOR`, `PG_SUPPORTED_VERSIONS`, and `PG_CONFIG`. Numbered
Makefile targets such as `build18` are selected explicitly.
Source-only `get` prepares all declared source variants without requiring
PostgreSQL development tools to be installed first.
DEB recipe inputs include auxiliary wheels, source descriptors, headers and
wrapper JARs. Recipe declarations take precedence over the fallback source
bundles. RPM URL renames use the same filename in both `get` and `mak`.
PG selection constrains build outputs; a recipe's dependency on a metapackage
such as `postgresql-all` can still install additional PostgreSQL toolchains.

Preview from macOS using local recipes:

```bash
pgext pkg vector --pg 18 --plan --platform u24.arm64 \
  --spec-dir ../deb/debbuild --json
```

Plans do not download, install, or build. RPM Source macros are expanded by
native `rpmspec` at execution, so a cross-platform preview may only show the
catalog's initial source list.

`get` fetches from `repo.pigsty.io/ext/src` with the China service as fallback;
`--mirror` reverses the preference and `--source-url` selects another base
URL. Downloads are atomic, cached inputs are reused, and `--force` refreshes
them. Supply `--checksums` to get/pkg/mak for a JSON object mapping every
source filename to its SHA256. Missing or mismatched entries fail; without
this option hashes are recorded for provenance, not checked against an
independent trusted manifest. Exact source names and native package versions
are retained independently of normalized catalog versions.

DEB builds bind source tarball variables to the checked copies in the isolated
`SOURCES` directory. DocumentDB also stages its four offline dependency archives
and gives its preparation script explicit paths. GraphBLAS, LAGraph and
pg_mooncake retain their shared `kitware-cmake` helper; on Ubuntu 22.04, `dep`
prepares that private CMake while keeping all other package constraints enforced.

The default workspace is `~/ext`, configurable with `--work-dir`:

```text
src/                         source archives
tmp/<package-random>/pg18/    isolated native build trees
pkg/<package-random>/pg18/    binary packages, including debug packages
log/run-*/                   per-stage logs and result.json
```

Source archives and recipes are copied into isolated build trees. Authoritative
recipes are not rewritten. Old artifacts cannot satisfy a new build. Native
package names, raw versions, architecture, expected companion packages,
nonempty payloads, ELF build IDs and matching DWARF are checked. RPM native
builds also require debugsource; LLVM-enabled extension recipes retain their
bitcode in the main package and must not add LLVM runtime dependencies.
Pure SQL packages need no artificial debug packages. `--no-debug` is an
explicit opt-out; debug output remains the default.

With `--json`, stdout contains the result while native output goes to stderr
and stage logs. The saved result includes catalog and recipe hashes, input
and artifact hashes, native package versions, tool versions and failure
details. The recipe tree fingerprint also covers packaging controls, patches
and helper scripts. Staged source hashes are recorded before invoking the
native build, including when compilation subsequently fails.
Set `PGEXT_BUILD_IMAGE` to the inspected container image ID to record
it. Build trees and logs remain for diagnosis; successful compilation and
package validation do not themselves install, import, sign or publish packages.

## Docker migration and verification scope

The sibling `pgsty/docker` repository uses `pgext.mk` to compile a local
Linux binary for each architecture and supplies it as the named `pgext`
Docker build context. All 16 standard build stages use `pgext env`.
`make meta` copies the local catalog and recipe inputs into a running builder;
`make refresh-build` updates its PGEXT binary and pgrx before committing the
build image. Base/latest and build tags remain distinct.

Representative acceptance covers C/PGXS, Rust/pgrx, pure SQL, multiple PG
majors and package payload/debug validation on Linux builders. Platform
planning tests cover all 16 targets; this is not proof that every extension
recipe has built on every target. Specialized PostgreSQL kernel suites such
as Cloudberry's compound top-level Make targets are outside the extension
command contract and retain their native repository workflows.

PIG's existing commands remain available for rollback until the replacement
has been released and the production builder matrix has been refreshed.

Initial acceptance on 2026-11-08 included these native builds, with PIG absent
from the command path and no metadata database connection:

| Builder | Extension | PG | Verified output |
| --- | --- | --- | --- |
| EL9 arm64 | pgvector 0.8.7 | 18 | main RPM, debuginfo, debugsource, LLVM bitcode/index |
| EL9 arm64 | pgmq 1.13.1 | 18 | one noarch RPM |
| Ubuntu 24.04 arm64 | pgvector 0.8.7 | 18 | main DEB and dbgsym |
| Ubuntu 24.04 arm64 | pgmq 1.13.1 | 17, 18 | exactly two Architecture: all packages |
| Ubuntu 24.04 arm64 | pg_eviltransform 0.0.7 | 18 | Rust/pgrx main DEB and matching DWARF |
| Ubuntu 24.04 arm64 | Citus 13.5.0 / 14.3.0 | 15 / 18 | version-specific main DEBs and dbgsym, including companion libraries |

Fresh EL9 and Ubuntu base containers also passed repository, tool and recipe
initialization. Ubuntu passed Rust and cargo-pgrx installation, then both
the pgvector and pg_eviltransform pipelines. EL9 completed the pgvector
pipeline from its fresh environment. A standalone `mak pgmq --pg 18` also
passed with only a make stage and an intentionally unreachable PGURL.
Legacy repository migration
was exercised on both package families. Existing Citus packaging emitted
debugedit warnings about a DWARF form; build ID pairing and `.debug_info`
verification passed, but this does not establish full debugger source mapping.

Code validation includes `go test ./...`, `go vet ./...`, race checks for
the build/command packages, and Linux/macOS compilation. All 16 Docker
Makefile workflows passed dry runs, with Dockerfile checks on u24a and el9a.
These checks did not replace or publish production image tags.

The follow-up review also checked source plans for 428 local DEB recipes,
native rejection of missing debug packages, and repository rollback after
injected APT/DNF cache failures. Revised pgvector packages passed installation,
`CREATE EXTENSION vector`, and a vector-distance query in isolated PG18 clusters
on both EL9 and Ubuntu 24.04 arm64. Standalone `mak` passed for EL9 pgmq and
Ubuntu pg_eviltransform. A GoReleaser snapshot produced all four platform
archives and both Linux architectures' RPM/DEB packages without publication.
