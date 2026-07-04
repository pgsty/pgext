---
title: Get Started
description: Quick start guide for PIG, the PostgreSQL package manager
icon: Play
weight: 200
breadcrumbs: false
---

Here is a simple getting started tutorial to help you experience the core capabilities of the PIG package manager.

## Short Version

```bash
curl -fsSL https://repo.pigsty.io/pig | bash   # Install PIG from Cloudflare
pig repo set                                   # One-time setup for Linux, Pigsty + PGDG repos (overwrites!)
pig install -v 18 -y pg18 pg_duckdb vector     # Install PG 18 kernel, pg_duckdb, pgvector extensions...
```


## Installation

You can [**install**](/pig/install/) `pig` with the following command:

**Global** (Cloudflare CDN):

```bash
curl -fsSL https://repo.pigsty.io/pig | bash
```

**China Mainland**:

```bash
curl -fsSL https://repo.pigsty.cc/pig | bash
```


PIG binary is about 5 MB. On Linux it uses `rpm` or `dpkg` to install the latest available version:

```bash
$ curl -fsSL https://repo.pigsty.io/pig | bash
[INFO] kernel = Linux
[INFO] machine = x86_64
[INFO] package = deb
[INFO] pkg_url = https://repo.pigsty.io/pkg/pig/v1.5.0/pig_1.5.0-1_amd64.deb
[INFO] download = /tmp/pig_1.5.0-1_amd64.deb
[INFO] downloading pig v1.5.0
curl -fSL https://repo.pigsty.io/pkg/pig/v1.5.0/pig_1.5.0-1_amd64.deb -o /tmp/pig_1.5.0-1_amd64.deb
######################################################################## 100.0%
[INFO] md5sum = b7d18865090df8c03523c79ec2744e91
[INFO] installing: dpkg -i /tmp/pig_1.5.0-1_amd64.deb
(Reading database ... 166001 files and directories currently installed.)
Preparing to unpack /tmp/pig_1.5.0-1_amd64.deb ...
Unpacking pig (1.5.0-1) ...
Setting up pig (1.5.0-1) ...
[INFO] pig v1.5.0 installed successfully
check https://pgext.cloud for details
```



## Check Environment

PIG is a Go-written binary program, installed by default at `/usr/bin/pig`. `pig version` prints version information:

```bash
$ pig version
pig version 1.5.0 linux/amd64
```

Use `pig status` to print the current environment status, OS code, PG installation status, repository accessibility and latency.

```bash
$ pig status

# [Configuration] ================================
Pig Version      : 1.5.0
Pig Config       : /home/vagrant/.pig/config.yml
Log Level        : info
Log Path         : stderr

# [OS Environment] ===============================
OS Distro Code   : u26
OS OSArch        : arm64
OS Package Type  : deb
OS Vendor ID     : ubuntu
OS Version       : 26
OS Version Full  : 26.04
OS Version Code  : resolute

# [PG Environment] ===============================
Installed:
* PostgreSQL 18.4 (Ubuntu 18.4-1.pgdg26.04+1)  72  Extensions

Active:
PG Version      :  PostgreSQL 18.4 (Ubuntu 18.4-1.pgdg26.04+1)
Config Path     :  /usr/lib/postgresql/18/bin/pg_config
Binary Path     :  /usr/lib/postgresql/18/bin
Library Path    :  /usr/lib/postgresql/18/lib
Extension Path  :  /usr/share/postgresql/18/extension

# [Pigsty Environment] ===========================
Inventory Path   : Not Found
Pigsty Home      : Not Found

# [Network Conditions] ===========================
pigsty.cc  ping ok: 376 ms
pigsty.io  ping ok: 1270 ms
Internet Access   :  true
Pigsty Repo       :  pigsty.io
Inferred Region   :  china
Latest Pigsty Ver :  v4.4.0
```

## Automation Tips

For production recovery tasks, it is recommended to run `--plan` first to preview the PITR execution plan before actually executing:

```bash
pig pitr -d --plan            # preview recovery steps only, no execution
pig pitr -d -y                # skip confirmation (automation scenarios)
```


## List Extensions

Use the `pig ext list` command to print the built-in PG extension catalog.

