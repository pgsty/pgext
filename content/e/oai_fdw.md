---
title: "oai_fdw"
linkTitle: "oai_fdw"
description: "Read-only OAI-PMH foreign data wrapper for metadata harvesting over HTTP/HTTPS."
weight: 8770
categories: ["FDW"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**oai_fdw**](https://github.com/jimjonesbr/oai_fdw) : Read-only OAI-PMH foreign data wrapper for metadata harvesting over HTTP/HTTPS.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8770** | {{< badge content="oai_fdw" link="https://github.com/jimjonesbr/oai_fdw" >}} | {{< ext "oai_fdw" >}} | `1.16` | {{< category "FDW" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "rdf_fdw" >}} {{< ext "http" >}} {{< ext "postgres_fdw" >}} |

> [!Note] SQL/control 1.16; package 1.16.0 (tag v1.16). Read-only OAI-PMH 2.0 FDW; libcurl/libxml2; no preload. Response limit has a 1 GB hard ceiling. PG14-18.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.16` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `oai_fdw` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.16.0` | {{< bg "18" "oai_fdw_18" "green" >}} {{< bg "17" "oai_fdw_17" "green" >}} {{< bg "16" "oai_fdw_16" "green" >}} {{< bg "15" "oai_fdw_15" "green" >}} {{< bg "14" "oai_fdw_14" "green" >}} | `oai_fdw_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.16.0` | {{< bg "18" "postgresql-18-oai-fdw" "green" >}} {{< bg "17" "postgresql-17-oai-fdw" "green" >}} {{< bg "16" "postgresql-16-oai-fdw" "green" >}} {{< bg "15" "postgresql-15-oai-fdw" "green" >}} {{< bg "14" "postgresql-14-oai-fdw" "green" >}} | `postgresql-$v-oai-fdw` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "oai_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-18-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-17-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-16-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-15-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-14-oai-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-18-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-17-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-16-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-15-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-14-oai-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-18-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-17-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-16-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-15-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-14-oai-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-18-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-17-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-16-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-15-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-14-oai-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-18-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-17-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-16-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-15-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-14-oai-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-18-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-17-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-16-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-15-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-14-oai-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-18-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-17-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-16-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-15-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-14-oai-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-18-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-17-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-16-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-15-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-14-oai-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-18-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-17-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-16-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-15-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-14-oai-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-18-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-17-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-16-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-15-oai-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.16.0" "postgresql-14-oai-fdw : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oai_fdw_18` | `1.16.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 115.9 KiB | [oai_fdw_18-1.16.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/oai_fdw_18-1.16.0-1PGSTY.el8.x86_64.rpm) |
| `oai_fdw_18` | `1.16.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 113.5 KiB | [oai_fdw_18-1.16.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/oai_fdw_18-1.16.0-1PGSTY.el8.aarch64.rpm) |
| `oai_fdw_18` | `1.16.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 117.5 KiB | [oai_fdw_18-1.16.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/oai_fdw_18-1.16.0-1PGSTY.el9.x86_64.rpm) |
| `oai_fdw_18` | `1.16.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 115.8 KiB | [oai_fdw_18-1.16.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/oai_fdw_18-1.16.0-1PGSTY.el9.aarch64.rpm) |
| `oai_fdw_18` | `1.16.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 117.9 KiB | [oai_fdw_18-1.16.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/oai_fdw_18-1.16.0-1PGSTY.el10.x86_64.rpm) |
| `oai_fdw_18` | `1.16.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 116.1 KiB | [oai_fdw_18-1.16.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/oai_fdw_18-1.16.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-oai-fdw` | `1.16.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 107.5 KiB | [postgresql-18-oai-fdw_1.16.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/oai-fdw/postgresql-18-oai-fdw_1.16.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-oai-fdw` | `1.16.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.1 KiB | [postgresql-18-oai-fdw_1.16.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/oai-fdw/postgresql-18-oai-fdw_1.16.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-oai-fdw` | `1.16.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 107.6 KiB | [postgresql-18-oai-fdw_1.16.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/oai-fdw/postgresql-18-oai-fdw_1.16.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-oai-fdw` | `1.16.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 104.8 KiB | [postgresql-18-oai-fdw_1.16.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/oai-fdw/postgresql-18-oai-fdw_1.16.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-oai-fdw` | `1.16.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 116.1 KiB | [postgresql-18-oai-fdw_1.16.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/oai-fdw/postgresql-18-oai-fdw_1.16.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-oai-fdw` | `1.16.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 114.4 KiB | [postgresql-18-oai-fdw_1.16.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/oai-fdw/postgresql-18-oai-fdw_1.16.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-oai-fdw` | `1.16.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 111.0 KiB | [postgresql-18-oai-fdw_1.16.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/oai-fdw/postgresql-18-oai-fdw_1.16.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-oai-fdw` | `1.16.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 109.4 KiB | [postgresql-18-oai-fdw_1.16.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/oai-fdw/postgresql-18-oai-fdw_1.16.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-oai-fdw` | `1.16.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 109.1 KiB | [postgresql-18-oai-fdw_1.16.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/oai-fdw/postgresql-18-oai-fdw_1.16.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-oai-fdw` | `1.16.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 108.0 KiB | [postgresql-18-oai-fdw_1.16.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/oai-fdw/postgresql-18-oai-fdw_1.16.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oai_fdw_17` | `1.16.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 115.8 KiB | [oai_fdw_17-1.16.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/oai_fdw_17-1.16.0-1PGSTY.el8.x86_64.rpm) |
| `oai_fdw_17` | `1.16.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 113.3 KiB | [oai_fdw_17-1.16.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/oai_fdw_17-1.16.0-1PGSTY.el8.aarch64.rpm) |
| `oai_fdw_17` | `1.16.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 117.5 KiB | [oai_fdw_17-1.16.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/oai_fdw_17-1.16.0-1PGSTY.el9.x86_64.rpm) |
| `oai_fdw_17` | `1.16.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 115.7 KiB | [oai_fdw_17-1.16.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/oai_fdw_17-1.16.0-1PGSTY.el9.aarch64.rpm) |
| `oai_fdw_17` | `1.16.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 117.9 KiB | [oai_fdw_17-1.16.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/oai_fdw_17-1.16.0-1PGSTY.el10.x86_64.rpm) |
| `oai_fdw_17` | `1.16.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 116.1 KiB | [oai_fdw_17-1.16.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/oai_fdw_17-1.16.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-oai-fdw` | `1.16.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 107.3 KiB | [postgresql-17-oai-fdw_1.16.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/oai-fdw/postgresql-17-oai-fdw_1.16.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-oai-fdw` | `1.16.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.0 KiB | [postgresql-17-oai-fdw_1.16.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/oai-fdw/postgresql-17-oai-fdw_1.16.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-oai-fdw` | `1.16.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 107.5 KiB | [postgresql-17-oai-fdw_1.16.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/oai-fdw/postgresql-17-oai-fdw_1.16.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-oai-fdw` | `1.16.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 104.6 KiB | [postgresql-17-oai-fdw_1.16.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/oai-fdw/postgresql-17-oai-fdw_1.16.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-oai-fdw` | `1.16.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 125.6 KiB | [postgresql-17-oai-fdw_1.16.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/oai-fdw/postgresql-17-oai-fdw_1.16.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-oai-fdw` | `1.16.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 123.8 KiB | [postgresql-17-oai-fdw_1.16.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/oai-fdw/postgresql-17-oai-fdw_1.16.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-oai-fdw` | `1.16.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 111.0 KiB | [postgresql-17-oai-fdw_1.16.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/oai-fdw/postgresql-17-oai-fdw_1.16.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-oai-fdw` | `1.16.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 109.4 KiB | [postgresql-17-oai-fdw_1.16.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/oai-fdw/postgresql-17-oai-fdw_1.16.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-oai-fdw` | `1.16.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 109.1 KiB | [postgresql-17-oai-fdw_1.16.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/oai-fdw/postgresql-17-oai-fdw_1.16.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-oai-fdw` | `1.16.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 108.0 KiB | [postgresql-17-oai-fdw_1.16.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/oai-fdw/postgresql-17-oai-fdw_1.16.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oai_fdw_16` | `1.16.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 115.7 KiB | [oai_fdw_16-1.16.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/oai_fdw_16-1.16.0-1PGSTY.el8.x86_64.rpm) |
| `oai_fdw_16` | `1.16.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 113.3 KiB | [oai_fdw_16-1.16.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/oai_fdw_16-1.16.0-1PGSTY.el8.aarch64.rpm) |
| `oai_fdw_16` | `1.16.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 117.4 KiB | [oai_fdw_16-1.16.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/oai_fdw_16-1.16.0-1PGSTY.el9.x86_64.rpm) |
| `oai_fdw_16` | `1.16.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 115.8 KiB | [oai_fdw_16-1.16.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/oai_fdw_16-1.16.0-1PGSTY.el9.aarch64.rpm) |
| `oai_fdw_16` | `1.16.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 117.9 KiB | [oai_fdw_16-1.16.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/oai_fdw_16-1.16.0-1PGSTY.el10.x86_64.rpm) |
| `oai_fdw_16` | `1.16.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 116.0 KiB | [oai_fdw_16-1.16.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/oai_fdw_16-1.16.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-oai-fdw` | `1.16.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 107.5 KiB | [postgresql-16-oai-fdw_1.16.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/oai-fdw/postgresql-16-oai-fdw_1.16.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-oai-fdw` | `1.16.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.0 KiB | [postgresql-16-oai-fdw_1.16.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/oai-fdw/postgresql-16-oai-fdw_1.16.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-oai-fdw` | `1.16.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 107.7 KiB | [postgresql-16-oai-fdw_1.16.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/oai-fdw/postgresql-16-oai-fdw_1.16.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-oai-fdw` | `1.16.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 104.7 KiB | [postgresql-16-oai-fdw_1.16.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/oai-fdw/postgresql-16-oai-fdw_1.16.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-oai-fdw` | `1.16.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 125.4 KiB | [postgresql-16-oai-fdw_1.16.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/oai-fdw/postgresql-16-oai-fdw_1.16.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-oai-fdw` | `1.16.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 123.5 KiB | [postgresql-16-oai-fdw_1.16.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/oai-fdw/postgresql-16-oai-fdw_1.16.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-oai-fdw` | `1.16.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 111.0 KiB | [postgresql-16-oai-fdw_1.16.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/oai-fdw/postgresql-16-oai-fdw_1.16.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-oai-fdw` | `1.16.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 111.9 KiB | [postgresql-16-oai-fdw_1.16.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/oai-fdw/postgresql-16-oai-fdw_1.16.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-oai-fdw` | `1.16.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 109.1 KiB | [postgresql-16-oai-fdw_1.16.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/oai-fdw/postgresql-16-oai-fdw_1.16.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-oai-fdw` | `1.16.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 108.0 KiB | [postgresql-16-oai-fdw_1.16.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/oai-fdw/postgresql-16-oai-fdw_1.16.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oai_fdw_15` | `1.16.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 115.7 KiB | [oai_fdw_15-1.16.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/oai_fdw_15-1.16.0-1PGSTY.el8.x86_64.rpm) |
| `oai_fdw_15` | `1.16.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 113.3 KiB | [oai_fdw_15-1.16.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/oai_fdw_15-1.16.0-1PGSTY.el8.aarch64.rpm) |
| `oai_fdw_15` | `1.16.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 117.4 KiB | [oai_fdw_15-1.16.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/oai_fdw_15-1.16.0-1PGSTY.el9.x86_64.rpm) |
| `oai_fdw_15` | `1.16.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 115.8 KiB | [oai_fdw_15-1.16.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/oai_fdw_15-1.16.0-1PGSTY.el9.aarch64.rpm) |
| `oai_fdw_15` | `1.16.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 117.8 KiB | [oai_fdw_15-1.16.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/oai_fdw_15-1.16.0-1PGSTY.el10.x86_64.rpm) |
| `oai_fdw_15` | `1.16.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 116.1 KiB | [oai_fdw_15-1.16.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/oai_fdw_15-1.16.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-oai-fdw` | `1.16.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 107.5 KiB | [postgresql-15-oai-fdw_1.16.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/oai-fdw/postgresql-15-oai-fdw_1.16.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-oai-fdw` | `1.16.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 116.0 KiB | [postgresql-15-oai-fdw_1.16.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/oai-fdw/postgresql-15-oai-fdw_1.16.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-oai-fdw` | `1.16.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 107.6 KiB | [postgresql-15-oai-fdw_1.16.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/oai-fdw/postgresql-15-oai-fdw_1.16.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-oai-fdw` | `1.16.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 104.8 KiB | [postgresql-15-oai-fdw_1.16.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/oai-fdw/postgresql-15-oai-fdw_1.16.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-oai-fdw` | `1.16.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 125.3 KiB | [postgresql-15-oai-fdw_1.16.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/oai-fdw/postgresql-15-oai-fdw_1.16.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-oai-fdw` | `1.16.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 123.6 KiB | [postgresql-15-oai-fdw_1.16.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/oai-fdw/postgresql-15-oai-fdw_1.16.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-oai-fdw` | `1.16.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 111.1 KiB | [postgresql-15-oai-fdw_1.16.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/oai-fdw/postgresql-15-oai-fdw_1.16.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-oai-fdw` | `1.16.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 109.4 KiB | [postgresql-15-oai-fdw_1.16.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/oai-fdw/postgresql-15-oai-fdw_1.16.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-oai-fdw` | `1.16.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 109.0 KiB | [postgresql-15-oai-fdw_1.16.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/oai-fdw/postgresql-15-oai-fdw_1.16.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-oai-fdw` | `1.16.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 108.0 KiB | [postgresql-15-oai-fdw_1.16.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/oai-fdw/postgresql-15-oai-fdw_1.16.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oai_fdw_14` | `1.16.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 115.7 KiB | [oai_fdw_14-1.16.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/oai_fdw_14-1.16.0-1PGSTY.el8.x86_64.rpm) |
| `oai_fdw_14` | `1.16.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 113.3 KiB | [oai_fdw_14-1.16.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/oai_fdw_14-1.16.0-1PGSTY.el8.aarch64.rpm) |
| `oai_fdw_14` | `1.16.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 117.4 KiB | [oai_fdw_14-1.16.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/oai_fdw_14-1.16.0-1PGSTY.el9.x86_64.rpm) |
| `oai_fdw_14` | `1.16.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 115.8 KiB | [oai_fdw_14-1.16.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/oai_fdw_14-1.16.0-1PGSTY.el9.aarch64.rpm) |
| `oai_fdw_14` | `1.16.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 117.8 KiB | [oai_fdw_14-1.16.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/oai_fdw_14-1.16.0-1PGSTY.el10.x86_64.rpm) |
| `oai_fdw_14` | `1.16.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 116.3 KiB | [oai_fdw_14-1.16.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/oai_fdw_14-1.16.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-oai-fdw` | `1.16.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 107.6 KiB | [postgresql-14-oai-fdw_1.16.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/oai-fdw/postgresql-14-oai-fdw_1.16.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-oai-fdw` | `1.16.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.0 KiB | [postgresql-14-oai-fdw_1.16.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/oai-fdw/postgresql-14-oai-fdw_1.16.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-oai-fdw` | `1.16.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 107.6 KiB | [postgresql-14-oai-fdw_1.16.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/oai-fdw/postgresql-14-oai-fdw_1.16.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-oai-fdw` | `1.16.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 104.9 KiB | [postgresql-14-oai-fdw_1.16.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/oai-fdw/postgresql-14-oai-fdw_1.16.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-oai-fdw` | `1.16.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 125.3 KiB | [postgresql-14-oai-fdw_1.16.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/oai-fdw/postgresql-14-oai-fdw_1.16.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-oai-fdw` | `1.16.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 123.5 KiB | [postgresql-14-oai-fdw_1.16.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/oai-fdw/postgresql-14-oai-fdw_1.16.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-oai-fdw` | `1.16.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 111.1 KiB | [postgresql-14-oai-fdw_1.16.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/oai-fdw/postgresql-14-oai-fdw_1.16.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-oai-fdw` | `1.16.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 109.4 KiB | [postgresql-14-oai-fdw_1.16.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/oai-fdw/postgresql-14-oai-fdw_1.16.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-oai-fdw` | `1.16.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 109.0 KiB | [postgresql-14-oai-fdw_1.16.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/oai-fdw/postgresql-14-oai-fdw_1.16.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-oai-fdw` | `1.16.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 107.9 KiB | [postgresql-14-oai-fdw_1.16.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/oai-fdw/postgresql-14-oai-fdw_1.16.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jimjonesbr/oai_fdw" title="Repository" icon="github" subtitle="github.com/jimjonesbr/oai_fdw" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="oai_fdw-1.16.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg oai_fdw;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install oai_fdw;		# install via package name, for the active PG version

pig install oai_fdw -v 18;   # install for PG 18
pig install oai_fdw -v 17;   # install for PG 17
pig install oai_fdw -v 16;   # install for PG 16
pig install oai_fdw -v 15;   # install for PG 15
pig install oai_fdw -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION oai_fdw;
```

## Usage

Sources:

- [CHANGELOG.md](https://github.com/jimjonesbr/oai_fdw/blob/206169ac1cc2ea9a35d92e9a3ade9d82ec01d18c/CHANGELOG.md)
- [Official v1.16 README](https://github.com/jimjonesbr/oai_fdw/blob/206169ac1cc2ea9a35d92e9a3ade9d82ec01d18c/README.md)
- [Official v1.16 control file](https://github.com/jimjonesbr/oai_fdw/blob/206169ac1cc2ea9a35d92e9a3ade9d82ec01d18c/oai_fdw.control)
- [Official v1.16 changelog](https://github.com/jimjonesbr/oai_fdw/blob/206169ac1cc2ea9a35d92e9a3ade9d82ec01d18c/CHANGELOG.md)
- [Official OAI-PMH specification](https://www.openarchives.org/OAI/openarchivesprotocol.html)

`oai_fdw` 1.16 is a foreign data wrapper for harvesting metadata from OAI-PMH 2.0 repositories over HTTP or HTTPS. It maps OAI record headers and XML content to foreign-table columns, supports automatic table generation for repositories or sets, and can copy harvested records into local tables.

### Core Workflow

Create a server for the repository, then either define a foreign table explicitly or use `IMPORT FOREIGN SCHEMA`. Every table needs a metadata format such as `oai_dc`.

```sql
CREATE EXTENSION oai_fdw;
CREATE SCHEMA oai_data;

CREATE SERVER museum_oai
FOREIGN DATA WRAPPER oai_fdw
OPTIONS (url 'https://example.org/oai');

CREATE FOREIGN TABLE oai_data.records (
    id text OPTIONS (oai_node 'identifier'),
    content xml OPTIONS (oai_node 'content'),
    sets text[] OPTIONS (oai_node 'setspec'),
    datestamp timestamp OPTIONS (oai_node 'datestamp'),
    format text OPTIONS (oai_node 'metadataprefix'),
    status boolean OPTIONS (oai_node 'status')
)
SERVER museum_oai
OPTIONS (metadataprefix 'oai_dc');

SELECT id, datestamp
FROM oai_data.records
LIMIT 100;
```

Without `setspec`, a scan can harvest every set exposed by the repository. Use `from`, `until`, or `setspec` table options and an outer `LIMIT` where appropriate, but verify request behavior with the remote service.

### Import and Support Interfaces

- `IMPORT FOREIGN SCHEMA oai_repository` creates one table spanning the repository; `oai_sets` creates a table per OAI set. Both require the `metadataprefix` import option and support `LIMIT TO` or `EXCEPT`.
- Column option `oai_node` maps `identifier`, `setspec`, `datestamp`, `content`, `metadataprefix`, or `status` to compatible PostgreSQL types.
- Server options include `url`, `http_proxy`, `connect_timeout`, `connect_retry`, `request_redirect`, and `request_max_redirect`.
- `OAI_HarvestTable` copies a foreign-table harvest into a local table; OAI support functions expose repository metadata such as sets and formats.

Version 1.13 moves `proxy_user` and `proxy_password` to the user mapping. Repository `user` and `password` also belong there and use HTTP Basic Authentication.

### Network, Credentials, and Consistency

Queries perform synchronous remote requests and parse untrusted XML. Latency, pagination, server throttling, redirects, malformed responses, and repository changes affect results. Store credentials in narrowly scoped user mappings, restrict who may create or alter servers, and treat configurable URLs/proxies as outbound-network access with SSRF implications.

OAI-PMH is a harvesting protocol, not a transactional database API. Repeated scans can observe changing pages and deletions; local copies become stale. Record harvest windows and identifiers, make refreshes restartable, and reconcile deleted/status records explicitly. Version 1.13 requires matching libcurl/libxml2 build dependencies and should be tested against each target repository’s metadata format.

### Version 1.15

HTTP 429 and 503 retries respect backoff/Retry-After; from/until validation and timestamp granularity are stricter. Server USAGE and view-owner user mappings are checked. Harvest windows and cancellation handling are corrected, and stalled transfers abort after the documented five-minute low-speed threshold. Install matching files and apply `ALTER EXTENSION oai_fdw UPDATE`. PGXN calls this distribution 1.15.0, while the SQL/control version is 1.15.

### 1.16 Response Limits

Use the server option `max_response_size` to bound response bytes; a hard 1 GB ceiling applies even without a smaller configured limit. Exceeding the cap aborts the request. The FDW is read-only: INSERT, UPDATE and DELETE now fail immediately. GetRecord row estimates are capped at one, and EXPLAIN includes the server identity. Distribution 1.16.0 uses SQL version 1.16; install matching files and run `ALTER EXTENSION oai_fdw UPDATE TO '1.16'`.
