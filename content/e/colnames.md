---
title: "colnames"
linkTitle: "colnames"
description: "Return column names of a record or composite value"
weight: 4360
categories: ["UTIL"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**colnames**](https://github.com/theory/colnames) : Return column names of a record or composite value


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4360** | {{< badge content="colnames" link="https://github.com/theory/colnames" >}} | {{< ext "colnames" >}} | `1.7.0` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_describe" >}} {{< ext "describe_resultset" >}} {{< ext "get_column" >}} {{< ext "pltoolbox" >}} {{< ext "hstore" >}} |

> [!Note] Distribution 1.7.1 retains extension version 1.7.0; colnames(record) returns name[]. No preload or runtime extension dependencies. Upstream release lists PG8.2-17; PG14-18 verified across all 16 PGSTY target platforms.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `colnames` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.1` | {{< bg "18" "colnames_18" "green" >}} {{< bg "17" "colnames_17" "green" >}} {{< bg "16" "colnames_16" "green" >}} {{< bg "15" "colnames_15" "green" >}} {{< bg "14" "colnames_14" "green" >}} | `colnames_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.1` | {{< bg "18" "postgresql-18-colnames" "green" >}} {{< bg "17" "postgresql-17-colnames" "green" >}} {{< bg "16" "postgresql-16-colnames" "green" >}} {{< bg "15" "postgresql-15-colnames" "green" >}} {{< bg "14" "postgresql-14-colnames" "green" >}} | `postgresql-$v-colnames` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.7.1" "colnames_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.7.1" "colnames_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.7.1" "colnames_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.7.1" "colnames_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.7.1" "colnames_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.7.1" "colnames_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "colnames_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-18-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-17-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-16-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-15-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-14-colnames : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-18-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-17-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-16-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-15-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-14-colnames : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-18-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-17-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-16-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-15-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-14-colnames : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-18-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-17-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-16-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-15-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-14-colnames : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-18-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-17-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-16-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-15-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-14-colnames : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-18-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-17-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-16-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-15-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-14-colnames : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-18-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-17-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-16-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-15-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-14-colnames : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-18-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-17-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-16-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-15-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-14-colnames : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-18-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-17-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-16-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-15-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-14-colnames : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-18-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-17-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-16-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-15-colnames : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.1" "postgresql-14-colnames : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `colnames_18` | `1.7.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.5 KiB | [colnames_18-1.7.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/colnames_18-1.7.1-1PGSTY.el8.x86_64.rpm) |
| `colnames_18` | `1.7.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.7 KiB | [colnames_18-1.7.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/colnames_18-1.7.1-1PGSTY.el8.aarch64.rpm) |
| `colnames_18` | `1.7.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.4 KiB | [colnames_18-1.7.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/colnames_18-1.7.1-1PGSTY.el9.x86_64.rpm) |
| `colnames_18` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.4 KiB | [colnames_18-1.7.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/colnames_18-1.7.1-1PGSTY.el9.aarch64.rpm) |
| `colnames_18` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.4 KiB | [colnames_18-1.7.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/colnames_18-1.7.1-1PGSTY.el10.x86_64.rpm) |
| `colnames_18` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.6 KiB | [colnames_18-1.7.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/colnames_18-1.7.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-colnames` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.4 KiB | [postgresql-18-colnames_1.7.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/colnames/postgresql-18-colnames_1.7.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-colnames` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.4 KiB | [postgresql-18-colnames_1.7.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/colnames/postgresql-18-colnames_1.7.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-colnames` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.3 KiB | [postgresql-18-colnames_1.7.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/colnames/postgresql-18-colnames_1.7.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-colnames` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.4 KiB | [postgresql-18-colnames_1.7.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/colnames/postgresql-18-colnames_1.7.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-colnames` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.6 KiB | [postgresql-18-colnames_1.7.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/colnames/postgresql-18-colnames_1.7.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-colnames` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.7 KiB | [postgresql-18-colnames_1.7.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/colnames/postgresql-18-colnames_1.7.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-colnames` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.7 KiB | [postgresql-18-colnames_1.7.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/colnames/postgresql-18-colnames_1.7.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-colnames` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.6 KiB | [postgresql-18-colnames_1.7.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/colnames/postgresql-18-colnames_1.7.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-colnames` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.8 KiB | [postgresql-18-colnames_1.7.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/colnames/postgresql-18-colnames_1.7.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-colnames` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.0 KiB | [postgresql-18-colnames_1.7.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/colnames/postgresql-18-colnames_1.7.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `colnames_17` | `1.7.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.5 KiB | [colnames_17-1.7.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/colnames_17-1.7.1-1PGSTY.el8.x86_64.rpm) |
| `colnames_17` | `1.7.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.7 KiB | [colnames_17-1.7.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/colnames_17-1.7.1-1PGSTY.el8.aarch64.rpm) |
| `colnames_17` | `1.7.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.3 KiB | [colnames_17-1.7.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/colnames_17-1.7.1-1PGSTY.el9.x86_64.rpm) |
| `colnames_17` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.3 KiB | [colnames_17-1.7.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/colnames_17-1.7.1-1PGSTY.el9.aarch64.rpm) |
| `colnames_17` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.4 KiB | [colnames_17-1.7.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/colnames_17-1.7.1-1PGSTY.el10.x86_64.rpm) |
| `colnames_17` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.6 KiB | [colnames_17-1.7.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/colnames_17-1.7.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-colnames` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.3 KiB | [postgresql-17-colnames_1.7.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/colnames/postgresql-17-colnames_1.7.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-colnames` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.3 KiB | [postgresql-17-colnames_1.7.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/colnames/postgresql-17-colnames_1.7.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-colnames` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.3 KiB | [postgresql-17-colnames_1.7.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/colnames/postgresql-17-colnames_1.7.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-colnames` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.4 KiB | [postgresql-17-colnames_1.7.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/colnames/postgresql-17-colnames_1.7.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-colnames` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 8.0 KiB | [postgresql-17-colnames_1.7.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/colnames/postgresql-17-colnames_1.7.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-colnames` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.9 KiB | [postgresql-17-colnames_1.7.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/colnames/postgresql-17-colnames_1.7.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-colnames` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.7 KiB | [postgresql-17-colnames_1.7.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/colnames/postgresql-17-colnames_1.7.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-colnames` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.6 KiB | [postgresql-17-colnames_1.7.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/colnames/postgresql-17-colnames_1.7.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-colnames` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.8 KiB | [postgresql-17-colnames_1.7.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/colnames/postgresql-17-colnames_1.7.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-colnames` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.0 KiB | [postgresql-17-colnames_1.7.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/colnames/postgresql-17-colnames_1.7.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `colnames_16` | `1.7.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.5 KiB | [colnames_16-1.7.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/colnames_16-1.7.1-1PGSTY.el8.x86_64.rpm) |
| `colnames_16` | `1.7.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.7 KiB | [colnames_16-1.7.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/colnames_16-1.7.1-1PGSTY.el8.aarch64.rpm) |
| `colnames_16` | `1.7.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.3 KiB | [colnames_16-1.7.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/colnames_16-1.7.1-1PGSTY.el9.x86_64.rpm) |
| `colnames_16` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.3 KiB | [colnames_16-1.7.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/colnames_16-1.7.1-1PGSTY.el9.aarch64.rpm) |
| `colnames_16` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.4 KiB | [colnames_16-1.7.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/colnames_16-1.7.1-1PGSTY.el10.x86_64.rpm) |
| `colnames_16` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.6 KiB | [colnames_16-1.7.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/colnames_16-1.7.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-colnames` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.3 KiB | [postgresql-16-colnames_1.7.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/colnames/postgresql-16-colnames_1.7.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-colnames` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.3 KiB | [postgresql-16-colnames_1.7.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/colnames/postgresql-16-colnames_1.7.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-colnames` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.3 KiB | [postgresql-16-colnames_1.7.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/colnames/postgresql-16-colnames_1.7.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-colnames` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.3 KiB | [postgresql-16-colnames_1.7.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/colnames/postgresql-16-colnames_1.7.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-colnames` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.9 KiB | [postgresql-16-colnames_1.7.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/colnames/postgresql-16-colnames_1.7.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-colnames` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.9 KiB | [postgresql-16-colnames_1.7.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/colnames/postgresql-16-colnames_1.7.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-colnames` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.7 KiB | [postgresql-16-colnames_1.7.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/colnames/postgresql-16-colnames_1.7.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-colnames` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.6 KiB | [postgresql-16-colnames_1.7.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/colnames/postgresql-16-colnames_1.7.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-colnames` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.8 KiB | [postgresql-16-colnames_1.7.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/colnames/postgresql-16-colnames_1.7.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-colnames` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.0 KiB | [postgresql-16-colnames_1.7.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/colnames/postgresql-16-colnames_1.7.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `colnames_15` | `1.7.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.5 KiB | [colnames_15-1.7.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/colnames_15-1.7.1-1PGSTY.el8.x86_64.rpm) |
| `colnames_15` | `1.7.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.7 KiB | [colnames_15-1.7.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/colnames_15-1.7.1-1PGSTY.el8.aarch64.rpm) |
| `colnames_15` | `1.7.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.3 KiB | [colnames_15-1.7.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/colnames_15-1.7.1-1PGSTY.el9.x86_64.rpm) |
| `colnames_15` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.3 KiB | [colnames_15-1.7.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/colnames_15-1.7.1-1PGSTY.el9.aarch64.rpm) |
| `colnames_15` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.4 KiB | [colnames_15-1.7.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/colnames_15-1.7.1-1PGSTY.el10.x86_64.rpm) |
| `colnames_15` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.6 KiB | [colnames_15-1.7.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/colnames_15-1.7.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-colnames` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.4 KiB | [postgresql-15-colnames_1.7.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/colnames/postgresql-15-colnames_1.7.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-colnames` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.3 KiB | [postgresql-15-colnames_1.7.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/colnames/postgresql-15-colnames_1.7.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-colnames` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.4 KiB | [postgresql-15-colnames_1.7.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/colnames/postgresql-15-colnames_1.7.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-colnames` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.4 KiB | [postgresql-15-colnames_1.7.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/colnames/postgresql-15-colnames_1.7.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-colnames` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 8.0 KiB | [postgresql-15-colnames_1.7.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/colnames/postgresql-15-colnames_1.7.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-colnames` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 8.0 KiB | [postgresql-15-colnames_1.7.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/colnames/postgresql-15-colnames_1.7.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-colnames` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.7 KiB | [postgresql-15-colnames_1.7.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/colnames/postgresql-15-colnames_1.7.1-1PGSTY~noble_amd64.deb) |
| `postgresql-15-colnames` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.6 KiB | [postgresql-15-colnames_1.7.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/colnames/postgresql-15-colnames_1.7.1-1PGSTY~noble_arm64.deb) |
| `postgresql-15-colnames` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.8 KiB | [postgresql-15-colnames_1.7.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/colnames/postgresql-15-colnames_1.7.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-colnames` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.0 KiB | [postgresql-15-colnames_1.7.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/colnames/postgresql-15-colnames_1.7.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `colnames_14` | `1.7.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.5 KiB | [colnames_14-1.7.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/colnames_14-1.7.1-1PGSTY.el8.x86_64.rpm) |
| `colnames_14` | `1.7.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.6 KiB | [colnames_14-1.7.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/colnames_14-1.7.1-1PGSTY.el8.aarch64.rpm) |
| `colnames_14` | `1.7.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.3 KiB | [colnames_14-1.7.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/colnames_14-1.7.1-1PGSTY.el9.x86_64.rpm) |
| `colnames_14` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.3 KiB | [colnames_14-1.7.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/colnames_14-1.7.1-1PGSTY.el9.aarch64.rpm) |
| `colnames_14` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.4 KiB | [colnames_14-1.7.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/colnames_14-1.7.1-1PGSTY.el10.x86_64.rpm) |
| `colnames_14` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.5 KiB | [colnames_14-1.7.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/colnames_14-1.7.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-colnames` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-14-colnames_1.7.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/colnames/postgresql-14-colnames_1.7.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-colnames` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.4 KiB | [postgresql-14-colnames_1.7.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/colnames/postgresql-14-colnames_1.7.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-colnames` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-14-colnames_1.7.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/colnames/postgresql-14-colnames_1.7.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-colnames` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.4 KiB | [postgresql-14-colnames_1.7.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/colnames/postgresql-14-colnames_1.7.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-colnames` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.0 KiB | [postgresql-14-colnames_1.7.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/colnames/postgresql-14-colnames_1.7.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-colnames` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.0 KiB | [postgresql-14-colnames_1.7.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/colnames/postgresql-14-colnames_1.7.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-colnames` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.7 KiB | [postgresql-14-colnames_1.7.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/colnames/postgresql-14-colnames_1.7.1-1PGSTY~noble_amd64.deb) |
| `postgresql-14-colnames` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.5 KiB | [postgresql-14-colnames_1.7.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/colnames/postgresql-14-colnames_1.7.1-1PGSTY~noble_arm64.deb) |
| `postgresql-14-colnames` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.8 KiB | [postgresql-14-colnames_1.7.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/colnames/postgresql-14-colnames_1.7.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-colnames` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.0 KiB | [postgresql-14-colnames_1.7.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/colnames/postgresql-14-colnames_1.7.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/theory/colnames" title="Repository" icon="github" subtitle="github.com/theory/colnames" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="colnames-1.7.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg colnames;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install colnames;		# install via package name, for the active PG version

pig install colnames -v 18;   # install for PG 18
pig install colnames -v 17;   # install for PG 17
pig install colnames -v 16;   # install for PG 16
pig install colnames -v 15;   # install for PG 15
pig install colnames -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION colnames;
```

## Usage

Sources:

- [Release documentation](https://github.com/theory/colnames/blob/v1.7.1/doc/colnames.md)
- [SQL function](https://github.com/theory/colnames/blob/v1.7.1/sql/colnames.sql)
- [C implementation](https://github.com/theory/colnames/blob/v1.7.1/src/colnames.c)
- [Extension control](https://github.com/theory/colnames/blob/v1.7.1/colnames.control)
- [Release notes](https://github.com/theory/colnames/blob/v1.7.1/Changes)

`colnames` provides `colnames(record)`, returning the input record's field names as a `name[]`. Use it in generic triggers or dynamic SQL helpers that need a composite value's structure without converting its values to JSON. Distribution version `1.7.1` retains extension version `1.7.0`.

### Core Workflow

```sql
CREATE EXTENSION colnames;

SELECT colnames(ROW(1, 'foo', 458.0));
SELECT colnames(t)
FROM (SELECT 1 AS id, 'Ada'::text AS name) AS t;

CREATE TYPE contact AS (id integer, name text);
SELECT colnames(NULL::contact);
```

The three calls return `{f1,f2,f3}`, `{id,name}`, and `{id,name}` respectively. Anonymous rows receive generated names; named rows preserve their attribute names. A typed null works because its composite type supplies the descriptor. An untyped null does not identify a composite type.

### Function Behavior

`colnames(record)` returns `name[]`, is `STABLE`, and is deliberately not `STRICT`. It preserves column order and quoted identifier spelling, omits dropped columns, and returns an empty array for an empty composite type. It reads the row descriptor rather than table data.

### Operational Notes

The extension has no runtime extension dependencies and needs no preload or restart. Its C library loads when the function is called. Creating the extension requires a superuser because the control file does not mark it trusted. It is relocatable and can be installed in a chosen schema or moved with `ALTER EXTENSION colnames SET SCHEMA`.

Release `1.7.1` documents PostgreSQL 8.2–17 compatibility. The PGSTY package targets PostgreSQL 14–18; PostgreSQL 18 acceptance comes from package and regression testing rather than that release's compatibility statement.

PostgreSQL 9.3 and later also offer `row_to_json` with `json_object_keys` as a pure SQL alternative. If discovered names are used in dynamic SQL, quote identifiers and keep values parameterized.
