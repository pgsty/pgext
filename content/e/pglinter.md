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
| **5090** | {{< badge content="pglinter" link="https://github.com/pmpetit/pglinter" >}} | {{< ext "pglinter" >}} | `1.1.0` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "amcheck" >}} {{< ext "supautils" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pglinter` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "pglinter_18" "green" >}} {{< bg "17" "pglinter_17" "green" >}} {{< bg "16" "pglinter_16" "green" >}} {{< bg "15" "pglinter_15" "green" >}} {{< bg "14" "pglinter_14" "green" >}} {{< bg "13" "pglinter_13" "red" >}} | `pglinter_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "postgresql-18-pglinter" "green" >}} {{< bg "17" "postgresql-17-pglinter" "green" >}} {{< bg "16" "postgresql-16-pglinter" "green" >}} {{< bg "15" "postgresql-15-pglinter" "green" >}} {{< bg "14" "postgresql-14-pglinter" "green" >}} {{< bg "13" "postgresql-13-pglinter" "red" >}} | `postgresql-$v-pglinter` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.1" "pglinter_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.1" "pglinter_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.1" "pglinter_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.1" "pglinter_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.1" "pglinter_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pglinter_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.1" "pglinter_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pglinter : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pglinter : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pglinter : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pglinter : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pglinter : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pglinter : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pglinter : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pglinter : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglinter_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 580.2 KiB | [pglinter_18-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_18-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_18` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 532.9 KiB | [pglinter_18-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_18-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 442.9 KiB | [pglinter_18-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_18-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_18` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 406.2 KiB | [pglinter_18-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_18-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 594.2 KiB | [pglinter_18-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_18-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_18` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 546.7 KiB | [pglinter_18-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_18-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 471.4 KiB | [pglinter_18-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_18-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_18` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 430.7 KiB | [pglinter_18-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_18-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 594.4 KiB | [pglinter_18-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_18-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_18` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 546.9 KiB | [pglinter_18-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_18-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 471.6 KiB | [pglinter_18-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_18-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `pglinter_18` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 431.5 KiB | [pglinter_18-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_18-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pglinter` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 489.2 KiB | [postgresql-18-pglinter_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-18-pglinter_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pglinter` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 358.9 KiB | [postgresql-18-pglinter_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-18-pglinter_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pglinter` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 489.2 KiB | [postgresql-18-pglinter_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-18-pglinter_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pglinter` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 359.0 KiB | [postgresql-18-pglinter_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-18-pglinter_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pglinter` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 542.1 KiB | [postgresql-18-pglinter_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-18-pglinter_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pglinter` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 417.8 KiB | [postgresql-18-pglinter_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-18-pglinter_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pglinter` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 537.3 KiB | [postgresql-18-pglinter_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-18-pglinter_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pglinter` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 412.6 KiB | [postgresql-18-pglinter_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-18-pglinter_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglinter_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 580.0 KiB | [pglinter_17-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_17-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_17` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 532.7 KiB | [pglinter_17-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_17-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 441.3 KiB | [pglinter_17-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_17-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_17` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 406.4 KiB | [pglinter_17-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_17-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 594.3 KiB | [pglinter_17-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_17-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_17` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 546.8 KiB | [pglinter_17-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_17-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 469.6 KiB | [pglinter_17-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_17-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_17` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 430.9 KiB | [pglinter_17-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_17-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 594.3 KiB | [pglinter_17-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_17-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_17` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 546.7 KiB | [pglinter_17-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_17-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 469.5 KiB | [pglinter_17-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_17-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `pglinter_17` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 430.9 KiB | [pglinter_17-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_17-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pglinter` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 489.3 KiB | [postgresql-17-pglinter_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-17-pglinter_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pglinter` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 358.0 KiB | [postgresql-17-pglinter_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-17-pglinter_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pglinter` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 489.2 KiB | [postgresql-17-pglinter_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-17-pglinter_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pglinter` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 357.9 KiB | [postgresql-17-pglinter_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-17-pglinter_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pglinter` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 541.8 KiB | [postgresql-17-pglinter_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-17-pglinter_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pglinter` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 416.4 KiB | [postgresql-17-pglinter_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-17-pglinter_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pglinter` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 537.2 KiB | [postgresql-17-pglinter_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-17-pglinter_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pglinter` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 411.4 KiB | [postgresql-17-pglinter_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-17-pglinter_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglinter_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 578.4 KiB | [pglinter_16-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_16-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_16` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 532.9 KiB | [pglinter_16-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_16-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 441.2 KiB | [pglinter_16-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_16-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_16` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 406.4 KiB | [pglinter_16-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_16-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 592.6 KiB | [pglinter_16-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_16-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_16` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 547.1 KiB | [pglinter_16-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_16-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 469.4 KiB | [pglinter_16-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_16-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_16` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 430.9 KiB | [pglinter_16-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_16-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 592.8 KiB | [pglinter_16-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_16-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_16` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 547.3 KiB | [pglinter_16-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_16-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 469.3 KiB | [pglinter_16-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_16-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `pglinter_16` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 430.9 KiB | [pglinter_16-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_16-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pglinter` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 487.8 KiB | [postgresql-16-pglinter_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-16-pglinter_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pglinter` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 358.0 KiB | [postgresql-16-pglinter_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-16-pglinter_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pglinter` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 487.8 KiB | [postgresql-16-pglinter_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-16-pglinter_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pglinter` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 357.8 KiB | [postgresql-16-pglinter_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-16-pglinter_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pglinter` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 540.7 KiB | [postgresql-16-pglinter_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-16-pglinter_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pglinter` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 416.2 KiB | [postgresql-16-pglinter_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-16-pglinter_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pglinter` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 535.8 KiB | [postgresql-16-pglinter_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-16-pglinter_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pglinter` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 410.8 KiB | [postgresql-16-pglinter_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-16-pglinter_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglinter_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 578.2 KiB | [pglinter_15-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_15-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_15` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 532.7 KiB | [pglinter_15-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_15-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 441.6 KiB | [pglinter_15-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_15-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_15` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 406.5 KiB | [pglinter_15-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_15-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 592.5 KiB | [pglinter_15-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_15-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_15` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 546.7 KiB | [pglinter_15-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_15-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 469.6 KiB | [pglinter_15-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_15-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_15` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 431.0 KiB | [pglinter_15-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_15-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 592.6 KiB | [pglinter_15-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_15-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_15` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 546.9 KiB | [pglinter_15-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_15-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 469.3 KiB | [pglinter_15-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_15-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `pglinter_15` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 431.6 KiB | [pglinter_15-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_15-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pglinter` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 487.6 KiB | [postgresql-15-pglinter_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-15-pglinter_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pglinter` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 358.1 KiB | [postgresql-15-pglinter_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-15-pglinter_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pglinter` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 487.8 KiB | [postgresql-15-pglinter_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-15-pglinter_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pglinter` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 358.0 KiB | [postgresql-15-pglinter_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-15-pglinter_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pglinter` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 540.5 KiB | [postgresql-15-pglinter_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-15-pglinter_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pglinter` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 416.2 KiB | [postgresql-15-pglinter_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-15-pglinter_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pglinter` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 535.6 KiB | [postgresql-15-pglinter_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-15-pglinter_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pglinter` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 411.4 KiB | [postgresql-15-pglinter_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-15-pglinter_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglinter_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 578.3 KiB | [pglinter_14-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_14-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_14` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 532.7 KiB | [pglinter_14-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_14-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 441.2 KiB | [pglinter_14-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_14-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_14` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 406.5 KiB | [pglinter_14-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_14-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 591.6 KiB | [pglinter_14-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_14-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_14` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 546.6 KiB | [pglinter_14-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_14-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 469.3 KiB | [pglinter_14-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_14-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_14` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 431.8 KiB | [pglinter_14-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_14-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 591.7 KiB | [pglinter_14-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_14-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_14` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 546.8 KiB | [pglinter_14-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_14-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 469.3 KiB | [pglinter_14-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_14-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `pglinter_14` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 431.6 KiB | [pglinter_14-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_14-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pglinter` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 487.4 KiB | [postgresql-14-pglinter_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-14-pglinter_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pglinter` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 357.9 KiB | [postgresql-14-pglinter_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-14-pglinter_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pglinter` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 487.7 KiB | [postgresql-14-pglinter_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-14-pglinter_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pglinter` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 357.7 KiB | [postgresql-14-pglinter_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-14-pglinter_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pglinter` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 540.3 KiB | [postgresql-14-pglinter_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-14-pglinter_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pglinter` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 416.0 KiB | [postgresql-14-pglinter_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-14-pglinter_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pglinter` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 535.3 KiB | [postgresql-14-pglinter_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-14-pglinter_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pglinter` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 410.8 KiB | [postgresql-14-pglinter_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-14-pglinter_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglinter_13` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 533.1 KiB | [pglinter_13-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglinter_13-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pglinter_13` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 406.5 KiB | [pglinter_13-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglinter_13-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pglinter_13` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 546.7 KiB | [pglinter_13-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglinter_13-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pglinter_13` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 431.0 KiB | [pglinter_13-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglinter_13-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pglinter_13` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 546.9 KiB | [pglinter_13-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglinter_13-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pglinter_13` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 431.0 KiB | [pglinter_13-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglinter_13-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pglinter` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 449.0 KiB | [postgresql-13-pglinter_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-13-pglinter_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pglinter` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 327.8 KiB | [postgresql-13-pglinter_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-13-pglinter_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pglinter` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 449.2 KiB | [postgresql-13-pglinter_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-13-pglinter_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pglinter` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 327.9 KiB | [postgresql-13-pglinter_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-13-pglinter_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pglinter` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 499.0 KiB | [postgresql-13-pglinter_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-13-pglinter_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pglinter` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 380.7 KiB | [postgresql-13-pglinter_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-13-pglinter_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pglinter` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 493.8 KiB | [postgresql-13-pglinter_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-13-pglinter_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pglinter` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 376.7 KiB | [postgresql-13-pglinter_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-13-pglinter_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pmpetit/pglinter" title="Repository" icon="github" subtitle="github.com/pmpetit/pglinter" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pglinter-1.1.0.tar.gz" >}}
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
