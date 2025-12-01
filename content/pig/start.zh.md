---
title: 上手
description: 快速上手 pig，PostgreSQL 包管理器
icon: Play
weight: 200
breadcrumbs: false
---

下面是一个简单的上手教程，带您体验 PIG 包管理器的核心能力。

## 简短版本

```bash
curl -fsSL https://repo.pigsty.io/pig | bash   # 从 Cloudflare 安装 PIG
pig repo set                                   # 一次性设置好 Linux, Pigsty + PGDG 仓库（覆盖式！）
pig install -v 18 -y pg18 pg_duckdb vector     # 安装 PG 18 内核，pg_duckdb, pgvector 扩展…… 
```


## 安装

您可以使用以下命令 [**一键安装**](/zh/pig/install/) `pig`：

{{< tabs items="默认,镜像" defaultIndex="1" >}}
{{< tab >}}
```bash
curl -fsSL https://repo.pigsty.io/pig | bash     # 从 Cloudflare 安装
```
{{< /tab >}}
{{< tab >}}
```bash
curl -fsSL https://repo.pigsty.cc/pig | bash     # 从中国 CDN 镜像站安装
```
{{< /tab >}}
{{< /tabs >}}

PIG 二进制包大约 4 MB，在 Linux 上会自动使用 `rpm` 或 `dpkg` 安装最新可用版本：

```bash
[INFO] kernel = Linux
[INFO] machine = x86_64
[INFO] package = rpm
[INFO] pkg_url = https://repo.pigsty.io/pkg/pig/v0.7.4/pig-0.7.4-1.x86_64.rpm
[INFO] download = /tmp/pig-0.7.2-1.x86_64.rpm
[INFO] downloading pig v0.7.2
curl -fSL https://repo.pigsty.io/pkg/pig/v0.7.2/pig-0.7.2-1.x86_64.rpm -o /tmp/pig-0.7.2-1.x86_64.rpm
######################################################################## 100.0%
[INFO] md5sum = 85d75c16dfd3ce935d9d889fae345430
[INFO] installing: rpm -ivh /tmp/pig-0.7.2-1.x86_64.rpm
Verifying...                          ################################# [100%]
Preparing...                          ################################# [100%]
Updating / installing...
   1:pig-0.7.2-1                      ################################# [100%]
[INFO] pig v0.7.2 installed successfully
check https://pgext.cloud for details
```



## 检查环境

PIG 是一个由 Go 编写的二进制程序，默认安装路径为 `/usr/bin/pig`，`pig version` 会打印版本信息：

```bash
$ pig version

pig version 0.7.2 linux/amd64
build: HEAD 9cdb57a 2025-11-10T11:14:17Z
```

使用 `pig status` 命令，会打印当前环境的状态，操作系统代码，PG的安装情况，仓库的可访问性与延迟。

```bash
$ pig status

# [Configuration] ================================
Pig Version      : 0.7.2
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


## 列出扩展

使用 `pig ext list` 命令，可以打印内置的 PG 扩展数据目录。

```bash
[root@pg-meta ~]# pig ext list

Name                            Version     Cate   Flags   License       RPM      DEB      PG Ver  Description
----                            -------     ----   ------  -------       ------   ------   ------  ---------------------
timescaledb                     2.23.0      TIME   -dsl--  Timescale     PIGSTY   PIGSTY   15-18   Enables scalable inserts and complex queries for time-series dat...
timescaledb_toolkit             1.22.0      TIME   -ds-t-  Timescale     PIGSTY   PIGSTY   15-18   Library of analytical hyperfunctions, time-series pipelining, an...
timeseries                      0.1.7       TIME   -d----  PostgreSQL    PIGSTY   PIGSTY   13-18   Convenience API for time series stack
periods                         1.2.3       TIME   -ds---  PostgreSQL    PGDG     PGDG     13-18   Provide Standard SQL functionality for PERIODs and SYSTEM VERSIO...
temporal_tables                 1.2.2       TIME   -ds--r  BSD 2-Clause  PIGSTY   PIGSTY   13-18   temporal tables
.........
pg_fact_loader                  2.0.1       ETL    -ds--x  MIT           PGDG     PGDG     13-18   build fact tables with Postgres
pg_bulkload                     3.1.22      ETL    bds---  BSD 3-Clause  PGDG     PIGSTY   13-17   pg_bulkload is a high speed data loading utility for PostgreSQL
test_decoding                   -           ETL    --s--x  PostgreSQL    CONTRIB  CONTRIB  13-18   SQL-based test/example module for WAL logical decoding
pgoutput                        -           ETL    --s---  PostgreSQL    CONTRIB  CONTRIB  13-18   Logical Replication output plugin

