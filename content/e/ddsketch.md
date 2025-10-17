---
title: "ddsketch"
linkTitle: "ddsketch"
description: "Provides ddsketch aggregate function"
weight: 4650
categories: ["Func"]
width: full
---

Provides ddsketch aggregate function

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4650** | {{< badge content="ddsketch" link="https://github.com/tvondra/ddsketch" >}} | {{< ext "ddsketch" "ddsketch" >}} | `1.0.1` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "omnisketch" >}} {{< ext "quantile" >}} {{< ext "lower_quantile" >}} {{< ext "topn" >}} {{< ext "count_distinct" >}} {{< ext "hll" >}} {{< ext "first_last_agg" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/ddsketch" >}} | `1.0.1` | {{< badge content="18" color="red" alt="ddsketch_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `ddsketch_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/ddsketch" >}} | `1.0.1` | {{< badge content="18" color="red" alt="postgresql-18-ddsketch" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-ddsketch` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "ddsketch_18" >}}     | {{< pkg "ddsketch_17" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_17-1.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "ddsketch_16" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_16-1.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "ddsketch_15" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_15-1.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "ddsketch_14" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_14-1.0.1-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "ddsketch_18" >}}     | {{< pkg "ddsketch_17" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_17-1.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "ddsketch_16" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_16-1.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "ddsketch_15" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_15-1.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "ddsketch_14" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_14-1.0.1-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "ddsketch_18" >}}     | {{< pkg "ddsketch_17" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_17-1.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "ddsketch_16" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_16-1.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "ddsketch_15" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_15-1.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "ddsketch_14" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_14-1.0.1-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "ddsketch_18" >}}     | {{< pkg "ddsketch_17" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_17-1.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "ddsketch_16" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_16-1.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "ddsketch_15" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_15-1.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "ddsketch_14" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_14-1.0.1-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-ddsketch" >}}     | {{< pkg "postgresql-17-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-ddsketch" >}}     | {{< pkg "postgresql-17-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-ddsketch" >}}     | {{< pkg "postgresql-17-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-ddsketch" >}}     | {{< pkg "postgresql-17-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-ddsketch" >}}     | {{< pkg "postgresql-17-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-ddsketch" >}}     | {{< pkg "postgresql-17-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-ddsketch" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `ddsketch_17` | 1.0.1 | `el8.x86_64` | pigsty | 32.5 KiB | [ddsketch_17-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_17-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_17` | 1.0.1 | `el8.aarch64` | pigsty | 31.0 KiB | [ddsketch_17-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_17-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_17` | 1.0.1 | `el9.aarch64` | pigsty | 32.0 KiB | [ddsketch_17-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_17-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `ddsketch_17` | 1.0.1 | `el9.x86_64` | pigsty | 33.2 KiB | [ddsketch_17-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_17-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-ddsketch` | 1.0.1 | `d12.x86_64` | pigsty | 66.0 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-ddsketch` | 1.0.1 | `d12.aarch64` | pigsty | 65.1 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-ddsketch` | 1.0.1 | `u22.x86_64` | pigsty | 69.7 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-ddsketch` | 1.0.1 | `u22.aarch64` | pigsty | 69.2 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-ddsketch` | 1.0.1 | `u24.x86_64` | pigsty | 64.2 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-ddsketch` | 1.0.1 | `u24.aarch64` | pigsty | 65.3 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `ddsketch_16` | 1.0.1 | `el8.x86_64` | pigsty | 32.5 KiB | [ddsketch_16-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_16-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_16` | 1.0.1 | `el8.aarch64` | pigsty | 31.0 KiB | [ddsketch_16-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_16-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_16` | 1.0.1 | `el9.x86_64` | pigsty | 33.2 KiB | [ddsketch_16-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_16-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `ddsketch_16` | 1.0.1 | `el9.aarch64` | pigsty | 31.9 KiB | [ddsketch_16-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_16-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-ddsketch` | 1.0.1 | `d12.x86_64` | pigsty | 66.0 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-ddsketch` | 1.0.1 | `d12.aarch64` | pigsty | 65.1 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-ddsketch` | 1.0.1 | `u22.aarch64` | pigsty | 69.2 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-ddsketch` | 1.0.1 | `u22.x86_64` | pigsty | 69.7 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-ddsketch` | 1.0.1 | `u24.x86_64` | pigsty | 64.2 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-ddsketch` | 1.0.1 | `u24.aarch64` | pigsty | 65.3 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `ddsketch_15` | 1.0.1 | `el8.x86_64` | pigsty | 32.5 KiB | [ddsketch_15-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_15-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_15` | 1.0.1 | `el8.aarch64` | pigsty | 31.0 KiB | [ddsketch_15-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_15-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_15` | 1.0.1 | `el9.x86_64` | pigsty | 33.2 KiB | [ddsketch_15-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_15-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `ddsketch_15` | 1.0.1 | `el9.aarch64` | pigsty | 32.0 KiB | [ddsketch_15-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_15-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-ddsketch` | 1.0.1 | `d12.aarch64` | pigsty | 65.3 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-ddsketch` | 1.0.1 | `d12.x86_64` | pigsty | 66.0 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-ddsketch` | 1.0.1 | `u22.aarch64` | pigsty | 69.4 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-ddsketch` | 1.0.1 | `u22.x86_64` | pigsty | 69.7 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-ddsketch` | 1.0.1 | `u24.x86_64` | pigsty | 64.2 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-ddsketch` | 1.0.1 | `u24.aarch64` | pigsty | 65.3 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `ddsketch_14` | 1.0.1 | `el8.x86_64` | pigsty | 32.5 KiB | [ddsketch_14-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_14-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_14` | 1.0.1 | `el8.aarch64` | pigsty | 31.0 KiB | [ddsketch_14-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_14-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_14` | 1.0.1 | `el9.x86_64` | pigsty | 33.2 KiB | [ddsketch_14-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_14-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `ddsketch_14` | 1.0.1 | `el9.aarch64` | pigsty | 32.0 KiB | [ddsketch_14-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_14-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-ddsketch` | 1.0.1 | `d12.x86_64` | pigsty | 65.9 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-ddsketch` | 1.0.1 | `d12.aarch64` | pigsty | 65.3 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-ddsketch` | 1.0.1 | `u22.x86_64` | pigsty | 69.7 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-ddsketch` | 1.0.1 | `u22.aarch64` | pigsty | 69.3 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-ddsketch` | 1.0.1 | `u24.x86_64` | pigsty | 64.2 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-ddsketch` | 1.0.1 | `u24.aarch64` | pigsty | 65.2 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `ddsketch_13` | 1.0.1 | `el8.aarch64` | pigsty | 31.0 KiB | [ddsketch_13-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_13-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_13` | 1.0.1 | `el8.x86_64` | pigsty | 32.0 KiB | [ddsketch_13-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_13-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_13` | 1.0.1 | `el9.aarch64` | pigsty | 31.9 KiB | [ddsketch_13-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_13-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `ddsketch_13` | 1.0.1 | `el9.x86_64` | pigsty | 33.1 KiB | [ddsketch_13-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_13-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-ddsketch` | 1.0.1 | `d12.aarch64` | pigsty | 65.5 KiB | [postgresql-13-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-13-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-ddsketch` | 1.0.1 | `d12.x86_64` | pigsty | 65.8 KiB | [postgresql-13-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-13-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-ddsketch` | 1.0.1 | `u22.aarch64` | pigsty | 69.0 KiB | [postgresql-13-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-13-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-ddsketch` | 1.0.1 | `u22.x86_64` | pigsty | 69.4 KiB | [postgresql-13-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-13-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-ddsketch` | 1.0.1 | `u24.aarch64` | pigsty | 65.3 KiB | [postgresql-13-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-13-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-ddsketch` | 1.0.1 | `u24.x86_64` | pigsty | 64.0 KiB | [postgresql-13-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-13-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/ddsketch" title="Repository" icon="github" subtitle="github.com/tvondra/ddsketch" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="ddsketch-1.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get ddsketch; # get ddsketch source code
pig build dep ddsketch; # install build dependencies
pig build pkg ddsketch; # build extension rpm or deb
pig build ext ddsketch; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install ddsketch; # install by extension name, for the current active PG version
pig ext install ddsketch; # install via package alias, for the active PG version
pig ext install ddsketch -v 18;   # install for PG 18
pig ext install ddsketch -v 17;   # install for PG 17
pig ext install ddsketch -v 16;   # install for PG 16
pig ext install ddsketch -v 15;   # install for PG 15
pig ext install ddsketch -v 14;   # install for PG 14
pig ext install ddsketch -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION ddsketch;
```

