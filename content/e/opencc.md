---
title: "opencc"
linkTitle: "opencc"
description: "OpenCC conversion between Simplified Chinese, Traditional Chinese, regional variants and Japanese Shinjitai"
weight: 4480
categories: ["UTIL"]
languages: ["C++"]
licenses: ["Apache-2.0 AND BSD-2-Clause AND MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_opencc**](https://github.com/pgsty/pg_opencc) : OpenCC conversion between Simplified Chinese, Traditional Chinese, regional variants and Japanese Shinjitai


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4480** | {{< badge content="opencc" link="https://github.com/pgsty/pg_opencc" >}} | {{< ext "opencc" "pg_opencc" >}} | `0.1.0` | {{< category "UTIL" >}} | {{< license "Apache-2.0 AND BSD-2-Clause AND MIT" >}} | {{< language "C++" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `opencc` |

> [!Note] UTF8; bundled OpenCC 1.4.2.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_opencc` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_opencc_18" "green" >}} {{< bg "17" "pg_opencc_17" "green" >}} {{< bg "16" "pg_opencc_16" "green" >}} {{< bg "15" "pg_opencc_15" "green" >}} {{< bg "14" "pg_opencc_14" "green" >}} | `pg_opencc_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-opencc" "green" >}} {{< bg "17" "postgresql-17-opencc" "green" >}} {{< bg "16" "postgresql-16-opencc" "green" >}} {{< bg "15" "postgresql-15-opencc" "green" >}} {{< bg "14" "postgresql-14-opencc" "green" >}} | `postgresql-$v-opencc` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_opencc_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-opencc : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-opencc : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-opencc : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-opencc : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-opencc : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-opencc : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-opencc : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-opencc : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-opencc : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-opencc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-opencc : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_opencc_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 890.4 KiB | [pg_opencc_18-0.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_opencc_18-0.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_opencc_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 850.5 KiB | [pg_opencc_18-0.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_opencc_18-0.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_opencc_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 799.1 KiB | [pg_opencc_18-0.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_opencc_18-0.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_opencc_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 782.2 KiB | [pg_opencc_18-0.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_opencc_18-0.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_opencc_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 813.3 KiB | [pg_opencc_18-0.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_opencc_18-0.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_opencc_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 790.4 KiB | [pg_opencc_18-0.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_opencc_18-0.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-opencc` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 784.1 KiB | [postgresql-18-opencc_0.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-opencc/postgresql-18-opencc_0.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-opencc` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 749.7 KiB | [postgresql-18-opencc_0.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-opencc/postgresql-18-opencc_0.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-opencc` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 802.8 KiB | [postgresql-18-opencc_0.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-opencc/postgresql-18-opencc_0.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-opencc` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 764.2 KiB | [postgresql-18-opencc_0.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-opencc/postgresql-18-opencc_0.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-opencc` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 774.9 KiB | [postgresql-18-opencc_0.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-opencc/postgresql-18-opencc_0.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-opencc` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 764.1 KiB | [postgresql-18-opencc_0.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-opencc/postgresql-18-opencc_0.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-opencc` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 764.1 KiB | [postgresql-18-opencc_0.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-opencc/postgresql-18-opencc_0.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-opencc` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 753.5 KiB | [postgresql-18-opencc_0.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-opencc/postgresql-18-opencc_0.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-opencc` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 770.2 KiB | [postgresql-18-opencc_0.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-opencc/postgresql-18-opencc_0.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-opencc` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 761.9 KiB | [postgresql-18-opencc_0.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-opencc/postgresql-18-opencc_0.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_opencc_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 890.3 KiB | [pg_opencc_17-0.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_opencc_17-0.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_opencc_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 850.5 KiB | [pg_opencc_17-0.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_opencc_17-0.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_opencc_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 798.8 KiB | [pg_opencc_17-0.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_opencc_17-0.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_opencc_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 782.1 KiB | [pg_opencc_17-0.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_opencc_17-0.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_opencc_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 813.3 KiB | [pg_opencc_17-0.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_opencc_17-0.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_opencc_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 790.4 KiB | [pg_opencc_17-0.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_opencc_17-0.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-opencc` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 784.0 KiB | [postgresql-17-opencc_0.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-opencc/postgresql-17-opencc_0.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-opencc` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 749.5 KiB | [postgresql-17-opencc_0.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-opencc/postgresql-17-opencc_0.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-opencc` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 802.8 KiB | [postgresql-17-opencc_0.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-opencc/postgresql-17-opencc_0.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-opencc` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 764.1 KiB | [postgresql-17-opencc_0.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-opencc/postgresql-17-opencc_0.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-opencc` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 778.5 KiB | [postgresql-17-opencc_0.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-opencc/postgresql-17-opencc_0.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-opencc` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 767.5 KiB | [postgresql-17-opencc_0.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-opencc/postgresql-17-opencc_0.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-opencc` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 763.2 KiB | [postgresql-17-opencc_0.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-opencc/postgresql-17-opencc_0.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-opencc` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 753.6 KiB | [postgresql-17-opencc_0.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-opencc/postgresql-17-opencc_0.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-opencc` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 769.6 KiB | [postgresql-17-opencc_0.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-opencc/postgresql-17-opencc_0.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-opencc` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 761.8 KiB | [postgresql-17-opencc_0.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-opencc/postgresql-17-opencc_0.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_opencc_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 890.3 KiB | [pg_opencc_16-0.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_opencc_16-0.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_opencc_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 850.5 KiB | [pg_opencc_16-0.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_opencc_16-0.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_opencc_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 798.8 KiB | [pg_opencc_16-0.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_opencc_16-0.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_opencc_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 782.1 KiB | [pg_opencc_16-0.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_opencc_16-0.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_opencc_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 813.4 KiB | [pg_opencc_16-0.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_opencc_16-0.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_opencc_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 790.4 KiB | [pg_opencc_16-0.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_opencc_16-0.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-opencc` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 784.1 KiB | [postgresql-16-opencc_0.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-opencc/postgresql-16-opencc_0.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-opencc` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 749.5 KiB | [postgresql-16-opencc_0.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-opencc/postgresql-16-opencc_0.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-opencc` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 802.9 KiB | [postgresql-16-opencc_0.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-opencc/postgresql-16-opencc_0.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-opencc` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 764.2 KiB | [postgresql-16-opencc_0.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-opencc/postgresql-16-opencc_0.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-opencc` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 778.4 KiB | [postgresql-16-opencc_0.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-opencc/postgresql-16-opencc_0.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-opencc` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 767.5 KiB | [postgresql-16-opencc_0.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-opencc/postgresql-16-opencc_0.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-opencc` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 763.3 KiB | [postgresql-16-opencc_0.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-opencc/postgresql-16-opencc_0.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-opencc` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 753.9 KiB | [postgresql-16-opencc_0.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-opencc/postgresql-16-opencc_0.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-opencc` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 769.6 KiB | [postgresql-16-opencc_0.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-opencc/postgresql-16-opencc_0.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-opencc` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 761.8 KiB | [postgresql-16-opencc_0.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-opencc/postgresql-16-opencc_0.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_opencc_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 890.5 KiB | [pg_opencc_15-0.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_opencc_15-0.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_opencc_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 850.6 KiB | [pg_opencc_15-0.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_opencc_15-0.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_opencc_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 799.0 KiB | [pg_opencc_15-0.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_opencc_15-0.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_opencc_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 782.2 KiB | [pg_opencc_15-0.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_opencc_15-0.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_opencc_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 813.5 KiB | [pg_opencc_15-0.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_opencc_15-0.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_opencc_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 790.4 KiB | [pg_opencc_15-0.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_opencc_15-0.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-opencc` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 784.0 KiB | [postgresql-15-opencc_0.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-opencc/postgresql-15-opencc_0.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-opencc` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 752.3 KiB | [postgresql-15-opencc_0.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-opencc/postgresql-15-opencc_0.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-opencc` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 802.8 KiB | [postgresql-15-opencc_0.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-opencc/postgresql-15-opencc_0.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-opencc` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 764.0 KiB | [postgresql-15-opencc_0.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-opencc/postgresql-15-opencc_0.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-opencc` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 778.4 KiB | [postgresql-15-opencc_0.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-opencc/postgresql-15-opencc_0.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-opencc` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 767.4 KiB | [postgresql-15-opencc_0.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-opencc/postgresql-15-opencc_0.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-opencc` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 763.4 KiB | [postgresql-15-opencc_0.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-opencc/postgresql-15-opencc_0.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-opencc` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 754.1 KiB | [postgresql-15-opencc_0.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-opencc/postgresql-15-opencc_0.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-opencc` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 769.7 KiB | [postgresql-15-opencc_0.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-opencc/postgresql-15-opencc_0.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-opencc` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 761.9 KiB | [postgresql-15-opencc_0.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-opencc/postgresql-15-opencc_0.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_opencc_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 890.5 KiB | [pg_opencc_14-0.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_opencc_14-0.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_opencc_14` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 850.6 KiB | [pg_opencc_14-0.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_opencc_14-0.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_opencc_14` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 798.9 KiB | [pg_opencc_14-0.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_opencc_14-0.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_opencc_14` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 782.2 KiB | [pg_opencc_14-0.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_opencc_14-0.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_opencc_14` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 813.5 KiB | [pg_opencc_14-0.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_opencc_14-0.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_opencc_14` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 790.2 KiB | [pg_opencc_14-0.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_opencc_14-0.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-opencc` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 784.0 KiB | [postgresql-14-opencc_0.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-opencc/postgresql-14-opencc_0.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-opencc` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 749.1 KiB | [postgresql-14-opencc_0.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-opencc/postgresql-14-opencc_0.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-opencc` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 802.7 KiB | [postgresql-14-opencc_0.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-opencc/postgresql-14-opencc_0.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-opencc` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 764.3 KiB | [postgresql-14-opencc_0.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-opencc/postgresql-14-opencc_0.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-opencc` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 778.4 KiB | [postgresql-14-opencc_0.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-opencc/postgresql-14-opencc_0.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-opencc` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 767.2 KiB | [postgresql-14-opencc_0.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-opencc/postgresql-14-opencc_0.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-opencc` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 763.2 KiB | [postgresql-14-opencc_0.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-opencc/postgresql-14-opencc_0.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-opencc` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 753.4 KiB | [postgresql-14-opencc_0.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-opencc/postgresql-14-opencc_0.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-opencc` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 769.7 KiB | [postgresql-14-opencc_0.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-opencc/postgresql-14-opencc_0.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-opencc` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 761.9 KiB | [postgresql-14-opencc_0.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-opencc/postgresql-14-opencc_0.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgsty/pg_opencc" title="Repository" icon="github" subtitle="github.com/pgsty/pg_opencc" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_opencc-0.1.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_opencc;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_opencc;		# install via package name, for the active PG version
pig install opencc;		# install by extension name, for the current active PG version

pig install opencc -v 18;   # install for PG 18
pig install opencc -v 17;   # install for PG 17
pig install opencc -v 16;   # install for PG 16
pig install opencc -v 15;   # install for PG 15
pig install opencc -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION opencc;
```

## Usage

Sources:

- [pg_opencc 0.1.0 README](https://github.com/pgsty/pg_opencc/blob/v0.1.0/README.md)
- [Extension control file](https://github.com/pgsty/pg_opencc/blob/v0.1.0/opencc.control)
- [SQL API](https://github.com/pgsty/pg_opencc/blob/v0.1.0/sql/opencc--0.1.0.sql)
- [Release notes](https://github.com/pgsty/pg_opencc/releases/tag/v0.1.0)

`opencc`, distributed as `pg_opencc`, converts Simplified Chinese, Traditional Chinese, Taiwan and Hong Kong conventions, and Japanese Shinjitai. Version 0.1.0 embeds OpenCC 1.4.2 and its dictionaries in the shared library; no system OpenCC or external dictionary files are required at runtime.

### Enable and convert

Supports PostgreSQL 14–18 and requires a `UTF8` database. This trusted extension can be installed by a user with `CREATE` privilege on the database. Objects use the fixed `opencc` schema; the extension is not relocatable. Preloading and a server restart are not required.

```sql
CREATE EXTENSION opencc;

SELECT opencc.s2twp('鼠标里面的硅二极管坏了，导致光标分辨率降低。');
SELECT opencc.convert('繁體中文', 't2s');
SELECT * FROM opencc.configs;
```

`opencc.convert(input, config)` selects a conversion explicitly. Wrappers such as `opencc.s2t()` and `opencc.t2s()` provide one function per configuration:

| Configurations | Conversion |
| --- | --- |
| `s2t`, `t2s` | Simplified and Traditional Chinese |
| `s2tw`, `tw2s`, `s2hk`, `hk2s` | Taiwan or Hong Kong character conventions |
| `s2twp`, `tw2sp` | Taiwan regional phrases |
| `s2hkp`, `hk2sp` | Hong Kong regional phrases; developing |
| `t2tw`, `tw2t`, `t2hk`, `hk2t` | Traditional Chinese variant conversion |
| `t2jp`, `jp2t` | Japanese old/new character forms; experimental |

### Regional conversion and inspection

`opencc.localize(input, target, source)` accepts `zh-Hans`, `zh-Hant`, `zh-CN`, `zh-TW`, `zh-HK`, and `ja`. Its default `source => 'auto'` detects the script, but cannot distinguish Taiwan from Hong Kong vocabulary. Supply the source region when regional phrases matter.

```sql
SELECT opencc.localize('滑鼠裡面的矽二極體壞了', 'zh-CN', source => 'zh-TW');
SELECT opencc.detect('中文'), opencc.script_stats('中文');
SELECT * FROM opencc.routes;
SELECT * FROM opencc.ambiguities('大战文丑的时候，他的头发很干燥', 's2t');
SELECT opencc.candidates('发', 's2t');
SELECT * FROM opencc.explain('数据库连接池', 's2twp');
```

| Object | Purpose |
| --- | --- |
| `opencc.detect()` | Classify the script, including mixed or neutral text |
| `opencc.script_stats()` | Count Han, simplified, traditional, and kana characters |
| `opencc.configs`, `opencc.routes` | Inspect configurations and regional conversion chains |
| `opencc.ambiguities()`, `opencc.candidates()` | Inspect ambiguous spans and dictionary candidates |
| `opencc.explain()` | Inspect each conversion stage |
| `opencc.version()` | Report extension, engine, and dictionary digest |
| `opencc.dependents()` | Find dependent database objects and suggested rebuild statements |

### Operation and upgrades

Conversion functions are `IMMUTABLE STRICT PARALLEL SAFE` within an extension version. They can be used in expression indexes and stored generated columns. PostgreSQL 18 generated columns must explicitly use `STORED` for user-defined conversion functions.

Dictionary changes in a future extension release can change conversion results. Inspect `opencc.version()` and review the statements returned by `opencc.dependents()` when upgrading. Dependency discovery does not find calls hidden inside other user-defined functions.

```sql
SELECT * FROM opencc.version();
SELECT * FROM opencc.dependents();
```

Each backend initializes converters on demand. Optional startup preloading can share their memory across backends; changing `shared_preload_libraries` requires a restart:

```ini
shared_preload_libraries = 'opencc'
opencc.preload = 's2t,t2s,s2twp,tw2sp'
```

Merge `opencc` into any existing preload list. `opencc.preload` also accepts `all`. Dropping the extension leaves its empty schema behind.
