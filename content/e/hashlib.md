---
title: "hashlib"
linkTitle: "hashlib"
description: "Stable hash functions for Postgres"
weight: 4400
categories: ["Util"]
width: full
---

Stable hash functions for Postgres

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4400** | {{< badge content="hashlib" link="https://github.com/markokr/pghashlib" >}} | {{< ext "hashlib" "pg_hashlib" >}} | `1.1` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dtr" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "xxhash" >}} {{< ext "shacrypt" >}} {{< ext "cryptint" >}} {{< ext "pguecc" >}} {{< ext "pgcrypto" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/hashlib" >}} | `1.1` | {{< badge content="18" color="red" alt="pg_hashlib_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_hashlib_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/hashlib" >}} | `1.1` | {{< badge content="18" color="red" alt="postgresql-18-pg-hashlib" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-hashlib` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_hashlib_18" >}}     | {{< pkg "pg_hashlib_17" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashlib_17-1.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_hashlib_16" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashlib_16-1.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_hashlib_15" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashlib_15-1.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_hashlib_14" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashlib_14-1.1-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_hashlib_18" >}}     | {{< pkg "pg_hashlib_17" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashlib_17-1.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_hashlib_16" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashlib_16-1.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_hashlib_15" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashlib_15-1.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_hashlib_14" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashlib_14-1.1-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_hashlib_18" >}}     | {{< pkg "pg_hashlib_17" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashlib_17-1.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_hashlib_16" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashlib_16-1.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_hashlib_15" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashlib_15-1.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_hashlib_14" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashlib_14-1.1-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_hashlib_18" >}}     | {{< pkg "pg_hashlib_17" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashlib_17-1.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_hashlib_16" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashlib_16-1.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_hashlib_15" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashlib_15-1.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_hashlib_14" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashlib_14-1.1-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-hashlib" >}}     | {{< pkg "postgresql-17-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-14-pg-hashlib_1.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-hashlib" >}}     | {{< pkg "postgresql-17-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-14-pg-hashlib_1.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-hashlib" >}}     | {{< pkg "postgresql-17-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-14-pg-hashlib_1.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-hashlib" >}}     | {{< pkg "postgresql-17-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-14-pg-hashlib_1.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-hashlib" >}}     | {{< pkg "postgresql-17-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-14-pg-hashlib_1.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-hashlib" >}}     | {{< pkg "postgresql-17-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-hashlib" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-14-pg-hashlib_1.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_hashlib_17` | 1.1 | `el8.x86_64` | pigsty | 26.7 KiB | [pg_hashlib_17-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashlib_17-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_hashlib_17` | 1.1 | `el8.aarch64` | pigsty | 27.3 KiB | [pg_hashlib_17-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashlib_17-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_hashlib_17` | 1.1 | `el9.aarch64` | pigsty | 27.4 KiB | [pg_hashlib_17-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashlib_17-1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_hashlib_17` | 1.1 | `el9.x86_64` | pigsty | 27.1 KiB | [pg_hashlib_17-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashlib_17-1.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-hashlib` | 1.1 | `d12.x86_64` | pigsty | 47.1 KiB | [postgresql-17-pg-hashlib_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-hashlib` | 1.1 | `d12.aarch64` | pigsty | 47.0 KiB | [postgresql-17-pg-hashlib_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-hashlib` | 1.1 | `u22.x86_64` | pigsty | 49.9 KiB | [postgresql-17-pg-hashlib_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-hashlib` | 1.1 | `u22.aarch64` | pigsty | 49.9 KiB | [postgresql-17-pg-hashlib_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-hashlib` | 1.1 | `u24.x86_64` | pigsty | 47.9 KiB | [postgresql-17-pg-hashlib_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-hashlib` | 1.1 | `u24.aarch64` | pigsty | 48.5 KiB | [postgresql-17-pg-hashlib_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_hashlib_16` | 1.1 | `el8.x86_64` | pigsty | 26.7 KiB | [pg_hashlib_16-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashlib_16-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_hashlib_16` | 1.1 | `el8.aarch64` | pigsty | 27.3 KiB | [pg_hashlib_16-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashlib_16-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_hashlib_16` | 1.1 | `el9.x86_64` | pigsty | 27.1 KiB | [pg_hashlib_16-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashlib_16-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_hashlib_16` | 1.1 | `el9.aarch64` | pigsty | 27.4 KiB | [pg_hashlib_16-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashlib_16-1.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-hashlib` | 1.1 | `d12.x86_64` | pigsty | 47.1 KiB | [postgresql-16-pg-hashlib_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-hashlib` | 1.1 | `d12.aarch64` | pigsty | 47.0 KiB | [postgresql-16-pg-hashlib_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-hashlib` | 1.1 | `u22.aarch64` | pigsty | 50.0 KiB | [postgresql-16-pg-hashlib_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-hashlib` | 1.1 | `u22.x86_64` | pigsty | 49.9 KiB | [postgresql-16-pg-hashlib_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-hashlib` | 1.1 | `u24.x86_64` | pigsty | 47.9 KiB | [postgresql-16-pg-hashlib_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-hashlib` | 1.1 | `u24.aarch64` | pigsty | 48.5 KiB | [postgresql-16-pg-hashlib_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_hashlib_15` | 1.1 | `el8.x86_64` | pigsty | 26.9 KiB | [pg_hashlib_15-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashlib_15-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_hashlib_15` | 1.1 | `el8.aarch64` | pigsty | 27.5 KiB | [pg_hashlib_15-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashlib_15-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_hashlib_15` | 1.1 | `el9.x86_64` | pigsty | 27.3 KiB | [pg_hashlib_15-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashlib_15-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_hashlib_15` | 1.1 | `el9.aarch64` | pigsty | 27.6 KiB | [pg_hashlib_15-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashlib_15-1.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-hashlib` | 1.1 | `d12.aarch64` | pigsty | 47.5 KiB | [postgresql-15-pg-hashlib_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-hashlib` | 1.1 | `d12.x86_64` | pigsty | 47.5 KiB | [postgresql-15-pg-hashlib_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-hashlib` | 1.1 | `u22.aarch64` | pigsty | 50.6 KiB | [postgresql-15-pg-hashlib_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-hashlib` | 1.1 | `u22.x86_64` | pigsty | 50.3 KiB | [postgresql-15-pg-hashlib_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-hashlib` | 1.1 | `u24.x86_64` | pigsty | 48.0 KiB | [postgresql-15-pg-hashlib_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-hashlib` | 1.1 | `u24.aarch64` | pigsty | 48.6 KiB | [postgresql-15-pg-hashlib_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_hashlib_14` | 1.1 | `el8.x86_64` | pigsty | 26.9 KiB | [pg_hashlib_14-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashlib_14-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_hashlib_14` | 1.1 | `el8.aarch64` | pigsty | 27.5 KiB | [pg_hashlib_14-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashlib_14-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_hashlib_14` | 1.1 | `el9.x86_64` | pigsty | 27.3 KiB | [pg_hashlib_14-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashlib_14-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_hashlib_14` | 1.1 | `el9.aarch64` | pigsty | 27.6 KiB | [pg_hashlib_14-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashlib_14-1.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-hashlib` | 1.1 | `d12.x86_64` | pigsty | 47.5 KiB | [postgresql-14-pg-hashlib_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-14-pg-hashlib_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-hashlib` | 1.1 | `d12.aarch64` | pigsty | 47.5 KiB | [postgresql-14-pg-hashlib_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-14-pg-hashlib_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-hashlib` | 1.1 | `u22.x86_64` | pigsty | 50.3 KiB | [postgresql-14-pg-hashlib_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-14-pg-hashlib_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-hashlib` | 1.1 | `u22.aarch64` | pigsty | 50.6 KiB | [postgresql-14-pg-hashlib_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-14-pg-hashlib_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-hashlib` | 1.1 | `u24.x86_64` | pigsty | 48.0 KiB | [postgresql-14-pg-hashlib_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-14-pg-hashlib_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-hashlib` | 1.1 | `u24.aarch64` | pigsty | 48.5 KiB | [postgresql-14-pg-hashlib_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-14-pg-hashlib_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_hashlib_13` | 1.1 | `el8.aarch64` | pigsty | 27.5 KiB | [pg_hashlib_13-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashlib_13-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_hashlib_13` | 1.1 | `el8.x86_64` | pigsty | 26.9 KiB | [pg_hashlib_13-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashlib_13-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_hashlib_13` | 1.1 | `el9.aarch64` | pigsty | 27.6 KiB | [pg_hashlib_13-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashlib_13-1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_hashlib_13` | 1.1 | `el9.x86_64` | pigsty | 27.3 KiB | [pg_hashlib_13-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashlib_13-1.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-hashlib` | 1.1 | `d12.aarch64` | pigsty | 47.5 KiB | [postgresql-13-pg-hashlib_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-13-pg-hashlib_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-hashlib` | 1.1 | `d12.x86_64` | pigsty | 47.5 KiB | [postgresql-13-pg-hashlib_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-13-pg-hashlib_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-hashlib` | 1.1 | `u22.aarch64` | pigsty | 50.4 KiB | [postgresql-13-pg-hashlib_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-13-pg-hashlib_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-hashlib` | 1.1 | `u22.x86_64` | pigsty | 50.4 KiB | [postgresql-13-pg-hashlib_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-13-pg-hashlib_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-hashlib` | 1.1 | `u24.aarch64` | pigsty | 48.5 KiB | [postgresql-13-pg-hashlib_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-13-pg-hashlib_1.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-hashlib` | 1.1 | `u24.x86_64` | pigsty | 48.1 KiB | [postgresql-13-pg-hashlib_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-13-pg-hashlib_1.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/markokr/pghashlib" title="Repository" icon="github" subtitle="github.com/markokr/pghashlib" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_hashlib-1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get hashlib; # get hashlib source code
pig build dep hashlib; # install build dependencies
pig build pkg hashlib; # build extension rpm or deb
pig build ext hashlib; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install hashlib; # install by extension name, for the current active PG version
pig ext install pg_hashlib; # install via package alias, for the active PG version
pig ext install hashlib -v 18;   # install for PG 18
pig ext install hashlib -v 17;   # install for PG 17
pig ext install hashlib -v 16;   # install for PG 16
pig ext install hashlib -v 15;   # install for PG 15
pig ext install hashlib -v 14;   # install for PG 14
pig ext install hashlib -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION hashlib;
```

