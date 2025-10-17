---
title: "pg_hashids"
linkTitle: "pg_hashids"
description: "Short unique id generator for PostgreSQL, using hashids"
weight: 4560
categories: ["Func"]
width: full
---

Short unique id generator for PostgreSQL, using hashids

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4560** | {{< badge content="pg_hashids" link="https://github.com/iCyberon/pg_hashids" >}} | {{< ext "pg_hashids" "pg_hashids" >}} | `1.3` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pg_base58" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "sequential_uuids" >}} {{< ext "permuteseq" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_hashids" >}} | `1.3` | {{< badge content="18" color="red" alt="pg_hashids_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_hashids_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_hashids" >}} | `1.3` | {{< badge content="18" color="red" alt="postgresql-18-pg-hashids" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-hashids` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_hashids_18" >}}     | {{< pkg "pg_hashids_17" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashids_17-1.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_hashids_16" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashids_16-1.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_hashids_15" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashids_15-1.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_hashids_14" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashids_14-1.3-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_hashids_18" >}}     | {{< pkg "pg_hashids_17" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashids_17-1.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_hashids_16" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashids_16-1.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_hashids_15" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashids_15-1.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_hashids_14" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashids_14-1.3-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_hashids_18" >}}     | {{< pkg "pg_hashids_17" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashids_17-1.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_hashids_16" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashids_16-1.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_hashids_15" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashids_15-1.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_hashids_14" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashids_14-1.3-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_hashids_18" >}}     | {{< pkg "pg_hashids_17" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashids_17-1.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_hashids_16" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashids_16-1.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_hashids_15" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashids_15-1.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_hashids_14" "1.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashids_14-1.3-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-hashids" >}}     | {{< pkg "postgresql-17-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-hashids" >}}     | {{< pkg "postgresql-17-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-hashids" >}}     | {{< pkg "postgresql-17-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-hashids" >}}     | {{< pkg "postgresql-17-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-hashids" >}}     | {{< pkg "postgresql-17-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-hashids" >}}     | {{< pkg "postgresql-17-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-hashids" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_hashids_17` | 1.3 | `el8.x86_64` | pigsty | 18.6 KiB | [pg_hashids_17-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashids_17-1.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_hashids_17` | 1.3 | `el8.aarch64` | pigsty | 17.9 KiB | [pg_hashids_17-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashids_17-1.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_hashids_17` | 1.3 | `el9.aarch64` | pigsty | 16.6 KiB | [pg_hashids_17-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashids_17-1.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_hashids_17` | 1.3 | `el9.x86_64` | pigsty | 17.2 KiB | [pg_hashids_17-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashids_17-1.3-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-hashids` | 1.3 | `d12.x86_64` | pigsty | 27.9 KiB | [postgresql-17-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-hashids` | 1.3 | `d12.aarch64` | pigsty | 27.3 KiB | [postgresql-17-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-hashids` | 1.3 | `u22.x86_64` | pigsty | 28.4 KiB | [postgresql-17-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-hashids` | 1.3 | `u22.aarch64` | pigsty | 27.6 KiB | [postgresql-17-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-hashids` | 1.3 | `u24.x86_64` | pigsty | 27.3 KiB | [postgresql-17-pg-hashids_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-hashids` | 1.3 | `u24.aarch64` | pigsty | 26.6 KiB | [postgresql-17-pg-hashids_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_hashids_16` | 1.3 | `el8.x86_64` | pigsty | 18.6 KiB | [pg_hashids_16-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashids_16-1.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_hashids_16` | 1.3 | `el8.aarch64` | pigsty | 17.9 KiB | [pg_hashids_16-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashids_16-1.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_hashids_16` | 1.3 | `el9.x86_64` | pigsty | 17.2 KiB | [pg_hashids_16-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashids_16-1.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_hashids_16` | 1.3 | `el9.aarch64` | pigsty | 16.6 KiB | [pg_hashids_16-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashids_16-1.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-hashids` | 1.3 | `d12.x86_64` | pigsty | 27.9 KiB | [postgresql-16-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-hashids` | 1.3 | `d12.aarch64` | pigsty | 27.3 KiB | [postgresql-16-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-hashids` | 1.3 | `u22.aarch64` | pigsty | 27.6 KiB | [postgresql-16-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-hashids` | 1.3 | `u22.x86_64` | pigsty | 28.4 KiB | [postgresql-16-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-hashids` | 1.3 | `u24.x86_64` | pigsty | 27.3 KiB | [postgresql-16-pg-hashids_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-hashids` | 1.3 | `u24.aarch64` | pigsty | 26.6 KiB | [postgresql-16-pg-hashids_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_hashids_15` | 1.3 | `el8.x86_64` | pigsty | 18.8 KiB | [pg_hashids_15-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashids_15-1.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_hashids_15` | 1.3 | `el8.aarch64` | pigsty | 18.2 KiB | [pg_hashids_15-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashids_15-1.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_hashids_15` | 1.3 | `el9.x86_64` | pigsty | 18.6 KiB | [pg_hashids_15-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashids_15-1.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_hashids_15` | 1.3 | `el9.aarch64` | pigsty | 18.0 KiB | [pg_hashids_15-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashids_15-1.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-hashids` | 1.3 | `d12.aarch64` | pigsty | 27.6 KiB | [postgresql-15-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-hashids` | 1.3 | `d12.x86_64` | pigsty | 28.2 KiB | [postgresql-15-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-hashids` | 1.3 | `u22.aarch64` | pigsty | 28.9 KiB | [postgresql-15-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-hashids` | 1.3 | `u22.x86_64` | pigsty | 29.7 KiB | [postgresql-15-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-hashids` | 1.3 | `u24.x86_64` | pigsty | 28.5 KiB | [postgresql-15-pg-hashids_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-hashids` | 1.3 | `u24.aarch64` | pigsty | 27.9 KiB | [postgresql-15-pg-hashids_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_hashids_14` | 1.3 | `el8.x86_64` | pigsty | 18.8 KiB | [pg_hashids_14-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashids_14-1.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_hashids_14` | 1.3 | `el8.aarch64` | pigsty | 18.2 KiB | [pg_hashids_14-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashids_14-1.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_hashids_14` | 1.3 | `el9.x86_64` | pigsty | 18.6 KiB | [pg_hashids_14-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashids_14-1.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_hashids_14` | 1.3 | `el9.aarch64` | pigsty | 18.0 KiB | [pg_hashids_14-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashids_14-1.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-hashids` | 1.3 | `d12.x86_64` | pigsty | 28.1 KiB | [postgresql-14-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-hashids` | 1.3 | `d12.aarch64` | pigsty | 27.5 KiB | [postgresql-14-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-hashids` | 1.3 | `u22.x86_64` | pigsty | 29.6 KiB | [postgresql-14-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-hashids` | 1.3 | `u22.aarch64` | pigsty | 28.8 KiB | [postgresql-14-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-hashids` | 1.3 | `u24.x86_64` | pigsty | 28.5 KiB | [postgresql-14-pg-hashids_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-hashids` | 1.3 | `u24.aarch64` | pigsty | 27.9 KiB | [postgresql-14-pg-hashids_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_hashids_13` | 1.3 | `el8.aarch64` | pigsty | 18.2 KiB | [pg_hashids_13-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashids_13-1.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_hashids_13` | 1.3 | `el8.x86_64` | pigsty | 18.7 KiB | [pg_hashids_13-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashids_13-1.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_hashids_13` | 1.3 | `el9.aarch64` | pigsty | 18.0 KiB | [pg_hashids_13-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashids_13-1.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_hashids_13` | 1.3 | `el9.x86_64` | pigsty | 18.6 KiB | [pg_hashids_13-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashids_13-1.3-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-hashids` | 1.3 | `d12.aarch64` | pigsty | 27.4 KiB | [postgresql-13-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-13-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-hashids` | 1.3 | `d12.x86_64` | pigsty | 28.1 KiB | [postgresql-13-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-13-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-hashids` | 1.3 | `u22.aarch64` | pigsty | 28.6 KiB | [postgresql-13-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-13-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-hashids` | 1.3 | `u22.x86_64` | pigsty | 29.6 KiB | [postgresql-13-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-13-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-hashids` | 1.3 | `u24.aarch64` | pigsty | 27.9 KiB | [postgresql-13-pg-hashids_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-13-pg-hashids_1.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-hashids` | 1.3 | `u24.x86_64` | pigsty | 28.5 KiB | [postgresql-13-pg-hashids_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-13-pg-hashids_1.3-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/iCyberon/pg_hashids" title="Repository" icon="github" subtitle="github.com/iCyberon/pg_hashids" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_hashids-1.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_hashids; # get pg_hashids source code
pig build dep pg_hashids; # install build dependencies
pig build pkg pg_hashids; # build extension rpm or deb
pig build ext pg_hashids; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_hashids; # install by extension name, for the current active PG version
pig ext install pg_hashids; # install via package alias, for the active PG version
pig ext install pg_hashids -v 18;   # install for PG 18
pig ext install pg_hashids -v 17;   # install for PG 17
pig ext install pg_hashids -v 16;   # install for PG 16
pig ext install pg_hashids -v 15;   # install for PG 15
pig ext install pg_hashids -v 14;   # install for PG 14
pig ext install pg_hashids -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_hashids;
```

