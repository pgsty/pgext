---
title: "ulak"
linkTitle: "ulak"
description: "Transactional Outbox extension for PostgreSQL with reliable asynchronous delivery"
weight: 2680
categories: ["FEAT"]
width: full
---

[**ulak**](https://github.com/zeybek/ulak) : Transactional Outbox extension for PostgreSQL with reliable asynchronous delivery


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2680** | {{< badge content="ulak" link="https://github.com/zeybek/ulak" >}} | {{< ext "ulak" >}} | `0.0.2` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `ulak` |

> [!Note] preload required


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `ulak` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.2` | {{< bg "18" "ulak_18" "green" >}} {{< bg "17" "ulak_17" "green" >}} {{< bg "16" "ulak_16" "green" >}} {{< bg "15" "ulak_15" "green" >}} {{< bg "14" "ulak_14" "green" >}} | `ulak_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.2` | {{< bg "18" "postgresql-18-ulak" "green" >}} {{< bg "17" "postgresql-17-ulak" "green" >}} {{< bg "16" "postgresql-16-ulak" "green" >}} {{< bg "15" "postgresql-15-ulak" "green" >}} {{< bg "14" "postgresql-14-ulak" "green" >}} | `postgresql-$v-ulak` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "ulak_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "ulak_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "ulak_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "ulak_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "ulak_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "ulak_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "ulak_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-ulak : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-ulak : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-ulak : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-ulak : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-ulak : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-ulak : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-ulak : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-ulak : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-ulak : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-ulak : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-ulak : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-ulak : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-ulak : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-ulak : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-ulak : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-ulak : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-ulak : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-ulak : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ulak_18` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 108.1 KiB | [ulak_18-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ulak_18-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `ulak_18` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 106.4 KiB | [ulak_18-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ulak_18-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `ulak_18` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 104.6 KiB | [ulak_18-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ulak_18-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `ulak_18` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 104.3 KiB | [ulak_18-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ulak_18-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `ulak_18` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 104.9 KiB | [ulak_18-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ulak_18-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `ulak_18` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 104.6 KiB | [ulak_18-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ulak_18-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-ulak` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 273.7 KiB | [postgresql-18-ulak_0.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-18-ulak_0.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-ulak` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 266.9 KiB | [postgresql-18-ulak_0.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-18-ulak_0.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-ulak` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 298.2 KiB | [postgresql-18-ulak_0.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-18-ulak_0.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-ulak` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 291.0 KiB | [postgresql-18-ulak_0.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-18-ulak_0.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-ulak` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 288.2 KiB | [postgresql-18-ulak_0.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-18-ulak_0.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-ulak` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 284.2 KiB | [postgresql-18-ulak_0.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-18-ulak_0.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-ulak` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 305.9 KiB | [postgresql-18-ulak_0.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-18-ulak_0.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-ulak` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 302.1 KiB | [postgresql-18-ulak_0.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-18-ulak_0.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ulak_17` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 108.1 KiB | [ulak_17-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ulak_17-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `ulak_17` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 106.4 KiB | [ulak_17-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ulak_17-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `ulak_17` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 105.1 KiB | [ulak_17-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ulak_17-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `ulak_17` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 104.3 KiB | [ulak_17-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ulak_17-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `ulak_17` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 105.0 KiB | [ulak_17-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ulak_17-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `ulak_17` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 104.6 KiB | [ulak_17-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ulak_17-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-ulak` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 273.9 KiB | [postgresql-17-ulak_0.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-17-ulak_0.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-ulak` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 266.8 KiB | [postgresql-17-ulak_0.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-17-ulak_0.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-ulak` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 298.2 KiB | [postgresql-17-ulak_0.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-17-ulak_0.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-ulak` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 291.0 KiB | [postgresql-17-ulak_0.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-17-ulak_0.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-ulak` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 313.7 KiB | [postgresql-17-ulak_0.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-17-ulak_0.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-ulak` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 309.9 KiB | [postgresql-17-ulak_0.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-17-ulak_0.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-ulak` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 306.1 KiB | [postgresql-17-ulak_0.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-17-ulak_0.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-ulak` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 302.1 KiB | [postgresql-17-ulak_0.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-17-ulak_0.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ulak_16` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 108.1 KiB | [ulak_16-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ulak_16-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `ulak_16` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 106.3 KiB | [ulak_16-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ulak_16-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `ulak_16` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 105.1 KiB | [ulak_16-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ulak_16-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `ulak_16` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 104.3 KiB | [ulak_16-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ulak_16-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `ulak_16` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 105.0 KiB | [ulak_16-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ulak_16-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `ulak_16` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 104.6 KiB | [ulak_16-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ulak_16-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-ulak` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 273.7 KiB | [postgresql-16-ulak_0.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-16-ulak_0.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-ulak` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 267.0 KiB | [postgresql-16-ulak_0.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-16-ulak_0.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-ulak` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 298.3 KiB | [postgresql-16-ulak_0.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-16-ulak_0.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-ulak` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 291.2 KiB | [postgresql-16-ulak_0.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-16-ulak_0.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-ulak` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 313.2 KiB | [postgresql-16-ulak_0.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-16-ulak_0.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-ulak` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 309.4 KiB | [postgresql-16-ulak_0.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-16-ulak_0.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-ulak` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 306.1 KiB | [postgresql-16-ulak_0.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-16-ulak_0.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-ulak` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 302.2 KiB | [postgresql-16-ulak_0.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-16-ulak_0.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ulak_15` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 112.1 KiB | [ulak_15-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ulak_15-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `ulak_15` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 109.9 KiB | [ulak_15-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ulak_15-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `ulak_15` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 115.0 KiB | [ulak_15-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ulak_15-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `ulak_15` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 114.7 KiB | [ulak_15-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ulak_15-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `ulak_15` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 115.7 KiB | [ulak_15-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ulak_15-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `ulak_15` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 115.2 KiB | [ulak_15-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ulak_15-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-ulak` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 277.7 KiB | [postgresql-15-ulak_0.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-15-ulak_0.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-ulak` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 270.1 KiB | [postgresql-15-ulak_0.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-15-ulak_0.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-ulak` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 302.2 KiB | [postgresql-15-ulak_0.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-15-ulak_0.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-ulak` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 294.8 KiB | [postgresql-15-ulak_0.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-15-ulak_0.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-ulak` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 321.1 KiB | [postgresql-15-ulak_0.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-15-ulak_0.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-ulak` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 317.5 KiB | [postgresql-15-ulak_0.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-15-ulak_0.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-ulak` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 314.1 KiB | [postgresql-15-ulak_0.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-15-ulak_0.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-ulak` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 310.7 KiB | [postgresql-15-ulak_0.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-15-ulak_0.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ulak_14` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 112.0 KiB | [ulak_14-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ulak_14-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `ulak_14` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 109.8 KiB | [ulak_14-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ulak_14-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `ulak_14` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 114.9 KiB | [ulak_14-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ulak_14-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `ulak_14` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 114.8 KiB | [ulak_14-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ulak_14-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `ulak_14` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 115.7 KiB | [ulak_14-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ulak_14-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `ulak_14` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 115.2 KiB | [ulak_14-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ulak_14-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-ulak` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 277.1 KiB | [postgresql-14-ulak_0.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-14-ulak_0.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-ulak` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 270.0 KiB | [postgresql-14-ulak_0.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-14-ulak_0.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-ulak` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 301.8 KiB | [postgresql-14-ulak_0.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-14-ulak_0.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-ulak` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 294.4 KiB | [postgresql-14-ulak_0.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-14-ulak_0.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-ulak` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 319.9 KiB | [postgresql-14-ulak_0.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-14-ulak_0.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-ulak` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 316.6 KiB | [postgresql-14-ulak_0.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-14-ulak_0.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-ulak` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 313.8 KiB | [postgresql-14-ulak_0.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-14-ulak_0.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-ulak` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 310.6 KiB | [postgresql-14-ulak_0.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-14-ulak_0.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/zeybek/ulak" title="Repository" icon="github" subtitle="github.com/zeybek/ulak" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="ulak-0.0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg ulak;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install ulak;		# install via package name, for the active PG version

pig install ulak -v 18;   # install for PG 18
pig install ulak -v 17;   # install for PG 17
pig install ulak -v 16;   # install for PG 16
pig install ulak -v 15;   # install for PG 15
pig install ulak -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'ulak';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION ulak;
```


## Usage

> Sources: [README](https://github.com/zeybek/ulak/blob/main/README.md), [wiki](https://github.com/zeybek/ulak/wiki), [v0.0.2 release](https://github.com/zeybek/ulak/releases/tag/v0.0.2)

`ulak` implements the transactional outbox pattern inside PostgreSQL. Messages are inserted transactionally, then background workers deliver them asynchronously with retries, circuit breaking, and dead-letter handling.

### Enable the Extension

```sql
-- postgresql.conf
shared_preload_libraries = 'ulak'

CREATE EXTENSION ulak;
```

The README also shows setting `ulak.database` when running the included Docker example.

### Define Endpoints and Send Messages

```sql
SELECT ulak.create_endpoint(
  'my-webhook',
  'http',
  '{"url":"https://httpbin.org/post","method":"POST"}'::jsonb
);

BEGIN;
  INSERT INTO orders (id, total) VALUES (1, 99.99);
  SELECT ulak.send('my-webhook', '{"order_id":1,"total":99.99}'::jsonb);
COMMIT;
```

The same API pattern is documented for `kafka`, `mqtt`, `redis`, `amqp`, and `nats` endpoints.

### Delivery Controls and Pub/Sub

```sql
SELECT ulak.send_with_options(
  'my-webhook',
  '{"event":"order.created"}'::jsonb,
  5,
  NOW() + INTERVAL '10 minutes',
  'order-123-created',
  '550e8400-e29b-41d4-a716-446655440000'::uuid,
  NOW() + INTERVAL '1 hour',
  'order-123'
);

SELECT ulak.send_batch('my-webhook', ARRAY['{"id":1}'::jsonb, '{"id":2}'::jsonb]);

SELECT ulak.create_event_type('order.created', 'New order placed');
SELECT ulak.subscribe('order.created', 'my-webhook');
SELECT ulak.publish('order.created', '{"order_id":123}'::jsonb);
```

### Monitoring and Recovery

```sql
SELECT * FROM ulak.health_check();
SELECT * FROM ulak.get_worker_status();
SELECT * FROM ulak.get_endpoint_health();
SELECT * FROM ulak.dlq_summary();

SELECT ulak.redrive_message(42);
SELECT ulak.redrive_endpoint('my-webhook');
SELECT ulak.redrive_all();
SELECT ulak.replay_message(100);
```

### Key Settings

The README highlights these `ulak.` GUCs:

- `workers`
- `poll_interval`
- `batch_size`
- `max_queue_size`
- `circuit_breaker_threshold`
- `circuit_breaker_cooldown`
- `http_timeout`
- `dlq_retention_days`

The `v0.0.2` release notes only document a Docker publish workflow fix, so the current README and wiki remain the authoritative user-facing usage sources.
