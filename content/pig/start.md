---
title: Get Started
description: Quick start guide for PIG, the PostgreSQL package manager
icon: Play
weight: 200
breadcrumbs: false
---

Here's a simple tutorial to experience the core capabilities of PIG package manager.

## Quick Start

```bash
curl -fsSL https://repo.pigsty.io/pig | bash   # Install PIG from Cloudflare
pig repo set                                   # Configure Linux, Pigsty + PGDG repos (overwrite mode!)
pig install -v 18 -y pg18 pg_duckdb vector     # Install PG 18 core, pg_duckdb, pgvector extensions...
```


## Installation

[**Install**](/pig/install/) `pig` with a single command:

{{< tabs items="Default,Mirror" defaultIndex="1" >}}
{{< tab >}}
```bash
curl -fsSL https://repo.pigsty.io/pig | bash     # Install from Cloudflare
```
{{< /tab >}}
{{< tab >}}
```bash
curl -fsSL https://repo.pigsty.cc/pig | bash     # Install from China CDN mirror
```
{{< /tab >}}
{{< /tabs >}}

The PIG binary is approximately 4 MB and will automatically install the latest available version using `rpm` or `dpkg` on Linux:

```bash
[INFO] kernel = Linux
[INFO] machine = x86_64
[INFO] package = rpm
[INFO] pkg_url = https://repo.pigsty.io/pkg/pig/v1.0.0/pig-1.0.0-1.x86_64.rpm
[INFO] download = /tmp/pig-1.0.0-1.x86_64.rpm
[INFO] downloading pig v1.0.0
curl -fSL https://repo.pigsty.io/pkg/pig/v1.0.0/pig-1.0.0-1.x86_64.rpm -o /tmp/pig-1.0.0-1.x86_64.rpm
######################################################################## 100.0%
[INFO] md5sum = 85d75c16dfd3ce935d9d889fae345430
[INFO] installing: rpm -ivh /tmp/pig-1.0.0-1.x86_64.rpm
Verifying...                          ################################# [100%]
Preparing...                          ################################# [100%]
Updating / installing...
   1:pig-1.0.0-1                      ################################# [100%]
[INFO] pig v1.0.0 installed successfully
check https://pgext.cloud for details
```



## Environment Check

PIG is a Go binary, installed by default at `/usr/bin/pig`. Use `pig version` to display version information:

```bash
$ pig version

pig version 1.0.0 linux/amd64
build: HEAD 35f3aac 2026-01-18T00:00:00Z
```

Use `pig status` to display the current environment status, OS code, PostgreSQL installation, repository accessibility and latency.

```bash
$ pig status

# [Configuration] ================================
Pig Version      : 1.0.0
Pig Config       : /root/.pig/config.yml
Log Level        : info
Log Path         : stderr

# [OS Environment] ===============================
OS Distro Code   : el10
OS OSArch        : amd64
OS Package Type  : rpm
OS Vendor ID     : rocky
OS Version       : 10
OS Version Full  : 10.0
OS Version Code  : el10

# [PG Environment] ===============================
No PostgreSQL installation found

No active PostgreSQL found in PATH:
- /root/.local/bin
- /root/bin
- /usr/local/sbin
- /usr/local/bin
- /usr/sbin
- /usr/bin

# [Pigsty Environment] ===========================
Inventory Path   : Not Found
Pigsty Home      : Not Found

# [Network Conditions] ===========================
pigsty.cc  ping ok: 612 ms
pigsty.io  ping ok: 1222 ms
google.com request error
Internet Access   :  true
Pigsty Repo       :  pigsty.io
Inferred Region   :  china
Latest Pigsty Ver :  v3.6.1
```


## List Extensions

Use `pig ext list` to display the built-in PostgreSQL extension catalog.

