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
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `storage_engine_18` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 202.5 KiB | [storage_engine_18-1.2.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/storage_engine_18-1.2.3-2PIGSTY.el8.x86_64.rpm) |
| `storage_engine_18` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 193.7 KiB | [storage_engine_18-1.2.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/storage_engine_18-1.2.3-2PIGSTY.el8.aarch64.rpm) |
| `storage_engine_18` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 172.7 KiB | [storage_engine_18-1.2.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/storage_engine_18-1.2.3-2PIGSTY.el9.x86_64.rpm) |
| `storage_engine_18` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 168.6 KiB | [storage_engine_18-1.2.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/storage_engine_18-1.2.3-2PIGSTY.el9.aarch64.rpm) |
| `storage_engine_18` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 175.9 KiB | [storage_engine_18-1.2.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/storage_engine_18-1.2.3-2PIGSTY.el10.x86_64.rpm) |
| `storage_engine_18` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 170.2 KiB | [storage_engine_18-1.2.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/storage_engine_18-1.2.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-storage-engine` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 444.3 KiB | [postgresql-18-storage-engine_1.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-18-storage-engine_1.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-storage-engine` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 432.4 KiB | [postgresql-18-storage-engine_1.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-18-storage-engine_1.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-storage-engine` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 445.6 KiB | [postgresql-18-storage-engine_1.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-18-storage-engine_1.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-storage-engine` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 434.5 KiB | [postgresql-18-storage-engine_1.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-18-storage-engine_1.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-storage-engine` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 485.6 KiB | [postgresql-18-storage-engine_1.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-18-storage-engine_1.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-storage-engine` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 465.0 KiB | [postgresql-18-storage-engine_1.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-18-storage-engine_1.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-storage-engine` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 467.3 KiB | [postgresql-18-storage-engine_1.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-18-storage-engine_1.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-storage-engine` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 449.4 KiB | [postgresql-18-storage-engine_1.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-18-storage-engine_1.2.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-storage-engine` | `1.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 450.9 KiB | [postgresql-18-storage-engine_1.2.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-18-storage-engine_1.2.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-storage-engine` | `1.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 446.5 KiB | [postgresql-18-storage-engine_1.2.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-18-storage-engine_1.2.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `storage_engine_17` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 202.2 KiB | [storage_engine_17-1.2.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/storage_engine_17-1.2.3-2PIGSTY.el8.x86_64.rpm) |
| `storage_engine_17` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 193.0 KiB | [storage_engine_17-1.2.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/storage_engine_17-1.2.3-2PIGSTY.el8.aarch64.rpm) |
| `storage_engine_17` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 172.5 KiB | [storage_engine_17-1.2.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/storage_engine_17-1.2.3-2PIGSTY.el9.x86_64.rpm) |
| `storage_engine_17` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 168.2 KiB | [storage_engine_17-1.2.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/storage_engine_17-1.2.3-2PIGSTY.el9.aarch64.rpm) |
| `storage_engine_17` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 175.5 KiB | [storage_engine_17-1.2.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/storage_engine_17-1.2.3-2PIGSTY.el10.x86_64.rpm) |
| `storage_engine_17` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 169.6 KiB | [storage_engine_17-1.2.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/storage_engine_17-1.2.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-storage-engine` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 443.2 KiB | [postgresql-17-storage-engine_1.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-17-storage-engine_1.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-storage-engine` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 431.4 KiB | [postgresql-17-storage-engine_1.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-17-storage-engine_1.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-storage-engine` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 444.3 KiB | [postgresql-17-storage-engine_1.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-17-storage-engine_1.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-storage-engine` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 433.3 KiB | [postgresql-17-storage-engine_1.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-17-storage-engine_1.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-storage-engine` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 549.9 KiB | [postgresql-17-storage-engine_1.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-17-storage-engine_1.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-storage-engine` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 528.8 KiB | [postgresql-17-storage-engine_1.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-17-storage-engine_1.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-storage-engine` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 465.7 KiB | [postgresql-17-storage-engine_1.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-17-storage-engine_1.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-storage-engine` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 448.0 KiB | [postgresql-17-storage-engine_1.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-17-storage-engine_1.2.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-storage-engine` | `1.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 449.3 KiB | [postgresql-17-storage-engine_1.2.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-17-storage-engine_1.2.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-storage-engine` | `1.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 445.2 KiB | [postgresql-17-storage-engine_1.2.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-17-storage-engine_1.2.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `storage_engine_16` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 201.9 KiB | [storage_engine_16-1.2.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/storage_engine_16-1.2.3-2PIGSTY.el8.x86_64.rpm) |
| `storage_engine_16` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 193.1 KiB | [storage_engine_16-1.2.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/storage_engine_16-1.2.3-2PIGSTY.el8.aarch64.rpm) |
| `storage_engine_16` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 172.4 KiB | [storage_engine_16-1.2.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/storage_engine_16-1.2.3-2PIGSTY.el9.x86_64.rpm) |
| `storage_engine_16` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 167.9 KiB | [storage_engine_16-1.2.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/storage_engine_16-1.2.3-2PIGSTY.el9.aarch64.rpm) |
| `storage_engine_16` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 175.3 KiB | [storage_engine_16-1.2.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/storage_engine_16-1.2.3-2PIGSTY.el10.x86_64.rpm) |
| `storage_engine_16` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 169.6 KiB | [storage_engine_16-1.2.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/storage_engine_16-1.2.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-storage-engine` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 442.7 KiB | [postgresql-16-storage-engine_1.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-16-storage-engine_1.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-storage-engine` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 430.8 KiB | [postgresql-16-storage-engine_1.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-16-storage-engine_1.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-storage-engine` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 443.5 KiB | [postgresql-16-storage-engine_1.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-16-storage-engine_1.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-storage-engine` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 432.8 KiB | [postgresql-16-storage-engine_1.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-16-storage-engine_1.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-storage-engine` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 547.0 KiB | [postgresql-16-storage-engine_1.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-16-storage-engine_1.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-storage-engine` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 525.3 KiB | [postgresql-16-storage-engine_1.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-16-storage-engine_1.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-storage-engine` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 465.4 KiB | [postgresql-16-storage-engine_1.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-16-storage-engine_1.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-storage-engine` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 447.3 KiB | [postgresql-16-storage-engine_1.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-16-storage-engine_1.2.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-storage-engine` | `1.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 448.5 KiB | [postgresql-16-storage-engine_1.2.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-16-storage-engine_1.2.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-storage-engine` | `1.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 444.3 KiB | [postgresql-16-storage-engine_1.2.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-16-storage-engine_1.2.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `storage_engine_15` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 205.6 KiB | [storage_engine_15-1.2.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/storage_engine_15-1.2.3-2PIGSTY.el8.x86_64.rpm) |
| `storage_engine_15` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 196.7 KiB | [storage_engine_15-1.2.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/storage_engine_15-1.2.3-2PIGSTY.el8.aarch64.rpm) |
| `storage_engine_15` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 191.8 KiB | [storage_engine_15-1.2.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/storage_engine_15-1.2.3-2PIGSTY.el9.x86_64.rpm) |
| `storage_engine_15` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 185.9 KiB | [storage_engine_15-1.2.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/storage_engine_15-1.2.3-2PIGSTY.el9.aarch64.rpm) |
| `storage_engine_15` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 194.3 KiB | [storage_engine_15-1.2.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/storage_engine_15-1.2.3-2PIGSTY.el10.x86_64.rpm) |
| `storage_engine_15` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 187.6 KiB | [storage_engine_15-1.2.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/storage_engine_15-1.2.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-storage-engine` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 450.1 KiB | [postgresql-15-storage-engine_1.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-15-storage-engine_1.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-storage-engine` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 435.8 KiB | [postgresql-15-storage-engine_1.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-15-storage-engine_1.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-storage-engine` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 451.8 KiB | [postgresql-15-storage-engine_1.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-15-storage-engine_1.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-storage-engine` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 438.4 KiB | [postgresql-15-storage-engine_1.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-15-storage-engine_1.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-storage-engine` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 554.9 KiB | [postgresql-15-storage-engine_1.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-15-storage-engine_1.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-storage-engine` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 547.6 KiB | [postgresql-15-storage-engine_1.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-15-storage-engine_1.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-storage-engine` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 472.2 KiB | [postgresql-15-storage-engine_1.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-15-storage-engine_1.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-storage-engine` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 465.3 KiB | [postgresql-15-storage-engine_1.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-15-storage-engine_1.2.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-storage-engine` | `1.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 469.0 KiB | [postgresql-15-storage-engine_1.2.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-15-storage-engine_1.2.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-storage-engine` | `1.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 463.8 KiB | [postgresql-15-storage-engine_1.2.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-15-storage-engine_1.2.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `storage_engine_14` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 205.7 KiB | [storage_engine_14-1.2.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/storage_engine_14-1.2.3-2PIGSTY.el8.x86_64.rpm) |
| `storage_engine_14` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 196.8 KiB | [storage_engine_14-1.2.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/storage_engine_14-1.2.3-2PIGSTY.el8.aarch64.rpm) |
| `storage_engine_14` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 192.8 KiB | [storage_engine_14-1.2.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/storage_engine_14-1.2.3-2PIGSTY.el9.x86_64.rpm) |
| `storage_engine_14` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 186.4 KiB | [storage_engine_14-1.2.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/storage_engine_14-1.2.3-2PIGSTY.el9.aarch64.rpm) |
| `storage_engine_14` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 194.9 KiB | [storage_engine_14-1.2.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/storage_engine_14-1.2.3-2PIGSTY.el10.x86_64.rpm) |
| `storage_engine_14` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 188.2 KiB | [storage_engine_14-1.2.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/storage_engine_14-1.2.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-storage-engine` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 450.8 KiB | [postgresql-14-storage-engine_1.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-14-storage-engine_1.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-storage-engine` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 436.6 KiB | [postgresql-14-storage-engine_1.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-14-storage-engine_1.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-storage-engine` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 452.2 KiB | [postgresql-14-storage-engine_1.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-14-storage-engine_1.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-storage-engine` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 439.6 KiB | [postgresql-14-storage-engine_1.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-14-storage-engine_1.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-storage-engine` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 556.8 KiB | [postgresql-14-storage-engine_1.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-14-storage-engine_1.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-storage-engine` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 549.5 KiB | [postgresql-14-storage-engine_1.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-14-storage-engine_1.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-storage-engine` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 472.7 KiB | [postgresql-14-storage-engine_1.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-14-storage-engine_1.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-storage-engine` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 467.1 KiB | [postgresql-14-storage-engine_1.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-14-storage-engine_1.2.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-storage-engine` | `1.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 470.4 KiB | [postgresql-14-storage-engine_1.2.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-14-storage-engine_1.2.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-storage-engine` | `1.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 464.7 KiB | [postgresql-14-storage-engine_1.2.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-14-storage-engine_1.2.3-1PIGSTY~resolute_arm64.deb) |

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