(431 Rows) (Flags: b = HasBin, d = HasDDL, s = HasLib, l = NeedLoad, t = Trusted, r = Relocatable, x = Unknown)
```

所有的扩展元数据都在一份名为 [`extension.csv`](https://github.com/pgsty/pig/blob/main/cli/ext/assets/extension.csv) 的数据文件中定义，
这份文件会随着 pig 版本发布不断更新，您可以直接使用 [`pig ext reload`](/zh/pig/cmd/ext#pig-ext-reload) 命令更新这份数据文件。
更新后的文件会默认放置于 `~/.pig/extension.csv` 中，您可以查阅与更改 —— 您也可以在本项目中找到该数据文件的 [**权威版本**](https://github.com/pgsty/pgext/blob/main/db/extension.csv)。




## 添加仓库

要想安装扩展，首先需要添加上游仓库。[`pig repo`](/zh/pig/cmd/repo) 可用于管理 Linux APT/YUM/DNF 软件仓库配置。

您可以使用简单粗暴直接的版本 [`pig repo set`](/zh/pig/cmd/repo#repo-set) 覆盖式写入现有仓库配置，该命令确保系统中只存在必须的仓库配置：

```bash
pig repo set                # 一次性配置好所有仓库，包括 Linux 系统仓库，PGDG，PIGSTY (PGSQL+INFRA) 仓库
```

{{< callout type="warning" >}}
`pig repo set` 会备份并清理现有的仓库配置，然后添加所需的仓库，实现 Overwrite 语义，请务必注意！
{{< /callout >}}


或者选择使用温和的 [`pig repo add`](/zh/pig/cmd/repo#repo-add) 添加所需的仓库：

```bash
pig repo add pgdg pigsty     # 添加 PGDG 官方仓库 和 PIGSTY 补充仓库 
pig repo add pgsql           # 【可选】您也可以选择将 PGDG 和 PIGSTY 合在一起，当成一个 “pgsql” 模块整体添加
pig repo update              # 更新缓存：apt update / yum makecache
```

PIG 会检测您的网络环境，并选择使用 Cloudflare 全球 CDN，或者中国境内云 CDN，但您可以通过 `--region` 参数强制指定区域。

```bash
pig repo set      --region=china              # 使用中国区域镜像仓库加速下载
pig repo add pgdg --region=default --update   # 强制指定使用 PGDG 上游仓库
```

PIG 本身不支持离线安装，您可以自行下载 RPM/DEB 包，拷贝到网络隔离的生产服务器安装。
相关项目 PIGSTY 提供本地软件仓库，支持可以使用 pig 从本地软件仓库安装已经下载好的扩展。



## 安装PG

添加仓库后，您可以使用 [`pig ext add`](/zh/pig/cmd/ext#ext-add) 子命令安装扩展（以及相关软件包）

```bash
pig ext add -v 18 -y pgsql timescaledb postgis vector pg_duckdb pg_mooncake # 安装 PG 18 内核与扩展，自动确认

