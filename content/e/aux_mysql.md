---
title: "aux_mysql"
linkTitle: "aux_mysql"
description: "MySQL Supplementary Extension"
weight: 9420
categories: ["SIM"]
languages: ["C"]
licenses: ["GPL-3.0-or-later"]
repos: ["PIGSTY"]
page_width: full
---

[**openhalo**](https://github.com/pgsty/openhalo) : MySQL Supplementary Extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9420** | {{< badge content="aux_mysql" link="https://github.com/pgsty/openhalo" >}} | {{< ext "aux_mysql" "openhalo" >}} | `2.0` | {{< category "SIM" >}} | {{< license "GPL-3.0-or-later" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `mysql` |
|   **See Also**    | {{< ext "mysql_fdw" >}} {{< ext "db_migrator" >}} {{< ext "postgres_fdw" >}} {{< ext "orafce" >}} {{< ext "db2fce" >}} {{< ext "tds_fdw" >}} {{< ext "db2_fdw" >}} |

> [!Note] Requires the OpenHalo kernel. OpenHalo 19 bundles aux_mysql 2.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0` | {{< bg "19" "" "green" >}} {{< bg "18" "" "red" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} | `openhalo` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `19.0.0` | {{< bg "19" "openhalodb-19" "green" >}} {{< bg "18" "openhalodb-19" "red" >}} {{< bg "17" "openhalodb-19" "red" >}} {{< bg "16" "openhalodb-19" "red" >}} {{< bg "15" "openhalodb-19" "red" >}} | `openhalodb-19` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `19.0.0` | {{< bg "19" "openhalodb-19" "green" >}} {{< bg "18" "openhalodb-19" "red" >}} {{< bg "17" "openhalodb-19" "red" >}} {{< bg "16" "openhalodb-19" "red" >}} {{< bg "15" "openhalodb-19" "red" >}} | `openhalodb-19` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 19.0.0" "openhalodb-19 : AVAIL 1" "green" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 19.0.0" "openhalodb-19 : AVAIL 1" "green" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 19.0.0" "openhalodb-19 : AVAIL 1" "green" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 19.0.0" "openhalodb-19 : AVAIL 1" "green" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 19.0.0" "openhalodb-19 : AVAIL 1" "green" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 19.0.0" "openhalodb-19 : AVAIL 1" "green" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 19.0.0" "openhalodb-19 : AVAIL 1" "green" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 19.0.0" "openhalodb-19 : AVAIL 1" "green" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 19.0.0" "openhalodb-19 : AVAIL 1" "green" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 19.0.0" "openhalodb-19 : AVAIL 1" "green" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 19.0.0" "openhalodb-19 : AVAIL 1" "green" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 19.0.0" "openhalodb-19 : AVAIL 1" "green" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 19.0.0" "openhalodb-19 : AVAIL 1" "green" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 19.0.0" "openhalodb-19 : AVAIL 1" "green" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} | {{< bg "N/A" "openhalodb-19 : N/A 0" "gray" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `openhalodb-19` | `19.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.2 MiB | [openhalodb-19-19.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/openhalodb-19-19.0.0-1PGSTY.el9.x86_64.rpm) |
| `openhalodb-19` | `19.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.8 MiB | [openhalodb-19-19.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/openhalodb-19-19.0.0-1PGSTY.el9.aarch64.rpm) |
| `openhalodb-19` | `19.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.3 MiB | [openhalodb-19-19.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/openhalodb-19-19.0.0-1PGSTY.el10.x86_64.rpm) |
| `openhalodb-19` | `19.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.9 MiB | [openhalodb-19-19.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/openhalodb-19-19.0.0-1PGSTY.el10.aarch64.rpm) |
| `openhalodb-19` | `19.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.0 MiB | [openhalodb-19_19.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/openhalodb-19/openhalodb-19_19.0.0-1PGSTY~bookworm_amd64.deb) |
| `openhalodb-19` | `19.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.3 MiB | [openhalodb-19_19.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/openhalodb-19/openhalodb-19_19.0.0-1PGSTY~bookworm_arm64.deb) |
| `openhalodb-19` | `19.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.9 MiB | [openhalodb-19_19.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/openhalodb-19/openhalodb-19_19.0.0-1PGSTY~trixie_amd64.deb) |
| `openhalodb-19` | `19.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.3 MiB | [openhalodb-19_19.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/openhalodb-19/openhalodb-19_19.0.0-1PGSTY~trixie_arm64.deb) |
| `openhalodb-19` | `19.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.0 MiB | [openhalodb-19_19.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/openhalodb-19/openhalodb-19_19.0.0-1PGSTY~jammy_amd64.deb) |
| `openhalodb-19` | `19.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.7 MiB | [openhalodb-19_19.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/openhalodb-19/openhalodb-19_19.0.0-1PGSTY~jammy_arm64.deb) |
| `openhalodb-19` | `19.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.2 MiB | [openhalodb-19_19.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/openhalodb-19/openhalodb-19_19.0.0-1PGSTY~noble_amd64.deb) |
| `openhalodb-19` | `19.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.9 MiB | [openhalodb-19_19.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/openhalodb-19/openhalodb-19_19.0.0-1PGSTY~noble_arm64.deb) |
| `openhalodb-19` | `19.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.7 MiB | [openhalodb-19_19.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/openhalodb-19/openhalodb-19_19.0.0-1PGSTY~resolute_amd64.deb) |
| `openhalodb-19` | `19.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 23.3 MiB | [openhalodb-19_19.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/openhalodb-19/openhalodb-19_19.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgsty/openhalo" title="Repository" icon="github" subtitle="github.com/pgsty/openhalo" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="openhalodb-19.0.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg openhalo;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install openhalo;		# install by package name
pig install aux_mysql;		# install by extension name

pig install aux_mysql -v 19;   # install for PG 19

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'aux_mysql';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION aux_mysql;
```

## Usage

Sources:

- [19.0.0 manual](https://github.com/pgsty/openhalo/blob/v19.0.0/README.md)
- [2.0 control file](https://github.com/pgsty/openhalo/blob/v19.0.0/contrib/aux_mysql/aux_mysql.control)

`aux_mysql` 2.0 ships with OpenHalo 19.0.0 and provides MySQL wire protocol, syntax, and functions. Run `pig install openhalo-19` to install the `openhalodb-19` kernel package. This command selects PG19 even when another PostgreSQL major is active.

### Enable and connect

New OpenHalo 19 clusters load `aux_mysql` through `shared_preload_libraries` by default. The extension creates ICU collations, so use a database with an ICU-compatible encoding such as `UTF8`. As an administrator, connect over PostgreSQL to the database named by `mysql.database` (default `postgres`) and create the extension.

```sql
CREATE EXTENSION aux_mysql;
SELECT extversion FROM pg_extension WHERE extname = 'aux_mysql';
```

On OpenHalo 19 the query returns `2.0`. With `mysql.listener_on` enabled, MySQL clients can connect to `mysql.port` (default 3306). MySQL databases map to schemas in the selected PostgreSQL database. The extension installs casts in that database. Use a separate database for workloads that require stock PostgreSQL casts.
