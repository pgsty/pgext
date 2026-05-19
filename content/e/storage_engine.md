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
| **2450** | {{< badge content="storage_engine" link="https://github.com/saulojb/storage_engine" >}} | {{< ext "storage_engine" >}} | `2.3.0` | {{< category "OLAP" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `engine` |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `storage_engine` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.3.0` | {{< bg "18" "storage_engine_18" "green" >}} {{< bg "17" "storage_engine_17" "green" >}} {{< bg "16" "storage_engine_16" "green" >}} {{< bg "15" "storage_engine_15" "green" >}} {{< bg "14" "storage_engine_14" "red" >}} | `storage_engine_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.3.0` | {{< bg "18" "postgresql-18-storage-engine" "green" >}} {{< bg "17" "postgresql-17-storage-engine" "green" >}} {{< bg "16" "postgresql-16-storage-engine" "green" >}} {{< bg "15" "postgresql-15-storage-engine" "green" >}} {{< bg "14" "postgresql-14-storage-engine" "red" >}} | `postgresql-$v-storage-engine` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "storage_engine_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "storage_engine_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-storage-engine : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.4" "postgresql-14-storage-engine : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `storage_engine_18` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 276.1 KiB | [storage_engine_18-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/storage_engine_18-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `storage_engine_18` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 264.6 KiB | [storage_engine_18-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/storage_engine_18-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `storage_engine_18` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 243.7 KiB | [storage_engine_18-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/storage_engine_18-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `storage_engine_18` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 237.7 KiB | [storage_engine_18-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/storage_engine_18-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `storage_engine_18` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 247.9 KiB | [storage_engine_18-2.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/storage_engine_18-2.3.0-1PIGSTY.el10.x86_64.rpm) |
| `storage_engine_18` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 239.7 KiB | [storage_engine_18-2.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/storage_engine_18-2.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-storage-engine` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 592.2 KiB | [postgresql-18-storage-engine_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-18-storage-engine_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-storage-engine` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 574.8 KiB | [postgresql-18-storage-engine_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-18-storage-engine_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-storage-engine` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 592.3 KiB | [postgresql-18-storage-engine_2.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-18-storage-engine_2.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-storage-engine` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 576.6 KiB | [postgresql-18-storage-engine_2.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-18-storage-engine_2.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-storage-engine` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 630.7 KiB | [postgresql-18-storage-engine_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-18-storage-engine_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-storage-engine` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 621.4 KiB | [postgresql-18-storage-engine_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-18-storage-engine_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-storage-engine` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 606.1 KiB | [postgresql-18-storage-engine_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-18-storage-engine_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-storage-engine` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 600.1 KiB | [postgresql-18-storage-engine_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-18-storage-engine_2.3.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-storage-engine` | `2.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 603.1 KiB | [postgresql-18-storage-engine_2.3.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-18-storage-engine_2.3.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-storage-engine` | `2.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 596.6 KiB | [postgresql-18-storage-engine_2.3.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-18-storage-engine_2.3.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `storage_engine_17` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 275.5 KiB | [storage_engine_17-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/storage_engine_17-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `storage_engine_17` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 264.0 KiB | [storage_engine_17-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/storage_engine_17-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `storage_engine_17` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 243.4 KiB | [storage_engine_17-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/storage_engine_17-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `storage_engine_17` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 237.2 KiB | [storage_engine_17-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/storage_engine_17-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `storage_engine_17` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 247.0 KiB | [storage_engine_17-2.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/storage_engine_17-2.3.0-1PIGSTY.el10.x86_64.rpm) |
| `storage_engine_17` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 239.3 KiB | [storage_engine_17-2.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/storage_engine_17-2.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-storage-engine` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 589.2 KiB | [postgresql-17-storage-engine_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-17-storage-engine_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-storage-engine` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 574.7 KiB | [postgresql-17-storage-engine_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-17-storage-engine_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-storage-engine` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 591.0 KiB | [postgresql-17-storage-engine_2.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-17-storage-engine_2.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-storage-engine` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 575.3 KiB | [postgresql-17-storage-engine_2.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-17-storage-engine_2.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-storage-engine` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 704.4 KiB | [postgresql-17-storage-engine_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-17-storage-engine_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-storage-engine` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 695.0 KiB | [postgresql-17-storage-engine_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-17-storage-engine_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-storage-engine` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 604.9 KiB | [postgresql-17-storage-engine_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-17-storage-engine_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-storage-engine` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 599.2 KiB | [postgresql-17-storage-engine_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-17-storage-engine_2.3.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-storage-engine` | `2.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 601.8 KiB | [postgresql-17-storage-engine_2.3.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-17-storage-engine_2.3.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-storage-engine` | `2.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 595.3 KiB | [postgresql-17-storage-engine_2.3.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-17-storage-engine_2.3.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `storage_engine_16` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 275.6 KiB | [storage_engine_16-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/storage_engine_16-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `storage_engine_16` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 264.2 KiB | [storage_engine_16-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/storage_engine_16-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `storage_engine_16` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 243.2 KiB | [storage_engine_16-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/storage_engine_16-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `storage_engine_16` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 237.5 KiB | [storage_engine_16-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/storage_engine_16-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `storage_engine_16` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 247.0 KiB | [storage_engine_16-2.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/storage_engine_16-2.3.0-1PIGSTY.el10.x86_64.rpm) |
| `storage_engine_16` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 239.3 KiB | [storage_engine_16-2.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/storage_engine_16-2.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-storage-engine` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 589.7 KiB | [postgresql-16-storage-engine_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-16-storage-engine_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-storage-engine` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 573.7 KiB | [postgresql-16-storage-engine_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-16-storage-engine_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-storage-engine` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 590.4 KiB | [postgresql-16-storage-engine_2.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-16-storage-engine_2.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-storage-engine` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 575.4 KiB | [postgresql-16-storage-engine_2.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-16-storage-engine_2.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-storage-engine` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 701.3 KiB | [postgresql-16-storage-engine_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-16-storage-engine_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-storage-engine` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 692.7 KiB | [postgresql-16-storage-engine_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-16-storage-engine_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-storage-engine` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 604.7 KiB | [postgresql-16-storage-engine_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-16-storage-engine_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-storage-engine` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 599.0 KiB | [postgresql-16-storage-engine_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-16-storage-engine_2.3.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-storage-engine` | `2.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 601.9 KiB | [postgresql-16-storage-engine_2.3.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-16-storage-engine_2.3.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-storage-engine` | `2.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 595.4 KiB | [postgresql-16-storage-engine_2.3.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-16-storage-engine_2.3.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `storage_engine_15` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 279.2 KiB | [storage_engine_15-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/storage_engine_15-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `storage_engine_15` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 268.0 KiB | [storage_engine_15-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/storage_engine_15-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `storage_engine_15` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 262.0 KiB | [storage_engine_15-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/storage_engine_15-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `storage_engine_15` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 255.5 KiB | [storage_engine_15-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/storage_engine_15-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `storage_engine_15` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 265.2 KiB | [storage_engine_15-2.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/storage_engine_15-2.3.0-1PIGSTY.el10.x86_64.rpm) |
| `storage_engine_15` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 257.6 KiB | [storage_engine_15-2.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/storage_engine_15-2.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-storage-engine` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 596.4 KiB | [postgresql-15-storage-engine_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-15-storage-engine_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-storage-engine` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 578.2 KiB | [postgresql-15-storage-engine_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-15-storage-engine_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-storage-engine` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 598.1 KiB | [postgresql-15-storage-engine_2.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-15-storage-engine_2.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-storage-engine` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 581.1 KiB | [postgresql-15-storage-engine_2.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-15-storage-engine_2.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-storage-engine` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 724.3 KiB | [postgresql-15-storage-engine_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-15-storage-engine_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-storage-engine` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 713.9 KiB | [postgresql-15-storage-engine_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-15-storage-engine_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-storage-engine` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 625.9 KiB | [postgresql-15-storage-engine_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-15-storage-engine_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-storage-engine` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 617.4 KiB | [postgresql-15-storage-engine_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-15-storage-engine_2.3.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-storage-engine` | `2.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 622.7 KiB | [postgresql-15-storage-engine_2.3.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-15-storage-engine_2.3.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-storage-engine` | `2.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 614.4 KiB | [postgresql-15-storage-engine_2.3.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-15-storage-engine_2.3.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `storage_engine_14` | `1.3.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 225.2 KiB | [storage_engine_14-1.3.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/storage_engine_14-1.3.4-1PIGSTY.el8.x86_64.rpm) |
| `storage_engine_14` | `1.3.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 216.0 KiB | [storage_engine_14-1.3.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/storage_engine_14-1.3.4-1PIGSTY.el8.aarch64.rpm) |
| `storage_engine_14` | `1.3.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 211.1 KiB | [storage_engine_14-1.3.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/storage_engine_14-1.3.4-1PIGSTY.el9.x86_64.rpm) |
| `storage_engine_14` | `1.3.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 204.9 KiB | [storage_engine_14-1.3.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/storage_engine_14-1.3.4-1PIGSTY.el9.aarch64.rpm) |
| `storage_engine_14` | `1.3.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 213.2 KiB | [storage_engine_14-1.3.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/storage_engine_14-1.3.4-1PIGSTY.el10.x86_64.rpm) |
| `storage_engine_14` | `1.3.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 206.2 KiB | [storage_engine_14-1.3.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/storage_engine_14-1.3.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-storage-engine` | `1.3.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 477.5 KiB | [postgresql-14-storage-engine_1.3.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-14-storage-engine_1.3.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-storage-engine` | `1.3.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 463.5 KiB | [postgresql-14-storage-engine_1.3.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/storage-engine/postgresql-14-storage-engine_1.3.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-storage-engine` | `1.3.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 479.8 KiB | [postgresql-14-storage-engine_1.3.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-14-storage-engine_1.3.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-storage-engine` | `1.3.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 466.0 KiB | [postgresql-14-storage-engine_1.3.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/storage-engine/postgresql-14-storage-engine_1.3.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-storage-engine` | `1.3.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 586.3 KiB | [postgresql-14-storage-engine_1.3.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-14-storage-engine_1.3.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-storage-engine` | `1.3.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 578.4 KiB | [postgresql-14-storage-engine_1.3.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/storage-engine/postgresql-14-storage-engine_1.3.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-storage-engine` | `1.3.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 501.5 KiB | [postgresql-14-storage-engine_1.3.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-14-storage-engine_1.3.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-storage-engine` | `1.3.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 494.8 KiB | [postgresql-14-storage-engine_1.3.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/storage-engine/postgresql-14-storage-engine_1.3.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-storage-engine` | `1.3.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 500.7 KiB | [postgresql-14-storage-engine_1.3.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-14-storage-engine_1.3.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-storage-engine` | `1.3.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 494.4 KiB | [postgresql-14-storage-engine_1.3.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/storage-engine/postgresql-14-storage-engine_1.3.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/saulojb/storage_engine" title="Repository" icon="github" subtitle="github.com/saulojb/storage_engine" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="storage_engine-2.3.0.tar.gz" >}}
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

Sources: [README v2.3.0](https://github.com/saulojb/storage_engine/blob/v2.3.0/README.md), [release v2.3.0](https://github.com/saulojb/storage_engine/releases/tag/v2.3.0), [PGXN 2.3.0](https://pgxn.org/dist/storage_engine/2.3.0/), [current README](https://github.com/saulojb/storage_engine/blob/main/README.md)

`storage_engine` 2.3.0 provides two PostgreSQL table access methods in the `engine` schema:

- `colcompress` for column-oriented compressed storage with vectorized filtering, vectorized aggregation, parallel scans, and stripe/chunk min/max pruning.
- `rowcompress` for row-batch compression with parallel scans, index scans, and batch metadata.

```sql
CREATE EXTENSION storage_engine;
```

### Quick Start

Create tables using either access method. Version 2.2 and later accept per-table options directly in `CREATE TABLE ... WITH (...)`.

```sql
CREATE TABLE events (
  ts timestamptz NOT NULL,
  user_id bigint,
  event_type text,
  value float8
) USING colcompress
  WITH (compression = 'zstd', compression_level = 9, orderby = 'ts ASC');

CREATE TABLE logs (
  id bigserial,
  logged_at timestamptz NOT NULL,
  message text
) USING rowcompress
  WITH (batch_size = 10000, compression = 'zstd');

SELECT event_type, count(*), avg(value)
FROM events
WHERE ts > now() - interval '1 day'
GROUP BY 1;
```

Version 2.3 expands `colcompress` vectorized aggregation with simple `sum(expression)` shapes such as `sum(amount + price)`, post-aggregation arithmetic such as `sum(amount) + count(*)`, and corrected `avg(int8)` behavior in parallel plans.

### Main Tuning Knobs

Session-level GUCs documented upstream include:

- `storage_engine.compression`
- `storage_engine.compression_level`
- `storage_engine.stripe_row_limit`
- `storage_engine.chunk_group_row_limit`
- `storage_engine.enable_parallel_execution`
- `storage_engine.min_parallel_processes`
- `storage_engine.enable_vectorization`
- `storage_engine.enable_vectorized_groupagg`
- `storage_engine.enable_automatic_plan`
- `storage_engine.enable_dml`
- `storage_engine.enable_custom_scan`
- `storage_engine.enable_qual_pushdown`
- `storage_engine.qual_pushdown_correlation_threshold`
- `storage_engine.max_custom_scan_paths`
- `storage_engine.enable_engine_index_scan`
- `storage_engine.enable_column_cache`
- `storage_engine.column_cache_size`
- `storage_engine.debug_vectorized_groupagg_fallback`
- `storage_engine.planner_debug_level`
- `storage_engine.maintenance_auto_enabled`
- `storage_engine.maintenance_auto_naptime`
- `storage_engine.maintenance_auto_database`

The README says these GUCs become visible once the shared library is loaded; add `storage_engine` to `shared_preload_libraries` if you want them available immediately in every session or need the built-in maintenance background worker.

### Types and Operators

`engine.uint8` stores unsigned 64-bit values for `colcompress` workloads that need the full `0` through `2^64 - 1` range. Upstream documents comparison operators (`=`, `<>`, `<`, `<=`, `>`, `>=`), B-tree and hash opclasses, casts to and from `bigint`, `numeric`, and `text`, plus `engine.min`, `engine.max`, and `engine.sum` aggregates. The vectorized planner can dispatch `engine.vmin`, `engine.vmax`, and `engine.vsum` on `colcompress` tables.

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
CALL engine.colcompress_repack('events');
CALL engine.colcompress_repack('events', min_fill_ratio => 0.7);
CALL engine.colcompress_merge_incremental('events', max_stripes => 64);
CALL engine.smart_update(
  'events'::regclass,
  'value = value * 1.1',
  'event_type = ''purchase'''
);
```

Use `engine.colcompress_merge()` after bulk loads when the `orderby` key should be globally sorted for pruning. Use `CALL engine.colcompress_repack()` to compact low-fill stripes, and `CALL engine.colcompress_merge_incremental()` for lower-lock maintenance that processes dirty stripes in batches.

For `rowcompress` tables:

```sql
SELECT engine.alter_rowcompress_table_set(
  'logs'::regclass,
  batch_size => 10000,
  compression => 'zstd',
  compression_level => 5
);

SELECT engine.rowcompress_repack('logs');
CALL engine.rowcompress_merge_incremental('logs', max_batches => 128);
SELECT * FROM engine.rowcompress_scan_stats();
```

Operational views include `engine.colcompress_options`, `engine.colcompress_stripes`, `engine.rowcompress_options`, `engine.rowcompress_batches`, and `engine.storage_health`. `engine.storage_maintenance_recommendation(table)` returns health metrics and a recommended action for one table, and `CALL engine.storage_maintenance_auto(...)` can dispatch maintenance manually or through the built-in background worker.

### When to Use Which AM

- Use `colcompress` for analytical scans, aggregates, and range predicates where projection, vectorization, and stripe/chunk pruning pay off.
- Use `rowcompress` for append-heavy logs or wide rows that are usually fetched together, where compression matters more than column projection.
- For point lookups on `colcompress`, use per-table `index_scan => true` or session-level `storage_engine.enable_engine_index_scan = on`; for analytical range scans, prefer `index_scan => false` with `engine.colcompress_merge()` and an `orderby` key.

### Caveats

- The packaged version in this repo is `2.3.0` for PostgreSQL 15 through 18. Upstream 2.x also tests PostgreSQL 19 devel, but PG19 is not in this repo's package matrix. PostgreSQL 12, 13, and 14 users should stay on upstream 1.3.4.
- The upstream default branch README has moved past the packaged 2.3.0 release; this stub follows `extension.csv` and the v2.3.0 release/PGXN docs.
- Upgrade existing installations with `ALTER EXTENSION storage_engine UPDATE TO '2.3.0';`.
- `colcompress` and `rowcompress` do not support foreign keys or `AFTER ROW` triggers.
- `pg_repack` cannot be used on these table access methods. `engine.colcompress_repack()` acquires `AccessExclusiveLock`, so schedule it during maintenance windows for large tables; the incremental merge procedures are the lower-lock option for dirty stripes or batches.
- `VACUUM FULL`, `CLUSTER`, and `CREATE UNLOGGED TABLE ... USING colcompress` are not supported; upstream recommends the extension's merge/repack functions instead.
- On `colcompress`, combining `orderby` with B-tree indexes can disable the sort-on-write path, and B-tree indexes on ordered columns can defeat stripe pruning for range queries. Use `engine.colcompress_merge()` after loading data when global ordering matters, and prefer `index_scan => false` for analytical tables.
- If `citus` or `pg_cron` is also preloaded, upstream documents the load order as `shared_preload_libraries = 'pg_cron,citus,storage_engine'`; `citus` must appear before `storage_engine`.