# 该命令会自动执行翻译，将软件包翻译为
INFO[20:34:44] translate alias 'pgsql' to package: postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit
INFO[20:34:44] translate extension 'timescaledb' to package: timescaledb-tsl_18*
INFO[20:34:44] translate extension 'postgis' to package: postgis36_18*
INFO[20:34:44] translate extension 'vector' to package: pgvector_18*
INFO[20:34:44] translate extension 'pg_duckdb' to package: pg_duckdb_18*
INFO[20:34:44] translate extension 'pg_mooncake' to package: pg_mooncake_18*
INFO[20:34:44] installing packages: dnf install -y postgresql18 postgresql18-server postgresql18-libs postgresql18-contrib postgresql18-plperl postgresql18-plpython3 postgresql18-pltcl postgresql18-llvmjit timescaledb-tsl_18* postgis36_18* pgvector_18* pg_duckdb_18* pg_mooncake_18*
```

这里使用了 “别名翻译” 机制，将清爽的 PG 内核/扩展 逻辑包名翻译为实际的 RPM/DEB 列表。如果您不需要别名翻译机制，可以直接使用 `apt/dnf` 安装，
或者使用变体 `pig install` 的  `-n|--no-translation` 参数：

```bash
pig install vector     # 带有翻译机制，安装当前 PG 18 对应的 pgvector_18 或 postgresql-18-pgvector 
pig install vector -n  # 关闭翻译机制，安装名为 vector 的日志收集组件（来自 pigsty-infra 仓库）
```




## 别名翻译

PostgreSQL 内核与扩展对应着一系列的 RPM/DEB 包，记住这些包是一件麻烦事，所以 pig 提供了许多常用的别名，帮助您简化安装过程：

例如在 EL 系统上， 下面的别名将会被翻译为右侧的对应 RPM 包列表：

```yaml
pgsql:        "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit"
pg18:         "postgresql18 postgresql18-server postgresql18-libs postgresql18-contrib postgresql18-plperl postgresql18-plpython3 postgresql18-pltcl postgresql18-llvmjit"
pg17-client:  "postgresql17"
pg17-server:  "postgresql17-server postgresql17-libs postgresql17-contrib"
pg17-devel:   "postgresql17-devel"
pg17-basic:   "pg_repack_17* wal2json_17* pgvector_17*"
pg16-mini:    "postgresql16 postgresql16-server postgresql16-libs postgresql16-contrib"
pg15-full:    "postgresql15 postgresql15-server postgresql15-libs postgresql15-contrib postgresql15-plperl postgresql15-plpython3 postgresql15-pltcl postgresql15-llvmjit postgresql15-test postgresql15-devel"
pg14-main:    "postgresql14 postgresql14-server postgresql14-libs postgresql14-contrib postgresql14-plperl postgresql14-plpython3 postgresql14-pltcl postgresql14-llvmjit pg_repack_14* wal2json_14* pgvector_14*"
pg13-core:    "postgresql13 postgresql13-server postgresql13-libs postgresql13-contrib postgresql13-plperl postgresql13-plpython3 postgresql13-pltcl postgresql13-llvmjit"
```

注意这里的 `$v` 占位符会被替换为 PG 大版本号，因此当您使用 `pgsql` 别名时，`$v` 会被实际替代为 18，17 这样的大版本号。
因此，当您安装 `pg17-server` 别名时，EL 上实际安装的是 `postgresql17-server`, `postgresql17-libs`, `postgresql17-contrib`，在 Debian / Ubuntu 上安装的是 `postgresql-17` ，pig 会处理好所有细节。

<br>
<details><summary> 常用 PostgreSQL 别名</summary>

[EL 使用的别名翻译列表](https://github.com/pgsty/pig/blob/main/cli/ext/catalog.go#L154)

```bash {base_url="https://github.com/pgsty/pig/blob/main/",filename="cli/ext/catalog.go"}
"pgsql":        "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit",
"pgsql-mini":   "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib",
"pgsql-core":   "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit",
"pgsql-full":   "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit postgresql$v-test postgresql$v-devel",
"pgsql-main":   "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit pg_repack_$v* wal2json_$v* pgvector_$v*",
"pgsql-client": "postgresql$v",
"pgsql-server": "postgresql$v-server postgresql$v-libs postgresql$v-contrib",
"pgsql-devel":  "postgresql$v-devel",
"pgsql-basic":  "pg_repack_$v* wal2json_$v* pgvector_$v*",
```

[Debian / Ubuntu 系统使用的别名翻译](https://github.com/pgsty/pig/blob/main/cli/ext/catalog.go#L260)

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

上面这些别名可以直接使用，并通过参数实例化大版本号，也可以使用另一种带有大版本号的别名变体：即将 `pgsql` 替换为 `pg18`, `pg17`, `pgxx` 等具体大版本号。
例如，对于 PostgreSQL 18，可以直接使用下面这些别名：

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



## 安装扩展

pig 会检测当前系统环境中的 PostgreSQL 安装情况。如果检测到环境中（以 `PATH` 中的 `pg_config` 为准）有活跃的 PG 安装，那么 pig 会自动安装对应 PG 大版本所需的扩展，无需您显式指定 PG 大版本。

```bash
pig install pg_smtp_client          # 更简单
pig install pg_smtp_client -v 18    # 显示指定大版本，更稳定可靠
pig install pg_smtp_client -p /usr/lib/postgresql/16/bin/pg_config   # 另一种指定 PG 版本的方式
dnf install pg_smtp_client_18       # 最直接……，但并非所有扩展都这么简单……
```

提示：如需将特定大版本的 PostgreSQL 内核二进制加入 `PATH`，可用 `pig ext link` 命令：

```bash
pig ext link pg17             # 创建 /usr/pgsql 软链接，并写入 /etc/profile.d/pgsql.sh
. /etc/profile.d/pgsql.sh     # 立即生效，更新 PATH 环境变量
```

如果你想要安装特定版本的软件，可以使用 `name=ver` 的语法：

```bash
pig ext add -v 17 pgvector=0.7.2 # install pgvector 0.7.2 for PG 17
pig ext add pg16=16.5            # install PostgreSQL 16 with a specific minor version
```

{{< callout type="warning" >}}
请注意，目前只有 PGDG YUM 仓库提供扩展历史版本，PIGSTY 仓库与 PGDG APT 仓库都只提供扩展的 **最新版本**。
{{< /callout >}} 




## 显示扩展

[`pig ext status`](/zh/pig/cmd/ext#ext-status) 命令可以用于显示当前安装的扩展。

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

如果您的当前系统路径中找不到 PostgreSQL（以 `PATH` 中的 `pg_config` 为准），那么请务必通过 `-v|-p` 指定 PG 大版本号或 `pg_config` 路径。



## 扫描扩展

[`pig ext scan`](/zh/pig/cmd/ext#ext-scan) 提供更底层的扩展扫描功能，将扫描指定 PostgreSQL 目录下的共享库，从而发现安装了哪些扩展：

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



## 容器实战

您可以创建一台全新的虚拟机，或者使用下面的 Docker 容器进行功能测试，创建一个 `d13` 目录，创建 `Dockerfile`：

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

pig repo set --region=china    # 添加中国区域的仓库
pig install -y pg18            # 安装 PGDG 18 内核包
pig install -y postgis timescaledb pgvector pg_duckdb
```