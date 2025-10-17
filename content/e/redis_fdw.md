---
title: "redis_fdw"
linkTitle: "redis_fdw"
description: "Foreign data wrapper for querying a Redis server"
weight: 8710
categories: ["Fdw"]
width: full
---

Foreign data wrapper for querying a Redis server

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8710** | {{< badge content="redis_fdw" link="https://github.com/pg-redis-fdw/redis_fdw" >}} | {{< ext "redis_fdw" "redis_fdw" >}} | `1.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "mongo_fdw" >}} {{< ext "redis" >}} {{< ext "kafka_fdw" >}} {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "spat" >}} {{< ext "pgmemcache" >}} {{< ext "odbc_fdw" >}} |

> [!Note] multiple branch for different pg major versions


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/redis_fdw" >}} | `1.0` | {{< badge content="18" color="red" alt="redis_fdw_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `redis_fdw_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/redis_fdw" >}} | `1.0` | {{< badge content="18" color="red" alt="postgresql-18-redis-fdw" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-redis-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "redis_fdw_18" >}}     | {{< pkg "redis_fdw_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/redis_fdw_17-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "redis_fdw_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/redis_fdw_16-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "redis_fdw_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/redis_fdw_15-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "redis_fdw_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/redis_fdw_14-1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "redis_fdw_18" >}}     | {{< pkg "redis_fdw_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/redis_fdw_17-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "redis_fdw_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/redis_fdw_16-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "redis_fdw_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/redis_fdw_15-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "redis_fdw_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/redis_fdw_14-1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "redis_fdw_18" >}}     | {{< pkg "redis_fdw_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/redis_fdw_17-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "redis_fdw_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/redis_fdw_16-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "redis_fdw_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/redis_fdw_15-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "redis_fdw_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/redis_fdw_14-1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "redis_fdw_18" >}}     | {{< pkg "redis_fdw_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/redis_fdw_17-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "redis_fdw_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/redis_fdw_16-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "redis_fdw_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/redis_fdw_15-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "redis_fdw_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/redis_fdw_14-1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-redis-fdw" >}}     | {{< pkg "postgresql-17-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-redis-fdw" >}}     | {{< pkg "postgresql-17-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-redis-fdw" >}}     | {{< pkg "postgresql-17-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-redis-fdw" >}}     | {{< pkg "postgresql-17-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-redis-fdw" >}}     | {{< pkg "postgresql-17-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-redis-fdw" >}}     | {{< pkg "postgresql-17-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-redis-fdw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `redis_fdw_17` | 1.0 | `el8.x86_64` | pigsty | 26.9 KiB | [redis_fdw_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/redis_fdw_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `redis_fdw_17` | 1.0 | `el8.aarch64` | pigsty | 26.0 KiB | [redis_fdw_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/redis_fdw_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `redis_fdw_17` | 1.0 | `el9.x86_64` | pigsty | 27.1 KiB | [redis_fdw_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/redis_fdw_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `redis_fdw_17` | 1.0 | `el9.aarch64` | pigsty | 26.3 KiB | [redis_fdw_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/redis_fdw_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-redis-fdw` | 1.0 | `d12.x86_64` | pigsty | 57.1 KiB | [postgresql-17-redis-fdw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-redis-fdw` | 1.0 | `d12.aarch64` | pigsty | 55.5 KiB | [postgresql-17-redis-fdw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-redis-fdw` | 1.0 | `u22.aarch64` | pigsty | 58.5 KiB | [postgresql-17-redis-fdw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-redis-fdw` | 1.0 | `u22.x86_64` | pigsty | 59.8 KiB | [postgresql-17-redis-fdw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-redis-fdw` | 1.0 | `u24.x86_64` | pigsty | 52.6 KiB | [postgresql-17-redis-fdw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-redis-fdw` | 1.0 | `u24.aarch64` | pigsty | 51.8 KiB | [postgresql-17-redis-fdw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `redis_fdw_16` | 1.0 | `el8.aarch64` | pigsty | 26.0 KiB | [redis_fdw_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/redis_fdw_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `redis_fdw_16` | 1.0 | `el8.x86_64` | pigsty | 26.9 KiB | [redis_fdw_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/redis_fdw_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `redis_fdw_16` | 1.0 | `el9.x86_64` | pigsty | 27.2 KiB | [redis_fdw_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/redis_fdw_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `redis_fdw_16` | 1.0 | `el9.aarch64` | pigsty | 26.5 KiB | [redis_fdw_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/redis_fdw_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-redis-fdw` | 1.0 | `d12.aarch64` | pigsty | 55.2 KiB | [postgresql-16-redis-fdw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-redis-fdw` | 1.0 | `d12.x86_64` | pigsty | 56.7 KiB | [postgresql-16-redis-fdw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-redis-fdw` | 1.0 | `u22.aarch64` | pigsty | 58.5 KiB | [postgresql-16-redis-fdw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-redis-fdw` | 1.0 | `u22.x86_64` | pigsty | 59.5 KiB | [postgresql-16-redis-fdw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-redis-fdw` | 1.0 | `u24.aarch64` | pigsty | 51.8 KiB | [postgresql-16-redis-fdw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-redis-fdw` | 1.0 | `u24.x86_64` | pigsty | 52.8 KiB | [postgresql-16-redis-fdw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `redis_fdw_15` | 1.0 | `el8.x86_64` | pigsty | 26.8 KiB | [redis_fdw_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/redis_fdw_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `redis_fdw_15` | 1.0 | `el8.aarch64` | pigsty | 26.0 KiB | [redis_fdw_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/redis_fdw_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `redis_fdw_15` | 1.0 | `el9.x86_64` | pigsty | 27.1 KiB | [redis_fdw_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/redis_fdw_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `redis_fdw_15` | 1.0 | `el9.aarch64` | pigsty | 26.4 KiB | [redis_fdw_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/redis_fdw_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-redis-fdw` | 1.0 | `d12.x86_64` | pigsty | 56.3 KiB | [postgresql-15-redis-fdw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-redis-fdw` | 1.0 | `d12.aarch64` | pigsty | 54.9 KiB | [postgresql-15-redis-fdw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-redis-fdw` | 1.0 | `u22.aarch64` | pigsty | 58.2 KiB | [postgresql-15-redis-fdw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-redis-fdw` | 1.0 | `u22.x86_64` | pigsty | 59.2 KiB | [postgresql-15-redis-fdw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-redis-fdw` | 1.0 | `u24.aarch64` | pigsty | 51.7 KiB | [postgresql-15-redis-fdw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-redis-fdw` | 1.0 | `u24.x86_64` | pigsty | 52.7 KiB | [postgresql-15-redis-fdw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `redis_fdw_14` | 1.0 | `el8.aarch64` | pigsty | 26.0 KiB | [redis_fdw_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/redis_fdw_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `redis_fdw_14` | 1.0 | `el8.x86_64` | pigsty | 26.8 KiB | [redis_fdw_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/redis_fdw_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `redis_fdw_14` | 1.0 | `el9.x86_64` | pigsty | 27.1 KiB | [redis_fdw_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/redis_fdw_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `redis_fdw_14` | 1.0 | `el9.aarch64` | pigsty | 26.4 KiB | [redis_fdw_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/redis_fdw_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-redis-fdw` | 1.0 | `d12.x86_64` | pigsty | 56.4 KiB | [postgresql-14-redis-fdw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-redis-fdw` | 1.0 | `d12.aarch64` | pigsty | 54.9 KiB | [postgresql-14-redis-fdw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-redis-fdw` | 1.0 | `u22.aarch64` | pigsty | 58.2 KiB | [postgresql-14-redis-fdw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-redis-fdw` | 1.0 | `u22.x86_64` | pigsty | 59.3 KiB | [postgresql-14-redis-fdw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-redis-fdw` | 1.0 | `u24.aarch64` | pigsty | 51.7 KiB | [postgresql-14-redis-fdw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-redis-fdw` | 1.0 | `u24.x86_64` | pigsty | 52.7 KiB | [postgresql-14-redis-fdw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `redis_fdw_13` | 1.1 | `el8.x86_64` | pgdg | 198.5 KiB | [redis_fdw_13-1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/redis_fdw_13-1.1-1.rhel8.x86_64.rpm) |
| `redis_fdw_13` | 1.0 | `el8.x86_64` | pigsty | 26.6 KiB | [redis_fdw_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/redis_fdw_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `redis_fdw_13` | 1.0 | `el8.aarch64` | pigsty | 25.9 KiB | [redis_fdw_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/redis_fdw_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `redis_fdw_13` | 1.0 | `el9.aarch64` | pigsty | 26.5 KiB | [redis_fdw_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/redis_fdw_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `redis_fdw_13` | 1.0 | `el9.x86_64` | pigsty | 27.2 KiB | [redis_fdw_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/redis_fdw_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-redis-fdw` | 1.0 | `d12.x86_64` | pigsty | 56.4 KiB | [postgresql-13-redis-fdw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-13-redis-fdw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-redis-fdw` | 1.0 | `d12.aarch64` | pigsty | 54.8 KiB | [postgresql-13-redis-fdw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-13-redis-fdw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-redis-fdw` | 1.0 | `u22.x86_64` | pigsty | 59.0 KiB | [postgresql-13-redis-fdw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-13-redis-fdw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-redis-fdw` | 1.0 | `u22.aarch64` | pigsty | 57.8 KiB | [postgresql-13-redis-fdw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-13-redis-fdw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-redis-fdw` | 1.0 | `u24.aarch64` | pigsty | 51.9 KiB | [postgresql-13-redis-fdw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-13-redis-fdw_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-redis-fdw` | 1.0 | `u24.x86_64` | pigsty | 52.7 KiB | [postgresql-13-redis-fdw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-13-redis-fdw_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pg-redis-fdw/redis_fdw" title="Repository" icon="github" subtitle="github.com/pg-redis-fdw/redis_fdw" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="redis_fdw-1.0-18.tar.gz" >}}
{{< /cards >}}


```bash
pig build get redis_fdw; # get redis_fdw source code
pig build dep redis_fdw; # install build dependencies
pig build pkg redis_fdw; # build extension rpm or deb
pig build ext redis_fdw; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install redis_fdw; # install by extension name, for the current active PG version
pig ext install redis_fdw; # install via package alias, for the active PG version
pig ext install redis_fdw -v 18;   # install for PG 18
pig ext install redis_fdw -v 17;   # install for PG 17
pig ext install redis_fdw -v 16;   # install for PG 16
pig ext install redis_fdw -v 15;   # install for PG 15
pig ext install redis_fdw -v 14;   # install for PG 14
pig ext install redis_fdw -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION redis_fdw;
```