```bash
$ pig ext list
✓ Found 531 extensions
Name                Status     Version     Cate   Flags   License         Repo     PGVer  Package                               Description
────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
timescaledb         available  2.28.2      TIME   -dslt-  Timescale       PIGSTY   15-18  postgresql-18-timescaledb-tsl         Enables scalable inserts and complex queries for time-series dat
timescaledb_toolkit available  1.23.0      TIME   -ds---  Timescale       PIGSTY   15-18  postgresql-18-timescaledb-toolkit     Library of analytical hyperfunctions, time-series pipelining, an
timeseries          available  0.2.1       TIME   -d----  PostgreSQL      PIGSTY   14-18  postgresql-18-pg-timeseries           Convenience API for time series stack
periods             available  1.2.3       TIME   -ds---  PostgreSQL      PGDG     14-18  postgresql-18-periods                 Provide Standard SQL functionality for PERIODs and SYSTEM VERSIO
........
pg_bulkload         available  3.1.23      ETL    bds---  BSD 3-Clause    PIGSTY   14-18  postgresql-18-pg-bulkload             pg_bulkload is a high speed data loading utility for PostgreSQL
test_decoding       available  -           ETL    --s--x  PostgreSQL      CONTRIB  14-18  postgresql-18                         SQL-based test/example module for WAL logical decoding
pgoutput            available  -           ETL    --s---  PostgreSQL      CONTRIB  14-18  postgresql-18                         Logical Replication output plugin

(531 Rows)
```

