---
title: "xxhash"
linkTitle: "xxhash"
description: "xxhash functions for PostgreSQL"
weight: 4430
categories: ["Util"]
width: full
---

xxhash functions for PostgreSQL

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4430** | {{< badge content="xxhash" link="https://github.com/hatarist/pg_xxhash" >}} | {{< ext "xxhash" "pg_xxhash" >}} | `0.0.1` | {{< category "UTIL" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dtr" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hashlib" >}} {{< ext "shacrypt" >}} {{< ext "pgcrypto" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/xxhash" >}} | `0.0.1` | {{< badge content="18" color="red" alt="pg_xxhash_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_xxhash_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/xxhash" >}} | `0.0.1` | {{< badge content="18" color="red" alt="postgresql-18-pg-xxhash" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-xxhash` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_xxhash_18" >}}     | {{< pkg "pg_xxhash_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xxhash_17-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_xxhash_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xxhash_16-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_xxhash_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xxhash_15-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_xxhash_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xxhash_14-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_xxhash_18" >}}     | {{< pkg "pg_xxhash_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xxhash_17-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_xxhash_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xxhash_16-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_xxhash_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xxhash_15-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_xxhash_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xxhash_14-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_xxhash_18" >}}     | {{< pkg "pg_xxhash_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xxhash_17-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_xxhash_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xxhash_16-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_xxhash_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xxhash_15-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_xxhash_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xxhash_14-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_xxhash_18" >}}     | {{< pkg "pg_xxhash_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xxhash_17-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_xxhash_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xxhash_16-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_xxhash_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xxhash_15-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_xxhash_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xxhash_14-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-xxhash" >}}     | {{< pkg "postgresql-17-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-xxhash" >}}     | {{< pkg "postgresql-17-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-xxhash" >}}     | {{< pkg "postgresql-17-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-xxhash" >}}     | {{< pkg "postgresql-17-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-xxhash" >}}     | {{< pkg "postgresql-17-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-xxhash" >}}     | {{< pkg "postgresql-17-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-xxhash" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_xxhash_17` | 0.0.1 | `el8.x86_64` | pigsty | 30.0 KiB | [pg_xxhash_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xxhash_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_xxhash_17` | 0.0.1 | `el8.aarch64` | pigsty | 29.8 KiB | [pg_xxhash_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xxhash_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_xxhash_17` | 0.0.1 | `el9.aarch64` | pigsty | 18.8 KiB | [pg_xxhash_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xxhash_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_xxhash_17` | 0.0.1 | `el9.x86_64` | pigsty | 18.5 KiB | [pg_xxhash_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xxhash_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-xxhash` | 0.0.1 | `d12.x86_64` | pigsty | 79.8 KiB | [postgresql-17-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-xxhash` | 0.0.1 | `d12.aarch64` | pigsty | 82.2 KiB | [postgresql-17-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-xxhash` | 0.0.1 | `u22.x86_64` | pigsty | 77.0 KiB | [postgresql-17-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-xxhash` | 0.0.1 | `u22.aarch64` | pigsty | 78.8 KiB | [postgresql-17-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-xxhash` | 0.0.1 | `u24.x86_64` | pigsty | 71.8 KiB | [postgresql-17-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-xxhash` | 0.0.1 | `u24.aarch64` | pigsty | 75.5 KiB | [postgresql-17-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_xxhash_16` | 0.0.1 | `el8.x86_64` | pigsty | 30.0 KiB | [pg_xxhash_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xxhash_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_xxhash_16` | 0.0.1 | `el8.aarch64` | pigsty | 29.8 KiB | [pg_xxhash_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xxhash_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_xxhash_16` | 0.0.1 | `el9.x86_64` | pigsty | 18.5 KiB | [pg_xxhash_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xxhash_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_xxhash_16` | 0.0.1 | `el9.aarch64` | pigsty | 18.8 KiB | [pg_xxhash_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xxhash_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-xxhash` | 0.0.1 | `d12.x86_64` | pigsty | 79.8 KiB | [postgresql-16-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-xxhash` | 0.0.1 | `d12.aarch64` | pigsty | 82.1 KiB | [postgresql-16-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-xxhash` | 0.0.1 | `u22.aarch64` | pigsty | 78.8 KiB | [postgresql-16-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-xxhash` | 0.0.1 | `u22.x86_64` | pigsty | 77.0 KiB | [postgresql-16-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-xxhash` | 0.0.1 | `u24.x86_64` | pigsty | 71.8 KiB | [postgresql-16-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-xxhash` | 0.0.1 | `u24.aarch64` | pigsty | 75.5 KiB | [postgresql-16-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_xxhash_15` | 0.0.1 | `el8.x86_64` | pigsty | 30.7 KiB | [pg_xxhash_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xxhash_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_xxhash_15` | 0.0.1 | `el8.aarch64` | pigsty | 30.4 KiB | [pg_xxhash_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xxhash_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_xxhash_15` | 0.0.1 | `el9.x86_64` | pigsty | 29.2 KiB | [pg_xxhash_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xxhash_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_xxhash_15` | 0.0.1 | `el9.aarch64` | pigsty | 30.6 KiB | [pg_xxhash_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xxhash_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-xxhash` | 0.0.1 | `d12.aarch64` | pigsty | 83.0 KiB | [postgresql-15-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-xxhash` | 0.0.1 | `d12.x86_64` | pigsty | 80.5 KiB | [postgresql-15-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-xxhash` | 0.0.1 | `u22.aarch64` | pigsty | 89.8 KiB | [postgresql-15-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-xxhash` | 0.0.1 | `u22.x86_64` | pigsty | 87.0 KiB | [postgresql-15-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-xxhash` | 0.0.1 | `u24.x86_64` | pigsty | 82.4 KiB | [postgresql-15-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-xxhash` | 0.0.1 | `u24.aarch64` | pigsty | 86.6 KiB | [postgresql-15-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_xxhash_14` | 0.0.1 | `el8.x86_64` | pigsty | 30.7 KiB | [pg_xxhash_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xxhash_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_xxhash_14` | 0.0.1 | `el8.aarch64` | pigsty | 30.4 KiB | [pg_xxhash_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xxhash_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_xxhash_14` | 0.0.1 | `el9.x86_64` | pigsty | 29.2 KiB | [pg_xxhash_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xxhash_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_xxhash_14` | 0.0.1 | `el9.aarch64` | pigsty | 30.6 KiB | [pg_xxhash_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xxhash_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-xxhash` | 0.0.1 | `d12.x86_64` | pigsty | 80.7 KiB | [postgresql-14-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-xxhash` | 0.0.1 | `d12.aarch64` | pigsty | 82.9 KiB | [postgresql-14-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-xxhash` | 0.0.1 | `u22.x86_64` | pigsty | 87.2 KiB | [postgresql-14-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-xxhash` | 0.0.1 | `u22.aarch64` | pigsty | 89.6 KiB | [postgresql-14-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-xxhash` | 0.0.1 | `u24.x86_64` | pigsty | 82.4 KiB | [postgresql-14-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-xxhash` | 0.0.1 | `u24.aarch64` | pigsty | 86.6 KiB | [postgresql-14-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_xxhash_13` | 0.0.1 | `el8.aarch64` | pigsty | 30.4 KiB | [pg_xxhash_13-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xxhash_13-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_xxhash_13` | 0.0.1 | `el8.x86_64` | pigsty | 30.7 KiB | [pg_xxhash_13-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xxhash_13-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_xxhash_13` | 0.0.1 | `el9.aarch64` | pigsty | 30.5 KiB | [pg_xxhash_13-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xxhash_13-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_xxhash_13` | 0.0.1 | `el9.x86_64` | pigsty | 29.2 KiB | [pg_xxhash_13-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xxhash_13-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-xxhash` | 0.0.1 | `d12.aarch64` | pigsty | 82.9 KiB | [postgresql-13-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-13-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-xxhash` | 0.0.1 | `d12.x86_64` | pigsty | 80.8 KiB | [postgresql-13-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-13-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-xxhash` | 0.0.1 | `u22.aarch64` | pigsty | 89.6 KiB | [postgresql-13-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-13-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-xxhash` | 0.0.1 | `u22.x86_64` | pigsty | 87.2 KiB | [postgresql-13-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-13-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-xxhash` | 0.0.1 | `u24.aarch64` | pigsty | 86.6 KiB | [postgresql-13-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-13-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-xxhash` | 0.0.1 | `u24.x86_64` | pigsty | 82.4 KiB | [postgresql-13-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-13-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/hatarist/pg_xxhash" title="Repository" icon="github" subtitle="github.com/hatarist/pg_xxhash" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_xxhash-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get xxhash; # get xxhash source code
pig build dep xxhash; # install build dependencies
pig build pkg xxhash; # build extension rpm or deb
pig build ext xxhash; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install xxhash; # install by extension name, for the current active PG version
pig ext install pg_xxhash; # install via package alias, for the active PG version
pig ext install xxhash -v 18;   # install for PG 18
pig ext install xxhash -v 17;   # install for PG 17
pig ext install xxhash -v 16;   # install for PG 16
pig ext install xxhash -v 15;   # install for PG 15
pig ext install xxhash -v 14;   # install for PG 14
pig ext install xxhash -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION xxhash;
```