```bash
$ pig ext list
Name                            Status              Version     Cate   Flags   License       Repo     PGVer  Package                    Description
----                            ------              -------     ----   ------  -------       ------   -----  ------------               ---------------------
timescaledb                     installed  2.24.0      TIME   -dsl--  Timescale     PIGSTY   15-18  timescaledb-tsl_18         Enables scalable inserts and complex queries for time-series dat
timescaledb_toolkit             installed  1.22.0      TIME   -ds-t-  Timescale     PIGSTY   15-18  timescaledb-toolkit_18     Library of analytical hyperfunctions, time-series pipelining, an
timeseries                      installed  0.2.0       TIME   -d----  PostgreSQL    PIGSTY   13-18  pg_timeseries_18           Convenience API for time series stack
periods                         installed  1.2.3       TIME   -ds---  PostgreSQL    PGDG     13-18  periods_18                 Provide Standard SQL functionality for PERIODs and SYSTEM VERSIO
temporal_tables                 installed  1.2.2       TIME   -ds--r  BSD 2-Clause  PIGSTY   13-18  temporal_tables_18         temporal tables
...
pg_bulkload                     installed  3.1.23      ETL    bds---  BSD 3-Clause  PGDG     13-18  pg_bulkload_18             pg_bulkload is a high speed data loading utility for PostgreSQL
test_decoding                   available  -           ETL    --s--x  PostgreSQL    CONTRIB  13-18  postgresql18-contrib       SQL-based test/example module for WAL logical decoding
pgoutput                        available  -           ETL    --s---  PostgreSQL    CONTRIB  13-18  postgresql18-contrib       Logical Replication output plugin

(451 Rows) (Status: installed, available, not avail | Flags: b = HasBin, d = HasDDL, s = HasLib, l = NeedLoad, t = Trusted, r = Relocatable, x = Unknown)

```

