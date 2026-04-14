---
title: "pglinter"
linkTitle: "pglinter"
description: "PostgreSQL Linting and Analysis Extension"
weight: 5090
categories: ["ADMIN"]
width: full
---

[**pglinter**](https://github.com/pmpetit/pglinter) : PostgreSQL Linting and Analysis Extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5090** | {{< badge content="pglinter" link="https://github.com/pmpetit/pglinter" >}} | {{< ext "pglinter" >}} | `1.1.2` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "amcheck" >}} {{< ext "supautils" >}} |

> [!Note] manually upgraded PGRX from 0.16.1 to 0.17.0 by Vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pglinter` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.2` | {{< bg "18" "pglinter_18" "green" >}} {{< bg "17" "pglinter_17" "green" >}} {{< bg "16" "pglinter_16" "green" >}} {{< bg "15" "pglinter_15" "green" >}} {{< bg "14" "pglinter_14" "green" >}} | `pglinter_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.2` | {{< bg "18" "postgresql-18-pglinter" "green" >}} {{< bg "17" "postgresql-17-pglinter" "green" >}} {{< bg "16" "postgresql-16-pglinter" "green" >}} {{< bg "15" "postgresql-15-pglinter" "green" >}} {{< bg "14" "postgresql-14-pglinter" "green" >}} | `postgresql-$v-pglinter` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pglinter_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-pglinter : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-pglinter : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-pglinter : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-pglinter : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-pglinter : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-pglinter : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-pglinter : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-pglinter : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglinter_18` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 587.9 KiB | [pglinter_18-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_18-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_18` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 449.1 KiB | [pglinter_18-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_18-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_18` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 601.7 KiB | [pglinter_18-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_18-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_18` | `1.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 477.2 KiB | [pglinter_18-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_18-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_18` | `1.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 601.5 KiB | [pglinter_18-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_18-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_18` | `1.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 476.6 KiB | [pglinter_18-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_18-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pglinter` | `1.1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 496.2 KiB | [postgresql-18-pglinter_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-18-pglinter_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pglinter` | `1.1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 364.7 KiB | [postgresql-18-pglinter_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-18-pglinter_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pglinter` | `1.1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 496.2 KiB | [postgresql-18-pglinter_1.1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-18-pglinter_1.1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pglinter` | `1.1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 364.5 KiB | [postgresql-18-pglinter_1.1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-18-pglinter_1.1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pglinter` | `1.1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 549.8 KiB | [postgresql-18-pglinter_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-18-pglinter_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pglinter` | `1.1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 424.9 KiB | [postgresql-18-pglinter_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-18-pglinter_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pglinter` | `1.1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 545.0 KiB | [postgresql-18-pglinter_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-18-pglinter_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pglinter` | `1.1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 418.6 KiB | [postgresql-18-pglinter_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-18-pglinter_1.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglinter_17` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 587.7 KiB | [pglinter_17-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_17-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_17` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 447.7 KiB | [pglinter_17-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_17-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_17` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 601.9 KiB | [pglinter_17-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_17-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_17` | `1.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 475.5 KiB | [pglinter_17-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_17-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_17` | `1.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 602.0 KiB | [pglinter_17-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_17-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_17` | `1.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 475.6 KiB | [pglinter_17-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_17-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pglinter` | `1.1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 496.3 KiB | [postgresql-17-pglinter_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-17-pglinter_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pglinter` | `1.1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 363.6 KiB | [postgresql-17-pglinter_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-17-pglinter_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pglinter` | `1.1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 495.7 KiB | [postgresql-17-pglinter_1.1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-17-pglinter_1.1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pglinter` | `1.1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 363.7 KiB | [postgresql-17-pglinter_1.1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-17-pglinter_1.1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pglinter` | `1.1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 550.0 KiB | [postgresql-17-pglinter_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-17-pglinter_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pglinter` | `1.1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 423.3 KiB | [postgresql-17-pglinter_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-17-pglinter_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pglinter` | `1.1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 545.1 KiB | [postgresql-17-pglinter_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-17-pglinter_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pglinter` | `1.1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 417.7 KiB | [postgresql-17-pglinter_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-17-pglinter_1.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglinter_16` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 587.6 KiB | [pglinter_16-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_16-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_16` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 447.5 KiB | [pglinter_16-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_16-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_16` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 601.3 KiB | [pglinter_16-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_16-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_16` | `1.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 475.2 KiB | [pglinter_16-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_16-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_16` | `1.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 601.5 KiB | [pglinter_16-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_16-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_16` | `1.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 475.5 KiB | [pglinter_16-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_16-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pglinter` | `1.1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 495.6 KiB | [postgresql-16-pglinter_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-16-pglinter_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pglinter` | `1.1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 363.7 KiB | [postgresql-16-pglinter_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-16-pglinter_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pglinter` | `1.1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 495.4 KiB | [postgresql-16-pglinter_1.1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-16-pglinter_1.1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pglinter` | `1.1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 363.6 KiB | [postgresql-16-pglinter_1.1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-16-pglinter_1.1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pglinter` | `1.1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 549.8 KiB | [postgresql-16-pglinter_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-16-pglinter_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pglinter` | `1.1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 423.3 KiB | [postgresql-16-pglinter_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-16-pglinter_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pglinter` | `1.1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 545.0 KiB | [postgresql-16-pglinter_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-16-pglinter_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pglinter` | `1.1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 417.6 KiB | [postgresql-16-pglinter_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-16-pglinter_1.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglinter_15` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 587.4 KiB | [pglinter_15-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_15-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_15` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 449.1 KiB | [pglinter_15-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_15-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_15` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 601.3 KiB | [pglinter_15-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_15-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_15` | `1.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 477.1 KiB | [pglinter_15-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_15-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_15` | `1.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 601.5 KiB | [pglinter_15-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_15-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_15` | `1.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 476.5 KiB | [pglinter_15-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_15-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pglinter` | `1.1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 495.9 KiB | [postgresql-15-pglinter_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-15-pglinter_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pglinter` | `1.1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 364.7 KiB | [postgresql-15-pglinter_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-15-pglinter_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pglinter` | `1.1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 495.6 KiB | [postgresql-15-pglinter_1.1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-15-pglinter_1.1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pglinter` | `1.1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 365.0 KiB | [postgresql-15-pglinter_1.1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-15-pglinter_1.1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pglinter` | `1.1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 549.7 KiB | [postgresql-15-pglinter_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-15-pglinter_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pglinter` | `1.1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 424.4 KiB | [postgresql-15-pglinter_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-15-pglinter_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pglinter` | `1.1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 544.8 KiB | [postgresql-15-pglinter_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-15-pglinter_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pglinter` | `1.1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 418.2 KiB | [postgresql-15-pglinter_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-15-pglinter_1.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglinter_14` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 589.3 KiB | [pglinter_14-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_14-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_14` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 447.8 KiB | [pglinter_14-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_14-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_14` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 603.2 KiB | [pglinter_14-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_14-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_14` | `1.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 475.6 KiB | [pglinter_14-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_14-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_14` | `1.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 603.3 KiB | [pglinter_14-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_14-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_14` | `1.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 475.3 KiB | [pglinter_14-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_14-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pglinter` | `1.1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 497.8 KiB | [postgresql-14-pglinter_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-14-pglinter_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pglinter` | `1.1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 363.8 KiB | [postgresql-14-pglinter_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-14-pglinter_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pglinter` | `1.1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 497.5 KiB | [postgresql-14-pglinter_1.1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-14-pglinter_1.1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pglinter` | `1.1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 364.5 KiB | [postgresql-14-pglinter_1.1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-14-pglinter_1.1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pglinter` | `1.1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 550.9 KiB | [postgresql-14-pglinter_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-14-pglinter_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pglinter` | `1.1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 423.7 KiB | [postgresql-14-pglinter_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-14-pglinter_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pglinter` | `1.1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 546.1 KiB | [postgresql-14-pglinter_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-14-pglinter_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pglinter` | `1.1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 417.7 KiB | [postgresql-14-pglinter_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-14-pglinter_1.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pmpetit/pglinter" title="Repository" icon="github" subtitle="github.com/pmpetit/pglinter" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pglinter-1.1.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pglinter;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pglinter;		# install via package name, for the active PG version

pig install pglinter -v 18;   # install for PG 18
pig install pglinter -v 17;   # install for PG 17
pig install pglinter -v 16;   # install for PG 16
pig install pglinter -v 15;   # install for PG 15
pig install pglinter -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pglinter;
```




## Usage

> [pglinter: PostgreSQL Linting and Analysis Extension](https://github.com/pmpetit/pglinter)

pglinter analyzes your database for potential issues, performance problems, and best practice violations. It outputs results in SARIF 2.1.0 format.

### Run Checks

```sql
SELECT pglinter.check();                                -- Run all enabled rules
SELECT pglinter.check_rule('B001');                     -- Run a specific rule
SELECT pglinter.check('/path/to/results.sarif');        -- Save SARIF report to file
SELECT pglinter.check_rule('B001', '/path/to/b001.sarif');
```

### Rule Management

```sql
SELECT pglinter.show_rules();                -- Show all rules and their status
SELECT pglinter.explain_rule('B001');        -- Get rule details and suggested fixes
SELECT pglinter.enable_rule('B001');         -- Enable a specific rule
SELECT pglinter.disable_rule('B001');        -- Disable a specific rule
SELECT pglinter.is_rule_enabled('B001');     -- Check if a rule is enabled
SELECT pglinter.enable_all_rules();
SELECT pglinter.disable_all_rules();
```

### Rule Configuration

```sql
SELECT pglinter.update_rule_levels('B001', 30, 70);   -- Set warning/error thresholds
SELECT pglinter.get_rule_levels('B001');               -- Get current thresholds
SELECT pglinter.export_rules_to_yaml();                -- Export rules to YAML
SELECT pglinter.import_rules_from_yaml('yaml...');     -- Import rules from YAML
```

### Available Rules

**Base (B-series):** B001 tables without PK, B002 redundant indexes, B003 missing FK indexes, B004 unused indexes, B005 uppercase names, B006 unused tables, B007 cross-schema FKs, B008 FK type mismatches, B009 shared trigger functions, B010 reserved keywords, B011 multiple owners per schema.

**Cluster (C-series):** C002 insecure pg_hba.conf entries, C003 MD5 password encryption.

**Schema (S-series):** S001 no default role grants, S002 env prefixes/suffixes, S003 unsecured public schema, S004 system role ownership, S005 multiple owners per schema.
