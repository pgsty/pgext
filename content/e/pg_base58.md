---
title: "pg_base58"
linkTitle: "pg_base58"
description: "Base58 Encoder/Decoder Extension for PostgreSQL"
weight: 4830
categories: ["Func"]
width: full
---

Base58 Encoder/Decoder Extension for PostgreSQL

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4830** | {{< badge content="pg_base58" link="https://github.com/Fell-x27/pg_base58" >}} | {{< ext "pg_base58" "pg_base58" >}} | `0.0.1` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "url_encode" >}} {{< ext "pg_cardano" >}} {{< ext "base36" >}} {{< ext "base62" >}} {{< ext "pg_polyline" >}} {{< ext "uri" >}} {{< ext "pg_curl" >}} {{< ext "pg_rewrite" >}} |

> [!Note] pgrx=0.12.1


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_base58" >}} | `0.0.1` | {{< badge content="18" color="red" alt="pg_base58_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_base58_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_base58" >}} | `0.0.1` | {{< badge content="18" color="red" alt="postgresql-18-pg-base58" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-base58` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_base58_18" >}}     | {{< pkg "pg_base58_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_17-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_base58_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_16-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_base58_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_15-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_base58_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_14-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_base58_18" >}}     | {{< pkg "pg_base58_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_17-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_base58_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_16-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_base58_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_15-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_base58_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_14-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_base58_18" >}}     | {{< pkg "pg_base58_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_17-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_base58_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_16-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_base58_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_15-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_base58_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_14-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_base58_18" >}}     | {{< pkg "pg_base58_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_17-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_base58_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_16-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_base58_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_15-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_base58_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_14-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-base58" >}}     | {{< pkg "postgresql-17-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-base58" >}}     | {{< pkg "postgresql-17-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-base58" >}}     | {{< pkg "postgresql-17-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-base58" >}}     | {{< pkg "postgresql-17-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-base58" >}}     | {{< pkg "postgresql-17-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-base58" >}}     | {{< pkg "postgresql-17-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-base58" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_base58_17` | 0.0.1 | `el8.x86_64` | pigsty | 197.0 KiB | [pg_base58_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_base58_17` | 0.0.1 | `el8.aarch64` | pigsty | 183.7 KiB | [pg_base58_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_base58_17` | 0.0.1 | `el9.aarch64` | pigsty | 197.1 KiB | [pg_base58_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_base58_17` | 0.0.1 | `el9.x86_64` | pigsty | 201.1 KiB | [pg_base58_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-base58` | 0.0.1 | `d12.x86_64` | pigsty | 159.2 KiB | [postgresql-17-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-base58` | 0.0.1 | `d12.aarch64` | pigsty | 141.9 KiB | [postgresql-17-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-base58` | 0.0.1 | `u22.x86_64` | pigsty | 173.6 KiB | [postgresql-17-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-base58` | 0.0.1 | `u22.aarch64` | pigsty | 164.6 KiB | [postgresql-17-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-base58` | 0.0.1 | `u24.x86_64` | pigsty | 172.3 KiB | [postgresql-17-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-base58` | 0.0.1 | `u24.aarch64` | pigsty | 163.6 KiB | [postgresql-17-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_base58_16` | 0.0.1 | `el8.x86_64` | pigsty | 197.0 KiB | [pg_base58_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_base58_16` | 0.0.1 | `el8.aarch64` | pigsty | 183.7 KiB | [pg_base58_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_base58_16` | 0.0.1 | `el9.x86_64` | pigsty | 201.2 KiB | [pg_base58_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_base58_16` | 0.0.1 | `el9.aarch64` | pigsty | 197.3 KiB | [pg_base58_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-base58` | 0.0.1 | `d12.x86_64` | pigsty | 159.2 KiB | [postgresql-16-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-base58` | 0.0.1 | `d12.aarch64` | pigsty | 141.9 KiB | [postgresql-16-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-base58` | 0.0.1 | `u22.aarch64` | pigsty | 164.6 KiB | [postgresql-16-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-base58` | 0.0.1 | `u22.x86_64` | pigsty | 173.6 KiB | [postgresql-16-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-base58` | 0.0.1 | `u24.x86_64` | pigsty | 172.2 KiB | [postgresql-16-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-base58` | 0.0.1 | `u24.aarch64` | pigsty | 163.2 KiB | [postgresql-16-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_base58_15` | 0.0.1 | `el8.x86_64` | pigsty | 197.0 KiB | [pg_base58_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_base58_15` | 0.0.1 | `el8.aarch64` | pigsty | 183.7 KiB | [pg_base58_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_base58_15` | 0.0.1 | `el9.x86_64` | pigsty | 201.1 KiB | [pg_base58_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_base58_15` | 0.0.1 | `el9.aarch64` | pigsty | 197.1 KiB | [pg_base58_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-base58` | 0.0.1 | `d12.aarch64` | pigsty | 141.9 KiB | [postgresql-15-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-base58` | 0.0.1 | `d12.x86_64` | pigsty | 159.2 KiB | [postgresql-15-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-base58` | 0.0.1 | `u22.aarch64` | pigsty | 164.6 KiB | [postgresql-15-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-base58` | 0.0.1 | `u22.x86_64` | pigsty | 173.6 KiB | [postgresql-15-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-base58` | 0.0.1 | `u24.x86_64` | pigsty | 172.2 KiB | [postgresql-15-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-base58` | 0.0.1 | `u24.aarch64` | pigsty | 163.2 KiB | [postgresql-15-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_base58_14` | 0.0.1 | `el8.x86_64` | pigsty | 197.0 KiB | [pg_base58_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_base58_14` | 0.0.1 | `el8.aarch64` | pigsty | 183.7 KiB | [pg_base58_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_base58_14` | 0.0.1 | `el9.x86_64` | pigsty | 201.2 KiB | [pg_base58_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_base58_14` | 0.0.1 | `el9.aarch64` | pigsty | 197.1 KiB | [pg_base58_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-base58` | 0.0.1 | `d12.x86_64` | pigsty | 159.2 KiB | [postgresql-14-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-base58` | 0.0.1 | `d12.aarch64` | pigsty | 141.9 KiB | [postgresql-14-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-base58` | 0.0.1 | `u22.x86_64` | pigsty | 173.6 KiB | [postgresql-14-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-base58` | 0.0.1 | `u22.aarch64` | pigsty | 164.4 KiB | [postgresql-14-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-base58` | 0.0.1 | `u24.x86_64` | pigsty | 172.3 KiB | [postgresql-14-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-base58` | 0.0.1 | `u24.aarch64` | pigsty | 163.6 KiB | [postgresql-14-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_base58_13` | 0.0.1 | `el8.aarch64` | pigsty | 183.7 KiB | [pg_base58_13-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_13-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_base58_13` | 0.0.1 | `el8.x86_64` | pigsty | 197.0 KiB | [pg_base58_13-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_13-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_base58_13` | 0.0.1 | `el9.aarch64` | pigsty | 197.0 KiB | [pg_base58_13-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_13-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_base58_13` | 0.0.1 | `el9.x86_64` | pigsty | 201.1 KiB | [pg_base58_13-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_13-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-base58` | 0.0.1 | `d12.aarch64` | pigsty | 141.9 KiB | [postgresql-13-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-13-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-base58` | 0.0.1 | `d12.x86_64` | pigsty | 159.2 KiB | [postgresql-13-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-13-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-base58` | 0.0.1 | `u22.aarch64` | pigsty | 164.4 KiB | [postgresql-13-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-13-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-base58` | 0.0.1 | `u22.x86_64` | pigsty | 173.6 KiB | [postgresql-13-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-13-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-base58` | 0.0.1 | `u24.aarch64` | pigsty | 163.6 KiB | [postgresql-13-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-13-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-base58` | 0.0.1 | `u24.x86_64` | pigsty | 172.3 KiB | [postgresql-13-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-13-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Fell-x27/pg_base58" title="Repository" icon="github" subtitle="github.com/Fell-x27/pg_base58" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_base58-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_base58; # get pg_base58 source code
pig build dep pg_base58; # install build dependencies
pig build pkg pg_base58; # build extension rpm or deb
pig build ext pg_base58; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_base58; # install by extension name, for the current active PG version
pig ext install pg_base58; # install via package alias, for the active PG version
pig ext install pg_base58 -v 17;   # install for PG 17
pig ext install pg_base58 -v 16;   # install for PG 16
pig ext install pg_base58 -v 15;   # install for PG 15
pig ext install pg_base58 -v 14;   # install for PG 14
pig ext install pg_base58 -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_base58;
```

