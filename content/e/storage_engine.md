---
title: "storage_engine"
linkTitle: "storage_engine"
description: "colcompress and rowcompress Table Access Methods with vectorized execution"
weight: 2450
categories: ["OLAP"]
width: full
---

[**storage_engine**](https://github.com/saulojb/storage_engine) : colcompress and rowcompress Table Access Methods with vectorized execution


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2450** | {{< badge content="storage_engine" link="https://github.com/saulojb/storage_engine" >}} | {{< ext "storage_engine" >}} | `1.2.3` | {{< category "OLAP" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `engine` |

> [!Note] release 1.2.3; SQL v1.2.1


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `storage_engine` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.3` | {{< bg "18" "storage_engine_18" "green" >}} {{< bg "17" "storage_engine_17" "green" >}} {{< bg "16" "storage_engine_16" "green" >}} {{< bg "15" "storage_engine_15" "green" >}} {{< bg "14" "storage_engine_14" "green" >}} | `storage_engine_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.3` | {{< bg "18" "postgresql-18-storage-engine" "green" >}} {{< bg "17" "postgresql-17-storage-engine" "green" >}} {{< bg "16" "postgresql-16-storage-engine" "green" >}} {{< bg "15" "postgresql-15-storage-engine" "green" >}} {{< bg "14" "postgresql-14-storage-engine" "green" >}} | `postgresql-$v-storage-engine` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-storage-engine : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-storage-engine : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-storage-engine : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-storage-engine : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-storage-engine : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-storage-engine : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-storage-engine : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-storage-engine : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-storage-engine : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-storage-engine : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `storage_engine_18` | `1.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 182.5 KiB | [storage_engine_18-1.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/storage_engine_18-1.0.7-1PIGSTY.el8.x86_64.rpm) |
| `storage_engine_18` | `1.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 174.4 KiB | [storage_engine_18-1.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/storage_engine_18-1.0.7-1PIGSTY.el8.aarch64.rpm) |
| `storage_engine_18` | `1.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 154.4 KiB | [storage_engine_18-1.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/storage_engine_18-1.0.7-1PIGSTY.el9.x86_64.rpm) |
| `storage_engine_18` | `1.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 150.0 KiB | [storage_engine_18-1.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/storage_engine_18-1.0.7-1PIGSTY.el9.aarch64.rpm) |
| `storage_engine_18` | `1.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 156.7 KiB | [storage_engine_18-1.0.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/storage_engine_18-1.0.7-1PIGSTY.el10.x86_64.rpm) |
| `storage_engine_18` | `1.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 151.0 KiB | [storage_engine_18-1.0.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/storage_engine_18-1.0.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-storage-engine` | `1.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 415.0 KiB | [postgresql-18-storage-engine_1.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-18-storage-engine_1.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-storage-engine` | `1.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 403.5 KiB | [postgresql-18-storage-engine_1.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-18-storage-engine_1.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-storage-engine` | `1.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 416.3 KiB | [postgresql-18-storage-engine_1.0.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-18-storage-engine_1.0.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-storage-engine` | `1.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 405.7 KiB | [postgresql-18-storage-engine_1.0.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-18-storage-engine_1.0.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-storage-engine` | `1.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 440.4 KiB | [postgresql-18-storage-engine_1.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-18-storage-engine_1.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-storage-engine` | `1.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 434.0 KiB | [postgresql-18-storage-engine_1.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-18-storage-engine_1.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-storage-engine` | `1.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 423.2 KiB | [postgresql-18-storage-engine_1.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-18-storage-engine_1.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-storage-engine` | `1.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 419.0 KiB | [postgresql-18-storage-engine_1.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-18-storage-engine_1.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `storage_engine_17` | `1.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 182.0 KiB | [storage_engine_17-1.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/storage_engine_17-1.0.7-1PIGSTY.el8.x86_64.rpm) |
| `storage_engine_17` | `1.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 173.9 KiB | [storage_engine_17-1.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/storage_engine_17-1.0.7-1PIGSTY.el8.aarch64.rpm) |
| `storage_engine_17` | `1.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 153.7 KiB | [storage_engine_17-1.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/storage_engine_17-1.0.7-1PIGSTY.el9.x86_64.rpm) |
| `storage_engine_17` | `1.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 149.5 KiB | [storage_engine_17-1.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/storage_engine_17-1.0.7-1PIGSTY.el9.aarch64.rpm) |
| `storage_engine_17` | `1.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 156.1 KiB | [storage_engine_17-1.0.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/storage_engine_17-1.0.7-1PIGSTY.el10.x86_64.rpm) |
| `storage_engine_17` | `1.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 150.4 KiB | [storage_engine_17-1.0.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/storage_engine_17-1.0.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-storage-engine` | `1.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 413.6 KiB | [postgresql-17-storage-engine_1.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-17-storage-engine_1.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-storage-engine` | `1.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 402.4 KiB | [postgresql-17-storage-engine_1.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-17-storage-engine_1.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-storage-engine` | `1.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 414.9 KiB | [postgresql-17-storage-engine_1.0.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-17-storage-engine_1.0.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-storage-engine` | `1.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 404.5 KiB | [postgresql-17-storage-engine_1.0.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-17-storage-engine_1.0.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-storage-engine` | `1.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 503.8 KiB | [postgresql-17-storage-engine_1.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-17-storage-engine_1.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-storage-engine` | `1.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 496.1 KiB | [postgresql-17-storage-engine_1.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-17-storage-engine_1.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-storage-engine` | `1.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 420.9 KiB | [postgresql-17-storage-engine_1.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-17-storage-engine_1.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-storage-engine` | `1.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 418.0 KiB | [postgresql-17-storage-engine_1.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-17-storage-engine_1.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `storage_engine_16` | `1.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 181.9 KiB | [storage_engine_16-1.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/storage_engine_16-1.0.7-1PIGSTY.el8.x86_64.rpm) |
| `storage_engine_16` | `1.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 173.8 KiB | [storage_engine_16-1.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/storage_engine_16-1.0.7-1PIGSTY.el8.aarch64.rpm) |
| `storage_engine_16` | `1.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 153.9 KiB | [storage_engine_16-1.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/storage_engine_16-1.0.7-1PIGSTY.el9.x86_64.rpm) |
| `storage_engine_16` | `1.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 149.4 KiB | [storage_engine_16-1.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/storage_engine_16-1.0.7-1PIGSTY.el9.aarch64.rpm) |
| `storage_engine_16` | `1.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 155.7 KiB | [storage_engine_16-1.0.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/storage_engine_16-1.0.7-1PIGSTY.el10.x86_64.rpm) |
| `storage_engine_16` | `1.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 150.3 KiB | [storage_engine_16-1.0.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/storage_engine_16-1.0.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-storage-engine` | `1.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 412.9 KiB | [postgresql-16-storage-engine_1.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-16-storage-engine_1.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-storage-engine` | `1.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 402.1 KiB | [postgresql-16-storage-engine_1.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-16-storage-engine_1.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-storage-engine` | `1.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 414.4 KiB | [postgresql-16-storage-engine_1.0.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-16-storage-engine_1.0.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-storage-engine` | `1.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 403.9 KiB | [postgresql-16-storage-engine_1.0.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-16-storage-engine_1.0.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-storage-engine` | `1.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 499.2 KiB | [postgresql-16-storage-engine_1.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-16-storage-engine_1.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-storage-engine` | `1.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 491.9 KiB | [postgresql-16-storage-engine_1.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-16-storage-engine_1.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-storage-engine` | `1.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 420.0 KiB | [postgresql-16-storage-engine_1.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-16-storage-engine_1.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-storage-engine` | `1.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 417.6 KiB | [postgresql-16-storage-engine_1.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-16-storage-engine_1.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `storage_engine_15` | `1.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 185.4 KiB | [storage_engine_15-1.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/storage_engine_15-1.0.7-1PIGSTY.el8.x86_64.rpm) |
| `storage_engine_15` | `1.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 177.2 KiB | [storage_engine_15-1.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/storage_engine_15-1.0.7-1PIGSTY.el8.aarch64.rpm) |
| `storage_engine_15` | `1.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 173.4 KiB | [storage_engine_15-1.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/storage_engine_15-1.0.7-1PIGSTY.el9.x86_64.rpm) |
| `storage_engine_15` | `1.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 167.6 KiB | [storage_engine_15-1.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/storage_engine_15-1.0.7-1PIGSTY.el9.aarch64.rpm) |
| `storage_engine_15` | `1.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 175.2 KiB | [storage_engine_15-1.0.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/storage_engine_15-1.0.7-1PIGSTY.el10.x86_64.rpm) |
| `storage_engine_15` | `1.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 168.4 KiB | [storage_engine_15-1.0.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/storage_engine_15-1.0.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-storage-engine` | `1.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 419.0 KiB | [postgresql-15-storage-engine_1.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-15-storage-engine_1.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-storage-engine` | `1.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 406.2 KiB | [postgresql-15-storage-engine_1.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-15-storage-engine_1.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-storage-engine` | `1.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 420.6 KiB | [postgresql-15-storage-engine_1.0.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-15-storage-engine_1.0.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-storage-engine` | `1.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 408.3 KiB | [postgresql-15-storage-engine_1.0.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-15-storage-engine_1.0.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-storage-engine` | `1.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 519.9 KiB | [postgresql-15-storage-engine_1.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-15-storage-engine_1.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-storage-engine` | `1.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 512.9 KiB | [postgresql-15-storage-engine_1.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-15-storage-engine_1.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-storage-engine` | `1.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 440.0 KiB | [postgresql-15-storage-engine_1.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-15-storage-engine_1.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-storage-engine` | `1.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 434.1 KiB | [postgresql-15-storage-engine_1.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-15-storage-engine_1.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `storage_engine_14` | `1.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 185.8 KiB | [storage_engine_14-1.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/storage_engine_14-1.0.7-1PIGSTY.el8.x86_64.rpm) |
| `storage_engine_14` | `1.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 177.5 KiB | [storage_engine_14-1.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/storage_engine_14-1.0.7-1PIGSTY.el8.aarch64.rpm) |
| `storage_engine_14` | `1.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 173.8 KiB | [storage_engine_14-1.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/storage_engine_14-1.0.7-1PIGSTY.el9.x86_64.rpm) |
| `storage_engine_14` | `1.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 167.5 KiB | [storage_engine_14-1.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/storage_engine_14-1.0.7-1PIGSTY.el9.aarch64.rpm) |
| `storage_engine_14` | `1.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 175.5 KiB | [storage_engine_14-1.0.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/storage_engine_14-1.0.7-1PIGSTY.el10.x86_64.rpm) |
| `storage_engine_14` | `1.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 168.9 KiB | [storage_engine_14-1.0.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/storage_engine_14-1.0.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-storage-engine` | `1.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 420.5 KiB | [postgresql-14-storage-engine_1.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-14-storage-engine_1.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-storage-engine` | `1.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 407.1 KiB | [postgresql-14-storage-engine_1.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-14-storage-engine_1.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-storage-engine` | `1.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 421.8 KiB | [postgresql-14-storage-engine_1.0.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-14-storage-engine_1.0.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-storage-engine` | `1.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 409.9 KiB | [postgresql-14-storage-engine_1.0.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-14-storage-engine_1.0.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-storage-engine` | `1.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 521.9 KiB | [postgresql-14-storage-engine_1.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-14-storage-engine_1.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-storage-engine` | `1.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 514.2 KiB | [postgresql-14-storage-engine_1.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-14-storage-engine_1.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-storage-engine` | `1.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 440.9 KiB | [postgresql-14-storage-engine_1.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-14-storage-engine_1.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-storage-engine` | `1.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 435.7 KiB | [postgresql-14-storage-engine_1.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-14-storage-engine_1.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/saulojb/storage_engine" title="Repository" icon="github" subtitle="github.com/saulojb/storage_engine" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="storage_engine-1.2.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg storage_engine;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install storage_engine;		# install via package name, for the active PG version

pig install storage_engine -v 18;   # install for PG 18
pig install storage_engine -v 17;   # install for PG 17
pig install storage_engine -v 16;   # install for PG 16
pig install storage_engine -v 15;   # install for PG 15
pig install storage_engine -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'storage_engine';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION storage_engine;
```

## Usage

Sources: [README](https://github.com/saulojb/storage_engine/blob/main/README.md), [release 1.0.7](https://github.com/saulojb/storage_engine/releases/tag/v1.0.7), [META.json](https://github.com/saulojb/storage_engine/blob/main/META.json)

`storage_engine` provides two PostgreSQL table access methods in the `engine` schema:

- `colcompress` for column-oriented compressed storage with vectorized execution, min/max pruning, and parallel scans.
- `rowcompress` for row-batch compression with parallel scans.

```sql
CREATE EXTENSION storage_engine;
```

### Quick Start

Create tables using either access method:

```sql
CREATE TABLE events (
  ts timestamptz NOT NULL,
  user_id bigint,
  event_type text,
  value float8
) USING colcompress;

CREATE TABLE logs (
  id bigserial,
  logged_at timestamptz NOT NULL,
  message text
) USING rowcompress;
```

### Main Tuning Knobs

Session-level GUCs documented upstream include:

- `storage_engine.enable_parallel_execution`
- `storage_engine.enable_vectorization`
- `storage_engine.enable_column_cache`
- `storage_engine.enable_columnar_index_scan`
- `storage_engine.enable_dml`
- `storage_engine.stripe_row_limit`
- `storage_engine.chunk_group_row_limit`
- `storage_engine.compression_level`

The README says these GUCs become visible once the library is loaded; add `storage_engine` to `shared_preload_libraries` if you want them available immediately in every session.

### Useful Management Functions

For `colcompress` tables:

```sql
SELECT engine.alter_colcompress_table_set(
  'events'::regclass,
  orderby => 'ts ASC, user_id ASC',
  compression => 'zstd',
  compression_level => 9
);

SELECT engine.colcompress_merge('events');
SELECT engine.colcompress_repack('events');
```

For `rowcompress` tables:

```sql
SELECT engine.alter_rowcompress_table_set(
  'logs'::regclass,
  batch_size => 10000,
  compression => 'zstd',
  compression_level => 5
);

SELECT engine.rowcompress_repack('logs');
```

### When to Use Which AM

- Use `colcompress` for analytical scans, aggregates, and range predicates where projection, vectorization, and stripe/chunk pruning pay off.
- Use `rowcompress` for append-heavy logs or wide rows that are usually fetched together, where compression matters more than column projection.
- For point lookups on `colcompress`, upstream recommends enabling `storage_engine.enable_columnar_index_scan` or per-table `index_scan`.

### Caveats

- `colcompress` and `rowcompress` do not support foreign keys or `AFTER ROW` triggers.
- `VACUUM FULL` and `CREATE UNLOGGED TABLE ... USING colcompress` are not supported; upstream recommends the extension's repack functions instead.
- On `colcompress`, combining `orderby` with B-tree indexes can disable the sort-on-write path; run `engine.colcompress_merge()` after loading data when global ordering matters.
