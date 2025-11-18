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
| **5090** | {{< badge content="pglinter" link="https://github.com/pmpetit/pglinter" >}} | {{< ext "pglinter" >}} | `1.0.0` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pglinter` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pglinter_18" "green" >}} {{< bg "17" "pglinter_17" "green" >}} {{< bg "16" "pglinter_16" "green" >}} {{< bg "15" "pglinter_15" "green" >}} {{< bg "14" "pglinter_14" "green" >}} {{< bg "13" "pglinter_13" "green" >}} | `pglinter_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pglinter" "green" >}} {{< bg "17" "postgresql-17-pglinter" "green" >}} {{< bg "16" "postgresql-16-pglinter" "green" >}} {{< bg "15" "postgresql-15-pglinter" "green" >}} {{< bg "14" "postgresql-14-pglinter" "green" >}} {{< bg "13" "postgresql-13-pglinter" "green" >}} | `postgresql-$v-pglinter` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pglinter_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pglinter_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pglinter_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pglinter_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "pglinter_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "pglinter_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglinter_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pglinter : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pglinter : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pglinter : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pglinter : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pglinter : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pglinter : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pglinter : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglinter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pglinter : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-pglinter` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 449.1 KiB | [postgresql-18-pglinter_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-18-pglinter_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pglinter` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 329.0 KiB | [postgresql-18-pglinter_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-18-pglinter_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pglinter` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 449.2 KiB | [postgresql-18-pglinter_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-18-pglinter_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pglinter` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 328.8 KiB | [postgresql-18-pglinter_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-18-pglinter_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pglinter` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 498.8 KiB | [postgresql-18-pglinter_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-18-pglinter_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pglinter` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 382.2 KiB | [postgresql-18-pglinter_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-18-pglinter_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pglinter` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 493.2 KiB | [postgresql-18-pglinter_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-18-pglinter_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pglinter` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 377.8 KiB | [postgresql-18-pglinter_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-18-pglinter_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-17-pglinter` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 449.3 KiB | [postgresql-17-pglinter_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-17-pglinter_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pglinter` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 328.9 KiB | [postgresql-17-pglinter_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-17-pglinter_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pglinter` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 449.3 KiB | [postgresql-17-pglinter_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-17-pglinter_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pglinter` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 329.1 KiB | [postgresql-17-pglinter_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-17-pglinter_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pglinter` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 498.8 KiB | [postgresql-17-pglinter_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-17-pglinter_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pglinter` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 381.9 KiB | [postgresql-17-pglinter_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-17-pglinter_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pglinter` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 493.1 KiB | [postgresql-17-pglinter_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-17-pglinter_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pglinter` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 378.0 KiB | [postgresql-17-pglinter_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-17-pglinter_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-16-pglinter` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 449.3 KiB | [postgresql-16-pglinter_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-16-pglinter_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pglinter` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 328.8 KiB | [postgresql-16-pglinter_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-16-pglinter_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pglinter` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 449.1 KiB | [postgresql-16-pglinter_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-16-pglinter_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pglinter` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 328.7 KiB | [postgresql-16-pglinter_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-16-pglinter_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pglinter` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 498.5 KiB | [postgresql-16-pglinter_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-16-pglinter_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pglinter` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 381.8 KiB | [postgresql-16-pglinter_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-16-pglinter_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pglinter` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 493.5 KiB | [postgresql-16-pglinter_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-16-pglinter_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pglinter` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 377.9 KiB | [postgresql-16-pglinter_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-16-pglinter_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-15-pglinter` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 448.6 KiB | [postgresql-15-pglinter_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-15-pglinter_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pglinter` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 328.8 KiB | [postgresql-15-pglinter_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-15-pglinter_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pglinter` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 448.6 KiB | [postgresql-15-pglinter_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-15-pglinter_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pglinter` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 328.9 KiB | [postgresql-15-pglinter_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-15-pglinter_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pglinter` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 498.2 KiB | [postgresql-15-pglinter_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-15-pglinter_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pglinter` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 381.8 KiB | [postgresql-15-pglinter_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-15-pglinter_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pglinter` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 492.7 KiB | [postgresql-15-pglinter_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-15-pglinter_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pglinter` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 377.7 KiB | [postgresql-15-pglinter_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-15-pglinter_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-14-pglinter` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 448.7 KiB | [postgresql-14-pglinter_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-14-pglinter_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pglinter` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 328.8 KiB | [postgresql-14-pglinter_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-14-pglinter_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pglinter` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 448.7 KiB | [postgresql-14-pglinter_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-14-pglinter_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pglinter` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 328.8 KiB | [postgresql-14-pglinter_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-14-pglinter_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pglinter` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 498.3 KiB | [postgresql-14-pglinter_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-14-pglinter_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pglinter` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 382.1 KiB | [postgresql-14-pglinter_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-14-pglinter_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pglinter` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 493.1 KiB | [postgresql-14-pglinter_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-14-pglinter_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pglinter` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 377.8 KiB | [postgresql-14-pglinter_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-14-pglinter_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-13-pglinter` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 449.2 KiB | [postgresql-13-pglinter_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-13-pglinter_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pglinter` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 328.7 KiB | [postgresql-13-pglinter_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglinter/postgresql-13-pglinter_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pglinter` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 449.2 KiB | [postgresql-13-pglinter_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-13-pglinter_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pglinter` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 328.9 KiB | [postgresql-13-pglinter_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglinter/postgresql-13-pglinter_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pglinter` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 499.1 KiB | [postgresql-13-pglinter_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-13-pglinter_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pglinter` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 381.9 KiB | [postgresql-13-pglinter_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglinter/postgresql-13-pglinter_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pglinter` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 493.7 KiB | [postgresql-13-pglinter_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-13-pglinter_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pglinter` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 377.8 KiB | [postgresql-13-pglinter_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglinter/postgresql-13-pglinter_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pmpetit/pglinter" title="Repository" icon="github" subtitle="github.com/pmpetit/pglinter" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pglinter-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pglinter;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pglinter;		# install via package name, for the active PG version

pig install pglinter -v 18;   # install for PG 18
pig install pglinter -v 17;   # install for PG 17
pig install pglinter -v 16;   # install for PG 16
pig install pglinter -v 15;   # install for PG 15
pig install pglinter -v 14;   # install for PG 14
pig install pglinter -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pglinter;
```