All extension metadata is defined in a file called [`extension.csv`](https://github.com/pgsty/pig/blob/main/cli/ext/assets/extension.csv),
which is continuously updated with pig releases. You can update this data file directly using [`pig ext reload`](/pig/cmd/ext#pig-ext-reload).
The updated file is placed at `~/.pig/extension.csv` by default, where you can review and modify it. The [**authoritative version**](https://github.com/pgsty/pgext/blob/main/db/extension.csv) of this data file can be found in this project.




## Add Repositories

To install extensions, you first need to add upstream repositories. [`pig repo`](/pig/cmd/repo) manages Linux APT/YUM/DNF repository configurations.

You can use the straightforward [`pig repo set`](/pig/cmd/repo#repo-set) to overwrite existing repository configurations, ensuring only necessary repositories exist:

```bash
pig repo set                # Configure all repos at once: Linux, PGDG, PIGSTY (PGSQL+INFRA) repositories
```

{{< callout type="warning" >}}
`pig repo set` will backup and clean existing repository configurations, then add required repositories with overwrite semantics. Use with caution!
{{< /callout >}}


Or choose the gentler [`pig repo add`](/pig/cmd/repo#repo-add) to add required repositories:

```bash
pig repo add pgdg pigsty     # Add official PGDG and PIGSTY supplementary repositories
pig repo add pgsql           # [Optional] Add PGDG and PIGSTY combined as a single "pgsql" module
pig repo update              # Update cache: apt update / yum makecache
```

PIG detects your network environment and chooses between Cloudflare global CDN or China mainland CDN, but you can force a region using the `--region` parameter.

```bash
pig repo set      --region=china              # Use China region mirror (pigsty/aliyun) for faster downloads
pig repo set      --region=europe             # Use Europe region mirror (xtom) for faster downloads
pig repo add pgdg --region=default --update   # Force using upstream PGDG repository
```

PIG doesn't support offline installation. You can download RPM/DEB packages yourself and copy them to network-isolated production servers.
The related PIGSTY project provides local repositories, supporting extension installation from local repositories using pig.



## Install PostgreSQL

After adding repositories, use [`pig ext add`](/pig/cmd/ext#ext-add) to install extensions (and related packages):

```bash
pig ext add -v 18 -y pgsql timescaledb postgis vector pg_duckdb pg_mooncake # Install PG 18 core and extensions with auto-confirm

# The command automatically performs translation, mapping packages to:
INFO[20:34:44] translate alias 'pgsql' to package: postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl
INFO[20:34:44] translate extension 'timescaledb' to package: timescaledb-tsl_18
INFO[20:34:44] translate extension 'postgis' to package: postgis36_18
INFO[20:34:44] translate extension 'vector' to package: pgvector_18
INFO[20:34:44] translate extension 'pg_duckdb' to package: pg_duckdb_18
INFO[20:34:44] translate extension 'pg_mooncake' to package: pg_mooncake_18
INFO[20:34:44] installing packages: dnf install -y postgresql18 postgresql18-server postgresql18-libs postgresql18-contrib postgresql18-plperl postgresql18-plpython3 postgresql18-pltcl timescaledb-tsl_18 postgis36_18 pgvector_18 pg_duckdb_18 pg_mooncake_18
```

This uses an "alias translation" mechanism, translating clean PostgreSQL core/extension logical names to actual RPM/DEB packages. To install without alias translation, use `apt/dnf` directly,
or use the `pig install` variant with `-n|--no-translation`:

```bash
pig install vector     # With translation, installs pgvector_18 or postgresql-18-pgvector for PG 18
pig install vector -n  # Without translation, installs a logging component named vector (from pigsty-infra repo)
```




## Alias Translation

PostgreSQL core and extensions correspond to various RPM/DEB packages. Remembering these packages is cumbersome, so pig provides many useful aliases to simplify installation:

For example on EL systems, the following aliases translate to the corresponding RPM packages on the right:

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

Note the `$v` placeholder is replaced with the PostgreSQL major version. When using the `pgsql` alias, `$v` is substituted with actual major versions like 18, 17.
So when you install the `pg18-server` alias, on EL you're actually installing `postgresql18-server`, `postgresql18-libs`, `postgresql18-contrib`, and on Debian/Ubuntu you're installing `postgresql-18`. PIG handles all the details.

<br>
<details><summary> Common PostgreSQL Aliases</summary>

[Alias translation list for EL](https://github.com/pgsty/pig/blob/main/cli/ext/catalog.go#L206)

```bash {base_url="https://github.com/pgsty/pig/blob/main/",filename="cli/ext/catalog.go"}
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

[Alias translation for Debian/Ubuntu](https://github.com/pgsty/pig/blob/main/cli/ext/catalog.go#L270)

```bash {base_url="https://github.com/pgsty/pig/blob/main/",filename="cli/ext/catalog.go"}
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

These aliases can be used directly and instantiated with parameters for major versions, or you can use variants with specific major version numbers: replace `pgsql` with `pg18`, `pg17`, `pgxx`, etc.
For example, for PostgreSQL 18, you can use these aliases directly:

| `pgsql`        | `pg18`        | `pg17`        | `pg16`        | `pg15`        | `pg14`        | `pg13`        |
|:---------------|:--------------|:--------------|:--------------|:--------------|:--------------|:--------------|
| `pgsql`        | **`pg18`**    | `pg17`        | `pg16`        | `pg15`        | `pg14`        | `pg13`        |
| `pgsql-mini`   | `pg18-mini`   | `pg17-mini`   | `pg16-mini`   | `pg15-mini`   | `pg14-mini`   | `pg13-mini`   |
| `pgsql-core`   | `pg18-core`   | `pg17-core`   | `pg16-core`   | `pg15-core`   | `pg14-core`   | `pg13-core`   |
| `pgsql-full`   | `pg18-full`   | `pg17-full`   | `pg16-full`   | `pg15-full`   | `pg14-full`   | `pg13-full`   |
| `pgsql-main`   | `pg18-main`   | `pg17-main`   | `pg16-main`   | `pg15-main`   | `pg14-main`   | `pg13-main`   |
| `pgsql-client` | `pg18-client` | `pg17-client` | `pg16-client` | `pg15-client` | `pg14-client` | `pg13-client` |
| `pgsql-server` | `pg18-server` | `pg17-server` | `pg16-server` | `pg15-server` | `pg14-server` | `pg13-server` |
| `pgsql-devel`  | `pg18-devel`  | `pg17-devel`  | `pg16-devel`  | `pg15-devel`  | `pg14-devel`  | `pg13-devel`  |
| `pgsql-basic`  | `pg18-basic`  | `pg17-basic`  | `pg16-basic`  | `pg15-basic`  | `pg14-basic`  | `pg13-basic`  |

</details>



## Install Extensions

PIG detects PostgreSQL installations in your system environment. If an active PG installation is detected (based on `pg_config` in `PATH`), pig automatically installs extensions for that PostgreSQL major version without explicit version specification.

```bash
pig install pg_smtp_client          # Simpler
pig install pg_smtp_client -v 18    # Explicitly specify major version, more reliable
pig install pg_smtp_client -p /usr/lib/postgresql/16/bin/pg_config   # Another way to specify PG version
dnf install pg_smtp_client_18       # Most direct... but not all extensions are this simple...
```

Tip: To add a specific PostgreSQL major version binaries to `PATH`, use `pig ext link`:

```bash
pig ext link pg18             # Create /usr/pgsql symlink and write to /etc/profile.d/pgsql.sh
. /etc/profile.d/pgsql.sh     # Take effect immediately, update PATH environment variable
```

To install a specific software version, use `name=ver` syntax:

```bash
pig ext add -v 17 pgvector=0.7.2 # install pgvector 0.7.2 for PG 17
pig ext add pg16=16.5            # install PostgreSQL 16 with a specific minor version
```

{{< callout type="warning" >}}
Note that currently only PGDG YUM repository provides historical extension versions. PIGSTY repository and PGDG APT repository only provide the **latest version** of extensions.
{{< /callout >}}




## Display Extensions

[`pig ext status`](/pig/cmd/ext#ext-status) displays currently installed extensions.

```bash
$ pig ext status -v 18

Installed:
- PostgreSQL 18.0  80  Extensions

No active PostgreSQL found in PATH:
- /root/.local/bin
- /root/bin
- /usr/local/sbin
- /usr/local/bin
- /usr/sbin
- /usr/bin
Extension Stat  :  11 Installed (PIGSTY 3, PGDG 8) + 69 CONTRIB = 80 Total

Name                          Version  Cate  Flags   License     Repo    Package              Description
----                          -------  ----  ------  -------     ------  ------------         ---------------------
timescaledb                   2.23.0   TIME  -dsl--  Timescale   PIGSTY  timescaledb-tsl_18*  Enables scalable inserts and complex queries for time-series dat
postgis                       3.6.0    GIS   -ds---  GPL-2.0     PGDG    postgis36_18*        PostGIS geometry and geography spatial types and functions
postgis_topology              3.6.0    GIS   -ds---  GPL-2.0     PGDG    postgis36_18*        PostGIS topology spatial types and functions
postgis_raster                3.6.0    GIS   -ds---  GPL-2.0     PGDG    postgis36_18*        PostGIS raster types and functions
postgis_sfcgal                3.6.0    GIS   -ds--r  GPL-2.0     PGDG    postgis36_18*        PostGIS SFCGAL functions
postgis_tiger_geocoder        3.6.0    GIS   -ds-t-  GPL-2.0     PGDG    postgis36_18*        PostGIS tiger geocoder and reverse geocoder
address_standardizer          3.6.0    GIS   -ds--r  GPL-2.0     PGDG    postgis36_18*        Used to parse an address into constituent elements. Generally us
address_standardizer_data_us  3.6.0    GIS   -ds--r  GPL-2.0     PGDG    postgis36_18*        Address Standardizer US dataset example
vector                        0.8.1    RAG   -ds--r  PostgreSQL  PGDG    pgvector_18*         vector data type and ivfflat and hnsw access methods
pg_duckdb                     1.1.0    OLAP  -dsl--  MIT         PIGSTY  pg_duckdb_18*        DuckDB Embedded in Postgres
pg_mooncake                   0.2.0    OLAP  -d----  MIT         PIGSTY  pg_mooncake_18*      Columnstore Table in Postgres
```

If PostgreSQL cannot be found in your current system PATH (based on `pg_config` in `PATH`), make sure to specify the PG major version or `pg_config` path via `-v|-p`.



## Scan Extensions

[`pig ext scan`](/pig/cmd/ext#ext-scan) provides lower-level extension scanning, scanning shared libraries in specified PostgreSQL directories to discover installed extensions:

```bash
root@s37451:~# pig ext scan
Installed:
* PostgreSQL 17.6 (Debian 17.6-2.pgdg13+1)    70  Extensions
- PostgreSQL 15.14 (Debian 15.14-1.pgdg13+1)  69  Extensions
- PostgreSQL 14.19 (Debian 14.19-1.pgdg13+1)  66  Extensions
- PostgreSQL 13.22 (Debian 13.22-1.pgdg13+1)  64  Extensions
- PostgreSQL 18.0 (Debian 18.0-1.pgdg13+3)    70  Extensions
- PostgreSQL 16.10 (Debian 16.10-1.pgdg13+1)  70  Extensions

Active:
PG Version      :  PostgreSQL 17.6 (Debian 17.6-2.pgdg13+1)
Config Path     :  /usr/lib/postgresql/17/bin/pg_config
Binary Path     :  /usr/lib/postgresql/17/bin
Library Path    :  /usr/lib/postgresql/17/lib
Extension Path  :  /usr/share/postgresql/17/extension
Name                 Version  SharedLibs                                       Description            Meta
----                 -------  ----------                                       ---------------------  ------
amcheck              1.4      functions for verifying relation integrity       relocatable=true module_pathname=$libdir/amcheck lib=amcheck.so
...
pg_duckdb            1.1.0    DuckDB Embedded in Postgres                      module_pathname=$libdir/pg_duckdb relocatable=false schema=public lib=libduckdb.so, pg_duckdb.so
pg_mooncake          0.2.0    Real-time analytics on Postgres tables           module_pathname=pg_mooncake relocatable=false requires=pg_duckdb superuser=true lib=pg_mooncake.so
pg_prewarm           1.2      prewarm relation data                            module_pathname=$libdir/pg_prewarm relocatable=true lib=pg_prewarm.so
pg_smtp_client       0.2.1    PostgreSQL extension to send email using SMTP    relocatable=false superuser=false schema=smtp_client module_pathname=$libdir/pg_smtp_client lib=pg_smtp_client.so
...
Encoding Libs: cyrillic_and_mic, euc2004_sjis2004, euc_cn_and_mic, euc_jp_and_sjis, euc_kr_and_mic, euc_tw_and_big5, latin2_and_win1250, latin_and_mic, utf8_and_big5, utf8_and_cyrillic, utf8_and_euc2004, utf8_and_euc_cn, utf8_and_euc_jp, utf8_and_euc_kr, utf8_and_euc_tw, utf8_and_gb18030, utf8_and_gbk, utf8_and_iso8859, utf8_and_iso8859_1, utf8_and_johab, utf8_and_sjis, utf8_and_sjis2004, utf8_and_uhc, utf8_and_win
Built-in Libs: dict_snowball, libpqwalreceiver, llvmjit
```



## Container Practice

You can create a fresh virtual machine or use the following Docker container for testing. Create a `d13` directory with a `Dockerfile`:

```dockerfile {filename="d13/Dockerfile"}
FROM debian:13
USER root
WORKDIR /root/
CMD ["/bin/bash"]

RUN apt update && apt install -y ca-certificates curl && curl https://repo.pigsty.io/pig | bash
```

```bash
docker build -t d13:latest .
docker run -it d13:latest /bin/bash

pig repo set --region=china    # Add China region repositories
pig install -y pg18            # Install PostgreSQL 18 core packages
pig install -y postgis timescaledb pgvector pg_duckdb
```
