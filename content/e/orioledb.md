---
title: "orioledb"
linkTitle: "orioledb"
description: "OrioleDB, the next generation transactional engine"
weight: 2910
categories: ["FEAT"]
width: full
---

[**orioledb**](https://github.com/orioledb/orioledb) : OrioleDB, the next generation transactional engine


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2910** | {{< badge content="orioledb" link="https://github.com/orioledb/orioledb" >}} | {{< ext "orioledb" >}} | `1.8` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_mooncake" >}} {{< ext "citus_columnar" >}} {{< ext "pg_analytics" >}} {{< ext "pg_duckdb" >}} {{< ext "timescaledb" >}} {{< ext "citus" >}} {{< ext "pg_strom" >}} |

> [!Note] patched kernel; beta16 for patchset 18.1/17.20/16.47


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.8` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `orioledb` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.8` | {{< bg "18" "orioledb-18" "green" >}} {{< bg "17" "orioledb-17" "green" >}} {{< bg "16" "orioledb-16" "green" >}} {{< bg "15" "orioledb-15" "red" >}} {{< bg "14" "orioledb-14" "red" >}} | `orioledb-$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.8` | {{< bg "18" "orioledb-18" "green" >}} {{< bg "17" "orioledb-17" "green" >}} {{< bg "16" "orioledb-16" "green" >}} {{< bg "15" "orioledb-15" "red" >}} {{< bg "14" "orioledb-14" "red" >}} | `orioledb-$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |
| {{< os "u26.x86_64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "PIGSTY 1.8" "orioledb-18 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-17 : FORK 1" >}}      |      {{< bg "PIGSTY 1.8" "orioledb-16 : FORK 1" >}}      |      {{< bg "MISS" "orioledb-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb-14 : FORK 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orioledb-18` | `1.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 MiB | [orioledb-18-1.8-beta16PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/orioledb-18-1.8-beta16PIGSTY.el8.x86_64.rpm) |
| `orioledb-18` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.1 MiB | [orioledb-18-1.8-beta16PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/orioledb-18-1.8-beta16PIGSTY.el8.aarch64.rpm) |
| `orioledb-18` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.1 MiB | [orioledb-18-1.8-beta16PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/orioledb-18-1.8-beta16PIGSTY.el9.x86_64.rpm) |
| `orioledb-18` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.9 MiB | [orioledb-18-1.8-beta16PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/orioledb-18-1.8-beta16PIGSTY.el9.aarch64.rpm) |
| `orioledb-18` | `1.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.3 MiB | [orioledb-18-1.8-beta16PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/orioledb-18-1.8-beta16PIGSTY.el10.x86_64.rpm) |
| `orioledb-18` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.1 MiB | [orioledb-18-1.8-beta16PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/orioledb-18-1.8-beta16PIGSTY.el10.aarch64.rpm) |
| `orioledb-18` | `1.8` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.3 MiB | [orioledb-18_1.8-0.beta16PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/orioledb-18/orioledb-18_1.8-0.beta16PIGSTY~bookworm_amd64.deb) |
| `orioledb-18` | `1.8` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.7 MiB | [orioledb-18_1.8-0.beta16PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/orioledb-18/orioledb-18_1.8-0.beta16PIGSTY~bookworm_arm64.deb) |
| `orioledb-18` | `1.8` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.3 MiB | [orioledb-18_1.8-0.beta16PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/orioledb-18/orioledb-18_1.8-0.beta16PIGSTY~trixie_amd64.deb) |
| `orioledb-18` | `1.8` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.8 MiB | [orioledb-18_1.8-0.beta16PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/orioledb-18/orioledb-18_1.8-0.beta16PIGSTY~trixie_arm64.deb) |
| `orioledb-18` | `1.8` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.7 MiB | [orioledb-18_1.8-0.beta16PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/orioledb-18/orioledb-18_1.8-0.beta16PIGSTY~jammy_amd64.deb) |
| `orioledb-18` | `1.8` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.5 MiB | [orioledb-18_1.8-0.beta16PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/orioledb-18/orioledb-18_1.8-0.beta16PIGSTY~jammy_arm64.deb) |
| `orioledb-18` | `1.8` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.5 MiB | [orioledb-18_1.8-0.beta16PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/orioledb-18/orioledb-18_1.8-0.beta16PIGSTY~noble_amd64.deb) |
| `orioledb-18` | `1.8` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.4 MiB | [orioledb-18_1.8-0.beta16PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/orioledb-18/orioledb-18_1.8-0.beta16PIGSTY~noble_arm64.deb) |
| `orioledb-18` | `1.8` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.6 MiB | [orioledb-18_1.8-0.beta16PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/orioledb-18/orioledb-18_1.8-0.beta16PIGSTY~resolute_amd64.deb) |
| `orioledb-18` | `1.8` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.3 MiB | [orioledb-18_1.8-0.beta16PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/orioledb-18/orioledb-18_1.8-0.beta16PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orioledb-17` | `1.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.0 MiB | [orioledb-17-1.8-beta16PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/orioledb-17-1.8-beta16PIGSTY.el8.x86_64.rpm) |
| `orioledb-17` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.6 MiB | [orioledb-17-1.8-beta16PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/orioledb-17-1.8-beta16PIGSTY.el8.aarch64.rpm) |
| `orioledb-17` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.8 MiB | [orioledb-17-1.8-beta16PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/orioledb-17-1.8-beta16PIGSTY.el9.x86_64.rpm) |
| `orioledb-17` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.6 MiB | [orioledb-17-1.8-beta16PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/orioledb-17-1.8-beta16PIGSTY.el9.aarch64.rpm) |
| `orioledb-17` | `1.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.9 MiB | [orioledb-17-1.8-beta16PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/orioledb-17-1.8-beta16PIGSTY.el10.x86_64.rpm) |
| `orioledb-17` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 MiB | [orioledb-17-1.8-beta16PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/orioledb-17-1.8-beta16PIGSTY.el10.aarch64.rpm) |
| `orioledb-17` | `1.8` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.9 MiB | [orioledb-17_1.8-0.beta16PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/orioledb-17/orioledb-17_1.8-0.beta16PIGSTY~bookworm_amd64.deb) |
| `orioledb-17` | `1.8` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.4 MiB | [orioledb-17_1.8-0.beta16PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/orioledb-17/orioledb-17_1.8-0.beta16PIGSTY~bookworm_arm64.deb) |
| `orioledb-17` | `1.8` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.0 MiB | [orioledb-17_1.8-0.beta16PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/orioledb-17/orioledb-17_1.8-0.beta16PIGSTY~trixie_amd64.deb) |
| `orioledb-17` | `1.8` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.5 MiB | [orioledb-17_1.8-0.beta16PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/orioledb-17/orioledb-17_1.8-0.beta16PIGSTY~trixie_arm64.deb) |
| `orioledb-17` | `1.8` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.3 MiB | [orioledb-17_1.8-0.beta16PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/orioledb-17/orioledb-17_1.8-0.beta16PIGSTY~jammy_amd64.deb) |
| `orioledb-17` | `1.8` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.1 MiB | [orioledb-17_1.8-0.beta16PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/orioledb-17/orioledb-17_1.8-0.beta16PIGSTY~jammy_arm64.deb) |
| `orioledb-17` | `1.8` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.1 MiB | [orioledb-17_1.8-0.beta16PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/orioledb-17/orioledb-17_1.8-0.beta16PIGSTY~noble_amd64.deb) |
| `orioledb-17` | `1.8` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.0 MiB | [orioledb-17_1.8-0.beta16PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/orioledb-17/orioledb-17_1.8-0.beta16PIGSTY~noble_arm64.deb) |
| `orioledb-17` | `1.8` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.2 MiB | [orioledb-17_1.8-0.beta16PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/orioledb-17/orioledb-17_1.8-0.beta16PIGSTY~resolute_amd64.deb) |
| `orioledb-17` | `1.8` | [u26.aarch64](/os/u26.aarch64) | pigsty | 10.9 MiB | [orioledb-17_1.8-0.beta16PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/orioledb-17/orioledb-17_1.8-0.beta16PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orioledb-16` | `1.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.3 MiB | [orioledb-16-1.8-beta16PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/orioledb-16-1.8-beta16PIGSTY.el8.x86_64.rpm) |
| `orioledb-16` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.9 MiB | [orioledb-16-1.8-beta16PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/orioledb-16-1.8-beta16PIGSTY.el8.aarch64.rpm) |
| `orioledb-16` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.3 MiB | [orioledb-16-1.8-beta16PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/orioledb-16-1.8-beta16PIGSTY.el9.x86_64.rpm) |
| `orioledb-16` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.1 MiB | [orioledb-16-1.8-beta16PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/orioledb-16-1.8-beta16PIGSTY.el9.aarch64.rpm) |
| `orioledb-16` | `1.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.4 MiB | [orioledb-16-1.8-beta16PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/orioledb-16-1.8-beta16PIGSTY.el10.x86_64.rpm) |
| `orioledb-16` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.2 MiB | [orioledb-16-1.8-beta16PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/orioledb-16-1.8-beta16PIGSTY.el10.aarch64.rpm) |
| `orioledb-16` | `1.8` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.4 MiB | [orioledb-16_1.8-0.beta16PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/orioledb-16/orioledb-16_1.8-0.beta16PIGSTY~bookworm_amd64.deb) |
| `orioledb-16` | `1.8` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.0 MiB | [orioledb-16_1.8-0.beta16PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/orioledb-16/orioledb-16_1.8-0.beta16PIGSTY~bookworm_arm64.deb) |
| `orioledb-16` | `1.8` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.5 MiB | [orioledb-16_1.8-0.beta16PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/orioledb-16/orioledb-16_1.8-0.beta16PIGSTY~trixie_amd64.deb) |
| `orioledb-16` | `1.8` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.0 MiB | [orioledb-16_1.8-0.beta16PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/orioledb-16/orioledb-16_1.8-0.beta16PIGSTY~trixie_arm64.deb) |
| `orioledb-16` | `1.8` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.8 MiB | [orioledb-16_1.8-0.beta16PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/orioledb-16/orioledb-16_1.8-0.beta16PIGSTY~jammy_amd64.deb) |
| `orioledb-16` | `1.8` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.6 MiB | [orioledb-16_1.8-0.beta16PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/orioledb-16/orioledb-16_1.8-0.beta16PIGSTY~jammy_arm64.deb) |
| `orioledb-16` | `1.8` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.7 MiB | [orioledb-16_1.8-0.beta16PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/orioledb-16/orioledb-16_1.8-0.beta16PIGSTY~noble_amd64.deb) |
| `orioledb-16` | `1.8` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.5 MiB | [orioledb-16_1.8-0.beta16PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/orioledb-16/orioledb-16_1.8-0.beta16PIGSTY~noble_arm64.deb) |
| `orioledb-16` | `1.8` | [u26.x86_64](/os/u26.x86_64) | pigsty | 10.7 MiB | [orioledb-16_1.8-0.beta16PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/orioledb-16/orioledb-16_1.8-0.beta16PIGSTY~resolute_amd64.deb) |
| `orioledb-16` | `1.8` | [u26.aarch64](/os/u26.aarch64) | pigsty | 10.5 MiB | [orioledb-16_1.8-0.beta16PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/orioledb-16/orioledb-16_1.8-0.beta16PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/orioledb/orioledb" title="Repository" icon="github" subtitle="github.com/orioledb/orioledb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="orioledb-beta16.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg orioledb;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install orioledb;		# install via package name, for the active PG version

pig install orioledb -v 18;   # install for PG 18
pig install orioledb -v 17;   # install for PG 17
pig install orioledb -v 16;   # install for PG 16

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'orioledb';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION orioledb;
```




## Usage

Sources: [README](https://github.com/orioledb/orioledb), [beta16 release](https://github.com/orioledb/orioledb/releases/tag/beta16), [patched PostgreSQL tree](https://github.com/orioledb/postgres)

OrioleDB is a new storage engine for PostgreSQL that provides modern approaches to database capacity, capabilities, and performance. It uses undo log-based MVCC, copy-on-write checkpoints, and row-level WAL to eliminate bloat and the need for VACUUM.

### Configuration

Add to `postgresql.conf` (requires restart):

```ini
shared_preload_libraries = 'orioledb.so'
```

Then enable the extension:

```sql
CREATE EXTENSION orioledb;
```

### Creating Tables

Use the `USING orioledb` clause to create tables with the OrioleDB storage engine:

```sql
CREATE TABLE my_table (
    id serial PRIMARY KEY,
    name text,
    value numeric
) USING orioledb;
```

All standard PostgreSQL operations work on OrioleDB tables:

```sql
INSERT INTO my_table (name, value) VALUES ('test', 42);
SELECT * FROM my_table WHERE id = 1;
UPDATE my_table SET value = 100 WHERE id = 1;
DELETE FROM my_table WHERE id = 1;
```

### Collation Requirements

OrioleDB tables support only **ICU**, **C**, and **POSIX** collations. To avoid specifying COLLATE on every text field, create the database with an appropriate default:

```sql
CREATE DATABASE mydb LOCALE 'C' TEMPLATE template0;
-- OR
CREATE DATABASE mydb LOCALE_PROVIDER icu ICU_LOCALE 'en' TEMPLATE template0;
```

### Key Benefits

- **No bloat**: Undo log-based MVCC means old tuple versions do not bloat main storage
- **No VACUUM needed**: Page-merging and undo log reclaim space automatically
- **No wraparound problem**: Native 64-bit transaction identifiers
- **Lock-less page reading**: In-memory pages linked directly to storage pages
- **Row-level WAL**: Compact write-ahead logging suitable for parallel apply

### Limitations

- Public beta status -- recommended for testing, not production
- Requires a patched PostgreSQL build from [orioledb/postgres](https://github.com/orioledb/postgres)
- Only ICU, C, and POSIX collations are supported

### Version Notes

OrioleDB 1.8-beta16 bumps the extension SQL version to `1.8`, bases patched PostgreSQL builds on 16.13, 17.9, and 18.4, and adds PostgreSQL 18 support. New user-facing surfaces include `orioledb.serializable` for SERIALIZABLE support and `verify_orioledb(regclass, boolean)` for `pg_amcheck` integration. The release also includes recovery, replication, index-scan, vacuum, and DDL correctness fixes.
