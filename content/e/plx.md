---
title: "plx"
linkTitle: "plx"
description: "Transpile multiple procedural dialects to PL/pgSQL"
weight: 3140
categories: ["LANG"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**plx**](https://github.com/commandprompt/plx) : Transpile multiple procedural dialects to PL/pgSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3140** | {{< badge content="plx" link="https://github.com/commandprompt/plx" >}} | {{< ext "plx" >}} | `2.0.1` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "plisql" >}} {{< ext "plpgsql_wrap" >}} {{< ext "orafce" >}} {{< ext "db2fce" >}} |

> [!Note] 2.x NULL interpolation: redeploy function DDL.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `plx` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.1` | {{< bg "18" "plx_18" "green" >}} {{< bg "17" "plx_17" "green" >}} {{< bg "16" "plx_16" "green" >}} {{< bg "15" "plx_15" "green" >}} {{< bg "14" "plx_14" "green" >}} | `plx_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.1` | {{< bg "18" "postgresql-18-plx" "green" >}} {{< bg "17" "postgresql-17-plx" "green" >}} {{< bg "16" "postgresql-16-plx" "green" >}} {{< bg "15" "postgresql-15-plx" "green" >}} {{< bg "14" "postgresql-14-plx" "green" >}} | `postgresql-$v-plx` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "plx_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "plx_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "plx_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "plx_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "plx_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "plx_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "plx_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plx_18` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 316.2 KiB | [plx_18-2.0.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plx_18-2.0.1-1PGSTY.el8.x86_64.rpm) |
| `plx_18` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 311.2 KiB | [plx_18-2.0.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plx_18-2.0.1-1PGSTY.el8.aarch64.rpm) |
| `plx_18` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 333.0 KiB | [plx_18-2.0.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plx_18-2.0.1-1PGSTY.el9.x86_64.rpm) |
| `plx_18` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 329.3 KiB | [plx_18-2.0.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plx_18-2.0.1-1PGSTY.el9.aarch64.rpm) |
| `plx_18` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 338.8 KiB | [plx_18-2.0.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plx_18-2.0.1-1PGSTY.el10.x86_64.rpm) |
| `plx_18` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 331.9 KiB | [plx_18-2.0.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plx_18-2.0.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-plx` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 309.2 KiB | [postgresql-18-plx_2.0.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-18-plx_2.0.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-plx` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 300.6 KiB | [postgresql-18-plx_2.0.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-18-plx_2.0.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-plx` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 310.7 KiB | [postgresql-18-plx_2.0.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-18-plx_2.0.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-plx` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 300.8 KiB | [postgresql-18-plx_2.0.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-18-plx_2.0.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-plx` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 335.5 KiB | [postgresql-18-plx_2.0.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-18-plx_2.0.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-plx` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 330.2 KiB | [postgresql-18-plx_2.0.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-18-plx_2.0.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-plx` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 330.1 KiB | [postgresql-18-plx_2.0.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-18-plx_2.0.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-plx` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 324.0 KiB | [postgresql-18-plx_2.0.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-18-plx_2.0.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-plx` | `2.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 325.3 KiB | [postgresql-18-plx_2.0.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-18-plx_2.0.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-plx` | `2.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 319.6 KiB | [postgresql-18-plx_2.0.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-18-plx_2.0.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plx_17` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 316.0 KiB | [plx_17-2.0.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plx_17-2.0.1-1PGSTY.el8.x86_64.rpm) |
| `plx_17` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 311.1 KiB | [plx_17-2.0.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plx_17-2.0.1-1PGSTY.el8.aarch64.rpm) |
| `plx_17` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 332.8 KiB | [plx_17-2.0.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plx_17-2.0.1-1PGSTY.el9.x86_64.rpm) |
| `plx_17` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 329.1 KiB | [plx_17-2.0.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plx_17-2.0.1-1PGSTY.el9.aarch64.rpm) |
| `plx_17` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 338.6 KiB | [plx_17-2.0.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plx_17-2.0.1-1PGSTY.el10.x86_64.rpm) |
| `plx_17` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 331.6 KiB | [plx_17-2.0.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plx_17-2.0.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-plx` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 309.4 KiB | [postgresql-17-plx_2.0.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-17-plx_2.0.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-plx` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 299.8 KiB | [postgresql-17-plx_2.0.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-17-plx_2.0.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-plx` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 310.9 KiB | [postgresql-17-plx_2.0.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-17-plx_2.0.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-plx` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 300.8 KiB | [postgresql-17-plx_2.0.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-17-plx_2.0.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-plx` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 352.4 KiB | [postgresql-17-plx_2.0.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-17-plx_2.0.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-plx` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 346.6 KiB | [postgresql-17-plx_2.0.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-17-plx_2.0.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-plx` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 329.9 KiB | [postgresql-17-plx_2.0.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-17-plx_2.0.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-plx` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 323.8 KiB | [postgresql-17-plx_2.0.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-17-plx_2.0.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-plx` | `2.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 325.3 KiB | [postgresql-17-plx_2.0.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-17-plx_2.0.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-plx` | `2.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 319.3 KiB | [postgresql-17-plx_2.0.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-17-plx_2.0.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plx_16` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 316.0 KiB | [plx_16-2.0.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plx_16-2.0.1-1PGSTY.el8.x86_64.rpm) |
| `plx_16` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 311.1 KiB | [plx_16-2.0.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plx_16-2.0.1-1PGSTY.el8.aarch64.rpm) |
| `plx_16` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 332.9 KiB | [plx_16-2.0.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plx_16-2.0.1-1PGSTY.el9.x86_64.rpm) |
| `plx_16` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 329.2 KiB | [plx_16-2.0.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plx_16-2.0.1-1PGSTY.el9.aarch64.rpm) |
| `plx_16` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 338.7 KiB | [plx_16-2.0.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plx_16-2.0.1-1PGSTY.el10.x86_64.rpm) |
| `plx_16` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 331.7 KiB | [plx_16-2.0.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plx_16-2.0.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-plx` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 309.0 KiB | [postgresql-16-plx_2.0.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-16-plx_2.0.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-plx` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 300.3 KiB | [postgresql-16-plx_2.0.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-16-plx_2.0.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-plx` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 310.7 KiB | [postgresql-16-plx_2.0.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-16-plx_2.0.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-plx` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 300.8 KiB | [postgresql-16-plx_2.0.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-16-plx_2.0.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-plx` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 350.8 KiB | [postgresql-16-plx_2.0.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-16-plx_2.0.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-plx` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 344.9 KiB | [postgresql-16-plx_2.0.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-16-plx_2.0.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-plx` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 329.9 KiB | [postgresql-16-plx_2.0.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-16-plx_2.0.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-plx` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 322.7 KiB | [postgresql-16-plx_2.0.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-16-plx_2.0.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-plx` | `2.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 325.3 KiB | [postgresql-16-plx_2.0.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-16-plx_2.0.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-plx` | `2.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 319.6 KiB | [postgresql-16-plx_2.0.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-16-plx_2.0.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plx_15` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 315.9 KiB | [plx_15-2.0.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plx_15-2.0.1-1PGSTY.el8.x86_64.rpm) |
| `plx_15` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 311.1 KiB | [plx_15-2.0.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plx_15-2.0.1-1PGSTY.el8.aarch64.rpm) |
| `plx_15` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 326.4 KiB | [plx_15-2.0.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plx_15-2.0.1-1PGSTY.el9.x86_64.rpm) |
| `plx_15` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 324.0 KiB | [plx_15-2.0.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plx_15-2.0.1-1PGSTY.el9.aarch64.rpm) |
| `plx_15` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 330.1 KiB | [plx_15-2.0.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plx_15-2.0.1-1PGSTY.el10.x86_64.rpm) |
| `plx_15` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 326.9 KiB | [plx_15-2.0.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plx_15-2.0.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-plx` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 308.3 KiB | [postgresql-15-plx_2.0.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-15-plx_2.0.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-plx` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 318.3 KiB | [postgresql-15-plx_2.0.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-15-plx_2.0.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-plx` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 309.3 KiB | [postgresql-15-plx_2.0.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-15-plx_2.0.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-plx` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 299.9 KiB | [postgresql-15-plx_2.0.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-15-plx_2.0.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-plx` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 343.5 KiB | [postgresql-15-plx_2.0.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-15-plx_2.0.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-plx` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 340.1 KiB | [postgresql-15-plx_2.0.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-15-plx_2.0.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-plx` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 321.8 KiB | [postgresql-15-plx_2.0.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-15-plx_2.0.1-1PGSTY~noble_amd64.deb) |
| `postgresql-15-plx` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 319.6 KiB | [postgresql-15-plx_2.0.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-15-plx_2.0.1-1PGSTY~noble_arm64.deb) |
| `postgresql-15-plx` | `2.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 318.4 KiB | [postgresql-15-plx_2.0.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-15-plx_2.0.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-plx` | `2.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 315.1 KiB | [postgresql-15-plx_2.0.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-15-plx_2.0.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plx_14` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 315.8 KiB | [plx_14-2.0.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plx_14-2.0.1-1PGSTY.el8.x86_64.rpm) |
| `plx_14` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 311.1 KiB | [plx_14-2.0.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plx_14-2.0.1-1PGSTY.el8.aarch64.rpm) |
| `plx_14` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 326.4 KiB | [plx_14-2.0.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plx_14-2.0.1-1PGSTY.el9.x86_64.rpm) |
| `plx_14` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 323.9 KiB | [plx_14-2.0.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plx_14-2.0.1-1PGSTY.el9.aarch64.rpm) |
| `plx_14` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 330.1 KiB | [plx_14-2.0.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plx_14-2.0.1-1PGSTY.el10.x86_64.rpm) |
| `plx_14` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 326.8 KiB | [plx_14-2.0.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plx_14-2.0.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-plx` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 308.5 KiB | [postgresql-14-plx_2.0.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-14-plx_2.0.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-plx` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 298.5 KiB | [postgresql-14-plx_2.0.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-14-plx_2.0.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-plx` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 309.3 KiB | [postgresql-14-plx_2.0.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-14-plx_2.0.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-plx` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 299.8 KiB | [postgresql-14-plx_2.0.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-14-plx_2.0.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-plx` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 343.3 KiB | [postgresql-14-plx_2.0.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-14-plx_2.0.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-plx` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 340.9 KiB | [postgresql-14-plx_2.0.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-14-plx_2.0.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-plx` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 321.8 KiB | [postgresql-14-plx_2.0.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-14-plx_2.0.1-1PGSTY~noble_amd64.deb) |
| `postgresql-14-plx` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 319.5 KiB | [postgresql-14-plx_2.0.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-14-plx_2.0.1-1PGSTY~noble_arm64.deb) |
| `postgresql-14-plx` | `2.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 318.3 KiB | [postgresql-14-plx_2.0.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-14-plx_2.0.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-plx` | `2.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 315.0 KiB | [postgresql-14-plx_2.0.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-14-plx_2.0.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/commandprompt/plx" title="Repository" icon="github" subtitle="github.com/commandprompt/plx" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plx-2.0.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg plx;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install plx;		# install via package name, for the active PG version

pig install plx -v 18;   # install for PG 18
pig install plx -v 17;   # install for PG 17
pig install plx -v 16;   # install for PG 16
pig install plx -v 15;   # install for PG 15
pig install plx -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION plx;
```

## Usage

Sources:

- [PGXN plx 2.0.1 README](https://pgxn.org/dist/plx/2.0.1/README.html)
- [plx 2.0.1 user guide](https://api.pgxn.org/src/plx/plx-2.0.1/doc/USERGUIDE.md)
- [plx 2.0.1 compatibility notes](https://api.pgxn.org/src/plx/plx-2.0.1/doc/COMPATIBILITY.md)
- [plx 2.0.1 changelog](https://api.pgxn.org/src/plx/plx-2.0.1/CHANGELOG.md)
- [plx control file](https://api.pgxn.org/src/plx/plx-2.0.1/plx.control)

`plx` provides familiar procedural-language dialects that transpile to ordinary PL/pgSQL when `CREATE FUNCTION` runs. PostgreSQL stores and executes the generated PL/pgSQL with its built-in trusted handler; no Ruby, PHP, JavaScript, Python, Go, COBOL, Oracle, or SQL Server runtime is loaded into the backend.

```sql
CREATE EXTENSION plx;
```

### Available Dialects

| Language | Surface syntax |
|---|---|
| `plxruby` | Ruby |
| `plxphp` | PHP |
| `plxjs` | JavaScript |
| `plxts` | TypeScript annotations over the JavaScript dialect |
| `plxpython3` | Python 3 |
| `plxgo` | Go |
| `plxcobol` | ISO COBOL |
| `plxplsql` | Oracle PL/SQL |
| `plxtsql` | Transact-SQL |

All dialects target the same PL/pgSQL statement surface, including assignments, conditionals, loops, query iteration, dynamic SQL, cursors, exceptions, triggers, and set-returning functions.

### Create a Function

Choose a dialect in the `LANGUAGE` clause while keeping the function signature in PostgreSQL types:

```sql
CREATE FUNCTION grade(score integer)
RETURNS text
LANGUAGE plxruby
AS $$
  grade #:: text
  if score >= 90
    grade = "A"
  elsif score >= 80
    grade = "B"
  else
    grade = "F"
  end
  return grade
$$;

SELECT grade(85);
```

Translation happens once, when the function is created. The executable body stored in `pg_proc.prosrc` is regular PL/pgSQL, so it can be dumped, reviewed, and run without a separate interpreter.

### Inspect and Debug Generated Code

```sql
SELECT pg_get_functiondef('grade(integer)'::regprocedure);
SELECT prosrc
FROM pg_proc
WHERE oid = 'grade(integer)'::regprocedure;

SELECT plx_source('grade(integer)'::regprocedure);
```

Runtime error line numbers refer to generated PL/pgSQL. `plx_source()` recovers the original embedded dialect body; use it together with `pg_get_functiondef()` when correlating an error with the source.

### SQL and String Building

Expressions retain PostgreSQL SQL semantics rather than emulating a complete source-language runtime. Use each dialect's query/execute form for SQL and explicit PostgreSQL types for non-literal expressions. The `plx_strbuild` expanded-object helper accelerates repeated string appends on PostgreSQL 18:

```sql
CREATE FUNCTION labels(n integer)
RETURNS text
LANGUAGE plxjs
AS $$
  let out: text = "";
  for (let i = 1; i <= n; i++) {
    out += `item-${i},`;
  }
  return out;
$$;
```

The builder remains correct on PostgreSQL 13-17, but its in-place optimization requires PostgreSQL 18.

### Boundaries and Caveats

- plx implements syntax surfaces, not the source languages' runtimes: there are no gems, Python modules, JavaScript imports, Go goroutines, PHP classes, Oracle packages, or SQL Server transaction commands.
- Functions run in PL/pgSQL's trusted sandbox, with no direct filesystem, network, arbitrary native-code, or transaction-control access.
- Parameters and return types must be PostgreSQL types. Type inference for locals is limited; explicitly declare types for calls and compound expressions.
- SQL uses three-valued logic and PostgreSQL numeric/string semantics. Source-language truthiness and string concatenation with `+` are not reproduced.
- Locals are hoisted into one PL/pgSQL `DECLARE` block, so block-local scope and redeclaration with a different type are unavailable.
- Distribution 2.0.1 installs extension version 2.0.0; 2.0.1 changes PGXN packaging only and has no SQL upgrade of its own. Version 2.0.0 changes interpolation in `plxruby`, `plxphp`, `plxjs`, `plxts`, `plxpython3`, and `plxgo`: a NULL operand now makes the whole interpolated string NULL instead of silently becoming an empty string. Run `ALTER EXTENSION plx UPDATE TO '2.0.0'`, then redeploy affected functions because existing generated PL/pgSQL in `pg_proc.prosrc` is not rewritten by the extension update.
