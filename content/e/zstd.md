---
title: "zstd"
linkTitle: "zstd"
description: "Zstandard compression algorithm implementation in PostgreSQL"
weight: 4030
categories: ["Util"]
width: full
---

Zstandard compression algorithm implementation in PostgreSQL

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4030** | {{< badge content="zstd" link="https://github.com/grahamedgecombe/pgzstd" >}} | {{< ext "zstd" "pg_zstd" >}} | `1.1.2` | {{< category "UTIL" >}} | {{< license "ISC" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} {{< ext "pgjq" >}} {{< ext "pgjwt" >}} {{< ext "pg_smtp_client" >}} |

> [!Note] +varatt.h


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/zstd" >}} | `1.1.2` | {{< badge content="18" color="red" alt="pg_zstd_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_zstd_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/zstd" >}} | `1.1.2` | {{< badge content="18" color="red" alt="postgresql-18-zstd" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-zstd` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_zstd_18" >}}     | {{< pkg "pg_zstd_17" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_zstd_17-1.1.2-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_zstd_16" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_zstd_16-1.1.2-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_zstd_15" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_zstd_15-1.1.2-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_zstd_14" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_zstd_14-1.1.2-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_zstd_18" >}}     | {{< pkg "pg_zstd_17" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_zstd_17-1.1.2-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_zstd_16" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_zstd_16-1.1.2-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_zstd_15" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_zstd_15-1.1.2-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_zstd_14" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_zstd_14-1.1.2-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_zstd_18" >}}     | {{< pkg "pg_zstd_17" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_zstd_17-1.1.2-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_zstd_16" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_zstd_16-1.1.2-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_zstd_15" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_zstd_15-1.1.2-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_zstd_14" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_zstd_14-1.1.2-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_zstd_18" >}}     | {{< pkg "pg_zstd_17" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_zstd_17-1.1.2-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_zstd_16" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_zstd_16-1.1.2-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_zstd_15" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_zstd_15-1.1.2-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_zstd_14" "1.1.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_zstd_14-1.1.2-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-zstd" >}}     | {{< pkg "postgresql-17-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-zstd" >}}     | {{< pkg "postgresql-17-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-zstd" >}}     | {{< pkg "postgresql-17-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-zstd" >}}     | {{< pkg "postgresql-17-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-zstd" >}}     | {{< pkg "postgresql-17-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-zstd" >}}     | {{< pkg "postgresql-17-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-zstd" "1.1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_zstd_17` | 1.1.2 | `el8.x86_64` | pigsty | 12.3 KiB | [pg_zstd_17-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_zstd_17-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_zstd_17` | 1.1.2 | `el8.aarch64` | pigsty | 12.1 KiB | [pg_zstd_17-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_zstd_17-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_zstd_17` | 1.1.2 | `el9.aarch64` | pigsty | 11.9 KiB | [pg_zstd_17-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_zstd_17-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_zstd_17` | 1.1.2 | `el9.x86_64` | pigsty | 12.2 KiB | [pg_zstd_17-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_zstd_17-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-zstd` | 1.1.2 | `d12.x86_64` | pigsty | 11.9 KiB | [postgresql-17-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-zstd` | 1.1.2 | `d12.aarch64` | pigsty | 11.8 KiB | [postgresql-17-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-zstd` | 1.1.2 | `u22.x86_64` | pigsty | 12.7 KiB | [postgresql-17-zstd_1.1.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-zstd` | 1.1.2 | `u22.aarch64` | pigsty | 12.5 KiB | [postgresql-17-zstd_1.1.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-zstd` | 1.1.2 | `u24.x86_64` | pigsty | 12.4 KiB | [postgresql-17-zstd_1.1.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-zstd` | 1.1.2 | `u24.aarch64` | pigsty | 12.3 KiB | [postgresql-17-zstd_1.1.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_zstd_16` | 1.1.2 | `el8.x86_64` | pigsty | 12.3 KiB | [pg_zstd_16-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_zstd_16-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_zstd_16` | 1.1.2 | `el8.aarch64` | pigsty | 12.1 KiB | [pg_zstd_16-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_zstd_16-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_zstd_16` | 1.1.2 | `el9.x86_64` | pigsty | 12.2 KiB | [pg_zstd_16-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_zstd_16-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_zstd_16` | 1.1.2 | `el9.aarch64` | pigsty | 11.9 KiB | [pg_zstd_16-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_zstd_16-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-zstd` | 1.1.2 | `d12.x86_64` | pigsty | 11.9 KiB | [postgresql-16-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-zstd` | 1.1.2 | `d12.aarch64` | pigsty | 11.8 KiB | [postgresql-16-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-zstd` | 1.1.2 | `u22.aarch64` | pigsty | 12.5 KiB | [postgresql-16-zstd_1.1.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-zstd` | 1.1.2 | `u22.x86_64` | pigsty | 12.7 KiB | [postgresql-16-zstd_1.1.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-zstd` | 1.1.2 | `u24.x86_64` | pigsty | 12.4 KiB | [postgresql-16-zstd_1.1.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-zstd` | 1.1.2 | `u24.aarch64` | pigsty | 12.3 KiB | [postgresql-16-zstd_1.1.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_zstd_15` | 1.1.2 | `el8.x86_64` | pigsty | 12.3 KiB | [pg_zstd_15-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_zstd_15-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_zstd_15` | 1.1.2 | `el8.aarch64` | pigsty | 12.1 KiB | [pg_zstd_15-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_zstd_15-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_zstd_15` | 1.1.2 | `el9.x86_64` | pigsty | 12.2 KiB | [pg_zstd_15-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_zstd_15-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_zstd_15` | 1.1.2 | `el9.aarch64` | pigsty | 11.9 KiB | [pg_zstd_15-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_zstd_15-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-zstd` | 1.1.2 | `d12.aarch64` | pigsty | 11.8 KiB | [postgresql-15-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-zstd` | 1.1.2 | `d12.x86_64` | pigsty | 11.9 KiB | [postgresql-15-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-zstd` | 1.1.2 | `u22.aarch64` | pigsty | 12.5 KiB | [postgresql-15-zstd_1.1.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-zstd` | 1.1.2 | `u22.x86_64` | pigsty | 12.7 KiB | [postgresql-15-zstd_1.1.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-zstd` | 1.1.2 | `u24.x86_64` | pigsty | 12.4 KiB | [postgresql-15-zstd_1.1.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-zstd` | 1.1.2 | `u24.aarch64` | pigsty | 12.2 KiB | [postgresql-15-zstd_1.1.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_zstd_14` | 1.1.2 | `el8.x86_64` | pigsty | 12.3 KiB | [pg_zstd_14-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_zstd_14-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_zstd_14` | 1.1.2 | `el8.aarch64` | pigsty | 12.1 KiB | [pg_zstd_14-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_zstd_14-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_zstd_14` | 1.1.2 | `el9.x86_64` | pigsty | 12.2 KiB | [pg_zstd_14-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_zstd_14-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_zstd_14` | 1.1.2 | `el9.aarch64` | pigsty | 11.8 KiB | [pg_zstd_14-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_zstd_14-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-zstd` | 1.1.2 | `d12.x86_64` | pigsty | 11.9 KiB | [postgresql-14-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-zstd` | 1.1.2 | `d12.aarch64` | pigsty | 11.8 KiB | [postgresql-14-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-zstd` | 1.1.2 | `u22.x86_64` | pigsty | 12.6 KiB | [postgresql-14-zstd_1.1.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-zstd` | 1.1.2 | `u22.aarch64` | pigsty | 12.4 KiB | [postgresql-14-zstd_1.1.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-zstd` | 1.1.2 | `u24.x86_64` | pigsty | 12.4 KiB | [postgresql-14-zstd_1.1.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-zstd` | 1.1.2 | `u24.aarch64` | pigsty | 12.2 KiB | [postgresql-14-zstd_1.1.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_zstd_13` | 1.1.2 | `el8.aarch64` | pigsty | 12.1 KiB | [pg_zstd_13-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_zstd_13-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_zstd_13` | 1.1.2 | `el8.x86_64` | pigsty | 12.2 KiB | [pg_zstd_13-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_zstd_13-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_zstd_13` | 1.1.2 | `el9.aarch64` | pigsty | 11.8 KiB | [pg_zstd_13-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_zstd_13-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_zstd_13` | 1.1.2 | `el9.x86_64` | pigsty | 12.2 KiB | [pg_zstd_13-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_zstd_13-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-zstd` | 1.1.2 | `d12.aarch64` | pigsty | 11.8 KiB | [postgresql-13-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-13-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-zstd` | 1.1.2 | `d12.x86_64` | pigsty | 11.8 KiB | [postgresql-13-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-13-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-zstd` | 1.1.2 | `u22.aarch64` | pigsty | 12.4 KiB | [postgresql-13-zstd_1.1.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-13-zstd_1.1.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-13-zstd` | 1.1.2 | `u22.x86_64` | pigsty | 12.4 KiB | [postgresql-13-zstd_1.1.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-13-zstd_1.1.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-13-zstd` | 1.1.2 | `u24.aarch64` | pigsty | 12.2 KiB | [postgresql-13-zstd_1.1.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-13-zstd_1.1.2-2PIGSTY~noble_arm64.deb) |
| `postgresql-13-zstd` | 1.1.2 | `u24.x86_64` | pigsty | 12.3 KiB | [postgresql-13-zstd_1.1.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-13-zstd_1.1.2-2PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/grahamedgecombe/pgzstd" title="Repository" icon="github" subtitle="github.com/grahamedgecombe/pgzstd" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgzstd-1.1.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get zstd; # get zstd source code
pig build dep zstd; # install build dependencies
pig build pkg zstd; # build extension rpm or deb
pig build ext zstd; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install zstd; # install by extension name, for the current active PG version
pig ext install pg_zstd; # install via package alias, for the active PG version
pig ext install zstd -v 18;   # install for PG 18
pig ext install zstd -v 17;   # install for PG 17
pig ext install zstd -v 16;   # install for PG 16
pig ext install zstd -v 15;   # install for PG 15
pig ext install zstd -v 14;   # install for PG 14
pig ext install zstd -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION zstd;
```




--------

## Usage

| Function                                                                             | Return Type |
|--------------------------------------------------------------------------------------|-------------|
| <code>zstd_compress(*data* bytea [, *dictionary* bytea [, *level* integer ]])</code> | `bytea`     |
| <code>zstd_decompress(*data* bytea [, *dictionary* bytea ])</code>                   | `bytea`     |
| <code>zstd_length(*data* bytea)</code>                                               | `integer`   |

`zstd_compress` compresses the provided `data` and returns a Zstandard frame. A
preset `dictionary` may also be provided. The default compression `level` may
also be overriden, valid values range from `1` (best speed) to `22` (best
compression). The default level is `3`.

If you want to override the compression level without using a dictionary, set
`dictionary` to `NULL`.

`zstd_decompress` decompresses the provided Zstandard frame in `data` and
returns the uncompressed data. A preset `dictionary`, matching the dictionary
used to compress the data, may also be provided.

`zstd_length` returns the decompressed length of the provided Zstandard frame.
If `ZSTD_getFrameContentSize()` is available it returns `NULL` if the length is
unknown. If unavailable, it isn't possible to distinguish the error, unknown
decompressed length and zero decompressed length cases.


### Example

Basic compress/decompress example:

```sql
CREATE EXTENSION zstd;

SELECT zstd_compress('hello world');
-- zstd_compress
-- --------------------------------------------
-- \x28b52ffd200b59000068656c6c6f20776f726c64

SELECT convert_from(zstd_decompress('\x28b52ffd200b59000068656c6c6f20776f726c64'), 'utf-8');
-- convert_from
-- --------------
--  hello world
```

Compress with level (`1` for best speed, `22` for best compression, default to `3`)

```sql
CREATE EXTENSION zstd;

SELECT zstd_compress('hello world',  NULL, 10);
-- zstd_compress
-- --------------------------------------------
-- \x28b52ffd200b59000068656c6c6f20776f726c64

SELECT convert_from(zstd_decompress('\x28b52ffd200b59000068656c6c6f20776f726c64'), 'utf-8');
-- convert_from
-- --------------
--  hello world
```