All extension metadata is defined in a data file named [`extension.csv`](https://github.com/pgsty/pig/blob/main/cli/ext/assets/extension.csv).
This file is updated with each pig release. You can update it directly using the [`pig ext reload`](/pig/cmd/ext/#ext-reload) command.
The updated file is placed in `~/.pig/extension.csv` by default. You can view and modify it, and the latest online catalog is available at [**pigsty.io/ext/data/extension.csv**](https://pigsty.io/ext/data/extension.csv).




## Add Repositories

To install extensions, you first need to add upstream repositories. [`pig repo`](/pig/cmd/repo/) can be used to manage Linux APT/YUM/DNF software repository configuration.

You can use the straightforward [`pig repo set`](/pig/cmd/repo/#repo-set) to overwrite existing repository configuration, ensuring only necessary repositories exist in the system:

```bash
pig repo set                # One-time setup for all repos including Linux system, PGDG, PIGSTY (PGSQL+INFRA)
```

> **Warning**: `pig repo set` will back up and clear existing repository configuration, then add required repositories with overwrite semantics.


Or choose the gentler [`pig repo add`](/pig/cmd/repo/#repo-add) to add needed repositories:

```bash
pig repo add pgdg pigsty     # Add PGDG official repo and PIGSTY supplementary repo
pig repo add pgsql           # [Optional] Add PGDG and PIGSTY together as one "pgsql" module
pig repo update              # Update cache: apt update / yum makecache
```

PIG detects your network environment and chooses Cloudflare global CDN or China cloud CDN, but you can force a specific region with `--region`:

```bash
pig repo set      --region=china              # use China mirror for faster downloads
pig repo add pgdg --region=default --update   # force PGDG upstream repo
```

PIG does not support offline installation. You can download RPM/DEB packages yourself and copy them to isolated servers for installation.
The related PIGSTY project provides local software repositories. You can use pig to install pre-downloaded extensions from local repos.



## Install PG

After adding repositories, you can use [`pig ext add`](/pig/cmd/ext/#ext-add) to install extensions (and related packages):

```bash
pig ext add -v 18 -y pgsql timescaledb postgis vector pg_duckdb pg_mooncake # install PG 18 kernel and extensions, auto-confirm

# This command will translate aliases to actual packages
INFO[20:34:44] translate alias 'pgsql' to package: postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl
INFO[20:34:44] translate extension 'timescaledb' to package: timescaledb-tsl_18
INFO[20:34:44] translate extension 'postgis' to package: postgis36_18
INFO[20:34:44] translate extension 'vector' to package: pgvector_18
INFO[20:34:44] translate extension 'pg_duckdb' to package: pg_duckdb_18
INFO[20:34:44] translate extension 'pg_mooncake' to package: pg_mooncake_18
INFO[20:34:44] installing packages: dnf install -y postgresql18 postgresql18-server postgresql18-libs postgresql18-contrib postgresql18-plperl postgresql18-plpython3 postgresql18-pltcl timescaledb-tsl_18 postgis36_18 pgvector_18 pg_duckdb_18 pg_mooncake_18
```

This uses the "alias translation" mechanism to map clean PG kernel/extension logical names into real RPM/DEB lists. If you do not need translation, use `apt/dnf` directly,
or use the `-n|--no-translation` option with the `pig install` variant:

```bash
pig install vector     # with translation, installs pgvector_18 or postgresql-18-pgvector for current PG 18
pig install vector -n  # no translation, installs a component named vector (from pigsty-infra repo)
```




## Alias Translation

PostgreSQL kernels and extensions map to many RPM/DEB packages. Remembering them is painful, so pig provides common aliases to simplify installation.

For example, on EL systems the following aliases translate to the RPM lists on the right:

```yaml
pgsql:        "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl"
pg18:         "postgresql18 postgresql18-server postgresql18-libs postgresql18-contrib postgresql18-plperl postgresql18-plpython3 postgresql18-pltcl"
pg18-client:  "postgresql18"
pg18-server:  "postgresql18-server postgresql18-libs postgresql18-contrib"
pg18-devel:   "postgresql18-devel"
pg18-basic:   "pg_repack_18 wal2json_18 pgvector_18"
pg17-mini:    "postgresql17 postgresql17-server postgresql17-libs postgresql17-contrib"
pg16-full:    "postgresql16 postgresql16-server postgresql16-libs postgresql16-contrib postgresql16-plperl postgresql16-plpython3 postgresql16-pltcl postgresql16-llvmjit postgresql16-test postgresql16-devel"
pg15-main:    "postgresql15 postgresql15-server postgresql15-libs postgresql15-contrib postgresql15-plperl postgresql15-plpython3 postgresql15-pltcl pg_repack_15 wal2json_15 pgvector_15"
pg14-core:    "postgresql14 postgresql14-server postgresql14-libs postgresql14-contrib postgresql14-plperl postgresql14-plpython3 postgresql14-pltcl"
```

Note the `$v` placeholder is replaced by the PG major version. When you use the `pgsql` alias, `$v` becomes 18, 17, etc.
So when you install the `pg18-server` alias, EL actually installs `postgresql18-server`, `postgresql18-libs`, `postgresql18-contrib`, while Debian/Ubuntu installs `postgresql-18`. Pig handles all details.

<details>
<summary>Common PostgreSQL Aliases</summary>

[Alias translation list for EL](https://github.com/pgsty/pig/blob/main/cli/ext/catalog.go#L206)

```bash
"pgsql":        "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl",
"pgsql-mini":   "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib",
"pgsql-core":   "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl",
"pgsql-full":   "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit postgresql$v-test postgresql$v-devel",
"pgsql-main":   "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl pg_repack_$v wal2json_$v pgvector_$v",
"pgsql-client": "postgresql$v",
"pgsql-server": "postgresql$v-server postgresql$v-libs postgresql$v-contrib",
"pgsql-devel":  "postgresql$v-devel",
"pgsql-basic":  "pg_repack_$v wal2json_$v pgvector_$v",
```

[Alias translation list for Debian/Ubuntu](https://github.com/pgsty/pig/blob/main/cli/ext/catalog.go#L270)

```bash
"pgsql":        "postgresql-$v postgresql-client-$v postgresql-plpython3-$v postgresql-plperl-$v postgresql-pltcl-$v",
"pgsql-mini":   "postgresql-$v postgresql-client-$v",
"pgsql-core":   "postgresql-$v postgresql-client-$v postgresql-plpython3-$v postgresql-plperl-$v postgresql-pltcl-$v",
"pgsql-full":   "postgresql-$v postgresql-client-$v postgresql-plpython3-$v postgresql-plperl-$v postgresql-pltcl-$v postgresql-server-dev-$v",
"pgsql-main":   "postgresql-$v postgresql-client-$v postgresql-plpython3-$v postgresql-plperl-$v postgresql-pltcl-$v postgresql-$v-repack postgresql-$v-wal2json postgresql-$v-pgvector",
"pgsql-client": "postgresql-client-$v",
"pgsql-server": "postgresql-$v",
"pgsql-devel":  "postgresql-server-dev-$v",
"pgsql-basic":  "postgresql-$v-repack postgresql-$v-wal2json postgresql-$v-pgvector",
```

These aliases can be instantiated with major versions, or you can use versioned aliases like `pg18`, `pg17`, and so on.
The actively supported PostgreSQL major versions are now **14-18**. For example, for PostgreSQL 18 you can use:

| `pgsql`        | `pg18`        | `pg17`        | `pg16`        | `pg15`        | `pg14`        |
|:---------------|:--------------|:--------------|:--------------|:--------------|:--------------|
| `pgsql`        | **`pg18`**    | `pg17`        | `pg16`        | `pg15`        | `pg14`        |
| `pgsql-mini`   | `pg18-mini`   | `pg17-mini`   | `pg16-mini`   | `pg15-mini`   | `pg14-mini`   |
| `pgsql-core`   | `pg18-core`   | `pg17-core`   | `pg16-core`   | `pg15-core`   | `pg14-core`   |
| `pgsql-full`   | `pg18-full`   | `pg17-full`   | `pg16-full`   | `pg15-full`   | `pg14-full`   |
| `pgsql-main`   | `pg18-main`   | `pg17-main`   | `pg16-main`   | `pg15-main`   | `pg14-main`   |
| `pgsql-client` | `pg18-client` | `pg17-client` | `pg16-client` | `pg15-client` | `pg14-client` |
| `pgsql-server` | `pg18-server` | `pg17-server` | `pg16-server` | `pg15-server` | `pg14-server` |
| `pgsql-devel`  | `pg18-devel`  | `pg17-devel`  | `pg16-devel`  | `pg15-devel`  | `pg14-devel`  |
| `pgsql-basic`  | `pg18-basic`  | `pg17-basic`  | `pg16-basic`  | `pg15-basic`  | `pg14-basic`  |
{.full-width}

</details>



## Install Extensions

Pig detects your PostgreSQL installation. If there is an active PG installation (detected via `pg_config` in `PATH`), pig installs extensions for that PG major by default.

```bash
pig install pg_smtp_client          # simplest
pig install pg_smtp_client -v 18    # specify major version (more stable)
pig install pg_smtp_client -p /usr/lib/postgresql/16/bin/pg_config   # another way to target PG

dnf install pg_smtp_client_18       # most direct, but not all extensions are that simple
```

Tip: to add a specific PG major version into `PATH`, use `pig ext link`:

```bash
pig ext link pg18             # create /usr/pgsql symlink and write /etc/profile.d/pgsql.sh
. /etc/profile.d/pgsql.sh     # take effect now and update PATH
```

If you want a specific package version, use `name=ver` syntax:

```bash
pig ext add -v 18 pgvector=0.7.2 # install pgvector 0.7.2 for PG 18
pig ext add pg16=16.5            # install PostgreSQL 16 with a specific minor version
```

> **Warning**: currently only PGDG YUM repositories provide historical extension versions. PIGSTY repo and PGDG APT repo only provide the **latest** extension versions.




## Show Extensions

[`pig ext status`](/pig/cmd/ext/#ext-status) shows installed extensions.

```bash
$ pig ext status

✓ PostgreSQL 18: 72 extensions installed (3 shown)
PostgreSQL 18: PostgreSQL 18.4 (Ubuntu 18.4-1.pgdg26.04+1)
  Binary: /usr/lib/postgresql/18/bin
  Extension: /usr/share/postgresql/18/extension

Extension Stat: 3 Installed (PGDG 3, PIGSTY 0) + 69 CONTRIB = 72 Total

Name       Version  Cate   Flags   License       Repo  Package                 Description
───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
pg_repack  1.5.3    ADMIN  bds---  PostgreSQL    PGDG  postgresql-18-repack    Reorganize tables in PostgreSQL databases with minimal locks
vector     0.8.4    RAG    -ds--r  PostgreSQL    PGDG  postgresql-18-pgvector  vector data type and ivfflat and hnsw access methods
wal2json   2.6      ETL    --s--x  BSD 3-Clause  PGDG  postgresql-18-wal2json  Changing data capture in JSON format
```

If PostgreSQL cannot be found in your current `PATH` (via `pg_config`), it is recommended to explicitly specify PG major with `-v|-p` to avoid version detection ambiguity.




## Scan Extensions

[`pig ext scan`](/pig/cmd/ext/#ext-scan) provides a lower-level scan. It scans shared libraries under the target PG directory to discover installed extensions:

```bash
$ pig ext scan

Installed:
- PostgreSQL 18.4 (Ubuntu 18.4-1.pgdg26.04+1)  72 Extensions

Active:
PG Version      :  PostgreSQL 18.4 (Ubuntu 18.4-1.pgdg26.04+1)
Config Path     :  /usr/lib/postgresql/18/bin/pg_config
Binary Path     :  /usr/lib/postgresql/18/bin
Library Path    :  /usr/lib/postgresql/18/lib
Extension Path  :  /usr/share/postgresql/18/extension

Name       Version  SharedLibs      Description                                                        Meta
----       -------  ----------      -----------                                                        ----
pg_repack  1.5.3    pg_repack.so    Reorganize tables in PostgreSQL databases with minimal locks        module_pathname=$libdir/pg_repack relocatable=false lib=pg_repack.so
vector     0.8.4    vector.so       vector data type and ivfflat and hnsw access methods                relocatable=true lib=vector.so
wal2json   2.6      wal2json.so     Changing data capture in JSON format                                lib=wal2json.so
...
```



## Container Practice

You can create a new VM or use the following Docker container for testing. Create a `d13` directory and a `Dockerfile`:

```dockerfile
FROM debian:13
USER root
WORKDIR /root/
CMD ["/bin/bash"]

RUN apt update && apt install -y ca-certificates curl && curl https://repo.pigsty.io/pig | bash
```

```bash
docker build -t d13:latest .
docker run -it d13:latest /bin/bash

pig repo set --region=china    # add China region repos
pig install -y pg18            # install PGDG 18 kernel packages
pig install -y postgis timescaledb pgvector pg_duckdb
```
