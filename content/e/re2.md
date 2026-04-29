---
title: "re2"
linkTitle: "re2"
description: "ClickHouse-compatible regex functions using RE2"
weight: 4235
categories: ["UTIL"]
width: full
---

[**re2**](https://github.com/ClickHouse/pg_re2) : ClickHouse-compatible regex functions using RE2


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4235** | {{< badge content="re2" link="https://github.com/ClickHouse/pg_re2" >}} | {{< ext "re2" >}} | `0.1.1` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |

> [!Note] release 0.1.1; SQL v0.1


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `re2` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "re2_18" "green" >}} {{< bg "17" "re2_17" "green" >}} {{< bg "16" "re2_16" "green" >}} {{< bg "15" "re2_15" "red" >}} {{< bg "14" "re2_14" "red" >}} | `re2_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "postgresql-18-re2" "green" >}} {{< bg "17" "postgresql-17-re2" "green" >}} {{< bg "16" "postgresql-16-re2" "green" >}} {{< bg "15" "postgresql-15-re2" "red" >}} {{< bg "14" "postgresql-14-re2" "red" >}} | `postgresql-$v-re2` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "re2_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "re2_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "re2_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "re2_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "re2_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "re2_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "re2_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "re2_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "re2_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "re2_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "re2_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "re2_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "re2_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "re2_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "re2_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "re2_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "re2_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "re2_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "re2_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "re2_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "re2_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "re2_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "re2_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "re2_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "re2_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "re2_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "re2_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "re2_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "re2_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "re2_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-re2 : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-re2 : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-re2 : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-re2 : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-re2 : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-re2 : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-re2 : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-re2 : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-re2 : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-re2 : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-re2 : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-re2 : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-re2 : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-re2 : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-re2 : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-re2 : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-re2 : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-re2 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-re2 : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `re2_18` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 25.7 KiB | [re2_18-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/re2_18-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `re2_18` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 25.1 KiB | [re2_18-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/re2_18-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `re2_18` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.8 KiB | [re2_18-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/re2_18-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `re2_18` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.1 KiB | [re2_18-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/re2_18-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `re2_18` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.8 KiB | [re2_18-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/re2_18-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `re2_18` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.3 KiB | [re2_18-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/re2_18-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-re2` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.3 KiB | [postgresql-18-re2_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/re2/postgresql-18-re2_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-re2` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 36.6 KiB | [postgresql-18-re2_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/re2/postgresql-18-re2_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-re2` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 38.6 KiB | [postgresql-18-re2_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/re2/postgresql-18-re2_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-re2` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.6 KiB | [postgresql-18-re2_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/re2/postgresql-18-re2_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-re2` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 38.1 KiB | [postgresql-18-re2_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/re2/postgresql-18-re2_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-re2` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.5 KiB | [postgresql-18-re2_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/re2/postgresql-18-re2_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-re2` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 37.6 KiB | [postgresql-18-re2_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/re2/postgresql-18-re2_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-re2` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 37.3 KiB | [postgresql-18-re2_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/re2/postgresql-18-re2_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `re2_17` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 25.7 KiB | [re2_17-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/re2_17-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `re2_17` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 25.1 KiB | [re2_17-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/re2_17-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `re2_17` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.8 KiB | [re2_17-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/re2_17-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `re2_17` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.0 KiB | [re2_17-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/re2_17-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `re2_17` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.8 KiB | [re2_17-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/re2_17-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `re2_17` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.3 KiB | [re2_17-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/re2_17-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-re2` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.1 KiB | [postgresql-17-re2_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/re2/postgresql-17-re2_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-re2` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 36.4 KiB | [postgresql-17-re2_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/re2/postgresql-17-re2_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-re2` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 38.3 KiB | [postgresql-17-re2_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/re2/postgresql-17-re2_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-re2` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.5 KiB | [postgresql-17-re2_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/re2/postgresql-17-re2_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-re2` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 39.6 KiB | [postgresql-17-re2_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/re2/postgresql-17-re2_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-re2` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 38.8 KiB | [postgresql-17-re2_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/re2/postgresql-17-re2_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-re2` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 37.4 KiB | [postgresql-17-re2_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/re2/postgresql-17-re2_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-re2` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 37.1 KiB | [postgresql-17-re2_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/re2/postgresql-17-re2_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `re2_16` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 25.7 KiB | [re2_16-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/re2_16-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `re2_16` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 25.1 KiB | [re2_16-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/re2_16-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `re2_16` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.8 KiB | [re2_16-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/re2_16-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `re2_16` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.0 KiB | [re2_16-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/re2_16-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `re2_16` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.8 KiB | [re2_16-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/re2_16-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `re2_16` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.3 KiB | [re2_16-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/re2_16-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-re2` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.1 KiB | [postgresql-16-re2_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/re2/postgresql-16-re2_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-re2` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 36.4 KiB | [postgresql-16-re2_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/re2/postgresql-16-re2_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-re2` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 38.3 KiB | [postgresql-16-re2_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/re2/postgresql-16-re2_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-re2` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.4 KiB | [postgresql-16-re2_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/re2/postgresql-16-re2_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-re2` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 39.6 KiB | [postgresql-16-re2_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/re2/postgresql-16-re2_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-re2` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 38.8 KiB | [postgresql-16-re2_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/re2/postgresql-16-re2_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-re2` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 37.4 KiB | [postgresql-16-re2_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/re2/postgresql-16-re2_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-re2` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 37.2 KiB | [postgresql-16-re2_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/re2/postgresql-16-re2_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ClickHouse/pg_re2" title="Repository" icon="github" subtitle="github.com/ClickHouse/pg_re2" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="re2-0.1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg re2;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install re2;		# install via package name, for the active PG version

pig install re2 -v 18;   # install for PG 18
pig install re2 -v 17;   # install for PG 17
pig install re2 -v 16;   # install for PG 16
pig install re2 -v 15;   # install for PG 15
pig install re2 -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION re2;
```

## Usage

Sources: [official README](https://github.com/ClickHouse/pg_re2/blob/main/README.md), [official reference doc](https://github.com/ClickHouse/pg_re2/blob/main/doc/re2.md), [v0.1.1 release](https://github.com/ClickHouse/pg_re2/releases/tag/v0.1.1)

`re2` provides ClickHouse-compatible regular expression functions backed by Google's RE2 engine. It exposes both `text` and `bytea` overloads, so binary data with `\\0` bytes can be searched too.

```sql
CREATE EXTENSION re2;

SELECT re2match('hello world', 'h.*o');
SELECT re2extract('Order #123', '(\\d+)');
SELECT re2countmatches('a1 b2 c3', '\\d');
```

### Core Functions

- `re2match(haystack, pattern) -> boolean`
- `re2extract(haystack, pattern) -> text|bytea`
- `re2extractall(haystack, pattern) -> text[]|bytea[]`
- `re2regexpextract(haystack, pattern, index default 1) -> text|bytea`
- `re2extractgroups(haystack, pattern) -> text[]|bytea[]`
- `re2replaceregexpone(haystack, pattern, replacement) -> text|bytea`
- `re2replaceregexpall(haystack, pattern, replacement) -> text|bytea`
- `re2countmatches(...)` and `re2countmatchescaseinsensitive(...)`

### Multi-Pattern Matching

The `re2multimatch*` family accepts either multiple pattern arguments or a `VARIADIC` array:

```sql
SELECT re2multimatchany('error: timeout', 'timeout', 'denied');
SELECT re2multimatchanyindex('error: timeout', VARIADIC ARRAY['timeout', 'denied']);
SELECT re2multimatchallindices('error: timeout', 'error', 'timeout', 'panic');
```

### Matching Semantics

- To match ClickHouse behavior, `.` matches line breaks by default.
- Prefix the pattern with `(?-s)` if you want `.` not to cross line breaks.
- Replacement strings support `\\0` through `\\9` backreferences.

### Caveats

- Upstream requires the system `re2` library at build/install time.
- The `v0.1.1` release is binary-only: it adds PostgreSQL 13+ support and documents `VARIADIC` use for multi-pattern functions, but existing `v0.1` SQL installations do not need `ALTER EXTENSION UPDATE`.
