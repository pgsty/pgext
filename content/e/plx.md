---
title: "plx"
linkTitle: "plx"
description: "Transpile multiple procedural dialects to PL/pgSQL"
weight: 3140
categories: ["LANG"]
width: full
---

[**plx**](https://github.com/commandprompt/plx) : Transpile multiple procedural dialects to PL/pgSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3140** | {{< badge content="plx" link="https://github.com/commandprompt/plx" >}} | {{< ext "plx" >}} | `1.3.1` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "plv8" >}} {{< ext "pllua" >}} {{< ext "pljava" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} {{< ext "plprql" >}} {{< ext "plsh" >}} |

> [!Note] Uses PostgreSQL's built-in PL/pgSQL call handler; no control-file dependency is declared.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `plx` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.1` | {{< bg "18" "plx_18" "green" >}} {{< bg "17" "plx_17" "green" >}} {{< bg "16" "plx_16" "green" >}} {{< bg "15" "plx_15" "green" >}} {{< bg "14" "plx_14" "green" >}} | `plx_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.1` | {{< bg "18" "postgresql-18-plx" "green" >}} {{< bg "17" "postgresql-17-plx" "green" >}} {{< bg "16" "postgresql-16-plx" "green" >}} {{< bg "15" "postgresql-15-plx" "green" >}} {{< bg "14" "postgresql-14-plx" "green" >}} | `postgresql-$v-plx` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "plx_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "plx_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "plx_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "plx_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "plx_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "plx_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "plx_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-plx : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-plx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-plx : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plx_18` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 103.9 KiB | [plx_18-1.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plx_18-1.3.1-1PIGSTY.el8.x86_64.rpm) |
| `plx_18` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 98.5 KiB | [plx_18-1.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plx_18-1.3.1-1PIGSTY.el8.aarch64.rpm) |
| `plx_18` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 110.8 KiB | [plx_18-1.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plx_18-1.3.1-1PIGSTY.el9.x86_64.rpm) |
| `plx_18` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 106.6 KiB | [plx_18-1.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plx_18-1.3.1-1PIGSTY.el9.aarch64.rpm) |
| `plx_18` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 116.2 KiB | [plx_18-1.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plx_18-1.3.1-1PIGSTY.el10.x86_64.rpm) |
| `plx_18` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 109.6 KiB | [plx_18-1.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plx_18-1.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-plx` | `1.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 307.9 KiB | [postgresql-18-plx_1.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-18-plx_1.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-plx` | `1.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 296.2 KiB | [postgresql-18-plx_1.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-18-plx_1.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-plx` | `1.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 309.6 KiB | [postgresql-18-plx_1.3.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-18-plx_1.3.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-plx` | `1.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 296.8 KiB | [postgresql-18-plx_1.3.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-18-plx_1.3.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-plx` | `1.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 324.8 KiB | [postgresql-18-plx_1.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-18-plx_1.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-plx` | `1.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 320.9 KiB | [postgresql-18-plx_1.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-18-plx_1.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-plx` | `1.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 321.2 KiB | [postgresql-18-plx_1.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-18-plx_1.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-plx` | `1.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 315.0 KiB | [postgresql-18-plx_1.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-18-plx_1.3.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-plx` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 318.0 KiB | [postgresql-18-plx_1.3.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-18-plx_1.3.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-plx` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 312.9 KiB | [postgresql-18-plx_1.3.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-18-plx_1.3.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plx_17` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 103.9 KiB | [plx_17-1.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plx_17-1.3.1-1PIGSTY.el8.x86_64.rpm) |
| `plx_17` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 98.4 KiB | [plx_17-1.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plx_17-1.3.1-1PIGSTY.el8.aarch64.rpm) |
| `plx_17` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 110.7 KiB | [plx_17-1.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plx_17-1.3.1-1PIGSTY.el9.x86_64.rpm) |
| `plx_17` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 106.5 KiB | [plx_17-1.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plx_17-1.3.1-1PIGSTY.el9.aarch64.rpm) |
| `plx_17` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 116.1 KiB | [plx_17-1.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plx_17-1.3.1-1PIGSTY.el10.x86_64.rpm) |
| `plx_17` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 109.6 KiB | [plx_17-1.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plx_17-1.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-plx` | `1.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 307.7 KiB | [postgresql-17-plx_1.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-17-plx_1.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-plx` | `1.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 296.2 KiB | [postgresql-17-plx_1.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-17-plx_1.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-plx` | `1.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 309.7 KiB | [postgresql-17-plx_1.3.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-17-plx_1.3.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-plx` | `1.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 296.9 KiB | [postgresql-17-plx_1.3.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-17-plx_1.3.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-plx` | `1.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 340.7 KiB | [postgresql-17-plx_1.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-17-plx_1.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-plx` | `1.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 337.8 KiB | [postgresql-17-plx_1.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-17-plx_1.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-plx` | `1.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 321.1 KiB | [postgresql-17-plx_1.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-17-plx_1.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-plx` | `1.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 315.0 KiB | [postgresql-17-plx_1.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-17-plx_1.3.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-plx` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 317.8 KiB | [postgresql-17-plx_1.3.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-17-plx_1.3.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-plx` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 312.8 KiB | [postgresql-17-plx_1.3.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-17-plx_1.3.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plx_16` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 103.9 KiB | [plx_16-1.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plx_16-1.3.1-1PIGSTY.el8.x86_64.rpm) |
| `plx_16` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 98.4 KiB | [plx_16-1.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plx_16-1.3.1-1PIGSTY.el8.aarch64.rpm) |
| `plx_16` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 110.7 KiB | [plx_16-1.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plx_16-1.3.1-1PIGSTY.el9.x86_64.rpm) |
| `plx_16` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 106.6 KiB | [plx_16-1.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plx_16-1.3.1-1PIGSTY.el9.aarch64.rpm) |
| `plx_16` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 116.1 KiB | [plx_16-1.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plx_16-1.3.1-1PIGSTY.el10.x86_64.rpm) |
| `plx_16` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 109.6 KiB | [plx_16-1.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plx_16-1.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-plx` | `1.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 307.7 KiB | [postgresql-16-plx_1.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-16-plx_1.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-plx` | `1.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 296.6 KiB | [postgresql-16-plx_1.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-16-plx_1.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-plx` | `1.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 309.5 KiB | [postgresql-16-plx_1.3.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-16-plx_1.3.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-plx` | `1.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 296.9 KiB | [postgresql-16-plx_1.3.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-16-plx_1.3.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-plx` | `1.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 339.1 KiB | [postgresql-16-plx_1.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-16-plx_1.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-plx` | `1.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 336.2 KiB | [postgresql-16-plx_1.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-16-plx_1.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-plx` | `1.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 321.1 KiB | [postgresql-16-plx_1.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-16-plx_1.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-plx` | `1.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 315.0 KiB | [postgresql-16-plx_1.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-16-plx_1.3.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-plx` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 317.8 KiB | [postgresql-16-plx_1.3.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-16-plx_1.3.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-plx` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 312.9 KiB | [postgresql-16-plx_1.3.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-16-plx_1.3.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plx_15` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 104.2 KiB | [plx_15-1.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plx_15-1.3.1-1PIGSTY.el8.x86_64.rpm) |
| `plx_15` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 98.6 KiB | [plx_15-1.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plx_15-1.3.1-1PIGSTY.el8.aarch64.rpm) |
| `plx_15` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 110.9 KiB | [plx_15-1.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plx_15-1.3.1-1PIGSTY.el9.x86_64.rpm) |
| `plx_15` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 106.6 KiB | [plx_15-1.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plx_15-1.3.1-1PIGSTY.el9.aarch64.rpm) |
| `plx_15` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 116.6 KiB | [plx_15-1.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plx_15-1.3.1-1PIGSTY.el10.x86_64.rpm) |
| `plx_15` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 109.9 KiB | [plx_15-1.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plx_15-1.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-plx` | `1.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 308.1 KiB | [postgresql-15-plx_1.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-15-plx_1.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-plx` | `1.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 296.1 KiB | [postgresql-15-plx_1.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-15-plx_1.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-plx` | `1.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 309.8 KiB | [postgresql-15-plx_1.3.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-15-plx_1.3.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-plx` | `1.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 297.1 KiB | [postgresql-15-plx_1.3.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-15-plx_1.3.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-plx` | `1.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 339.4 KiB | [postgresql-15-plx_1.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-15-plx_1.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-plx` | `1.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 336.6 KiB | [postgresql-15-plx_1.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-15-plx_1.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-plx` | `1.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 321.5 KiB | [postgresql-15-plx_1.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-15-plx_1.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-plx` | `1.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 315.5 KiB | [postgresql-15-plx_1.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-15-plx_1.3.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-plx` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 318.4 KiB | [postgresql-15-plx_1.3.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-15-plx_1.3.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-plx` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 313.1 KiB | [postgresql-15-plx_1.3.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-15-plx_1.3.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plx_14` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 104.2 KiB | [plx_14-1.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plx_14-1.3.1-1PIGSTY.el8.x86_64.rpm) |
| `plx_14` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 98.6 KiB | [plx_14-1.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plx_14-1.3.1-1PIGSTY.el8.aarch64.rpm) |
| `plx_14` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 110.9 KiB | [plx_14-1.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plx_14-1.3.1-1PIGSTY.el9.x86_64.rpm) |
| `plx_14` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 106.8 KiB | [plx_14-1.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plx_14-1.3.1-1PIGSTY.el9.aarch64.rpm) |
| `plx_14` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 116.5 KiB | [plx_14-1.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plx_14-1.3.1-1PIGSTY.el10.x86_64.rpm) |
| `plx_14` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 109.9 KiB | [plx_14-1.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plx_14-1.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-plx` | `1.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 308.2 KiB | [postgresql-14-plx_1.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-14-plx_1.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-plx` | `1.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 296.3 KiB | [postgresql-14-plx_1.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plx/postgresql-14-plx_1.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-plx` | `1.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 310.0 KiB | [postgresql-14-plx_1.3.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-14-plx_1.3.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-plx` | `1.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 297.6 KiB | [postgresql-14-plx_1.3.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plx/postgresql-14-plx_1.3.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-plx` | `1.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 339.4 KiB | [postgresql-14-plx_1.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-14-plx_1.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-plx` | `1.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 336.5 KiB | [postgresql-14-plx_1.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plx/postgresql-14-plx_1.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-plx` | `1.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 321.4 KiB | [postgresql-14-plx_1.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-14-plx_1.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-plx` | `1.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 315.4 KiB | [postgresql-14-plx_1.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plx/postgresql-14-plx_1.3.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-plx` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 318.4 KiB | [postgresql-14-plx_1.3.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-14-plx_1.3.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-plx` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 313.1 KiB | [postgresql-14-plx_1.3.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plx/postgresql-14-plx_1.3.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/commandprompt/plx" title="Repository" icon="github" subtitle="github.com/commandprompt/plx" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plx-1.3.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg plx;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install plx;		# install via package name, for the active PG version

pig install plx -v 18;   # install for PG 18
pig install plx -v 17;   # install for PG 17
pig install plx -v 16;   # install for PG 16
pig install plx -v 15;   # install for PG 15
pig install plx -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plx;
```

## Usage

Sources:

- [plx 1.3.1 README](https://github.com/commandprompt/plx/blob/v1.3.1/README.md)
- [plx documentation](https://commandprompt.github.io/plx/)
- [plx user guide](https://github.com/commandprompt/plx/blob/v1.3.1/doc/USERGUIDE.md)
- [plx limitations](https://github.com/commandprompt/plx/blob/v1.3.1/doc/LIMITATIONS.md)
- [plx 1.3.1 release](https://github.com/commandprompt/plx/releases/tag/v1.3.1)

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
- Version 1.3.1 is a code-only safety release: it adds lexer/string-builder capacity guards, stack-depth checks, bounded indentation handling, and fixes for raw-string, PHP interpolation, and non-decimal integer literal parsing. After installing the binary, run `ALTER EXTENSION plx UPDATE TO '1.3.1'`.
