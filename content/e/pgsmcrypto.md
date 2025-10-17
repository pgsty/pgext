---
title: "pgsmcrypto"
linkTitle: "pgsmcrypto"
description: "PostgreSQL SM Algorithm Extension"
weight: 7070
categories: ["Sec"]
width: full
---

PostgreSQL SM Algorithm Extension

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7070** | {{< badge content="pgsmcrypto" link="https://github.com/zhuobie/pgsmcrypto" >}} | {{< ext "pgsmcrypto" "pgsmcrypto" >}} | `0.1.0` | {{< category "SEC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgsodium" >}} {{< ext "pgcryptokey" >}} {{< ext "pgcrypto" >}} {{< ext "pg_tde" >}} {{< ext "sslutils" >}} {{< ext "faker" >}} {{< ext "uuid-ossp" >}} {{< ext "lo" >}} |

> [!Note] pgrx=0.12.6


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgsmcrypto" >}} | `0.1.0` | {{< badge content="18" color="red" alt="pgsmcrypto_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgsmcrypto_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgsmcrypto" >}} | `0.1.0` | {{< badge content="18" color="red" alt="postgresql-18-pgsmcrypto" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pgsmcrypto` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pgsmcrypto_18" >}}     | {{< pkg "pgsmcrypto_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_17-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgsmcrypto_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_16-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgsmcrypto_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_15-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgsmcrypto_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_14-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pgsmcrypto_18" >}}     | {{< pkg "pgsmcrypto_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_17-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgsmcrypto_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_16-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgsmcrypto_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_15-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgsmcrypto_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_14-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pgsmcrypto_18" >}}     | {{< pkg "pgsmcrypto_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_17-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgsmcrypto_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_16-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgsmcrypto_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_15-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgsmcrypto_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_14-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pgsmcrypto_18" >}}     | {{< pkg "pgsmcrypto_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_17-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgsmcrypto_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_16-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgsmcrypto_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_15-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgsmcrypto_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_14-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pgsmcrypto" >}}     | {{< pkg "postgresql-17-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pgsmcrypto" >}}     | {{< pkg "postgresql-17-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pgsmcrypto" >}}     | {{< pkg "postgresql-17-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pgsmcrypto" >}}     | {{< pkg "postgresql-17-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pgsmcrypto" >}}     | {{< pkg "postgresql-17-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pgsmcrypto" >}}     | {{< pkg "postgresql-17-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pgsmcrypto" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsmcrypto_17` | 0.1.0 | `el8.x86_64` | pigsty | 764.8 KiB | [pgsmcrypto_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_17` | 0.1.0 | `el8.aarch64` | pigsty | 670.8 KiB | [pgsmcrypto_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_17` | 0.1.0 | `el9.aarch64` | pigsty | 732.6 KiB | [pgsmcrypto_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgsmcrypto_17` | 0.1.0 | `el9.x86_64` | pigsty | 783.5 KiB | [pgsmcrypto_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pgsmcrypto` | 0.1.0 | `d12.x86_64` | pigsty | 639.8 KiB | [postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgsmcrypto` | 0.1.0 | `d12.aarch64` | pigsty | 540.5 KiB | [postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgsmcrypto` | 0.1.0 | `u22.x86_64` | pigsty | 719.0 KiB | [postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgsmcrypto` | 0.1.0 | `u22.aarch64` | pigsty | 656.2 KiB | [postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgsmcrypto` | 0.1.0 | `u24.x86_64` | pigsty | 715.9 KiB | [postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgsmcrypto` | 0.1.0 | `u24.aarch64` | pigsty | 646.8 KiB | [postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsmcrypto_16` | 0.1.0 | `el8.x86_64` | pigsty | 764.9 KiB | [pgsmcrypto_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_16` | 0.1.0 | `el8.aarch64` | pigsty | 670.6 KiB | [pgsmcrypto_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_16` | 0.1.0 | `el9.x86_64` | pigsty | 783.7 KiB | [pgsmcrypto_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsmcrypto_16` | 0.1.0 | `el9.aarch64` | pigsty | 733.5 KiB | [pgsmcrypto_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pgsmcrypto` | 0.1.0 | `d12.x86_64` | pigsty | 640.3 KiB | [postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgsmcrypto` | 0.1.0 | `d12.aarch64` | pigsty | 540.4 KiB | [postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgsmcrypto` | 0.1.0 | `u22.aarch64` | pigsty | 656.0 KiB | [postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgsmcrypto` | 0.1.0 | `u22.x86_64` | pigsty | 720.7 KiB | [postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgsmcrypto` | 0.1.0 | `u24.x86_64` | pigsty | 715.9 KiB | [postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgsmcrypto` | 0.1.0 | `u24.aarch64` | pigsty | 643.7 KiB | [postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsmcrypto_15` | 0.1.0 | `el8.x86_64` | pigsty | 764.8 KiB | [pgsmcrypto_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_15` | 0.1.0 | `el8.aarch64` | pigsty | 670.8 KiB | [pgsmcrypto_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_15` | 0.1.0 | `el9.x86_64` | pigsty | 783.6 KiB | [pgsmcrypto_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsmcrypto_15` | 0.1.0 | `el9.aarch64` | pigsty | 734.0 KiB | [pgsmcrypto_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pgsmcrypto` | 0.1.0 | `d12.aarch64` | pigsty | 540.5 KiB | [postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgsmcrypto` | 0.1.0 | `d12.x86_64` | pigsty | 640.3 KiB | [postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgsmcrypto` | 0.1.0 | `u22.aarch64` | pigsty | 656.0 KiB | [postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgsmcrypto` | 0.1.0 | `u22.x86_64` | pigsty | 721.4 KiB | [postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgsmcrypto` | 0.1.0 | `u24.x86_64` | pigsty | 716.0 KiB | [postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgsmcrypto` | 0.1.0 | `u24.aarch64` | pigsty | 644.7 KiB | [postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsmcrypto_14` | 0.1.0 | `el8.x86_64` | pigsty | 764.8 KiB | [pgsmcrypto_14-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_14-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_14` | 0.1.0 | `el8.aarch64` | pigsty | 670.7 KiB | [pgsmcrypto_14-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_14-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_14` | 0.1.0 | `el9.x86_64` | pigsty | 785.3 KiB | [pgsmcrypto_14-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_14-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsmcrypto_14` | 0.1.0 | `el9.aarch64` | pigsty | 733.7 KiB | [pgsmcrypto_14-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_14-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pgsmcrypto` | 0.1.0 | `d12.x86_64` | pigsty | 640.1 KiB | [postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgsmcrypto` | 0.1.0 | `d12.aarch64` | pigsty | 540.2 KiB | [postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgsmcrypto` | 0.1.0 | `u22.x86_64` | pigsty | 721.4 KiB | [postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgsmcrypto` | 0.1.0 | `u22.aarch64` | pigsty | 655.9 KiB | [postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgsmcrypto` | 0.1.0 | `u24.x86_64` | pigsty | 716.2 KiB | [postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgsmcrypto` | 0.1.0 | `u24.aarch64` | pigsty | 646.7 KiB | [postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsmcrypto_13` | 0.1.0 | `el8.aarch64` | pigsty | 670.6 KiB | [pgsmcrypto_13-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_13-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_13` | 0.1.0 | `el8.x86_64` | pigsty | 765.0 KiB | [pgsmcrypto_13-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_13-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_13` | 0.1.0 | `el9.aarch64` | pigsty | 733.7 KiB | [pgsmcrypto_13-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_13-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgsmcrypto_13` | 0.1.0 | `el9.x86_64` | pigsty | 785.4 KiB | [pgsmcrypto_13-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_13-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pgsmcrypto` | 0.1.0 | `d12.aarch64` | pigsty | 540.2 KiB | [postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgsmcrypto` | 0.1.0 | `d12.x86_64` | pigsty | 639.9 KiB | [postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgsmcrypto` | 0.1.0 | `u22.aarch64` | pigsty | 655.9 KiB | [postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgsmcrypto` | 0.1.0 | `u22.x86_64` | pigsty | 721.3 KiB | [postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgsmcrypto` | 0.1.0 | `u24.aarch64` | pigsty | 646.6 KiB | [postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pgsmcrypto` | 0.1.0 | `u24.x86_64` | pigsty | 716.0 KiB | [postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/zhuobie/pgsmcrypto" title="Repository" icon="github" subtitle="github.com/zhuobie/pgsmcrypto" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgsmcrypto-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgsmcrypto; # get pgsmcrypto source code
pig build dep pgsmcrypto; # install build dependencies
pig build pkg pgsmcrypto; # build extension rpm or deb
pig build ext pgsmcrypto; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgsmcrypto; # install by extension name, for the current active PG version
pig ext install pgsmcrypto; # install via package alias, for the active PG version
pig ext install pgsmcrypto -v 17;   # install for PG 17
pig ext install pgsmcrypto -v 16;   # install for PG 16
pig ext install pgsmcrypto -v 15;   # install for PG 15
pig ext install pgsmcrypto -v 14;   # install for PG 14
pig ext install pgsmcrypto -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgsmcrypto;
```

