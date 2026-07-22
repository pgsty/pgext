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
| **4235** | {{< badge content="re2" link="https://github.com/ClickHouse/pg_re2" >}} | {{< ext "re2" >}} | `0.4.1` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |

> [!Note] Stable PGXN and PIGSTY package release 0.4.1 for PostgreSQL 16 through 18.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `re2` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.1` | {{< bg "18" "re2_18" "green" >}} {{< bg "17" "re2_17" "green" >}} {{< bg "16" "re2_16" "green" >}} {{< bg "15" "re2_15" "red" >}} {{< bg "14" "re2_14" "red" >}} | `re2_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.1` | {{< bg "18" "postgresql-18-re2" "green" >}} {{< bg "17" "postgresql-17-re2" "green" >}} {{< bg "16" "postgresql-16-re2" "green" >}} {{< bg "15" "postgresql-15-re2" "red" >}} {{< bg "14" "postgresql-14-re2" "red" >}} | `postgresql-$v-re2` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "re2_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "re2_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "re2_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "re2_15 : N/A 0" "gray" >}} | {{< bg "N/A" "re2_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "re2_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "re2_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "re2_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "re2_15 : N/A 0" "gray" >}} | {{< bg "N/A" "re2_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "re2_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "re2_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "re2_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "re2_15 : N/A 0" "gray" >}} | {{< bg "N/A" "re2_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "re2_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "re2_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "re2_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "re2_15 : N/A 0" "gray" >}} | {{< bg "N/A" "re2_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "re2_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "re2_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "re2_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "re2_15 : N/A 0" "gray" >}} | {{< bg "N/A" "re2_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "re2_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "re2_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "re2_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "re2_15 : N/A 0" "gray" >}} | {{< bg "N/A" "re2_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-re2 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-re2 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-re2 : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-re2 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-re2 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-re2 : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-re2 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-re2 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-re2 : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-re2 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-re2 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-re2 : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-re2 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-re2 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-re2 : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-re2 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-re2 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-re2 : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-re2 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-re2 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-re2 : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-re2 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-re2 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-re2 : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-re2 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-re2 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-re2 : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-re2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-re2 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-re2 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-re2 : N/A 0" "gray" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `re2_18` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 41.9 KiB | [re2_18-0.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/re2_18-0.4.1-1PIGSTY.el8.x86_64.rpm) |
| `re2_18` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 39.8 KiB | [re2_18-0.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/re2_18-0.4.1-1PIGSTY.el8.aarch64.rpm) |
| `re2_18` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 41.8 KiB | [re2_18-0.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/re2_18-0.4.1-1PIGSTY.el9.x86_64.rpm) |
| `re2_18` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 40.5 KiB | [re2_18-0.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/re2_18-0.4.1-1PIGSTY.el9.aarch64.rpm) |
| `re2_18` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.8 KiB | [re2_18-0.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/re2_18-0.4.1-1PIGSTY.el10.x86_64.rpm) |
| `re2_18` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 40.8 KiB | [re2_18-0.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/re2_18-0.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-re2` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 74.6 KiB | [postgresql-18-re2_0.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/re2/postgresql-18-re2_0.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-re2` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 72.8 KiB | [postgresql-18-re2_0.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/re2/postgresql-18-re2_0.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-re2` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 76.9 KiB | [postgresql-18-re2_0.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/re2/postgresql-18-re2_0.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-re2` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 74.6 KiB | [postgresql-18-re2_0.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/re2/postgresql-18-re2_0.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-re2` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 77.2 KiB | [postgresql-18-re2_0.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/re2/postgresql-18-re2_0.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-re2` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 74.8 KiB | [postgresql-18-re2_0.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/re2/postgresql-18-re2_0.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-re2` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 74.9 KiB | [postgresql-18-re2_0.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/re2/postgresql-18-re2_0.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-re2` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 74.1 KiB | [postgresql-18-re2_0.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/re2/postgresql-18-re2_0.4.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-re2` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 76.1 KiB | [postgresql-18-re2_0.4.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/re2/postgresql-18-re2_0.4.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-re2` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 74.4 KiB | [postgresql-18-re2_0.4.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/re2/postgresql-18-re2_0.4.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `re2_17` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 41.8 KiB | [re2_17-0.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/re2_17-0.4.1-1PIGSTY.el8.x86_64.rpm) |
| `re2_17` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 39.8 KiB | [re2_17-0.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/re2_17-0.4.1-1PIGSTY.el8.aarch64.rpm) |
| `re2_17` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 41.8 KiB | [re2_17-0.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/re2_17-0.4.1-1PIGSTY.el9.x86_64.rpm) |
| `re2_17` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 40.5 KiB | [re2_17-0.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/re2_17-0.4.1-1PIGSTY.el9.aarch64.rpm) |
| `re2_17` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.9 KiB | [re2_17-0.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/re2_17-0.4.1-1PIGSTY.el10.x86_64.rpm) |
| `re2_17` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 40.7 KiB | [re2_17-0.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/re2_17-0.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-re2` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 74.5 KiB | [postgresql-17-re2_0.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/re2/postgresql-17-re2_0.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-re2` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 72.8 KiB | [postgresql-17-re2_0.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/re2/postgresql-17-re2_0.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-re2` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 76.8 KiB | [postgresql-17-re2_0.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/re2/postgresql-17-re2_0.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-re2` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 74.6 KiB | [postgresql-17-re2_0.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/re2/postgresql-17-re2_0.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-re2` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 83.2 KiB | [postgresql-17-re2_0.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/re2/postgresql-17-re2_0.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-re2` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 81.0 KiB | [postgresql-17-re2_0.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/re2/postgresql-17-re2_0.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-re2` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 74.8 KiB | [postgresql-17-re2_0.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/re2/postgresql-17-re2_0.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-re2` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 73.9 KiB | [postgresql-17-re2_0.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/re2/postgresql-17-re2_0.4.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-re2` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 76.0 KiB | [postgresql-17-re2_0.4.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/re2/postgresql-17-re2_0.4.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-re2` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 74.4 KiB | [postgresql-17-re2_0.4.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/re2/postgresql-17-re2_0.4.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `re2_16` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 41.8 KiB | [re2_16-0.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/re2_16-0.4.1-1PIGSTY.el8.x86_64.rpm) |
| `re2_16` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 39.8 KiB | [re2_16-0.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/re2_16-0.4.1-1PIGSTY.el8.aarch64.rpm) |
| `re2_16` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 41.8 KiB | [re2_16-0.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/re2_16-0.4.1-1PIGSTY.el9.x86_64.rpm) |
| `re2_16` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 40.5 KiB | [re2_16-0.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/re2_16-0.4.1-1PIGSTY.el9.aarch64.rpm) |
| `re2_16` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.9 KiB | [re2_16-0.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/re2_16-0.4.1-1PIGSTY.el10.x86_64.rpm) |
| `re2_16` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 40.7 KiB | [re2_16-0.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/re2_16-0.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-re2` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 74.5 KiB | [postgresql-16-re2_0.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/re2/postgresql-16-re2_0.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-re2` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 72.8 KiB | [postgresql-16-re2_0.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/re2/postgresql-16-re2_0.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-re2` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 76.8 KiB | [postgresql-16-re2_0.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/re2/postgresql-16-re2_0.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-re2` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 74.6 KiB | [postgresql-16-re2_0.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/re2/postgresql-16-re2_0.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-re2` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 83.1 KiB | [postgresql-16-re2_0.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/re2/postgresql-16-re2_0.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-re2` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 80.9 KiB | [postgresql-16-re2_0.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/re2/postgresql-16-re2_0.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-re2` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 74.8 KiB | [postgresql-16-re2_0.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/re2/postgresql-16-re2_0.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-re2` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 73.9 KiB | [postgresql-16-re2_0.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/re2/postgresql-16-re2_0.4.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-re2` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 76.1 KiB | [postgresql-16-re2_0.4.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/re2/postgresql-16-re2_0.4.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-re2` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 74.4 KiB | [postgresql-16-re2_0.4.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/re2/postgresql-16-re2_0.4.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ClickHouse/pg_re2" title="Repository" icon="github" subtitle="github.com/ClickHouse/pg_re2" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="re2-0.4.1.tar.gz" >}}
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

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION re2;
```




## Usage

Sources:

- [pg_re2 v0.4.1 README](https://github.com/ClickHouse/pg_re2/blob/v0.4.1/README.md)
- [SQL reference](https://github.com/ClickHouse/pg_re2/blob/v0.4.1/doc/re2.md)
- [v0.4.0 release](https://github.com/ClickHouse/pg_re2/releases/tag/v0.4.0)
- [v0.4.1 release](https://github.com/ClickHouse/pg_re2/releases/tag/v0.4.1)

`re2` provides ClickHouse-compatible regular-expression functions backed by Google's RE2 engine. It exposes both `text` and `bytea` overloads, so binary data containing `\\0` bytes can be searched too. Version `0.4.1` also adds index-assisted matching and reports the linked RE2 version.

```sql
CREATE EXTENSION re2;

SELECT re2match('hello world', 'h.*o');
SELECT re2extract('Order #123', '(\\d+)');
SELECT re2countmatches('a1 b2 c3', '\\d');
SELECT re2_version();
```

### Core Functions

- `re2match(haystack, pattern) -> boolean`
- `re2extract(haystack, pattern) -> text|bytea`
- `re2extractall(haystack, pattern) -> text[]|bytea[]`
- `re2regexpextract(haystack, pattern, index default 1) -> text|bytea`
- `re2extractgroups(haystack, pattern) -> text[]|bytea[]`
- `re2extractallgroupsvertical(haystack, pattern) -> text[]|bytea[]`
- `re2extractallgroupshorizontal(haystack, pattern) -> text[]|bytea[]`
- `re2regexpquotemeta(haystack) -> text|bytea`
- `re2splitbyregexp(pattern, haystack, max_substrings default 0) -> text[]|bytea[]`
- `re2replaceregexpone(haystack, pattern, replacement) -> text|bytea`
- `re2replaceregexpall(haystack, pattern, replacement) -> text|bytea`
- `re2countmatches(...)` and `re2countmatchescaseinsensitive(...)`

```sql
SELECT re2extractallgroupsvertical('a=1 b=2', '(\\w)=(\\d)');
SELECT re2regexpquotemeta('a+b?');
SELECT re2splitbyregexp('\\s+', 'one two three', 2);
```

### Multi-Pattern Matching

The `re2multimatch*` family accepts either multiple pattern arguments or a `VARIADIC` array:

```sql
SELECT re2multimatchany('error: timeout', 'timeout', 'denied');
SELECT re2multimatchanyindex('error: timeout', VARIADIC ARRAY['timeout', 'denied']);
SELECT re2multimatchallindices('error: timeout', 'error', 'timeout', 'panic');
```

### Index Support

Version `0.4.0` adds two complementary index paths:

```sql
-- Anchored constant patterns can use a normal btree prefix scan.
CREATE INDEX docs_body_btree ON docs (body);
SELECT * FROM docs WHERE re2match(body, '^order_2025');

-- The @~ operator can use the extension's GIN operator class.
CREATE INDEX docs_body_re2_gin ON docs USING gin (body gin_re2_ops);
SELECT * FROM docs WHERE body @~ 'timeout|denied';
```

The extension also provides selectivity estimation for RE2 predicates. Check `EXPLAIN` with representative data before choosing between btree, GIN, and a sequential scan.

### Matching Semantics

- To match ClickHouse behavior, `.` matches line breaks by default.
- Prefix the pattern with `(?-s)` if you want `.` not to cross line breaks.
- Replacement strings support `\\0` through `\\9` backreferences.

### Caveats

- Upstream requires the system `re2` library at build/install time.
- The `v0.4.x` binaries use SQL extension version `0.4`; after replacing an older binary, run `ALTER EXTENSION re2 UPDATE TO '0.4'` when an upgrade is pending.
- `v0.4.1` fixes a cache-related use-after-free and improves stable-pattern and multi-match performance; use it instead of `v0.4.0`.
- `re2splitbyregexp` uses `pattern, haystack[, max_substrings]`. Builds older than `0.3.0` used the reverse order.
- RE2 deliberately excludes features such as backreferences in patterns and look-around assertions; its bounded-time behavior differs from PostgreSQL's native regular-expression engine.
