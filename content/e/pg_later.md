---
title: "pg_later"
linkTitle: "pg_later"
description: "Run queries now and get results later"
weight: 1090
categories: ["Time"]
width: full
---

Run queries now and get results later

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1090** | {{< badge content="pg_later" link="https://github.com/ChuckHend/pg_later" >}} | {{< ext "pg_later" "pg_later" >}} | `0.3.0` | {{< category "TIME" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pgmq" >}} |
|   **See Also**    | {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pg_background" >}} {{< ext "timescaledb" >}} {{< ext "timescaledb_toolkit" >}} {{< ext "timeseries" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} |

> [!Note] pgrx=0.12.5


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_later" >}} | `0.3.0` | {{< badge content="18" color="red" alt="pg_later_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_later_$v` | `pgmq_$v` |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_later" >}} | `0.3.0` | {{< badge content="18" color="red" alt="postgresql-18-pg-later" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-later` | `postgresql-$v-pgmq` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_later_18" >}}     | {{< pkg "pg_later_17" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_later_17-0.3.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_later_16" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_later_16-0.3.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_later_15" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_later_15-0.3.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_later_14" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_later_14-0.3.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_later_18" >}}     | {{< pkg "pg_later_17" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_later_17-0.3.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_later_16" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_later_16-0.3.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_later_15" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_later_15-0.3.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_later_14" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_later_14-0.3.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_later_18" >}}     | {{< pkg "pg_later_17" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_later_17-0.3.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_later_16" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_later_16-0.3.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_later_15" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_later_15-0.3.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_later_14" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_later_14-0.3.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_later_18" >}}     | {{< pkg "pg_later_17" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_17-0.2.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_later_16" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_16-0.2.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_later_15" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_15-0.2.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_later_14" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_14-0.2.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-later" >}}     | {{< pkg "postgresql-17-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-17-pg-later_0.3.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-16-pg-later_0.3.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-15-pg-later_0.3.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-14-pg-later_0.3.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-later" >}}     | {{< pkg "postgresql-17-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-17-pg-later_0.3.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-16-pg-later_0.3.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-15-pg-later_0.3.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-14-pg-later_0.3.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-later" >}}     | {{< pkg "postgresql-17-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-17-pg-later_0.3.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-16-pg-later_0.3.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-15-pg-later_0.3.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-14-pg-later_0.3.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-later" >}}     | {{< pkg "postgresql-17-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-17-pg-later_0.3.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-16-pg-later_0.3.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-15-pg-later_0.3.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-14-pg-later_0.3.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-later" >}}     | {{< pkg "postgresql-17-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-17-pg-later_0.3.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-16-pg-later_0.3.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-15-pg-later_0.3.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-14-pg-later_0.3.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-later" >}}     | {{< pkg "postgresql-17-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-17-pg-later_0.3.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-16-pg-later_0.3.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-15-pg-later_0.3.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-later" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-14-pg-later_0.3.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_later_17` | 0.3.0 | `el8.aarch64` | pigsty | 1.2 MiB | [pg_later_17-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_later_17-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_later_17` | 0.3.0 | `el8.x86_64` | pigsty | 1.3 MiB | [pg_later_17-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_later_17-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_later_17` | 0.3.0 | `el9.x86_64` | pigsty | 1.3 MiB | [pg_later_17-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_later_17-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_later_17` | 0.3.0 | `el9.aarch64` | pigsty | 1.3 MiB | [pg_later_17-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_17-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_17` | 0.2.0 | `el9.aarch64` | pigsty | 1.3 MiB | [pg_later_17-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_17-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-later` | 0.3.0 | `d12.aarch64` | pigsty | 973.4 KiB | [postgresql-17-pg-later_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-17-pg-later_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-later` | 0.3.0 | `d12.x86_64` | pigsty | 1.1 MiB | [postgresql-17-pg-later_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-17-pg-later_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-later` | 0.3.0 | `u22.x86_64` | pigsty | 1.2 MiB | [postgresql-17-pg-later_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-17-pg-later_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-later` | 0.3.0 | `u22.aarch64` | pigsty | 1.1 MiB | [postgresql-17-pg-later_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-17-pg-later_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-later` | 0.3.0 | `u24.aarch64` | pigsty | 1.1 MiB | [postgresql-17-pg-later_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-17-pg-later_0.3.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-later` | 0.3.0 | `u24.x86_64` | pigsty | 1.2 MiB | [postgresql-17-pg-later_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-17-pg-later_0.3.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_later_16` | 0.3.0 | `el8.x86_64` | pigsty | 1.3 MiB | [pg_later_16-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_later_16-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_later_16` | 0.3.0 | `el8.aarch64` | pigsty | 1.2 MiB | [pg_later_16-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_later_16-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_later_16` | 0.3.0 | `el9.aarch64` | pigsty | 1.3 MiB | [pg_later_16-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_16-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_16` | 0.3.0 | `el9.x86_64` | pigsty | 1.3 MiB | [pg_later_16-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_later_16-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_later_16` | 0.2.0 | `el9.aarch64` | pigsty | 1.3 MiB | [pg_later_16-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_16-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-later` | 0.3.0 | `d12.x86_64` | pigsty | 1.1 MiB | [postgresql-16-pg-later_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-16-pg-later_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-later` | 0.3.0 | `d12.aarch64` | pigsty | 973.2 KiB | [postgresql-16-pg-later_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-16-pg-later_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-later` | 0.3.0 | `u22.aarch64` | pigsty | 1.1 MiB | [postgresql-16-pg-later_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-16-pg-later_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-later` | 0.3.0 | `u22.x86_64` | pigsty | 1.2 MiB | [postgresql-16-pg-later_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-16-pg-later_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-later` | 0.3.0 | `u24.aarch64` | pigsty | 1.1 MiB | [postgresql-16-pg-later_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-16-pg-later_0.3.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-later` | 0.3.0 | `u24.x86_64` | pigsty | 1.2 MiB | [postgresql-16-pg-later_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-16-pg-later_0.3.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_later_15` | 0.3.0 | `el8.x86_64` | pigsty | 1.3 MiB | [pg_later_15-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_later_15-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_later_15` | 0.3.0 | `el8.aarch64` | pigsty | 1.2 MiB | [pg_later_15-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_later_15-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_later_15` | 0.3.0 | `el9.x86_64` | pigsty | 1.3 MiB | [pg_later_15-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_later_15-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_later_15` | 0.3.0 | `el9.aarch64` | pigsty | 1.3 MiB | [pg_later_15-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_15-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_15` | 0.2.0 | `el9.aarch64` | pigsty | 1.3 MiB | [pg_later_15-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_15-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-later` | 0.3.0 | `d12.aarch64` | pigsty | 973.4 KiB | [postgresql-15-pg-later_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-15-pg-later_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-later` | 0.3.0 | `d12.x86_64` | pigsty | 1.1 MiB | [postgresql-15-pg-later_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-15-pg-later_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-later` | 0.3.0 | `u22.aarch64` | pigsty | 1.1 MiB | [postgresql-15-pg-later_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-15-pg-later_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-later` | 0.3.0 | `u22.x86_64` | pigsty | 1.2 MiB | [postgresql-15-pg-later_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-15-pg-later_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-later` | 0.3.0 | `u24.aarch64` | pigsty | 1.1 MiB | [postgresql-15-pg-later_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-15-pg-later_0.3.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-later` | 0.3.0 | `u24.x86_64` | pigsty | 1.2 MiB | [postgresql-15-pg-later_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-15-pg-later_0.3.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_later_14` | 0.3.0 | `el8.x86_64` | pigsty | 1.3 MiB | [pg_later_14-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_later_14-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_later_14` | 0.3.0 | `el8.aarch64` | pigsty | 1.2 MiB | [pg_later_14-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_later_14-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_later_14` | 0.3.0 | `el9.x86_64` | pigsty | 1.3 MiB | [pg_later_14-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_later_14-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_later_14` | 0.3.0 | `el9.aarch64` | pigsty | 1.3 MiB | [pg_later_14-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_14-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_14` | 0.2.0 | `el9.aarch64` | pigsty | 1.3 MiB | [pg_later_14-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_14-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-later` | 0.3.0 | `d12.x86_64` | pigsty | 1.1 MiB | [postgresql-14-pg-later_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-14-pg-later_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-later` | 0.3.0 | `d12.aarch64` | pigsty | 973.2 KiB | [postgresql-14-pg-later_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-14-pg-later_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-later` | 0.3.0 | `u22.x86_64` | pigsty | 1.2 MiB | [postgresql-14-pg-later_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-14-pg-later_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-later` | 0.3.0 | `u22.aarch64` | pigsty | 1.1 MiB | [postgresql-14-pg-later_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-14-pg-later_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-later` | 0.3.0 | `u24.x86_64` | pigsty | 1.2 MiB | [postgresql-14-pg-later_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-14-pg-later_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-later` | 0.3.0 | `u24.aarch64` | pigsty | 1.1 MiB | [postgresql-14-pg-later_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-14-pg-later_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_later_13` | 0.3.0 | `el8.aarch64` | pigsty | 1.2 MiB | [pg_later_13-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_later_13-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_later_13` | 0.3.0 | `el8.x86_64` | pigsty | 1.3 MiB | [pg_later_13-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_later_13-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_later_13` | 0.3.0 | `el9.x86_64` | pigsty | 1.3 MiB | [pg_later_13-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_later_13-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_later_13` | 0.3.0 | `el9.aarch64` | pigsty | 1.3 MiB | [pg_later_13-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_13-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_13` | 0.2.0 | `el9.aarch64` | pigsty | 1.3 MiB | [pg_later_13-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_13-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pg-later` | 0.3.0 | `d12.aarch64` | pigsty | 973.6 KiB | [postgresql-13-pg-later_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-13-pg-later_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-later` | 0.3.0 | `d12.x86_64` | pigsty | 1.1 MiB | [postgresql-13-pg-later_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-13-pg-later_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-later` | 0.3.0 | `u22.x86_64` | pigsty | 1.2 MiB | [postgresql-13-pg-later_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-13-pg-later_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-later` | 0.3.0 | `u22.aarch64` | pigsty | 1.1 MiB | [postgresql-13-pg-later_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-13-pg-later_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-later` | 0.3.0 | `u24.x86_64` | pigsty | 1.2 MiB | [postgresql-13-pg-later_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-13-pg-later_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-later` | 0.3.0 | `u24.aarch64` | pigsty | 1.1 MiB | [postgresql-13-pg-later_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-13-pg-later_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ChuckHend/pg_later" title="Repository" icon="github" subtitle="github.com/ChuckHend/pg_later" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_later-0.3.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_later; # get pg_later source code
pig build dep pg_later; # install build dependencies
pig build pkg pg_later; # build extension rpm or deb
pig build ext pg_later; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_later; # install by extension name, for the current active PG version
pig ext install pg_later; # install via package alias, for the active PG version
pig ext install pg_later -v 17;   # install for PG 17
pig ext install pg_later -v 16;   # install for PG 16
pig ext install pg_later -v 15;   # install for PG 15
pig ext install pg_later -v 14;   # install for PG 14
pig ext install pg_later -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_later CASCADE SCHEMA pglater;
```

