---
title: "vasco"
linkTitle: "vasco"
description: "discover hidden correlations in your data with MIC"
weight: 4660
categories: ["Func"]
width: full
---

discover hidden correlations in your data with MIC

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4660** | {{< badge content="vasco" link="https://github.com/Florents-Tselai/vasco" >}} | {{< ext "vasco" "vasco" >}} | `0.1.0` | {{< category "FUNC" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "ddsketch" >}} {{< ext "tdigest" >}} {{< ext "uuid-ossp" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/vasco" >}} | `0.1.0` | {{< badge content="18" color="red" alt="vasco_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `vasco_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/vasco" >}} | `0.1.0` | {{< badge content="18" color="red" alt="postgresql-18-vasco" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-vasco` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "vasco_18" >}}     | {{< pkg "vasco_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/vasco_17-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "vasco_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/vasco_16-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "vasco_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/vasco_15-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "vasco_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/vasco_14-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "vasco_18" >}}     | {{< pkg "vasco_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/vasco_17-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "vasco_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/vasco_16-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "vasco_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/vasco_15-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "vasco_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/vasco_14-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "vasco_18" >}}     | {{< pkg "vasco_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/vasco_17-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "vasco_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/vasco_16-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "vasco_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/vasco_15-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "vasco_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/vasco_14-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "vasco_18" >}}     | {{< pkg "vasco_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/vasco_17-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "vasco_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/vasco_16-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "vasco_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/vasco_15-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "vasco_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/vasco_14-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-vasco" >}}     | {{< pkg "postgresql-17-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-17-vasco_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-16-vasco_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-15-vasco_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-14-vasco_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-vasco" >}}     | {{< pkg "postgresql-17-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-17-vasco_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-16-vasco_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-15-vasco_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-14-vasco_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-vasco" >}}     | {{< pkg "postgresql-17-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-17-vasco_0.1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-16-vasco_0.1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-15-vasco_0.1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-14-vasco_0.1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-vasco" >}}     | {{< pkg "postgresql-17-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-17-vasco_0.1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-16-vasco_0.1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-15-vasco_0.1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-14-vasco_0.1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-vasco" >}}     | {{< pkg "postgresql-17-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-17-vasco_0.1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-16-vasco_0.1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-15-vasco_0.1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-14-vasco_0.1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-vasco" >}}     | {{< pkg "postgresql-17-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-17-vasco_0.1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-16-vasco_0.1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-15-vasco_0.1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-vasco" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-14-vasco_0.1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `vasco_17` | 0.1.0 | `el8.x86_64` | pigsty | 40.1 KiB | [vasco_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vasco_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `vasco_17` | 0.1.0 | `el8.aarch64` | pigsty | 38.7 KiB | [vasco_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vasco_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `vasco_17` | 0.1.0 | `el9.aarch64` | pigsty | 37.1 KiB | [vasco_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vasco_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `vasco_17` | 0.1.0 | `el9.x86_64` | pigsty | 38.2 KiB | [vasco_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vasco_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-vasco` | 0.1.0 | `d12.x86_64` | pigsty | 56.0 KiB | [postgresql-17-vasco_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-17-vasco_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-vasco` | 0.1.0 | `d12.aarch64` | pigsty | 54.9 KiB | [postgresql-17-vasco_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-17-vasco_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-vasco` | 0.1.0 | `u22.x86_64` | pigsty | 57.8 KiB | [postgresql-17-vasco_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-17-vasco_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-vasco` | 0.1.0 | `u22.aarch64` | pigsty | 56.2 KiB | [postgresql-17-vasco_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-17-vasco_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-vasco` | 0.1.0 | `u24.x86_64` | pigsty | 54.7 KiB | [postgresql-17-vasco_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-17-vasco_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-vasco` | 0.1.0 | `u24.aarch64` | pigsty | 53.0 KiB | [postgresql-17-vasco_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-17-vasco_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `vasco_16` | 0.1.0 | `el8.x86_64` | pigsty | 40.1 KiB | [vasco_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vasco_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `vasco_16` | 0.1.0 | `el8.aarch64` | pigsty | 38.7 KiB | [vasco_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vasco_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `vasco_16` | 0.1.0 | `el9.x86_64` | pigsty | 38.2 KiB | [vasco_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vasco_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `vasco_16` | 0.1.0 | `el9.aarch64` | pigsty | 37.1 KiB | [vasco_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vasco_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-vasco` | 0.1.0 | `d12.x86_64` | pigsty | 55.9 KiB | [postgresql-16-vasco_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-16-vasco_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-vasco` | 0.1.0 | `d12.aarch64` | pigsty | 54.8 KiB | [postgresql-16-vasco_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-16-vasco_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-vasco` | 0.1.0 | `u22.aarch64` | pigsty | 56.2 KiB | [postgresql-16-vasco_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-16-vasco_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-vasco` | 0.1.0 | `u22.x86_64` | pigsty | 57.8 KiB | [postgresql-16-vasco_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-16-vasco_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-vasco` | 0.1.0 | `u24.x86_64` | pigsty | 54.7 KiB | [postgresql-16-vasco_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-16-vasco_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-vasco` | 0.1.0 | `u24.aarch64` | pigsty | 53.0 KiB | [postgresql-16-vasco_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-16-vasco_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `vasco_15` | 0.1.0 | `el8.x86_64` | pigsty | 40.6 KiB | [vasco_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vasco_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `vasco_15` | 0.1.0 | `el8.aarch64` | pigsty | 39.3 KiB | [vasco_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vasco_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `vasco_15` | 0.1.0 | `el9.x86_64` | pigsty | 40.4 KiB | [vasco_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vasco_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `vasco_15` | 0.1.0 | `el9.aarch64` | pigsty | 39.8 KiB | [vasco_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vasco_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-vasco` | 0.1.0 | `d12.aarch64` | pigsty | 55.6 KiB | [postgresql-15-vasco_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-15-vasco_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-vasco` | 0.1.0 | `d12.x86_64` | pigsty | 56.6 KiB | [postgresql-15-vasco_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-15-vasco_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-vasco` | 0.1.0 | `u22.aarch64` | pigsty | 58.5 KiB | [postgresql-15-vasco_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-15-vasco_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-vasco` | 0.1.0 | `u22.x86_64` | pigsty | 59.4 KiB | [postgresql-15-vasco_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-15-vasco_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-vasco` | 0.1.0 | `u24.x86_64` | pigsty | 56.3 KiB | [postgresql-15-vasco_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-15-vasco_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-vasco` | 0.1.0 | `u24.aarch64` | pigsty | 55.6 KiB | [postgresql-15-vasco_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-15-vasco_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `vasco_14` | 0.1.0 | `el8.x86_64` | pigsty | 40.6 KiB | [vasco_14-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vasco_14-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `vasco_14` | 0.1.0 | `el8.aarch64` | pigsty | 39.3 KiB | [vasco_14-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vasco_14-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `vasco_14` | 0.1.0 | `el9.x86_64` | pigsty | 40.4 KiB | [vasco_14-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vasco_14-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `vasco_14` | 0.1.0 | `el9.aarch64` | pigsty | 39.7 KiB | [vasco_14-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vasco_14-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-vasco` | 0.1.0 | `d12.x86_64` | pigsty | 56.6 KiB | [postgresql-14-vasco_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-14-vasco_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-vasco` | 0.1.0 | `d12.aarch64` | pigsty | 55.5 KiB | [postgresql-14-vasco_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-14-vasco_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-vasco` | 0.1.0 | `u22.x86_64` | pigsty | 59.4 KiB | [postgresql-14-vasco_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-14-vasco_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-vasco` | 0.1.0 | `u22.aarch64` | pigsty | 58.5 KiB | [postgresql-14-vasco_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-14-vasco_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-vasco` | 0.1.0 | `u24.x86_64` | pigsty | 56.3 KiB | [postgresql-14-vasco_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-14-vasco_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-vasco` | 0.1.0 | `u24.aarch64` | pigsty | 55.5 KiB | [postgresql-14-vasco_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-14-vasco_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `vasco_13` | 0.1.0 | `el8.aarch64` | pigsty | 39.2 KiB | [vasco_13-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vasco_13-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `vasco_13` | 0.1.0 | `el8.x86_64` | pigsty | 40.5 KiB | [vasco_13-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vasco_13-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `vasco_13` | 0.1.0 | `el9.aarch64` | pigsty | 39.7 KiB | [vasco_13-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vasco_13-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `vasco_13` | 0.1.0 | `el9.x86_64` | pigsty | 40.4 KiB | [vasco_13-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vasco_13-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-vasco` | 0.1.0 | `d12.aarch64` | pigsty | 55.5 KiB | [postgresql-13-vasco_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-13-vasco_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-vasco` | 0.1.0 | `d12.x86_64` | pigsty | 56.4 KiB | [postgresql-13-vasco_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vasco/postgresql-13-vasco_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-vasco` | 0.1.0 | `u22.aarch64` | pigsty | 58.3 KiB | [postgresql-13-vasco_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-13-vasco_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-vasco` | 0.1.0 | `u22.x86_64` | pigsty | 59.3 KiB | [postgresql-13-vasco_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vasco/postgresql-13-vasco_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-vasco` | 0.1.0 | `u24.aarch64` | pigsty | 55.3 KiB | [postgresql-13-vasco_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-13-vasco_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-vasco` | 0.1.0 | `u24.x86_64` | pigsty | 56.1 KiB | [postgresql-13-vasco_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vasco/postgresql-13-vasco_0.1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Florents-Tselai/vasco" title="Repository" icon="github" subtitle="github.com/Florents-Tselai/vasco" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="vasco-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get vasco; # get vasco source code
pig build dep vasco; # install build dependencies
pig build pkg vasco; # build extension rpm or deb
pig build ext vasco; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install vasco; # install by extension name, for the current active PG version
pig ext install vasco; # install via package alias, for the active PG version
pig ext install vasco -v 18;   # install for PG 18
pig ext install vasco -v 17;   # install for PG 17
pig ext install vasco -v 16;   # install for PG 16
pig ext install vasco -v 15;   # install for PG 15
pig ext install vasco -v 14;   # install for PG 14
pig ext install vasco -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION vasco;
```

