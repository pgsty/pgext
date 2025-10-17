---
title: "pg_idkit"
linkTitle: "pg_idkit"
description: "multi-tool for generating new/niche universally unique identifiers (ex. UUIDv6, ULID, KSUID)"
weight: 4500
categories: ["Func"]
width: full
---

multi-tool for generating new/niche universally unique identifiers (ex. UUIDv6, ULID, KSUID)

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4500** | {{< badge content="pg_idkit" link="https://github.com/VADOSWARE/pg_idkit" >}} | {{< ext "pg_idkit" "pg_idkit" >}} | `0.3.1` | {{< category "FUNC" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "uuid-ossp" >}} {{< ext "permuteseq" >}} {{< ext "pg_cardano" >}} {{< ext "pg_base58" >}} |

> [!Note] pgrx=0.15.0


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_idkit" >}} | `0.3.1` | {{< badge content="18" color="red" alt="pg_idkit_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_idkit_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_idkit" >}} | `0.3.1` | {{< badge content="18" color="red" alt="postgresql-18-pg-idkit" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-idkit` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_idkit_18" >}}     | {{< pkg "pg_idkit_17" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_17-0.3.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_idkit_16" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_16-0.3.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_idkit_15" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_15-0.3.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_idkit_14" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_14-0.3.1-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_idkit_18" >}}     | {{< pkg "pg_idkit_17" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_17-0.3.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_idkit_16" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_16-0.3.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_idkit_15" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_15-0.3.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_idkit_14" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_14-0.3.1-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_idkit_18" >}}     | {{< pkg "pg_idkit_17" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_17-0.3.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_idkit_16" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_16-0.3.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_idkit_15" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_15-0.3.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_idkit_14" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_14-0.3.1-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_idkit_18" >}}     | {{< pkg "pg_idkit_17" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_17-0.3.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_idkit_16" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_16-0.3.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_idkit_15" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_15-0.3.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_idkit_14" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_14-0.3.1-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-idkit" >}}     | {{< pkg "postgresql-17-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.3.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.3.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.3.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.3.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-idkit" >}}     | {{< pkg "postgresql-17-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.3.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.3.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.3.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.3.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-idkit" >}}     | {{< pkg "postgresql-17-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.3.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.3.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.3.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.3.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-idkit" >}}     | {{< pkg "postgresql-17-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.3.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.3.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.3.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.3.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-idkit" >}}     | {{< pkg "postgresql-17-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.3.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.3.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.3.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.3.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-idkit" >}}     | {{< pkg "postgresql-17-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.3.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.3.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.3.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-idkit" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.3.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_idkit_17` | 0.3.1 | `el8.x86_64` | pigsty | 375.0 KiB | [pg_idkit_17-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_17-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_17` | 0.3.1 | `el8.aarch64` | pigsty | 356.7 KiB | [pg_idkit_17-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_17-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_17` | 0.3.1 | `el9.aarch64` | pigsty | 379.1 KiB | [pg_idkit_17-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_17-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_idkit_17` | 0.3.1 | `el9.x86_64` | pigsty | 380.0 KiB | [pg_idkit_17-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_17-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-idkit` | 0.3.1 | `d12.x86_64` | pigsty | 312.0 KiB | [postgresql-17-pg-idkit_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-idkit` | 0.3.1 | `d12.aarch64` | pigsty | 289.7 KiB | [postgresql-17-pg-idkit_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-idkit` | 0.3.1 | `u22.x86_64` | pigsty | 340.3 KiB | [postgresql-17-pg-idkit_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-idkit` | 0.3.1 | `u22.aarch64` | pigsty | 334.8 KiB | [postgresql-17-pg-idkit_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-idkit` | 0.3.1 | `u24.x86_64` | pigsty | 336.1 KiB | [postgresql-17-pg-idkit_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-idkit` | 0.3.1 | `u24.aarch64` | pigsty | 330.3 KiB | [postgresql-17-pg-idkit_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_idkit_16` | 0.3.1 | `el8.x86_64` | pigsty | 375.2 KiB | [pg_idkit_16-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_16-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_16` | 0.3.1 | `el8.aarch64` | pigsty | 356.6 KiB | [pg_idkit_16-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_16-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_16` | 0.3.1 | `el9.x86_64` | pigsty | 380.3 KiB | [pg_idkit_16-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_16-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_idkit_16` | 0.3.1 | `el9.aarch64` | pigsty | 379.0 KiB | [pg_idkit_16-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_16-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-idkit` | 0.3.1 | `d12.x86_64` | pigsty | 312.1 KiB | [postgresql-16-pg-idkit_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-idkit` | 0.3.1 | `d12.aarch64` | pigsty | 289.7 KiB | [postgresql-16-pg-idkit_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-idkit` | 0.3.1 | `u22.aarch64` | pigsty | 334.8 KiB | [postgresql-16-pg-idkit_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-idkit` | 0.3.1 | `u22.x86_64` | pigsty | 340.1 KiB | [postgresql-16-pg-idkit_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-idkit` | 0.3.1 | `u24.x86_64` | pigsty | 336.0 KiB | [postgresql-16-pg-idkit_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-idkit` | 0.3.1 | `u24.aarch64` | pigsty | 330.3 KiB | [postgresql-16-pg-idkit_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_idkit_15` | 0.3.1 | `el8.x86_64` | pigsty | 375.1 KiB | [pg_idkit_15-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_15-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_15` | 0.3.1 | `el8.aarch64` | pigsty | 356.8 KiB | [pg_idkit_15-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_15-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_15` | 0.3.1 | `el9.x86_64` | pigsty | 380.2 KiB | [pg_idkit_15-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_15-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_idkit_15` | 0.3.1 | `el9.aarch64` | pigsty | 379.1 KiB | [pg_idkit_15-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_15-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-idkit` | 0.3.1 | `d12.aarch64` | pigsty | 289.9 KiB | [postgresql-15-pg-idkit_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-idkit` | 0.3.1 | `d12.x86_64` | pigsty | 312.3 KiB | [postgresql-15-pg-idkit_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-idkit` | 0.3.1 | `u22.aarch64` | pigsty | 334.7 KiB | [postgresql-15-pg-idkit_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-idkit` | 0.3.1 | `u22.x86_64` | pigsty | 340.4 KiB | [postgresql-15-pg-idkit_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-idkit` | 0.3.1 | `u24.x86_64` | pigsty | 336.3 KiB | [postgresql-15-pg-idkit_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-idkit` | 0.3.1 | `u24.aarch64` | pigsty | 330.3 KiB | [postgresql-15-pg-idkit_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_idkit_14` | 0.3.1 | `el8.x86_64` | pigsty | 374.9 KiB | [pg_idkit_14-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_14-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_14` | 0.3.1 | `el8.aarch64` | pigsty | 356.5 KiB | [pg_idkit_14-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_14-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_14` | 0.3.1 | `el9.x86_64` | pigsty | 380.1 KiB | [pg_idkit_14-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_14-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_idkit_14` | 0.3.1 | `el9.aarch64` | pigsty | 379.2 KiB | [pg_idkit_14-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_14-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-idkit` | 0.3.1 | `d12.x86_64` | pigsty | 311.9 KiB | [postgresql-14-pg-idkit_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-idkit` | 0.3.1 | `d12.aarch64` | pigsty | 289.9 KiB | [postgresql-14-pg-idkit_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-idkit` | 0.3.1 | `u22.x86_64` | pigsty | 340.3 KiB | [postgresql-14-pg-idkit_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-idkit` | 0.3.1 | `u22.aarch64` | pigsty | 335.0 KiB | [postgresql-14-pg-idkit_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-idkit` | 0.3.1 | `u24.x86_64` | pigsty | 336.1 KiB | [postgresql-14-pg-idkit_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-idkit` | 0.3.1 | `u24.aarch64` | pigsty | 330.2 KiB | [postgresql-14-pg-idkit_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_idkit_13` | 0.3.1 | `el8.aarch64` | pigsty | 356.7 KiB | [pg_idkit_13-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_13-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_13` | 0.3.1 | `el8.x86_64` | pigsty | 374.8 KiB | [pg_idkit_13-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_13-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_13` | 0.3.1 | `el9.aarch64` | pigsty | 379.4 KiB | [pg_idkit_13-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_13-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_idkit_13` | 0.3.1 | `el9.x86_64` | pigsty | 380.0 KiB | [pg_idkit_13-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_13-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-idkit` | 0.3.1 | `d12.aarch64` | pigsty | 289.6 KiB | [postgresql-13-pg-idkit_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-13-pg-idkit_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-idkit` | 0.3.1 | `d12.x86_64` | pigsty | 312.1 KiB | [postgresql-13-pg-idkit_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-13-pg-idkit_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-idkit` | 0.3.1 | `u22.aarch64` | pigsty | 334.9 KiB | [postgresql-13-pg-idkit_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-13-pg-idkit_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-idkit` | 0.3.1 | `u22.x86_64` | pigsty | 340.4 KiB | [postgresql-13-pg-idkit_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-13-pg-idkit_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-idkit` | 0.3.1 | `u24.aarch64` | pigsty | 330.2 KiB | [postgresql-13-pg-idkit_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-13-pg-idkit_0.3.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-idkit` | 0.3.1 | `u24.x86_64` | pigsty | 336.2 KiB | [postgresql-13-pg-idkit_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-13-pg-idkit_0.3.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/VADOSWARE/pg_idkit" title="Repository" icon="github" subtitle="github.com/VADOSWARE/pg_idkit" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_idkit-0.3.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_idkit; # get pg_idkit source code
pig build dep pg_idkit; # install build dependencies
pig build pkg pg_idkit; # build extension rpm or deb
pig build ext pg_idkit; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_idkit; # install by extension name, for the current active PG version
pig ext install pg_idkit; # install via package alias, for the active PG version
pig ext install pg_idkit -v 17;   # install for PG 17
pig ext install pg_idkit -v 16;   # install for PG 16
pig ext install pg_idkit -v 15;   # install for PG 15
pig ext install pg_idkit -v 14;   # install for PG 14
pig ext install pg_idkit -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_idkit;
```

